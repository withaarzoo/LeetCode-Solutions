import "math"

func countTriples(n int) int {
    count := 0

    // Try all possible pairs (a, b)
    for a := 1; a <= n; a++ {
        for b := 1; b <= n; b++ {
            sumSquares := a*a + b*b // this should be c^2

            c := int(math.Sqrt(float64(sumSquares))) // integer square root

            // Check if c is within range and forms a perfect square
            if c <= n && c*c == sumSquares {
                count++ // (a, b, c) is a valid square triple
            }
        }
    }

    return count
}
