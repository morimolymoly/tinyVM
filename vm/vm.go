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
	reg    *Reg
	memory *Memory
	disc   *hw.Disc
}

// New ... Create New Virtual Machine
func New() VM {
	fmt.Println("[*] Creating VM...")
	vm := VM{
		reg:    &Reg{},
		memory: &Memory{},
		disc:   &hw.Disc{},
	}
	vm.memory.Init()
	vm.disc.Init()
	vm.memory.FetchFromDisc(vm.disc, 0x0, 0x4, 0x0)
	vm.reg.ip = 0x0
	return vm
}

// Run ... run
func (vm *VM) Run() {
	fmt.Println("[*] tinyVM start!")
	for {
		i := vm.fetchCode()
		if !vm.executeCode(i) {
			break
		}
	}
	fmt.Println("[*] tinyVM end!")
	return
}

func (vm *VM) executeCode(code []byte) bool {
	op := code[0] & 0xf0
	if op == 0x0 {
		return cHlt()
	}
	return badInstruction()
}

func (vm *VM) fetchCode() []byte {
	ret := vm.memory.GetContent(vm.reg.ip, 2)
	vm.reg.ip += 2
	return ret
}
