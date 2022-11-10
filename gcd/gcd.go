package gcd

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrZeroValues = errors.New("both values cannot be zero")
)

type Integer interface {
	constraints.Integer
}

func GCD[T Integer](a, b T) (T, error) {
	TZero := *new(T)
	if a == TZero && b == TZero {
		return TZero, ErrZeroValues
	}
	return gcd(a, b), nil
}

func gcd[T Integer](a, b T) T {
	TZero := *new(T)
	if b == TZero {
		return abs(a)
	}
	return gcd(b, abs(a)%b)
}

func abs[T Integer](a T) T {
	TZero := *new(T)
	if a < TZero {
		return -a
	}
	return a
}
