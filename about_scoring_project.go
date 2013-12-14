package main

/*
Greed is a dice game where you roll up to five dice to accumulate
points.  The following "score" function will be used to calculate
the score of a single roll of the dice.

A greed roll is scored as follows:

* A set of three ones is 1000 points

* A set of three numbers (other than ones) is worth 100 times the
  number. (e.g. three fives is 500 points).

* A one (that is not part of a set of three) is worth 100 points.

* A five (that is not part of a set of three) is worth 50 points.

* Everything else is worth 0 points.


Examples:

score([1,1,1,5,1]) => 1150 points
score([2,3,4,6,2]) => 0 points
score([3,4,5,3,3]) => 350 points
score([1,5,1,2,4]) => 250 points

More scoring examples are given in the tests below:

Your goal is to write the score method.
*/

func score(diceArg []int) int {
	//You need to write this method
	return -1
}

func TestScoreOfAnEmptyListIsZero(t *T) {
	t.AssertEquals(0, score([]int{}))
}

func TestScoreOfASingleRollOf_5_Is_50(t *T) {
	t.AssertEquals(50, score([]int{5}))
}

func TestScoreOfASingleRollOf_1_Is_100(t *T) {
	t.AssertEquals(100, score([]int{1}))
}

func TestScoreOfMultiple_1sAnd_5sIsTheSumOfIndividualScores(t *T) {
	t.AssertEquals(300, score([]int{1, 5, 5, 1}))
}

func TestScoreOfSingle_2s_3s_4s_And_6s_AreZero(t *T) {
	t.AssertEquals(0, score([]int{2, 3, 4, 6}))
}

func TestScoreOfTriple_1_Is_1000(t *T) {
	t.AssertEquals(1000, score([]int{1, 1, 1}))
}

func TestScoreOfOtherTriplesIs_100x(t *T) {
	t.AssertEquals(200, score([]int{2, 2, 2}))
	t.AssertEquals(300, score([]int{3, 3, 3}))
	t.AssertEquals(400, score([]int{4, 4, 4}))
	t.AssertEquals(500, score([]int{5, 5, 5}))
	t.AssertEquals(600, score([]int{6, 6, 6}))
}

func TestScoreOfMixedIsSum(t *T) {
	t.AssertEquals(250, score([]int{2, 5, 2, 2, 3}))
	t.AssertEquals(550, score([]int{5, 5, 5, 5}))
}
