package o

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrSet(t *testing.T) {
	r := require.New(t)
	set := NewStrSet(false, "abC")
	set.Add("abc")
	set.Add("Abc")
	r.Equal(1, set.Len())
	r.True(set.Contains("aBc"))

	set2 := NewStrSet(false)
	set2.Add("abC")
	set2.Add("bcd")
	set.AddSet(set2)
	r.Equal(2, set.Len())

	slice := set2.Raw()
	r.Contains(slice, "abC")
	r.Contains(slice, "bcd")

	r.True(set.Contains("ABC"))
	set.Remove("abc")
	r.False(set.Contains("ABC"))
}
