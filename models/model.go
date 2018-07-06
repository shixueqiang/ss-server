package models

type Profile struct {
	OriginUrl     string
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
	VpnType       int //1 ss,2 brook,3 strongswan
	BrookType     string
}

type Brook struct {
	OriginUrl string
	Name      string
	BrookType string
	IP        string
	Port      int
	Password  string
}

type Package struct {
	// Brooks []Brook
	Profiles []Profile
}

type User struct {
	ID       int
	Name     string
	Account  string
	Password string
}

func (brook *Brook) ToProfile() *Profile {
	profile := new(Profile)
	profile.OriginUrl = brook.OriginUrl
	profile.Name = brook.Name
	profile.VpnType = 2
	profile.BrookType = brook.BrookType
	profile.Host = brook.IP
	profile.RemotePort = brook.Port
	profile.Password = brook.Password
	return profile
}
