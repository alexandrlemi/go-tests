package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var sp *string

	var s string
	s= "test string"

	sp=&s

	runes := []rune(s)
	firstRPtr:=  &runes[0]
	secondCharPtr := &runes[1]
	fmt.Println(firstRPtr)
	fmt.Println(secondCharPtr)
	
	ptr := unsafe.Pointer(firstRPtr)
	addr := uintptr(ptr)
	newAddr := addr + (2)

	newPtr := unsafe.Pointer(newAddr)

	value := *(*rune)(newPtr)

	fmt.Println(value)
	fmt.Printf("Символ: %c\n", value)
	strTestP(sp)
	fmt.Println(s)

	///////
	/////

}

func strTestP(str *string) {
	*str = "test from func"
}

func strTest(str string) {
	str = "test from func"
}