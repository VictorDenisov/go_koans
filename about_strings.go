package about_strings

import "./koans"
import "strings"
import "fmt"

func TestDoubleQuotedStringsAreStrings(t *koans.T) {
	str := "hello" //String can be defined with literal

	t.AssertEquals(koans.String__, str)
}

func TestBackQuotedStringsAreStrings(t *koans.T) {
	str := `hello\n
world`

	t.AssertEquals(koans.Int__, len(str))
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

func TestYouCanGetASubstringFromAString(t *koans.T) {
	str := "Bacon, lettuce and tomato"
	t.AssertTrue(koans.String__ == str[7:10])
}

func TestYouCanGetASingleCharacterFromAString(t *koans.T) {
	str := "Bacon, lettuce and tomato"
	t.AssertTrue(koans.Char__ == str[1])
}

func TestCharactersAreBytesActually(t *koans.T) {
	t.AssertTrue(koans.Char__ == 'a' + 1)
}

func TestStringsCanBeSplit(t *koans.T) {
	str := "Sausage Egg Cheese"
	words := strings.Split(str, " ", -1)
	t.AssertEqualsStringSlices(koans.StringSlice__, words)
}

func TestStringsCanBeSplitWithDifferentPatterns(t *koans.T) {
	str := "the,|;rain;in,spain"
	words := strings.Split(str, ",|;", -1)


	t.AssertEqualsStringSlices(koans.StringSlice__, words)
}

