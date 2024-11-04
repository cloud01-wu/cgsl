package env

import (
	"os"
	"testing"
)

const key = "_GO_TEST_ENV_VALUE"

func TestGetBool(t *testing.T) {
	const value = true
	err := os.Setenv(key, "true")
	if err != nil {
		t.Error(err)
	}

	output := GetBool(key, false)
	if output != value {
		t.Errorf("mismatched")
	}
}

func TestGetString(t *testing.T) {
	const value = "abcdefg"
	err := os.Setenv(key, value)
	if err != nil {
		t.Error(err)
	}

	output := GetString(key, "fallback")
	if output != value {
		t.Errorf("mismatched")
	}
}

func TestGetInt(t *testing.T) {
	const value = 5566
	err := os.Setenv(key, "5566")
	if err != nil {
		t.Error(err)
	}

	output := GetInt(key, 0)
	if output != value {
		t.Errorf("mismatched")
	}
}

func TestGetInt64(t *testing.T) {
	const value = -33333556677
	err := os.Setenv(key, "-33333556677")
	if err != nil {
		t.Error(err)
	}

	output := GetInt64(key, 0)
	if output != value {
		t.Errorf("mismatched")
	}
}
