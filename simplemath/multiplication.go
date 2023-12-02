package main

type nums struct {
	firstNum  int
	secondNum int
}

func newNums() *nums {
	return &nums{}
}

func (n nums) addition(a int, b int) int {
	return a + b
}

func (n nums) mul(a int, b int) int {
	return a * b
}
