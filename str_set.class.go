package o

func NewStrSet(caseSensitive bool, es ...string) *StrSet {
	s := &StrSet{}
	s.m_Set = extendSet[string](s, NewStrKeyMap[any](caseSensitive))
	s.Add(es...)
	return s
}

type StrSet struct {
	*m_Set[string]
}
