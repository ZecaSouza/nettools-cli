// app/netresolver_test.go
package app

import (
	"testing"
)

func TestLookupIPWithMock(t *testing.T) {
	mock := &MockResolver{
		IPResult: []string{"192.168.0.1"},
	}

	ips, err := mock.LookupIP("example.com")
	if err != nil {
		t.Fatal("Erro inesperado:", err)
	}
	if len(ips) != 1 || ips[0] != "192.168.0.1" {
		t.Errorf("Resultado inesperado: %v", ips)
	}
}

func TestTCPScanWithMock(t *testing.T) {
	mock := &MockResolver{
		TCPPortsOpen: map[int]bool{
			80:  true,
			443: false,
		},
	}

	open, _ := mock.TCPScan("example.com", 80)
	if !open {
		t.Error("Porta 80 deveria estar aberta")
	}
	open, _ = mock.TCPScan("example.com", 443)
	if open {
		t.Error("Porta 443 deveria estar fechada")
	}
}
