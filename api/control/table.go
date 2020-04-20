package control

import (
	"context"
	"encoding/json"
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SubRegionTable(c *gin.Context) {
	name := c.Param("name")
	res := CreateSearchResHandler().SubRegionHouseTable(name, "2020-04-19")
	c.JSON(http.StatusOK, res)
}

func (h SearchResultHandler) SubRegionHouseTable(subRegion string, date string) []map[string]interface{} {
	boolQuery := CommonBoolQuery("SubRegion.keyword", subRegion, date, date)
	data, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).Sort("UnitPrice", false).
		Do(context.Background())
	log.Print(data.TotalHits())

	var res []map[string]interface{}
	for _, hit := range data.Hits.Hits {
		var dat map[string]interface{}
		json.Unmarshal(hit.Source, &dat)
		res = append(res, dat)
	}
	return res
}
