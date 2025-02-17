package main

import (
	"fmt"
	"sync"
)

func print_func(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}

	wg.Add(len(words))

	for i, x := range words {
		go print_func(fmt.Sprintf("%d: %s", i, x), &wg) // don't copy the wg, thus, pass the pointer of wg
	}

	wg.Wait()

	// worst case to use : time.Sleep(1 * time.Second)

	wg.Add(1)
	print_func("Second one", &wg)

}
