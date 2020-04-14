package main

import (
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"github.com/TDTzzz/crawlerLianjia/parser"
	"github.com/TDTzzz/crawlerLianjia/persist"
	"github.com/TDTzzz/crawlerLianjia/schedular"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &schedular.QueuedScheduler{},
		WorkerCount: 12,
		ItemChan:    itemChan,
	}

	req := engine.Request{
		Url:       "https://wh.lianjia.com/ershoufang/",
		ParseFunc: parser.RegionList,
	}

	e.Run(req)
	//engine.SimpleEngine{}.Run(req) //单机版爬虫
}
