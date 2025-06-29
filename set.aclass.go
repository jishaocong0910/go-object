package o

func extendSet[T comparable](i i_Set[T], im i_Map[T, any]) *m_Set[T] {
	return &m_Set[T]{i: i, im: im}
}

type i_Set[T comparable] interface {
	m_Set_() *m_Set[T]
}

type m_Set[T comparable] struct {
	i  i_Set[T]
	im i_Map[T, any]
}

func (this *m_Set[T]) m_Set_() *m_Set[T] {
	return this
}

func (this *m_Set[T]) Add(es ...T) {
	for _, e := range es {
		this.im.m_Map_().Put(e, struct{}{})
	}
}

func (this *m_Set[T]) AddSet(i i_Set[T]) {
	for k := range i.m_Set_().im.m_Map_().m {
		this.im.m_Map_().Put(k, struct{}{})
	}
}

func (this *m_Set[T]) Remove(e T) bool {
	return this.im.m_Map_().Remove(e)
}

func (this *m_Set[T]) RemoveAll(e ...T) {
	this.im.m_Map_().RemoveAll(e...)
}

func (this *m_Set[T]) Contains(ts ...T) bool {
	return this.im.m_Map_().ContainsKeys(ts...)
}

func (this *m_Set[T]) ContainsAny(ts ...T) bool {
	return this.im.m_Map_().ContainsAnyKey(ts...)
}

func (this *m_Set[T]) Len() int {
	return this.im.m_Map_().Len()
}

func (this *m_Set[T]) Empty() bool {
	return this.im.m_Map_().Empty()
}

func (this *m_Set[T]) Raw() []T {
	return this.im.m_Map_().Keys()
}

func (this *m_Set[T]) Range(f func(t T)) {
	this.im.m_Map_().Range(func(k T, v any) {
		f(k)
	})
}
