package manta

import (
	"testing"
)

func TestGetAccessTokenSuccess(t *testing.T) {
	var w Wrapper = Wrapper{access_token: "~/access_token"}

	// check the size of the token
	if len(w.GetAccessToken()) != 1156 {
		t.Errorf("error: get access token")
	}
}
