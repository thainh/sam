package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Func made by Ba Thai
func clearScreen() {
	cmd := exec.Command("clear") //Linux clear screen
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Init() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()
}

// Identify a number is odd number or even number
func Practice1() {
	fmt.Printf("Practice 1: Identify a number is odd number or even number\n")
	for i := 1; i <= 10; i++ {
		if i == 1 {
			fmt.Printf("\n\nTime  %d:\n", i)
		} else {
			fmt.Printf("\n\nTimes %d:\n", i)
		}
		n := rand.Intn(1000)
		fmt.Print(n)
		if n%2 == 0 {
			fmt.Printf(" la so chan\n")
		} else {
			fmt.Printf(" la so le\n")
		}
	}
}

func Practice2() {
	count := 0
	for n := 0; n <= 1000; n++ {
		if n%5 == 0 && n%3 == 0 && n%7 == 0 {
			count++
			fmt.Printf("Lan %d: %d can / to 7,5 and 3\n", count, n)
		}
	}
}
func Practice3() {
	fmt.Printf("This is fionacci from 1 to 50 \n\n")
	i1 := 0
	i2 := 1
	fmt.Printf("So thu 1: %d \n", i1)
	fmt.Printf("So thu 2: %d \n", i2)
	for count := 3; count <= 50; count++ {
		i3 := i1 + i2
		fmt.Printf("So thu %d: %d \n", count, i3)
		i1 = i2
		i2 = i3
	}
}
func main() {
	Init()

	Practice1()
	Practice2()
	Practice3()
}
