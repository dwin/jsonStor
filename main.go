package main

import (
	"github.com/dwin/jsonStor/config"
	"github.com/dwin/jsonStor/controller"
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
)

func main() {
	r := gin.Default()

	// Static Route
	r.Static("/static", config.GetStaticPath())

	// Load Templates
	r.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "view",
	})

	//
	// Web Routes
	//
	r.GET("/", controller.Index)

	//
	// JSON API Routes
	//

	a0 := r.Group("/api/v0")

	r.Run(config.GetHostPort)
}
