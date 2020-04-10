package router

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	//测试路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/test", control.Test)
	}

	//房价
	price := r.Group("/price")
	{
		price.GET("/subregion/:name", control.SubRegionPrice)
		price.GET("/region/:name", control.RegionPrice)
		price.GET("/community/:name", control.RegionPrice)
	}

	r.Run(":8080")
	return r
}
