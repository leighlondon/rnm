package main

import (
	"testing"
)

func TestVersionIsDefined(t *testing.T) {
	if version == "" {
		t.Errorf("expected version, got '%s'", version)
	}
}
