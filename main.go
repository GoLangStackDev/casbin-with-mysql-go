package main

import (
	"casbin-with-mysql-go/mid"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.New()
	// 加入中间件
	r.Use(mid.CheekLogin(),mid.RBAC())

	r.GET("/posts", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"msg": "所有文章",
		})
	})
	r.POST("/posts", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"msg": "批量修改文章",
		})
	})
	if err:=r.Run(":8082");err!=nil {
		log.Panicln(err)
	}
}
// 统一错误检查
func checkError(err error)  {
	if err!=nil {
		println(err.Error())
	}
}
