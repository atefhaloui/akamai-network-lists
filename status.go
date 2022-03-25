package v2

// Status Represents a network list's activation status, as the response to the
// [Get activation status](#getactivationstatus) operation.
// (This object may also appear within an outer [ActivationDetails](#activationdetails) envelope.)
type Status struct {
	// Further information related to the activation. This reflects the `comments` from the initial [ActivationRequest](#activationrequest).
	ActivationComments string `json:"activationComments,omitempty"`

	// Unique identifier for the most recent activation request. May be absent if the list is inactive, or if this object in embedded within an outer [ActivationDetails](#activationdetails) envelope.
	ActivationId int `json:"activationId,omitempty"`

	// This network list's current activation status in the specified `environment`. See [Activation States](#activationvalues) for details on each activation state.
	ActivationStatus string `json:"activationStatus"`

	// Number of times we have attempted to deploy this security configuration on the network.
	DispatchCount int `json:"dispatchCount,omitempty"`

	// True when _fast metadata_ activation is in effect. False when this is the first deployment of a network list to an environment using _fast metadata_ activation.
	Fast bool `json:"fast,omitempty"`

	// Encapsulates the set of [API hypermedia](#apihypermedia) to access a set of resources related to this activation. The object is arranged as a hash of keys, each of which represents a link relation.
	Links *Links `json:"links,omitempty"`

	// The version of the currently activated network list. See [Concurrency control](#concurrency) for details.
	SyncPoint int `json:"syncPoint"`

	// Unique identifier for this network list, corresponding to the `networkListId` URL parameter.
	UniqueId string `json:"uniqueId"`
}

// Check if the activationStatus is in enumeration
func isValidStatus(v string) bool {
	if v != "INACTIVE" &&
		v != "ACTIVE" &&
		v != "MODIFIED" &&
		v != "PENDING_ACTIVATION" &&
		v != "FAILED" &&
		v != "PENDING_DEACTIVATION" {
		return false
	}

	return true
}

/*
func (s *Status) unmarshal(b []byte) error {
	activationStatusReceived := false
	syncPointReceived := false
	uniqueIdReceived := false
	var jsonMap map[string]json.RawMessage

	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}

	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "activationComments":
			if err := json.Unmarshal(v, &s.ActivationComments); err != nil {
				return err
			}
		case "activationId":
			if err := json.Unmarshal(v, &s.ActivationId); err != nil {
				return err
			}
		case "activationStatus":
			if err := json.Unmarshal(v, &s.ActivationStatus); err != nil {
				return err
			}
			if isValidStatus(s.ActivationStatus) {
				return errors.New("\"activationStatus\" value is not recognized")
			}
			activationStatusReceived = true
		case "dispatchCount":
			if err := json.Unmarshal(v, &s.DispatchCount); err != nil {
				return err
			}
		case "fast":
			if err := json.Unmarshal(v, &s.Fast); err != nil {
				return err
			}
		case "links":
			if err := s.Links.unmarshal(v); err != nil {
				return err
			}
		case "syncPoint":
			if err := json.Unmarshal(v, &s.SyncPoint); err != nil {
				return err
			}
			syncPointReceived = true
		case "uniqueId":
			if err := json.Unmarshal(v, &s.UniqueId); err != nil {
				return err
			}
			uniqueIdReceived = true
		}
	}
	// check if activationStatus (a required property) was received
	if !activationStatusReceived {
		return errors.New("\"activationStatus\" is required but was not present")
	}
	// check if syncPoint (a required property) was received
	if !syncPointReceived {
		return errors.New("\"syncPoint\" is required but was not present")
	}
	// check if uniqueId (a required property) was received
	if !uniqueIdReceived {
		return errors.New("\"uniqueId\" is required but was not present")
	}

	return nil
}
*/
