package utils

import (
	"strings"
	"testing"
)

func TestCurrentFunctionName(t *testing.T) {
	currentFunctionName := CurrentFunctionName()
	if strings.EqualFold(currentFunctionName, "sercomm.com/ai/commons/util.TestCurrentFunctionName") {
		t.Logf("Get current function name success, name:%s \n", currentFunctionName)
	} else {
		t.Errorf("Get current function name error, expect %s but get %s", "sercomm.com/ai/commons/util.TestCurrentFunctionName", currentFunctionName)
	}
}

func TestCurrentCallerName(t *testing.T) {
	currentCallerName := CurrentCallerName()
	if strings.EqualFold(currentCallerName, "testing.tRunner") {
		t.Logf("Get current caller name success, name:%s \n", currentCallerName)
	} else {
		t.Errorf("Get current caller name error, expect %s but get %s", "testing.tRunner", currentCallerName)
	}
}
