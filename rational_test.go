package gorational

import (
	"testing"
)

const a, b int = 1, 2
const x, y int = 2, 4

// TestNew will test to see if the constructor can not only handle any amount
// of arguments, but also whether or not the rational number returned is in
// lowest terms.
func TestNew(t *testing.T) {

	// Test if the number has been reduced
	if r := New(2, 4); r.x != 1 || r.y != 2 {
		t.Errorf("New(2, 4) = %v / %v, expected %v / %v. Must be reduced",
			r.x, r.y, 1, 2)
	}

	// Test number that is already in lowest terms
	if r := New(1, 3); r.x != 1 || r.y != 3 {
		t.Errorf("New(1, 3) = %v / %v, expected %v / %v. Must not be reduced",
			r.x, r.y, 1, 3)
	}

	// Test a single argument (whole number)
	if r := New(3); r.x != 3 || r.y != 1 {
		t.Errorf("New(3) = %v / %v, expected %v / %v. Must have denominator of 1",
			r.x, r.y, 3, 1)
	}

	// Test no arguments (0)
	if r := New(); r.x != 0 || r.y != 1 {
		t.Errorf("New() = %v / %v, expected %v / %v. Must have a denominator of 1",
			r.x, r.y, 0, 1)
	}

	// Test more than 2 arguments (should ignore extras)
	if r := New(1, 3, 6); r.x != 1 || r.y != 3 {
		t.Errorf("New(1, 3, 6) = %v / %v, expected %v / %v.",
			r.x, r.y, 1, 3)
	}
}

func TestRational_Numerator(t *testing.T) {
	if r := New(a, b); r.Numerator() != a {
		t.Errorf("Unable to get numerator. Expected %v, got %v", a, r.Numerator())
	}
}

func TestRational_Denominator(t *testing.T) {
	if r := New(a, b); r.Denominator() != b {
		t.Errorf("Unable to get denominator. Expected %v, got %v", b, r.Denominator())
	}
}

func TestRational_SetNumerator(t *testing.T) {
	r1 := Rational{a, b}
	r1.SetNumerator(x)
	if r1.x != x {
		t.Errorf("Unable to set numerator to %v.", x)
	}
}

func TestRational_SetDenominator(t *testing.T) {
	r1 := Rational{a, b}
	r1.SetDenominator(y)
	if r1.y != y {
		t.Errorf("Unable to set denominator to %v.", y)
	}
}

func TestRational_Normalize(t *testing.T) {
	r := Rational{x, y}
	r.Normalize()
	if r.x != a || r.y != b {
		t.Errorf("r.Normalize() = %v/%v. Expected %v/%v", r.x, r.y, a, b)
	}
}

func TestGCD(t *testing.T) {
	if i := GCD(x, y); i != b {
		t.Errorf("GCD(%v, %v) = %v, expected %v", x, y, i, b)
	}
}
