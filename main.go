package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, fmt.Errorf("incorrect slice size")
	}

	src := rand.NewSource(time.Now().Unix())
	slice := make([]int, size)
	for i := range slice {
		slice[i] = int(src.Int63()) + 1
	}
	return slice, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("empty slice passed")
	}
	if len(data) == 1 {
		if data[0] < 0 {
			return 0, fmt.Errorf("slice contains negative numbers")
		}
		return data[0], nil
	}
	for _, val := range data {
		if val < 0 {
			return 0, fmt.Errorf("slice contains negative numbers")
		}
	}

	max := 0
	for _, val := range data {
		if val > max {
			max = val
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("empty slice passed")
	}
	if len(data) == 1 {
		if data[0] < 0 {
			return 0, fmt.Errorf("slice contains negative numbers")
		}
		return data[0], nil
	}

	sliceOfMax := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		lenSlice := len(data) / CHUNKS

		startIndex := i * lenSlice
		endIndex := startIndex + lenSlice
		tempSlice := data[startIndex:endIndex]

		if i == CHUNKS-1 {
			tempSlice = data[startIndex:]
		}

		wg.Add(1)
		go func(i int, tempSlice []int) {
			defer wg.Done()
			max, err := maximum(tempSlice)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			sliceOfMax[i] = max
		}(i, tempSlice)
	}
	wg.Wait()
	return maximum(sliceOfMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice, err := generateRandomElements(SIZE)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, err := maximum(slice)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max, err = maxChunks(slice)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
