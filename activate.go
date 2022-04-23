package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

// ActivateRequest Requests a new activation.
type ActivateRequest struct {
	// Descriptive text to accompany the activation. This is reflected in the [ActivationStatus](#activationstatus) object's `activationComments` member.
	Comments string `json:"comments,omitempty"`

	// List of email addresses of Portal users who receive an email when this activation of this list is complete. Do not add addresses to this list without the recipients' consent.
	NotificationRecipients []string `json:"notificationRecipients,omitempty"`

	// If the activation is linked to a Siebel ticket, this identifies the ticket.
	SiebelTicketId string `json:"siebelTicketId,omitempty"`
}

// GetActivationStatus Returns the Activation status in production or staging.
func (nls *NetworkListEndpoint) GetActivationStatus(env string) (*string, error) {
	var (
		ref    Link
		err    error
		req    *http.Request
		resp   *http.Response
		data   []byte
		status Status
	)

	if env == Staging {
		ref = nls.refs.stagingStatus
	} else if env == Production {
		ref = nls.refs.productionStatus
	} else {
		return nil, UnsupportedEnvironmentError
	}

	req, err = client.NewRequest(nls.config, ref.Method, ref.Href, nil)
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

	if err = json.Unmarshal(data, &status); err != nil {
		return nil, JsonError
	}

	if env == Staging {
		return &status.ActivationStatus, nil
	}

	return &status.ActivationStatus, nil
}

// Activate activates the Staging or Production environment
func (nls *NetworkListEndpoint) Activate(env, comment string, recipients []string) error {
	var (
		ref  Link
		req  *http.Request
		resp *http.Response
		data []byte
	)

	if env == Staging {
		ref = nls.refs.activateStaging
	} else if env == Production {
		ref = nls.refs.activateProduction
	} else {
		return UnsupportedEnvironmentError
	}

	acReq := ActivateRequest{Comments: comment, NotificationRecipients: recipients}
	body, err := json.Marshal(&acReq)
	if err != nil {
		return JsonError
	}

	req, err = client.NewRequest(nls.config, ref.Method, ref.Href, bytes.NewReader(body))
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
