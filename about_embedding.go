package main

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

func TestEmbeddingInterface(t *T) {
	var ab AB = &ABimpl{}
	t.AssertEquals(String__, ab.getA())
	t.AssertEquals(String__, ab.getB())
}

func TestEmbeddingStruct(t *T) {
	ab := &ABembed{}
	t.AssertEquals(String__, ab.getA())
	t.AssertEquals(String__, ab.getB())
}
