package hardware

import (
	"fmt"
)

// Memory ... memory
type Memory struct {
	memory []byte
}

// Init ... init memory
func (m *Memory) Init() {
	fmt.Println("[*] Init memory...")
	m.memory = make([]byte, 0x200)
}

// Dump ... dump memory
func (m *Memory) Dump() {
	fmt.Println(m.memory)
}

// FetchFromDisc ... Fetch data from Disc and fill memory
func (m *Memory) FetchFromDisc(d *Disc, position, len, moffset int) {
	copy(m.memory[moffset:], d.GetContent(position, len))
}

// GetContent ... get memory contents
func (m *Memory) GetContent(position, len int) []byte {
	return m.memory[position : position+len]
}
