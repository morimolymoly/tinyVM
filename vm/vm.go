package vm

import (
	"fmt"

	hw "github.com/morimolymoly/tinyVM/hardware"
)

// Reg ... Registers
type Reg struct {
	ip int
	sp int
	r1 int
	r2 int
	r3 int
	r4 int
	r5 int
	r6 int
}

// VM ... Virtual Machine
type VM struct {
	reg    Reg
	memory Memory
	disc   hw.Disc
}

// New ... Create New Virtual Machine
func New() VM {
	fmt.Println("[*] Creating VM...")
	vm := VM{}
	vm.memory.Init()
	vm.disc.Init()
	return vm
}
