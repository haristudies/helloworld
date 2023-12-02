package main

type nums struct {
	firstNum, secondNum int
}

func (n nums) addition() int {
	return n.firstNum + n.secondNum
}

func (n nums) mul() int {
	return n.firstNum * n.secondNum
}
