package main

import (
	"fmt"
)

// Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
// If the string is not camelCase, return the string unchanged.
// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase (PascalCase)
// Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).

// ConvertCamelToSnake converts a camelCase string to snake_case.
// If the string is empty, it returns an empty string.
// If the string is not camelCase, it returns the string unchanged.
func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Check if the string is in camelCase
	if !isCamelCase(s) {
		return s
	}

	result := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		// Check if the character is uppercase and not the first character
		if char >= 'A' && char <= 'Z' && i > 0 {
			result += "_" // Add underscore before uppercase letters
		}
		// Convert to lowercase
		if char >= 'A' && char <= 'Z' {
			char += 32 // Convert uppercase to lowercase (ASCII)
		}
		result += string(char)
	}

	return result
}

// isCamelCase checks if a given string is in camelCase format.
func isCamelCase(s string) bool {
	if len(s) == 0 || s[0] >= 'A' && s[0] <= 'Z' || s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return false
		}
		if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= 'A' && s[i] <= 'Z') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}