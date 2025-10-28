package manta

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (w *Wrapper) DeleteBss(params BssParams) (string, error) {
	client := &http.Client{}

	jData, err := json.Marshal(params)

	req, err := http.NewRequest("DELETE", "http://localhost:27778/boot/v1/bootparameters", bytes.NewBuffer(jData))
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := io.ReadAll(resp.Body)

	return string(body), nil
}
