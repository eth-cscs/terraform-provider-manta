package manta

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func (w *Wrapper) GetAccessToken() string {
	if w.access_token_content != "" {
		return w.access_token_content
	}

	var path string = w.access_token

	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}

	byteFile, _ := os.ReadFile(path)
	var stringFile string = string(byteFile)

	var splitFile []string = strings.Split(stringFile, "\n")

	w.access_token_content = splitFile[0]

	return w.access_token_content
}
