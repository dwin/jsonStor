package controller

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data := pongo2.Context{
		"title": "photos2go",
	}
	c.HTML(200, "index.html", data)
	return
}
