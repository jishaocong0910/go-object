package o

func NewMap[K comparable, V any]() *Map[K, V] {
	m := &Map[K, V]{}
	m.m_Map = extendMap[K, V](m)
	return m
}

type Map[K comparable, V any] struct {
	*m_Map[K, V]
}
