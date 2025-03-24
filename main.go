package main

import (
	"fmt"
	"os"
	"strconv"
)

func printNumber(number ...int) {
	for _, n := range number {
		fmt.Println(n)
	}
}

func changeNumber(n *int) {
	*n = 100
}

func calculateDegrees(val *int, to *string) error {
	switch *to {
	case "F":
		*val = (*val * 9 / 5) + 32
		return nil
	case "C":
		*val = (*val - 32) * 5 / 9
		return nil
	default:
		return fmt.Errorf("invalid target degree")
	}
}

func main() {
	//  NOTE: value category targetcategory
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguments")
		return
	}
	initialValue := os.Args[1]
	value, err := strconv.Atoi(initialValue)
	if err != nil {
		fmt.Println(err)
		return
	}
	target := os.Args[2]
	calculateDegrees(&value, &target)
	fmt.Println(calculateDegrees(&value, &target))
	printNumber(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	printNumber([]int{1, 2, 3, 4, 5}...)
}
