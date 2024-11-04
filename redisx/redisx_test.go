package redisx

import (
	"testing"
)

func TestNew(t *testing.T) {
	Init("192.168.2.236:6379", "", 16, 16)
	_, err := New()

	if err != nil {
		t.Error("Get Error when New Database")
	} else {
		t.Log("New Database success")
	}
}
