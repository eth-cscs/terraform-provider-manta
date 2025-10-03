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
