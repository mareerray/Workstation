package main

import "fmt"
// Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
// If the string is not camelCase, return the string unchanged.
// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase
// Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).
func CamelToSnakeCase(s string) string{
	// If the string is empty, return an empty string.
	// If the string is not camelCase, return the string unchanged.
	if len(s) == 0 || !CheckCamelCase(s){
		return s
	}

	var result string
	for i := 0; i < len(s); i++{
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z'{
			result += "_" + fmt.Sprintf("%c", s[i]+32) // Convert to lowercase directly
		}else{
			result += fmt.Sprintf("%c", s[i])
		}
		
	}

	return result
}

func CheckCamelCase (s string) bool{
	if len(s) == 0{
		return false
	}
	for i, char := range s{ 
		//Check if the first character is uppercase for UpperCamelCase
		if i == 0 && char >= 'A' && char <= 'Z'{ 
			continue
		}
		if char >= 'A' && char <= 'Z'{
			if i > 0 && (s[i-1]>= 'A' && s[i-1] <= 'Z'){
				return false //not camelcase ex. two capitalized letters followed each other (CamelCAse)
		}
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z'){
			return false //not camelcase ex. found numbers or punctuation(camelCase1)
		}
		if s[len(s)-1]>= 'A' && s[len(s)-1]>= 'Z'{
			return false //not camelcase ex.word ends on a capitalized letter (CamelCasE)
		}
	}
	return true	//it is camelcase
}

func main(){
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}