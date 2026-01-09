package hash

import "math/big"

type Hasher interface {
	Init(s string) uint64 // начальный подсчет хэша образца
	Roll(prev uint64, left byte, right byte) uint64 // реализация скользящего хэша
	Name() string // название используемой хэш-функции
}

type BigHasher interface {
	Init(s string) *big.Int
	Roll(prev *big.Int, left, right byte) *big.Int
	Name() string
}