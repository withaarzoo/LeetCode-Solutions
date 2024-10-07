package main

func minLength(s string) int {
    stack := []rune{} // Stack to store characters
    
    // Traverse through each character in the string
    for _, ch := range s {
        // If top of the stack forms "AB" or "CD" with the current character, pop the stack
        if len(stack) > 0 && ((stack[len(stack)-1] == 'A' && ch == 'B') || 
                              (stack[len(stack)-1] == 'C' && ch == 'D')) {
            stack = stack[:len(stack)-1]  // Remove the substring
        } else {
            stack = append(stack, ch)  // Push current character onto the stack if no pair
        }
    }
    
    return len(stack)  // The size of the stack is the minimum length of the string
}