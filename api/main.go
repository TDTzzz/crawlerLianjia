package main

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"log"
)

func main() {


	res := control.CreateSearchResHandler().AvgPriceSearchV2("Region.keyword", "武昌", "2020-04-05", "2020-04-10")
	log.Println(res)
}
