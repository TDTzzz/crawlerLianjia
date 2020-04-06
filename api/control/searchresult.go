package control

import (
	"context"
	"github.com/TDTzzz/crawlerLianjia/config"
	model2 "github.com/TDTzzz/crawlerLianjia/model"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

type SearchResultHandler struct {
	client *elastic.Client
}

func CreateSearchResHandler() SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{client: client}
}

func (h SearchResultHandler) GetSearchRes() (interface{}, error) {

	//resp := h.Search()
	resp := h.AggsSearch()
	total := resp.TotalHits()
	for _, item := range resp.Each(reflect.TypeOf(model2.HouseDetail{})) {
		log.Print(item)
	}

	return total, nil
}

func (h SearchResultHandler) Search() *elastic.SearchResult {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewMatchQuery("SubRegion", "中南丁字"))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gt("2020-04-04"))
	searchRes, _ := h.client.Search(config.ElasticIndex).Query(boolQuery).Do(context.Background())
	return searchRes
}

func (h SearchResultHandler) AggsSearch() *elastic.SearchResult {
	minAgg := elastic.NewMinAggregation().Field("unitPrice")
	build := h.client.Search(config.ElasticIndex).Pretty(true)

	minRes, _ := build.Aggregation("aggsRegion", minAgg).Do(context.Background())
	return minRes
}
