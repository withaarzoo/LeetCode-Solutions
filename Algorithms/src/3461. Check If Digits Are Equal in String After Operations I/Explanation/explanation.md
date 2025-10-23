# 3461. Check If Digits Are Equal in String After Operations I

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

I am given a string `s` consisting of digits. I repeatedly perform the following operation until the string has exactly two digits:

* For each pair of consecutive digits in `s` (starting from the first digit), I compute a new digit as `(digit_i + digit_{i+1}) % 10`.
* I replace `s` with the sequence of newly computed digits, preserving order.

I must return `true` if the final two digits are the same; otherwise return `false`.

---

## Constraints

* `3 <= s.length <= 100`
* `s` consists of only digits (`'0'` to `'9'`).

---

## Intuition

I thought about what each operation does: it replaces each adjacent pair with their sum modulo 10. If I simulate this reduction step-by-step, the string length decreases by 1 each round until it becomes length 2. Since `n <= 100`, direct simulation is simple and fast — no need for a fancy mathematical trick for this constraint.

---

## Approach

1. Convert the string `s` into an array (or list) of integer digits.
2. While the array length is greater than 2:

   * Build a new array where each element is `(current[i] + current[i+1]) % 10`.
   * Replace the working array with this new array.
3. After the loop, check whether the two remaining digits are equal and return that boolean.

This approach is straightforward, easy to implement, and clearly matches the problem description.

---

## Data Structures Used

* Array / List of integers to store digits.
* Temporary array / list for each reduction step.

---

## Operations & Behavior Summary

* Conversion: string → digits (O(n))
* Iterative reduction: repeatedly compute pairwise sums mod 10 and create a new list of size `k-1` from current length `k`.
* Termination when exactly two digits remain.
* Compare final two digits to return boolean.

---

## Complexity

* **Time Complexity:** `O(n^2)` where `n` is the initial length of `s`.
  Explanation: Each iteration reduces length by 1. The cost per iteration is proportional to current length. Summing costs: `n + (n-1) + ... + 3 ≈ O(n^2)`. With `n ≤ 100`, this is fast.

* **Space Complexity:** `O(n)` extra space for digit lists.
  I allocate a new list each iteration (or replace an array), the maximum extra space is proportional to `n`.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    bool hasSameDigits(string s) {
        // convert string to vector<int> for easy arithmetic
        vector<int> digits;
        digits.reserve(s.size());
        for (char c : s) digits.push_back(c - '0');

        // reduce until exactly two digits remain
        while (digits.size() > 2) {
            vector<int> next;
            next.reserve(digits.size() - 1);
            for (size_t i = 0; i + 1 < digits.size(); ++i) {
                next.push_back((digits[i] + digits[i+1]) % 10);
            }
            digits.swap(next); // efficient replace
        }

        return digits.size() == 2 && digits[0] == digits[1];
    }
};
```

---

### Java

```java
class Solution {
    public boolean hasSameDigits(String s) {
        int n = s.length();
        int[] digits = new int[n];
        for (int i = 0; i < n; ++i) digits[i] = s.charAt(i) - '0';

        while (digits.length > 2) {
            int m = digits.length - 1;
            int[] next = new int[m];
            for (int i = 0; i < m; ++i) {
                next[i] = (digits[i] + digits[i+1]) % 10;
            }
            digits = next;
        }

        return digits.length == 2 && digits[0] == digits[1];
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {boolean}
 */
var hasSameDigits = function(s) {
    // convert to array of numbers
    let digits = new Array(s.length);
    for (let i = 0; i < s.length; ++i) digits[i] = s.charCodeAt(i) - 48;

    // reduce until length 2
    while (digits.length > 2) {
        const next = new Array(digits.length - 1);
        for (let i = 0; i + 1 < digits.length; ++i) {
            next[i] = (digits[i] + digits[i+1]) % 10;
        }
        digits = next;
    }

    return digits.length === 2 && digits[0] === digits[1];
};
```

---

### Python3

```python3
class Solution:
    def hasSameDigits(self, s: str) -> bool:
        # convert to list of ints
        digits = [ord(c) - 48 for c in s]

        # reduce until exactly two digits remain
        while len(digits) > 2:
            next_digits = [(digits[i] + digits[i+1]) % 10 for i in range(len(digits)-1)]
            digits = next_digits

        return len(digits) == 2 and digits[0] == digits[1]
```

---

### Go

```go
package main

// hasSameDigits checks if final two digits after repeated operations are equal
func hasSameDigits(s string) bool {
    // convert to slice of ints
    digits := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        digits[i] = int(s[i] - '0')
    }

    // reduce until length == 2
    for len(digits) > 2 {
        next := make([]int, len(digits)-1)
        for i := 0; i+1 < len(digits); i++ {
            next[i] = (digits[i] + digits[i+1]) % 10
        }
        digits = next
    }

    return len(digits) == 2 && digits[0] == digits[1]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the core idea and map it to the code in small steps. The logic is identical in all implementations — only language syntax differs.

1. **Convert the input string to digits**

   * I transform each character `'0'..'9'` into an integer `0..9`.

     * C++: `digits.push_back(c - '0');`
     * Java: `digits[i] = s.charAt(i) - '0';`
     * JS: `s.charCodeAt(i) - 48`
     * Python: `[ord(c) - 48 for c in s]`
     * Go: `int(s[i] - '0')`
   * This makes mathematical operations `(a + b) % 10` easy and clear.

2. **Iteratively compute the new list**

   * While the current digits array has length > 2:

     * I create `next` of length `current_length - 1`.
     * For each `i` from `0` to `current_length - 2`:

       * Compute `next[i] = (current[i] + current[i+1]) % 10`.
     * Replace `current` with `next`.
   * This mirrors the problem rule: each new digit is the sum modulo 10 of consecutive digits.

3. **Stop when there are exactly two digits**

   * After the loop, I expect the digits list to be of length 2.
   * I return `digits[0] == digits[1]`.

4. **Example walkthrough**

   * Given `s = "3902"`:

     * digits = `[3,9,0,2]`
     * next = `[(3+9)%10=2, (9+0)%10=9, (0+2)%10=2]` → `[2,9,2]`
     * next = `[(2+9)%10=1, (9+2)%10=1]` → `[1,1]`
     * Final two digits equal → `true`.

5. **Why I used new arrays each iteration**

   * Simplicity and clarity. Each iteration reduces length, so allocating a small array is cheap (n ≤ 100).
   * In languages like C++, `swap(next)` avoids an expensive copy.

---

## Examples

1. Input: `"3902"`

   * Output: `true`
   * Explanation: becomes `"[2,9,2]"` → `"[1,1]"` → equal.

2. Input: `"34789"`

   * Output: `false`
   * Explanation: becomes `"[7,1,5,7]"` → `"[8,6,2]"` → `"[4,8]"` → not equal.

---

## How to use / Run locally

* **C++**

  * Place the `Solution` class in a file, write a `main` to read input and call `hasSameDigits`.
  * Compile: `g++ -std=c++17 -O2 solution.cpp -o solution`
  * Run: `./solution`

* **Java**

  * Place class in `Solution.java` and add a `public static void main` to test.
  * Compile: `javac Solution.java`
  * Run: `java Solution`

* **JavaScript (Node)**

  * Save the function in `solution.js`, export or call it with test inputs.
  * Run: `node solution.js`

* **Python3**

  * Save in `solution.py` and call `Solution().hasSameDigits("3902")`.
  * Run: `python3 solution.py`

* **Go**

  * Save function in a `.go` file, add `main()` to test.
  * Build & Run: `go run solution.go`

I kept the examples and test harness out of the README to keep it concise. If you want, I can add test harnesses for each language.

---

## Notes & Optimizations

* For `n ≤ 100`, the `O(n^2)` approach is perfectly fine and simple.
* Micro-optimization: reuse a single buffer and overwrite elements in place to reduce allocations. That is doable but makes the code slightly less readable.
* Advanced math approach: it's possible to compute the final pair directly using binomial coefficients modulo 10 (because each final digit is a linear combination of original digits with binomial weights modulo 10). This is more complex and error-prone; I avoid it unless we must support very large `n`.
* In C++ I used `swap` to avoid copying; in languages with cheap small-array allocation (Python/JS/Go) readable reassignments are okay.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
