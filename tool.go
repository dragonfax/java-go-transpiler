package main

func mustByteListErr(buf []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buf
}
