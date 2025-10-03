package manta

import (
	"fmt"
	"testing"
)

func testAddBssSuccess(t *testing.T, added, expected BssParams) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	added, err := w.AddBss(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	_, err = w.GetBssParams()
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	_, err = w.DeleteBss(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if added.Cmp(&expected) {
		fmt.Println(added)
		fmt.Println(expected)
		t.Errorf(`error: the added is not equal to the expected`)
	}
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

	testAddBssSuccess(t, added, expected)
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

	testAddBssSuccess(t, added, expected)
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

	testAddBssSuccess(t, added, expected)
}
