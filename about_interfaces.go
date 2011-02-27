package about_interfaces

import "./koans"

type Hellower interface {
	Hello() string
}

type HellowerImpl string

func (t HellowerImpl) Hello() string {
	return "Hello " + string(t)
}

func TestAnyStructWithRequiredMethodsSatisfiesInterface(t *koans.T) {
	s := HellowerImpl("world").Hello()
	t.AssertEquals(koans.String__, s)
}
