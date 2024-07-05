package middleware

import (
	"net/http"

	"gin-rest-api/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		token := c.GetHeader("x-token")
		_, err := util.ParseToken(token)

		var msg string

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				msg = "Token已过期"
			default:
				msg = "Token错误"
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":     401,
				"msg":      msg,
				"data":     data,
				"trace_id": c.GetString("trace_id"),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
