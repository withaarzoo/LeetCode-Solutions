const MOD = 1000000007n;

// Computes (base^exp) % MOD using binary exponentiation.
function modPow(base: bigint, exp: bigint): bigint {
    let result = 1n;

    while (exp > 0n) {
        // Multiply by the current power when this exponent bit is set.
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

function numberOfSets(n: number, k: number): number {
    // The combinatorial formula is C(n + k - 1, 2k).
    const N = BigInt(n + k - 1);
    let R = BigInt(2 * k);

    // Use the smaller side because C(N, R) = C(N, N-R).
    if (R > N - R) {
        R = N - R;
    }

    let numerator = 1n;
    let denominator = 1n;

    // Compute the numerator of C(N, R).
    for (let i = 1n; i <= R; i++) {
        numerator = numerator * (N - R + i) % MOD;

        // Compute R!.
        denominator = denominator * i % MOD;
    }

    // Find the modular inverse of R!.
    const inverseDenominator = modPow(denominator, MOD - 2n);

    // The result is below MOD, so converting it to number is safe.
    return Number(numerator * inverseDenominator % MOD);
}