package about_testing_test

import "testing"
import "fmt"

/* Test functions should have the following signature:
   func TestXxx(*testing.T)*/
func TestFailFunction(t *testing.T) {
	__ := 1
	fmt.Println("hello" == __)
	if !(false) { // change this line to pass the test
		t.Fail()
	}
}

/* Function Fail doesn't stop execution of the current test. */
func TestFailFunctionDoesNotStopExceutionOfTest(t *testing.T) {
	if !(false) { // change this line to pass the test
		t.Fail()
	}
	fmt.Println("Fail doesn't stop execution")
}

/* Function FailNow stops execution of the current test. */
func TestFailNow(t *testing.T) {
	if !(false) { // change this line to pass the test
		t.FailNow()
	}
	fmt.Println("FailNow stops execution of current test")
}
