package v2

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

func TestNetworkListEndpoint_Append(t *testing.T) {
	var (
		status   *string
		required string
	)

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

	if err := nls.Append([]string{"1.2.3.4/32", "1.2.3.5/32", "1.2.3.6/32"}); err != nil {
		t.Errorf("Add() error = %s", err)
	}

	if status, err = nls.GetActivationStatus(Staging); err != nil {
		t.Errorf("GetActivationStatus(Staging) error = %s", err)
	}

	required = "MODIFIED"
	if *status != required {
		t.Errorf("Invalid status: found = %v, required = %v", *status, required)
	}

	if err := nls.Activate(Staging, "test only", nil); err != nil {
		t.Errorf("Activate(Staging) error = %s", err)
	}

	if status, err = nls.GetActivationStatus(Staging); err != nil {
		t.Errorf("GetActivationStatus(Staging) error = %s", err)
	}

	required = "PENDING_ACTIVATION"
	if *status != required {
		t.Errorf("Invalid status: found = %v, required = %v", *status, required)
	}
}
