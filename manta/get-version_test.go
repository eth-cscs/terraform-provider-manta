package manta

import (
	"testing"
)

func TestGetVersionSuccess(t *testing.T) {
	var w Wrapper = Wrapper{base_url: "http://localhost:3000"}
	const expectedVersion string = "0.1.13"

	out, err := w.Version()

	if err != nil {
		t.Errorf(`error: %s`, err)
		return
	}

	if out != expectedVersion {
		t.Errorf("error: version is incorrect\nexpected %s\nreceived %s",
			expectedVersion,
			out,
		)
		return
	}
}

func TestGetVersionFail(t *testing.T) {
	var w Wrapper = Wrapper{base_url: "http://bad-address"}

	_, err := w.Version()

	if err == nil {
		t.Errorf(`error: get version should fail`)
		return
	}
}
