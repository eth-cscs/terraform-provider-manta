package manta

import (
	"encoding/json"
)

type Components struct {
	Components []Component `json:"Components"`
}

type Component struct {
	ID                  string      `json:"ID"`
	Type                string      `json:"Type"`
	State               string      `json:"State,omitempty"`
	Flag                string      `json:"Flag,omitempty"`
	Enabled             *bool       `json:"Enabled,omitempty"`
	SwStatus            string      `json:"SoftwareStatus,omitempty"`
	Role                string      `json:"Role,omitempty"`
	SubRole             string      `json:"SubRole,omitempty"`
	NID                 json.Number `json:"NID,omitempty"`
	Subtype             string      `json:"Subtype,omitempty"`
	NetType             string      `json:"NetType,omitempty"`
	Arch                string      `json:"Arch,omitempty"`
	Class               string      `json:"Class,omitempty"`
	ReservationDisabled bool        `json:"ReservationDisabled,omitempty"`
	Locked              bool        `json:"Locked,omitempty"`
}

func (w *Wrapper) CreateComponent(newComponents Components) error {
	return requestPost(
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
		newComponents,
	)
}

func (w *Wrapper) GetComponent() (Components, error) {
	return requestGet[Components](
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
	)
}

func (w *Wrapper) DeleteComponent(deleteComponents Components) error {
	return requestDeleteBody(
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
		deleteComponents,
	)
}
