package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if CFG == nil {
		t.Error("configuration is nil")
	}

	t.Log("finish configuration test")
}
