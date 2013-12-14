package main

import "runtime"
import "os"

var (
	Int__         int    = -10000
	Intp__        *int   = new(int)
	String__      string = ""
	Boolean__     bool   = false
	Char__        byte   = '0'
	StringSlice__ []string
)

type T struct {
	errors string
	failed bool
	ch     chan *T
}

func (t *T) Fail() { t.failed = true }

func (t *T) FailNow() {
	t.Fail()
	t.ch <- t
	runtime.Goexit()
}

func (t *T) AssertTrue(value bool) {
	if value == false {
		t.FailNow()
	}
}

func (t *T) AssertEquals(a, b interface{}) {
	switch v := a.(type) {
	case []int:
		n := len(v)
		w := b.([]int)
		m := len(w)
		t.AssertTrue(n == m)
		for i := 0; i < n; i++ {
			t.AssertTrue(v[i] == w[i])
		}
	case int:
		t.AssertTrue(v == b.(int))
	case string:
		t.AssertTrue(v == b.(string))
	}
}

func (t *T) AssertNotEqual(a, b interface{}) {
	switch v := a.(type) {
	case []int:
		equal := true
		n := len(v)
		w := b.([]int)
		m := len(w)
		equal = equal && (n == m)
		for i := 0; i < n; i++ {
			equal = equal && (v[i] == w[i])
		}
		t.AssertTrue(!equal)
	case int:
		t.AssertTrue(v != b.(int))
	case string:
		t.AssertTrue(v != b.(string))
	}
}

func (t *T) AssertEqualInt(expected int, received int) {
	if expected != received {
		t.FailNow()
	}
}

func (t *T) AssertTrueWithMessage(value bool, message string) {
	if value == false {
		t.FailNow()
	}
}

func (t *T) AssertEqualsStringSlices(expected []string, received []string) {
	n := len(expected)
	if n != len(received) {
		t.FailNow()
	}
	for i := 0; i < n; i++ {
		if expected[i] != received[i] {
			t.FailNow()
		}
	}
}

func (t *T) Failed() bool { return t.failed }

type Test struct {
	Name string
	F    func(*T)
}

func tRunner(t *T, test *Test) {
	test.F(t)
	t.ch <- t
}

func Main(tests []Test) {
	ok := true
	for i := 0; i < len(tests); i++ {
		t := new(T)
		t.ch = make(chan *T)
		go tRunner(t, &tests[i])
		<-t.ch
		if t.failed {
			print("\033[1;31m")
			println(tests[i].Name, " has damaged your karma.")
			println("Please keep meditating")
			print("\033[0m")
			println(t.errors)
			ok = false
			break
		} else {
			print("\033[1;32m")
			println(tests[i].Name, " has expanded your awareness.")
			print("\033[0m")
			println(t.errors)
		}
	}
	if !ok {
		print("\033[1;33m")
		println("You have not yet reached enlightnment ...")
		print("\033[0m")
		os.Exit(1)
	}
	print("\033[1;34m")
	println("Satori!!!")
	print("\033[0m")
}
