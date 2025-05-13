package manta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func marshalNode(node NodeItem) string {
	str, _ := json.MarshalIndent(node, "", " ")
	return string(str)
}

func TestPowerNodeBadPowerStatus(t *testing.T) {
	var w Wrapper

	out, err := w.PowerNodeId("x1000c0s0b1n1", "BadPowerStatus")

	fmt.Println("node:", marshalNode(out))
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Errorf(`error: the function don't fail after bad power status`)
}

func TestPowerNodeOff(t *testing.T) {
	var w Wrapper

	out, err := w.PowerNodeId("x1000c0s0b1n1", "off")

	fmt.Println("node:", marshalNode(out))
	if err != nil {
		fmt.Println(err)
		t.Errorf(`error: cannot turn off`)
		return
	}
}

func TestPowerNodeOn(t *testing.T) {
	var w Wrapper

	out, err := w.PowerNodeId("x1000c0s0b1n1", "on")

	fmt.Println("node:", marshalNode(out))
	if err != nil {
		fmt.Println(err)
		t.Errorf(`error: cannot turn on`)
		return
	}
}
