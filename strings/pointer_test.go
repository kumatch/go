package strings

import (
	"testing"
)

func TestPointer(t *testing.T) {
	str := "foo"
	ret := Pointer(str)

	if str != *ret {
		t.Errorf("want %v, got %v", str, *ret)
	}

	ret = Pointer("")

	if ret != nil {
		t.Errorf("want nil, got %v", *ret)
	}
}
