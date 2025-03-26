package o

type i_EnumValue interface {
	m_EnumValue_() *M_EnumValue
}

type M_EnumValue struct {
	id string
}

func (this *M_EnumValue) m_EnumValue_() *M_EnumValue {
	return this
}

// ID 枚举值ID，值为枚举集合中的字段名
func (this *M_EnumValue) ID() string {
	var id string
	if this != nil {
		id = this.id
	}
	return id
}

// Undefined 是否未定义的枚举
func (this *M_EnumValue) Undefined() bool {
	return this == nil
}
