package about_closures;

import "./koans"

func TestClosuresCanBeAssignedToVariablesAndCalledExplicitly(t *koans.T) {
	addOne := func (n int) int { return n + 1}
	t.AssertEquals(koans.Int__, addOne(10))
}
