package engine

import "github.com/TDTzzz/crawlerLianjia/model"

type Request struct {
	Url       string
	ParseFunc func(contents []byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []model.HouseDetail
}
