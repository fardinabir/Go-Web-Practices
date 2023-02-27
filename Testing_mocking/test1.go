package main

import "fmt"

type testIntfc interface {
	add() int
}

type Obj struct {
	x, y int
}

type ObjThree struct {
	x, y, z int
}

func (t *Obj) add() int {
	//fmt.Printf("Invoked with x: %v, y:%v\n", t.x, t.y)
	return t.x + t.y
}

func (w *ObjThree) add() int {
	return w.x + w.y + w.z
}

func main() {
	firstSt := Obj{3, 4}
	secondSt := ObjThree{3, 4, 5}
	fmt.Println(firstSt.add())
	fmt.Println(secondSt.add())

}
