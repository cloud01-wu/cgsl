package env

import (
	"os"
	"strconv"
)

func GetBool(key string, fallback bool) bool {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseBool(s); err == nil {
			return value
		}
	}

	return fallback
}

func GetString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, strconv.IntSize); err == nil {
			return int(value)
		}
	}

	return fallback
}

func GetInt8(key string, fallback int8) int8 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, 8); err == nil {
			return int8(value)
		}
	}

	return fallback
}

func GetInt16(key string, fallback int16) int16 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, 32); err == nil {
			return int16(value)
		}
	}

	return fallback
}

func GetInt32(key string, fallback int32) int32 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, 32); err == nil {
			return int32(value)
		}
	}

	return fallback
}

func GetInt64(key string, fallback int64) int64 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, 64); err == nil {
			return value
		}
	}

	return fallback
}

func GetUint(key string, fallback uint) uint {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseUint(s, 10, strconv.IntSize); err == nil {
			return uint(value)
		}
	}

	return fallback
}

func GetUint8(key string, fallback uint8) uint8 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseUint(s, 10, 8); err == nil {
			return uint8(value)
		}
	}

	return fallback
}

func GetUint16(key string, fallback uint16) uint16 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseUint(s, 10, 16); err == nil {
			return uint16(value)
		}
	}

	return fallback
}

func GetUint32(key string, fallback uint32) uint32 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseUint(s, 10, 32); err == nil {
			return uint32(value)
		}
	}

	return fallback
}

func GetUint64(key string, fallback uint64) uint64 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseUint(s, 10, 64); err == nil {
			return value
		}
	}

	return fallback
}

func GetFloat32(key string, fallback float32) float32 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseFloat(s, 32); err == nil {
			return float32(value)
		}
	}

	return fallback
}

func GetFloat64(key string, fallback float64) float64 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseFloat(s, 64); err == nil {
			return value
		}
	}

	return fallback
}

func GetComplex64(key string, fallback complex64) complex64 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseComplex(s, 64); err == nil {
			return complex64(value)
		}
	}

	return fallback
}

func GetComplex128(key string, fallback complex128) complex128 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseComplex(s, 64); err == nil {
			return value
		}
	}

	return fallback
}
