package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (w *Wrapper) requestTls(url, method string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	//req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *Wrapper) requestJsonTls(url, method string, data []byte) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *Wrapper) DeleteEthernetInterface(mac string) error {
	ethInterfaceID := strings.Replace(mac, ":", "", -1)

	resp, err := w.requestTls(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		"DELETE",
	)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var ochamiError OchamiError

	err = json.Unmarshal(body, &ochamiError)

	if err != nil {
		return err
	}
	if ochamiError.Status == 404 {
		return errors.New(ochamiError.Detail)
	}

	return nil
}

func (w *Wrapper) GetEthernetInterface(mac string) (NodeInterface, error) {
	ethInterfaceID := strings.Replace(mac, ":", "", -1)

	resp, err := w.requestTls(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		"GET",
	)

	if err != nil {
		return NodeInterface{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return NodeInterface{}, errors.New(string(body))
	}

	if err != nil {
		return NodeInterface{}, err
	}

	var comp NodeInterface

	err = json.Unmarshal(body, &comp)

	if err != nil {
		return NodeInterface{}, err
	}

	return comp, nil
}

func (w *Wrapper) AddEthernetInterface(nodeInterface NodeInterface) error {
	_, err := w.GetEthernetInterface(nodeInterface.MAC)

	if err != nil {
		var ochamiError OchamiError

		errUnmarshal := json.Unmarshal([]byte(err.Error()), &ochamiError)

		if errUnmarshal != nil {
			return errUnmarshal
		}

		// not found means not already existing
		if ochamiError.Status != 404 {
			return err
		}
	} else {
		// patch already existing interface
		_, err := w.PatchEthernetInterface(
			NodeInterfacePatch{
				MAC: nodeInterface.MAC,
				IPs: nodeInterface.IPs,
			},
		)
		return err
	}

	jData, err := json.Marshal(nodeInterface)

	if err != nil {
		return err
	}

	resp, err := w.requestJsonTls(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/",
		"POST",
		jData,
	)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New(string(body))
	}

	var comp NodeInterface

	err = json.Unmarshal(body, &comp)

	if err != nil {
		return err
	}

	return nil
}

func (w *Wrapper) PatchEthernetInterface(nodeInterface NodeInterfacePatch) (NodeInterface, error) {
	jData, err := json.Marshal(nodeInterface)

	if err != nil {
		return NodeInterface{}, err
	}

	ethInterfaceID := strings.Replace(nodeInterface.MAC, ":", "", -1)
	resp, err := w.requestJsonTls(
		"https://foobar.openchami.cluster:8443/hsm/v2/Inventory/EthernetInterfaces/"+ethInterfaceID,
		"PATCH",
		jData,
	)

	if err != nil {
		return NodeInterface{}, err
	}

	var comp NodeInterface

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return NodeInterface{}, err
	}

	if resp.StatusCode != 200 {
		return NodeInterface{}, errors.New(string(body))
	}

	err = json.Unmarshal(body, &comp)

	if err != nil {
		return NodeInterface{}, err
	}

	return comp, nil
}
