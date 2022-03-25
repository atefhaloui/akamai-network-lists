package v2

// Link Specifies each hypermedia link.
type Link struct {
	// Any additional information about the target of the link.
	Detail string `json:"detail,omitempty"`

	// URL to access or perform the action on a related resource. May be expressed as an absolute server path, or relative to the current URL call.
	Href string `json:"href"`

	// The HTTP method with which to call the `href`, `GET` by default.
	Method string `json:"method,omitempty"`
}

// Links Encapsulates the set of [API hypermedia](#apihypermedia) to access a set of related resources. The object is arranged as a hash of keys, each of which represents a link relation.
type Links struct {
	// A link to [Activate a network list](#postactivate) in the `PRODUCTION` environment.
	ActivateInProduction Link `json:"activateInProduction,omitempty"`

	// A link to [Activate a network list](#postactivate) in the `STAGING` environment.
	ActivateInStaging Link `json:"activateInStaging,omitempty"`

	// A link to [Get activation details](#getactivationrequeststatus).
	ActivationDetails Link `json:"activationDetails,omitempty"`

	// A link to [Append elements to a network list](#postappend).
	AppendItems Link `json:"appendItems,omitempty"`

	// A link to [Get a network list](#getlist).
	Retrieve Link `json:"retrieve,omitempty"`

	// A link to [Get activation status](#getactivationstatus) for the `PRODUCTION` environment.
	StatusInProduction Link `json:"statusInProduction,omitempty"`

	// A link to [Get activation status](#getactivationstatus) for the `STAGING` environment.
	StatusInStaging Link `json:"statusInStaging,omitempty"`

	// A link to [Update a network list](#putlist).
	Update Link `json:"update,omitempty"`
}
