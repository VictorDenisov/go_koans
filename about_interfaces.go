package main

type Hellower interface {
	Hello() string
}

type HellowerImpl string

func (t HellowerImpl) Hello() string {
	return "Hello " + string(t)
}

func TestAnyStructWithRequiredMethodsSatisfiesInterface(t *T) {
	s := HellowerImpl("world").Hello()
	t.AssertEquals(String__, s)
}
