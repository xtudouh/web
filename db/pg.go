package db

import (
    "database/sql"
    "github.com/go-xorm/xorm"
    _ "github.com/lib/pq"
    "xtudouh/common/conf"
    "xtudouh/common/log"
    "github.com/go-xorm/core"
)

var (
    Engine *xorm.Engine
    l = log.NewLogger()
)

type logAdapter struct  {}

func (da *logAdapter)Debug(v ...interface{}) (err error) {
    if len(v) > 1 {
        l.Debug(v[0].(string), v[1:]...)
        return
    }
    l.Debug(v[0].(string))
    return
}

func (da *logAdapter)Debugf(format string, v ...interface{}) (err error) {
    l.Debug(format, v...)
    return
}

func (da *logAdapter) Err(v ...interface{}) (err error) {
    if len(v) > 1 {
        l.Error(v[0].(string), v[1:]...)
        return
    }
    l.Error(v[0].(string))
    return
}
func (da *logAdapter) Errf(format string, v ...interface{}) (err error) {
    l.Error(format, v...)
    return
}
func (da *logAdapter) Info(v ...interface{}) (err error) {
    if len(v) > 1 {
        l.Info(v[0].(string), v[1:]...)
        return
    }
    l.Info(v[0].(string))
    return
}
func (da *logAdapter) Infof(format string, v ...interface{}) (err error) {
    l.Info(format, v...)
    return
}
func (da *logAdapter) Warning(v ...interface{}) (err error) {
    if len(v) > 1 {
        l.Warn(v[0].(string), v[1:]...)
        return
    }
    l.Warn(v[0].(string))
    return
}
func (da *logAdapter) Warningf(format string, v ...interface{}) (err error) {
    l.Warn(format, v...)
    return
}

func (da *logAdapter) Level() core.LogLevel {
    return core.LOG_DEBUG
}
func (da *logAdapter) SetLevel(l core.LogLevel) (err error) {
    return nil
}

func Init() {
    var (
        err error
    )
    Engine, err = xorm.NewEngine(sql.Drivers()[0], conf.String("database", "DSN"))
    if err != nil {
        panic(err)
    }
    if conf.ENV != "release" {
        Engine.ShowSQL = true
    }
    Engine.SetMaxOpenConns(conf.Int("database", "MAX_CONNECTION", 10))
    Engine.SetMaxIdleConns(conf.Int("database", "MAX_IDLE_CONNECTION", 50))
    Engine.SetLogger(&logAdapter{})
}
