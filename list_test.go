package v2

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

type args struct {
	config          edgegrid.Config
	includeElements bool
	search          *string
	listType        *string
	extended        bool
}

func TestListNetworkLists(t *testing.T) {
	// Load config
	config, err := edgegrid.InitEdgeRc("./edgerc", "ccu")
	config.Debug = true
	if err != nil {
		t.Fatalf("Cannot load config file: %v", err)
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Get'em all",
			args: args{
				config:          config,
				includeElements: false,
				extended:        false,
				listType:        nil,
				search:          nil,
			},
			want:    20, /* Adjust the value if required */
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListNetworkLists(tt.args.config, tt.args.includeElements, tt.args.search, tt.args.listType, tt.args.extended)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListNetworkLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("ListNetworkLists size mismatch got = %v, want %v", len(got), tt.want)
			}
		})
	}
}
