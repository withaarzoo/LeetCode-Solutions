# 166. Fraction to Recurring Decimal

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
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given two integers representing the `numerator` and `denominator` of a fraction, return the fraction in string format. If the fractional part is repeating, enclose the repeating part in parentheses.

Examples:

* `numerator = 1, denominator = 2` → `"0.5"`
* `numerator = 2, denominator = 1` → `"2"`
* `numerator = 4, denominator = 333` → `"0.(012)"`

It is guaranteed that the length of the answer string is less than 10^4 for all given inputs.

## Constraints

* `-2^31 <= numerator, denominator <= 2^31 - 1`
* `denominator != 0`
* Output string length < 10^4

## Intuition

I thought about how I do long division on paper. First I take the integer part by dividing the absolute numerator by the absolute denominator. Then I look at the remainder and simulate the long division: I multiply the remainder by 10, extract the next digit, and update the remainder. If the same remainder appears again during this process, the digits between the two occurrences repeat forever. So I record where each remainder first appeared and when I see it again, I put parentheses around the repeating section.

## Approach

1. Handle the trivial case: if `numerator` is `0`, return `"0"`.
2. Determine the sign of the result. If numerator and denominator have different signs, the result is negative.
3. Convert `numerator` and `denominator` to their absolute values using a type that avoids overflow (e.g., `long long` / `int64` / `BigInt` for JavaScript).
4. Compute the integer part: `integer = abs(numerator) / abs(denominator)` and the initial remainder `rem = abs(numerator) % abs(denominator)`.
5. If `rem == 0` return the integer (with sign) — no fractional part.
6. Otherwise append `.` and simulate long division:

   * Before computing the digit for a remainder, record the current position (index) of the output where the next fractional digit will go.
   * Multiply remainder by 10, compute next digit `rem / denominator`, append digit, set `rem = rem % denominator`.
   * If any remainder repeats, insert `(` at the position recorded for that remainder and append `)` to close the repeating portion. Stop.
7. Return the constructed string.

## Data Structures Used

* **Hash map** (remainder → index in result string): to detect repeating remainders and know where to insert `(`.
* **String-builder / list / string**: to accumulate characters efficiently.
* **64-bit integer (or BigInt in JavaScript)**: to handle overflow (i.e., `abs(INT_MIN)` edge case).

## Operations & Behavior Summary

* Sign detection using XOR of sign bits.
* Absolute conversion to a wide numeric type.
* Integer division for integer part and modulus for remainder.
* Long-division loop generating fractional digits until remainder becomes `0` or repeats.
* Map lookup/insertion per remainder for O(1) expected time per digit.

## Complexity

* **Time Complexity:** `O(k)` where `k` is the number of digits produced in the fractional part until it terminates or repeats. Each remainder produces at most one digit before a repeat; we do constant work per digit.
* **Space Complexity:** `O(k)` for the map that stores seen remainders and the output string (fractional digits).

## Multi-language Solutions

Below are the reference implementations for LeetCode-style function signatures. All implementations follow the same algorithm described earlier and include explanatory comments.

### C++

```c++
#include <string>
#include <unordered_map>
#include <cstdlib>
using namespace std;

class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) return "0";
        string res;
        // Determine sign
        if ((numerator < 0) ^ (denominator < 0)) res.push_back('-');

        // Use long long to avoid overflow when taking abs(INT_MIN)
        long long n = llabs((long long)numerator);
        long long d = llabs((long long)denominator);

        // Integer part
        res += to_string(n / d);
        long long rem = n % d;
        if (rem == 0) return res;

        res.push_back('.');
        unordered_map<long long, int> seen; // remainder -> index in result string

        // Simulate long division
        while (rem != 0) {
            if (seen.find(rem) != seen.end()) {
                int pos = seen[rem];
                res.insert(pos, "(");  // insert '(' at first index where this remainder appeared
                res.push_back(')');      // close parentheses at end
                break;
            }
            seen[rem] = res.size();
            rem *= 10;
            int digit = rem / d;
            res.push_back(char('0' + digit));
            rem = rem % d;
        }
        return res;
    }
};
```

### Java

```java
import java.util.HashMap;
import java.util.Map;

class Solution {
    public String fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0) return "0";

        StringBuilder res = new StringBuilder();
        // sign
        if ((numerator < 0) ^ (denominator < 0)) res.append('-');

        // use long to avoid overflow
        long n = Math.abs((long) numerator);
        long d = Math.abs((long) denominator);

        // integer part
        res.append(n / d);
        long rem = n % d;
        if (rem == 0) return res.toString();

        res.append('.');
        Map<Long, Integer> seen = new HashMap<>(); // remainder -> index in res

        while (rem != 0) {
            if (seen.containsKey(rem)) {
                int pos = seen.get(rem);
                res.insert(pos, "(");
                res.append(')');
                break;
            }
            seen.put(rem, res.length());
            rem *= 10;
            res.append(rem / d);
            rem = rem % d;
        }
        return res.toString();
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} numerator
 * @param {number} denominator
 * @return {string}
 */
var fractionToDecimal = function(numerator, denominator) {
    if (numerator === 0) return "0";

    // Use BigInt to avoid overflow for extreme values
    let n = BigInt(numerator);
    let d = BigInt(denominator);

    let res = "";
    // sign: different signs -> negative
    if ((n < 0n) !== (d < 0n)) res += "-";

    // work with absolute values
    if (n < 0n) n = -n;
    if (d < 0n) d = -d;

    // integer part
    res += (n / d).toString();
    let rem = n % d;
    if (rem === 0n) return res;

    res += ".";
    const seen = new Map(); // remainder (BigInt) -> index in res string

    while (rem !== 0n) {
        if (seen.has(rem)) {
            const pos = seen.get(rem);
            res = res.slice(0, pos) + "(" + res.slice(pos) + ")";
            break;
        }
        seen.set(rem, res.length);
        rem *= 10n;
        res += (rem / d).toString();
        rem = rem % d;
    }
    return res;
};
```

### Python3

```python3
class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        # if numerator is zero, result is 0
        if numerator == 0:
            return "0"

        res = []
        # sign
        if (numerator < 0) ^ (denominator < 0):
            res.append('-')

        n = abs(numerator)
        d = abs(denominator)

        # integer part
        res.append(str(n // d))
        rem = n % d
        if rem == 0:
            return ''.join(res)

        res.append('.')
        seen = {}  # remainder -> index in res list

        # long division - detect repeating remainder
        while rem:
            if rem in seen:
                idx = seen[rem]
                res.insert(idx, '(')
                res.append(')')
                break
            seen[rem] = len(res)
            rem *= 10
            res.append(str(rem // d))
            rem = rem % d

        return ''.join(res)
```

### Go

```go
package main

import (
    "strconv"
    "strings"
)

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    var sb strings.Builder
    // sign
    if (numerator < 0) != (denominator < 0) {
        sb.WriteByte('-')
    }

    // convert to int64 and take absolute values to be safe
    n := int64(numerator)
    d := int64(denominator)
    if n < 0 {
        n = -n
    }
    if d < 0 {
        d = -d
    }

    // integer part
    sb.WriteString(strconv.FormatInt(n/d, 10))
    rem := n % d
    if rem == 0 {
        return sb.String()
    }

    sb.WriteByte('.')
    posMap := make(map[int64]int) // remainder -> index in current sb string

    for rem != 0 {
        if p, ok := posMap[rem]; ok {
            s := sb.String()
            // insert '(' at p and ')' at end
            return s[:p] + "(" + s[p:] + ")"
        }
        posMap[rem] = sb.Len()
        rem *= 10
        sb.WriteString(strconv.FormatInt(rem/d, 10))
        rem = rem % d
    }
    return sb.String()
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I describe the algorithm in detail and then explain key lines in each language.

**High-level algorithm (same across all languages)**

1. If numerator is `0`, return `"0"`.
2. Determine sign of the result by checking if numerator and denominator have different signs.
3. Convert both numbers to their absolute values in a type that prevents overflow (use 64-bit where needed).
4. Compute integer part (`abs(numerator) / abs(denominator)`) and the remainder (`abs(numerator) % abs(denominator)`).
5. If remainder is 0, return integer part (with sign). Else append `.` and start the long-division simulation.
6. Use a hash map `seen` mapping `remainder -> position` (the index in the resulting string where the next digit will be placed). For each step:

   * If current `rem` was seen before: insert `(` at `seen[rem]` and append `)` and stop.
   * Else record `seen[rem] = current index`, `rem *= 10`, append `rem / denom` to result, update `rem = rem % denom`.

**Language-specific notes & key lines**

* **C++**

  * Use `long long` and `llabs` to safely take `abs` of `INT_MIN`.
  * `unordered_map<long long, int> seen` stores remainder positions (index in string `res`).
  * `res.insert(pos, "(")` is used to put the opening parenthesis at the correct index.

* **Java**

  * Use `long` to store absolute values (`Math.abs((long) numerator)`) to avoid overflow.
  * `Map<Long, Integer> seen` stores positions.
  * `StringBuilder res` is used for efficient string concatenation and `res.insert(pos, "(")` inserts parenthesis.

* **JavaScript**

  * Use `BigInt` to safely represent `abs(INT_MIN)` and keep integer division precise.
  * `Map` is used with `BigInt` keys for remainders.
  * Strings are immutable, so when a repeat is detected we build a new string using `slice` to insert parentheses.

* **Python3**

  * Python `int` is arbitrary precision so overflow is not an issue.
  * `seen` is a dict mapping remainder to the index in the `res` list where the next digit will be inserted.
  * Use a list `res` and `''.join(res)` for efficient concatenation.

* **Go**

  * Convert to `int64` for absolute computation safety.
  * `map[int64]int` is used to store remainder to index mapping.
  * `strings.Builder` is used for efficient string building. When repeat is found we reconstruct the string with parentheses insertion.

## Examples

| Input (numerator, denominator) | Output      |
| -----------------------------: | :---------- |
|                           1, 2 | `"0.5"`     |
|                           2, 1 | `"2"`       |
|                         4, 333 | `"0.(012)"` |
|                         -50, 8 | `"-6.25"`   |
|                           1, 6 | `"0.1(6)"`  |

## How to use / Run locally

* **C++ (g++)**

  1. Put the `Solution` class into a `.cpp` file with a `main()` that calls your test cases.
  2. Compile: `g++ -std=c++17 -O2 -o frac fraction.cpp`
  3. Run: `./frac`

* **Java**

  1. Put the `Solution` class in `Solution.java` and add a `main` to test.
  2. Compile: `javac Solution.java`
  3. Run: `java Solution`

* **JavaScript (Node.js)**

  1. Save the JS function in `fraction.js` and write a small harness that calls the function and prints the result.
  2. Run: `node fraction.js`

* **Python3**

  1. Save the class in `solution.py` and add a test harness (`if __name__ == '__main__': ...`).
  2. Run: `python3 solution.py`

* **Go**

  1. Save the function in a Go file (e.g., `main.go`) and add a `main()` to call and print results.
  2. Run: `go run main.go`

## Notes & Optimizations

* Use a type that safely handles `abs(INT_MIN)` — `long long` in C++, `long` in Java, `int64` in Go, or `BigInt` in JS. Python's `int` is safe.
* The map stores remainders; each remainder can only appear once before repeating, so space usage is bounded by the number of digits in the fractional part.
* We avoid building many intermediate strings (use `StringBuilder`/`strings.Builder`/list+join) for performance.
* This algorithm is optimal in the sense that it does the minimal necessary work to generate the fractional representation.

## Author

* Author: [Aarzoo](https://bento.me/withaarzoo)
* Problem: LeetCode `166. Fraction to Recurring Decimal`
* Implementation & writeup: I wrote this solution following long-division with remainder tracking.
