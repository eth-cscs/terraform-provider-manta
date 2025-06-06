package manta

import (
	"testing"
)

func TestGetNodeSuccess(t *testing.T) {
	var w Wrapper

	out, err := w.GetNodeId(testXnameNode)

	out.State = "On"

	correctNodeItem := NodeItem{
		ID:      testXnameNode,
		Type:    "Node",
		State:   "On",
		Flag:    "OK",
		Enabled: true,
		Role:    "Compute",
		NID:     16400389,
		NetType: "Sling",
		Arch:    "X86",
		Class:   "River",
	}

	if err != nil {
		t.Errorf(`error: Get node should has failed`)
	}

	if out != correctNodeItem {
		t.Errorf(`error: Node received is incorrect`)
	}
}
