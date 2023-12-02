package main

type Arithmatic interface {
	add() float32
	multiplication() float32
}

type integerStruct struct {
	a, b int
}

type floatStruct struct {
	a, b float32
}

func (i integerStruct) add() float32 {
	return float32(i.a + i.b)
}

func (i integerStruct) multiplication() float32 {
	return float32(i.a * i.b)
}

func (f floatStruct) add() float32 {
	return f.a + f.b
}

func (f floatStruct) multiplication() float32 {
	return f.a * f.b
}
