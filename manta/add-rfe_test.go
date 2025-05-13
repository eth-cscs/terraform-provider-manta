package manta

import (
	"fmt"
	"testing"
)

func TestAddRfeSuccess(t *testing.T) {
	var w Wrapper

	var rfe RfeItem = RfeItem{ID: "adsf"}

	_, err := w.AddRfe(rfe)
	if err != nil {
		t.Errorf(`error: err`)
	}

	fmt.Println(err)
}
