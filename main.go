package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Variable Declaration
	var prize, choice, door = 0, 0, 0
	var iterations = 1000
	var switchwin, staywin = 0, 0
	var switchwins, staywins = false, false
	// Seed the random number generator
	rand.Seed(time.Now().Unix())
	//Start pritnting the table
	fmt.Println("Iteration\tPrize\tChoice\tDoor\tStayWins\tSwitchWins")
	// Start Iteration
	for i := 0; i < iterations; i++ {
		prize = random(0, 3)
		choice = random(0, 3)
		door = random(0, 3)
		// Continue this loop until door does not equal prize
		for {
			if door == prize || door == choice {
				door = random(0, 3)
			} else {
				break
			}
		}
		// Switch Wins is true if choice != prize
		switchwins = choice != prize
		// Stay wins if choice = prize
		staywins = choice == prize
		fmt.Printf("%v\t%v\t%v\t%v\t%t\t%t\n", i, prize, choice, door, staywins, switchwins)
		if staywins {
			staywin++
		} else {
			switchwin++
		}
	}
	fmt.Printf("Switch Win: %v\nStay Win: %v\n", switchwin, staywin)
	var switchpercent = float64(switchwin) / float64(iterations)
	fmt.Printf("Switch Win: %%%.2f", switchpercent*100)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
