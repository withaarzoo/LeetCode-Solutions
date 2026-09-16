# 1621. Number of Sets of K Non-Overlapping Line Segments

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [TypeScript](#typescript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 1621, **Number of Sets of K Non-Overlapping Line Segments**, asks us to count how many different ways we can draw exactly `k` non-overlapping line segments using `n` points placed on a 1-D number line.

The points are:

```text
0, 1, 2, ..., n - 1
```

Every segment must cover at least two points, so its two endpoints must be different.

The segments are allowed to share an endpoint, but they cannot overlap each other.

The final answer can become very large, so we return it modulo:

```text
10^9 + 7
```

For example:

```text
Input:  n = 4, k = 2
Output: 5
```

There are exactly 5 valid ways to choose two non-overlapping segments.

This problem can be solved using a combinatorial transformation instead of building a large dynamic programming table. The final formula is:

```text
C(n + k - 1, 2k)
```

where `C(a, b)` represents the binomial coefficient.

## Constraints

| Constraint             | Value             |
| ---------------------- | ----------------- |
| Number of points `n`   | `2 <= n <= 1000`  |
| Number of segments `k` | `1 <= k <= n - 1` |
| Coordinates            | `0` to `n - 1`    |
| Modulus                | `10^9 + 7`        |

The relatively small value of `n` makes an `O(n)` style solution possible, but the combinatorial solution can be implemented with even less extra space.

## Intuition

I first looked at the endpoints of all the segments.

Suppose I have `k` segments. If I write all their endpoints in order, they must satisfy:

```text
x1 < x2 <= x3 < x4 <= x5 < x6 ...
```

The `<` inside each pair is important because every segment needs to cover at least two points.

The `<=` between two segments is allowed because two segments may share an endpoint.

The problem is that this pattern of strict and non-strict inequalities is not easy to count directly.

So I change the endpoint positions slightly.

For endpoint index `i`, I subtract:

```text
floor(i / 2)
```

This removes the minimum spacing that is forced between endpoints of the same segment.

After this transformation, the complicated condition becomes a simple non-decreasing sequence:

```text
y1 <= y2 <= y3 <= ... <= y2k
```

Now the problem is much easier.

I am choosing `2k` values from `n-k` possible values, with repetition allowed and in non-decreasing order.

That is a standard combinations-with-repetition problem.

So the answer becomes:

```text
C((n - k) + 2k - 1, 2k)
```

which simplifies to:

```text
C(n + k - 1, 2k)
```

For example, when `n = 4` and `k = 2`:

```text
C(4 + 2 - 1, 4)
= C(5, 4)
= 5
```

That matches the expected result.

## Approach

I solve the problem in four main steps.

### 1. Represent every segment by its endpoints

For `k` segments, there are `2k` endpoint positions.

Because segments cannot overlap:

```text
x1 < x2 <= x3 < x4 <= ... < x(2k-1) < x(2k)
```

### 2. Remove the forced gaps

I transform every endpoint using:

```text
yi = xi - floor(i / 2)
```

This handles the strict inequalities inside the segments.

After the transformation, I only need:

```text
y1 <= y2 <= ... <= y2k
```

So the original geometric problem becomes a non-decreasing sequence problem.

### 3. Count the sequences

Each transformed value can be one of `n-k` possible values.

I am choosing `2k` values with repetition allowed.

The combinations-with-repetition formula is:

```text
C(m + r - 1, r)
```

Here:

```text
m = n - k
r = 2k
```

Therefore:

```text
C((n-k) + 2k - 1, 2k)
= C(n + k - 1, 2k)
```

### 4. Calculate the combination modulo `10^9 + 7`

I use:

```text
C(N, R) = N! / (R! * (N-R)!)
```

but I cannot perform normal division under modulo arithmetic.

Since `10^9 + 7` is prime, I use Fermat's Little Theorem:

```text
a^(-1) = a^(MOD - 2) mod MOD
```

So division becomes multiplication by a modular inverse.

I also use:

```text
C(N, R) = C(N, N-R)
```

and choose the smaller value of `R` and `N-R` to reduce the number of iterations.

## Data Structures Used

No special data structure is required.

The solution only uses a few integer variables to store:

* `N = n + k - 1`
* `R = 2k`
* The numerator of the combination
* The denominator
* The modular inverse

I do not use arrays, vectors, maps, sets, or a DP table.

This keeps the extra space usage at `O(1)`.

## Operations & Behavior Summary

The algorithm can be summarized like this:

```text
Start
  |
  v
Read n and k
  |
  v
Convert the geometry problem
into a combinations problem
  |
  v
N = n + k - 1
R = 2k
  |
  v
Use R = min(R, N - R)
  |
  v
Compute:
(N-R+1) * ... * N
        /
       R!
  |
  v
Find modular inverse of R!
using exponent MOD - 2
  |
  v
Multiply numerator by
the modular inverse
  |
  v
Return answer modulo 10^9 + 7
```

The most important step is the transformation from non-overlapping line segments to a non-decreasing sequence. Once that is done, the rest is just efficient binomial coefficient calculation.

## Complexity

| Type  | Complexity       | Explanation                                                                        |
| ----- | ---------------- | ---------------------------------------------------------------------------------- |
| Time  | `O(R + log MOD)` | `R = min(2k, n-k-1)` after using combination symmetry, plus modular exponentiation |
| Space | `O(1)`           | Only a fixed number of integer variables are used                                  |

Since `MOD = 10^9 + 7` is fixed, the modular exponentiation part takes about 30 iterations.

There is no `O(nk)` dynamic programming table and no `O(n)` factorial array.

## Multi-language Solutions

### C++

```cpp
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
```

### Java

```java
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
```

### JavaScript

```javascript
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
```

### TypeScript

```typescript
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
```

### Python3

```python
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
```

### Go

```go
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
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is the same in all six languages. The main differences are syntax and numeric handling.

### C++

The C++ solution stores the modulus as a `long long` constant because multiplication can temporarily produce values around `10^18`.

The helper for modular exponentiation uses binary exponentiation. Instead of multiplying `base` `MOD - 2` times, it repeatedly squares the base and processes the bits of the exponent.

The main function first calculates:

```text
N = n + k - 1
R = 2k
```

Then it replaces `R` with `min(R, N-R)` using the symmetry of binomial coefficients.

The numerator is calculated as:

```text
(N-R+1) * (N-R+2) * ... * N
```

while the denominator is:

```text
R!
```

The denominator is converted into a modular inverse with exponent `MOD - 2`.

Finally:

```text
answer = numerator * inverse(denominator) % MOD
```

### Java

The Java version follows exactly the same mathematical steps.

The important detail is using `long` instead of `int` for multiplication. An `int` cannot safely store products close to `10^18`.

The modular power function uses repeated squaring.

The combination is calculated without constructing full factorial values. This is important because factorials become extremely large even when the final answer is only needed modulo `10^9 + 7`.

The result is converted back to `int` only at the end because the final answer is always less than the modulus.

### JavaScript

JavaScript needs extra care because the normal `Number` type uses floating-point arithmetic and cannot safely represent every integer around `10^18`.

Because the combination calculation contains large intermediate products, the implementation uses `BigInt`.

All modular multiplication is therefore done with values such as:

```text
1000000007n
```

The algorithm itself does not change.

The final result is below `10^9 + 7`, so it can safely be converted back to a normal JavaScript `Number`.

### TypeScript

The TypeScript implementation uses `bigint` for the same reason as JavaScript.

The algorithm still uses:

```text
N = n + k - 1
R = 2k
```

followed by the combination calculation and modular inverse.

The TypeScript type annotations make the input and output types clear, while `bigint` protects multiplication from JavaScript's safe-integer limitations.

### Python3

Python integers can grow automatically, so there is no need for a special large-integer type.

The solution still performs modulo after every multiplication to keep the values small and efficient.

The modular exponentiation function uses binary exponentiation.

Python also has a built-in `pow(a, b, mod)` that can perform modular exponentiation efficiently, but implementing the helper manually makes the underlying idea easier to understand and keeps the algorithm consistent with the other languages.

### Go

The Go solution uses `int64` for all calculations.

This is necessary because multiplication can reach values close to `10^18`.

The modular exponentiation function uses the same binary exponentiation technique as the C++, Java, JavaScript, TypeScript, and Python3 versions.

The final result is converted to `int` because the returned value is smaller than `10^9 + 7`.

## Examples

### Example 1

```text
Input:
n = 4
k = 2

Output:
5
```

The formula gives:

```text
C(n + k - 1, 2k)
= C(5, 4)
= 5
```

So there are 5 different ways to choose two non-overlapping line segments.

The transformed sequences can be viewed as:

```text
0000
0001
0011
0111
1111
```

There are exactly 5 of them.

### Example 2

```text
Input:
n = 3
k = 1

Output:
3
```

There is only one segment, so I simply need to choose two different endpoints from the three points.

That gives:

```text
(0,1)
(0,2)
(1,2)
```

The formula gives:

```text
C(3 + 1 - 1, 2)
= C(3, 2)
= 3
```

### Example 3

```text
Input:
n = 30
k = 7

Output:
796297179
```

The total number of valid sets before taking modulo is:

```text
3796297200
```

The required result is:

```text
3796297200 mod (10^9 + 7)
= 796297179
```

Using the formula:

```text
C(30 + 7 - 1, 14)
= C(36, 14)
```

and calculating it modulo `10^9 + 7` gives the expected answer.

## How to Use / Run Locally

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```bash
./solution
```

For a full standalone program, add your own `main()` function and pass values such as `n = 4` and `k = 2`.

### Java

Save the file as:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For local testing, you can create a `main()` method and call `numberOfSets(n, k)`.

### JavaScript

Save the code as:

```text
solution.js
```

Run it with:

```bash
node solution.js
```

You can add a small test call such as:

```text
numberOfSets(4, 2)
```

to check the result locally.

### TypeScript

Save the file as:

```text
solution.ts
```

Install TypeScript if needed:

```bash
npm install -g typescript
```

Compile it with:

```bash
tsc solution.ts
```

Then run the generated JavaScript:

```bash
node solution.js
```

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

You can create a small test section to call:

```text
Solution().numberOfSets(4, 2)
```

### Go

Save the file as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

For local testing, add a `main()` function and call `numberOfSets(4, 2)`.

## Notes & Optimizations

The biggest optimization is avoiding a two-dimensional dynamic programming table.

A natural DP solution can track the current point and the number of segments already created, but that requires additional memory and more transitions.

The combinatorial transformation is much simpler once the endpoint inequalities are understood.

Another useful optimization is:

```text
C(N, R) = C(N, N-R)
```

I always use the smaller side so the multiplication loop does less work.

The JavaScript and TypeScript implementations should use `BigInt`. Using normal `Number` arithmetic for the intermediate multiplication can lose integer precision.

The solution also avoids explicitly calculating factorials such as `N!`. Those numbers become far too large very quickly. Instead, I calculate only the portion needed for the selected binomial coefficient.

The transformation:

```text
yi = xi - floor(i / 2)
```

is the key observation behind the entire solution. It removes the mandatory gap inside every segment and turns the original constraints into a simple non-decreasing sequence.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
