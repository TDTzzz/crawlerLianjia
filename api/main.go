package main

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/avg_price/:region", func(c *gin.Context) {
		region := c.Param("region")
		searchClient := control.CreateSearchResHandler()
		res := searchClient.AvgPriceSearch(region, "2020-04-07", "2020-04-08")
		c.JSON(http.StatusOK, res)
	})
	router.Run()
}
