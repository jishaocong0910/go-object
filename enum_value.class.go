package o

type i_EnumValue interface {
	m_092ACD12CAAC() *M_EnumValue
}

type M_EnumValue struct {
	Id string
}

func (this *M_EnumValue) m_092ACD12CAAC() *M_EnumValue {
	return this
}

func (this *M_EnumValue) Undefined() bool {
	return this == nil
}
