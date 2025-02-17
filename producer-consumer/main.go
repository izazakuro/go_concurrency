package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfProducts = 10

var productsMade, productsFailed, total int

type Producer struct {
	data chan ProductOrder
	quit chan chan error
}

type ProductOrder struct {
	productNumber int
	message       string
	success       bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makeProduct(number int) *ProductOrder {
	number++
	if number <= NumberOfProducts {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order #%d\n", number)

		rnd := rand.Intn(12) + 1
		msg := ""
		isSuccess := false

		if rnd < 5 {
			productsFailed++
		} else {
			productsMade++
		}

		total++

		fmt.Printf("Making product #%d. It will take %d seconds\n", number, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** No Material for #%d *** ", number)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** No Worker for %d ***", number)
		} else {
			isSuccess = true
			msg = fmt.Sprintf("Product #%d is ready", number)
		}

		p := ProductOrder{
			productNumber: number,
			message:       msg,
			success:       isSuccess,
		}

		return &p

	}

	return &ProductOrder{
		productNumber: number,
	}
}

func shop(maker *Producer) {
	// keep track of which product we are making
	var i = 0

	// run forever or until we recieve a quit notification
	// try to make products
	for {

		currentProduct := makeProduct(i)
		if currentProduct != nil {
			i = currentProduct.productNumber
			select {
			case maker.data <- *currentProduct:

			case quitChan := <-maker.quit:

				close(maker.data)
				close(quitChan)

				return
			}
		}

	}
}

func main() {
	// seed the number generator
	rand.NewSource(time.Now().UnixNano())
	// print out a message

	color.Cyan("The Shop is open for business!")
	color.Cyan("------------------------------")

	// create a producer
	producer := &Producer{
		data: make(chan ProductOrder),
		quit: make(chan chan error),
	}
	// run producer in the background

	go shop(producer)

	// create and run consumer

	for i := range producer.data {
		if i.productNumber <= NumberOfProducts {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d completed!", i.productNumber)
			} else {
				color.Red(i.message)
				color.Red("Failed to make!")
			}
		} else {
			color.Cyan("Done making products...")
			err := producer.Close()
			if err != nil {
				color.Red("Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-------------")

	color.Cyan("Done all Job")

	color.Cyan("Made %d products, Failed to make %d, Total: %d", productsMade, productsFailed, total)

	switch {
	case productsFailed > 9:
		color.Red("Awful")
	case productsFailed >= 6:
		color.Blue("Ma-ma-")
	case productsFailed >= 4:
		color.Yellow("Fine")
	case productsFailed >= 2:
		color.Green("Nice")
	default:
		color.Green("Perfect")
	}
}
