package manta

import (
	"testing"
)

func TestGetPowerStatusNode(t *testing.T) {
	var w Wrapper

	// TODO make a better test
	_, err := w.GetPowerStatusNodeId(testXnameNode)

	if err != nil {
		return
	}
}
