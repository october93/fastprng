package fastprng

import (
	"testing"
	// "log"
)

func TestCircularShiftLeft(t *testing.T) {
	input := uint64(0xff00440000000000)
	if circular_shift_left(input, 24) != uint64(0x0000000000ff0044) {
		t.Fatalf("Circular Shift Left Failed")
	}

}

func TestCircularShiftRight(t *testing.T) {
	input := uint64(0x0000000000ff0044)
	if circular_shift_right(input, 24) != uint64(0xff00440000000000) {
		t.Fatalf("Circular Shift Right Failed")
	}
}

// func TestCircularShiftRight(t *testing.T) {
//     input := uint64(1)
//     output := circular_shift_right(input, 65)
//     log.Println("INPUT", input, "OUTPUT", output)
// }
