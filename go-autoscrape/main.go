package main

import (
	"context"
	"log"
	"time"

	"github.com/zeindevs/goautoscrape/config"
	"github.com/zeindevs/goautoscrape/internal/service"
)

type Worker[T any] struct {
	Name     string
	Interval time.Duration
	Func     func() ([]*T, error)
	StoreDB  func(ctx context.Context, data []*T) error
	QuitCh   chan bool
}

func (w *Worker[T]) Run() {
	log.Println("Worker", w.Name, "running...")
	ticker := time.NewTicker(w.Interval)
	ctx, cancel := context.WithCancel(context.Background())

	for {
		select {
		case <-w.QuitCh:
			log.Println("Quit")
			cancel()
			return
		case t := <-ticker.C:
			log.Println("Fetch data", t)

			start := time.Now()
			data, err := w.Func()
			if err != nil {
				log.Println(err.Error())
			}
			if err := w.StoreDB(ctx, data); err != nil {
				log.Println(err.Error())
			}

			log.Println("Worker", w.Name, "took:", time.Since(start).String())
		}
	}
}

func main() {
	cfg := config.NewConfig()
	db := config.ConnectDB(cfg)

	otakudesuScraper := service.NewOtakudesuScraper()
	otakudesuService := service.NewOtakudesuService(db)

	quitCh1 := make(chan bool)
	worker1 := Worker[service.Otakudesu]{
		Name:     "Otakudesu Ongoing",
		Interval: time.Second * 30,
		Func:     otakudesuScraper.GetOngoingAll,
		StoreDB:  otakudesuService.Save,
		QuitCh:   quitCh1,
	}

 //  quitCh2 := make(chan bool)
	// worker2 := Worker[service.Otakudesu]{
	// 	Name:     "Otakudesu Complete",
	// 	Interval: time.Second * 30,
	// 	Func:     otakudesuScraper.GetCompleteAll,
	// 	StoreDB:  otakudesuService.Save,
	// 	QuitCh:   quitCh2,
	// }

	go worker1.Run()
	// go worker2.Run()

	time.Sleep(50 * time.Second)
	quitCh1 <- true
	// quitCh2 <- true
	time.Sleep(2 * time.Second)

	log.Println("Exiting")
}
