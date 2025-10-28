package manta

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (w *Wrapper) GetBssParams() ([]BssParams, error) {
	var bss []BssParams

	req, err := http.NewRequest("GET", "http://localhost:27778/boot/v1/bootparameters", nil)
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return make([]BssParams, 0), err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return make([]BssParams, 0), err
	}

	if resp.StatusCode == 401 {
		return make([]BssParams, 0), errors.New(string(body))
	}

	err = json.Unmarshal(body, &bss)
	if err != nil {
		return make([]BssParams, 0), err
	}

	return bss, nil
}
