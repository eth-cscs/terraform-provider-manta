package manta

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (w *Wrapper) AddBss(params BssParams) (BssParams, error) {
	client := &http.Client{}

	jData, err := json.Marshal(params)

	if err != nil {
		return BssParams{}, err
	}

	req, err := http.NewRequest("POST", "http://localhost:27778/boot/v1/bootparameters", bytes.NewBuffer(jData))

	if err != nil {
		return BssParams{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return BssParams{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return BssParams{}, err
	}

	if resp.StatusCode == 401 {
		return BssParams{}, errors.New(string(body))
	}

	return params, nil
}

func (w *Wrapper) UpdateBss(params BssParams) (BssParams, error) {
	client := &http.Client{}

	jData, err := json.Marshal(params)

	if err != nil {
		return BssParams{}, err
	}

	req, err := http.NewRequest("PATCH", "http://localhost:27778/boot/v1/bootparameters", bytes.NewBuffer(jData))

	if err != nil {
		return BssParams{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return BssParams{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return BssParams{}, err
	}

	if resp.StatusCode == 401 {
		return BssParams{}, errors.New(string(body))
	}

	return params, nil
}
