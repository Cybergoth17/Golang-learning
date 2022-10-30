package main

import (
	"fmt"
	"test/types"
)

func main() {
	//interface
	var iface types.Interface
	iface = &types.String{"hello world"}
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)
	iface = types.Integer(100)
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)
	//assertion
	types.Assertion()
	types.Switches("108880")
	types.Switches(types.Integer(100))
	types.Switches(types.String{"hello world"})
	types.Switches("testing again")
}
