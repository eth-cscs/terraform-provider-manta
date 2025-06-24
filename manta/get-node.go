package manta

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

func (w *Wrapper) GetNodeId(id string) (NodeItem, error) {
	var node NodeItem
	var url = "https://foobar.openchami.cluster:8443/hsm/v2/State/Components/" + id

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	if err != nil {
		return NodeItem{}, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return NodeItem{}, err
	}

	err = json.Unmarshal(body, &node)

	if err != nil {
		return NodeItem{}, err
	}

	node.State, err = w.GetPowerStatusNodeId(id)

	if err != nil {
		return NodeItem{}, err
	}

	return node, err
}
