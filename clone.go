package interfaceCopier

import (
	"reflect"
)

/* To clone an interface, you need to know the target structure :-(
 *
 *	a := interface1.(*Struct1)
 *	b := *a
 *	interface2 = &b
 *
 *  This class help for it
 */

type rType struct {
	reflect.Type
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New(v interface{}) rType {
	return rType{reflect.ValueOf(v).Elem().Type()}
}

func (m rType) CloneInterface() interface{} {
	return reflect.New(m.Type).Interface()
}
