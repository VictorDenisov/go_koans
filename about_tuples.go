package main

func sample() (string, bool) {
	return "hello", true
}

func TestTupleAssignment(t *T) {
	result, errcode := sample()

	t.AssertTrue(String__ == result)
	t.AssertTrue(Boolean__ == errcode)
}

func TestTupleAssignmentBlankIdentifier(t *T) {
	result, _ := sample()

	t.AssertTrue(String__ == result)
}

func TestSwapWithTuples(t *T) {
	a := 1
	b := 2
	a, b = b, a
	t.AssertTrue(Int__ == a)
	t.AssertTrue(Int__ == b)

}
