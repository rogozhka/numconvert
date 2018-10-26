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
		return 0, convertErrorf(s, "uint64", nil)
	}

	return bi.Uint64(), nil
}

func (s NumberInString) AsInt64() (int64, error) {

	bi := big.NewInt(0)
	if _, ok := bi.SetString(string(s), 10); !ok {
		return 0, convertErrorf(s, "int64", nil)
	}

	return bi.Int64(), nil
}

func (s NumberInString) AsInt() (int, error) {
	return s.AsInt32()
}

func (s NumberInString) AsInt32() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "int", err)
	}
	if res > MaxInt32 || res < -MaxInt32 {
		return 0, overflowErrorf(s, "MaxInt32")
	}

	return res, nil
}

func (s NumberInString) AsUint() (uint, error) {
	return s.AsUint32()
}
func (s NumberInString) AsUint32() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "uint32", nil)
	}
	if res > MaxUint32 || res < 0 {
		return 0, overflowErrorf(s, "MaxUint32")
	}

	return uint(res), nil
}

func (s NumberInString) AsUint16() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "uint16", nil)
	}
	if res > MaxUint16 || res < 0 {
		return 0, overflowErrorf(s, "MaxUint16")
	}

	return uint(res), nil
}

func (s NumberInString) AsUint8() (uint, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "uint8", nil)
	}
	if res > MaxUint8 || res < 0 {
		return 0, overflowErrorf(s, "MaxUint8")
	}

	return uint(res), nil
}

func (s NumberInString) AsInt16() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "int16", nil)
	}
	if res > MaxInt16 || res < -MaxInt16 {
		return 0, overflowErrorf(s, "MaxInt16")
	}

	return res, nil
}

func (s NumberInString) AsInt8() (int, error) {

	res, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, convertErrorf(s, "int8", nil)
	}
	if res > MaxInt8 || res < -MaxInt8 {
		return 0, overflowErrorf(s, "MaxInt8")
	}

	return res, nil
}

func convertErrorf(from interface{}, to string, err error) error {

	var strErr string
	if err != nil {
		strErr = err.Error()
	}

	return fmt.Errorf("Cannot convert to | %s | %s | %s", to, from, strErr)
}

func overflowErrorf(from interface{}, subj string) error {
	return fmt.Errorf("Overflow | %s | %s", subj, from)
}
