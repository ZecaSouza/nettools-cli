// app/netresolver.go
package app

import (
	"fmt"
	"net"
	"time"
)

type netResolver struct{}

func NewNetResolver() Resolver {
	return &netResolver{}
}

func (r *netResolver) LookupIP(host string) ([]string, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, ip := range ips {
		list = append(list, ip.String())
	}
	return list, nil
}

func (r *netResolver) LookupNS(host string) ([]string, error) {
	ns, err := net.LookupNS(host)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, n := range ns {
		list = append(list, n.Host)
	}
	return list, nil
}

func (r *netResolver) LookupMX(host string) ([]MXRecord, error) {
	mx, err := net.LookupMX(host)
	if err != nil {
		return nil, err
	}
	var list []MXRecord
	for _, m := range mx {
		list = append(list, MXRecord{Host: m.Host, Pref: m.Pref})
	}
	return list, nil
}

func (r *netResolver) LookupCNAME(host string) (string, error) {
	return net.LookupCNAME(host)
}

func (r *netResolver) LookupAddr(ip string) ([]string, error) {
	return net.LookupAddr(ip)
}

func (r *netResolver) Whois(host string) (string, error) {
	conn, err := net.Dial("tcp", "whois.verisign-grs.com:43")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	fmt.Fprintf(conn, "%s\r\n", host)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return result, nil
}

func (r *netResolver) Ping(host string) (int64, error) {
	start := time.Now()
	_, err := net.LookupIP(host)
	if err != nil {
		return 0, err
	}
	return time.Since(start).Milliseconds(), nil
}

func (r *netResolver) TCPScan(host string, port int) (bool, error) {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false, nil
	}
	defer conn.Close()
	return true, nil
}
