package main

import (
    "fmt"
    "xtudouh/common/conf"
    "xtudouh/web/controllers"
    "github.com/gin-gonic/gin"
    "xtudouh/web/sessions"
    "xtudouh/web/midwares"
    "xtudouh/web/db"
    "xtudouh/common/redis"
    "xtudouh/common/idg"
    "xtudouh/common/log"
    "xtudouh/common/wechat"
)

func start() {
    app := gin.New()
    gin.SetMode(conf.ENV)
    appContext := fmt.Sprintf("/%s", conf.AppName)
    rootGroup := app.Group(appContext)
    version := fmt.Sprintf("/%s", conf.AppVer)
    verGroup := rootGroup.Group(version)
    verGroup.Use(
        midwares.LoggerHandler(),
        midwares.ErrorHandler(),
        midwares.Cros(),
        sessions.Sessions())

    controllers.RegisterHandlers(verGroup)
    app.Run(conf.ListeningPort)
}

func init() {
    conf.Init("app.ini")
    log.Init()
    idg.Init()
    redis.Init()
    sessions.Init()
    midwares.Init()
    db.Init()
    wechat.Init()
    fmt.Println("Init successfully.")
}

func main() {
    start()
}