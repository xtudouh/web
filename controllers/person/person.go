package person
import (
    "github.com/gin-gonic/gin"
    "xtudouh/common/log"
)

var l = log.NewLogger()

type demo struct {
    Id int64 `json:"id,string"`
}


func Hello(c *gin.Context) {
    d := demo{Id:689696318607855616}
    c.JSON(200, d)
}