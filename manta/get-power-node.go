package manta

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (w *Wrapper) GetPowerStatusNodeId(id string) (string, error) {
	var pcs PcsStatus

	// TODO: stop using the PCS url directly
	// PCS to not update HSM about the state of the node.
	// PCS and HSM doesn't have the same information
	// export XNAME=x0c0s0b0n0
	// curl -s localhost:28007/power-status\?xname\="${XNAME}" | jq | grep 'powerState'
	// curl -sk https://foobar.openchami.cluster:8443/hsm/v2/State/Components/"${XNAME}" | jq | grep 'State'
	var url string = `http://localhost:28007/power-status?xname=` + id

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	json.Unmarshal(body, &pcs)

	return strings.Title(pcs.Status[0].PowerState), err
}
