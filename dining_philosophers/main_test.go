package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		order = []string{}
		dine()
		if len(order) != 5 {
			t.Errorf("incorrect length of slice")
		}
	}

}

func Test_dineWithVaryingDelay(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		order = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay
	}

	dine()
	if len(order) != 5 {
		t.Errorf("incorrect length of slice")
	}
}
