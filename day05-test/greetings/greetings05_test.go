package greetings05

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Peter"
	want := regexp.MustCompile(`\b`+name+`\b`)

	message, error := Hello("Gladys")

	if(!want.MatchString(message) || error != nil) {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if(msg != "" || err == nil) {
		t.Errorf(`Hello("") = %q , %v, want "", error`, msg, err )
	}
}