package midwares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"runtime"
)

func stack(skip int) string {
	stk := make([]uintptr, 32)
	str := ""
	l := runtime.Callers(skip, stk[:])
	for i := 0; i < l; i++ {
		f := runtime.FuncForPC(stk[i])
		name := f.Name()
		file, line := f.FileLine(stk[i])
		str += fmt.Sprintf("\n    %-30s [%s:%d]", name, path.Base(file), line)
	}
	return str
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				l.Error("%v\n%v", err, stack(3))
				c.JSON(500, err)
				c.Abort()
			}
		}()

		c.Next()
	}
}
