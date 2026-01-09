package hash

type SumHash struct{}

func (h *SumHash) Init(s string) uint64 {
	var sum uint64
	for i := 0; i < len(s); i++ {
		sum += uint64(s[i])
	}
	return sum
}

func (h *SumHash) Roll(prev uint64, left byte, right byte) uint64 {
	return prev - uint64(left) + uint64(right)
}

func (h *SumHash) Name() string {
	return "sum"
}
