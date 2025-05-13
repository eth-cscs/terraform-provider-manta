package manta

import (
	"fmt"
	"testing"
)

func TestDeleteRfeSuccess(t *testing.T) {
	var w Wrapper

	out, _ := w.DeleteRfe("x0c0s0b0")

	fmt.Println(out)
}
