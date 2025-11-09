func countOperations(num1 int, num2 int) int {
    a, b := num1, num2
    ops := 0
    for a > 0 && b > 0 {
        if a < b {
            a, b = b, a // ensure a >= b
        }
        ops += a / b     // batch subtractions
        a = a % b        // remainder
    }
    return ops
}
