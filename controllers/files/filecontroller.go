package files

import (
	"github.com/gin-gonic/gin"
	"xtudouh/common/log"
	"xtudouh/common/fs"
	"mime/multipart"
)

var l = log.NewLogger()

func Upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		l.Error("%v", err)
		c.AbortWithStatus(500)
		return
	}
	m := c.Request.MultipartForm
	files := m.File["files"]

	fidList, err := saveFiles(files)
	if err != nil {
		l.Error("%v", err)
		c.AbortWithStatus(500)
		return
	}
	//TODO Save fids
	l.Debug("file list: %v", fidList)

	c.JSON(200, nil)
}

type multiForm struct {
	UserName string 	`form:"username" json:"username"`
	Address  string		`form:"address"  json:"address"`
	Email    string		`form:"email"    json:"email"`
}

func MultiUpload(c *gin.Context) {
	form := new(multiForm)
	if err := c.Bind(form); err != nil {
		l.Error("%v", err)
		c.AbortWithStatus(500)
		return
	}
	l.Debug("form data: %+v", form)
	files := c.Request.MultipartForm.File["files"]
	fidList, err := saveFiles(files)
	if err != nil {
		l.Error("%v", err)
		c.AbortWithStatus(500)
		return
	}
	//TODO Save data
	l.Debug("file list: %v", fidList)
	c.JSON(200, nil)
}

func saveFiles(files []*multipart.FileHeader) ([]string, error) {
	list := make([]string, len(files))
	for i, f := range files {
		src, err := f.Open()
		defer src.Close()
		if err != nil {
			l.Error("%v", err)
			return nil, err
		}
		mimeType := f.Header.Get("Content-Type")
		fid, err := fs.SaveFile(f.Filename, mimeType, src)
		if err != nil {
			l.Error("%v", err)
			return nil, err
		}
		list[i] = fid
	}

	return list, nil
}
