package main

import (
	"github.com/TDTzzz/crawlerLianjia/router"
)

func main() {

	//res := control.CreateSearchResHandler().AvgPriceSearchV2("Region.keyword", "武昌", "2020-04-05", "2020-04-13")
	//log.Println(res)

	//切片
	//var (
	//	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
	//		"delta": 87, "echo": 56, "foxtrot": 12,
	//		"golf": 34, "hotel": 16, "indio": 87,
	//		"juliet": 65, "kili": 43, "lima": 98}
	//)
	//
	//keys := make([]string, len(barVal))
	//i := 0
	//for k, _ := range barVal {
	//	keys[i] = k
	//	i++
	//}
	//sort.Strings(keys)
	//
	//for k, v := range keys {
	//	log.Print(k, v)
	//}
	router.InitRouter()
}
