package hash

type DJB2Hash struct {
	base  uint64
	power uint64
}

func NewDJB2Hash() *DJB2Hash {
	return &DJB2Hash{base: 33}
}

func (h *DJB2Hash) Init(s string) uint64 {
	var hash uint64 = 5381
	h.power = 1

	for i := 0; i < len(s)-1; i++ {
		h.power *= h.base
	}

	for i := 0; i < len(s); i++ {
		hash = hash*h.base + uint64(s[i])
	}
	return hash
}

func (h *DJB2Hash) Roll(prev uint64, left, right byte) uint64 {
	hash := prev - uint64(left)*h.power
	return hash*h.base + uint64(right)
}

func (h *DJB2Hash) Name() string {
	return "djb2"
}
