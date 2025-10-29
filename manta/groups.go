package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MemberAddBody struct {
	ID string `json:"id"`
}

type Members struct {
	IDs []string `json:"ids,omitempty"`
}

type Group struct {
	Label          string   `json:"label"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags,omitempty"`
	ExclusiveGroup string   `json:"exclusiveGroup,omitempty"`
	Members        `json:"members,omitempty"`
}

func deleteGroup(url, token string) error {
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

func deleteMemberToGroup(url, token string) error {
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

func createGroup(url, token string, newGroup Group) error {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(newGroup)
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

func addMemberToGroup(url, token string, newMember MemberAddBody) error {
	client := &http.Client{}

	jData, err := json.Marshal(newMember)
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

func (w *Wrapper) AddMemberToGroup(xname, group_label string) error {
	return addMemberToGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label+"/members",
		w.GetAccessToken(),
		MemberAddBody{ID: xname},
	)
}

func (w *Wrapper) DeleteMemberToGroup(xname, group_label string) error {
	return deleteMemberToGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label+"/members/"+xname,
		w.GetAccessToken(),
	)
}

func (w *Wrapper) UpdateGroup(newGroup Group) error {
	err := deleteGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+newGroup.Label,
		w.GetAccessToken(),
	)

	if err != nil {
		return err
	}

	return createGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups",
		w.GetAccessToken(),
		newGroup,
	)
}

func (w *Wrapper) CreateGroup(newGroup Group) error {
	return createGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups",
		w.GetAccessToken(),
		newGroup,
	)
}

func (w *Wrapper) DeleteGroup(group_label string) error {
	return deleteGroup(
		"https://foobar.openchami.cluster:8443/hsm/v2/groups/"+group_label,
		w.GetAccessToken(),
	)
}
