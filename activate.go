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

// ActivateRequest Requests a new activation.
type ActivateRequest struct {
	// Descriptive text to accompany the activation. This is reflected in the [ActivationStatus](#activationstatus) object's `activationComments` member.
	Comments string `json:"comments,omitempty"`

	// List of email addresses of Portal users who receive an email when this activation of this list is complete. Do not add addresses to this list without the recipients' consent.
	NotificationRecipients []string `json:"notificationRecipients,omitempty"`

	// If the activation is linked to a Siebel ticket, this identifies the ticket.
	SiebelTicketId string `json:"siebelTicketId,omitempty"`
}

// GetActivationStatus Returns the Activation status in production or staging
func (nls *NetworkListEndpoint) GetActivationStatus(env string) (*string, error) {
	var (
		ref    string
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

	req, err = client.NewRequest(nls.config, "GET", ref, nil)
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

	if err = json.Unmarshal(data, &status); err != nil {
		log.Debugf("%s", err)
		return nil, JsonError
	}

	if env == Staging {
		log.Debugf("Staging activation status: %s", status.ActivationStatus)
		return &status.ActivationStatus, nil
	}

	log.Debugf("Production activation status: %s", status.ActivationStatus)
	return &status.ActivationStatus, nil
}

// Activate activates the Staging or Production environment
func (nls *NetworkListEndpoint) Activate(env, comment string, recipients []string) error {
	var (
		ref  string
		req  *http.Request
		resp *http.Response
		data []byte
	)

	if env == Staging {
		ref = nls.refs.stagingStatus
	} else if env == Production {
		ref = nls.refs.productionStatus
	} else {
		return UnsupportedEnvironmentError
	}

	acReq := ActivateRequest{Comments: comment, NotificationRecipients: recipients}
	body, err := json.Marshal(&acReq)
	if err != nil {
		log.Errorf("%s", err)
		return JsonError
	}

	log.Tracef("Activation Body: %s", string(body))

	req, err = client.NewRequest(nls.config, "POST", ref, bytes.NewReader(body))
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

	log.Tracef("Activation response: %s", string(data))

	if client.IsError(resp) {
		log.Errorf("%s", string(data))
		return fmt.Errorf("%v", data)
	}

	if env == Staging {
		log.Debugf("STAGING has been activated")
	} else {
		log.Debugf("PRODUCTION has been activated")
	}

	return nil
}
