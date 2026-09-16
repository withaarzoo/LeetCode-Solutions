class Solution:
    MOD = 10**9 + 7

    def mod_pow(self, base: int, exp: int) -> int:
        # This binary exponentiation computes (base^exp) % MOD
        # in O(log(exp)) time.
        result = 1

        while exp > 0:
            # If the current bit is 1, include this power in the result.
            if exp & 1:
                result = result * base % self.MOD

            # Square the base for the next bit.
            base = base * base % self.MOD

            # Move to the next binary bit.
            exp >>= 1

        return result

    def numberOfSets(self, n: int, k: int) -> int:
        # The problem reduces to C(n + k - 1, 2k).
        N = n + k - 1
        R = 2 * k

        # Use the smaller side because C(N, R) = C(N, N-R).
        R = min(R, N - R)

        numerator = 1
        denominator = 1

        # Compute the numerator of C(N, R).
        for i in range(1, R + 1):
            numerator = numerator * (N - R + i) % self.MOD

            # Compute R!.
            denominator = denominator * i % self.MOD

        # Find the modular inverse of R! using Fermat's Little Theorem.
        inverse_denominator = self.mod_pow(
            denominator,
            self.MOD - 2
        )

        # Multiply by the modular inverse to perform division modulo MOD.
        return numerator * inverse_denominator % self.MOD