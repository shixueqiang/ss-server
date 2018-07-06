package apis

import (
	"fmt"
	"log"
	"net/http"
	profileDao "ss-server/database"
	profile "ss-server/models"
	"ss-server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	log.Printf("account:%s,password:%s\n", account, password)
	if account != "shixq" || password != "shixq321" {
		c.String(http.StatusOK, "fail")
	}
	// user, err := profileDao.QueryUser(account, password)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	c.String(http.StatusOK, "fail")
	// }
	// var userStr []byte
	// userStr, err = json.Marshal(user)
	// log.Printf(string(userStr))
	session := sessions.Default(c)
	// session.Set("user", string(userStr))
	session.Set("user", "{\"Name\":shixq}")
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
	p, err := profileDao.QueryAll()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"profiles": p.Profiles,
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

func ImportProfile(c *gin.Context) {
	encodeUrl := c.PostForm("url")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if strings.Index(encodeUrl, "brook://") != -1 {
		brook, e := utils.DecodeBrookUrl(encodeUrl)
		e = profileDao.InsertProfileToMsgpack(brook.ToProfile())
		if e != nil {
			c.String(http.StatusOK, "导入失败")
		} else {
			c.String(http.StatusOK, "导入成功")
		}
	} else if strings.Index(encodeUrl, "ss://") != -1 {
		profile, e := utils.DecodeShadowSocksUrl(encodeUrl)
		e = profileDao.InsertProfileToMsgpack(profile)
		if e != nil {
			c.String(http.StatusOK, "导入失败")
		} else {
			c.String(http.StatusOK, "导入成功")
		}
	} else {
		c.String(http.StatusOK, "请输入合法的url")
	}
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
	BrookType := c.PostForm("BrookType")
	var mProfile *profile.Profile = new(profile.Profile)
	mProfile.Name = Name
	mProfile.Host = Host
	localPort, err := strconv.Atoi(LocalPort)
	mProfile.LocalPort = localPort
	remotePort, err := strconv.Atoi(RemotePort)
	mProfile.RemotePort = remotePort
	mProfile.Password = utils.AesEncrypt(Password)
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
	mProfile.BrookType = BrookType
	mProfile.OriginUrl = utils.AesEncrypt(utils.ToShadowSocksUrl(mProfile))
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	err = profileDao.InsertProfileToMsgpack(mProfile)
	if err != nil {
		c.String(http.StatusOK, "插入失败")
	} else {
		c.String(http.StatusOK, "插入成功")
	}
}

func UpdateProfile(c *gin.Context) {
	OriginUrl := c.PostForm("OriginUrl")
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
	BrookType := c.PostForm("BrookType")
	var mProfile *profile.Profile = new(profile.Profile)
	mProfile.OriginUrl = OriginUrl
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
	mProfile.BrookType = BrookType
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	err = profileDao.UpdateProfileFromMsgpack(mProfile)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, "更新失败")
	} else {
		c.String(http.StatusOK, "更新成功")
	}
}

func RemoveProfile(c *gin.Context) {
	// idArray := c.PostForm("removeIds")
	// ids := strings.Split(idArray, ",")
	urlArray := c.PostForm("removeUrls")
	fmt.Printf("%s\n", urlArray)
	urls := strings.Split(urlArray, " ")
	// rowsCnt, err := profileDao.RemoveProfiles(ids)
	var err error
	for i := 0; i < len(urls); i++ {
		err = profileDao.RemoveProfileFromMsgpack(urls[i])
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	if err != nil {
		c.String(http.StatusOK, "移除失败")
	} else {
		c.String(http.StatusOK, "移除成功")
	}
}
