package main

func TestEmptyPointerIsNil(t *T) {
	var p *int
	t.AssertTrue(Intp__ == p) //Put nil here instead of Intp__
}
