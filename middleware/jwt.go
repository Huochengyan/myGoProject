package jwt

import (
	"github.com/gin-gonic/gin"
	"myGoProjectNew/pkg/e"
	"myGoProjectNew/pkg/util"
	"net/http"
	"strings"
	"time"
)

/**
  每一个请求前验证Token
*/
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")

		//获取header的token-===========================

		//authString := c.Request.Header.Get("Authorization")
		//kv := strings.Split(authString, " ")
		//if len(kv) != 2 || kv[0] != "Bearer" {
		//	code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		//	return
		//} else {
		//	token = kv[1]
		//}
		tokenstr := c.Request.Header.Get("token")
		if tokenstr != "" {
			token = tokenstr
			print(tokenstr)
		}
		//================================================

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			//上边的超时条件
			if err != nil {
				var timeout int
				timeout = strings.Index(err.Error(), "token is expired")
				if timeout != -1 {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
