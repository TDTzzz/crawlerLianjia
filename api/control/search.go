package control

import (
	"github.com/olivere/elastic/v7"
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

func CommonBoolQuery(name string, value string, st string, ed string) *elastic.BoolQuery {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery(name, value))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gte(st))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Lte(ed))
	return boolQuery
}
