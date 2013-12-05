package main

import "reflect"

type Dog struct {
}

func TestInstancesOfStructsCanBeGeneratedWithFigureBrackets(t *T) {
	fido := Dog{}
	t.AssertEquals(String__, reflect.TypeOf(fido).Name())
}

type Dog2 struct {
	name string
}

func NewDog2() *Dog2 {
	return &Dog2{"Paul"}
}

func (d *Dog2) setName(name string) {
	d.name = name
}

func TestConstructorIsJustASeparateFunction(t *T) {
	dog := NewDog2()
	t.AssertEquals(String__, dog.name)
}

func TestAnyoneCanAccessField(t *T) {
	dog := &Dog2{"Fido"}
	t.AssertEquals(String__, dog.name)
}
