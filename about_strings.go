package main

import "strings"
import "fmt"

func TestDoubleQuotedStringsAreStrings(t *T) {
	str := "hello" //A string can be defined with a literal

	t.AssertEquals(String__, str)
}

func TestBackQuotedStringsAreStrings(t *T) {
	str := `hello\n
world`

	t.AssertEquals(Int__, len(str))
}

func TestPlusConcatenatesString(t *T) {
	str := "Hello " + "World"
	t.AssertEquals(String__, str)
}

func TestPlusWillNotModifyOriginalStrings(t *T) {
	hi := "Hello, "
	there := "world"
	str := hi + there
	t.AssertEquals(String__, hi)
	t.AssertEquals(String__, there)
	t.AssertEquals(String__, str)
}

func TestPlusEqualsWillAppendToEndOfString(t *T) {
	hi := "Hello, "
	there := "world"
	hi += there
	t.AssertEquals(String__, hi)
}

func TestPlusEqualsAlsoLeavesOriginalStringUnmodified(t *T) {
	original := "Hello, "
	hi := original
	there := "world"
	hi += there
	t.AssertEquals(String__, original)
}

func TestUseSprintfToInterpolateVaribales(t *T) {
	value1 := "one"
	value2 := 2
	str := fmt.Sprintf("The values are %s and %d", value1, value2)
	t.AssertEquals(String__, str)
}

func TestYouCanGetASubstringFromAString(t *T) {
	str := "Bacon, lettuce and tomato"
	t.AssertEquals(String__, str[7:10])
}

func TestYouCanGetASingleCharacterFromAString(t *T) {
	str := "Bacon, lettuce and tomato"
	t.AssertTrue(Char__ == str[1])
}

func TestCharactersAreBytesActually(t *T) {
	t.AssertTrue(Char__ == 'a'+1)
}

func TestStringsCanBeSplit(t *T) {
	str := "Sausage Egg Cheese"
	words := strings.Split(str, " ")
	t.AssertEqualsStringSlices(StringSlice__, words)
}

func TestStringsCanBeSplitWithDifferentPatterns(t *T) {
	str := "the,|;rain;in,spain"
	words := strings.Split(str, ",|;")

	t.AssertEqualsStringSlices(StringSlice__, words)
}
