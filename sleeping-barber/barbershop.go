package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity       int
	HairCutDuration    time.Duration
	NumberOfBarbers    int
	BarbersDoneChannel chan bool
	ClientChannel      chan string
	Open               bool
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(shop.ClientChannel) == 0 {
				color.Yellow("No costumer, %s go to sleep", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientChannel

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)

			} else {
				// shop is closed
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finishied cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChannel <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop")

	close(shop.ClientChannel)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChannel
	}

	close(shop.BarbersDoneChannel)

	color.Green("--------------------------------------------")

	color.Green("Barbershop is now closed , everyone go home")
}

func (shop *BarberShop) addClient(client string) {
	// print out a message
	color.Green("--- %s arrives! ---", client)

	if shop.Open {
		select {
		case shop.ClientChannel <- client:
			color.Blue("%s takes a seat in waiting room", client)
		default:
			color.Red("Waiting room is full")
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
