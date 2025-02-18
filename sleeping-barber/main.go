package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variales

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.NewSource(time.Now().UnixNano())

	// print message
	color.Yellow("Sleeping Barber Problem")
	color.Yellow("-----------------------")

	// create channels
	clientChannel := make(chan string, seatingCapacity)
	doneChannel := make(chan bool)

	// create barber
	shop := BarberShop{
		ShopCapacity:       seatingCapacity,
		HairCutDuration:    cutDuration,
		NumberOfBarbers:    0,
		ClientChannel:      clientChannel,
		BarbersDoneChannel: doneChannel,
		Open:               true,
	}

	color.Green("Shop is opened")

	// add barbers
	shop.AddBarber("No.1")
	shop.AddBarber("No.2")
	shop.AddBarber("No.3")
	shop.AddBarber("No.4")
	shop.AddBarber("No.5")
	shop.AddBarber("No.6")

	// start barber as a goroutine

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {

		for {
			randomMiliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMiliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}

	}()

	// block untile barbers closed
	<-closed

}
