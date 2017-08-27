package database

import (
	"log"
	profile "ss-server/models"
	cryptoUtil "ss-server/utils"
)

/*
查询所有的profile配置信息
*/
func QueryAllProfile(isCrypto bool) (profiles []profile.Profile, err error) {
	profiles = make([]profile.Profile, 0)
	rows, err := Db.Query("SELECT id,name,host,local_port,remote_port,password,protocol,protocol_param,obfs,obfs_param,method,route,remote_dns,proxy_apps,bypass,udpdns,ipv6,individual,date,user_order,plugin,country,type,ikev2_type FROM vpn_profile")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var profile profile.Profile
		rows.Scan(&profile.ID, &profile.Name, &profile.Host, &profile.LocalPort, &profile.RemotePort, &profile.Password, &profile.Protocol, &profile.ProtocolParam, &profile.Obfs, &profile.ObfsParam, &profile.Method,
			&profile.Route, &profile.RemoteDNS, &profile.ProxyApps, &profile.Bypass, &profile.Udpdns, &profile.Ipv6, &profile.Individual,
			&profile.Date, &profile.UserOrder, &profile.Plugin, &profile.Country, &profile.VpnType, &profile.Ikev2Type)
		if isCrypto {
			profile.Host = cryptoUtil.AesEncrypt(profile.Host)
			profile.Password = cryptoUtil.AesEncrypt(profile.Password)
		}
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

func InsertProfile(profile *profile.Profile) (id int64, err error) {
	rows, err := Db.Exec("INSERT INTO vpn_profile(name,host,local_port,remote_port,password,protocol,protocol_param,obfs,obfs_param,method,route,remote_dns,proxy_apps,bypass,udpdns,ipv6,individual,date,user_order,plugin,country,type,ikev2_type) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		profile.Name, profile.Host, profile.LocalPort, profile.RemotePort, profile.Password, profile.Protocol, profile.ProtocolParam, profile.Obfs, profile.ObfsParam, profile.Method, profile.Route, profile.RemoteDNS, profile.ProxyApps, profile.Bypass, profile.Udpdns, profile.Ipv6, profile.Individual, profile.Date, profile.UserOrder, profile.Plugin, profile.Country, profile.VpnType, profile.Ikev2Type)
	if err != nil {
		return
	}
	id, err = rows.LastInsertId()
	return
}

func UpdateProfile(profile *profile.Profile) (rowsCnt int64, err error) {
	rows, err := Db.Exec("UPDATE vpn_profile SET name = ?,host = ?,local_port = ?,remote_port = ?,password = ?,protocol = ?,protocol_param = ?,obfs = ?,obfs_param = ?,method = ?,route = ?,remote_dns = ?,proxy_apps = ?,bypass = ?,udpdns = ?,ipv6 = ?,individual = ?,date = ?,user_order = ?,plugin = ?,country = ?,type = ?,ikev2_type = ?",
		profile.Name, profile.Host, profile.LocalPort, profile.RemotePort, profile.Password, profile.Protocol, profile.ProtocolParam, profile.Obfs, profile.ObfsParam, profile.Method, profile.Route, profile.RemoteDNS, profile.ProxyApps, profile.Bypass, profile.Udpdns, profile.Ipv6, profile.Individual, profile.Date, profile.UserOrder, profile.Plugin, profile.Country, profile.VpnType, profile.Ikev2Type)
	if err != nil {
		return
	}
	rowsCnt, err = rows.RowsAffected()
	return
}

func RemoveProfiles(ids []int) (cnt int64, err error) {
	cnt = int64(0)
	for i := 0; i < len(ids); i++ {
		rows, err := Db.Exec("DELETE FROM vpn_profile where id = ?", ids[i])
		if err != nil {
			log.Fatal(err)
		}
		rowsCnt := int64(0)
		rowsCnt, err = rows.RowsAffected()
		cnt += rowsCnt
	}
	return
}
