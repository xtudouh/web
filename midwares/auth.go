package midwares

import (
    "github.com/gin-gonic/gin"
    "xtudouh/web/sessions"
    "xtudouh/common/conf"
    "strings"
)

var (
    skip_urls [][]string
)


func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if isSkipUrl(c.Request.URL.Path, c.Request.Method) {
            c.Next()
            return
        }
        user := sessions.Get("user", c)
        if user == nil {
            c.AbortWithStatus(401)
            return
        }
        c.Next()
    }
}

func isSkipUrl(urlStr, method string) bool {
    for _, urlAndMethod := range skip_urls {
        if urlStr == urlAndMethod[1] && method == urlAndMethod[0] {
            return true
        }
    }

    return false
}


func init() {
    registerInitFun(func() {
        urls := conf.Strings("auth", "SKIP_URLS", ",")
        for _, urlStr := range urls {
            skip_urls = append(skip_urls, strings.Fields(urlStr))
        }
    })
}
