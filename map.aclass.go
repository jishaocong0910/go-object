package o

func extendMap[K comparable, V any](i i_Map[K, V]) *m_Map[K, V] {
	return &m_Map[K, V]{i: i, m: map[K]*Entry[K, V]{}}
}

type i_Map[K comparable, V any] interface {
	m_Map_() *m_Map[K, V]
	key(K) K
}

type m_Map[K comparable, V any] struct {
	i i_Map[K, V]
	m map[K]*Entry[K, V]
}

func (this *m_Map[K, V]) m_Map_() *m_Map[K, V] {
	return this
}

func (this *m_Map[K, V]) key(k K) K {
	return k
}

func (this *m_Map[K, V]) Put(k K, v V) {
	this.m[this.i.key(k)] = &Entry[K, V]{k, v}
}

func (this *m_Map[K, V]) PutAll(other i_Map[K, V]) {
	for k, v := range other.m_Map_().m {
		this.m[this.i.key(k)] = v
	}
}

func (this *m_Map[K, V]) GetEntry(k K) *Entry[K, V] {
	return this.m[this.i.key(k)]
}

func (this *m_Map[K, V]) Get(k K) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	}
	return
}

func (this *m_Map[K, V]) GetIfAbsent(k K, f func(k K) V) (v V) {
	entry := this.GetEntry(k)
	if entry != nil {
		v = entry.Value
	} else {
		v = f(k)
		this.Put(k, v)
	}
	return
}

func (this *m_Map[K, V]) Remove(k K) bool {
	if _, ok := this.m[k]; ok {
		delete(this.m, this.i.key(k))
		return true
	}
	return false
}

func (this *m_Map[K, V]) RemoveAll(ks ...K) {
	for _, k := range ks {
		delete(this.m, this.i.key(k))
	}
}

func (this *m_Map[K, V]) ContainsKeys(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; !ok {
			return false
		}
	}
	return true
}

func (this *m_Map[K, V]) ContainsAnyKey(ks ...K) bool {
	for _, k := range ks {
		if _, ok := this.m[this.i.key(k)]; ok {
			return true
		}
	}
	return false
}

func (this *m_Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(this.m))
	for _, e := range this.m {
		keys = append(keys, e.Key)
	}
	return keys
}

func (this *m_Map[K, V]) Values() []V {
	values := make([]V, 0, len(this.m))
	for _, e := range this.m {
		values = append(values, e.Value)
	}
	return values
}

func (this *m_Map[K, V]) Len() int {
	return len(this.m)
}

func (this *m_Map[K, V]) Empty() bool {
	return len(this.m) == 0
}

func (this *m_Map[K, V]) Raw() map[K]V {
	raw := make(map[K]V, len(this.m))
	for _, e := range this.m {
		raw[e.Key] = e.Value
	}
	return raw
}

func (this *m_Map[K, V]) Range(f func(k K, v V)) {
	for _, e := range this.m {
		f(e.Key, e.Value)
	}
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
