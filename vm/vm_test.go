package vm

import (
	"testing"
)

func TestExecuteCode(t *testing.T) {
	vm := VM{}
	ret := vm.executeCode([]byte{0x00, 0x00})
	if ret {
		t.Fatal("HLT faled")
	}

	ret = vm.executeCode([]byte{0x11, 0x11})
	if !ret {
		t.Fatal("Bad Instruction")
	}
}
