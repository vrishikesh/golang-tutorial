/*

Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


Example 1:

Input: columnTitle = "A"
Output: 1
Example 2:

Input: columnTitle = "AB"
Output: 28
Example 3:

Input: columnTitle = "ZY"
Output: 701

*/

package main

import "log"

func main() {
	columnTitle := "ZY"
	log.Println(titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {
	sum, pow := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		digit := int(columnTitle[i] - 'A' + 1)
		sum += digit * pow
		pow *= 26
	}

	return sum
}
