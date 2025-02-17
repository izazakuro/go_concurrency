package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "A", leftFork: 4, rightFork: 0},
	{name: "B", leftFork: 0, rightFork: 1},
	{name: "C", leftFork: 1, rightFork: 2},
	{name: "D", leftFork: 2, rightFork: 3},
	{name: "E", leftFork: 3, rightFork: 4},
}

// define some varibles
var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var order []string
var orderMutex sync.Mutex

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Probelm")
	fmt.Println("---------------------------")
	fmt.Println("Tabel is empty.")

	// start
	dine()
	// print out finished message
	fmt.Println("Table is empty.")
	for i, name := range order {
		fmt.Printf("%d: %s\n", i+1, name)
	}

}

func dine() {
	wg := sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start meal
	for i := 0; i < len(philosophers); i++ {
		// fire off goroutine
		go diningProblem(philosophers[i], &wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("%s eating\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s thinking\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("%s finished meal\n", philosopher.name)

	}

	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")
	orderMutex.Lock()
	order = append(order, philosopher.name)
	orderMutex.Unlock()

}
