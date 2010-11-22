package main

import "./about_asserts"
import "./koans"

var tests = []koans.Test{
	{"about_asserts.TestAssertTruth", about_asserts.TestAssertTruth},
	{"about_asserts.TestAssertWithMessage", about_asserts.TestAssertWithMessage},
	{"about_asserts.TestFailNow", about_asserts.TestFailNow},
}

func main() {
	koans.Main(tests)
}
