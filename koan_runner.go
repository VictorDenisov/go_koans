package main

import "./about_asserts"
import "./about_nil"
import "./about_arrays"
import "./about_strings"
import "./about_maps"
import "./about_tuples"
import "./about_methods"
import "./about_control_statements"
import "./about_triangle_project"
import "./about_triangle_project2"
import "./about_closures"
import "./about_structs"
import "./about_scoring_project"
import "./about_dice_project"
import "./about_channels"
import "./about_interfaces"
import "./about_embedding"
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
	{"about_strings.TestPlusEqualsAlsoLeavesOriginalStringUnmodified",
	about_strings.TestPlusEqualsAlsoLeavesOriginalStringUnmodified},
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
	{"about_triangle_project.TestEquilateralTrianglesHaveEqualSides", about_triangle_project.TestEquilateralTrianglesHaveEqualSides},
	{"about_triangle_project.TestIsoscelesTrianglesHaveExactlyTwoSidesEqual",
	about_triangle_project.TestIsoscelesTrianglesHaveExactlyTwoSidesEqual},
	{"about_triangle_project.TestScaleneTrianglesHaveNoEqualSides", about_triangle_project.TestScaleneTrianglesHaveNoEqualSides},
	{"about_triangle_project2.TestIllegalTrianglesReturnErrors", about_triangle_project2.TestIllegalTrianglesReturnErrors},
	{"about_closures.TestClosuresCanBeAssignedToVariablesAndCalledExplicitly",
	about_closures.TestClosuresCanBeAssignedToVariablesAndCalledExplicitly},
	{"about_closures.TestAccessingClosureViaAssignment",
	about_closures.TestAccessingClosureViaAssignment},
	{"about_closures.TestAccessingClosureWithoutAssignment",
	about_closures.TestAccessingClosureWithoutAssignment},
	{"about_scoring_project.TestScoreOfAnEmptyListIsZero",
	about_scoring_project.TestScoreOfAnEmptyListIsZero},
	{"about_scoring_project.TestScoreOfASingleRollOf_5_Is_50",
	about_scoring_project.TestScoreOfASingleRollOf_5_Is_50},
	{"about_scoring_project.TestScoreOfASingleRollOf_1_Is_100",
	about_scoring_project.TestScoreOfASingleRollOf_1_Is_100},
	{"about_scoring_project.TestScoreOfMultiple_1sAnd_5sIsTheSumOfIndividualScores",
	about_scoring_project.TestScoreOfMultiple_1sAnd_5sIsTheSumOfIndividualScores},
	{"about_scoring_project.TestScoreOfSingle_2s_3s_4s_And_6s_AreZero",
	about_scoring_project.TestScoreOfSingle_2s_3s_4s_And_6s_AreZero},
	{"about_scoring_project.TestScoreOfTriple_1_Is_1000",
	about_scoring_project.TestScoreOfTriple_1_Is_1000},
	{"about_scoring_project.TestScoreOfOtherTriplesIs_100x",
	about_scoring_project.TestScoreOfOtherTriplesIs_100x},
	{"about_scoring_project.TestScoreOfMixedIsSum",
	about_scoring_project.TestScoreOfMixedIsSum},
	{"about_structs.TestInstancesOfStructsCanBeGeneratedWithFigureBrackets",
	about_structs.TestInstancesOfStructsCanBeGeneratedWithFigureBrackets},
	{"about_structs.TestConstructorIsJustASeparateFunction",
	about_structs.TestConstructorIsJustASeparateFunction},
	{"about_structs.TestAnyoneCanAccessField",
	about_structs.TestAnyoneCanAccessField},
	{"about_dice_project.TestCanCreateADiceSet",
	about_dice_project.TestCanCreateADiceSet},
	{"about_dice_project.TestRollingTheDiceReturnsASetOfIntegersBetween_1_And_6",
	about_dice_project.TestRollingTheDiceReturnsASetOfIntegersBetween_1_And_6},
	{"about_dice_project.TestDiceValuesDoNotChangeUnlessExplicitlyRolled",
	about_dice_project.TestDiceValuesDoNotChangeUnlessExplicitlyRolled},
	{"about_dice_project.TestDiceValuesShouldChangeBetweenRolls",
	about_dice_project.TestDiceValuesShouldChangeBetweenRolls},
	{"about_dice_project.TestYouCanRollDifferentNumbersOfDice",
	about_dice_project.TestYouCanRollDifferentNumbersOfDice},
	{"about_channels.TestEveryChannelReceiveValue",
	about_channels.TestEveryChannelReceiveValue},
	{"about_interfaces.TestAnyStructWithRequiredMethodsSatisfiesInterface",
	about_interfaces.TestAnyStructWithRequiredMethodsSatisfiesInterface},
	{"about_embedding.TestEmbeddingInterface",
	about_embedding.TestEmbeddingInterface},
	{"about_embedding.TestEmbeddingStruct",
	about_embedding.TestEmbeddingStruct},
}

var sampleTests = []koans.Test {
	{"about_embedding.TestEmbeddingStruct",
	about_embedding.TestEmbeddingStruct},
}

func main() {
	koans.Main(tests)
}
