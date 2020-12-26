package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

//Func made by Ba Thai
func clearScreen() {
	cmd := exec.Command("clear") //Linux clear screen
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Init() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()
}

//Identify a number is odd number or even number
func Problem1() {
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

func Problem2() {
	count := 0
	for n := 0; n <= 1000; n++ {
		if n%5 == 0 && n%3 == 0 && n%7 == 0 {
			count++
			fmt.Printf("Lan %d: %d can / to 7,5 and 3\n", count, n)
		}
	}
}
func Problem3() {
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
func Problem4() {
	count := 0
	n := rand.Intn(1000)
	fmt.Printf("cac thuong cua %d la :\n", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
			if count <= 9 {
				fmt.Printf("lan %d : %d\n", count, i)
			} else {
				fmt.Printf("lan %d: %d\n", count, i)
			}
		}
	}
}
func Problem5() {

	for n := 1; n <= 200000; n++ {

		// Kiem tra xem n co phai la so hoan hao khong ?
		if n%10000 == 0 {
			fmt.Printf("Ktra %d: \n", n)
		}
		tong := 0
		for i := 1; i < n; i++ {
			if n%i == 0 {
				tong = tong + i
			}
		}
		if tong == n {
			fmt.Printf("%d la so hoan hao\n", n)
		}
	}
}
func Problem6() {
	//Tim so cap doi
	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {

		}
	}


}

func main() {
	Init()
	fmt.Printf("\nStart the program\n")
	//Problem1()
	//Problem2()
	//Problem3()
	//Problem4()
	//Problem5()
	Problem6()

	fmt.Printf("\nEnd the program\n")
}
