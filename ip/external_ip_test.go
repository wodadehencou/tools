package ip_test

import (
	"testing"

	"github.com/wodadehencou/tools/ip"
)

func TestExternalIP(t *testing.T) {
	ipAddr, err := ip.ExternalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ipAddr)
}
