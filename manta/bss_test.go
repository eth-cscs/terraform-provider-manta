package manta

import (
	"fmt"
	"testing"
)

func testAddBss(t *testing.T, added BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.AddBss(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
		return
	}
}

func testUpdateBss(t *testing.T, added BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.UpdateBss(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
		return
	}
}

func testDeleteBss(t *testing.T, added BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	err := w.DeleteBss(added)

	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func testGetBss(t *testing.T, expected BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	params, err := w.GetBssParams()

	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if len(params) != 1 {
		t.Errorf(`error: boot parameters' length is not equal to 1`)
		return
	}

	if params[0].Cmp(&expected) {
		fmt.Println(params)
		fmt.Println(expected)
		t.Errorf(`error: the added is not equal to the expected`)
	}
}

func testAddGetUpdateDelete(t *testing.T, added, updated BssParams) {
	testAddBss(t, added)
	testGetBss(t, added)
	testUpdateBss(t, updated)
	testGetBss(t, updated)
	testDeleteBss(t, updated)
}

func TestUpdateBssSuccessPaths(t *testing.T) {
	var added = BssParams{
		Macs:   []string{"00:de:ad:be:ef:00"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
	}

	var updated = BssParams{
		Macs:   []string{"00:de:ad:be:ef:00"},
		Kernel: "https://another.com/kernel",
		Initrd: "https://another.com/initrd",
	}

	testAddGetUpdateDelete(t, added, updated)
}

func testAddGetDelete(t *testing.T, added, expected BssParams) {
	testAddBss(t, added)
	testGetBss(t, expected)
	testDeleteBss(t, added)
}

func TestAddBssSuccessMinimun(t *testing.T) {
	var added = BssParams{
		Macs:   []string{"00:de:ad:be:ef:00"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
	}

	var expected = BssParams{
		Macs:   []string{"00:de:ad:be:ef:00"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
	}

	testAddGetDelete(t, added, expected)
}

func TestAddBssSuccessParams(t *testing.T) {
	var added = BssParams{
		Macs:   []string{"00:de:ad:be:ef:01"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
		Params: "console=ttyS0,115200 console=tty0",
	}

	var expected = BssParams{
		Macs:   []string{"00:de:ad:be:ef:01"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
		Params: "console=ttyS0,115200 console=tty0",
	}

	testAddGetDelete(t, added, expected)
}

func TestAddBssSuccessTwoMacs(t *testing.T) {
	var added = BssParams{
		Macs:   []string{"00:de:ad:be:ef:02", "00:de:ad:be:ef:03"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
		Params: "console=ttyS0,115200 console=tty0",
	}

	var expected = BssParams{
		Macs:   []string{"00:de:ad:be:ef:02", "00:de:ad:be:ef:03"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
		Params: "console=ttyS0,115200 console=tty0",
	}

	testAddGetDelete(t, added, expected)
}

func TestDeleteBss(t *testing.T) {
	var added = BssParams{
		Macs:   []string{"00:de:ad:be:ef:02", "00:de:ad:be:ef:03"},
		Kernel: "https://example.com/kernel",
		Initrd: "https://example.com/initrd",
		Params: "console=ttyS0,115200 console=tty0",
	}

	testAddBss(t, added)
	testDeleteBss(t, added)
}
