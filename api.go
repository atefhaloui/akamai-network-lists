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

const (
	Staging    = "staging"
	Production = "production"
)

// NetworkListEndpoint Network List API parameters.
type NetworkListEndpoint struct {
	config      edgegrid.Config
	networkName string
	refs        struct {
		append             Link
		add                Link
		remove             Link
		stagingStatus      Link
		productionStatus   Link
		activateStaging    Link
		activateProduction Link
	}
}

// New init and crete a new Network List endpoint.
func New(cfg *edgegrid.Config, networkName string) (*NetworkListEndpoint, error) {
	var (
		nls  NetworkListEndpoint
		err  error
		req  *http.Request
		resp *http.Response
		data []byte
		list NetworkList
	)

	nls.config = *cfg
	nls.networkName = networkName

	retrieve := fmt.Sprintf("/network-list/v2/network-lists/%s?extended=false&includeElements=false", nls.networkName)

	req, err = client.NewRequest(nls.config, "GET", retrieve, nil)
	if err != nil {
		return nil, CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
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

	if err := json.Unmarshal(data, &list); err != nil {
		return nil, JsonError
	}

	// Update url: append
	if list.Links.AppendItems.Href != "" {
		nls.refs.append = list.Links.AppendItems
	} else {
		nls.refs.append.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/append", nls.networkName)
		nls.refs.append.Method = "POST"
	}

	// Update url: status in staging
	if list.Links.StatusInStaging.Href != "" {
		nls.refs.stagingStatus = list.Links.StatusInStaging
	} else {
		nls.refs.stagingStatus.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/STAGING/status", nls.networkName)
		nls.refs.stagingStatus.Method = "GET"
	}

	// Update url: status in production
	if list.Links.StatusInProduction.Href != "" {
		nls.refs.productionStatus = list.Links.StatusInProduction
	} else {
		nls.refs.productionStatus.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/PRODUCTION/status", nls.networkName)
		nls.refs.productionStatus.Method = "GET"
	}

	// Update url: activate staging
	if list.Links.ActivateInStaging.Href != "" {
		nls.refs.activateStaging = list.Links.ActivateInStaging
	} else {
		nls.refs.activateStaging.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/STAGING/activate", nls.networkName)
		nls.refs.activateStaging.Method = "POST"
	}

	// Update url: activate production
	if list.Links.ActivateInProduction.Href != "" {
		nls.refs.activateProduction = list.Links.ActivateInProduction
	} else {
		nls.refs.activateProduction.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/PRODUCTION/activate", nls.networkName)
		nls.refs.activateStaging.Method = "POST"
	}

	// This is a prefix
	nls.refs.add.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/elements", nls.networkName)
	nls.refs.add.Method = "PUT"

	nls.refs.remove.Href = fmt.Sprintf("/network-list/v2/network-lists/%s/elements", nls.networkName)
	nls.refs.remove.Method = "DELETE"

	return &nls, nil
}

// Add append an IP address, subnet or country to the network list.
func (nls *NetworkListEndpoint) Add(item string) error {
	var (
		err   error
		req   *http.Request
		resp  *http.Response
		data  []byte
	)

	query := fmt.Sprintf("%s?element=%s", nls.refs.add.Href, url.QueryEscape(item))

	req, err = client.NewRequest(nls.config, nls.refs.add.Method, query, nil)
	if err != nil {
		return CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
	if err != nil {
		return ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ReadBodyFailed
	}

	if client.IsError(resp) {
		return fmt.Errorf("%v", string(data))
	}

	return nil
}

// Delete remove an IP address, subnet or country from the network list.
func (nls *NetworkListEndpoint) Delete(item string) error {
	var (
		err   error
		req   *http.Request
		resp  *http.Response
		data  []byte
	)

	query := fmt.Sprintf("%s?element=%s", nls.refs.remove.Href, url.QueryEscape(item))

	req, err = client.NewRequest(nls.config, nls.refs.remove.Method, query, nil)
	if err != nil {
		return CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
	if err != nil {
		return ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ReadBodyFailed
	}

	if client.IsError(resp) {
		return fmt.Errorf("%s", string(data))
	}

	return nil
}
