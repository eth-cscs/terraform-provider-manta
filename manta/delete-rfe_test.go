package manta

import (
	"testing"
)

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
	out, err = w.DeleteRfe(xname)
	if err != nil {
		t.Errorf("error: delete RFE has not been successfully completed\noutput: %s\nerror: %s",
			out,
			err,
		)
	}
}
