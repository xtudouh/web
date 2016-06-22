package activityservice

import (
    "math"
    "xtudouh/web/domain"
    "xtudouh/common/idg"
    "xtudouh/web/db"
    "xtudouh/common/log"
    "strings"
    "xtudouh/common/conf"
)

var l = log.NewLogger()

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
    radius := float64(6371000) // 6378137
    rad := math.Pi/180.0

    lat1 = lat1 * rad
    lng1 = lng1 * rad
    lat2 = lat2 * rad
    lng2 = lng2 * rad

    theta := lng2 - lng1
    dist := math.Acos(math.Sin(lat1) * math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2) * math.Cos(theta))

    return dist * radius
}

const select_activity = `select id, longitude, latitude, address activityname, brief from p_activity where date between date_trunc('day', now()) and date_trunc('day', now()) + interval '1 day'`

func filterByDistance(acts []domain.Activity, lat1, lng1 float64) *domain.Activity {
    var PRECISION = conf.Float64("distance", "PRECISION", 200.0)
    for _, act := range acts {
        if dis := EarthDistance(lat1, lng1, act.Latitude, act.Longitude); dis < PRECISION {
            l.Debug("Distance between (%f, %f) and (%f, %f) is %f", lat1, lng1, act.Latitude, act.Longitude, dis)
            return &act
        }
    }
    return nil
}

func QueryActivity(lat1, lng1 float64) (*domain.Activity, error) {
    var acts []domain.Activity
    if err := db.Engine.Sql(select_activity).Find(&acts); err != nil {
        l.Error("%v", err)
        return nil, err
    }

    return filterByDistance(acts, lat1, lng1), nil
}

func AddActivity(act *domain.Activity) error {
    if act.Id == 0 {
        act.Id, _ = idg.Id()
    }
    if _, err := db.Engine.Insert(act); err != nil {
        l.Error("%v", err)
        return err
    }
    return nil
}

func SignIn(signin *domain.SignIn) (string, error) {
    if signin.Id == 0 {
        signin.Id, _ = idg.Id()
    }
    if _, err := db.Engine.Insert(signin); err != nil {
        l.Error("%v", err)
        if strings.Contains(err.Error(), "unique") {
            return domain.SIGNED_IN, nil
        }
        return "", err
    }
    return domain.OK, nil
}
