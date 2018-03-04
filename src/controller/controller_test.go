package controller

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	if result != `{"message":"Hello World"}` {
		t.Error("Message was: " +  result)
	}
}
