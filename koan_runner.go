package main

import "./about_asserts"
import "./about_nil"
import "./about_arrays"
import "./about_strings"
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
	{"about_arrays.TestSlicingArrays", about_arrays.TestSlicingArrays},
}

var sampleTests = []koans.Test{
	{"about_strings.TestDoubleQuotedStringsAreStrings", about_strings.TestDoubleQuotedStringsAreStrings},
}

func main() {
	koans.Main(sampleTests)
}
