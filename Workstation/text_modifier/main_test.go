package main

import (
    "testing"
)

func TestModifyText(t *testing.T) {
    input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
    expected := "Simply add 66 and 2 and you will see the result is 68."
    result := modifyText(input)
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}

func TestReplaceHexAndBin(t *testing.T) {
    input := "The value is 1A (hex) and 10 (bin)."
    expected := "The value is 26 and 2."
    result := modifyText(input)
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}

func TestHandleCaseModifications(t *testing.T) {
    input := "This is a test (up) and another (low) and one (cap)."
    expected := "THIS is a test and another and One."
    result := modifyText(input)
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}

func TestAdjustPunctuation(t *testing.T) {
    input := "Hello , world ! How are you ?"
    expected := "Hello, world! How are you?"
    result := modifyText(input)
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}

func TestModifyArticles(t *testing.T) {
    input := "A apple and a orange."
    expected := "An apple and an orange."
    result := modifyText(input)
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}
/* */