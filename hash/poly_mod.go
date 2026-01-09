package hash

type PolyHash struct {
	base  uint64
	mod   uint64
	power uint64
}

func NewPolyHash(base, mod uint64) *PolyHash {
	return &PolyHash{
		base: base,
		mod:  mod,
	}
}

func (h *PolyHash) Init(s string) uint64 {
	var hash uint64 = 0
	h.power = 1

	for i := 0; i < len(s)-1; i++ {
		h.power = (h.power * h.base) % h.mod
	}

	for i := 0; i < len(s); i++ {
		hash = (hash*h.base + uint64(s[i])) % h.mod
	}

	return hash
}

func (h *PolyHash) Roll(prev uint64, left byte, right byte) uint64 {
	hash := prev + h.mod
	// в go модуль от отрицательного числа дает отрицательный результат поэтому необходимо прибавлять его
	hash = (hash - uint64(left)*h.power%h.mod) % h.mod
	hash = (hash*h.base + uint64(right)) % h.mod
	return hash
}

func (h *PolyHash) Name() string {
	return "poly_mod"
}
