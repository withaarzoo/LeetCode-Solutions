# 3228. Maximum Number of Operations to Move Ones to the End

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

You are given a binary string `s`. You can perform the following operation any number of times:

* Choose any index `i` where `i + 1 < s.length`, `s[i] == '1'` and `s[i+1] == '0'`.
* Move the character `s[i]` (the `'1'`) to the right until it reaches either the end of the string or the next `'1'` — i.e., slide it over any zeros until blocked.

Return the **maximum** number of operations you can perform.

---

## Constraints

* `1 <= s.length <= 10^5`
* `s[i]` is either `'0'` or `'1'`

**Important note about numeric sizes:** the answer can grow as O(n²) in the worst case (for example many ones on left and many zeros). Use 64-bit integers (or language equivalent big integer) to avoid overflow.

---

## Intuition

I thought about what one operation actually does. An operation needs a `1` followed immediately by a `0` (pattern `"10"`). When I move that `1` right, it can slide across zeros until it meets another `1` or the end. Over the whole process the only zeros that matter for counting operations are those zeros that are immediate boundaries where `1`s can "use" them to land (commonly zeros that lie right after some `1` at some point).

So I realized: scanning left-to-right, if I keep a running count `ones` of how many `1`s I've seen so far, then when I meet a zero that is immediately after a `1` (i.e., a boundary zero), that zero can be involved with each of the previously seen `1`s in an operation (in some order). So I add `ones` to the answer for such zeros.

This gives a linear time, constant extra space solution.

---

## Approach

1. Initialize `ans = 0` (total operations) and `ones = 0` (how many `'1'` seen so far).
2. Iterate the string left-to-right by index `i`:

   * If `s[i] == '1'` then `ones += 1`.
   * Else (it's `'0'`), check whether `s[i-1] == '1'` (i.e., this zero is right after a `1`). If yes, `ans += ones`.
3. At the end return `ans`.

Why this works (intuitive): each boundary zero right after a `1` represents a place that every `1` to its left could use as their operation's "movement spot" at some step. Summing `ones` for every such zero accumulates the maximum number of valid operations.

---

## Data Structures Used

* Simple integer counters:

  * `ones` (64-bit) — counts `'1'` seen so far.
  * `ans` (64-bit) — accumulates the result.
* No arrays or extra containers are necessary beyond reading the string.

---

## Operations & Behavior Summary

* For every `'1'` encountered: increment the `ones` counter.
* For every `'0'` encountered that is immediately after a `'1'`: add the current `ones` to the `ans`.
* End: `ans` is the maximum number of operations.

---

## Complexity

* **Time Complexity:** `O(n)` where `n = s.length`. We make a single pass over the string.
* **Space Complexity:** `O(1)` extra space — only a few integer variables are used.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxOperations(string s) {
        long long ans = 0;    // 64-bit to avoid overflow
        long long ones = 0;   // count of '1' seen so far
        int n = s.size();
        for (int i = 0; i < n; ++i) {
            if (s[i] == '1') {
                ++ones;
            } else { // s[i] == '0'
                if (i > 0 && s[i-1] == '1') ans += ones;
            }
        }
        return (int)ans;
    }
};
```

---

### Java

```java
public class Solution {
    public int maxOperations(String s) {
        long ans = 0L;    // use long to avoid overflow
        long ones = 0L;
        int n = s.length();
        for (int i = 0; i < n; ++i) {
            char c = s.charAt(i);
            if (c == '1') {
                ++ones;
            } else { // c == '0'
                if (i > 0 && s.charAt(i - 1) == '1') ans += ones;
            }
        }
        return (int)ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {string} s
 * @return {number}
 */
var maxOperations = function(s) {
    // Use BigInt to be safe with large intermediate sums
    let ans = 0n;
    let ones = 0n;
    for (let i = 0; i < s.length; ++i) {
        if (s[i] === '1') {
            ones += 1n;
        } else { // '0'
            if (i > 0 && s[i-1] === '1') ans += ones;
        }
    }
    return Number(ans); // convert back to Number for return (safe under constraints)
};
```

---

### Python3

```python
class Solution:
    def maxOperations(self, s: str) -> int:
        ans = 0          # Python int is unbounded
        ones = 0         # count of '1's seen so far
        for i, c in enumerate(s):
            if c == '1':
                ones += 1
            else:  # c == '0'
                if i > 0 and s[i-1] == '1':
                    ans += ones
        return ans
```

---

### Go

```go
package main

func maxOperations(s string) int {
    var ans int64 = 0
    var ones int64 = 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            ones++
        } else { // s[i] == '0'
            if i > 0 && s[i-1] == '1' {
                ans += ones
            }
        }
    }
    return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the approach line-by-line for each language using the same logical mapping.

### Core idea (applies to all versions)

* Maintain `ones`: number of `'1'` characters seen so far while scanning left to right.
* When you find a `'0'` that follows a `'1'` (i.e., `s[i-1] == '1'`), add `ones` to `ans`. This counts how many different `1`s to the left could use this zero in some operation order.
* Return `ans` at the end.

---

### C++ Explanation (mapping to code)

```c++
long long ans = 0;    // store final number of operations (64-bit)
long long ones = 0;   // number of '1's seen so far
int n = s.size();
for (int i = 0; i < n; ++i) {
    if (s[i] == '1') ++ones;                 // increment ones when we see '1'
    else {
        if (i > 0 && s[i-1] == '1') ans += ones; // if zero just after a '1', add ones
    }
}
return (int)ans; // cast result to int (values expected to fit in problem constraints)
```

* `ans` is 64-bit to avoid overflow in large `n`.
* Only scanning once — O(n).

---

### Java Explanation

```java
long ans = 0L;         // accumulator (use long)
long ones = 0L;        // count of '1's seen
for (int i = 0; i < n; ++i) {
    char c = s.charAt(i);
    if (c == '1') ++ones;                      // saw a '1'
    else {
        if (i > 0 && s.charAt(i - 1) == '1')   // zero after a one?
            ans += ones;                       // collect contribution
    }
}
return (int) ans;
```

* Same logic as C++ but Java-specific string access.

---

### JavaScript Explanation

```javascript
let ans = 0n;   // BigInt accumulator
let ones = 0n;  // BigInt ones counter
for (let i = 0; i < s.length; ++i) {
    if (s[i] === '1') ones += 1n;                   // increment ones
    else {
        if (i > 0 && s[i-1] === '1') ans += ones;   // boundary zero adds ones
    }
}
return Number(ans);
```

* `BigInt` used because JavaScript `Number` can lose precision for very large integers. Convert result back to `Number` for typical uses.

---

### Python3 Explanation

```python
ans = 0
ones = 0
for i, c in enumerate(s):
    if c == '1':
        ones += 1
    else:  # c == '0'
        if i > 0 and s[i-1] == '1':
            ans += ones
return ans
```

* Python `int` is unbounded so no overflow worry.

---

### Go Explanation

```go
var ans int64 = 0
var ones int64 = 0
for i := 0; i < len(s); i++ {
    if s[i] == '1' { ones++ }               // count ones
    else {
        if i > 0 && s[i-1] == '1' {
            ans += ones                      // add ones when zero follows a one
        }
    }
}
return int(ans)
```

* Use `int64` for intermediate sums, return `int`.

---

## Examples

1. Example 1:

   * Input: `s = "1001101"`
   * Walkthrough (short): scan and accumulate `ones` and add `ones` at boundary zeros.
   * Output: computed by the algorithm (you can test with code).

2. Example 2:

   * Input: `s = "00111"`
   * Output: `0` (there is no `"10"` boundary initially so no operations possible).

(You can run the provided code with these inputs to see the concrete numbers.)

---

## How to use / Run locally

### Prerequisites

* A C++ compiler (`g++`), Java JDK, Node.js for JS, Python3, and Go installed if you want to run all languages.

### C++

Save code to `solution.cpp` (wrap the method into `main` if you want to test), then:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

Save class in `Solution.java` (ensure class name matches file). Compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node.js)

Save as `solution.js` and include a small harness to call `maxOperations` and print result:

```bash
node solution.js
```

### Python3

Save as `solution.py` with driver code:

```bash
python3 solution.py
```

### Go

Save as `main.go`, ensure function is called from `main`, then:

```bash
go run main.go
```

**Note:** The snippets above contain only the solution function/class. To test, wrap the function with a small `main`/driver that reads a sample string, calls the function, and prints the result.

---

## Notes & Optimizations

* This solution is already optimal in time `O(n)` and space `O(1)`.
* Using 64-bit integers (or equivalent) is important because in worst-case scenarios the answer can be as large as O(n²) and exceed 32-bit limits.
* The logic depends on scanning left-to-right to ensure `ones` counts only `1`s to the left of the current position.
* The check `if i > 0 && s[i-1] == '1'` ensures we only add contributions for "boundary zeros" which actually enable an operation.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
