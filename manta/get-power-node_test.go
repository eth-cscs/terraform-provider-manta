package manta

import (
	"testing"
)

func TestGetPowerStatusNode(t *testing.T) {
	var w Wrapper

	_, err := w.GetPowerStatusNodeId(testXnameNode)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestGetPowerStatus404(t *testing.T) {
	var w Wrapper

	_, err := w.GetPowerStatusNodeId(testXnameNode)
	if err == nil {
		t.Errorf(`Error should not be null`)
		return
	}

	if err.Error() != "Component not found in component map." {
		t.Errorf(`error: %s`, err)
	}
}
