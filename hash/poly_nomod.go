package hash

import "math/big"

type PolyHashBig struct {
	base  *big.Int
	power *big.Int
}

func NewPolyHashBig(base int64) *PolyHashBig {
	return &PolyHashBig{
		base: big.NewInt(base),
	}
}

func (h *PolyHashBig) Init(s string) *big.Int {
	hash := big.NewInt(0)
	h.power = big.NewInt(1)

	// power = base^(m-1)
	for i := 0; i < len(s)-1; i++ {
		h.power.Mul(h.power, h.base)
	}

	for i := 0; i < len(s); i++ {
		hash.Mul(hash, h.base)
		hash.Add(hash, big.NewInt(int64(s[i])))
	}

	return hash
}

func (h *PolyHashBig) Roll(prev *big.Int, left, right byte) *big.Int {
	hash := new(big.Int).Set(prev)

	// hash = hash - left * power
	tmp := new(big.Int).Mul(big.NewInt(int64(left)), h.power)
	hash.Sub(hash, tmp)

	// hash = hash * base + right
	hash.Mul(hash, h.base)
	hash.Add(hash, big.NewInt(int64(right)))

	return hash
}

func (h *PolyHashBig) Name() string {
	return "poly_nomod"
}