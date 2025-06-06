package manta

import (
	"encoding/json"
	"testing"
)

func marshalNode(node NodeItem) string {
	str, _ := json.MarshalIndent(node, "", " ")
	return string(str)
}

func TestPowerNodeBadPowerStatus(t *testing.T) {
	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	out, err := w.PowerNodeId(testXnameNode, "BadPowerStatus")

	correctNodeItem := NodeItem{
		ID:      "",
		Type:    "",
		State:   "",
		Flag:    "",
		Enabled: false,
		Role:    "",
		NID:     0,
		NetType: "",
		Arch:    "",
		Class:   "",
	}

	if out != correctNodeItem {
		t.Errorf("error: Node received is incorrect\nexpected: %s,\nreceived: %s",
			correctNodeItem.String(),
			out.String(),
		)
	}

	if err != nil {
		return
	}

	t.Errorf(`error: the function don't fail after bad power status`)
}

func testPowerNodePower(t *testing.T, powerStatus string) {
	var w Wrapper = Wrapper{access_token: "~/access_token", base_url: "http://localhost:3000"}

	out, err := w.PowerNodeId(testXnameNode, powerStatus)

	correctNodeItem := NodeItem{
		ID:      testXnameNode,
		Type:    "Node",
		State:   powerStatus,
		Flag:    "OK",
		Enabled: true,
		Role:    "Compute",
		NID:     16400389,
		NetType: "Sling",
		Arch:    "X86",
		Class:   "River",
	}

	if err != nil {
		t.Errorf(`error: %s`, err)
		return
	}

	if out != correctNodeItem {
		t.Errorf("error: Node received is incorrect\nexpected: %s,\nreceived: %s",
			correctNodeItem.String(),
			out.String(),
		)
	}
}

func TestPowerNodeOff(t *testing.T) {
	testPowerNodePower(t, "Off")
}

func TestPowerNodeOn(t *testing.T) {
	testPowerNodePower(t, "On")
}
