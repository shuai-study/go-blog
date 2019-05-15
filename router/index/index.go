/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-20
 * Time: 23:36
 */
package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Web struct {
}

func NewIndex() Home {
	return &Web{}
}

func (w *Web)Index(c *gin.Context) {
	data := make(map[string]string)
	data["he"] = "开玩笑"
	data["ha"] = "大小"

	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title": "M3ain website",
		})
	return
}