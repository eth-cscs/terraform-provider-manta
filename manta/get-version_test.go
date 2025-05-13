package manta

import (
	"fmt"
	"testing"
)

func TestGetVersionSuccess(t *testing.T) {
	var w Wrapper = Wrapper{base_url: "http://localhost:3000"}

	out, err := w.Version()

	if err != nil {
		t.Errorf(`error: %s`, err)
		return
	}

	fmt.Println("version: " + out)
}

func TestGetVersionFail(t *testing.T) {
	var w Wrapper = Wrapper{base_url: "http://bad-address"}

	_, err := w.Version()

	if err == nil {
		t.Errorf(`error: get version should fail`)
		return
	}
}
