package middleware

import (
	"fmt"
	"strings"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := GenerateJWT()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("token:", token)

	c, err := ParseJWT(token)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(c)
}

func TestAA(t *testing.T) {
	str := "Bearer ccc"
	seg := strings.Split(str, "Bearer ")
	fmt.Println(seg[0])
	fmt.Println(seg[1])
}
