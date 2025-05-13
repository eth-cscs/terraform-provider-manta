package manta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetRfeSuccess(t *testing.T) {
	var w Wrapper

	out, _ := w.GetRfe()

	fmt.Println(out)
	str, _ := json.MarshalIndent(out, "", " ")
	fmt.Println(string(str))

	for _, rfe := range out {
		fmt.Println(rfe)
	}
}
