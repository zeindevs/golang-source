package main

import (
	"context"
	"log"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

func main() {
	ctx := context.Background()

	workers := river.NewWorkers()

	river.AddWorker(workers, &SortWorker{})

	dbPool, err := pgxpool.New(ctx, "postgresql://postgres:root@localhost:5432/goriverqueue?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers: workers,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := riverClient.Stop(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := riverClient.Start(ctx); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := riverClient.Insert(ctx, SortArgs{
			Strings: []string{
				"whale", "tiger", "bear",
			},
		}, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Println("server listening on port :3000")
	http.ListenAndServe(":3000", r)
}

type SortArgs struct {
	Strings []string `json:"strings"`
}

func (SortArgs) Kind() string { return "sort" }

type SortWorker struct {
	river.WorkerDefaults[SortArgs]
}

func (w *SortWorker) Work(ctx context.Context, job *river.Job[SortArgs]) error {
	sort.Strings(job.Args.Strings)
	log.Printf("Sorted strings: %+v\n", job.Args.Strings)
	return nil
}
