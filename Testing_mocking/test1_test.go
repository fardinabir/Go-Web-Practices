package main

import "testing"

type testCase struct {
	arg1, arg2, exp int
}

var testCases = []testCase{
	{2, 3, 5},
	{5, 4, 9},
	{9, 10, 19},
	{1, 11, 12},
}

func TestAdd(t *testing.T) {

	for _, test := range testCases {
		demoObj1 := Obj{test.arg1, test.arg2}
		//demoObj2 := ObjThree{5, 4, 3}

		got := demoObj1.add()
		exp := test.exp

		if got != exp {
			t.Errorf("Got %v Expected %v", got, exp)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demoObj1 := Obj{3, 5}

		demoObj1.add()
	}
}
