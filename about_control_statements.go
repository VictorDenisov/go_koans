package main

import "fmt"

func TestIfThenElseStatements(t *T) {
	var result string
	if true {
		result = "true value"
	} else {
		result = "false value"
	}

	t.AssertTrue(String__ == result)
}

func TestIfThenStatement(t *T) {
	var result string = "default value"
	if true {
		result = "true value"
	}
	t.AssertTrue(String__ == result)
}

func TestWhileStatement(t *T) {
	i := 1
	result := 1
	for i <= 10 {
		result = result * i
		i += 1
	}
	t.AssertEqualInt(Int__, result)
}

func TestBreakStatement(t *T) {
	i := 1
	result := 1
	for {
		if i > 10 {
			break
		}
		result = result * i
		i += 1
	}
	t.AssertEqualInt(Int__, result)
}

func TestContinueStatement(t *T) {
	i := 0
	result := make([]string, 0)
	for i < 10 {
		i += 1
		if i%2 == 0 {
			continue
		}
		result = append(result, fmt.Sprintf("%d", i))
	}
	t.AssertEqualsStringSlices(StringSlice__, result)
}

func TestForStatement(t *T) {
	result := 1
	for i := 1; i <= 10; i++ {
		result = result * i
	}
	t.AssertEqualInt(Int__, result)
}

func TestForEachStatement(t *T) {
	list := []int{1, 2, 3, 4}
	result := 0
	for _, value := range list {
		result += value
	}
	t.AssertEqualInt(Int__, result)
}
