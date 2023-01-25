package gopher

import (
	"fmt"
	"time"
)

type Hallway struct {
	GopherCrossing *Gopher
}

type Gopher struct {
	Name           string
	IsNeedCrossing bool
}

func (s Hallway) Crossing() {
	fmt.Printf("%s has crossing!\n", s.GopherCrossing.Name)
}

func (g *Gopher) Walk(hallway *Hallway, otherGopher *Gopher) {
	for g.IsNeedCrossing {
		if hallway.GopherCrossing != g {
			time.Sleep(1 * time.Millisecond * 100)
			continue
		}

		/* if otherGopher.IsNeedCrossing {
			fmt.Printf("%s: please, you crossing first %s!\n", g.Name, otherGopher.Name)
			hallway.GopherCrossing = otherGopher
			continue
		} */

		hallway.Crossing()
		g.IsNeedCrossing = false
		fmt.Printf("%s: I have crossed the hallway, thanks %s!\n", g.Name, otherGopher.Name)
		hallway.GopherCrossing = otherGopher
		return
	}
}
