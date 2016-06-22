package sessions
import (
    "github.com/astaxie/beego/session"
    "xtudouh/common/log"
    "xtudouh/common/conf"
    "fmt"
    _ "github.com/astaxie/beego/session/redis"
    "github.com/gin-gonic/gin"
)


/**************************EXAMPLE**************************

    sess := GetSession(c)
    defer sess.SessionRelease(c.Writer)
    sess.Set("username", time.Now().Format(timeLayout))

************************************************************/

const sessionKey = "prosnav.session.key"

var (
    sessionManager *session.Manager
    logger = log.NewLogger()
)

func Init() {
    var err error
    sessionConfig := fmt.Sprintf(`{"cookieName":"gosessionid","gclifetime":%d,"ProviderConfig":"%s"}`,
        conf.Int("session", "GC_LIFETIME"),
        fmt.Sprintf("%s, %d", conf.String("redis", "REDIS_SERVER"), conf.Int("session", "SESSION_LENGTH")),
    )
    sessionManager, err = session.NewManager(conf.String("session", "PROVIDER"), sessionConfig)
    if err != nil {
        panic(err)
    }

    //go sessionManager.GC()
}

func startSession(c *gin.Context) session.SessionStore {
    session, err := sessionManager.SessionStart(c.Writer, c.Request)
    if err != nil {
        logger.Error("failed to start session, %v", err)
        panic(err)
    }
    return session
}

func Sessions() gin.HandlerFunc{
    return func(c *gin.Context) {
        sess := startSession(c)
        c.Set(sessionKey, sess)
        defer sess.SessionRelease(c.Writer)
        c.Next()
    }
}

func Get(key string, c *gin.Context) interface{} {
    return c.MustGet(sessionKey).(session.SessionStore).Get(key)
}

func Set(key string, val interface{}, c *gin.Context) {
    c.MustGet(sessionKey).(session.SessionStore).Set(key, val)
}

func SessionId(c *gin.Context) string {
    return c.MustGet(sessionKey).(session.SessionStore).SessionID()
}
