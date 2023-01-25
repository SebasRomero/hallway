package main

import (
	"time"

	"github.com/sebasromero/hallway/gopher"
)

func main() {
	gopherPurple := &gopher.Gopher{
		Name:           "Gopher purple",
		IsNeedCrossing: true,
	}
	gopherGreen := &gopher.Gopher{
		Name:           "Gopher green",
		IsNeedCrossing: true,
	}

	hallway := &gopher.Hallway{GopherCrossing: gopherPurple}

	go func() {
		gopherPurple.Walk(hallway, gopherGreen)
	}()

	go func() {
		gopherGreen.Walk(hallway, gopherPurple)
	}()

	time.Sleep(time.Second)
}
