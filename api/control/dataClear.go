package control

import (
	"context"
	"encoding/json"
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/olivere/elastic/v7"
	"log"
)

//清楚重复的数据
func (h SearchResultHandler) RepeatData(region string, date string) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("Region.keyword", region))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Gte(date))
	boolQuery.Filter(elastic.NewRangeQuery("Date").Lte(date))
	topHitsAggs := elastic.NewTopHitsAggregation().Size(10)
	termsAggs := elastic.NewTermsAggregation().Field("Id").Size(10000).SubAggregation("b", topHitsAggs)
	res, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).Size(0).
		Aggregation("a", termsAggs).Do(context.Background())
	data, _ := res.Aggregations["a"].MarshalJSON()
	log.Println("总数：", res.TotalHits())
	var dat map[string]interface{}
	json.Unmarshal(data, &dat)
	//取出
	h.parseRepeatData(dat["buckets"])
}

func (h SearchResultHandler) parseRepeatData(buckets interface{}) {
	aa := buckets.([]interface{})
	for _, v := range aa {
		tmp := v.(map[string]interface{})
		cnt := int(tmp["doc_count"].(float64))
		if cnt == 1 {
			//log.Println("cnt==1 okk！！")
			continue
		}
		needDelCnt := cnt - 1
		hitData := tmp["b"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})
		for i := 0; i < needDelCnt; i++ {
			delId := hitData[i].(map[string]interface{})["_id"].(string)
			h.DelById(delId)
		}
	}
}

func (h SearchResultHandler) DelById(id string) {
	res, err := h.client.Delete().Index(config.ElasticIndex).Id(id).Do(context.Background())
	if err != nil {
		log.Println("del err:", err)
	}
	if res.Result == "deleted" {
		log.Println("del success +1!!!")
	} else {
		log.Println("del res ???")
	}
}
