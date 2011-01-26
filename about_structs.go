package about_structs

import "./koans"
import "reflect"

type Dog struct {
}

func TestInstancesOfStructsCanBeGeneratedWithFigureBrackets(t *koans.T) {
	fido := Dog{}
	t.AssertEquals(koans.String__, reflect.Typeof(fido).Name())
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

func TestConstructorIsJustASeparateFunction(t *koans.T) {
	dog := NewDog2()
	t.AssertEquals(koans.String__, dog.name)
}

func TestAnyoneCanAccessField(t *koans.T) {
	dog := &Dog2{"Fido"}
	t.AssertEquals(koans.String__, dog.name)
}
