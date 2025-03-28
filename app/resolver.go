package app

type Resolver interface {
	LookupIP(host string) ([]string, error)
	LookupNS(host string) ([]string, error)
	LookupMX(host string) ([]MXRecord, error)
	LookupCNAME(host string) (string, error)
	LookupAddr(ip string) ([]string, error)
	Whois(host string) (string, error)
	Ping(host string) (int64, error)
	TCPScan(host string, port int) (bool, error)
}

// MXRecord representa o registro MX
type MXRecord struct {
	Host string
	Pref uint16
}
