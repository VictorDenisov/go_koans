package about_strings

import "./koans"

func TestDoubleQuotedStringsAreStrings(t *koans.T) {

	str := "hello" //String can be defined with literal

	t.AssertTrue("hello" == str)
}

