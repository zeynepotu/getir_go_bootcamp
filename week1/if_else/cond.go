package if_else

import "fmt"

func CompareStrings(str1, str2, str3 int) {
	if str1 < str2 {
		fmt.Println("str1 is shorter than str2.")
	}
	if str2 > str3 {
		fmt.Println("str2 is longer than str3.")
	}
	if str1 == str3 {
		fmt.Println("str3 and str1 are of equal length.")
	}
	fmt.Println()
}
