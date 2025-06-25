package manta

import (
	"testing"
)

func TestAddRfeFail(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfe = RfeItem{ID: "adsf"}

	_, err := w.AddRfe(rfe)
	if err == nil {
		return
	}

	t.Errorf(`error: add rfe should fail`)
}

func TestAddRfeSuccess(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfe = RfeItem{ID: testXnameRfe}

	_, err := w.AddRfe(rfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	out, err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf("error: delete RFE has not been successfully completed\noutput: %s\nerror: %s",
			out,
			err,
		)
	}
}

func TestAddRfeSuccessAll(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfe = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		Enabled:            true,
		User:               "user",
		Password:           "password",
		RediscoverOnUpdate: true,
	}

	// TODO password is broken

	rfeAdded, err := w.AddRfe(rfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	var rfeExpected = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		FQDN:               "hostname",
		Type:               "NodeBMC",
		Enabled:            true,
		User:               "user",
		Password:           "",
		RediscoverOnUpdate: true,
	}

	if rfeAdded != rfeExpected {
		t.Errorf("error: rfe added is not as expected\nadded: %s\nexpected: %s\ndiff: %s",
			rfeAdded.String(),
			rfeExpected.String(),
			GetDiff(rfeAdded, rfeExpected),
		)
	}

	out, err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf("error: delete RFE has not been successfully completed\noutput: %s\nerror: %s",
			out,
			err,
		)
	}
}
