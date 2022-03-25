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
