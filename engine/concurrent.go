package engine

import (
	"github.com/TDTzzz/crawlerLianjia/model"
	"log"
)

//并发版爬虫引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan model.HouseDetail
}

type Scheduler interface {
	Submit(Request)
	Run()
	NewWorkerChan() chan Request
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//开启scheduler队列
	out := make(chan ParseResult)
	e.Scheduler.Run()
	//开启worker
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.NewWorkerChan(), out, e.Scheduler)
	}
	for _, v := range seeds {
		e.Scheduler.Submit(v)
	}

	for {
		res := <-out
		for _, item := range res.Items {
			e.ItemChan <- item
		}
		for _, v := range res.Requests {
			e.Scheduler.Submit(v)
		}
	}
}

func (ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			req := <-in
			res, err := Worker(req)
			if err != nil {
				log.Println("creatWorker Err %v", err)
				continue
			}
			out <- res
		}
	}()
}
