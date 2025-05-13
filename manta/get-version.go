package manta

import (
	"io"
	"net/http"
)

// Return the version of the Manta WS
func (w *Wrapper) Version() (string, error) {
	// Authorization: Bearer $(<access_token)"
	// req.Header.Set("name", "value")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", w.base_url+"/version", nil)
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
