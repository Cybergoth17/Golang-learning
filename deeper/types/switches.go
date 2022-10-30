package types

import "fmt"

func Switches(i interface{}) {
	switch v := i.(type) {
	case Integer:
		fmt.Printf("Integer: %v\n", v)
	case String:
		fmt.Printf("string: %v\n", v)
	default:
		fmt.Printf("unknown: %v\n", v)
	}
}
