package goconcurrencytrial

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Implement thread pool

func DoWorkOnList(inputList []string) string {
	var wg sync.WaitGroup
	wg.Add(len(inputList))

	for _, elem := range inputList {
		go workAsync(&wg, elem)
	}

	// Wait for each goroutine to finish
	wg.Wait()
	fmt.Println("We are done")

	return ""
}

func workAsync(wg *sync.WaitGroup, input string) {
	wg.Done()
	fmt.Println(input)
	time.Sleep(1 * time.Second)
}
