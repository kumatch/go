package env

import (
	"os"
	"strconv"
)

func Get(name string) (string, bool) {
	v := os.Getenv(name)

	return v, v != ""
}

func GetDefault(name string, defaultValue string) string {
	if v, ok := Get(name); ok {
		return v
	}

	return defaultValue
}

func GetInt(name string) (int, bool) {
	v := os.Getenv(name)

	if v == "" {
		return 0, false
	}

	num, err := strconv.Atoi(v)
	if err != nil {
		return 0, false
	}

	return num, true
}

func GetIntDefault(name string, defaultValue int) int {
	if v, ok := GetInt(name); ok {
		return v
	}

	return defaultValue
}
