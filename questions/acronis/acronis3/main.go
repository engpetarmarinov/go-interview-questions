package main

import (
	"fmt"
	"io"
	"reflect"
)

type test struct {
}

func (t *test) Read(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	var a = 2
	typeA := reflect.TypeOf(a)
	readerType := reflect.TypeOf((*io.Reader)(nil)).Elem()
	fmt.Println("typeA implements reader int", typeA.Implements(readerType))
	typeTest := reflect.TypeOf(&test{})
	fmt.Println("test implements reader int", typeTest.Implements(readerType))
	fmt.Println("typeA:", typeA)
}
