// https://www.hackerrank.com/challenges/pangrams/problem

package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	var letterCountMap = make(map[rune]int)
	var letterCount = 0

	// Loop through all the characters after converting the string to lowercase.
	for _, chr := range strings.ToLower(s) {
		if chr >= 'a' && chr <= 'z' {
			count, _ := letterCountMap[chr]
			letterCountMap[chr] += 1

			if count == 0 {
				// New character found
				letterCount++

				if letterCount == 26 {
					// Stop reading the string when all characters have been found
					break
				}
			}
		}
	}

	if letterCount == 26 {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}
