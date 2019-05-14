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
		return 0, fmt.Errorf("cannot convert to bigInt | %v", s)
	}

	if "-0" == s {
		return 0, nil
	}

	if '-' == []rune(s)[0] {
		return 0, fmt.Errorf("negative overflow| %v", s)
	}

	return bi.Uint64(), nil
}

func (s NumberInString) AsInt64() (int64, error) {

	bi := big.NewInt(0)
	if _, ok := bi.SetString(string(s), 10); !ok {
		return 0, fmt.Errorf("cannot convert to int64 | %v", s)
	}

	return bi.Int64(), nil
}

func (s NumberInString) AsInt() (int, error) {
	res, err := s.AsInt32()
	return int(res), err
}

func (s NumberInString) AsInt32() (int32, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("cannot convert to int32 | %v | %v", s, err)
	}
	if res > MaxInt32 || res < -MaxInt32 {
		return 0, fmt.Errorf("overflow maxInt32 | %v", s)
	}

	return int32(res), nil
}

func (s NumberInString) AsUint() (uint, error) {
	res, err := s.AsUint32()
	return uint(res), err
}
func (s NumberInString) AsUint32() (uint32, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to uint from: %v | %v", s, err)
	}
	if res > MaxUint32 || res < 0 {
		return 0, fmt.Errorf("overflow maxUnt32 | %v", s)
	}

	return uint32(res), nil
}

func (s NumberInString) AsUint16() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to uint | %v | %v", s, err)
	}
	if res > MaxUint16 || res < 0 {
		return 0, fmt.Errorf("overflow maxUint16 | %v", s)
	}

	return uint(res), nil
}

func (s NumberInString) AsUint8() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to uint from: %v | %v", s, err)
	}
	if res > MaxUint8 || res < 0 {
		return 0, fmt.Errorf("Overflow maxUint8:%v", s)
	}

	return uint(res), nil
}

func (s NumberInString) AsInt16() (int16, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("cannot convert to int | %v | %v", s, err)
	}
	if res > MaxInt16 || res < -MaxInt16 {
		return 0, fmt.Errorf("overflow maxInt16 | %v", s)
	}

	return int16(res), nil
}

func (s NumberInString) AsInt8() (int8, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("cannot convert to int | %v | %v", s, err)
	}
	if res > MaxInt8 || res < -MaxInt8 {
		return 0, fmt.Errorf("overflow maxInt8 | %v", s)
	}

	return int8(res), nil
}
