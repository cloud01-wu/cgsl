package utils

import "reflect"

func GetAsByte(arguments []interface{}, idx int, defaultValue int8) int8 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(int8)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsShort(arguments []interface{}, idx int, defaultValue int16) int16 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(int16)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsInt(arguments []interface{}, idx int, defaultValue int32) int32 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(int32)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsLong(arguments []interface{}, idx int, defaultValue int64) int64 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(int64)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsFloat(arguments []interface{}, idx int, defaultValue float32) float32 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(float32)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsDouble(arguments []interface{}, idx int, defaultValue float64) float64 {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(float64)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsString(arguments []interface{}, idx int, defaultValue string) string {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	value, ok := arguments[idx].(string)

	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetAsObject(arguments []interface{}, idx int, defaultValue interface{}) interface{} {
	if len(arguments) < (idx + 1) {
		return defaultValue
	}
	if reflect.TypeOf(defaultValue) == reflect.TypeOf(arguments[idx]) {
		return arguments[idx]
	} else {
		return defaultValue
	}
}
