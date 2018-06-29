package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"ss-server/models"
	"strconv"
	"strings"
)

func DecodeBrookUrl(encodeurl string) (*models.Brook, error) {
	decodeurl, err := url.QueryUnescape(encodeurl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	array := strings.Split(decodeurl, " ")
	brook := new(models.Brook)
	brook.OriginUrl = url.QueryEscape(encodeurl)
	brook.BrookType = string([]byte(array[0])[8:len(array[0])])
	brook.IP = strings.Split(array[1], ":")[0]
	portStr := strings.Split(array[1], ":")[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println(err)
	}
	brook.Port = port
	brook.Password = AesEncrypt(array[2])
	fmt.Printf("%+v\n", *brook)
	return brook, nil
}

func DecodeShadowSocksUrl(encodeurl string) (*models.Profile, error) {
	// var pattern = "(?i)ss://[-a-zA-Z0-9+&@#/%?=~_|!:,.;\\[\\]]*[-a-zA-Z0-9+&@#/%=~_|\\[\\]]"
	// var userInfoPattern = "^(.+?):(.*)$"
	var legacyPattern = "^(.+?):(.*)@(.+?):(\\d+?)$"
	_url, err := url.Parse(encodeurl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if _url.User == nil {
		decodeHost, err := base64.StdEncoding.DecodeString(_url.Host)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		r, err := regexp.Compile(legacyPattern)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		match := r.FindAllStringSubmatch(string(decodeHost), 1)
		profile := new(models.Profile)
		profile.OriginUrl = url.QueryEscape(encodeurl)
		profile.Method = match[0][1]
		profile.Password = AesEncrypt(match[0][2])
		profile.Host = match[0][3]
		remotePort, err := strconv.Atoi(match[0][4])
		profile.RemotePort = remotePort
		profile.Plugin = _url.Query().Get("plugin")
		profile.Name = _url.Fragment
		profile.VpnType = 1
		fmt.Printf("%+v\n", *profile)
		return profile, nil
	} else {
		fmt.Println("need userInfoPattern")
	}
	return nil, errors.New("can't parse the url")
}
