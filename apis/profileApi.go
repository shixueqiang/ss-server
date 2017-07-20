package apis

import (
	"log"
	"net/http"
	profileDao "ss-server/database"

	"github.com/gin-gonic/gin"
)

func GetAllProfileAPI(c *gin.Context) {
	profiles, err := profileDao.QueryAllProfile()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}
