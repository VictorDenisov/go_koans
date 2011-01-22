package about_closures;

import "fmt"
import "./koans"

func TestClosuresCanBeAssignedToVariablesAndCalledExplicitly(t *koans.T) {
	addOne := func (n int) int { return n + 1}
	t.AssertEquals(koans.Int__, addOne(10))
}

func MakeOrder(order string) (func (qty int) string) {
	return func (qty int) string { return fmt.Sprintf("%d", qty) + " " + order + "s"}
}

func TestAccessingClosureViaAssignment(t *koans.T) {
	sausages := MakeOrder("sausage")
	eggs := MakeOrder("egg")

	t.AssertEquals(koans.String__, sausages(3))
	t.AssertEquals(koans.String__, eggs(2))
}

func TestAccessingClosureWithoutAssignment(t *koans.T) {
	t.AssertEquals(koans.String__, MakeOrder("spam")(39823))
}
