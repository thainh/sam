package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Func made by Ba Thai
func clearScreen(){
	cmd := exec.Command("clear") //Linux clear screen
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Init(){
	rand.Seed(time.Now().UnixNano())
	clearScreen()
}

// Identify a number is odd number or even number
func Practice1(){
	fmt.Printf("Practice 1: Identify a number is odd number or even number\n")
	for i := 1; i <= 10; i++ {
		if i == 1 {
			fmt.Printf("\n\nTime  %d:\n", i)
		}else{
			fmt.Printf("\n\nTimes %d:\n", i)
		}
		n := rand.Intn(1000)
		fmt.Print(n)
		if n%2 == 0 {
			fmt.Println(" la so chan")
		} else {
			fmt.Println(" la so le")
		}
	}
}

func main() {
	Init()
	fmt.Printf("I'm learn Golang \n\n")

	//Practice1()

	i := rand.Intn(1000)
	if i%5 == 0 {
		fmt.Printf( "%d can * to 5\n", i)
	} else {
		fmt.Printf("%d is can't * to 5\n", i)
	}
	if i%3 == 0 {
		fmt.Printf( "%d can * to 3\n", i)
	} else {
		fmt.Printf("%d is can't * to 3\n", i)
	}
}
