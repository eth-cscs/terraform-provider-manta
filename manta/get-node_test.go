package manta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetNodeSuccess(t *testing.T) {
	var w Wrapper

	out, err := w.GetNodeId("x1000c0s0b1n1")

	if err != nil {
		return
	}

	fmt.Println(out)
	str, _ := json.MarshalIndent(out, "", " ")
	fmt.Println(string(str))
}
