String and number manipulation in Go (Golang) refers to the processes of modifying and interacting with string and numerical data types. Understanding these concepts is essential for any programming task, as strings often represent text data, while numbers are used for calculations.

## String Manipulation in Go

In Go, strings are immutable, meaning once a string is created, it cannot be changed. However, you can perform various operations to create new strings or manipulate existing ones. Here are some common string manipulation techniques:

### Creating Strings

You can create strings using double quotes or backticks. Double quotes allow for escape sequences, while backticks create raw strings that interpret characters literally.

```go
myString := "Hello, World!" // Using double quotes
rawString := `Hello, "World!"` // Using backticks
```

### Common Operations

1. **Splitting Strings**: You can split a string into a slice of substrings using the `strings.Split` function.

   ```go
   import "strings"

   fruits := "apple,orange,kiwi"
   fruitSlice := strings.Split(fruits, ",")
   fmt.Println(fruitSlice) // Output: [apple orange kiwi]
   ```

2. **Changing Case**: You can convert strings to upper or lower case using `strings.ToUpper` and `strings.ToLower`.

   ```go
   upper := strings.ToUpper("hello")
   fmt.Println(upper) // Output: HELLO
   ```

3. **Trimming Whitespace**: Use `strings.TrimSpace` to remove leading and trailing whitespace.

   ```go
   trimmed := strings.TrimSpace("  Hello  ")
   fmt.Println(trimmed) // Output: Hello
   ```

4. **Finding Substrings**: Check if a substring exists within a string using `strings.Contains`.

   ```go
   exists := strings.Contains("Hello, World!", "World")
   fmt.Println(exists) // Output: true
   ```

## Number Manipulation in Go

In Go, numbers can be manipulated using various functions from the `strconv` package for converting between strings and numbers.

### Converting Numbers to Strings

To convert an integer to a string, use the `strconv.Itoa` function:

```go
import "strconv"

num := 42
numStr := strconv.Itoa(num)
fmt.Println(numStr) // Output: "42"
```

### Converting Strings to Numbers

To convert a string back to an integer, use the `strconv.Atoi` function:

```go
strNum := "123"
num, err := strconv.Atoi(strNum)
if err == nil {
    fmt.Println(num) // Output: 123
}
```

## Summary

String and number manipulation in Go involves creating strings, performing operations like splitting and trimming, and converting between strings and numbers. These operations are crucial for handling user input and processing data effectively in your applications. By mastering these concepts, you'll be well on your way to becoming proficient in Go programming.

Citations:
[1] https://golongwithgolang.com/string-manipulation-in-golang
[2] https://www.kelche.co/blog/go/golang-strings/
[3] https://www.codingexplorations.com/blog/mastering-string-manipulation-go
[4] https://www.meetgor.com/golang-strings/
[5] https://www.geeksforgeeks.org/strings-in-golang/
[6] https://pkg.go.dev/strings
[7] https://www.reddit.com/r/golang/comments/xwzc22/is_string_manipulation_supposed_to_be_this_hard/