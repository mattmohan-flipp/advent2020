package intutils

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	if Min(1, -1) != -1 {
		t.Fatalf("Failed @ min(1, -1)")
	}
	if Min(math.MaxInt64, math.MinInt64) != math.MinInt64 {
		t.Fatalf("Failed @ min(max, min)")
	}
}

func TestMax(t *testing.T) {
	if Max(1, -1) != 1 {
		t.Fatalf("Failed @ max(1, -1)")
	}
	if Max(math.MaxInt64, math.MinInt64) != math.MaxInt64 {
		t.Fatalf("Failed @ max(max, min)")
	}
}

func TestAbs(t *testing.T) {
	if Abs(1) != 1 {
		t.Fatalf("Failed @ abs(1)")
	}
	if Abs(-1) != 1 {
		t.Fatalf("Failed @ abs(-1)")
	}
	if Abs(math.MaxInt64) != math.MaxInt64 {
		t.Fatalf("Failed @ abs(max)")
	}
	if Abs(math.MinInt64+1) != -(math.MinInt64 + 1) {
		t.Fatalf("Failed @ abs(min+1)")
	}
}
func TestAbsMin(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic not reached")
		}
	}()
	Abs(math.MinInt64)

}
