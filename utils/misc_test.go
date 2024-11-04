package utils

import (
	"strings"
	"testing"
)

// func TestGetCurrentDirectory(t *testing.T) {
// 	directory, err := GetCurrentDirectory()
// 	if err == nil {
// 		t.Error(err.Error())
// 	} else {
// 		t.Logf("Get current directory success, directory: %s", directory)
// 	}
// }

func TestGetAppName(t *testing.T) {
	appName := GetAppName()
	if strings.EqualFold(appName, "util.test") {
		t.Logf("Get app name success, %s \n", appName)
	} else {
		t.Errorf("Get app name error, expect get %s but get %s\n", "util.test", appName)
	}

}
