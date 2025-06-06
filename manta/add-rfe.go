package manta

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (w *Wrapper) AddRfe(rfeItem RfeItem) (RfeItem, error) {
	var rfes RedfishEndpointArray

	rfes.RedfishEndpoints = append(rfes.RedfishEndpoints, rfeItem)

	client := &http.Client{}

	jData, err := json.Marshal(rfes)
	if err != nil {
		return RfeItem{}, err
	}

	req, err := http.NewRequest("POST", w.base_url+"/redfish", bytes.NewBuffer(jData))
	if err != nil {
		return RfeItem{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return RfeItem{}, err
	}

	_, err = io.ReadAll(resp.Body)

	rfeReturn, _ := w.GetRfeId(rfeItem.ID)

	return rfeReturn, nil
}
