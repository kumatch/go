package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	testKey := "foo"
	testValue := "bar"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	v1, ok1 := Get(testKey)
	if v1 != testValue {
		t.Errorf("want %s, got %s.", testValue, v1)
	}
	if ok1 == false {
		t.Errorf("want true, got false.")
	}

	v2, ok2 := Get("unknown")
	if v2 != "" {
		t.Errorf("want blank, got %s.", v2)
	}
	if ok2 == true {
		t.Errorf("want false, got true.")
	}
}

func TestGetDefault(t *testing.T) {
	testKey := "foo"
	testValue := "bar"
	defaultValue := "baz"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	v1 := GetDefault(testKey, defaultValue)
	if v1 != testValue {
		t.Errorf("want %s, got %s.", testValue, v1)
	}

	v2 := GetDefault("unknown", defaultValue)
	if v2 != defaultValue {
		t.Errorf("want %s, got %s.", defaultValue, v2)
	}
}

func TestGetInt(t *testing.T) {
	testKey := "foo"
	testValue := "42"
	expectsValue := 42
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	v1, ok1 := GetInt(testKey)
	if v1 != expectsValue {
		t.Errorf("want %d, got %d.", expectsValue, v1)
	}
	if ok1 == false {
		t.Errorf("want true, got false.")
	}

	v2, ok2 := GetInt("unknown")
	if v2 != 0 {
		t.Errorf("want 0, got %d.", v2)
	}
	if ok2 == true {
		t.Errorf("want false, got true.")
	}
}

func TestGetIntDefault(t *testing.T) {
	testKey := "foo"
	testValue := "42"
	expectsValue := 42
	defaultValue := 128

	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	v1 := GetIntDefault(testKey, defaultValue)
	if v1 != expectsValue {
		t.Errorf("want %d, got %d.", expectsValue, v1)
	}

	v2 := GetIntDefault("unknown", defaultValue)
	if v2 != defaultValue {
		t.Errorf("want %d, got %d.", defaultValue, v2)
	}
}
