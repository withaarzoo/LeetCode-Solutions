package main

import (
    "strconv" // Import the strconv package for string conversions
    "strings" // Import the strings package for string manipulation
)

// Define arrays for number words below 20, tens, and thousands
var below_20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var thousands = []string{"", "Thousand", "Million", "Billion"}

// Convert a number to its English words representation
func numberToWords(num int) string {
    if num == 0 {
        return "Zero" // If the number is 0, return "Zero"
    }
    result := "" // Initialize the result string
    i := 0       // Initialize the index for thousands array
    // Iterate over the number in chunks of 1000
    for num > 0 {
        if num % 1000 != 0 {
            // Convert the last three digits to words and prepend to the result
            result = helper(num % 1000) + thousands[i] + " " + result
        }
        num /= 1000 // Remove the last three digits from the number
        i++         // Move to the next thousands place
    }
    return strings.TrimSpace(result) // Remove any leading/trailing spaces and return the result
}

// Helper function to convert a number less than 1000 to words
func helper(num int) string {
    if num == 0 {
        return "" // If the number is 0, return an empty string
    } else if num < 20 {
        // If the number is less than 20, return the corresponding word from below_20 array
        return below_20[num] + " "
    } else if num < 100 {
        // If the number is less than 100, convert the tens place and the units place
        return tens[num / 10] + " " + helper(num % 10)
    } else {
        // If the number is 100 or more, convert the hundreds place and the rest
        return below_20[num / 100] + " Hundred " + helper(num % 100)
    }
}
