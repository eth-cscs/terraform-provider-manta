package manta

import (
	"encoding/json"
	"fmt"
	"testing"
)

func testDeleteEthernetInterface(t *testing.T, mac string) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	w.DeleteEthernetInterface(mac)
}

func TestDeleteEthernetInterface404(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteEthernetInterface("aa:aa:aa:aa:aa:aa")

	if err == nil {
		t.Errorf(`error: missing error`)
	}

	if err.Error() != "no such component ethernet interface." {
		t.Errorf(`error: %s`, err)
	}
}

func testGetEthernetInterface(t *testing.T, mac string, expected NodeInterface) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	comp, err := w.GetEthernetInterface(mac)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if comp.Cmp(&expected) {
		{
			node, _ := json.Marshal(comp)
			fmt.Println("received: " + string(node))
		}
		{
			node, _ := json.Marshal(expected)
			fmt.Println("expected: " + string(node))
		}
		t.Errorf(`error: not equal`)
	}
}

func TestGetEthernetInterfaceBlank(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	comp, err := w.GetEthernetInterface("aa:aa:aa:aa:aa:aa")
	var expected = NodeInterface{}

	if comp.Cmp(&expected) {
		{
			node, _ := json.Marshal(comp)
			fmt.Println("received: " + string(node))
		}
		{
			node, _ := json.Marshal(expected)
			fmt.Println("expected: " + string(node))
		}
		t.Errorf(`error: not equal`)
	}
	if err == nil {
		t.Errorf(`error: missing error`)
	}
}

func testAddEthernetInterface(t *testing.T, added NodeInterface) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.AddEthernetInterface(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestEthernetInterface(t *testing.T) {
	var added = NodeInterface{
		MAC: "00:de:ad:be:ef:00",
		IPs: []IP{IP{"10.0.0.5"}},
	}
	var patch = NodeInterfacePatch{
		MAC: "00:de:ad:be:ef:00",
		IPs: []IP{IP{"10.0.0.0"}},
	}

	testAddEthernetInterface(t, added)
	added.ID = "00deadbeef00"
	testGetEthernetInterface(t, added.MAC, added)

	added.IPs[0] = IP{"10.0.0.0"}

	testPatchEthernetInterface(t, patch)
	testGetEthernetInterface(t, patch.MAC, added)
	testDeleteEthernetInterface(t, added.MAC)
}

func testPatchEthernetInterface(t *testing.T, added NodeInterfacePatch) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	_, err := w.PatchEthernetInterface(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestAddTwoTimesEthernetInterface(t *testing.T) {
	var added = NodeInterface{
		MAC: "00:de:ad:be:ef:00",
		IPs: []IP{IP{"10.0.0.5"}},
	}

	testAddEthernetInterface(t, added)
	testAddEthernetInterface(t, added)
	testDeleteEthernetInterface(t, added.MAC)
}

func TestAddEthernetInterfaceCompID(t *testing.T) {
	var added = NodeInterface{
		MAC:    "00:de:ad:be:ef:00",
		CompID: "x0c0s0b0n0",
		IPs:    []IP{IP{"10.0.0.5"}},
	}

	testAddEthernetInterface(t, added)

	added.ID = "00deadbeef00"
	added.Type = "Node"

	testGetEthernetInterface(t, added.MAC, added)

	testDeleteEthernetInterface(t, added.MAC)
}
