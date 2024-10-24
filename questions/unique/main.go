package main

import (
	"fmt"
	"reflect"
	"unique"
)

func main() {
	handler := unique.Make[int](1)
	v := handler.Value()

	fmt.Println(reflect.TypeOf(v))
}
