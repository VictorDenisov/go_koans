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
