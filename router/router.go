package router

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"github.com/TDTzzz/crawlerLianjia/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	//中间件
	r.Use(middleware.Options)
	//房价
	price := r.Group("/price")
	{
		price.GET("/subregion/:name", control.SubRegionPrice)
		price.GET("/region/:name", control.RegionPrice)
		price.GET("/community/:name", control.CommunityPrice)
	}
	//房源数据
	table := r.Group("/table")
	{
		table.GET("/subregion/:name", control.SubRegionTable)
	}

	//区域信息
	region := r.Group("/region")
	{
		region.GET("/info",control.RegionInfo)
	}
	r.Run(":8081")
	return r
}
