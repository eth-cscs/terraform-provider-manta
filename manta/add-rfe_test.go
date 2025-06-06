package manta

import (
	"testing"
)

func TestAddRfeFail(t *testing.T) {
	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	var rfe RfeItem = RfeItem{ID: "adsf"}

	_, err := w.AddRfe(rfe)
	if err == nil {
		return
	}

	t.Errorf(`error: add rfe should fail`)
}

func TestAddRfeSuccess(t *testing.T) {
	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	var rfe RfeItem = RfeItem{ID: testXnameRfe}

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
