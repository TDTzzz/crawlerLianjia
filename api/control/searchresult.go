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

func (h SearchResultHandler) AvgPriceSearch(region string, st string, ed string) PriceResults {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("Region.keyword", region))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gte(st))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Lte(ed))
	
	sumPrice := elastic.NewSumAggregation().Field("TotalPrice")
	sumArea := elastic.NewSumAggregation().Field("Area")
	regionAgg := elastic.NewTermsAggregation().Field("SubRegion.keyword").Size(20).
		SubAggregation("sumPrice", sumPrice).SubAggregation("sumArea", sumArea)

	res, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).Size(0).Aggregation("a", regionAgg).Do(context.Background())
	data, _ := res.Aggregations["a"].MarshalJSON()
	var dat CommonBucket
	var priceRes PriceResults

	json.Unmarshal(data, &dat)
	for _, v := range dat.Buckets {
		avgPrice, err := strconv.ParseFloat(fmt.Sprintf("%.2f", v.SumPrice.Val*10000/v.SumArea.Val), 64)
		if err != nil {
			avgPrice = float64(0)
		}
		priceRes = append(priceRes, PriceRes{
			Key:      v.Key,
			Cnt:      v.Cnt,
			AvgPrice: avgPrice,
		})
	}

	sort.Sort(priceRes)
	return priceRes
}

type CommonBucket struct {
	SumCnt  int64          `json:"sum_other_doc_count"`
	DocCnt  int64          `json:"doc_count_error_upper_bound"`
	Buckets []PriceBuckets `json:"buckets"`
}

type PriceBuckets struct {
	Key      string      `json:"key"`
	Cnt      int64       `json:"doc_count"`
	SumPrice SubAggValue `json:"sumPrice"`
	SumArea  SubAggValue `json:"sumArea"`
}

type SubAggValue struct {
	Val float64 `json:"value"`
}

type PriceRes struct {
	Key      string
	Cnt      int64
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
