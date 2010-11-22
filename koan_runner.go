package main

import "./about_testing_test"
import "./koans"

var tests = []koans.Test{
	{"about_testing_test.TestAssertTruth", about_testing_test.TestAssertTruth},
	{"about_testing_test.TestAssertWithMessage", about_testing_test.TestAssertWithMessage},
	{"about_testing_test.TestFailNow", about_testing_test.TestFailNow},
}

func main() {
	koans.Main(tests)
}
