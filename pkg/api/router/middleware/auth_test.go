package middleware

import "testing"

func TestCheckToken(t *testing.T) {
	if err := CheckToken("xxx"); err != nil {
		t.Fatal(err)
	}
}
