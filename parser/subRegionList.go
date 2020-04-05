package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/TDTzzz/crawlerLianjia/engine"
	"log"
	"strings"
)

func SubRegionList(contents []byte) engine.ParseResult {
	var res engine.ParseResult
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		panic(err)
	}
	dom.Find("div[data-role=ershoufang]>div:nth-child(2)>a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		subRegionName := selection.Text()
		res.Requests = append(res.Requests, engine.Request{
			Url:       host + url,
			ParseFunc: PageList,
		})
		log.Printf(subRegionName)
	})
	return res
}
