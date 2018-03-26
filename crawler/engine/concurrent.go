package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		Scheduler.Submit(r)
	}
	in := make(chan Request)
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for {
		result := <-out
		for _,item:=range result.Items {
			fmt.Printf("%v",item)
		}
	}
}
