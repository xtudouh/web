package activity

import (
    "github.com/gin-gonic/gin"
    "xtudouh/common/wechat"
    "xtudouh/web/services/activityservice"
    "xtudouh/web/domain"
    "time"
    "fmt"
    "xtudouh/common/utils"
    "crypto/sha1"
    "xtudouh/common/log"
)

var l = log.NewLogger()

type signForm struct {
    ActivityId int64 `form:"activityid" json"activityid"`
    CustName   string `form:"custname" json:"custname"`
    Mobile     string   `form:"mobile" json:"mobile"`
}

func SignIn(c *gin.Context) {
    var form signForm
    if err := c.Bind(&form); err != nil {
        c.JSON(500, nil)
        return
    }
    sign := &domain.SignIn{
        ActivityId:form.ActivityId,
        Custname: form.CustName,
        Mobile: form.Mobile,
    }
    ret, err := activityservice.SignIn(sign)
    if err != nil {
        c.JSON(500, err)
        return
    }
    if ret == domain.SIGNED_IN {
        c.JSON(200, domain.SIGNED_IN)
        return
    }

    c.JSON(200, domain.OK)
}

type address struct {
    Latitude  float64   `json:"latitude" form:"latitude"`
    Longitude float64   `json:"longitude" form:"longitude"`
}

func QueryActivity(c *gin.Context) {
    var addr address
    c.Bind(&addr)
    act, err := activityservice.QueryActivity(addr.Latitude, addr.Longitude)
    if err != nil {
        c.JSON(200, nil)
        return
    }
    c.JSON(200, act)
}

type initData struct {
    JsapiTicket string `json:"jsapi_ticket"`
    AppId       string `json:"appid"`
    Timestamp   int64  `json:"timestamp"`
    NonceStr    string `json:"nonceStr"`
    Signature   string `json:"signature"`
}


func InitData(c *gin.Context) {
    url := c.Query("url")
    l.Debug("location url: %s", url)
    initData := new(initData)
    initData.JsapiTicket = wechat.JsTicket
    initData.AppId = wechat.Appid
    initData.Timestamp = time.Now().Unix()
    initData.NonceStr = utils.RandomStr(32)
    toSignStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", initData.JsapiTicket, initData.NonceStr, initData.Timestamp, url)
    initData.Signature = fmt.Sprintf("%x", sha1.Sum([]byte(toSignStr)))
    c.JSON(200, initData)
}
