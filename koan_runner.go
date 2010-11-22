package main

import "./about_asserts"
import "./koans"

var tests = []koans.Test{
	{"about_asserts.TestAssertTruth", about_asserts.TestAssertTruth},
	{"about_asserts.TestAssertWithMessage", about_asserts.TestAssertWithMessage},
	{"about_asserts.TestFillInValues", about_asserts.TestFillInValues},
}

func main() {
	koans.Main(tests)
}
