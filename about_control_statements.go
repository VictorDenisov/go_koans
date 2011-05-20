package about_control_statements

import "./koans"
import "fmt"

func TestIfThenElseStatements(t *koans.T) {
	var result string
	if true {
		result = "true value"
	} else {
		result = "false value"
	}

	t.AssertTrue(koans.String__ == result)
}

func TestIfThenStatement(t *koans.T) {
	var result string = "default value"
	if true {
		result = "true value"
	}
	t.AssertTrue(koans.String__ == result)
}

func TestWhileStatement(t *koans.T) {
	i := 1
	result := 1
	for i <= 10 {
		result = result * i
		i += 1
	}
	t.AssertEqualInt(koans.Int__, result)
}

func TestBreakStatement(t *koans.T) {
	i := 1
	result := 1
	for {
		if i > 10 {
			break
		}
		result = result * i
		i += 1
	}
	t.AssertEqualInt(koans.Int__, result)
}

func TestContinueStatement(t *koans.T) {
	i := 0
	result := make([]string, 0)
	for i < 10 {
		i += 1
		if i % 2 == 0 {
			continue
		}
		result = append(result, fmt.Sprintf("%d", i))
	}
	t.AssertEqualsStringSlices(koans.StringSlice__, result)
}

func TestForStatement(t *koans.T) {
	result := 1
	for i := 1; i <= 10; i++ {
		result = result * i
	}
	t.AssertEqualInt(koans.Int__, result)
}

func TestForEachStatement(t *koans.T) {
	list := []int{1, 2, 3, 4}
	result := 0
	for _, value := range list {
		result += value
	}
	t.AssertEqualInt(koans.Int__, result)
}

