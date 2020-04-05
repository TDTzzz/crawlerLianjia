package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"log"
	"strings"
)

var host = "https://wh.lianjia.com"

func RegionList(contents []byte) engine.ParseResult {
	var res engine.ParseResult
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}

	dom.Find("div.position>dl:nth-child(2)>dd>div[data-role=ershoufang]>div>a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		regionName := selection.Text()
		res.Requests = append(res.Requests, engine.Request{
			Url:       host + url,
			ParseFunc: SubRegionList,
		})
		log.Printf(regionName)
	})
	return res
}
