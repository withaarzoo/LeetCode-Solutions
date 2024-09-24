package main

import (
    "strconv"
)

// longestCommonPrefix function takes two integer arrays (arr1 and arr2) and 
// finds the length of the longest common prefix between any number in arr1 and arr2.
func longestCommonPrefix(arr1 []int, arr2 []int) int {
    // Initialize a map to store all possible prefixes from the numbers in arr1
    // The key is the prefix string, and the value is the count of occurrences of that prefix.
    prefixMap := make(map[string]int)
    
    // Step 1: Build the prefix map for arr1
    // Loop through each number in arr1
    for _, num := range arr1 {
        // Convert the current number to a string so we can iterate through its digits.
        strNum := strconv.Itoa(num)
        
        // Initialize an empty prefix string to gradually build up the prefix character by character.
        prefix := ""
        
        // Loop through each character in the string representation of the number.
        for _, ch := range strNum {
            // Append the current character to the prefix.
            prefix += string(ch)
            
            // Increment the count of this prefix in the map.
            prefixMap[prefix]++
        }
    }
    
    // Initialize a variable to keep track of the maximum length of common prefixes found.
    maxLength := 0
    
    // Step 2: Check for common prefixes in arr2
    // Loop through each number in arr2 to compare its prefixes with the ones in arr1.
    for _, num := range arr2 {
        // Convert the current number to a string to iterate through its digits.
        strNum := strconv.Itoa(num)
        
        // Initialize an empty prefix string to build up the prefix for arr2's number.
        prefix := ""
        
        // Loop through each character in the string representation of the number.
        for _, ch := range strNum {
            // Append the current character to the prefix.
            prefix += string(ch)
            
            // Check if this prefix exists in the map (i.e., it was a prefix from some number in arr1).
            if _, found := prefixMap[prefix]; found {
                // If the current prefix length is greater than the previously found maximum length,
                // update maxLength to store the current prefix length.
                if len(prefix) > maxLength {
                    maxLength = len(prefix)
                }
            }
        }
    }
    
    // Return the maximum length of the common prefix found between arr1 and arr2.
    return maxLength
}
