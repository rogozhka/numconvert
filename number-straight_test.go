package numconvert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestInString_Int(t *testing.T) {

	type testCase struct {
		s        string
		expected int
		e        error
	}

	cases := []testCase{
		{
			s:        "-2",
			expected: -2,
			e:        nil,
		},
		{
			s:        "-9223372036854775807",
			expected: 0,
			e:        fmt.Errorf("must overflow int32"),
		},
		{
			s:        "9223372036854775807",
			expected: 0,
			e:        fmt.Errorf("must overflow int32"),
		},
	}

	for _, tc := range commonCases {

		src := InString(tc.s)

		if tc.err != nil {
			assertPanic(t, func() { src.Int() })
		} else {
			assert.Equal(t, int(tc.expected), src.Int(), "res")
		}
	}

	for _, tc := range cases {
		src := InString(tc.s)

		if tc.e != nil {
			assertPanic(t, func() { src.Int() })
		} else {
			assert.Equal(t, int(tc.expected), src.Int(), "res")
		}
	}
}
