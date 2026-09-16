class Solution {
public:
    static const long long MOD = 1000000007LL;

    // Computes (base^exp) % MOD using binary exponentiation.
    long long modPow(long long base, long long exp) {
        long long result = 1;

        while (exp > 0) {
            // If the current bit is set, multiply the current power.
            if (exp & 1LL) {
                result = result * base % MOD;
            }

            // Square the base for the next bit.
            base = base * base % MOD;

            // Move to the next bit of the exponent.
            exp >>= 1LL;
        }

        return result;
    }

    int numberOfSets(int n, int k) {
        // From the combinatorial transformation, the answer is:
        // C(n + k - 1, 2k).
        long long N = n + k - 1;
        long long R = 2LL * k;

        // Use C(N, R) = C(N, N-R) to minimize the work.
        R = min(R, N - R);

        long long numerator = 1;
        long long denominator = 1;

        // Build the numerator: N * (N-1) * ... for R terms.
        for (long long i = 1; i <= R; ++i) {
            numerator = numerator * (N - R + i) % MOD;

            // Build R! for the denominator.
            denominator = denominator * i % MOD;
        }

        // denominator^(MOD-2) is the modular inverse of denominator.
        long long inverseDenominator = modPow(denominator, MOD - 2);

        // Multiply numerator by the modular inverse of the denominator.
        return static_cast<int>(numerator * inverseDenominator % MOD);
    }
};