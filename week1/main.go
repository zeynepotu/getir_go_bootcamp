package main

import (
	"fmt"
	"getir_golang/if_else"
	"getir_golang/sLength"
)

const MaxStringLength = 10

func main() {
	str1 := "Hello"
	str2 := "World!"
	str3 := "E.T.."
	fmt.Println("--- Task 1: Variable Declaration and String Operations ---")
	sLength.StringLength(str1)
	sLength.StringLength(str2)
	sLength.StringLength(str3)
	fmt.Println()
	fmt.Println("--- Task 2: If-Else Conditions ---")
	if_else.CompareStrings(len(str1), len(str2), len(str3))
	fmt.Println("--- Task 3:  Constants and IOTA ---")
	fmt.Printf("Max String Length: %d\n\n", MaxStringLength)
	fmt.Println("--- Days of the Week ---")
	printDay()
}
func printDay() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Printf("Monday: %d ", Monday)
	fmt.Println()
	fmt.Printf("Tuesday: %d ", Tuesday)
	fmt.Println()
	fmt.Printf("Wednesday: %d ", Wednesday)
	fmt.Println()
	fmt.Printf("Thursday: %d ", Thursday)
	fmt.Println()
	fmt.Printf("Friday: %d ", Friday)
	fmt.Println()
	fmt.Printf("Saturday: %d ", Saturday)
	fmt.Println()
	fmt.Printf("Sunday: %d ", Sunday)
	fmt.Println()

}
