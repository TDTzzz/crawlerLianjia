package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"log"
	"math"
	"strconv"
	"strings"
)

func PageList(contents []byte) engine.ParseResult {
	var res engine.ParseResult
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}
	totalStr := dom.Find("h2.total.fl>span").Text()
	totalStr = strings.Replace(totalStr, " ", "", -1)
	total, err := strconv.ParseFloat(totalStr, 64)
	log.Printf(totalStr)
	pg := int(math.Ceil(float64(total / 30)))
	//当前的url
	url, _ := dom.Find("div.position>dl:nth-child(2)>dd>div>div:nth-child(2)>a.selected").Attr("href")
	for i := 1; i <= pg; i++ {
		res.Requests = append(res.Requests, engine.Request{
			Url:       host + url + "pg" + strconv.Itoa(i),
			ParseFunc: HouseDetail,
		})
	}
	return res

}
