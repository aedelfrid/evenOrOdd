package main

import "fmt"

// in this package i will check to see if an int is even or odd and return the result

func evenOrOdd(arr []int) (evenOddArr []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			evenOddArr = append(evenOddArr, "Even");
		} else {
			evenOddArr = append(evenOddArr, "Odd");
		}
	}

	return evenOddArr
}

func main() {
	arr := []int{0,1,2,3,4,5,6,7,8,9,10};

	fmt.Println(evenOrOdd(arr));
}