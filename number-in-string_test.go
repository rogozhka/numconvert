package numconvert

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

//
// testCaseCommon is common case for all the types
//
type testCaseCommon struct {
	s        string
	expected uint
	err      error
}

var commonCases = []testCaseCommon{
	{
		s:        "0",
		expected: 0,
		err:      nil,
	},
	{
		s:        "0xaaa",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "aaa",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "-aaa",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "-",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "+",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "--0",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "++0",
		expected: 0,
		err:      fmt.Errorf("nonformat"),
	},
	{
		s:        "+0",
		expected: 0,
		err:      nil,
	},
	{
		s:        "-0",
		expected: 0,
		err:      nil,
	},
	{
		s:        "1",
		expected: 1,
		err:      nil,
	},
	{
		s:        "2",
		expected: 2,
		err:      nil,
	},
}

func TestNumberInString_AsInt(t *testing.T) {

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

		src := NumberInString(tc.s)
		res, err := src.AsInt()

		if tc.err == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, int(tc.expected), res, "res")
		} else {
			assert.NotNil(t, err, fmt.Sprintf("err | %v", tc.err))
		}
	}

	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsInt()

		if s.e == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.expected, res, "res")
		} else {
			assert.NotNil(t, err, fmt.Sprintf("err | %v", s.e))
		}
	}
}

func TestNumberInString_AsInt64(t *testing.T) {

	type testCase struct {
		s   string
		exp int64
		e   error
	}

	cases := []testCase{
		{
			s:   "-2",
			exp: -2,
			e:   nil,
		},
		{
			s:   "-9223372036854775807", // maxInt64
			exp: -9223372036854775807,
			e:   nil,
		},
		{
			s:   "9223372036854775807",
			exp: 9223372036854775807,
			e:   nil,
		},
	}

	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsInt64()

		if s.e == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.exp, res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
}

func TestNumberInString_AsUint64(t *testing.T) {

	type testCase struct {
		s   string
		exp uint64
		e   error
	}

	cases := []testCase{
		{
			s:   "18446744073709551615", // maxUint64
			exp: 18446744073709551615,
			e:   nil,
		},
		{
			s:   "-100",
			exp: 0,
			e:   fmt.Errorf("must overflow minUint64"),
		},
		{
			s:   "-1",
			exp: 0,
			e:   fmt.Errorf("must overflow minUint64"),
		},
	}

	for _, tc := range commonCases {

		src := NumberInString(tc.s)
		res, err := src.AsUint64()

		if tc.err == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, uint64(tc.expected), res, "res")
		} else {
			assert.NotNil(t, err, fmt.Sprintf("err | %v", tc.err))
		}
	}

	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsUint64()

		if s.e == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.exp, res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
}

func TestNumberInString_AsUint(t *testing.T) {

	type testCase struct {
		s   string
		exp uint
		e   error
	}

	cases := []testCase{
		{
			s:   "4294967295",
			exp: 4294967295,
			e:   nil,
		},
		{
			s:   "9223372036854775807",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
		{
			s:   "-1",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
		{
			s:   "-2",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
		{
			s:   "-10",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
	}

	for _, s := range commonCases {

		src := NumberInString(s.s)
		res, err := src.AsUint()

		if s.err == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.expected, res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsUint()

		if s.e == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.exp, res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
}

func TestNumberInString_AsInt8(t *testing.T) {

	type testCase struct {
		s   string
		exp int8
		e   error
	}

	cases := []testCase{
		{
			s:   "100",
			exp: 100,
			e:   nil,
		},
		{
			s:   "127",
			exp: 127,
			e:   nil,
		},
		{
			s:   "-127",
			exp: -127,
			e:   nil,
		},
		{
			s:   "200",
			exp: 0,
			e:   fmt.Errorf("overflow maxInt8"),
		},
		{
			s:   "-150",
			exp: 0,
			e:   fmt.Errorf("overflow minInt8"),
		},
		{
			s:   "9223372036854775807",
			exp: 0,
			e:   fmt.Errorf("Overflow maxInt8"),
		},
	}

	for _, s := range commonCases {

		src := NumberInString(s.s)
		res, err := src.AsInt8()

		if s.err == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, int8(s.expected), res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsInt8()

		if nil == s.e {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.exp, res, "res")
		} else {
			assert.NotNil(t, err, "err")
		}
	}
}
