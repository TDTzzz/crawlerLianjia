//价格相关的api
package control

import (
	"context"
	"encoding/json"
	"github.com/TDTzzz/crawlerLianjia/config"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"net/http"
	"sort"
	"time"
)

func RegionPrice(c *gin.Context) {
	name := c.Param("name")
	st, ed := getStEd()
	data := CreateSearchResHandler().AvgPriceSearchV2("Region.keyword", name, st, ed)
	res := formatData(data)
	c.JSON(http.StatusOK, res)
}

func SubRegionPrice(c *gin.Context) {
	name := c.Param("name")
	st, ed := getStEd()
	data := CreateSearchResHandler().AvgPriceSearchV2("SubRegion.keyword", name, st, ed)
	res := formatData(data)
	c.JSON(http.StatusOK, res)
}

//小区价格
func CommunityPrice(c *gin.Context) {

}

func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//格式转换
func formatData(data map[string]PriceResults) [2]interface{} {
	var rows []map[string]interface{}
	var columns []string
	columns = append(columns, "日期")
	for _, v := range data {
		var tmp = make(map[string]interface{})
		for _, vv := range v {
			tmp["日期"] = vv.Date
			columns = append(columns, vv.Key)
			tmp[vv.Key] = vv.AvgPrice
		}
		rows = append(rows, tmp)
	}
	columns = removeDuplicateElement(columns)
	res := [2]interface{}{0: columns, 1: rows}
	return res
}

func getStEd() (string, string) {
	ed := time.Now().Format("2006-01-02")
	dd, _ := time.ParseDuration("-168h")
	st := time.Now().Add(dd).Format("2006-01-02")
	return st, ed
}

//用pipeline Aggregation算各区域的平均价格日期直方图
func (h SearchResultHandler) AvgPriceSearchV2(name string, value string, st string, ed string) map[string]PriceResults {
	boolQuery := CommonBoolQuery(name, value, st, ed)

	pipelineScript := elastic.NewScript("params.A/params.B*10000")
	pipeAgg := elastic.NewBucketScriptAggregation().AddBucketsPath("A", "sumPrice").
		AddBucketsPath("B", "sumArea").Script(pipelineScript)

	termsAgg := elastic.NewTermsAggregation().Field("SubRegion.keyword").
		SubAggregation("sumPrice", elastic.NewSumAggregation().Field("TotalPrice")).
		SubAggregation("sumArea", elastic.NewSumAggregation().Field("Area")).
		SubAggregation("avgPrice", pipeAgg).Size(100)

	dateAgg := elastic.NewDateHistogramAggregation().
		Field("Date").CalendarInterval("day").Format("yyyy-MM-dd").
		SubAggregation("regions", termsAgg).
		SubAggregation("sumPrice", elastic.NewSumAggregation().Field("TotalPrice")).
		SubAggregation("sumArea", elastic.NewSumAggregation().Field("Area")).
		SubAggregation("avgPrice", pipeAgg)

	data, _ := h.client.Search().Index(config.ElasticIndex).Query(boolQuery).Size(0).
		Aggregation("per_day", dateAgg).Do(context.Background())
	var dat map[string]interface{}
	aa, _ := data.Aggregations["per_day"].MarshalJSON()
	json.Unmarshal(aa, &dat)
	return parsePriceV2(dat)
}

func parsePriceV2(raw map[string]interface{}) map[string]PriceResults {
	var res = make(map[string]PriceResults)
	data := raw["buckets"]
	for _, v := range data.([]interface{}) {
		tmp := v.(map[string]interface{})
		date := tmp["key_as_string"].(string)
		regions := tmp["regions"].(map[string]interface{})
		for _, vv := range regions["buckets"].([]interface{}) {
			tmp2 := vv.(map[string]interface{})
			res[date] = append(res[date], PriceRes{
				Date:     date,
				Cnt:      int(tmp2["doc_count"].(float64)),
				Key:      tmp2["key"].(string),
				AvgPrice: tmp2["avgPrice"].(map[string]interface{})["value"].(float64),
			})
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
