package apis

import (
	"encoding/json"
	"log"
	"net/http"
	profileDao "ss-server/database"
	profile "ss-server/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	user, err := profileDao.QueryUser(account, password)
	if err != nil {
		log.Fatalln(err)
		c.String(http.StatusOK, "fail")
	}
	var userStr []byte
	userStr, err = json.Marshal(user)
	log.Printf(string(userStr))
	session := sessions.Default(c)
	session.Set("user", string(userStr))
	session.Save()
	c.String(http.StatusOK, "success")
}

func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	c.String(http.StatusOK, "success")
}

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
		c.String(http.StatusOK, "插入成功")
	} else {
		c.String(http.StatusOK, "插入失败")
	}
}

func UpdateProfile(c *gin.Context) {
	ID := c.PostForm("ID")
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
	mId, err := strconv.Atoi(ID)
	mProfile.ID = mId
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
	rowsCnt, err := profileDao.UpdateProfile(mProfile)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if rowsCnt > 0 {
		c.String(http.StatusOK, "更新成功")
	} else {
		c.String(http.StatusOK, "更新失败")
	}
}

func RemoveProfile(c *gin.Context) {
	idArray := c.PostForm("removeIds")
	ids := strings.Split(idArray, ",")
	rowsCnt, err := profileDao.RemoveProfiles(ids)
	if err != nil {
		log.Fatalln(err)
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if rowsCnt > 0 {
		c.String(http.StatusOK, "移除成功")
	} else {
		c.String(http.StatusOK, "移除失败")
	}
}
