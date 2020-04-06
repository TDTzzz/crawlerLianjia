package main

import (
	"github.com/TDTzzz/crawlerLianjia/api/control"
	"log"
)

func main() {
	//router := gin.Default()
	//router.GET("/gua", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"msg": "aa",
	//	})
	//})
	//router.Run()

	searchClient := control.CreateSearchResHandler()
	res, _ := searchClient.GetSearchRes()
	log.Print(res)
}
