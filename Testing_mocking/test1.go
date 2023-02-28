package testingmocking

import "fmt"

type Cal interface {
	Add() int
	Sub(int) int
}

type Obj struct {
	x, y int
}

type ObjThree struct {
	x, y, z int
}

func (t *Obj) Add() int {
	fmt.Printf("Invoked with x: %v, y:%v\n", t.x, t.y)
	return t.x + t.y
}

func (t *Obj) Sub(val int) int {
	return val - 1
}

func (w *ObjThree) Add() int {
	return w.x + w.y + w.z
}

func Add(s Cal) int {
	return s.Add()
}

func main() {
	firstSt := Obj{x: 3, y: 4}
	//secondSt := ObjThree{3, 4, 5}
	fmt.Println(Add(&firstSt))
	fmt.Println(firstSt.Sub(10))
}
