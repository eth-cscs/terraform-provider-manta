package manta

import (
	"testing"
)

func testGetGroupData(t *testing.T, groupData GroupData) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	cd, err := w.GetGroupData(groupData.Name)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if cd.Cmp(&groupData) {
		t.Errorf(`error: received and expected aren't equal`)
	}
}

func testCreateGroupData(t *testing.T, groupData GroupData) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.CreateGroupData(groupData)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testDeleteGroupData(t *testing.T, groupDataName string) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteGroupData(groupDataName)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestCreateGetGroupData(t *testing.T) {
	groupData := GroupData{
		Description: "The compute group",
		File: CloudConfigFile{
			Content:  []byte("IyMgdGVtcGxhdGU6IGppbmphCiNjbG91ZC1jb25maWcKbWVyZ2VfaG93OgotIG5hbWU6IGxpc3QKICBzZXR0aW5nczogW2FwcGVuZF0KLSBuYW1lOiBkaWN0CiAgc2V0dGluZ3M6IFtub19yZXBsYWNlLCByZWN1cnNlX2xpc3RdCnVzZXJzOgogIC0gbmFtZTogcm9vdAogICAgc3NoX2F1dGhvcml6Z      WRfa2V5czoge3sgZHMubWV0YV9kYXRhLmluc3RhbmNlX2RhdGEudjEucHVibGljX2tleXMgfX0KZGlzYWJsZV9yb290OiBmYWxzZQo="),
			Encoding: "base64",
		},
		Name: "compute",
	}

	testCreateGroupData(t, groupData)
	testGetGroupData(t, groupData)
	testDeleteGroupData(t, groupData.Name)
}
