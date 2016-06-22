package midwares

import (
	"github.com/gin-gonic/gin"
	"xtudouh/common/conf"
)

var securityKey = ""
func SecurityAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get("Security-Key")
		if key == securityKey {
			c.Set("Security", key)
		}
		c.Next()
	}
}

func init() {
	registerInitFun(func() {
		securityKey = conf.String("auth", "KEYSECRET")
	})
}
