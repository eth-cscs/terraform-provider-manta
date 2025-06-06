package manta

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func getrfes(url, token string) ([]RfeItem, error) {
	var getrfe RedfishEndpointArray

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		return make([]RfeItem, 0), err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return make([]RfeItem, 0), err
	}

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &getrfe)

	return getrfe.RedfishEndpoints, nil
}

func (w *Wrapper) GetRfe() ([]RfeItem, error) {
	var rfeArray []RfeItem
	var err error

	rfeArray, err = getrfes(w.base_url+"/redfish/", w.GetAccessToken())

	if err != nil {
		return make([]RfeItem, 0), err
	}

	return rfeArray, nil
}

func (w *Wrapper) GetRfeId(rfeID string) (RfeItem, error) {
	var rfeArray []RfeItem
	var err error

	rfeArray, err = getrfes(w.base_url+"/redfish/"+rfeID, w.GetAccessToken())

	if err != nil {
		return RfeItem{}, err
	}

	if len(rfeArray) != 1 {
		return RfeItem{}, errors.New("error: length of array should be one")
	}

	return rfeArray[0], nil
}
