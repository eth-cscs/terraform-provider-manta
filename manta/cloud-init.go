package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ClusterDefaults struct {
	CloudProvider    string   `json:"cloud_provider,omitempty"`
	Region           string   `json:"region,omitempty"`
	AvailabilityZone string   `json:"availability-zone,omitempty"`
	ClusterName      string   `json:"cluster-name,omitempty"`
	PublicKeys       []string `json:"public-keys,omitempty"`
	BaseUrl          string   `json:"base-url,omitempty"`
	BootSubnet       string   `json:"boot-subnet,omitempty"`
	WGSubnet         string   `json:"wg-subnet,omitempty"`
	ShortName        string   `json:"short-name,omitempty"`
	NidLength        int      `json:"nid-length,omitempty"`
}

func (self *ClusterDefaults) Cmp(other *ClusterDefaults) bool {
	if cmpArray(self.PublicKeys, other.PublicKeys) {
		return true
	}
	if self.BaseUrl != other.BaseUrl {
		return true
	}

	return false
}

func getClusterDefault(url, token string) (ClusterDefaults, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return ClusterDefaults{}, err
	}

	//req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return ClusterDefaults{}, err
	}

	body, err := io.ReadAll(resp.Body)
	var clusterDefault ClusterDefaults
	err = json.Unmarshal(body, &clusterDefault)

	if err != nil {
		return ClusterDefaults{}, err
	}

	return clusterDefault, nil
}

func postClusterDefault(url, token string, ClusterDefaults ClusterDefaults) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(ClusterDefaults)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jData))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	if err != nil {
		return err
	}

	return nil
}

func (w *Wrapper) CreateClusterDefault(newClusterDefaults ClusterDefaults) error {
	return postClusterDefault(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/cluster-defaults",
		w.GetAccessToken(),
		newClusterDefaults,
	)
}

func (w *Wrapper) GetClusterDefault() (ClusterDefaults, error) {
	return getClusterDefault(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/cluster-defaults",
		w.GetAccessToken(),
	)
}
