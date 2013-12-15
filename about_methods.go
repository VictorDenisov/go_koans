package main

type Int int

func (t Int) Add(v Int) Int {
	return Int(int(t) + int(v))
}

func AGlobalFunction(a, b int) int {
	return a + b
}

func FunctionWithMultipleReturnTypes() (int, string) {
	return 3, "hello"
}

func FunctionWithNamedReturnValues() (a int, b string) {
	a = 4
	b = "bye"
	return
}

func TestFunctionWithNamedReturnValues(t *T) {
	intValue, stringValue := FunctionWithNamedReturnValues()

	t.AssertEqualInt(Int__, intValue)
	t.AssertTrue(String__ == stringValue)
}

func TestFunctionWithMultipleReturnTypes(t *T) {
	intValue, stringValue := FunctionWithMultipleReturnTypes()

	t.AssertEqualInt(Int__, intValue)
	t.AssertTrue(String__ == stringValue)
}

func TestCallingAGlobalFunction(t *T) {
	result := AGlobalFunction(3, 4)
	t.AssertEqualInt(Int__, int(result))
}

func TestEveryTypeCanHaveMethods(t *T) {
	v := Int(4)
	result := v.Add(3)
	t.AssertEqualInt(Int__, int(result))
}

type Interface interface {
	ReturnValue() int
}

type Implementation struct {
}

// The declaration of this function makes
// only a pointer to Implementation satisfy Interface.
func (t *Implementation) ReturnValue() int {
	return 1
}

func TestInterfaces(t *T) {
	var i Interface
	v := Implementation{}
	i = &v //Only a pointer to Implementation satisfies Interface.
	value := i.ReturnValue()

	t.AssertEquals(2, value)
}
