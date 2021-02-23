package util

import "unsafe"

func Str2Byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{x[0], x[1], x[1]}))
}

func Byte2Str(b []byte) string {
	x := (*[3]uintptr)(unsafe.Pointer(&b))
	return *(*string)(unsafe.Pointer(&[3]uintptr{x[0], x[1]}))
}
