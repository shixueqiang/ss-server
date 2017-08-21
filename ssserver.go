package main

import (
	"net/http"
	profileApi "ss-server/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getAllprofile", profileApi.GetAllProfileAPI)
	router.Static("/static", "./my-app/static")
	router.StaticFile("/favicon.ico", "./my-app/favicon.ico")
	router.StaticFile("/manifest.json", "./my-app/manifest.json")
	router.LoadHTMLGlob("my-app/*.html")
	// router.LoadHTMLFiles("my-app/index.html", "my-app/test139.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return router
}
