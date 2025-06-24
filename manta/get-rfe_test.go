package manta

import (
	"testing"
)

func TestGetRfeSuccess(t *testing.T) {
	const xname string = "x0c0s0b0"

	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfeitem RfeItem
	var err error

	// Create a new RFE
	_, err = w.AddRfe(RfeItem{ID: xname})

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

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
	_, err = w.DeleteRfe(xname)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestGetRfeZero(t *testing.T) {
	const xname string = "x431c0s0b0" // a not existing RFE

	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfeitem RfeItem
	var err error

	// Get the new RFE
	rfeitem, err = w.GetRfeId(xname)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	correctRfeItem := RfeItem{}

	if rfeitem != correctRfeItem {
		t.Errorf(`error: rfe got isn't correct`)
	}

	if err != nil {
		return
	}
}
