package about_structs

import "./koans"
import "reflect"

type Dog struct {
}

func TestInstancesOfStructsCanBeGeneratedWithFigureBrackets(t *koans.T) {
	fido := Dog{}
	t.AssertEquals("Dog", reflect.Typeof(fido).Name())
}
