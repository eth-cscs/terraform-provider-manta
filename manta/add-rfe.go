package manta

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func execOutErr(command string, arg ...string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, arg...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(stderr.String())
		return stderr.String(), errors.New(strings.Split(stderr.String(), "\n")[1])
	} else {
		fmt.Println(stdout.String())
	}

	return stdout.String(), nil
}

func (w *Wrapper) AddRfe(rfeItem RfeItem) (RfeItem, error) {
	var arguments []string

	if rfeItem.ID == "" {
		return RfeItem{}, errors.New("ID cannot be empty")
	}

	arguments = append(arguments, "add")
	arguments = append(arguments, "redfish-endpoint")

	arguments = append(arguments, "--id")
	arguments = append(arguments, rfeItem.ID)

	if rfeItem.FQDN != "" {
		arguments = append(arguments, "--fqdn")
		arguments = append(arguments, rfeItem.FQDN)
	}

	if rfeItem.User != "" {
		arguments = append(arguments, "--user")
		arguments = append(arguments, rfeItem.User)
	}

	if rfeItem.Hostname != "" {
		arguments = append(arguments, "--hostname")
		arguments = append(arguments, rfeItem.Hostname)
	}

	if rfeItem.Enabled != false {
		arguments = append(arguments, "--enabled")
	}

	if rfeItem.RediscoverOnUpdate != false {
		arguments = append(arguments, "--rediscover-on-update")
	}

	_, err := execOutErr("manta", arguments...)

	if err != nil {
		return RfeItem{}, err
	}

	rfeReturn, _ := w.GetRfeId(rfeItem.ID)

	return rfeReturn, nil
}
