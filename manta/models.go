package manta

import (
	"encoding/json"
	"fmt"
)

// Wrapper is the main struct for interacting with the API
type Wrapper struct {
	Base_url             string
	Access_token         string
	access_token_content string
}

func NewWrapper(base_url, access_token string) *Wrapper {
	w := new(Wrapper)
	w.Access_token = access_token
	w.Base_url = base_url
	return w
}

type DiscoveryInfo struct {
	RedfishVersion string `json:"RedfishVersion"`
	LastAttempt    string `json:"LastAttempt"`
	LastStatus     string `json:"LastStatus"`
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

func (rfe *RfeItem) String() string {
	outjson, _ := json.MarshalIndent(rfe, "", " ")
	return string(outjson)
}

func (rfe *RfeItem) Print() {
	outjson, _ := json.MarshalIndent(rfe, "", " ")
	fmt.Println(string(outjson))
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

func (node *NodeItem) String() string {
	outjson, _ := json.MarshalIndent(node, "", " ")
	return string(outjson)
}

type BssParams struct {
	Hosts  []string `json:"hosts"`
	Macs   []string `json:"macs"`
	Nids   []string `json:"nids"`
	Params string   `json:"params"`
	Kernel string   `json:"kernel"`
	Initrd string   `json:"initrd"`
}

// TODO should I use pointer on array? like this? *[]T
func cmpArray[T comparable](arrA, arrB []T) bool {
	if len(arrA) != len(arrB) {
		return true
	}

	for i, str := range arrA {
		if str != arrB[i] {
			return true
		}
	}

	return false
}

func (bss *BssParams) Cmp(other *BssParams) bool {
	if cmpArray(bss.Hosts, other.Hosts) {
		return true
	}
	if cmpArray(bss.Macs, other.Macs) {
		return true
	}
	if cmpArray(bss.Nids, other.Nids) {
		return true
	}

	if bss.Params != other.Params {
		return true
	}

	if bss.Kernel != other.Kernel {
		return true
	}

	if bss.Initrd != other.Initrd {
		return true
	}

	return false
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

type NodeInterfacePatch struct {
	MAC string `json:"MACAddress"`
	IPs []IP   `json:"IPAddresses"`
}

type NodeInterface struct {
	ID         string `json:"ID"`
	Desc       string `json:"Description"`
	MAC        string `json:"MACAddress"`
	IPs        []IP   `json:"IPAddresses"`
	LastUpdate string `json:"LastUpdate"`
	CompID     string `json:"ComponentID"`
	Type       string `json:"Type"`
}

func (nodeInterface *NodeInterface) Cmp(other *NodeInterface) bool {
	if cmpArray(nodeInterface.IPs, other.IPs) {
		return true
	}

	if nodeInterface.ID != other.ID {
		return true
	}

	if nodeInterface.Type != other.Type {
		return true
	}

	if nodeInterface.CompID != other.CompID {
		return true
	}

	return false
}

type IP struct {
	IP string `json:"IPAddress"`
}

type OchamiError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}
