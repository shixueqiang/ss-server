package main

import (
	"net/http"
	profileApi "ss-server/apis"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/getAllprofile", profileApi.GetAllProfileAPICrypto)
	router.GET("/getAllprofileNotCrypto", profileApi.GetAllProfileAPINotCrypto)
	router.Static("/static", "./app/build/static")
	router.StaticFile("/favicon.ico", "./app/build/favicon.ico")
	router.StaticFile("/manifest.json", "./app/build/manifest.json")
	// router.LoadHTMLGlob("app/build/*/*.html")
	router.LoadHTMLFiles("app/build/index.html", "app/build/login.html", "app/build/profile/profiles.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/toLogin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("login", profileApi.Login)
	//profile相关的做react单页面应用
	router.GET("/toProfiles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profiles.html", nil)
	})
	router.GET("/toProfileEdit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profiles.html", nil)
	})
	router.GET("/toProfileInsert", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profiles.html", nil)
	})

	router.POST("/profileInsert", profileApi.InsertProfile)
	router.POST("/profileUpdate", profileApi.UpdateProfile)
	router.POST("/profileRemove", profileApi.RemoveProfile)
	return router
}
