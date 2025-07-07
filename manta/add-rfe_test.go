package manta

import (
	"testing"
)

func TestAddRfeFail(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	var rfe = RfeItem{ID: "adsf"}

	_, err := w.AddRfe(rfe)
	if err == nil {
		return
	}

	t.Errorf(`error: add rfe should fail`)
}

func testAddRfeSuccess(t *testing.T, added, expected RfeItem) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}

	_, err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	added, err = w.AddRfe(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	if added != expected {
		t.Errorf("error: rfe added is not as expected\nadded: %s\nexpected: %s\ndiff: %s",
			added.String(),
			expected.String(),
			GetDiff(added, expected),
		)
	}

	_, err = w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestAddTwiceRfe(t *testing.T) {
	var w = Wrapper{Access_token: "~/access_token", Base_url: "http://localhost:3000"}
	var added = RfeItem{ID: testXnameRfe}

	_, err := w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}

	added, err = w.AddRfe(added)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
	added, err = w.AddRfe(added)
	if err == nil {
		t.Errorf(`error: add rfe should fail`)
	}

	_, err = w.DeleteRfe(testXnameRfe)
	if err != nil {
		t.Errorf(`error: %s`, err)
	}
}

func TestAddRfeSuccess(t *testing.T) {
	var added = RfeItem{ID: testXnameRfe}
	var expected = RfeItem{ID: testXnameRfe, Type: "NodeBMC", FQDN: testXnameRfe}

	testAddRfeSuccess(t, added, expected)
}

func TestAddRfeSuccessAll(t *testing.T) {
	var added = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		Enabled:            true,
		User:               "user",
		Password:           "password",
		RediscoverOnUpdate: true,
	}
	var expected = RfeItem{ID: testXnameRfe,
		Hostname:           "hostname",
		FQDN:               "hostname",
		Type:               "NodeBMC",
		Enabled:            true,
		User:               "user",
		Password:           "",
		RediscoverOnUpdate: true,
	}

	testAddRfeSuccess(t, added, expected)
}
