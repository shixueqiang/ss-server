package database

import (
	"log"
	profile "ss-server/models"
	cryptoUtil "ss-server/utils"
)

/*
查询所有的profile配置信息
*/
func QueryAllProfile() (profiles []profile.Profile, err error) {
	profiles = make([]profile.Profile, 0)
	rows, err := Db.Query("SELECT id,name,host,local_port,remote_port,password,method,route,remote_dns,proxy_apps,bypass,udpdns,ipv6,individual,date,user_order,plugin,country,type,ikev2_type FROM vpn_profile")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var profile profile.Profile
		rows.Scan(&profile.ID, &profile.Name, &profile.Host, &profile.LocalPort, &profile.RemotePort, &profile.Password, &profile.Method,
			&profile.Route, &profile.RemoteDNS, &profile.ProxyApps, &profile.Bypass, &profile.Udpdns, &profile.Ipv6, &profile.Individual,
			&profile.Date, &profile.UserOrder, &profile.Plugin, &profile.Country, &profile.VpnType, &profile.Ikev2Type)
		profile.Host = cryptoUtil.AesEncrypt(profile.Host)
		profile.Password = cryptoUtil.AesEncrypt(profile.Password)
		profiles = append(profiles, profile)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//defer 延迟执行即QueryAllProfile方法结束后才执行
	// defer Db.Close()
	return
}

func InsertProfile(profile *profile.Profile) {

}
