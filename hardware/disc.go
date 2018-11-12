package hardware

import "fmt"

var startup = []byte{
	0x10, 0x20, 0x00, 0x00,
}

// Disc ... disc
type Disc struct {
	content []byte
}

// Init ... Init
func (d *Disc) Init() {
	fmt.Println("[*] Init Disc...")
	d.content = make([]byte, 0x100)
	fmt.Println("[*] Copying startup content...")
	copy(d.content, startup)
}

// GetContent ... get contents
func (d *Disc) GetContent(position, len int) []byte {
	return d.content[position : position+len]
}
