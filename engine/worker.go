package engine

import "github.com/TDTzzz/crawlerLianjia/fetcher"

func Worker(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParseFunc(contents), nil
}
