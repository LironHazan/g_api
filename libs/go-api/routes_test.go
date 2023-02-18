package go_api

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello("world")
	if result != "Hello world" {
		t.Error("Expected Hello to append 'world'")
	}
}
