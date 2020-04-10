package control

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/olivere/elastic/v7"
	"sort"
	"strconv"
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

func (h SearchResultHandler) AvgPriceSearch(name string, value string, st string, ed string) map[string]PriceResults {
	boolQuery := commonBoolQuery(name, value, st, ed)
	regionAgg := commonAggQuery()
	dateAgg := elastic.NewDateHistogramAggregation().Field("Date").
		CalendarInterval("day").Format("yyyy-MM-dd").SubAggregation("a", regionAgg)

	res, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).Size(0).
		Aggregation("by_day", dateAgg).Do(context.Background())
	data, _ := res.Aggregations["by_day"].MarshalJSON()

	var dat map[string]interface{}
	json.Unmarshal(data, &dat)
	return parsePrice(dat)
}

func commonBoolQuery(name string, value string, st string, ed string) *elastic.BoolQuery {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery(name, value))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gte(st))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Lte(ed))
	return boolQuery
}

func commonAggQuery() *elastic.TermsAggregation {
	regionAgg := elastic.NewTermsAggregation().Field("SubRegion.keyword").Size(100).
		SubAggregation("sumPrice", elastic.NewSumAggregation().Field("TotalPrice")).
		SubAggregation("sumArea", elastic.NewSumAggregation().Field("Area"))
	return regionAgg
}

//[]byte解析 -> struct
func parsePrice(raw map[string]interface{}) map[string]PriceResults {
	var res = make(map[string]PriceResults)

	for _, v := range raw["buckets"].([]interface{}) {
		tmp := v.(map[string]interface{})
		date := tmp["key_as_string"].(string)
		subBuckets := tmp["a"].(map[string]interface{})
		for _, v2 := range subBuckets {
			switch vv := v2.(type) {
			case []interface{}:
				for _, u := range vv {
					keyRes := u.(map[string]interface{})
					sumArea := keyRes["sumArea"].(map[string]interface{})["value"].(float64)
					SumPrice := keyRes["sumPrice"].(map[string]interface{})["value"].(float64)
					avgPrice, err := strconv.ParseFloat(fmt.Sprintf("%.2f", SumPrice*10000/sumArea), 64)
					if err != nil {
						avgPrice = float64(0)
					}
					res[date] = append(res[date], PriceRes{
						Cnt:      int(keyRes["doc_count"].(float64)),
						Key:      keyRes["key"].(string),
						AvgPrice: avgPrice,
						Date:     date,
					})
				}
			}
		}
		sort.Sort(res[date])
	}
	return res
}

type PriceRes struct {
	Date     string
	Cnt      int
	Key      string
	AvgPrice float64
}

type PriceResults []PriceRes

func (p PriceResults) Len() int {
	return len(p)
}

func (p PriceResults) Less(i, j int) bool {
	return p[i].AvgPrice > p[j].AvgPrice
}

func (p PriceResults) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
