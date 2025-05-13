package manta

import (
	"log"
	"os/exec"
)

// NewWrapper creates a new wrapper with the given base_url and access_token
func NewWrapper(base_url string, access_token string) (*Wrapper, error) {
	c := Wrapper{base_url, access_token}
	cmd := exec.Command("manta", "--version")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Unable to get Manta CLI version: " + err.Error())
		return nil, err
	} else {
		log.Println("Manta CLI version: " + string(out))
	}
	return &c, nil
}
