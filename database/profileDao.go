package database

import (
	"database/sql"
	"log"
	profile "ss-server/models"
	cryptoUtil "ss-server/utils"
)

var Db *sql.DB

func initDB() {
	var err error
	//打开数据库
	//account:password@tcp(host:3306)/dbname?charset=utf8
	Db, err = sql.Open("mysql", "shixq:shixq1207@tcp(www.shixq.com:3306)/ssvpn?charset=utf8")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	//连接数据库
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

/*
查询所有的profile配置信息
*/
func QueryAllProfile() (profiles []profile.Profile, err error) {
	initDB()
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
		profile.Host = cryptoUtil.AesEncrypt1([]byte(profile.Host))
		profile.Password = cryptoUtil.AesEncrypt1([]byte(profile.Password))
		profiles = append(profiles, profile)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//defer 延迟执行即QueryAllProfile方法结束后才执行
	defer Db.Close()
	return
}

func InsertProfile(profile *profile.Profile) {

}
