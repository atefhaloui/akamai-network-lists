package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
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
		return JsonError
	}

	req, err = client.NewRequest(nls.config, nls.refs.append.Method, nls.refs.append.Href, bytes.NewReader(body))
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
