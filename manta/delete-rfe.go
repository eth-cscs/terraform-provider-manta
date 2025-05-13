package manta

import (
	"bytes"
	"os/exec"
)

func (w *Wrapper) DeleteRfe(rfeID string) (string, error) {
	var stdout bytes.Buffer

	cmd := exec.Command("manta", "delete", "redfish-endpoint", "--id", rfeID)

	cmd.Stdout = &stdout

	cmd.Run()

	return string(stdout.Bytes()), nil
}
