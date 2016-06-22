package controllers
import (
    "github.com/gin-gonic/gin"
    "xtudouh/web/controllers/activity"
)


func RegisterHandlers(verGroup *gin.RouterGroup) {
    verGroup.GET("/init", activity.InitData)
    verGroup.POST("/activity", activity.QueryActivity)
    verGroup.POST("/signin", activity.SignIn)
}