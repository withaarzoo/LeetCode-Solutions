/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */

// Use BigInt because multiplying two values near 1e9
// can exceed JavaScript's safe integer range for Number.
const MOD = 1000000007n;

// Computes (base^exp) % MOD using binary exponentiation.
function modPow(base, exp) {
    let result = 1n;

    while (exp > 0n) {
        // Multiply by base when the current exponent bit is 1.
        if (exp & 1n) {
            result = result * base % MOD;
        }

        // Square the base for the next bit.
        base = base * base % MOD;

        // Move to the next binary bit.
        exp >>= 1n;
    }

    return result;
}

var numberOfSets = function(n, k) {
    // The answer is C(n + k - 1, 2k).
    const N = BigInt(n + k - 1);
    let R = BigInt(2 * k);

    // Use C(N, R) = C(N, N-R) to minimize the loop size.
    if (R > N - R) {
        R = N - R;
    }

    let numerator = 1n;
    let denominator = 1n;

    // Compute the numerator of the combination.
    for (let i = 1n; i <= R; i++) {
        numerator = numerator * (N - R + i) % MOD;

        // Compute R!.
        denominator = denominator * i % MOD;
    }

    // Compute the modular inverse of R!.
    const inverseDenominator = modPow(denominator, MOD - 2n);

    // BigInt is converted back to Number because the final answer
    // is always smaller than MOD and therefore safely representable.
    return Number(numerator * inverseDenominator % MOD);
};