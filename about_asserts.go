package about_asserts

import "./koans"

/* 
   This unit testing infrastructure doesn't corellate gotest infrastructure.
   Please refer to golang.org for documentation about gotest. Here we are using
   unit testing infrastrucutre special for go koans.

   This file contains koans for go koans unit testing infrastructure.
   */
func TestAssertTruth(t *koans.T) {
	t.AssertTrue(false) // This should be true
}

/* Go doesn't allow method and function overloading. That is why 
   AssertTrue with message argument should have another name.
   */
func TestAssertWithMessage(t *koans.T) {
	t.AssertTrueWithMessage(false, "This should be true. Please fix.") // This should be true
}

/* Function FailNow stops execution of the current test. */
func TestFailNow(t *koans.T) {
	if !(false) { // change this line to pass the test
		t.FailNow()
	}
}
