package manta

import (
	"testing"
)

func TestDeleteRfeSuccess(t *testing.T) {
	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	const xname string = "x0c0s0b0"

	var out string
	var err error

	// Create a new RFE
	var rfe RfeItem = RfeItem{ID: xname}
	rfe, err = w.AddRfe(rfe)

	// Delete the new RFE
	out, err = w.DeleteRfe(xname)
	if err != nil {
		t.Errorf("error: delete RFE has not been successfully completed\noutput: %s\nerror: %s",
			out,
			err,
		)
	}
}
