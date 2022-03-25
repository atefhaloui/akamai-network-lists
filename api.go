package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	log "github.com/sirupsen/logrus"
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
		retrieve           string
		append             string
		addRemove          string
		stagingStatus      string
		productionStatus   string
		activateStaging    string
		activateProduction string
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

	nls.refs.retrieve = fmt.Sprintf("/network-list/v2/network-lists/%s?extended=true&includeElements=true", nls.networkName)

	req, err = client.NewRequest(nls.config, "GET", nls.refs.retrieve, nil)
	if err != nil {
		log.Errorf("%s", err)
		return nil, CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
	if err != nil {
		log.Errorf("%s", err)
		return nil, ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s", err)
		return nil, ReadBodyFailed
	}

	log.Tracef("Received body: %s", string(data))

	if client.IsError(resp) {
		log.Errorf("%s", string(data))
		return nil, fmt.Errorf("%v", data)
	}

	if err := json.Unmarshal(data, &list); err != nil {
		log.Errorf("%s", err)
		return nil, JsonError
	}

	// Update url: append
	if list.Links.AppendItems.Href != "" {
		nls.refs.append = list.Links.AppendItems.Href
	} else {
		nls.refs.append = fmt.Sprintf("/network-list/v2/network-lists/%s/append", nls.networkName)
	}

	// Update url: status in staging
	if list.Links.StatusInStaging.Href != "" {
		nls.refs.stagingStatus = list.Links.StatusInStaging.Href
	} else {
		nls.refs.stagingStatus = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/STAGING/status", nls.networkName)
	}

	// Update url: status in production
	if list.Links.StatusInProduction.Href != "" {
		nls.refs.productionStatus = list.Links.StatusInProduction.Href
	} else {
		nls.refs.productionStatus = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/PRODUCTION/status", nls.networkName)
	}

	// Update url: activate staging
	if list.Links.ActivateInStaging.Href != "" {
		nls.refs.activateStaging = list.Links.ActivateInStaging.Href
	} else {
		nls.refs.activateStaging = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/STAGING/activate", nls.networkName)
	}

	// Update url: activate production
	if list.Links.ActivateInProduction.Href != "" {
		nls.refs.activateProduction = list.Links.ActivateInProduction.Href
	} else {
		nls.refs.activateProduction = fmt.Sprintf("/network-list/v2/network-lists/%s/environments/PRODUCTION/activate", nls.networkName)
	}

	// This is a prefix
	nls.refs.addRemove = fmt.Sprintf("/network-list/v2/network-lists/%s/elements", nls.networkName)

	log.Debug("Akamai network-list Hrefs has been fetched")

	return &nls, nil
}

// Add append an IP address, subnet or country to the network list.
func (nls *NetworkListEndpoint) Add(item string) error {
	var (
		query bytes.Buffer
		err   error
		req   *http.Request
		resp  *http.Response
		data  []byte
	)

	query.WriteString(fmt.Sprintf("%s?element=%s", nls.refs.addRemove, url.QueryEscape(item)))

	req, err = client.NewRequest(nls.config, "PUT", query.String(), nil)
	if err != nil {
		log.Errorf("%s", err)
		return CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
	if err != nil {
		log.Errorf("%s", err)
		return ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s", err)
		return ReadBodyFailed
	}

	log.Tracef("Received body: %s", string(data))

	if client.IsError(resp) {
		log.Errorf("%s", string(data))
		return fmt.Errorf("%v", data)
	}

	log.Infof("'%v' has been added", item)

	return nil
}

// Delete remove an IP address, subnet or country from the network list.
func (nls *NetworkListEndpoint) Delete(item string) error {
	var (
		query bytes.Buffer
		err   error
		req   *http.Request
		resp  *http.Response
		data  []byte
	)

	query.WriteString(fmt.Sprintf("%s?element=%s", nls.refs.addRemove, url.QueryEscape(item)))

	req, err = client.NewRequest(nls.config, "DELETE", query.String(), nil)
	if err != nil {
		log.Errorf("%s", err)
		return CreateRequestFailed
	}

	resp, err = client.Do(nls.config, req)
	if err != nil {
		log.Errorf("%s", err)
		return ExecRequestFailed
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("%s", err)
		return ReadBodyFailed
	}

	log.Tracef("Received body: %s", string(data))

	if client.IsError(resp) {
		log.Errorf("%s", string(data))
		return fmt.Errorf("%v", data)
	}

	log.Infof("'%v' has been removed", item)

	return nil
}
