package o

func extendSet[T comparable](i i_Set[T], m i_Map[T, any]) *m_Set[T] {
	return &m_Set[T]{i: i, m: m}
}

type i_Set[T comparable] interface {
	m_Set() *m_Set[T]
}

type m_Set[T comparable] struct {
	i i_Set[T]
	m i_Map[T, any]
}

func (this *m_Set[T]) m_Set() *m_Set[T] {
	return this
}

func (this *m_Set[T]) Add(es ...T) {
	for _, e := range es {
		this.m.m_Map().Put(e, struct{}{})
	}
}

func (this *m_Set[T]) AddSet(i i_Set[T]) {
	for k := range i.m_Set().m.m_Map().m {
		this.m.m_Map().Put(k, struct{}{})
	}
}

func (this *m_Set[T]) Remove(e T) {
	this.m.m_Map().Remove(e)
}

func (this *m_Set[T]) Contains(ts ...T) bool {
	return this.m.m_Map().ContainsKeys(ts...)
}

func (this *m_Set[T]) ContainsAny(ts ...T) bool {
	return this.m.m_Map().ContainsAnyKey(ts...)
}

func (this *m_Set[T]) Len() int {
	return this.m.m_Map().Len()
}

func (this *m_Set[T]) Empty() bool {
	return this.m.m_Map().Empty()
}

func (this *m_Set[T]) Raw() []T {
	return this.m.m_Map().Keys()
}
