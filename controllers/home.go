package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//new template engine
	// router.HTMLRender = ginview.Default()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
	// router.GET("/", func(ctx *gin.Context) {
	// 	//render with master
	// 	ctx.HTML(http.StatusOK, "index", gin.H{
	// 		"title": "Index title!",
	// 		"add": func(a int, b int) int {
	// 			return a + b
	// 		},
	// 	})
	// })

	// //new template engine
	// router.HTMLRender = ginview.New(goview.Config{
	// 	Root:      "views",
	// 	Extension: ".html",
	// 	Master:    "layouts/master",
	// 	Partials:  []string{"partials/ad"},
	// 	Funcs: template.FuncMap{
	// 		"copy": func() string {
	// 			return time.Now().Format("2006")
	// 		},
	// 	},
	// 	DisableCache: true,
	// })

	// router.GET("/", func(ctx *gin.Context) {
	// 	// `HTML()` is a helper func to deal with multiple TemplateEngine's.
	// 	// It detects the suitable TemplateEngine for each path automatically.
	// 	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
	// 		"title": "Frontend title!",
	// 	})
	// })

}
