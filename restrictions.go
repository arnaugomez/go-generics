package main

func divideByTwo[T any] (a T) T {
	return a / 2 // Incorrecto: a podría ser una string
}

type intOrFloat interface {
	int | float64
}

func divideByTwo2[T intOrFloat] (a T) T {
	return a / 2
}