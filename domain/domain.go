package domain

import "time"

type Activity struct {
    Id           int64 `xorm:"id pk" json:"id"`
    Latitude     float64   `xorm:"latitude" json:"-"`
    Longitude    float64   `xorm:"longitude" json:"-"`
    Address      string `xorm:"address" json:"address"`
    ActivityName string `xorm:"activityname" json:"activityname"`
    Brief        string `xorm:"brief" json:"brief"`
    Date         time.Time `xorm:"date timestamp" json:"-"`
    Crt          time.Time `xorm:"crt timestamp created" json:"-"`
    Lut          time.Time `xorm:"lut timestamp updated" json:"-"`
    Del          bool      `xorm:"del" json:"-"`
}

func (c *Activity) TableName() string {
    return "p_activity"
}

type SignIn struct {
    Id         int64 `xorm:"id pk"`
    ActivityId int64 `xorm:"activityid" json:"activityid"`
    Custname   string `xorm:"custname" json:"custname"`
    Mobile     string `xorm:"mobile" json:"mobile"`
    Crt        time.Time `xorm:"crt timestamp created"`
    Lut        time.Time `xorm:"lut timestamp updated"`
}

func (c *SignIn) TableName() string {
    return "p_signin"
}