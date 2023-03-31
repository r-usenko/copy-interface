package main

import (
	"fmt"

	interfaceCopier "github.com/r-usenko/copy-interface"
)

type module struct {
	val string
}
type IModule interface {
	SetString(val string)
}

func (m *module) SetString(val string) {
	m.val = val
}

func main() {
	var interface1, interface2 IModule
	interface1 = &module{"hello"}

	t := interfaceCopier.New(interface1)
	interface2 = t.CloneInterface().(IModule)

	fmt.Println(interface1, interface2)

	interface2.SetString("world")
	fmt.Println(interface1, interface2)

}
