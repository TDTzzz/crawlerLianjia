package engine

import (
	"github.com/TDTzzz/crawlerLianjia/fetcher"
	"log"
	"strings"
	"time"
)

func Worker(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	str := string(contents)
	if strings.Contains(str, "人机认证") {
		//给个延迟，重启这个ParFunc
		log.Println("哦豁了：" + r.Url)
		time.Sleep(10 * time.Second)
		return Worker(r)
	}
	return r.ParseFunc(contents), nil
}
