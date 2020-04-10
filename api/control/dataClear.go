package control

import "github.com/olivere/elastic/v7"

//清楚重复的数据
func (h SearchResultHandler) RepeatData(date string) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gte(date))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Lte(date))
}
