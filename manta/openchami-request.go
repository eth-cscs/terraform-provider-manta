package manta

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func openchamiRequestBodyOchami[T any](url, token, method string, obj T) OchamiError {
	ochamiError := OchamiError{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(obj)

	if err != nil {
		return ochamiError
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jData))

	if err != nil {
		return ochamiError
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)

	if err != nil {
		return ochamiError
	}

	body, err := io.ReadAll(resp.Body)
	//fmt.Println(string(body))

	err = json.Unmarshal(body, &ochamiError)
	return ochamiError
}

func openchamiRequestBody[T any](url, token, method string, obj T) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	jData, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jData))

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

func openchamiRequestOchami[T any](url, token, method string) (T, error, OchamiError) {
	var returned T
	ochamiError := OchamiError{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return returned, err, ochamiError
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return returned, err, ochamiError
	}

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &returned)
	err = json.Unmarshal(body, &ochamiError)

	return returned, err, ochamiError
}

func openchamiRequest[T any](url, token, method string) (T, error) {
	var returned T
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return returned, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return returned, err
	}

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &returned)

	if err != nil {
		return returned, err
	}

	return returned, nil
}

func openchamiRequestPost[T any](url, token string, obj T) error {
	return openchamiRequestBody(url, token, "POST", obj)
}

func openchamiRequestPostOchami[T any](url, token string, obj T) OchamiError {
	return openchamiRequestBodyOchami(url, token, "POST", obj)
}

func openchamiRequestPut[T any](url, token string, obj T) error {
	return openchamiRequestBody(url, token, "PUT", obj)
}

func openchamiRequestPatch[T any](url, token string, obj T) error {
	return openchamiRequestBody(url, token, "PATCH", obj)
}

func openchamiRequestGet[T any](url, token string) (T, error) {
	return openchamiRequest[T](url, token, "GET")
}

func openchamiRequestDelete[T any](url, token string) (T, error) {
	return openchamiRequest[T](url, token, "DELETE")
}

func openchamiRequestDeleteBody[T any](url, token string, obj T) error {
	return openchamiRequestBody(url, token, "DELETE", obj)
}
