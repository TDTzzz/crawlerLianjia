package control

import (
	"context"
	"encoding/json"
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"net/http"
)

func RegionInfo(c *gin.Context) {
	data := CreateSearchResHandler().Region()
	c.JSON(http.StatusOK, data)
}

//获得区域信息
func (h SearchResultHandler) Region() []string {
	termsAggs := elastic.NewTermsAggregation().Field("Region.keyword")
	res, _ := h.client.Search().Index(config.ElasticIndex).Aggregation("regions", termsAggs).Do(context.Background())
	data, _ := res.Aggregations["regions"].MarshalJSON()
	var dat map[string]interface{}
	json.Unmarshal(data, &dat)
	var regions []string
	for _, v := range dat["buckets"].([]interface{}) {
		//log.Println(v)
		tmp := v.(map[string]interface{})
		regions = append(regions, tmp["key"].(string))
	}
	return regions
}

//获得区域下的subRegion信息
func (h SearchResultHandler) SubRegion(region string) []string {
	boolQuery := elastic.NewBoolQuery().Must(elastic.NewTermQuery("Region.keyword", region))

	termsAggs := elastic.NewTermsAggregation().Field("SubRegion.keyword")
	res, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).
		Aggregation("subRegions", termsAggs).Do(context.Background())
	data, _ := res.Aggregations["subRegions"].MarshalJSON()
	var dat map[string]interface{}
	json.Unmarshal(data, &dat)
	var regions []string
	for _, v := range dat["buckets"].([]interface{}) {
		//log.Println(v)
		tmp := v.(map[string]interface{})
		regions = append(regions, tmp["key"].(string))
	}
	return regions
}
