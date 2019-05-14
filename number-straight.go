package numconvert

type InString string

func (s InString) Int() int {
	n, err := NumberInString(s).Int()
	checkErrPanic(err)

	return n
}

func (s InString) Uint() uint {
	n, err := NumberInString(s).Uint()
	checkErrPanic(err)

	return n
}

func (s InString) Int64() int64 {
	n, err := NumberInString(s).Int64()
	checkErrPanic(err)

	return n
}

func (s InString) Uint64() uint64 {
	n, err := NumberInString(s).Uint64()
	checkErrPanic(err)

	return n
}

func (s InString) Int32() int32 {
	n, err := NumberInString(s).Int32()
	checkErrPanic(err)

	return n
}

func (s InString) Uint32() uint32 {
	n, err := NumberInString(s).Uint32()
	checkErrPanic(err)

	return n
}

func (s InString) Int16() int16 {
	n, err := NumberInString(s).Int16()
	checkErrPanic(err)

	return n
}

func (s InString) Uint16() uint16 {
	n, err := NumberInString(s).Uint16()
	checkErrPanic(err)

	return n
}

func (s InString) Int8() int8 {
	n, err := NumberInString(s).Int8()
	checkErrPanic(err)

	return n
}

func (s InString) Uint8() uint8 {
	n, err := NumberInString(s).Uint8()
	checkErrPanic(err)

	return n
}

func checkErrPanic(err error) {
	if nil != err {
		panic(err)
	}
}
