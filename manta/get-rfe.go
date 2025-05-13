package manta

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func (w *Wrapper) GetRfe() ([]RfeItem, error) {
	var stdout bytes.Buffer
	var getrfe RedfishEndpointArray

	cmd := exec.Command("manta", "get", "redfish-endpoints")

	cmd.Stdout = &stdout

	cmd.Run()
	json.Unmarshal(stdout.Bytes(), &getrfe)

	return getrfe.RedfishEndpoints, nil
}

func (w *Wrapper) GetRfeId(id string) (RfeItem, error) {
	var stdout bytes.Buffer
	var getrfe RedfishEndpointArray

	cmd := exec.Command("manta", "get", "redfish-endpoints", "--id", id)

	cmd.Stdout = &stdout

	cmd.Run()

	json.Unmarshal(stdout.Bytes(), &getrfe)

	return getrfe.RedfishEndpoints[0], nil
}
