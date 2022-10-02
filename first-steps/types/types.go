package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// unsigned int: uint8, uint16, uint32, uint64

	var baseSalary uint32 = 1000

	fmt.Println("baseSalary:", baseSalary)

	// signed int: int8, int16, int32, int64

	var maxInt32 int32 = math.MaxInt32
	var minInt32 int32 = math.MinInt32

	fmt.Println("int 32 range:", minInt32, "to", maxInt32)

	// rune = int32 that represents a Unicode code point
	var myChar rune = '1'
	fmt.Println("myChar:", myChar)

	// float32, float64
	// var x float32 = 49.99
	var x = float32(49.99)

	fmt.Println("type of x:", reflect.TypeOf(x))
	fmt.Println("literal 49.99:", reflect.TypeOf(49.99))

	// bool
	var isTrue bool = true
	fmt.Println("isTrue:", isTrue)
	fmt.Println("type of isTrue:", reflect.TypeOf(isTrue))

	// string, delimitated by double quotes only
	var name string = "Igor"
	fmt.Println("name:", name)
	fmt.Println("type of name:", reflect.TypeOf(name))
	fmt.Println("length of name:", len(name))

	// string with multiple lines
	var text = `This is a text
				with multiple lines
				and it is a raw string literal
				that is not interpreted
				and it is not escaped
				"Hello, World!"`

	fmt.Println("text:", text)

	// char is not a type in Go
	// char is a rune, int32
	var char byte = 'b'
	fmt.Println("char:", char)
}
