package main

func finalValueAfterOperations(operations []string) int {
    X := 0 // initial value
    for _, op := range operations {
        // if '+' appears it's increment, else decrement
        if containsPlus(op) {
            X++
        } else {
            X--
        }
    }
    return X
}

// small helper: checks if '+' is in the string
func containsPlus(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] == '+' {
            return true
        }
    }
    return false
}
