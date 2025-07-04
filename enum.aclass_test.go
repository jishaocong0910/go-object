package o_test

import (
	"testing"

	. "github.com/jishaocong0910/go-object"
	"github.com/stretchr/testify/require"
)

type Animal struct {
	*M_EnumElem
}

func TestEnum(t *testing.T) {
	enum := NewEnum[Animal](struct {
		*M_Enum[Animal]
		cat, DOG, bird Animal
	}{
		cat:  Animal{},
		DOG:  Animal{},
		bird: Animal{},
	})

	r := require.New(t)
	r.True(enum.Is(Animal{}, Animal{}))
	r.True(enum.Is(enum.OfId("cat"), enum.cat))
	r.True(enum.Is(enum.OfId("DOG"), enum.DOG))
	r.True(enum.Is(enum.OfId("bird"), enum.bird))
	r.Equal("cat", enum.cat.ID())
	r.Equal("DOG", enum.DOG.ID())
	r.Equal("bird", enum.bird.ID())
	r.True(enum.Is(enum.cat, enum.cat, enum.DOG))
	r.False(enum.Is(enum.cat, enum.bird, enum.DOG))
	r.True(enum.Not(enum.cat, enum.bird, enum.DOG))
	r.False(enum.Not(enum.cat, enum.cat, enum.DOG))
	r.True(enum.OfId("SNAKE").Undefined())
	r.True(enum.OfIdIgnoreCase("SNAKE").Undefined())
	r.True(enum.OfId("BIRD").Undefined())
	r.False(enum.OfIdIgnoreCase("BIRD").Undefined())

	values := enum.Elems()
	r.Len(values, 3)

	a := Animal{}
	r.Equal("", a.ID())
}

func TestEnumPanic(t *testing.T) {
	type animals_ struct {
		*M_Enum[Animal]
		CAT  Animal
		DOG  Animal
		BIRD *Animal
	}

	r := require.New(t)
	r.PanicsWithValue("parameter's type must not be a pointer", func() {
		NewEnum[Animal](&animals_{})
	})

	r.PanicsWithValue("o_test.animals_.BIRD must not be a pointer type", func() {
		NewEnum[Animal](animals_{
			CAT:  Animal{},
			DOG:  Animal{},
			BIRD: &Animal{},
		})
	})
}
