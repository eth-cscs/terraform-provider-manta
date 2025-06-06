package manta

import (
	"io"
	"net/http"
)

func (w *Wrapper) DeleteRfe(rfeID string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", w.base_url+"/redfish/"+rfeID, nil)
	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	return string(body), nil
}
