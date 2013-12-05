package main

// You need to write the triangle method in the file 'triangle.go'

func TestEquilateralTrianglesHaveEqualSides(t *T) {
	var result string
	result, _ = Triangle(2, 2, 2)
	t.AssertEquals("equilateral", result)
	result, _ = Triangle(10, 10, 10)
	t.AssertEquals("equilateral", result)
}

func TestIsoscelesTrianglesHaveExactlyTwoSidesEqual(t *T) {
	var result string
	result, _ = Triangle(3, 4, 4)
	t.AssertEquals("isosceles", result)
	result, _ = Triangle(4, 3, 4)
	t.AssertEquals("isosceles", result)
	result, _ = Triangle(4, 4, 3)
	t.AssertEquals("isosceles", result)
	result, _ = Triangle(10, 10, 2)
	t.AssertEquals("isosceles", result)
}

func TestScaleneTrianglesHaveNoEqualSides(t *T) {
	var result string
	result, _ = Triangle(3, 4, 5)
	t.AssertEquals("scalene", result)
	result, _ = Triangle(10, 11, 12)
	t.AssertEquals("scalene", result)
	result, _ = Triangle(5, 4, 2)
	t.AssertEquals("scalene", result)
}
