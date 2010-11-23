package main

import "./about_asserts"
import "./koans"

var tests = []koans.Test{
	{"about_asserts.TestAssertTruth", about_asserts.TestAssertTruth},
	{"about_asserts.TestAssertWithMessage", about_asserts.TestAssertWithMessage},
	{"about_asserts.TestFillInValues", about_asserts.TestFillInValues},
	{"about_asserts.TestAssertEquality", about_asserts.TestAssertEquality},
	{"about_asserts.TestABetterWayOfAssertingEquality", about_asserts.TestABetterWayOfAssertingEquality},
}

var sampleTests = []koans.Test{
	{"about_asserts.TestABetterWayOfAssertingEquality", about_asserts.TestABetterWayOfAssertingEquality},
}

func main() {
	koans.Main(tests)
}
