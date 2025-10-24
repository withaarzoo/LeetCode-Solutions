# 2048. Next Greater Numerically Balanced Number

---

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

An integer `x` is **numerically balanced** if for every digit `d` present in `x`, the digit `d` appears exactly `d` times in `x`.
Given an integer `n`, return the **smallest numerically balanced number strictly greater than `n`**.

Example:

* Input: `n = 1` → Output: `22` (digit `2` appears two times)
* Input: `n = 1000` → Output: `1333` (digit `1` once, digit `3` three times)

---

## Constraints

* `0 <= n <= 10^6` (as per the problem statement on LeetCode)

Important observation: digit `0` cannot appear in a numerically balanced number because, if it appears, it appears > 0 times, which contradicts the rule that it must appear 0 times.

---

## Intuition

I thought about what the rule means: if digit `d` appears at all in the number, it must appear exactly `d` times. So digits allowed are 1..9 and their counts must match their values. Because `n` is at most `10^6`, a simple and clear approach is to test each integer starting from `n+1` and check whether it satisfies the property. The digit counting check is cheap (at most ~7 digits), so brute force search is fast and straightforward.

---

## Approach

1. Start from `x = n + 1`.
2. For each candidate `x`:

   * Count occurrences of each digit `0..9` (use integer math or string methods).
   * If digit `0` appears, reject `x` immediately.
   * For each digit `d` from 1 to 9:

     * If `count[d] != 0` and `count[d] != d`, reject `x`.
   * If all digits pass the test, return `x`.
3. Repeat until a valid number is found.

This approach is easy to implement and fast enough for the input limits. A more advanced approach would precompute all numerically balanced numbers (there are very few) and do a binary search; that is a valid optimization but adds generation complexity.

---

## Data Structures Used

* Fixed-size integer array of length 10 for digit counts (index 0..9).
* A few scalar integers (for candidate number, temporary digit extraction).

All are O(1) extra memory.

---

## Operations & Behavior Summary

* Repeatedly increment candidate integer and check digits.
* Digit extraction uses modulo (`%`) and integer division (`/`) for speed and no extra memory allocation.
* The validation checks are constant-time with respect to the number of digit types (10), and linear in the number of digits of `x` (≤ ~7).

---

## Complexity

* **Time Complexity:**
  Let `gap = answer - n`. For each tested number we examine its digits (call that `D`, number of digits). Time ≈ `O(gap * D)`. For `n <= 10^6`, `D` ≤ 7. Practically fast.
  If we think of the worst theoretical search up to some upper bound `U`, complexity is `O(U * log10(U))`.
* **Space Complexity:**
  `O(1)` — fixed small-size array of 10 integers and a few extra scalars.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    // Check if x is numerically balanced
    bool isBalanced(int x) {
        int cnt[10] = {0};
        int t = x;
        while (t > 0) {
            cnt[t % 10]++;
            t /= 10;
        }
        if (cnt[0] > 0) return false; // 0 cannot appear
        for (int d = 1; d <= 9; ++d) {
            if (cnt[d] != 0 && cnt[d] != d) return false;
        }
        return true;
    }

    int nextBeautifulNumber(int n) {
        int x = n + 1;
        while (true) {
            if (isBalanced(x)) return x;
            ++x;
        }
        return -1; // unreachable
    }
};

// Example main (optional)
int main() {
    Solution sol;
    cout << sol.nextBeautifulNumber(1) << '\n';    // 22
    cout << sol.nextBeautifulNumber(1000) << '\n'; // 1333
    return 0;
}
```

### Java

```java
class Solution {
    private boolean isBalanced(int x) {
        int[] cnt = new int[10];
        int t = x;
        while (t > 0) {
            cnt[t % 10]++;
            t /= 10;
        }
        if (cnt[0] > 0) return false; // 0 cannot appear
        for (int d = 1; d <= 9; ++d) {
            if (cnt[d] != 0 && cnt[d] != d) return false;
        }
        return true;
    }

    public int nextBeautifulNumber(int n) {
        int x = n + 1;
        while (true) {
            if (isBalanced(x)) return x;
            x++;
        }
    }

    // Example main (optional)
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.nextBeautifulNumber(1));    // 22
        System.out.println(s.nextBeautifulNumber(1000)); // 1333
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var nextBeautifulNumber = function(n) {
    const isBalanced = (x) => {
        const cnt = new Array(10).fill(0);
        let t = x;
        while (t > 0) {
            cnt[t % 10]++;
            t = Math.floor(t / 10);
        }
        if (cnt[0] > 0) return false;
        for (let d = 1; d <= 9; ++d) {
            if (cnt[d] !== 0 && cnt[d] !== d) return false;
        }
        return true;
    };

    let x = n + 1;
    while (true) {
        if (isBalanced(x)) return x;
        x++;
    }
};

// Example usage:
console.log(nextBeautifulNumber(1));    // 22
console.log(nextBeautifulNumber(1000)); // 1333
```

### Python3

```python
class Solution:
    def isBalanced(self, x: int) -> bool:
        cnt = [0]*10
        t = x
        while t > 0:
            cnt[t % 10] += 1
            t //= 10
        if cnt[0] > 0:  # 0 cannot appear
            return False
        for d in range(1, 10):
            if cnt[d] != 0 and cnt[d] != d:
                return False
        return True

    def nextBeautifulNumber(self, n: int) -> int:
        x = n + 1
        while True:
            if self.isBalanced(x):
                return x
            x += 1

# Example usage
if __name__ == "__main__":
    sol = Solution()
    print(sol.nextBeautifulNumber(1))    # 22
    print(sol.nextBeautifulNumber(1000)) # 1333
```

### Go

```go
package main

import "fmt"

func isBalanced(x int) bool {
    var cnt [10]int
    t := x
    for t > 0 {
        cnt[t%10]++
        t /= 10
    }
    if cnt[0] > 0 {
        return false
    }
    for d := 1; d <= 9; d++ {
        if cnt[d] != 0 && cnt[d] != d {
            return false
        }
    }
    return true
}

func nextBeautifulNumber(n int) int {
    x := n + 1
    for {
        if isBalanced(x) {
            return x
        }
        x++
    }
}

func main() {
    fmt.Println(nextBeautifulNumber(1))    // 22
    fmt.Println(nextBeautifulNumber(1000)) // 1333
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic for the C++ version line-by-line. The other versions are direct translations of the same logic.

**Helper function `isBalanced(x)`**

* `int cnt[10] = {0};` — I create an array to store counts for digits 0..9.
* `int t = x; while (t > 0) { cnt[t % 10]++; t /= 10; }` — I extract digits by repeatedly taking `t % 10` (last digit) and then divide `t` by 10. This counts how many times each digit appears.
* `if (cnt[0] > 0) return false;` — If `0` appears at all, `x` cannot be numerically balanced. So I immediately return `false`.
* `for (int d = 1; d <= 9; ++d) { if (cnt[d] != 0 && cnt[d] != d) return false; }` — For every digit `d` from 1 to 9: if it appears (`cnt[d] != 0`), it must appear exactly `d` times. If not, `x` is invalid.
* `return true;` — If all checks pass, `x` is numerically balanced.

**Main function `nextBeautifulNumber(n)`**

* `int x = n + 1;` — I start search from `n+1` because the answer must be strictly greater than `n`.
* `while (true) { if (isBalanced(x)) return x; ++x; }` — For each integer `x`, I check if it is balanced. If yes, I return it immediately (this ensures it is the smallest such number greater than `n`). Otherwise I increment `x` and continue.

**Why this is correct**

* We are searching integers in increasing order and returning on the first valid one; by enumeration, it must be the smallest.
* Validation checks exactly enforce the definition of numerically balanced numbers.

**Translation notes**

* In Java, Python, JavaScript, and Go I follow identical steps: build a small counter for digits, reject if `0` appears, check counts for digits 1..9, iterate candidates increasing from `n+1`.

---

## Examples

* Input: `n = 1` → Output: `22`
  Explanation: `22` has digit `2` exactly two times.
* Input: `n = 1000` → Output: `1333`
  Explanation: `1333` has digit `1` once and digit `3` three times.
* Input: `n = 3000` → Output: `3133`
  Explanation: `3133` has `1` once and `3` three times.

---

## How to use / Run locally

### C++

1. Save the C++ code to `solution.cpp`.
2. Compile and run:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

1. Save to `Solution.java`.
2. Compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node)

1. Save the JS code to `solution.js`.
2. Run:

```bash
node solution.js
```

### Python3

1. Save to `solution.py`.
2. Run:

```bash
python3 solution.py
```

### Go

1. Save to `main.go`.
2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* This brute-force check is sufficient and simple for the problem because `n` is limited to `10^6` and valid numbers are not far apart in practice.
* **Optimization alternative:** Precompute all numerically balanced numbers by enumerating possible digit-multisets where for each chosen digit `d` you include exactly `d` copies. Then generate permutations (unique) to form integers, collect and sort them; finally binary-search the answer. This precomputation is small because the total length of such numbers is limited. This yields `O(1)` query time afterwards but needs careful implementation for generation and de-duplication.
* Avoid converting to strings repeatedly if you need absolute speed — integer digit extraction is faster and uses less memory.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
