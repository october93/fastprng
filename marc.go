package fastprng

import (
	"encoding/binary"
	"log"
)

type MARCRandSource struct {
	S             []uint8
	i             uint8
	j             uint8
	k             uint8
	ks_iterations uint16
}

func (z *MARCRandSource) KSA(key []uint8) {
	z.S = make([]uint8, 256)
	z.i = uint8(0)
	keylen := uint8(len(key))
	for {
		z.S[z.i] = uint8(z.i)
		z.i++
		if z.i == uint8(0) {
			break
		}
	}
	z.i = uint8(0)
	z.j = uint8(0)
	z.k = uint8(0)
	for r := uint16(0); r < z.ks_iterations; r++ {
		z.j = z.j + z.S[z.i] + key[z.i%keylen]
		z.k = z.k ^ z.j

		// left_rotate(S[i], S[j], S[k])
		tmp := z.S[z.i]
		z.S[z.i] = z.S[z.j]
		z.S[z.j] = z.S[z.k]
		z.S[z.k] = tmp

		z.i++
	}
}

func (z *MARCRandSource) prepare_iteration() {
	z.i = z.j + z.k
}

func (z *MARCRandSource) PRGA_iteration() []uint8 {
	output := make([]uint8, 4)

	z.i++
	z.j += z.S[z.i]
	z.k ^= z.j

	/* swap(S[i], S[j])after iteration k --- what does after iteration k mean here? */
	tmp := z.S[z.i]
	z.S[z.i] = z.S[z.j]
	z.S[z.j] = tmp

	m := z.S[z.j] + z.S[z.k]
	n := z.S[z.i] + z.S[z.j]

	output[0] = z.S[m]
	output[1] = z.S[n]
	output[2] = z.S[m^n]
	output[3] = z.S[n^z.k]
	return output
}

func (r *MARCRandSource) SeedWithIterationCount(seed int64, ks_iterations uint16) {
	_ = log.Println
	/* MaD0 accepts a seed/key that is no larger than 64 bytes (512 bits) */

	r.ks_iterations = ks_iterations
	// r.ks_iterations = 320 /* MARC-bb only uses 320 iterations */

	key := make([]uint8, 8)
	binary.LittleEndian.PutUint64(key, uint64(seed))
	r.KSA(key)
	r.prepare_iteration()
}

func (r *MARCRandSource) Seed(seed int64) {
	r.SeedWithIterationCount(seed, 576)
}

func (r *MARCRandSource) SeedMARCBB(seed int64) {
	r.SeedWithIterationCount(seed, 320)
}

func NewMARC(seed int64) *MARCRandSource {
	marc := MARCRandSource{}
	marc.Seed(seed)
	return &marc
}

func NewMARCBB(seed int64) *MARCRandSource {
	marc := MARCRandSource{}
	marc.SeedMARCBB(seed)
	return &marc
}

func (r *MARCRandSource) Next() uint32 {
	input := r.PRGA_iteration()
	data := binary.BigEndian.Uint32(input)
	return data
}

func (r *MARCRandSource) Int63() int64 {
	n := int64(r.Next())<<32 | int64(r.Next())
	n &= 0x7FFFFFFFFFFFFFFF
	return n
}

func (r *MARCRandSource) UInt64() uint64 {
	n := uint64(r.Next())<<32 | uint64(r.Next())
	return n
}
