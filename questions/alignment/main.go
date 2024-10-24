package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	A int32  // 4 bytes
	B int64  // 8 bytes
	C string // 16 byte, 8 alignment
}

func main() {
	var s MyStruct

	s = MyStruct{
		A: 1,
		B: 2,
		C: "a",
	}
	a := "a"
	var emptyStr string
	char := 'a'
	fmt.Println("Size of MyStruct:", unsafe.Sizeof(s))        // Should print 32
	fmt.Println("Size of MyStruct.C:", unsafe.Sizeof(s.C))    // Should print 16
	fmt.Println("Size of a:", unsafe.Sizeof(a))               // Should print 16
	fmt.Println("Size of emptyStr:", unsafe.Sizeof(emptyStr)) // Should print 16
	fmt.Println("Size of char:", unsafe.Sizeof(char))         // Should print 4
	fmt.Println("Alignment of MyStruct:", unsafe.Alignof(s))  // Should print 8
	fmt.Println("Offset of A:", unsafe.Offsetof(s.A))         // Should print 0
	fmt.Println("Offset of B:", unsafe.Offsetof(s.B))         // Should print 8
	fmt.Println("Offset of C:", unsafe.Offsetof(s.C))         // Should print 16
}
