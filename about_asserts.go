package main

/*
   This unit testing infrastructure doesn't corellate with the gotest infrastructure.
   Please refer to golang.org for documentation about gotest. Here we are using
   a special unit testing infrastructure for go koans.

   This file contains koans for go koans' unit testing infrastructure.
*/
func TestAssertTruth(t *T) {
	t.AssertTrue(false) // This should be true
}

/*  Go doesn't allow method and function overloading. That is why
    AssertTrue with message argument should have another name.
*/
func TestAssertWithMessage(t *T) {
	t.AssertTrueWithMessage(false, "This should be true. Please fix.") // This should be true
}

/*  Function FailNow stops execution of the current test. */
func TestFillInValues(t *T) {
	t.AssertEquals(Int__, 1+1)
}

func TestAssertEquality(t *T) {
	expected_value := Int__
	actual_value := 1 + 1
	t.AssertTrue(expected_value == actual_value)
}

func TestABetterWayOfAssertingEquality(t *T) {
	expected_value := Int__
	actual_value := 1 + 1
	t.AssertEquals(expected_value, actual_value)
}
