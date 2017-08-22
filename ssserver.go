package main

import (
	"net/http"
	profileApi "ss-server/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getAllprofile", profileApi.GetAllProfileAPI)
	router.Static("/static", "./app/build/static")
	router.StaticFile("/favicon.ico", "./app/build/favicon.ico")
	router.StaticFile("/manifest.json", "./app/build/manifest.json")
	// router.LoadHTMLGlob("app/build/*/*.html")
	router.LoadHTMLFiles("app/build/index.html", "app/build/login.html", "app/build/profile/profile_list.html", "app/build/profile/profile_edit.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/profileList", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile_list.html", nil)
	})
	router.GET("/profileEdit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile_edit.html", nil)
	})
	return router
}
