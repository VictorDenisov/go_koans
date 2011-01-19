package main

import "./about_asserts"
import "./about_nil"
import "./about_arrays"
import "./about_strings"
import "./about_maps"
import "./about_tuples"
import "./about_methods"
import "./about_control_statements"
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
	{"about_strings.TestBackQuotedStringsAreStrings", about_strings.TestBackQuotedStringsAreStrings},
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
	{"about_tuples.TestTupleAssignment", about_tuples.TestTupleAssignment},
	{"about_tuples.TestTupleAssignmentBlankIdentifier", about_tuples.TestTupleAssignmentBlankIdentifier},
	{"about_tuples.TestSwapWithTuples", about_tuples.TestSwapWithTuples},
	{"about_methods.TestEveryTypeCanHaveMethods", about_methods.TestEveryTypeCanHaveMethods},
	{"about_methods.TestCallingAGlobalFunction", about_methods.TestCallingAGlobalFunction},
	{"about_methods.TestFunctionWithMultipleReturnTypes", about_methods.TestFunctionWithMultipleReturnTypes},
	{"about_methods.TestFunctionWithNamedReturnValues", about_methods.TestFunctionWithNamedReturnValues},
	{"about_control_statements.TestIfThenElseStatements", about_control_statements.TestIfThenElseStatements},
	{"about_control_statements.TestIfThenStatement", about_control_statements.TestIfThenStatement},
	{"about_control_statements.TestWhileStatement", about_control_statements.TestWhileStatement},
	{"about_control_statements.TestBreakStatement", about_control_statements.TestBreakStatement},
	{"about_control_statements.TestContinueStatement", about_control_statements.TestContinueStatement},
	{"about_control_statements.TestForStatement", about_control_statements.TestForStatement},
	{"about_control_statements.TestForEachStatement", about_control_statements.TestForEachStatement},
}

var sampleTests = []koans.Test {
	{"about_control_statements.TestForEachStatement", about_control_statements.TestForEachStatement},
}

func main() {
	koans.Main(tests)
}
