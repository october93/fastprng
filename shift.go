package fastprng

import (
// "log"
)

func circular_shift_left(input uint64, count uint) uint64 {
	input_mask := ((uint64(1) << count) - 1) << (64 - count)

	saved_data := (input & input_mask) >> (64 - count)

	output := input << count

	output |= saved_data

	return output
}

func circular_shift_right(input uint64, count uint) uint64 {
	input_mask := ((uint64(1) << count) - 1)

	saved_data := (input & input_mask) << (64 - count)

	output := input >> count
	output |= saved_data

	return output
}
