package about_maps

import "./koans"

func TestCreatingMap(t *koans.T) {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	t.AssertTrue(koans.Int__ == len(timeZone))
	t.AssertTrue(koans.Int__ == timeZone["UTC"])
}

func TestMapIsReferenceType(t *koans.T) {
	var array [10]int
	t.AssertTrue(koans.Int__ == len(array)) //Length of array is part of its type.
}
