package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("|| Welcome to GoCalc ||")
	fmt.Println("1-Add; 2-Subtract; 3-Multiply; 4-Divide; 5-End\n")

	var m float64
	var n float64

	for {
		option := check_input()
		if option == 5 {
			break
		} else {
			if option <= 0 || option > 4 {
				fmt.Println("Invalid option.")
			} else {
				for {
					m, n = check_values()
					if m != 0 && n != 0 {
						break
					}
				}
				if option == 1 {
					add(m, n)
				} else if option == 2{
					sub(m, n)
				} else if option == 3 {
					mul(m, n)
				} else if option == 4 {
					div(m, n)
				}
			}
		}
	}
}

func check_input() int{
	var input string
	fmt.Print("\nEnter your option: ")
	//fmt.Scanf("%s\n", &input) // Remember to "%s\n"
	fmt.Scanln(&input)
	number,_ := strconv.Atoi(input) // Alphabet to Integer // error = _
	return number
}

func check_values() (float64, float64){
	var x, y string

	fmt.Print("Enter X:")
	fmt.Scanln(&x)
	fmt.Print("Enter Y:")
	fmt.Scanln(&y)

	a,_ := strconv.ParseFloat(x, 64)
	b,_ := strconv.ParseFloat(y, 64)

	return a, b
}

func add(m float64, n float64){
	fmt.Printf("%.2f + %.2f = %.2f", m, n, m+n)
}

func sub(m float64, n float64){
	fmt.Printf("%.2f - %.2f = %.2f", m, n, m-n)
}

func mul(m float64, n float64){
	fmt.Printf("%.2f x %.2f = %.2f", m, n, m*n)
}

func div(m float64, n float64){
	fmt.Printf("%.2f / %.2f = %.2f", m, n, m/n)
}