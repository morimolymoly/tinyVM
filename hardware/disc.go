package hardware

import "fmt"

// Disc ... disc
type Disc struct {
	content []byte
}

// Init ... Init
func (d *Disc) Init() {
	fmt.Println("[*] Init Disc...")
	d.content = make([]byte, 10000000)
}
