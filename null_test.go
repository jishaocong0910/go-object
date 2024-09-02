package o

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNull(t *testing.T) {
	r := require.New(t)
	var a any
	r.True(a == nil)
	r.True(IsNull(a))
	var i *int
	a = i
	r.False(a == nil)
	r.True(IsNull(a))
	var i2 int
	a = i2
	r.False(a == nil)
	r.False(IsNull(a))
}

func TestNonNull(t *testing.T) {
	r := require.New(t)
	var a any
	r.True(a == nil)
	r.False(NotNull(a))
	var i *int
	a = i
	r.False(a == nil)
	r.False(NotNull(a))
	var i2 int
	a = i2
	r.False(a == nil)
	r.True(NotNull(a))
}
