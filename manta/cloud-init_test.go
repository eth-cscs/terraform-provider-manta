package manta

import (
	"testing"
)

func testGetClusterDefault(t *testing.T, clusterDefault ClusterDefaults) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	cd, err := w.GetClusterDefault()

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if cd.Cmp(&clusterDefault) {
		t.Errorf(`error: received and expected aren't equal`)
	}
}

func testCreateClusterDefault(t *testing.T, clusterDefault ClusterDefaults) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.CreateClusterDefault(clusterDefault)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestCreateGetClusterDefault(t *testing.T) {
	clusterDefault := ClusterDefaults{
		BaseUrl: "http://cloud-init/cloud-init",
		PublicKeys: []string{
			"ssh-ed25519 AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA user1@demo-head",
		},
	}

	testCreateClusterDefault(t, clusterDefault)
	testGetClusterDefault(t, clusterDefault)
}
