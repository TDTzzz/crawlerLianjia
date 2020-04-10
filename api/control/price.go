//价格相关的api
package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegionPrice(c *gin.Context) {
	name := c.Param("name")
	ed := time.Now().Format("2006-01-02")
	dd, _ := time.ParseDuration("-168h")
	st := time.Now().Add(dd).Format("2006-01-02")
	res := CreateSearchResHandler().AvgPriceSearch("Region.keyword", name, st, ed)
	c.JSON(http.StatusOK, res)
}

func SubRegionPrice(c *gin.Context) {
	name := c.Param("name")
	ed := time.Now().Format("2006-01-02")
	dd, _ := time.ParseDuration("-168h")
	st := time.Now().Add(dd).Format("2006-01-02")
	res := CreateSearchResHandler().AvgPriceSearch("SubRegion.keyword", name, st, ed)
	c.JSON(http.StatusOK, res)
}

func CommunityPrice(c *gin.Context) {

}
