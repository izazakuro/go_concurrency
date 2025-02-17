package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for balance
	balance := 0
	var mutex sync.Mutex
	// print out starting values

	fmt.Printf("Initial balance: %d\n", balance)

	// define weekly revenue

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Part time job", Amount: 10},
		{Source: "Gifts", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 week print out amount ; keep a runing total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {

				mutex.Lock()
				temp := balance
				temp += income.Amount
				balance = temp
				mutex.Unlock()

				fmt.Printf("On week %d, you earned $%d, from %s\n", week, income.Amount, income.Source)

			}

		}(i, income)
	}

	wg.Wait()
	// pringt out final balance
	fmt.Printf("Final Aount: $%d\n", balance)
}
