package models

type Profile struct {
	ID            int
	Name          string
	Host          string
	LocalPort     int
	RemotePort    int
	Password      string
	Protocol      string
	ProtocolParam string
	Obfs          string
	ObfsParam     string
	Method        string
	Route         string
	RemoteDNS     string
	ProxyApps     int
	Bypass        int
	Udpdns        int
	Ipv6          int
	Individual    string
	Date          string
	UserOrder     int
	Plugin        string
	Country       string
	VpnType       int
	Ikev2Type     int
}
