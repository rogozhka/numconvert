package numconvert

import (
	"fmt"
	"math/big"
	"strconv"
)

//
// NumberInString is a string alias for
// converting to numbers
//
// example: NumberInString("123").AsUint()
//
type NumberInString string

//
// Int number limits
//
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func (s NumberInString) AsUint64() (uint64, error) {

	bi := big.NewInt(0)
	if _, ok := bi.SetString(string(s), 10); !ok {
		return 0, errorConvert("uint64", s, nil)
	}

	if "-0" == s {
		return 0, nil
	}

	if '-' == []rune(s)[0] {
		return 0, fmt.Errorf("negative overflow | %v", s)
	}

	return bi.Uint64(), nil
}

func (s NumberInString) AsInt64() (int64, error) {

	bi := big.NewInt(0)
	if _, ok := bi.SetString(string(s), 10); !ok {
		return 0, errorConvert("int64", s, fmt.Errorf("cannot SetString"))
	}

	return bi.Int64(), nil
}

func (s NumberInString) AsInt() (int, error) {
	res, err := s.AsInt32()
	return int(res), err
}

func (s NumberInString) AsInt32() (int32, error) {
	tag := "int32"

	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}

	if res > MaxInt32 || res < MinInt32 {
		return 0, errorRange(tag, s)
	}

	return int32(res), nil
}

func (s NumberInString) AsUint() (uint, error) {
	res, err := s.AsUint32()
	return uint(res), err
}

func (s NumberInString) AsUint32() (uint32, error) {
	tag := "uint32"

	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}
	if res > MaxUint32 || res < 0 {
		return 0, errorRange(tag, s)
	}

	return uint32(res), nil
}

func (s NumberInString) AsUint16() (uint16, error) {
	tag := "uint16"

	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}
	if res > MaxUint16 || res < 0 {
		return 0, errorRange("uint16", s)
	}

	return uint16(res), nil
}

func (s NumberInString) AsUint8() (uint8, error) {
	tag := "uint8"

	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}
	if res > MaxUint8 || res < 0 {
		return 0, errorRange(tag, s)
	}

	return uint8(res), nil
}

func (s NumberInString) AsInt16() (int16, error) {
	tag := "int16"
	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}
	if res > MaxInt16 || res < MinInt16 {
		return 0, errorRange(tag, s)
	}

	return int16(res), nil
}

func (s NumberInString) AsInt8() (int8, error) {
	tag := "int8"

	res, err := atoi(s, tag)
	if err != nil {
		return 0, err
	}
	if res > MaxInt8 || res < MinInt8 {
		return 0, errorRange(tag, s)
	}

	return int8(res), nil
}

func errorConvert(tag string, src interface{}, err error) error {
	return fmt.Errorf("cannot convert to %v | %v | %v", tag, src, err)
}

func errorRange(tag string, src interface{}) error {
	return fmt.Errorf("overflow %v | %v ", tag, src)
}

func atoi(s NumberInString, tag string) (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, errorConvert(tag, s, err)
	}
	return res, nil
}
