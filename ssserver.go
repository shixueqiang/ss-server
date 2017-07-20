package main

import (
	profileApi "ss-server/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getAllprofile", profileApi.GetAllProfileAPI)

	return router
}
