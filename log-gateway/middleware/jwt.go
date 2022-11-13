package middleware

import (
	"time"

	"log_gateway/utils"
	"log_gateway/utils/errmsg"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = errmsg.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errmsg.ErrorAuthCheckTokenTimeout
			}
		}
		if code != errmsg.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    errmsg.GetMsg(uint(code)),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
