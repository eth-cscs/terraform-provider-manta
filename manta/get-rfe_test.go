package manta

import (
	"testing"
)

func TestGetRfeSuccess(t *testing.T) {
	const xname string = "x0c0s0b0"

	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	var rfeitem RfeItem
	var err error

	// Create a new RFE
	rfeitem, err = w.AddRfe(RfeItem{ID: xname})

	// Get the new RFE
	rfeitem, err = w.GetRfeId(xname)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	correctRfeItem := RfeItem{
		ID:                 xname,
		Type:               "NodeBMC",
		Hostname:           "",
		Domain:             "",
		FQDN:               "x0c0s0b0",
		Enabled:            false,
		User:               "",
		Password:           "",
		RediscoverOnUpdate: false,
		DiscoveryInfo: DiscoveryInfo{
			RedfishVersion: "",
		},
	}

	if rfeitem != correctRfeItem {
		t.Errorf(`error: rfe got isn't correct`)
	}

	// Delete the new RFE
	w.DeleteRfe(xname)

	if err != nil {
		return
	}
}
