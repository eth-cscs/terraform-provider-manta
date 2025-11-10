package manta

import (
	"fmt"
	"testing"
)

func TestAddRfeFail(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfe = RfeItem{ID: "adsf"}

	_, err := w.AddRfe(rfe)
	if err == nil {
		t.Errorf(`error: add rfe should fail`)
	}

}

func testAddRfeSuccess(t *testing.T, added, expected RfeItem) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	added, err = w.AddRfe(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if added != expected {
		t.Errorf("error: rfe added is not as expected\nadded: %s\nexpected: %s\ndiff: %s",
			added.String(),
			expected.String(),
			GetDiff(added, expected),
		)
	}

	err = w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestAddTwiceRfe(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}
	var added = RfeItem{ID: testXnameRfe}

	err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	added, err = w.AddRfe(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
	added, err = w.AddRfe(added)
	if err == nil {
		t.Errorf(`error: add rfe should fail`)
	}

	err = w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestAddRfeSuccess(t *testing.T) {
	var added = RfeItem{ID: testXnameRfe}
	var expected = RfeItem{ID: testXnameRfe, Type: "NodeBMC", FQDN: testXnameRfe}
	testAddRfeSuccess(t, added, expected)
}

func TestAddRfeSuccessAll(t *testing.T) {
	var added = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		Enabled:            false,
		User:               "user",
		Password:           "password",
		RediscoverOnUpdate: true,
	}
	var expected = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		FQDN:               "hostname",
		Type:               "NodeBMC",
		Enabled:            false,
		User:               "user",
		Password:           "",
		RediscoverOnUpdate: true,
	}

	testAddRfeSuccess(t, added, expected)
}

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
		fmt.Println(rfeitem)
		fmt.Println(correctRfeItem)
	}

	// Delete the new RFE
	err = w.DeleteRfe(xname)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestGetRfeZero(t *testing.T) {
	const xname string = "x431c0s0b0" // a not existing RFE

	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var err error

	// Get the new RFE
	_, err = w.GetRfeId(xname)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestDeleteRfeSuccess(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	const xname string = "x0c0s0b0"

	var out string
	var err error

	// Create a new RFE
	var rfe = RfeItem{ID: xname}
	_, err = w.AddRfe(rfe)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	// Delete the new RFE
	err = w.DeleteRfe(xname)
	if err != nil {
		t.Errorf("error: delete RFE has not been successfully completed\noutput: %s\nerror: %s",
			out,
			err,
		)
	}
}
