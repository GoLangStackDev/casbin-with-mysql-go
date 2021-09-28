package mid

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 校验是否登录
func CheekLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("token")=="" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "token required",
			})
		}else{
			c.Set("user_name", c.Request.Header.Get("token"))
			c.Next()
		}
	}
}

func RBAC() gin.HandlerFunc {
	e,_ := casbin.NewEnforcer("resources/model.conf","resources/policy.csv")

	return func(c *gin.Context) {
		user,_ := c.Get("user_name")
		if has,err:=e.Enforce(user,c.Request.RequestURI,c.Request.Method);err!=nil||!has {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"msg": "forbidden",
			})
		}else{
			c.Next()
		}
	}
}