package manta

import (
	"fmt"
	"testing"
)

func testAddBss(t *testing.T, added, expected BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	added, err := w.AddBss(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if added.Cmp(&expected) {
		fmt.Println(added)
		fmt.Println(expected)
		t.Errorf(`error: the added is not equal to the expected`)
	}
}

func testDeleteBss(t *testing.T, added BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	body, err := w.DeleteBss(added)

	if body != "" {
		t.Errorf(`error: body should be empty: %s`, body)
	}

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

func testAddGetDelete(t *testing.T, added, expected BssParams) {
	testAddBss(t, added, expected)
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

	testAddBss(t, added, added)
	testDeleteBss(t, added)
}
