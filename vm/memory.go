package vm

import "fmt"

// Memory ... memory
type Memory struct {
	memory []byte
}

// Init ... init memory
func (m *Memory) Init() {
	fmt.Println("[*] Init memory...")
	m.memory = make([]byte, 100000)
}
