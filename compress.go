package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := "you say yes, I say no you say stop and I say go go go"
	fmt.Println(minimize(input))
}

func minimize(input string) string {
	/*
		1. Identify identifiers (string of letters)
		2. Store identifiers in map with location (index of the first appearance of identifier - not character index but identifier index)
		3. Replace any duplicate identifiers with location reference, using StringBuilder for efficiency
	*/

	identifiers := make(map[string]int)
	identifierCount := 0
	inWord := false
	startIndex := 0
	var b strings.Builder

	for i, r := range input {
		if unicode.IsLetter(r) && i != len(input)-1 {
			if !inWord {
				inWord = true
				startIndex = i
			}
		} else {
			if inWord {
				stopIndex := i
				if i == len(input)-1 {
					stopIndex++
				}
				identifier := string(input[startIndex:stopIndex])
				position, ok := identifiers[identifier]

				if !ok {
					identifiers[identifier] = identifierCount
					b.WriteString(identifier)
				} else {
					b.WriteString("$" + strconv.Itoa(position))
				}
				identifierCount++
			}

			if i != len(input)-1 || !inWord {
				b.WriteRune(r)
			}
			inWord = false
		}
	}

	return b.String()
}
