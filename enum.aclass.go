package o

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type i_Enum[V i_EnumValue] interface {
	m_Enum_() *M_Enum[V]
}

type M_Enum[V i_EnumValue] struct {
	values []V
	idMap  map[string]V
}

func (this *M_Enum[V]) m_Enum_() *M_Enum[V] {
	return this
}

func (this *M_Enum[V]) Values() []V {
	var result []V
	if this != nil {
		result = this.values
	}
	return result
}

func (this *M_Enum[V]) OfId(id string) (value V) {
	if this != nil {
		if v, ok := this.idMap[id]; ok {
			value = v
		}
	}
	return
}

func (this *M_Enum[V]) OfIdIgnoreCase(id string) (value V) {
	if this != nil {
		for _, v := range this.values {
			if strings.EqualFold(v.m_EnumValue_().id, id) {
				return v
			}
		}
	}
	return
}

func (this *M_Enum[V]) Is(source V, targets ...V) bool {
	if this != nil {
		for _, t := range targets {
			if source.m_EnumValue_().Undefined() {
				if t.m_EnumValue_().Undefined() {
					return true
				}
			} else if !t.m_EnumValue_().Undefined() {
				if t.m_EnumValue_().id == source.m_EnumValue_().id {
					return true
				}
			}
		}
	}
	return false
}

func (this *M_Enum[V]) Not(source V, targets ...V) bool {
	return !this.Is(source, targets...)
}

func NewEnum[V i_EnumValue, E i_Enum[V]](e E) E {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() == reflect.Pointer {
		panic("parameter's type must not be a pointer")
	}
	t = reflect.TypeOf(&e).Elem()
	v = reflect.ValueOf(&e).Elem()
	expectedType := reflect.TypeOf((*V)(nil)).Elem()
	v.FieldByName("M_Enum").Set(reflect.ValueOf(&M_Enum[V]{}))

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

		var value V
		evField := vf.FieldByName("M_EnumValue")
		if !tf.IsExported() {
			reflect.NewAt(evField.Type(), unsafe.Pointer(evField.UnsafeAddr())).Elem().Set(reflect.ValueOf(&M_EnumValue{}))
			value = reflect.NewAt(vf.Type(), unsafe.Pointer(vf.UnsafeAddr())).Elem().Interface().(V)
		} else {
			evField.Set(reflect.ValueOf(&M_EnumValue{}))
			value = vf.Interface().(V)
		}

		mEv := value.m_EnumValue_()
		mEv.id = tf.Name

		mE := e.m_Enum_()
		mE.values = append(mE.values, value)
	}

	mE := e.m_Enum_()
	mE.idMap = make(map[string]V, len(mE.values))
	for _, value := range mE.values {
		mE.idMap[value.m_EnumValue_().id] = value
	}

	return v.Interface().(E)
}
