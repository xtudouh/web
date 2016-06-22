package nsq

import (
	"github.com/gin-gonic/gin"
	"xtudouh/common/nsq"
)

func Pm(c *gin.Context) {
	nsq.TextMsg("what?").Push()
}
