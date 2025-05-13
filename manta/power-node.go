package manta

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func (w *Wrapper) PowerNodeId(id string, powerStatus string) (NodeItem, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if (powerStatus != "Off") && (powerStatus != "On") {
		return NodeItem{}, errors.New(`powerStatus ins't "Off" or "On"`)
	}

	cmd := exec.Command("manta", "power", strings.ToLower(powerStatus), "nodes", "--assume-yes", id)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
	cmd.Wait()
	fmt.Println(string(stdout.Bytes()))
	fmt.Println(string(stderr.Bytes()))
	return w.GetNodeId(id)
}
