package o_test

import (
	"testing"

	. "github.com/jishaocong9010/go-object"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	r := require.New(t)
	m := NewMap[string, any]()
	m.Put("aaa", "111")
	r.Equal("111", m.Get("aaa"))
	m.Put("AAA", "222")
	r.Equal("111", m.Get("aaa"))
	r.Equal("222", m.Get("AAA"))

	r.Equal("bbb", m.GetIfAbsent("333", func(k string) any {
		return "bbb"
	}))
	r.Equal(3, m.Len())

	r.Equal("bbb", m.GetIfAbsent("333", func(k string) any {
		return "ccc"
	}))
	r.Equal(3, m.Len())
	r.True(m.ContainsKeys("aaa", "333"))
	r.False(m.ContainsKeys("aaa", "333", "x"))
	r.True(m.ContainsAnyKey("x", "y", "aaa"))
	r.False(m.ContainsAnyKey("x", "y", "z"))
}
