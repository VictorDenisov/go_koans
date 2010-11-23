package main

import "./about_asserts"
import "./about_nil"
import "./about_arrays"
import "./koans"

var tests = []koans.Test{
	{"about_asserts.TestAssertTruth", about_asserts.TestAssertTruth},
	{"about_asserts.TestAssertWithMessage", about_asserts.TestAssertWithMessage},
	{"about_asserts.TestFillInValues", about_asserts.TestFillInValues},
	{"about_asserts.TestAssertEquality", about_asserts.TestAssertEquality},
	{"about_asserts.TestABetterWayOfAssertingEquality", about_asserts.TestABetterWayOfAssertingEquality},
	{"about_nil.TestEmptyPointerIsNil", about_nil.TestEmptyPointerIsNil},
	{"about_arrays.TestCreatingArray", about_arrays.TestCreatingArray},
	{"about_arrays.TestArraysAreValues", about_arrays.TestArraysAreValues},
	{"about_arrays.TestAccessingArrayElements", about_arrays.TestAccessingArrayElements},
}

var sampleTests = []koans.Test{
	{"about_arrays.TestAccessingArrayElements", about_arrays.TestAccessingArrayElements},
}

func main() {
	koans.Main(tests)
}
