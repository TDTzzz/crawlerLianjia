package main

import (
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"github.com/TDTzzz/crawlerLianjia/parser"
	"github.com/TDTzzz/crawlerLianjia/persist"
	"github.com/TDTzzz/crawlerLianjia/schedular"
)

func main() {
	//1.建立ElasticSearch连接
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	//创建Engine（引擎）
	e := engine.ConcurrentEngine{
		Scheduler:   &schedular.QueuedScheduler{},
		WorkerCount: 13,
		ItemChan:    itemChan,
	}

	//种子URL
	req := engine.Request{
		Url:       "https://wh.lianjia.com/ershoufang/",
		ParseFunc: parser.RegionList,
	}

	//启动Engine引擎
	e.Run(req)
	//engine.SimpleEngine{}.Run(req) //单机版爬虫
}
