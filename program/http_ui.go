package program

import (
	"mime"
	"net/http"
	"path"
	"strings"

	gin "github.com/gin-gonic/gin"
	"github.com/shiguanghuxian/etcd-manage/program/logger"
	"github.com/shiguanghuxian/etcd-manage/tpls"
)

// ui 界面
// 处理静态文件
func (p *Program) handlerStatic(c *gin.Context) {
	uri := strings.TrimLeft(c.Request.RequestURI, "/")
	uri = strings.TrimRight(uri, "?")
	if uri == "ui/" || uri == "ui" {
		uri = "dist/index.html"
	} else {
		uri = strings.Replace(uri, "ui", "dist", 1)
	}
	// log.Println(uri)
	// 读取模版内容
	body, err := tpls.Asset(uri)
	if err != nil {
		logger.Log.Errorw("UI static file reading error", "err", err)
		c.Status(http.StatusNotFound)
		return
	}
	mimetype := mime.TypeByExtension(path.Ext(uri))
	if mimetype != "" {
		c.Header("Content-Type", mimetype)
	}

	c.Writer.Write(body)
}
