package manta

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func (w *Wrapper) AddRfe(rfeItem RfeItem) (RfeItem, error) {
	var rfes RedfishEndpointArray

	rfes.RedfishEndpoints = append(rfes.RedfishEndpoints, rfeItem)

	client := &http.Client{}

	jData, err := json.Marshal(rfes)
	if err != nil {
		return RfeItem{}, err
	}

	req, err := http.NewRequest("POST", w.Base_url+"/redfish", bytes.NewBuffer(jData))

	if err != nil {
		return RfeItem{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return RfeItem{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return RfeItem{}, err
	}

	// if resp.StatusCode != 201 -> handle error

	if string(body) == `"ERROR - Message: OCHAMI-RS: {\"type\":\"about:blank\",\"title\":\"Conflict\",\"detail\":\"operation would conflict with an existing resource that has the same FQDN or xname ID.\",\"status\":409}\n"` {
		return RfeItem{}, errors.New("rfe item already exist")
	}

	rfeReturn, _ := w.GetRfeId(rfeItem.ID)

	if rfeReturn.Enabled {
		for rfeReturn.DiscoveryInfo.LastStatus != "DiscoverOK" {
			if rfeReturn.DiscoveryInfo.LastStatus == "ChildVerificationFailed" {
				break
			}
			if rfeReturn.DiscoveryInfo.LastStatus == "HTTPsGetFailed" {
				break
			}
			time.Sleep(1 * time.Second)
			rfeReturn, _ = w.GetRfeId(rfeReturn.ID)
		}
	}

	return rfeReturn, nil
}
