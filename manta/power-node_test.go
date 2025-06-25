package manta

import (
	"testing"
)

func TestPowerNodeBadPowerStatus(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

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
		t.Errorf("error: Node received is incorrect\nexpected: %s,\nreceived: %s\ndiff: %s",
			correctNodeItem.String(),
			out.String(),
			GetDiff(correctNodeItem, out),
		)
	}

	if err != nil {
		return
	}

	t.Errorf(`error: the function don't fail after bad power status`)
}

func testPowerNodePower(t *testing.T, powerStatus string) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

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
		t.Errorf("error: Node received is incorrect\nexpected: %s,\nreceived: %s\ndiff: %s",
			correctNodeItem.String(),
			out.String(),
			GetDiff(correctNodeItem, out),
		)
	}
}

func TestPowerNodeOff(t *testing.T) {
	testPowerNodePower(t, "Off")
}

func TestPowerNodeOn(t *testing.T) {
	testPowerNodePower(t, "On")
}
