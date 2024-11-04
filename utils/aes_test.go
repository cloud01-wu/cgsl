package utils

import (
	"encoding/base64"
	"strings"
	"testing"
)

var expectEncryptString = "WkaOhqSK03Z1pSuPOdc03w=="
var textToEncrypt = "test"
var keySuccess = "1234567890123456"
var keyError = "1"

func TestAesEncryptCBC(t *testing.T) {
	encryptByte, err := AesEncryptCBC([]byte(textToEncrypt), []byte(keySuccess))
	if err != nil {
		t.Errorf("AesEncryptCBC Error, error: %s", err.Error())
	}
	encryptString := base64.StdEncoding.EncodeToString(encryptByte)
	if !strings.EqualFold(expectEncryptString, encryptString) {
		t.Errorf("AesEncryptCBC Error, expect %s but get %s", expectEncryptString, encryptString)
	}
	_, err = AesEncryptCBC([]byte(textToEncrypt), []byte(keyError))
	if err == nil {
		t.Errorf("AesEncryptCBC should be error when key error")
	}
	t.Log("AesEncryptCBC Success")
}

func TestAesDecryptCBC(t *testing.T) {
	encryptByte, _ := base64.StdEncoding.DecodeString(expectEncryptString)
	decryptByte, err := AesDecryptCBC(encryptByte, []byte(keySuccess))
	if err != nil {
		t.Errorf("AesDecryptCBC Error, error: %s", err.Error())
	}
	if !strings.EqualFold(textToEncrypt, string(decryptByte)) {
		t.Errorf("AesDecryptCBC Error, expect %s but get %s", textToEncrypt, string(decryptByte))
	}
	_, err = AesDecryptCBC(encryptByte, []byte(keyError))
	if err == nil {
		t.Errorf("AesDecryptCBC should be error when key error")
	}
	t.Log("AesDecryptCBC Success")
}
