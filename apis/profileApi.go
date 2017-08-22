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

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}
