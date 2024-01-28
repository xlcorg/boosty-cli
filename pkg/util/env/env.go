package env

import (
	"errors"
	"os"
	"strconv"
)

func GetStringOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetIntOrDefault(key string, defaultValue int) int {
	if v, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(v); err == nil {
			return value
		}
	}
	return defaultValue
}

func GetString(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	throwError(key)
	return ""
}

func GetBoolOrDefault(key string, defaultValue bool) bool {
	val, err := TryGetBool(key)
	if err != nil {
		return defaultValue
	}

	return val
}

func TryGetString(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", errors.New(`Environment variable not found: "` + key + `"`)
}

func TryGetBool(key string) (bool, error) {
	val, err := TryGetString(key)
	if err != nil {
		return false, err
	}

	res, err := strconv.ParseBool(val)
	if err != nil {
		return false, errors.New(`Environment variable is not boolean: ` + key + `=` + val)
	}

	return res, nil
}

func GetInt(key string) int {
	return int(GetInt64(key))
}

func GetInt64(key string) int64 {
	if s, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseInt(s, 10, 64); err == nil {
			return value
		}
	}
	throwError(key)
	return 0
}

func throwError(key string) {
	panic(`Environment variable not found: "` + key + `"`)
}
