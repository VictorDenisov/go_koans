package main

type DiceSet struct {
	values []int
}

func (t *DiceSet) roll(n int) {
	//Needs implementing
	//Tip: rand.Int
}

func TestCanCreateADiceSet(t *T) {
	dice := &DiceSet{}
	t.AssertTrue(dice != nil)
}

func TestRollingTheDiceReturnsASetOfIntegersBetween_1_And_6(t *T) {
	dice := &DiceSet{}
	dice.roll(5)
	t.AssertEquals(5, len(dice.values))
	for _, value := range dice.values {
		t.AssertTrue(1 <= value && value <= 6)
	}
}

func TestDiceValuesDoNotChangeUnlessExplicitlyRolled(t *T) {
	dice := &DiceSet{}
	dice.roll(5)
	first_time := dice.values
	second_time := dice.values
	t.AssertEquals(5, len(first_time))
	t.AssertEquals(first_time, second_time)
}

func TestDiceValuesShouldChangeBetweenRolls(t *T) {
	dice := &DiceSet{}

	dice.roll(5)
	first_time := dice.values

	dice.roll(5)
	second_time := dice.values

	t.AssertNotEqual(first_time, second_time)
}

func TestYouCanRollDifferentNumbersOfDice(t *T) {
	dice := &DiceSet{}

	dice.roll(3)
	t.AssertEquals(3, len(dice.values))

	dice.roll(1)
	t.AssertEquals(1, len(dice.values))
}
