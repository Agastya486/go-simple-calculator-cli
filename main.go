package main

import (
	"errors"
	"fmt"
	"strconv"
)

func addition(num1, num2 int) int {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	return num1 + num2
}

func subtraction(num1, num2 int) int {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	return num1 - num2
}

func multiplication(num1, num2 int) int {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	return num1 * num2
}

func division(num1, num2 int) (int, error) {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)

	if num1 == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	if num2 == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return num1 / num2, nil
}

func modulo(num1, num2 int) int {
	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	return num1 % num2
}

func main() {
	var choice int
	var num1 int
	var num2 int

	for {
		fmt.Println("\n== Calculator == \n1.Addition \n2.Subtraction \n3.Multiplication \n4.Division \n5.Modulo \n6.End")
		fmt.Print("Choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			result := addition(num1, num2)
			fmt.Println("Result is: " + strconv.Itoa(result))
		case 2:
			result := subtraction(num1, num2)
			fmt.Println("Result is: " + strconv.Itoa(result))
		case 3:
			result := multiplication(num1, num2)
			fmt.Println("Result is: " + strconv.Itoa(result))
		case 4:
			result, err := division(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Result is: " + strconv.Itoa(result))
		case 5:
			result := modulo(num1, num2)
			fmt.Println("Result is: " + strconv.Itoa(result))
		}
		if choice == 6 {
			break
		}
	}

}
