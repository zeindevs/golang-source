package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type JobStatus string

const (
	StatusPending JobStatus = "pending"
	StatusRunning JobStatus = "running"
	StatusDone    JobStatus = "done"
	StatusFailed  JobStatus = "failed"
)

type Job struct {
	ID        int64
	Type      string
	Payload   string
	Status    JobStatus
	RunAt     time.Time
	Retry     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobHandler func(payload map[string]any) error

type Queue struct {
	DB       *sql.DB
	Workers  int
	Handlers map[string]JobHandler
	mu       sync.Mutex
	wg       sync.WaitGroup
	stopCh   chan struct{}
}

func NewQueue(dbPath string, workers int) (*Queue, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	db.Exec(`PRAGMA journal_mode=WAL;`)
	db.Exec(`PRAGMA busy_timeout=5000;`)
	schema := `
		CREATE TABLE IF NOT EXISTS jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT,
			payload TEXT,
			status TEXT,
			retry INTEGER DEFAULT 0,
			run_at DATETIME,
			created_at DATETIME,
			updated_at DATETIME
		);
		CREATE INDEX IF NOT EXISTS idx_jobs_status_runat ON jobs (status, run_at);
	`
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}
	return &Queue{
		DB:       db,
		Workers:  workers,
		Handlers: make(map[string]JobHandler),
		stopCh:   make(chan struct{}),
	}, nil
}

func (q *Queue) RegisterHandler(jobType string, handler JobHandler) {
	q.Handlers[jobType] = handler
}

func (q *Queue) Enqueue(jobType string, payload any, runAt time.Time) error {
	data, _ := json.Marshal(payload)
	now := time.Now()
	_, err := q.DB.Exec(`
		INSERT INTO jobs (type, payload, status, run_at, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)`, jobType, string(data), StatusPending, runAt, now, now)
	return err
}

func (q *Queue) fetchNextJob() (*Job, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	tx, err := q.DB.Begin()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(`
		SELECT id, type, payload, status, run_at, retry, created_at, updated_at 
		FROM jobs 
		WHERE status = ? AND run_at <= ? 
		ORDER BY run_at ASC
		LIMIT 1`, StatusPending, time.Now())
	var j Job
	if err := row.Scan(&j.ID, &j.Type, &j.Payload, &j.Status, &j.RunAt, &j.Retry, &j.CreatedAt, &j.UpdatedAt); err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.Exec("UPDATE jobs SET status = ?, updated_at = ? WHERE id = ?", StatusRunning, time.Now(), j.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &j, nil
}

func (q *Queue) processJob(job *Job) {
	handler, ok := q.Handlers[job.Type]
	if !ok {
		log.Printf("[Worker] No handler for job type %s", job.Type)
		q.markFailed(job, errors.New("no handler"))
		return
	}
	var payload map[string]any
	if err := json.Unmarshal([]byte(job.Payload), &payload); err != nil {
		q.markFailed(job, err)
		return
	}
	log.Printf("[Worker] Running job #%d type=%s", job.ID, job.Type)
	err := handler(payload)
	if err != nil {
		log.Printf("[Worker] Job #%d failed: %s", job.ID, err)
		q.retryOrFail(job)
		return
	}
	q.markDone(job)
}

func (q *Queue) markDone(job *Job) {
	_, _ = q.DB.Exec(`UPDATE jobs SET status = ?, updated_at = ? WHERE id = ?`, StatusDone, time.Now(), job.ID)
}

func (q *Queue) markFailed(job *Job, err error) {
	log.Printf("[Worker] Job #%d failed permanently: %v", job.ID, err)
	_, _ = q.DB.Exec(`UPDATE jobs SET status = ?, updated_at = ? WHERE id = ?`, StatusFailed, time.Now(), job.ID)
}

func (q *Queue) retryOrFail(job *Job) {
	if job.Retry > 3 {
		now := time.Now()
		_, _ = q.DB.Exec(`
			UPDATE jobs SET status = ?, retry = retry + 1, run_at = DATETIME(?, '+5 seconds'), updated_at = ?
			WHERE id = ?`, StatusPending, now, now, job.ID)
	} else {
		q.markFailed(job, errors.New("max retry exceeded"))
	}
}

func (q *Queue) Start() {
	for i := range q.Workers {
		q.wg.Add(1)
		go func(id int) {
			defer q.wg.Done()
			for {
				select {
				case <-q.stopCh:
					log.Printf("[Worker] Stopping worker #%d", id)
					return
				default:
					job, err := q.fetchNextJob()
					if err != nil {
						if !errors.Is(err, sql.ErrNoRows) {
							fmt.Printf("🔴 Error fetching next job: %v\n", err)
						}
						time.Sleep(1 * time.Second)
						continue
					}
					q.processJob(job)
				}
			}
		}(i)
	}
}

func (q *Queue) Stop() {
	close(q.stopCh)
	q.wg.Wait()
}

func main() {
	queue, err := NewQueue("jobs.db", 2)
	if err != nil {
		panic(err)
	}

	queue.RegisterHandler("email", func(payload map[string]any) error {
		fmt.Printf("📧 Sending email to %v with subject %v\n", payload["to"], payload["subject"])
		time.Sleep(1 * time.Second)
		return nil
	})

	now := time.Now()
	queue.Enqueue("email", map[string]any{"to": "john@example.com", "subject": "Hello"}, now)
	queue.Enqueue("email", map[string]any{"to": "jane@example.com", "subject": "Hello"}, now.Add(5*time.Second))

	queue.Start()

	time.Sleep(60 * time.Second)
	queue.Stop()
}
