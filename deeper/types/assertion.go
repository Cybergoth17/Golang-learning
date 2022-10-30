package types

import "fmt"

func Assertion() {
	var iface interface{} = Integer(100)
	t, ok := iface.(Integer)
	fmt.Printf("Ok? %v, Value: %v, Type:%T\n", ok, t, t)
	iface = "hello"
	t, ok = iface.(Integer)
	fmt.Printf("Ok? %v, Value: %v, Type:%T\n", ok, t, t)

}
