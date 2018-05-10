package numconvert

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestNumberInString_AsInt(t *testing.T) {

	type theCase struct {
		s   string
		exp int
		e   error
	}

	cases := []theCase{
		theCase{
			s:   "0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "0xaaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "--0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "++0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "-0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "1",
			exp: 1,
			e:   nil,
		},
		theCase{
			s:   "2",
			exp: 2,
			e:   nil,
		},
		theCase{
			s:   "-2",
			exp: -2,
			e:   nil,
		},
		theCase{
			s:   "-9223372036854775807",
			exp: 0,
			e:   fmt.Errorf("maxInt64 must overflow int32"),
		},
		theCase{
			s:   "9223372036854775807",
			exp: 0,
			e:   fmt.Errorf("maxInt64 must overflow int32"),
		},
	}

	for _, s := range cases {

		src := NumberInString(s.s)
		res, err := src.AsInt()

		if s.e == nil {
			assert.Nil(t, err, "err")
			assert.Equal(t, s.exp, res, "res")
		} else {
			assert.NotNil(t, err, fmt.Sprintf("err:%v", s.e))
		}
	}
}

func TestNumberInString_AsInt64(t *testing.T) {

	type theCase struct {
		s   string
		exp int64
		e   error
	}

	cases := []theCase{
		theCase{
			s:   "0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "0xaaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "--0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "++0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "-0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "1",
			exp: 1,
			e:   nil,
		},
		theCase{
			s:   "2",
			exp: 2,
			e:   nil,
		},
		theCase{
			s:   "-2",
			exp: -2,
			e:   nil,
		},
		theCase{
			s:   "-9223372036854775807", // maxInt64
			exp: -9223372036854775807,
			e:   nil,
		},
		theCase{
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

	type theCase struct {
		s   string
		exp uint64
		e   error
	}

	cases := []theCase{
		theCase{
			s:   "0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "0xaaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "--0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "++0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "-0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "1",
			exp: 1,
			e:   nil,
		},
		theCase{
			s:   "2",
			exp: 2,
			e:   nil,
		},
		theCase{
			s:   "18446744073709551615",
			exp: 18446744073709551615,
			e:   nil,
		},
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

	type theCase struct {
		s   string
		exp uint
		e   error
	}

	cases := []theCase{
		theCase{
			s:   "0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "0xaaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-aaa",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "-",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "--0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "++0",
			exp: 0,
			e:   fmt.Errorf("nonformat"),
		},
		theCase{
			s:   "+0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "-0",
			exp: 0,
			e:   nil,
		},
		theCase{
			s:   "1",
			exp: 1,
			e:   nil,
		},
		theCase{
			s:   "2",
			exp: 2,
			e:   nil,
		},
		theCase{
			s:   "4294967295",
			exp: 4294967295,
			e:   nil,
		},
		theCase{
			s:   "9223372036854775807",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
		theCase{
			s:   "-1",
			exp: 0,
			e:   fmt.Errorf("Overflow maxUint32"),
		},
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
