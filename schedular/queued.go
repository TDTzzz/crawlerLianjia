package schedular

import (
	"github.com/TDTzzz/crawlerLianjia/engine"
)

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (s *QueuedScheduler) NewWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeReq engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeReq = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-s.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeReq: //这里把req分发给你worker
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
