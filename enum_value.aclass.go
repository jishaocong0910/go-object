package o

type i_EnumValue interface {
	mEnumValue() *M_EnumValue
}

type M_EnumValue struct {
	id string
}

func (this *M_EnumValue) mEnumValue() *M_EnumValue {
	return this
}

func (this *M_EnumValue) ID() string {
	var id string
	if this != nil {
		id = this.id
	}
	return id
}

func (this *M_EnumValue) Undefined() bool {
	return this == nil
}
