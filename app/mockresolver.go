// app/mockresolver.go
package app

type MockResolver struct {
	IPResult      []string
	NSResult      []string
	MXResult      []MXRecord
	CNAMEResult   string
	WhoisResult   string
	AddrResult    []string
	PingLatency   int64
	TCPPortsOpen  map[int]bool
	Error         error
}

func (m *MockResolver) LookupIP(host string) ([]string, error)         { return m.IPResult, m.Error }
func (m *MockResolver) LookupNS(host string) ([]string, error)         { return m.NSResult, m.Error }
func (m *MockResolver) LookupMX(host string) ([]MXRecord, error)       { return m.MXResult, m.Error }
func (m *MockResolver) LookupCNAME(host string) (string, error)        { return m.CNAMEResult, m.Error }
func (m *MockResolver) LookupAddr(ip string) ([]string, error)         { return m.AddrResult, m.Error }
func (m *MockResolver) Whois(host string) (string, error)              { return m.WhoisResult, m.Error }
func (m *MockResolver) Ping(host string) (int64, error)                { return m.PingLatency, m.Error }
func (m *MockResolver) TCPScan(host string, port int) (bool, error) {
	return m.TCPPortsOpen[port], m.Error
}
