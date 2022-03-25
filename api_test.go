package v2

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func TestNetworkListEndpoint_Add(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = false
	if err != nil {
		t.Fatalf("Cannot load config file: %v", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %v", err)
	}

	if err := nls.Add("195.109.4.59/32"); err != nil {
		t.Errorf("Add() error = %v", err)
	}

	if err := nls.Activate(Staging, "test only", nil); err != nil {
		t.Errorf("Activate(Staging) error = %v", err)
	}

	var status *string
	required := "MODIFIED"
	if status, err = nls.GetActivationStatus(Staging); err != nil {
		t.Errorf("GetActivationStatus(Staging) error = %v", err)
	}

	if *status != required {
		t.Errorf("Invalid status: found = %v, required = %v", *status, required)
	}
}
