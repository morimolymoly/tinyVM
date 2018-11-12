package vm

import "fmt"

func cHlt() bool {
	fmt.Println("HLT!")
	return false
}

func badInstruction() bool {
	fmt.Println("Bad Instruction")
	return true
}
