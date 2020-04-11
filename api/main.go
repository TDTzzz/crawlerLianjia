package main

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"log"
)

func main() {
	res := control.CreateSearchResHandler().SubRegion("武昌")
	log.Println(res)
}
