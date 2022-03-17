package main


import "fmt"

func printAllNumbers(a []int) {
	for _, s := range a {
		fmt.Println(s)
	}
}

func printAll[T any](a []T) {
	for _, s := range a {
		fmt.Println(s)
	}
}