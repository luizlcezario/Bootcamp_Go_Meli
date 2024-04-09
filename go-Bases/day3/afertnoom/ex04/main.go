package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

// fiquei na duvida de o ordenacao de grupo seria o merge ou o bubble fiz o bubble
func bubbleSort(arr []int) {
	n := len(arr)
	swapped := true

	for m := 0; swapped; m++ {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				// Swap elements if they are in the wrong order
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--
	}
}

func benchMark(ex func([]int), sl []int) {
	t := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ex func([]int), co []int) {
		ex(co)
		defer wg.Done()

	}(ex, slices.Clone(sl))
	wg.Wait()
	fmt.Println("passou :", time.Since(t))
}

func main() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)
	fmt.Println("Para Insertion com 100 numeros:")
	benchMark(insertionSort, variavel)
	fmt.Println("Para Selection com 100 numeros:")
	benchMark(selectionsort, variavel)
	fmt.Println("Para Bubble com 100 numeros:")
	benchMark(bubbleSort, variavel)

	fmt.Println("Para Insertion com 1000 numeros:")
	benchMark(insertionSort, variavel2)
	fmt.Println("Para Selection com 1000 numeros:")
	benchMark(selectionsort, variavel2)
	fmt.Println("Para Bubble com 1000 numeros:")
	benchMark(bubbleSort, variavel2)

	fmt.Println("Para Insertion com 10000 numeros:")
	benchMark(insertionSort, variavel3)
	fmt.Println("Para Selection com 10000 numeros:")
	benchMark(selectionsort, variavel3)
	fmt.Println("Para Bubble com 10000 numeros:")
	benchMark(bubbleSort, variavel3)
}
