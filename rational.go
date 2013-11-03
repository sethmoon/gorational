package gorational

import (
	"fmt"
	"math"
)

// A type that represents a single rational number
type Rational struct {
	x int // Numerator
	y int // Denominator
}

// New accepts either 0, 1, or 2 arguments and creates a new rational number
// with them. One argument indicates a whole number and two indicates a
// fraction.
func New(vals ...int) *Rational {
	r := new(Rational)
	r.y = 1 // Denominator is always 1 by default
	// If there is 1 or more arguments, set the numerator to the first
	// argument.
	if len(vals) >= 1 {
		r.x = vals[0]
	}
	// If there are two or more arguments, set the second argument to the
	// denominator. Other arguments are ignored
	if len(vals) >= 2 {
		r.y = vals[1]
	}

	r.Normalize() // Reduce the number
	return r
}

func (r Rational) Numerator() int {
	return r.x
}

func (r *Rational) SetNumerator(x int) {
	r.x = x
}

func (r Rational) Denominator() int {
	return r.y
}

func (r *Rational) SetDenominator(y int) {
	r.y = y
}

// Repr returns a floating point representation of the rational number.
func (r Rational) Repr() float64 {
	return float64(r.x) / float64(r.y)
}

// normalize will bring the rational number to its lowest terms (reducing it)
func (r *Rational) Normalize() {
	if r.x == 0 {
		r.y = 1
	}

	if (r.x < 0 && r.y < 0) || (r.y < 0) {
		r.x = -r.x
		r.y = -r.y
	}
	gcd := r.gcd()
	r.x /= gcd
	r.y /= gcd
}

// String formats the rational number as a fraction
func (r Rational) String() string {
	return fmt.Sprintf("%v/%v", r.x, r.y)
}

// gcd finds the greatest common divisor between the numerator and denominator.
func (r Rational) gcd() int {
	return GCD(r.x, r.y)
}

// GCD finds the greatest common divisor between two integers using Euclid's
// algorithm.
func GCD(a, b int) int {
	if a == 0 {
		return b
	}

	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	remainder := b % a
	if remainder != 0 {
		return GCD(remainder, a)
	}
	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
