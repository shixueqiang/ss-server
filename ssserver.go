package main

import (
	"log"
	"net/http"
	"os"
	profileApi "ss-server/apis"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var APP_PATH string

func init() {
	APP_PATH = os.Getenv("GOBIN") + "/ss-server"
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		session := sessions.Default(c)
		mUser := session.Get("user")
		log.Printf("mUser:%s", mUser)
		uri := c.Request.RequestURI
		if mUser == nil && uri != "/toLogin" && uri != "/login" {
			//301永久重定向会被浏览器缓存，302临时重定向
			c.Redirect(http.StatusFound, "/toLogin")
		} else {
			// Set example variable
			c.Set("loginUser", mUser)
		}

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", APP_PATH+"/app/build/static")
	router.StaticFile("/favicon.ico", APP_PATH+"/app/build/favicon.ico")
	router.StaticFile("/manifest.json", APP_PATH+"/app/build/manifest.json")
	// router.LoadHTMLGlob("app/build/*/*.html")
	router.LoadHTMLFiles(APP_PATH+"/app/build/index.html", APP_PATH+"/app/build/login.html", APP_PATH+"/app/build/profile/profiles.html")

	router.GET("/getAllprofile", profileApi.GetAllProfileAPICrypto)

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(Logger())

	router.GET("/getAllprofileNotCrypto", profileApi.GetAllProfileAPINotCrypto)
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/toLogin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("login", profileApi.Login)
	router.POST("signOut", profileApi.SignOut)
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
	router.POST("/profileImport", profileApi.ImportProfile)
	return router
}
