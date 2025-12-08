class Solution {
    public int countTriples(int n) {
        int count = 0;

        // Try all possible pairs (a, b)
        for (int a = 1; a <= n; a++) {
            for (int b = 1; b <= n; b++) {
                int sumSquares = a * a + b * b; // this should be c^2

                int c = (int) Math.sqrt(sumSquares); // integer square root

                // Check if c is within range and forms a perfect square
                if (c <= n && c * c == sumSquares) {
                    count++; // (a, b, c) is a valid square triple
                }
            }
        }

        return count;
    }
}
