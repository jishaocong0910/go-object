package o_test

import (
	"testing"

	. "github.com/jishaocong0910/go-object"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	r := require.New(t)
	s := NewSet[string]()
	r.True(s.Empty())
	r.Equal(0, s.Len())
	s2 := NewSet[string]("a", "b", "c")
	r.False(s2.Empty())
	r.Equal(3, s2.Len())

	s.Add("str")
	r.Equal(1, s.Len())
	r.True(s.Contains("str"))
	s.AddSet(s2)
	r.Equal(4, s.Len())
	validateElement(r, s, "a", "str", "b", "c")

	r.True(s.Contains("a", "b"))
	r.False(s.Contains("a", "b", "d"))
	r.True(s.ContainsAny("a", "d"))
	r.True(s.ContainsAny("a", "b", "d"))

	s.Remove("a")
	r.False(s.Contains("a"))
}

func validateElement[T comparable](r *require.Assertions, s *Set[T], contains ...T) {
	r.Equal(len(contains), s.Len(), "expected: %v, actual: %v", contains, s.Raw())
	for _, c := range contains {
		r.True(s.Contains(c), "not contains: %v, actual: %v", c, s.Raw())
	}
}
