package about_nil

import "./koans"

func TestEmptyPointerIsNil(t *koans.T) {
	var p *int
	t.AssertTrue(koans.Intp__ == p) //Put nil here instead of koans.Intp__
}

