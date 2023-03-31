package interfaceCopier_test

import (
	"bytes"
	"testing"

	"github.com/r-usenko/copy-interface"
)

type inter interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

type dummy struct {
	val []byte
}

func (m *dummy) Marshal() ([]byte, error) {
	return m.val, nil
}

func (m *dummy) Unmarshal(data []byte) error {
	m.val = data
	return nil
}

var (
	input  = []byte("hello")
	output = []byte("null")
)

func test(b *testing.B, interface1, interface2 inter) {
	var err error
	var raw [2][]byte

	if err = interface2.Unmarshal(output); err != nil {
		b.Fatal(err)
	}

	if raw[0], err = interface1.Marshal(); err != nil {
		b.Fatal(err)
	}
	if raw[1], err = interface2.Marshal(); err != nil {
		b.Fatal(err)
	}

	if bytes.Compare(raw[0], raw[1]) == 0 {
		b.Fatalf("the source had changed [%q<-%q]", string(raw[0]), string(raw[1]))
	}
}

func BenchmarkCloneNew(b *testing.B) {
	var interface1, interface2 inter
	interface1 = &dummy{val: input}

	for i := 0; i < b.N; i++ {
		interface2 = interfaceCopier.New(interface1).CloneInterface().(inter)
	}

	test(b, interface1, interface2)
}

func BenchmarkCloneCache(b *testing.B) {
	var interface1, interface2 inter
	interface1 = &dummy{val: input}

	rt := interfaceCopier.New(interface1)
	for i := 0; i < b.N; i++ {
		interface2 = rt.CloneInterface().(inter)
	}

	test(b, interface1, interface2)
}

func BenchmarkStruct(b *testing.B) {
	var interface1, interface2 inter

	interface1 = &dummy{val: input}

	for i := 0; i < b.N; i++ {
		x := interface1.(*dummy)
		y := *x
		interface2 = &y
		_ = interface2
	}

	test(b, interface1, interface2)
}
