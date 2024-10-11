package main

import (
	"fmt"
	"strings"
)

// hexToDecimal converts a hexadecimal string to its decimal equivalent.
func hexToDecimal(hexString string) (int, error) {
	var decimalValue int
	length := len(hexString)

	for i := 0; i < length; i++ {
		currentChar := hexString[length-1-i]
		var currentValue int

		switch {
		case currentChar >= '0' && currentChar <= '9':
			currentValue = int(currentChar - '0')
		case currentChar >= 'A' && currentChar <= 'F':
			currentValue = int(currentChar-'A') + 10
		case currentChar >= 'a' && currentChar <= 'f':
			currentValue = int(currentChar-'a') + 10
		default:
			return 0, fmt.Errorf("invalid hexadecimal string: %s", hexString)
		}

		decimalValue += currentValue * (1 << (i * 4)) // Equivalent to currentValue * (16^i)
	}

	return decimalValue, nil
}

// replaceHexInSentence replaces all occurrences of hexadecimal numbers followed by "(hex)" in a sentence with their decimal equivalents.
func replaceHexInSentence(sentence string) string {
	const hexSuffix = " (hex)"
	// Find all words ending with the hex suffix
	for {
		// Search for the next occurrence of the pattern " (hex)"
		index := strings.Index(sentence, hexSuffix)
		if index == -1 {
			break
		}

		// Extract the hex number before the suffix
		hexNumber := strings.TrimSpace(sentence[:index])

		// Convert hexNumber to decimal
		decimalNumber, err := hexToDecimal(hexNumber)
		if err != nil {
			fmt.Println(err)
			return sentence // If there's an error, return the original sentence
		}

		// Replace the hex number and suffix with the decimal equivalent
		sentence = strings.Replace(sentence, hexNumber+hexSuffix, fmt.Sprintf("%d", decimalNumber), 1)
	}

	return sentence
}

func main() {
	inputSentence := "1E (hex) files were added"
	outputSentence := replaceHexInSentence(inputSentence)
	fmt.Println(outputSentence) // Output should be: "30 files were added"
}
