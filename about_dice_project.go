package about_dice_project

import "./koans"

type DiceSet struct {
	values []int
}

func (t *DiceSet) roll(n int) {
	//Need implementing
	//Tip: rand.Int
}

func TestCanCreateADiceSet(t *koans.T) {
	dice := &DiceSet{}
	t.AssertTrue(dice != nil)
}

func TestRollingTheDiceReturnsASetOfIntegersBetween_1_And_6(t *koans.T) {
	dice :=  &DiceSet{}
	dice.roll(5)
	t.AssertEquals(5, len(dice.values))
	for _, value := range dice.values {
		t.AssertTrue(1 <= value && value <= 6)
	}
}

func TestDiceValuesDoNotChangeUnlessExplicitlyRolled(t *koans.T) {
	dice := &DiceSet{}
	dice.roll(5)
	first_time := dice.values
	second_time := dice.values
	t.AssertEquals(5, len(first_time))
	t.AssertEquals(first_time, second_time)
}
