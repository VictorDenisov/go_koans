package about_asserts_test

import "testing"

func TestFailFunction(t *testing.T) {
	if !(false) { // change this line pass the test
		t.Fail()
	}

}
