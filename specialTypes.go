package main

type customFloat float64

func addTwoFloats[T float64](a T, b T) T {
	return a + b
}

func addTwoFloatsOrSimilar[T ~float64](a T, b T) T {
	return a + b
}

func main() {
	var c1 customFloat = 1
	var c2 customFloat = 16

	var c3 = addTwoFloats(c1, c2)
	var c4 = addTwoFloatsOrSimilar(c1, c2)

}