package greeting_test

import (
	"greeting"
	"testing"
)

func TestGreeting(t *testing.T) {
	result := greeting.Greeting("Hello")
	if result != "<b>Hello</b>" {
		t.Errorf("Não teve resultado esperado")
	}
}
