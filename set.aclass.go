package o

func extendSet[T comparable](i i_Set[T], m i_Map[T, any]) *m_Set[T] {
	return &m_Set[T]{i: i, m: m}
}

type i_Set[T comparable] interface {
	mSet() *m_Set[T]
}

type m_Set[T comparable] struct {
	i i_Set[T]
	m i_Map[T, any]
}

func (this *m_Set[T]) mSet() *m_Set[T] {
	return this
}

func (this *m_Set[T]) Add(es ...T) {
	for _, e := range es {
		this.m.mMap().Put(e, struct{}{})
	}
}

func (this *m_Set[T]) AddSet(i i_Set[T]) {
	for k := range i.mSet().m.mMap().m {
		this.m.mMap().Put(k, struct{}{})
	}
}

func (this *m_Set[T]) Remove(e T) {
	this.m.mMap().Remove(e)
}

func (this *m_Set[T]) Contains(ts ...T) bool {
	return this.m.mMap().ContainsKeys(ts...)
}

func (this *m_Set[T]) ContainsAny(ts ...T) bool {
	return this.m.mMap().ContainsAnyKey(ts...)
}

func (this *m_Set[T]) Len() int {
	return this.m.mMap().Len()
}

func (this *m_Set[T]) Empty() bool {
	return this.m.mMap().Empty()
}

func (this *m_Set[T]) Raw() []T {
	return this.m.mMap().Keys()
}
