package o

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrKeyMap(t *testing.T) {
	r := require.New(t)
	m := NewStrKeyMap[string](true)
	m.Put("aaa", "111")
	m.Put("AAA", "111")
	r.Equal(2, m.Len())

	m2 := NewStrKeyMap[string](false)
	m2.Put("aaa", "111")
	r.Equal(m2.Get("aaa"), "111")
	m2.Put("AAA", "222")
	r.Equal(m2.Get("aaa"), "222")
	r.Equal(1, m2.Len())
	m2.Put("bbb", "333")
	m2.Put("BBB", "444")
	r.Contains(m2.Keys(), "AAA")
	r.Contains(m2.Keys(), "BBB")

	mm := m2.Get("ccc")
	fmt.Printf("%T\n", mm)
	r.Equal("", m2.Get("ccc"))

	m2.Remove("bBb")
	r.Equal(1, m2.Len())

	m3 := NewStrKeyMap[string](true)
	m3.Put("aaa", "111")
	m3.Put("Aaa", "222")
	m3.Put("CCC", "333")
	r.Equal(3, m3.Len())
	m2.PutAll(m3)
	r.Equal(2, m2.Len())

	toMap := m3.Raw()
	r.Equal(3, len(toMap))
	toMap["aaa"] = "111"
	toMap["Aaa"] = "222"
	toMap["CCC"] = "333"
}
