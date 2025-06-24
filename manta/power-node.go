package manta

import (
	"errors"
	"net/http"
)

func (w *Wrapper) PowerNodeId(id, powerStatus string) (NodeItem, error) {
	var state string

	switch powerStatus {
	case "Off":
		state = "power-off"
	case "On":
		state = "power-on"
	default:
		return NodeItem{}, errors.New(`powerStatus ins't "Off" or "On"`)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", w.Base_url+"/node/"+id+"/"+state, nil)

	if err != nil {
		return NodeItem{}, err
	}

	req.Header.Set("Authorization", "Bearer "+w.GetAccessToken())

	_, err = client.Do(req)
	if err != nil {
		return NodeItem{}, err
	}

	return w.GetNodeId(id)
}
