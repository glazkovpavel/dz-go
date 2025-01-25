package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sliceCh := make(chan []int)
	resultCh := make(chan []int)
	go createSlice(sliceCh)
	go power(sliceCh, resultCh)

	result := <-resultCh
	fmt.Println(result)

}

func createSlice(sliceCh chan []int) {
	var slice []int
	for i := 0; i < 10; i++ {
		randomInt := rand.Intn(100)

		slice = append(slice, randomInt)
	}
	sliceCh <- slice
}

func power(sliceCh chan []int, resultCh chan []int) {
	slice := <-sliceCh
	for i, value := range slice {
		slice[i] = value * value
	}
	resultCh <- slice
}
