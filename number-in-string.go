package numconvert

import (
	"fmt"
	"math/big"
	"strconv"
)

//
// NumberInString is an alias for
// convert string to numbers
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
		return 0, fmt.Errorf("Cannot convert to uint64 from: %v", s)
	}

	return bi.Uint64(), nil
}

func (s NumberInString) AsInt64() (int64, error) {

	bi := big.NewInt(0)
	if _, ok := bi.SetString(string(s), 10); !ok {
		return 0, fmt.Errorf("Cannot convert to int64 from: %v", s)
	}

	return bi.Int64(), nil
}

func (s NumberInString) AsInt() (int, error) {
	return s.AsInt32()
}

func (s NumberInString) AsInt32() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to int from: %v | %v", s, err)
	}
	if res > MaxInt32 || res < -MaxInt32 {
		return 0, fmt.Errorf("Overflow maxInt32:%v", s)
	}

	return res, nil
}

func (s NumberInString) AsUint() (uint, error) {
	return s.AsUint32()
}
func (s NumberInString) AsUint32() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to uint from: %v | %v", s, err)
	}
	if res > MaxUint32 || res < 0 {
		return 0, fmt.Errorf("Overflow maxUnt32:%v", s)
	}

	return uint(res), nil
}

func (s NumberInString) AsUint16() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to uint from: %v | %v", s, err)
	}
	if res > MaxUint16 || res < 0 {
		return 0, fmt.Errorf("Overflow maxUint16:%v", s)
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

func (s NumberInString) AsInt16() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to int from: %v | %v", s, err)
	}
	if res > MaxInt16 || res < -MaxInt16 {
		return 0, fmt.Errorf("Overflow maxInt16:%v", s)
	}

	return res, nil
}

func (s NumberInString) AsInt8() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("Cannot convert to int from: %v | %v", s, err)
	}
	if res > MaxInt8 || res < -MaxInt8 {
		return 0, fmt.Errorf("Overflow maxInt8:%v", s)
	}

	return res, nil
}
