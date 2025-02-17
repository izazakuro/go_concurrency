package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_print_func(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go print_func("a", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "a") {
		t.Errorf("No a in result")
	}

}
