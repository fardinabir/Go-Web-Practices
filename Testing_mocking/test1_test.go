package testingmocking

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
	mock_testingmocking "testingmocking/mocks"
	"time"
)

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

	//for _, test := range testCases {
	//	demoObj1 := Obj{test.arg1, test.arg2}
	//	//demoObj2 := ObjThree{5, 4, 3}
	//
	//	got := demoObj1.Add()
	//	exp := test.exp
	//
	//	if got != exp {
	//		t.Errorf("Got %v Expected %v", got, exp)
	//	}
	//}

	// -------------------------------------------------- Mock ----------------------
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockCal := mock_testingmocking.NewMockCal(mockCtl)
	gomock.InOrder(
		mockCal.EXPECT().Add().Return(100).Times(1),
		mockCal.EXPECT().Add().DoAndReturn(func() int {
			time.Sleep(2 * time.Second)
			return 101
		}).AnyTimes(),

		mockCal.EXPECT().Sub(gomock.Any()).DoAndReturn(func(x int) int {
			time.Sleep(2 * time.Second)
			return x * 100
		}).AnyTimes(),
	)
	val := Add(mockCal)
	fmt.Println("This is from mock :", val)

	val = Add(mockCal)
	fmt.Println("This is from mock :", val)

	val = mockCal.Sub(5)
	fmt.Println("This is from mock sub :", val)
	if val != 101 {
		t.Fatal("Error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demoObj1 := Obj{3, 5}

		demoObj1.Add()
	}
}
