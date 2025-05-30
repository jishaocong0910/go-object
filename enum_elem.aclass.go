package o

type i_EnumElem interface {
	m_EnumElem_() *M_EnumElem
}

type M_EnumElem struct {
	id string
}

func (this *M_EnumElem) m_EnumElem_() *M_EnumElem {
	return this
}

// ID 枚举值ID，值为枚举集合中的字段名
func (this *M_EnumElem) ID() string {
	var id string
	if this != nil {
		id = this.id
	}
	return id
}

// Undefined 是否未定义的枚举
func (this *M_EnumElem) Undefined() bool {
	return this == nil
}
