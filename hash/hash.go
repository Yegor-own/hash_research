package hash


type Haser interface {
	Init(s string) uint64 // начальный подсчет хэша образца
	Roll(prev uint64, left byte, right byte) uint64 // реализация скользящего хэша
	Name() string // название используемой хэш-функции
}