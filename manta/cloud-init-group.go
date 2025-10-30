package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CloudConfigFile struct {
	Content  []byte `json:"content"`
	Name     string `json:"filename"`
	Encoding string `json:"encoding,omitempty"`
}

type GroupData struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Data        map[string]interface{} `json:"meta-data,omitempty"`
	File        CloudConfigFile        `json:"file,omitempty"`
	Versions    map[string]string      `json:"versions,omitempty"`
}

func (self *GroupData) Cmp(other *GroupData) bool {
	if cmpArray(self.File.Content, other.File.Content) {
		return true
	}

	return false
}

func getGroupData(url, token string) (GroupData, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return GroupData{}, err
	}

	//req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return GroupData{}, err
	}

	var groupData GroupData

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &groupData)

	if err != nil {
		return GroupData{}, err
	}

	return groupData, nil
}

func deleteGroupData(url, token string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

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

func postGroupData(url, token string, GroupData GroupData) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(GroupData)
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

func (w *Wrapper) CreateGroupData(newGroupData GroupData) error {
	return postGroupData(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups",
		w.GetAccessToken(),
		newGroupData,
	)
}

func (w *Wrapper) DeleteGroupData(groupDataName string) error {
	return deleteGroupData(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups/"+groupDataName,
		w.GetAccessToken(),
	)
}

func (w *Wrapper) GetGroupData(group_name string) (GroupData, error) {
	return getGroupData(
		"https://foobar.openchami.cluster:8443/cloud-init/admin/groups/"+group_name,
		w.GetAccessToken(),
	)
}
