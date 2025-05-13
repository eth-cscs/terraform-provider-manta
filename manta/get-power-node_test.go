package manta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetPowerStatusNode(t *testing.T) {
	var w Wrapper

	out, err := w.GetPowerStatusNodeId("x1000c0s0b1n1")

	if err != nil {
		return
	}

	str, _ := json.MarshalIndent(out, "", " ")
	fmt.Println(string(str))
	fmt.Println(out.PowerState)
}
