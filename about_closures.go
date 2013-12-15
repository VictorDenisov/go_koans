package main

import "fmt"

func TestClosuresCanBeAssignedToVariablesAndCalledExplicitly(t *T) {
	addOne := func(n int) int { return n + 1 }
	t.AssertEquals(Int__, addOne(10))
}

func MakeOrder(order string) func(qty int) string {
	return func(qty int) string { return fmt.Sprintf("%d", qty) + " " + order + "s" }
}

func TestAccessingClosureViaAssignment(t *T) {
	sausages := MakeOrder("sausage")
	eggs := MakeOrder("egg")

	t.AssertEquals(String__, sausages(3))
	t.AssertEquals(String__, eggs(2))
}

func TestAccessingClosureWithoutAssignment(t *T) {
	t.AssertEquals(String__, MakeOrder("spam")(39823))
}
