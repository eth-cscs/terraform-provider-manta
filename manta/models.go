package manta

// Wrapper is the main struct for interacting with the API
type Wrapper struct {
	base_url     string
	access_token string
}

type DiscoveryInfo struct {
	RedfishVersion string `json:"RedfishVersion"`
}

type RfeItem struct {
	ID                 string        `json:"ID"`
	Type               string        `json:"Type"`
	Hostname           string        `json:"Hostname"`
	Domain             string        `json:"Domain"`
	FQDN               string        `json:"FQDN"`
	Enabled            bool          `json:"Enabled"`
	User               string        `json:"User"`
	Password           string        `json:"Password"`
	RediscoverOnUpdate bool          `json:"RediscoverOnUpdate"`
	DiscoveryInfo      DiscoveryInfo `json:"DiscoveryInfo"`
}

type NodeItem struct {
	ID      string `json:"ID"`
	Type    string `json:"Type"`
	State   string `json:"State"`
	Flag    string `json:"Flag"`
	Enabled bool   `json:"Enabled"`
	Role    string `json:"Role"`
	NID     int    `json:"NID"`
	NetType string `json:"NetType"`
	Arch    string `json:"Arch"`
	Class   string `json:"Class"`
}

type RedfishEndpointArray struct {
	RedfishEndpoints []RfeItem `json:"RedfishEndpoints"`
}

type NodeStatus struct {
	Xname                     string   `json:"xname"`
	PowerState                string   `json:"powerState"`
	ManagementState           string   `json:"managementState"`
	Error                     string   `json:"error"`
	SupportedPowerTransitions []string `json:"supportedPowerTransitions"`
	LastUpdated               string   `json:"lastUpdated"`
}

type PcsStatus struct {
	Status []NodeStatus `json:"status"`
}
