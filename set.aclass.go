package o

func extendSet[T comparable](i i_Set[T], m i_Map[T, any]) *m_Set[T] {
	return &m_Set[T]{i: i, m: m}
}

type i_Set[T comparable] interface {
	m_0685668D24AC() *m_Set[T]
}

type m_Set[T comparable] struct {
	i i_Set[T]
	m i_Map[T, any]
}

func (this *m_Set[T]) m_0685668D24AC() *m_Set[T] {
	return this
}

func (this *m_Set[T]) Add(es ...T) {
	for _, e := range es {
		this.m.m_8CC656601595().Put(e, struct{}{})
	}
}

func (this *m_Set[T]) AddSet(i i_Set[T]) {
	for k := range i.m_0685668D24AC().m.m_8CC656601595().m {
		this.m.m_8CC656601595().Put(k, struct{}{})
	}
}

func (this *m_Set[T]) Remove(e T) {
	this.m.m_8CC656601595().Remove(e)
}

func (this *m_Set[T]) Contains(ts ...T) bool {
	return this.m.m_8CC656601595().ContainsKeys(ts...)
}

func (this *m_Set[T]) ContainsAny(ts ...T) bool {
	return this.m.m_8CC656601595().ContainsAnyKey(ts...)
}

func (this *m_Set[T]) Len() int {
	return this.m.m_8CC656601595().Len()
}

func (this *m_Set[T]) Empty() bool {
	return this.m.m_8CC656601595().Empty()
}

func (this *m_Set[T]) Raw() []T {
	return this.m.m_8CC656601595().Keys()
}
