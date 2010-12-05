package about_strings

import "./koans"
import "fmt"

func TestDoubleQuotedStringsAreStrings(t *koans.T) {
	str := "hello" //String can be defined with literal

	t.AssertTrue(koans.String__ == str)
}

func TestPlusConcatenatesString(t *koans.T) {
	str := "Hello " + "World"
	t.AssertTrue(koans.String__ == str)
}

func TestPlusWillNotModifyOriginalStrings(t *koans.T) {
    hi := "Hello, "
	there := "world"
	str := hi + there
	t.AssertTrue(koans.String__ == hi)
	t.AssertTrue(koans.String__ == there)
	t.AssertTrue(koans.String__ == str)
}

func TestPlusEqualsWillAppendToEndOfString(t *koans.T) {
	hi := "Hello, "
	there := "world"
	hi += there
	t.AssertTrue(koans.String__ == hi)
}

func TestPlusEqualsAlsoLeavesOriginalStringUnmodified(t *koans.T) {
	original := "Hello, "
	hi := original
	there := "world"
	hi += there
	t.AssertTrue(koans.String__ == original)
}

func TestUseSprintfToInterpolateVaribales(t *koans.T) {
	value1 := "one"
	value2 := 2
	str := fmt.Sprintf("The values are %s and %d", value1, value2)
	t.AssertTrue(koans.String__ == str)

}

