package apis

import (
	"log"
	"net/http"
	profileDao "ss-server/database"
	profile "ss-server/models"

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

func InsertProfile(c *gin.Context, profile *profile.Profile) {
	id, err := profileDao.InsertProfile(profile)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
}

func UpdateProfile(c *gin.Context, profile *profile.Profile) {
	rowsCnt, err := profileDao.UpdateProfile(profile)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if rowsCnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
}

func RemoveProfile(c *gin.Context, ids []int) {

	rowsCnt, err := profileDao.RemoveProfiles(ids)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if rowsCnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
		})
	}
}
