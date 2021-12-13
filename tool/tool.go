package tool

import "reflect"

func MustByteListErr(buf []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return buf
}

func IsNilInterface(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}
