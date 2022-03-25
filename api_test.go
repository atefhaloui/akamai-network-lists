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
		t.Fatalf("Cannot load config file: %s", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %s", err)
	}

	if err := nls.Add("195.109.4.59/32"); err != nil {
		t.Errorf("Add() error = %s", err)
	}

	if err := nls.Activate(Staging, "test only", nil); err != nil {
		t.Errorf("Activate(Staging) error = %s", err)
	}

	var status *string
	required := "MODIFIED"
	if status, err = nls.GetActivationStatus(Staging); err != nil {
		t.Errorf("GetActivationStatus(Staging) error = %s", err)
	}

	if *status != required {
		t.Errorf("Invalid status: found = %s, required = %s", *status, required)
	}
}

func TestNetworkListEndpoint_Remove_WithError(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = false
	if err != nil {
		t.Fatalf("Cannot load config file: %s", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %s", err)
	}

	if err := nls.Delete("1.2.3.4/32"); err == nil {
		t.Errorf("Delete() didn't fail as expected")
	}
}
