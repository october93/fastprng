package fastprng

import (
// "log"
// "encoding/binary"
)

type MaD0Source struct {
	S   []uint64
	a   uint64
	b   uint64
	c   uint64
	d   uint64
	T   []uint64
	pos uint16
}

const Tlen = 512

func (mad *MaD0Source) Seed(seed int64) {
	marcbb := NewMARCBB(seed)

	mad.a = marcbb.UInt64()
	mad.b = marcbb.UInt64()
	mad.c = marcbb.UInt64()
	mad.d = marcbb.UInt64()

	mad.S = make([]uint64, len(marcbb.S))

	for pos, val := range marcbb.S {
		mad.S[pos] = uint64(val)
	}

	mad.T = make([]uint64, Tlen)
	mad.pos = Tlen
}

func NewMaD0(seed int64) *MaD0Source {
	mad := MaD0Source{}
	mad.Seed(seed)
	return &mad
}

func (r *MaD0Source) Generate() {
	r.a = r.a + r.c
	r.b = r.b + r.d

	ta := r.a
	tb := r.b

	for i := 0; i <= 31; i++ {
		r.c = r.c ^ (r.S[i] + r.a)
		r.T[2*i] = r.c
		r.c += (ta ^ tb)
		r.d ^= (r.c + r.b)
		// ta = ta <<< 3
		ta = circular_shift_left(ta, 3)

		r.d += (ta ^ tb)
		r.T[(2*i)+1] = r.d
		r.S[i] = r.d
		// tb = tb >>> 5
		tb = circular_shift_right(tb, 5)
	}
}

func (r *MaD0Source) Next() uint64 {
	if r.pos >= Tlen {
		r.Generate()
		r.pos = 0
	}
	data := r.T[r.pos]
	r.pos++
	return data
}

func (r *MaD0Source) Int63() int64 {
	n := int64(r.Next())
	n &= 0x7FFFFFFFFFFFFFFF
	return n
}

func (r *MaD0Source) UInt64() uint64 {
	n := uint64(r.Next())<<32 | uint64(r.Next())
	return n
}
