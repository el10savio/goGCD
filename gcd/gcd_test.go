package gcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		name        string
		a           int
		b           int
		expectedGCD int
		expectedErr error
	}{
		{"base case", 8, 12, 4, nil},
		{"base case ii", 48, 18, 6, nil},
		{"base case ii reversed", 18, 48, 6, nil},
		{"zero values", 0, 0, 0, ErrZeroValues},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualGCD, err := GCD(tc.a, tc.b)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.expectedGCD, actualGCD)
		})
	}
}

func FuzzGCD(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		gcd, err := GCD(a, b)

		if a == 0 && b == 0 {
			assert.ErrorIs(t, err, ErrZeroValues)
			return
		}

		if a > 0 && b > 0 {
			assert.True(t, gcd <= a && gcd <= b)
		}

		if b == 0 {
			assert.Equal(t, abs(a), gcd)
		}

		gcdReversed, err := GCD(b, a)
		assert.NoError(t, err)
		assert.True(t, gcd == gcdReversed)

		assert.Equal(t, 0, a%gcd)
		assert.Equal(t, 0, b%gcd)

		if gcd == 1 {
			testNum := 12

			gcdABTest, err := GCD(a*b, testNum)
			assert.NoError(t, err)
			gcdATest, err := GCD(a, testNum)
			assert.NoError(t, err)
			gcdBTest, err := GCD(b, testNum)
			assert.NoError(t, err)

			assert.Equal(t, gcdABTest, gcdATest*gcdBTest)
		}
	})
}
