package tool

func MustByteListErr(buf []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buf
}
