package main

import "fmt"

func main() {
	slice := make([]int, 4)
	slice[0] = 56
	slice[1] = 56
	newFunc(slice)
	fmt.Println(&slice)
}

func newFunc(nvar []int) {
	nvar[0] = 66
	nvar[1] = 76
	fmt.Println(&nvar)
}
