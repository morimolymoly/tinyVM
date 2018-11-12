package main

import (
	"fmt"

	vm "github.com/morimolymoly/tinyVM/vm"
)

func main() {
	fmt.Println("tinyVM")
	m := vm.New()
	m.Run()
}
