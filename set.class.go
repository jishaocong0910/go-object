package o

func NewSet[T comparable](es ...T) *Set[T] {
	s := &Set[T]{}
	s.m_Set = extendSet[T](s, NewMap[T, any]())
	s.Add(es...)
	return s
}

type Set[T comparable] struct {
	*m_Set[T]
}
