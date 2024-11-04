package utils

import (
	"strings"
	"testing"
)

func TestGetAsByte(t *testing.T) {
	testSlice := []interface{}{int8(0), "1"}
	getByte := GetAsByte(testSlice, 0, int8(1))
	if getByte != int8(0) {
		t.Errorf("Get as byte error, expect %d but get %d", int8(0), getByte)
	}
	getByte = GetAsByte(testSlice, 1, int8(0))
	if getByte != int8(0) {
		t.Errorf("Get as byte error, expect %d but get %d", int8(0), getByte)
	}
	t.Log("Get byte success")
}

func TestGetAsShort(t *testing.T) {
	testSlice := []interface{}{int16(0), "1"}
	getShort := GetAsShort(testSlice, 0, int16(1))
	if getShort != int16(0) {
		t.Errorf("Get as short error, expect %d but get %d", int16(0), getShort)
	}
	getShort = GetAsShort(testSlice, 1, int16(0))
	if getShort != int16(0) {
		t.Errorf("Get as short error, expect %d but get %d", int16(0), getShort)
	}
	t.Log("Get short success")
}

func TestGetAsInt(t *testing.T) {
	testSlice := []interface{}{int32(0), "1"}
	getInt := GetAsInt(testSlice, 0, int32(1))
	if getInt != int32(0) {
		t.Errorf("Get as Int error, expect %d but get %d", int32(0), getInt)
	}
	getInt = GetAsInt(testSlice, 1, int32(0))
	if getInt != int32(0) {
		t.Errorf("Get as Int error, expect %d but get %d", int32(0), getInt)
	}
	t.Log("Get int success")
}

func TestGetAsLong(t *testing.T) {
	testSlice := []interface{}{int64(0), "1"}
	getLong := GetAsLong(testSlice, 0, int64(1))
	if getLong != int64(0) {
		t.Errorf("Get as Long error expect %d but get %d", int64(0), getLong)
	}
	getLong = GetAsLong(testSlice, 1, int64(0))
	if getLong != int64(0) {
		t.Errorf("Get as Long error expect %d but get %d", int64(0), getLong)
	}
	t.Log("Get long success")
}

func TestGetAsFloat(t *testing.T) {
	testSlice := []interface{}{float32(0.1), "1"}
	getFloat := GetAsFloat(testSlice, 0, float32(0.5))
	if getFloat != float32(0.1) {
		t.Errorf("Get as float error expect %g but get %g", float32(0.1), getFloat)
	}
	getFloat = GetAsFloat(testSlice, 1, float32(1.5))
	if getFloat != float32(1.5) {
		t.Errorf("Get as float error expect %g but get %g", float32(1.5), getFloat)
	}
	t.Log("Get float success")
}

func TestGetAsDouble(t *testing.T) {
	testSlice := []interface{}{float64(0.1), "1"}
	getDouble := GetAsDouble(testSlice, 0, float64(0.5))
	if getDouble != float64(0.1) {
		t.Errorf("Get as double error expect %g but get %g", float64(0.1), getDouble)
	}
	getDouble = GetAsDouble(testSlice, 1, float64(1.5))
	if getDouble != float64(1.5) {
		t.Errorf("Get as double error expect %g but get %g", float64(1.5), getDouble)
	}
	t.Log("Get double success")
}

func TestGetAsString(t *testing.T) {
	testSlice := []interface{}{"0", 1}
	getString := GetAsString(testSlice, 0, "1")
	if !strings.EqualFold(getString, "0") {
		t.Errorf("Get as string error expect %s but get %s", "0", getString)
	}
	getString = GetAsString(testSlice, 1, "0")
	if !strings.EqualFold("0", getString) {
		t.Errorf("Get as string error expect %s but get %s", "0", getString)
	}
	t.Log("Get string success")
}
