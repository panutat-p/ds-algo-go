package radix_sort

import (
	"fmt"
	"math"
)

/*
https://www.programiz.com/dsa/radix-sort
*/

// RadixSort
// use stable counting sort, not inplace
// all values must have same radix and same width
// radix of decimal is 10, radix of binary is 2
func RadixSort(sl []int, radix int, width int) []int {
	rslt := make([]int, len(sl))

	for i := 0; i < width; i += 1 {
		// count the digit to counts slice
		counts := make([]int, radix) // possible value: 0 to radix-1
		for _, v := range sl {
			digit := getDigit(v, i, radix)
			counts[digit] += 1
		}

		// change count value to cumulative value
		for j := 1; j < len(counts); j += 1 {
			counts[j] += counts[j-1] // ⚠️ counts slice becomes cumulative slice
		}

		for j := len(sl) - 1; j >= 0; j -= 1 { // loop backward
			index := getDigit(sl[j], i, radix) // 🦊 sort the current digit
			rslt[counts[index]-1] = sl[j]      // 🦄 formula
			counts[index] -= 1
			fmt.Println("counts:", counts)
			fmt.Println("rslt:", rslt)
			fmt.Println()
		}
		// reset counts slice and do this again on greater position
	}
	return rslt
}

// 4283/100 = 42, 42 % 10 = 2
func getDigit(num int, position int, radix int) int {
	return (num / int(math.Pow(10, float64(position)))) % radix
}

// call radixSort with radix=26, because ASCII lower case contains 26 characters
// "a" >> 0
// "b" >> 1
func getChar(text string, position int) int {
	return int(text[position] - "a"[0])
}
