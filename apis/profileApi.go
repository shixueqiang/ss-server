package apis

import (
	"log"
	"net/http"
	profileDao "ss-server/database"

	"github.com/gin-gonic/gin"
)

func GetAllProfileAPICrypto(c *gin.Context) {
	profiles, err := profileDao.QueryAllProfile(true)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}

func GetAllProfileAPINotCrypto(c *gin.Context) {
	profiles, err := profileDao.QueryAllProfile(false)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}
