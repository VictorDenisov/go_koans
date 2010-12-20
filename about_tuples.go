package about_tuples

import "./koans"

func sample() (string, bool) {
	return "hello", true
}

func TestTupleAssignment(t *koans.T) {
	result, errcode := sample()

	t.AssertTrue(koans.String__ == result)
	t.AssertTrue(koans.Boolean__ == errcode)
}

func TestTupleAssignmentBlankIdentifier(t *koans.T) {
	result, _ := sample()

	t.AssertTrue(koans.String__ == result)
}

func TestSwapWithTuples(t *koans.T) {
	a := 1
	b := 2
	a, b = b, a
	t.AssertTrue(koans.Int__ == a)
	t.AssertTrue(koans.Int__ == b)

}
