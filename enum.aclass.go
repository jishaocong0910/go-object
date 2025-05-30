package o

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type i_Enum[E i_EnumElem] interface {
	m_Enum_() *M_Enum[E]
}

type M_Enum[E i_EnumElem] struct {
	elems []E
	idMap map[string]E
}

func (this *M_Enum[E]) m_Enum_() *M_Enum[E] {
	return this
}

// Elems 获取所有枚举值
func (this *M_Enum[E]) Elems() []E {
	var result []E
	if this != nil {
		result = this.elems
	}
	return result
}

// Undefined 获取一个未定义的枚举值
func (this *M_Enum[E]) Undefined() E {
	var v E
	return v
}

// OfId 根据ID查找枚举值
func (this *M_Enum[E]) OfId(id string) (value E) {
	if this != nil {
		if v, ok := this.idMap[id]; ok {
			value = v
		}
	}
	return
}

// OfIdIgnoreCase 根据ID查找枚举值，不区分大小写
func (this *M_Enum[E]) OfIdIgnoreCase(id string) (value E) {
	if this != nil {
		for _, v := range this.elems {
			if strings.EqualFold(v.m_EnumElem_().id, id) {
				return v
			}
		}
	}
	return
}

// Is 判断是否存在指定枚举值
func (this *M_Enum[E]) Is(source E, targets ...E) bool {
	if this != nil {
		for _, t := range targets {
			if t.m_EnumElem_().ID() == source.m_EnumElem_().ID() {
				return true
			}
		}
	}
	return false
}

// Not 与Is方法相反
func (this *M_Enum[E]) Not(source E, targets ...E) bool {
	return !this.Is(source, targets...)
}

func NewEnum[E i_EnumElem, ES i_Enum[E]](e ES) ES {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() == reflect.Pointer {
		panic("parameter's type must not be a pointer")
	}
	t = reflect.TypeOf(&e).Elem()
	v = reflect.ValueOf(&e).Elem()
	expectedType := reflect.TypeOf((*E)(nil)).Elem()
	v.FieldByName("M_Enum").Set(reflect.ValueOf(&M_Enum[E]{}))

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		actualType := tf.Type
		if actualType.Kind() == reflect.Pointer {
			actualType = actualType.Elem()
		}
		if !actualType.AssignableTo(expectedType) {
			continue
		}
		if vf.Kind() == reflect.Pointer {
			panic(fmt.Sprintf("%s.%s must not be a pointer type", t.String(), tf.Name))
		}

		var elem E
		evField := vf.FieldByName("M_EnumElem")
		if !tf.IsExported() {
			reflect.NewAt(evField.Type(), unsafe.Pointer(evField.UnsafeAddr())).Elem().Set(reflect.ValueOf(&M_EnumElem{}))
			elem = reflect.NewAt(vf.Type(), unsafe.Pointer(vf.UnsafeAddr())).Elem().Interface().(E)
		} else {
			evField.Set(reflect.ValueOf(&M_EnumElem{}))
			elem = vf.Interface().(E)
		}

		mEv := elem.m_EnumElem_()
		mEv.id = tf.Name

		mE := e.m_Enum_()
		mE.elems = append(mE.elems, elem)
	}

	mE := e.m_Enum_()
	mE.idMap = make(map[string]E, len(mE.elems))
	for _, elem := range mE.elems {
		mE.idMap[elem.m_EnumElem_().id] = elem
	}

	return v.Interface().(ES)
}
