package main

import "unicode"

// Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
//    * The numbers can be assumed to be unsigned integers
//    * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

// function `testValidity`  takes the string as an input, and returns boolean flag `true` if the given string complies with the format, or `false` if the string does not comply
// Difficulty: Normal
// Estimated time: 15mins
// Used time: ~13mins
func testValidity(s string) bool {
	charType := 0
	isEmpty := true
	if len(s) == 0 {
		return false
	}
	// Loop through the string
	for i, r := range s {
		if i == len(s)-1 {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		//check the charType: 0 - number and 1 - text
		if charType == 0 {
			if !unicode.IsDigit(r) {
				// r is not a Digit, check if empty number
				if isEmpty {
					return false
				} else {
					// check the dash
					if r != '-' {
						return false
					} else {
						// move to check the text
						charType = 1
						isEmpty = true
					}
				}
			} else {
				// r is a Digit, continue to check the next char
				isEmpty = false
			}
		} else {
			if !unicode.IsLetter(r) {
				// r is not a letter, check if empty text
				// check empty text
				if isEmpty {
					return false
				} else {
					// check the dash
					if r != '-' {
						return false
					} else {
						// move to check the number
						charType = 0
						isEmpty = true
					}
				}
			} else {
				// r is a letter, continue to check the next char
				isEmpty = false
			}
		}
	}
	return true
}

func main() {
}
