package midwares
import (
    "time"
    "github.com/gin-gonic/gin"
    "xtudouh/common/conf"
    "strings"
)

func LoggerHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        end := time.Now()
        latency := end.Sub(start)
        clientIp := c.ClientIP()
        method := c.Request.Method
        statusCode := c.Writer.Status()
        appName := conf.AppName
        l.Info("[%s] %d  %v | %s %s %s", appName, statusCode, latency,
        clientIp, method, c.Request.URL.Path)
    }
}

func init() {
    registerInitFun(func() {
        urls := conf.Strings("auth", "SKIP_URLS", ",")
        for _, urlStr := range urls {
            skip_urls = append(skip_urls, strings.Fields(urlStr))
        }
    })
}