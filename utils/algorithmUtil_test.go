package utils

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	length := 8
	randomString := RandString(length)
	if length != len(randomString) {
		t.Errorf("Random String length error, expect %d but get str: %s len: %d", length, randomString, len(randomString))
	}
	t.Log(randomString)
}

func TestUUID(t *testing.T) {
	length := 36
	uuidString := RandomUUIDString()
	if length != len(uuidString) {
		t.Errorf("UUID String length error, expect %d but get str: %s len: %d", length, uuidString, len(uuidString))
	}

	splits := strings.Split(uuidString, "-")
	//UUID must be format in 8-4-4-4-12
	if len(splits) != 5 {
		t.Errorf("UUID string format error, get string: %s", uuidString)
	} else if len(splits[0]) != 8 || len(splits[1]) != 4 || len(splits[2]) != 4 || len(splits[3]) != 4 || len(splits[4]) != 12 {
		t.Errorf("UUID string format error, get string: %s", uuidString)
	}

	t.Log(uuidString)
}
