package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	log "github.com/sirupsen/logrus"
)

// AppendRequest definition
type AppendRequest struct {
	List []string `json:"list"`
}

// Append Add a new IP addresses, subnets or country to the network list.
func (nls *NetworkListEndpoint) Append(list []string) error {
	var (
		req  *http.Request
		resp *http.Response
		data []byte
	)

	appReq := AppendRequest{list}
	body, err := json.Marshal(appReq)
	if err != nil {
		log.Errorf("%s", err)
		return JsonError
	}

	req, err = client.NewRequest(nls.config, "POST", nls.refs.append, bytes.NewReader(body))
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

	log.Infof("'%v' had been added", list)

	return nil
}
