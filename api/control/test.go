package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	msg := "OK"
	c.String(http.StatusOK, msg)
}
