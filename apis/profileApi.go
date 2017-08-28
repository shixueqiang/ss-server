package apis

import (
	"log"
	"net/http"
	profileDao "ss-server/database"
	profile "ss-server/models"
	"strconv"
	"time"

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

func InsertProfile(c *gin.Context) {
	Name := c.PostForm("Name")
	Host := c.PostForm("Host")
	LocalPort := c.PostForm("LocalPort")
	RemotePort := c.PostForm("RemotePort")
	Password := c.PostForm("Password")
	Protocol := c.PostForm("Protocol")
	ProtocolParam := c.PostForm("ProtocolParam")
	Obfs := c.PostForm("Obfs")
	ObfsParam := c.PostForm("ObfsParam")
	Method := c.PostForm("Method")
	Route := c.PostForm("Route")
	RemoteDNS := c.PostForm("RemoteDNS")
	VpnType := c.PostForm("VpnType")
	Ikev2Type := c.PostForm("Ikev2Type")
	var mProfile *profile.Profile = new(profile.Profile)
	mProfile.Name = Name
	mProfile.Host = Host
	localPort, err := strconv.Atoi(LocalPort)
	mProfile.LocalPort = localPort
	remotePort, err := strconv.Atoi(RemotePort)
	mProfile.RemotePort = remotePort
	mProfile.Password = Password
	mProfile.Protocol = Protocol
	mProfile.ProtocolParam = ProtocolParam
	mProfile.Obfs = Obfs
	mProfile.ObfsParam = ObfsParam
	mProfile.Method = Method
	mProfile.Route = Route
	mProfile.RemoteDNS = RemoteDNS
	mProfile.ProxyApps = 0
	mProfile.Bypass = 0
	mProfile.Udpdns = 0
	mProfile.Ipv6 = 0
	mProfile.Individual = ""
	mProfile.Date = time.Now().Format("2006-01-02 15:04:05")
	mProfile.UserOrder = 0
	mProfile.Plugin = ""
	mProfile.Country = ""
	vpnType, err := strconv.Atoi(VpnType)
	mProfile.VpnType = vpnType
	ikev2Type, err := strconv.Atoi(Ikev2Type)
	mProfile.Ikev2Type = ikev2Type
	id, err := profileDao.InsertProfile(mProfile)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if id > 0 {
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusOK, "fail")
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
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusOK, "fail")
	}
}

func RemoveProfile(c *gin.Context, ids []int) {

	rowsCnt, err := profileDao.RemoveProfiles(ids)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if rowsCnt > 0 {
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusOK, "fail")
	}
}
