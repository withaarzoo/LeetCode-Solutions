package main

const MOD int64 = 1000000007

// Computes (base^exp) % MOD using binary exponentiation.
func modPow(base, exp int64) int64 {
	result := int64(1)

	for exp > 0 {
		// Multiply by base when the current exponent bit is 1.
		if exp&1 == 1 {
			result = result * base % MOD
		}

		// Square the base for the next bit.
		base = base * base % MOD

		// Move to the next binary bit.
		exp >>= 1
	}

	return result
}

func numberOfSets(n int, k int) int {
	// The combinatorial formula is C(n + k - 1, 2k).
	N := int64(n + k - 1)
	R := int64(2 * k)

	// Use the smaller side of the combination to reduce iterations.
	if R > N-R {
		R = N - R
	}

	numerator := int64(1)
	denominator := int64(1)

	// Compute the numerator of C(N, R).
	for i := int64(1); i <= R; i++ {
		numerator = numerator * (N - R + i) % MOD

		// Compute R!.
		denominator = denominator * i % MOD
	}

	// Compute the modular inverse of R! using Fermat's Little Theorem.
	inverseDenominator := modPow(denominator, MOD-2)

	// Multiply by the modular inverse to divide under modulo.
	return int(numerator * inverseDenominator % MOD)
}