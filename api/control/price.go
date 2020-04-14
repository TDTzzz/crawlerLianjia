//价格相关的api
package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegionPrice(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	method := c.Request.Method
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()

	name := c.Param("name")
	ed := time.Now().Format("2006-01-02")
	dd, _ := time.ParseDuration("-168h")
	st := time.Now().Add(dd).Format("2006-01-02")
	data := CreateSearchResHandler().AvgPriceSearch("Region.keyword", name, st, ed)

	//处理下数据
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
	c.JSON(http.StatusOK, res)
}

func SubRegionPrice(c *gin.Context) {

	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()

	name := c.Param("name")
	ed := time.Now().Format("2006-01-02")
	dd, _ := time.ParseDuration("-168h")
	st := time.Now().Add(dd).Format("2006-01-02")
	res := CreateSearchResHandler().AvgPriceSearch("SubRegion.keyword", name, st, ed)
	c.JSON(http.StatusOK, res)
}

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
