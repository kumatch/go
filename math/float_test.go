package math

import (
	"testing"
)

func TestRoundFloat64(t *testing.T) {
	fixtures := []struct {
		source  float64
		place   int
		expects float64
	}{
		{1.14, 0, 1},
		{1.14, 1, 1.1},
		{1.14, 2, 1.14},
		{1.15, 0, 1},
		{1.15, 1, 1.2},
		{1.15, 2, 1.15},
		{11.4, 0, 11},
		{11.4, 1, 11.4},
		{11.4, 2, 11.4},
		{11.5, 0, 12},
		{11.5, 1, 11.5},
		{11.5, 2, 11.5},

		{1.24, 0, 1},
		{1.24, 1, 1.2},
		{1.24, 2, 1.24},
		{1.25, 0, 1},
		{1.25, 1, 1.3},
		{1.25, 2, 1.25},
		{12.4, 0, 12},
		{12.4, 1, 12.4},
		{12.4, 2, 12.4},
		{12.5, 0, 13},
		{12.5, 1, 12.5},
		{12.5, 2, 12.5},
	}

	for i, f := range fixtures {
		r := RoundFloat64(f.source, f.place)
		if f.expects != r {
			t.Errorf("Error index %d: want %v, got %v", i, f.expects, r)
		}
	}
}

func TestFloorFloat64(t *testing.T) {
	fixtures := []struct {
		source  float64
		place   int
		expects float64
	}{
		{1.14, 0, 1},
		{1.14, 1, 1.1},
		{1.14, 2, 1.14},
		{1.15, 0, 1},
		{1.15, 1, 1.1},
		{1.15, 2, 1.15},
		{11.4, 0, 11},
		{11.4, 1, 11.4},
		{11.4, 2, 11.4},
		{11.5, 0, 11},
		{11.5, 1, 11.5},
		{11.5, 2, 11.5},

		{1.24, 0, 1},
		{1.24, 1, 1.2},
		{1.24, 2, 1.24},
		{1.25, 0, 1},
		{1.25, 1, 1.2},
		{1.25, 2, 1.25},
		{12.4, 0, 12},
		{12.4, 1, 12.4},
		{12.4, 2, 12.4},
		{12.5, 0, 12},
		{12.5, 1, 12.5},
		{12.5, 2, 12.5},
	}

	for i, f := range fixtures {
		r := FloorFloat64(f.source, f.place)
		if f.expects != r {
			t.Errorf("Error index %d: want %v, got %v", i, f.expects, r)
		}
	}
}
