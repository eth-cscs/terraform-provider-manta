package manta

import (
	//"encoding/json"
	"fmt"
	"testing"
)

func testCreateComponent(t *testing.T, components Components) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.CreateComponent(components)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testGetComponent(t *testing.T, components Components) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	received, err := w.GetComponent()

	if len(components.Components) != len(received.Components) {
		t.Errorf(`not same length`)
	}

	for _, component := range components.Components {
		fmt.Println("value " + component.ID)
	}

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testDeleteComponent(t *testing.T, components Components) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteComponent(components)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestCreateGetDeleteComponent(t *testing.T) {
	components := Components{[]Component{
		{
			ID:    "x1000c0s0b0n0",
			State: "Ready",
			Role:  "Compute",
			Arch:  "X86",
		},
	},
	}

	testCreateComponent(t, components)
	testGetComponent(t, components)
	testDeleteComponent(t, components)
}
