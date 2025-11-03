package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func getComponent(url, token string) (Components, error) {
	components := Components{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return components, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return components, err
	}

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &components)

	if err != nil {
		return components, err
	}

	if err != nil {
		return components, err
	}

	return components, nil
}

func requestComponent(url, token, method string, components Components) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(components)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jData))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	if err != nil {
		return err
	}

	return nil
}

func deleteComponent(url, token string, newComponents Components) error {
	return requestComponent(url, token, "DELETE", newComponents)
}

func createComponent(url, token string, deleteComponents Components) error {
	return requestComponent(url, token, "POST", deleteComponents)
}

func (w *Wrapper) CreateComponent(newComponents Components) error {
	return createComponent(
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
		newComponents,
	)
}

func (w *Wrapper) GetComponent() (Components, error) {
	return getComponent(
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
	)
}

func (w *Wrapper) DeleteComponent(deleteComponents Components) error {
	return deleteComponent(
		"https://foobar.openchami.cluster:8443/hsm/v2/State/Components",
		w.GetAccessToken(),
		deleteComponents,
	)
}
