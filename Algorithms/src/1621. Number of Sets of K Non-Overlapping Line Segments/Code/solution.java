class Solution {
    private static final long MOD = 1_000_000_007L;

    // Computes (base^exp) % MOD using binary exponentiation.
    private long modPow(long base, long exp) {
        long result = 1;

        while (exp > 0) {
            // Multiply by the current base when this bit is set.
            if ((exp & 1L) != 0) {
                result = result * base % MOD;
            }

            // Square the base for the next bit.
            base = base * base % MOD;

            // Move to the next bit.
            exp >>= 1;
        }

        return result;
    }

    public int numberOfSets(int n, int k) {
        // The combinatorial result is C(n + k - 1, 2k).
        long N = n + k - 1L;
        long R = 2L * k;

        // Use the smaller side of the combination to reduce iterations.
        R = Math.min(R, N - R);

        long numerator = 1;
        long denominator = 1;

        // Compute the numerator of C(N, R).
        for (long i = 1; i <= R; i++) {
            numerator = numerator * (N - R + i) % MOD;

            // Compute R!.
            denominator = denominator * i % MOD;
        }

        // Find the modular inverse of R! using Fermat's Little Theorem.
        long inverseDenominator = modPow(denominator, MOD - 2);

        // Divide under modulo by multiplying with the inverse.
        return (int) (numerator * inverseDenominator % MOD);
    }
}