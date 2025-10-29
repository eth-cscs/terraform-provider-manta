package manta

import (
	"testing"
)

func testCreateGroup(t *testing.T, group Group) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.CreateGroup(group)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testUpdateGroup(t *testing.T, group Group) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.UpdateGroup(group)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testDeleteGroup(t *testing.T, group_label string) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteGroup(group_label)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestCreateDeleteGroup(t *testing.T) {
	const xname = "x1000c0s0b0n0"
	components := Components{[]Component{
		{
			ID:    xname,
			State: "Ready",
			Role:  "Compute",
			Arch:  "X86",
		},
	},
	}
	testCreateComponent(t, components)

	group := Group{
		Label:       "test",
		Description: "Test group",
		Members: Members{
			IDs: []string{xname},
		},
	}
	testCreateGroup(t, group)
	testDeleteGroup(t, group.Label)

	testDeleteComponent(t, components)
}

func TestCreateUpdateDeleteGroup(t *testing.T) {
	var xname1 = "x1000c0s0b0n0"
	var xname2 = "x2000c0s0b0n0"
	components := Components{[]Component{
		{
			ID:    xname1,
			State: "Ready",
			Role:  "Compute",
			Arch:  "X86",
		},
	},
	}

	testCreateComponent(t, components)

	group := Group{
		Label:       "test",
		Description: "Test group",
		Members: Members{
			IDs: []string{xname1},
		},
	}

	// create group
	testCreateGroup(t, group)

	// delete xname1
	testDeleteComponent(t, components)

	// create xname2
	components.Components[0].ID = xname2
	testCreateComponent(t, components)

	group.Members = Members{IDs: []string{xname2}}

	// update group
	testUpdateGroup(t, group)

	// delete group
	testDeleteGroup(t, group.Label)

	// delete xname2
	testDeleteComponent(t, components)
}
