package engine

import (
	"github.com/TDTzzz/crawlerLianjia/fetcher"
	"log"
)

//单机办爬虫引擎
type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {

	log.Print("---")
	for _, v := range seeds {
		res := worker(v)
		for _, v2 := range res.Requests {
			res2 := worker(v2)
			for _, v3 := range res2.Requests {
				res3 := worker(v3)
				for _, v4 := range res3.Items {
					log.Print(v4)
				}
			}
		}
	}
}

func worker(r Request) ParseResult {
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		panic(err)
	}
	res := r.ParseFunc(contents)
	return res
}
