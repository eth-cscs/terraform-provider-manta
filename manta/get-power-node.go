package manta

import (
	"encoding/json"
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"net/http"
)

func (w *Wrapper) GetPowerStatusNodeId(id string) (string, error) {
	var pcs PcsStatus

	// TODO: stop using the PCS url directly
	// PCS to not update HSM about the state of the node.
	// PCS and HSM doesn't have the same information
	// export XNAME=x0c0s0b0n0
	// curl -s localhost:28007/power-status\?xname\="${XNAME}" | jq | grep 'powerState'
	// curl -sk https://foobar.openchami.cluster:8443/hsm/v2/State/Components/"${XNAME}" | jq | grep 'State'
	var url = `http://localhost:28007/power-status?xname=` + id

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &pcs)

	if err != nil {
		return "", err
	}

	if pcs.Status[0].Error == `Component not found in component map.` {
		return "", errors.New(pcs.Status[0].Error)
	}

	var c = cases.Title(language.English)
	return c.String(pcs.Status[0].PowerState), err
}
