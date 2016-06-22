package main

import (
    "prosnav.com/common/conf"
    "prosnav.com/common/log"
    "prosnav.com/common/idg"
    "prosnav.com/web/db"
    "prosnav.com/web/services/activityservice"
    "prosnav.com/web/domain"
    "fmt"
    "time"
    "crypto/sha1"
    "io"
)

func main() {
    addActivity()
}

func sign() {
    h := sha1.New();
    str := "jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg&noncestr=Wm3WZYTPz0wzccnW&timestamp=1414587457&url=http://mp.weixin.qq.com"
    io.WriteString(h, str)
    fmt.Printf("%x", sha1.Sum([]byte(str)))
}

func queryActivity() {
    act, err := activityservice.QueryActivity(121.318353,31.212041)
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }

    fmt.Printf("%+v", act)
}

func addActivity() {
    act := domain.Activity{
        Latitude     : 121.320967,
        Longitude    : 31.201574,
        Address      : "上海市闵行区华漕镇阿里巴巴上海虹桥",
        ActivityName : "帆茂公开演讲",
        Brief        : "哈哈哈哈",
        Date         : time.Now(),
    }
    if err := activityservice.AddActivity(&act); err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Printf("Ok")
}

func signin() {
    sign := domain.SignIn{
        ActivityId: 741838044944142336,
        Custname: "老王",
        Mobile: "13111111111",
    }
    if _, err := activityservice.SignIn(&sign); err != nil {
        fmt.Printf(">>>>>>>>>>>>>%s<<<<<<<<<<<<<<", err.Error())
        return
    }
    fmt.Printf("Sign successfully!")
}

func init() {
    conf.Init("app.ini")
    log.Init()
    idg.Init()
    db.Init()
}