package main

func TestIllegalTrianglesReturnErrors(t *T) {
	_, err := Triangle(0, 0, 0)
	t.AssertTrue(err != nil)
	_, err = Triangle(3, 4, -5)
	t.AssertTrue(err != nil)
	_, err = Triangle(1, 1, 3)
	t.AssertTrue(err != nil)
	_, err = Triangle(2, 4, 2)
	t.AssertTrue(err != nil)
}
