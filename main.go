package main

import "unicode"

// Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
//    * The numbers can be assumed to be unsigned integers
//    * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

// function `testValidity`  takes the string as an input, and returns boolean flag `true` if the given string complies with the format, or `false` if the string does not comply
// Difficulty: Normal
// Estimated time: 15mins
// Used time: 
func testValidity(s string) bool {
	charType := 0
	isEmpty := true
	if(len(s) == 0) {
		return false
	}
	// Loop through the string
	for _, r := range s {
		//check the charType: 0 - number and 1 - text
		if(charType == 0) {
			if !unicode.IsDigit(r) {
				// r is not a Digit, check if empty number
				if isEmpty {
					return false;
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
				isEmpty = false;
			}
		} else {

		}
    }
	return true
}

func main() {
}
