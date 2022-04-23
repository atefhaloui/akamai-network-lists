package v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

// NetworkList Encapsulates information about each network list.
type NetworkList struct {
	// Name of this network list's access control group (ACG).
	AccessControlGroup string `json:"accessControlGroup,omitempty"`

	// Encapsulates the set of [API hypermedia](#apihypermedia) to access a set of resources related to this network list. The object is arranged as a hash of keys, each of which represents a link relation.
	Links Links `json:"links,omitempty"`

	// ISO 8601 timestamp indicating when the network list was first created. Available only when using the `extended` query parameter to retrieve network list data.
	CreateDate string `json:"createDate,omitempty"`

	// Username of this list's creator. Available only when using the `extended` query parameter to retrieve network list data.
	CreatedBy string `json:"createdBy,omitempty"`

	// Detailed description of the list.
	Description string `json:"description,omitempty"`

	// Reflects the number of elements in the `list` array, which may not necessarily appear in the object when retrieving the list with the `includeElements` query parameter set to `false`.
	ElementCount int `json:"elementCount,omitempty"`

	// For clients with access to _expedited_ activations on select servers, provides the most recent activation status in the `PRODUCTION` environment. See [Activation States](#activationvalues) for details on each activation state. Available only when using the `extended` query parameter to retrieve network list data.
	ExpeditedProductionActivationStatus string `json:"expeditedProductionActivationStatus,omitempty"`

	// For clients with access to _expedited_ activations on select servers, provides the most recent activation status in the `STAGING` environment. See [Activation States](#activationvalues) for details on each activation state. Available only when using the `extended` query parameter to retrieve network list data.
	ExpeditedStagingActivationStatus string `json:"expeditedStagingActivationStatus,omitempty"`

	// List of IPs or Countries
	List []string `json:"list,omitempty"`

	// Display name of the network list.
	Name string `json:"name"`

	// If set to `extendedNetworkListResponse`, indicates that the current data features members enabled with the `extended` query parameter. Otherwise a plain `networkListResponse` value indicates this additional data is absent.
	NetworkListType string `json:"networkListType,omitempty"`

	// The most recent activation status of the current list in the `PRODUCTION` environment.  See [Activation States](#activationvalues) for details on each activation state. Available only when using the `extended` query parameter to retrieve network list data.
	ProductionActivationStatus string `json:"productionActivationStatus,omitempty"`

	// If `true`, indicates that you do not have permission to modify the network list. This may indicate either a network list that Akamai manages, or insufficient permission for your API client's identity to modify a customer-managed list. The default value is `false`.
	ReadOnly bool `json:"readOnly,omitempty"`

	// If `true`, indicates that this list has been shared with you by Akamai or some other account. The default value is `false`. Shared lists are always read only
	Shared bool `json:"shared,omitempty"`

	// The most recent activation status of the current list in the `STAGING` environment. See [Activation States](#activationvalues) for details on each activation state. Available only when using the `extended` query parameter to retrieve network list data.
	StagingActivationStatus string `json:"stagingActivationStatus,omitempty"`

	// Identifies each version of the network list, which increments each time it's modified. You need to include this value in any requests to modify the list. See [Concurrency control](#concurrency) for details.
	SyncPoint uint `json:"syncPoint"`

	// The network list type, either `IP` for IP addresses and CIDR blocks, or `GEO` for two-letter country codes.
	Type string `json:"type"`

	// A unique identifier for each network list, corresponding to the `networkListId` URL parameter.
	UniqueId string `json:"uniqueId"`

	// ISO 8601 timestamp indicating when the network list was last modified. Available only when using the `extended` query parameter to retrieve network list data.
	UpdateDate string `json:"updateDate,omitempty"`

	// Username of this list's creator. Available only when using the `extended` query parameter to retrieve network list data.
	UpdatedBy string `json:"updatedBy,omitempty"`
}

// ListNetworkLists Get all network list according to search criteria.
// [API](https://techdocs.akamai.com/network-lists/reference/get-network-lists)
func ListNetworkLists(config edgegrid.Config, includeElements bool, search *string, listType *string, extended bool) ([]NetworkList, error) {
	type listResponse struct {
		Links struct {
			Create struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"create"`
		} `json:"links"`
		NetworkLists []NetworkList `json:"networkLists"`
	}

	var (
		err     error
		req     *http.Request
		resp    *http.Response
		data    []byte
		query   string
		wrapper listResponse
	)

	query = fmt.Sprintf("/network-list/v2/network-lists?extended=%v&includeElements=%v", extended, includeElements)
	if search != nil {
		query = query + "&search=" + url.QueryEscape(*search)
	}
	if listType != nil {
		query = query + "&listType=" + *listType
	}

	req, err = client.NewRequest(config, "GET", query, nil)
	if err != nil {
		return nil, CreateRequestFailed
	}

	resp, err = client.Do(config, req)
	if err != nil {
		return nil, ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ReadBodyFailed
	}

	if client.IsError(resp) {
		return nil, fmt.Errorf("%s", string(data))
	}

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, JsonError
	}

	return wrapper.NetworkLists, nil
}
