package main

import "./about_asserts"
import "./about_nil"
import "./about_arrays"
import "./about_strings"
import "./about_maps"
import "./koans"

var tests = []koans.Test {
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
	{"about_maps.TestCreatingMap", about_maps.TestCreatingMap},
	{"about_strings.TestDoubleQuotedStringsAreStrings", about_strings.TestDoubleQuotedStringsAreStrings},
	{"about_strings.TestPlusConcatenatesString", about_strings.TestPlusConcatenatesString},
	{"about_strings.TestPlusWillNotModifyOriginalStrings", about_strings.TestPlusWillNotModifyOriginalStrings},
	{"about_strings.TestPlusEqualsWillAppendToEndOfString", about_strings.TestPlusEqualsWillAppendToEndOfString},
	{"about_strings.TestPlusEqualsAlsoLeavesOriginalStringUnmodified", about_strings.TestPlusEqualsAlsoLeavesOriginalStringUnmodified},
	{"about_strings.TestUseSprintfToInterpolateVaribales", about_strings.TestUseSprintfToInterpolateVaribales},
	{"about_strings.TestYouCanGetASubstringFromAString", about_strings.TestYouCanGetASubstringFromAString},
	{"about_strings.TestYouCanGetASingleCharacterFromAString", about_strings.TestYouCanGetASingleCharacterFromAString},
	{"about_strings.TestCharactersAreBytesActually", about_strings.TestCharactersAreBytesActually},
	{"about_strings.TestStringsCanBeSplit", about_strings.TestStringsCanBeSplit},
	{"about_strings.TestStringsCanBeSplitWithDifferentPatterns", about_strings.TestStringsCanBeSplitWithDifferentPatterns},
}

var sampleTests = []koans.Test {
	{"about_strings.TestStringsCanBeSplitWithDifferentPatterns", about_strings.TestStringsCanBeSplitWithDifferentPatterns},
}

func main() {
	koans.Main(tests)
}
