package v2

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func TestNetworkListEndpoint_Add(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = true
	if err != nil {
		t.Fatalf("Cannot load config file: %s", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %s", err)
	}

	if err := nls.Add("1.3.4.5/32"); err != nil {
		t.Fatalf("Add() error = %s", err)
	}
}

func TestNetworkListEndpoint_Remove(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = true
	if err != nil {
		t.Fatalf("Cannot load config file: %s", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %s", err)
	}

	// Delete an existing entry (make sure
	_ = nls.Add("1.3.4.5/32")
	if err := nls.Delete("1.3.4.5/32"); err != nil {
		t.Errorf("Delete() failed: %s", err)
	}

	// Delete a non-existing entry
	if err := nls.Delete("100.3.4.5/32"); err == nil {
		t.Errorf("Delete() didn't fail as expected")
	}
}

func TestNetworkListEndpoint_Activate(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = true
	if err != nil {
		t.Fatalf("Cannot load config file: %s", err)
	}

	// Select Network list: IP Blacklist
	nls, err := New(&config, "XXXXX_IPBLACKLIST") // @todo: Update the ID
	if err != nil {
		t.Fatalf("Cannot access the network list: %s", err)
	}

	// Add a fake entry
	if err := nls.Add("1.3.4.5/32"); err != nil {
		t.Errorf("Add() failed: %s", err)
	}

	if err := nls.Activate(Staging, "test only", nil); err != nil {
		t.Fatalf("Activate(Staging) error = %s", err)
	}

	var status *string
	required := "PENDING_ACTIVATION"
	if status, err = nls.GetActivationStatus(Staging); err != nil {
		t.Errorf("GetActivationStatus(Staging) error = %s", err)
	}

	if *status != required {
		t.Errorf("Invalid status: found = %s, required = %s", *status, required)
	}
}
