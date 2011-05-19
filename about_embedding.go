package about_embedding

import "./koans"

type A interface {
	getA() string
}

type B interface {
	getB() string
}

type AB interface {
	A
	B
}

type Aimpl struct {
}

type Bimpl struct {
}

func (t *Aimpl) getA() string {
	return "Aimpl"
}

func (t *Bimpl) getB() string {
	return "Bimpl"
}

type ABembed struct {
	Aimpl
	Bimpl
}

type ABimpl struct {
}

func (t *ABimpl) getA() string {
	return "A"
}

func (t *ABimpl) getB() string {
	return "B"
}

func TestEmbeddingInterface(t *koans.T) {
	var ab AB = &ABimpl{}
	t.AssertEquals(koans.String__, ab.getA())
	t.AssertEquals(koans.String__, ab.getB())
}

func TestEmbeddingStruct(t *koans.T) {
	ab := &ABembed{}
	t.AssertEquals(koans.String__, ab.getA())
	t.AssertEquals(koans.String__, ab.getB())
}
