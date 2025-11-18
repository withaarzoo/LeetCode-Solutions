# 717. 1-bit and 2-bit Characters

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

We are given an array `bits` consisting of only `0` and `1`. The array represents characters encoded as:

* A **one-bit character** is represented by a single `0`.
* A **two-bit character** is represented by `10` or `11`.

We are guaranteed that the given `bits` array ends with a `0`. Determine whether the last character is a one-bit character (`0`) or if it is part of a two-bit character that ends at the final index.

Return `true` if the last character must be a one-bit character; otherwise return `false`.

---

## Constraints

* `1 <= bits.length <= 1000`
* Each `bits[i]` is either `0` or `1`.
* The last element in `bits` is always `0` (problem guarantee).

---

## Intuition

I thought about simulating the decoding process from left to right. If I see a `1`, it must be the start of a two-bit character, so I skip two positions. If I see a `0`, it is a one-bit character, so I skip one position. I keep doing this until I reach or pass the final bit. If my pointer lands exactly on the last index, the last character is a one-bit character. If I pass it, then the last `0` was part of a two-bit character.

---

## Approach

1. Initialize an index pointer `i = 0`.
2. While `i < n - 1` (stop before the last element):

   * If `bits[i] == 1`, then this must be a two-bit character → `i += 2`.
   * Else (`bits[i] == 0`) → one-bit character → `i += 1`.
3. After the loop ends:

   * If `i == n - 1`, return `true` (last bit is a one-bit character).
   * Else return `false` (last bit was consumed as part of a two-bit character).

This approach reads the array once and uses only constant extra memory.

---

## Data Structures Used

* Input: array / list of integers (`bits[]`).
* No extra data structures are required.
* A few integer variables for indices and lengths.

---

## Operations & Behavior Summary

* Single traverse of the `bits` array from left to right.
* Conditional steps:

  * On `1` → move two indices forward.
  * On `0` → move one index forward.
* Final check to see whether pointer points to the last index.

---

## Complexity

* **Time Complexity:** O(n), where `n = bits.length`. Each index is processed at most once.
* **Space Complexity:** O(1). Only a few integer variables are used; no extra arrays or recursion.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    bool isOneBitCharacter(vector<int>& bits) {
        int n = bits.size();
        int i = 0;
        // Process until before the last bit
        while (i < n - 1) {
            if (bits[i] == 1) {
                // 1 starts a two-bit character => skip two
                i += 2;
            } else {
                // 0 is a one-bit character => skip one
                i += 1;
            }
        }
        // If we land exactly at last index, last char is one-bit
        return i == n - 1;
    }
};
```

---

### Java

```java
class Solution {
    public boolean isOneBitCharacter(int[] bits) {
        int n = bits.length;
        int i = 0;
        while (i < n - 1) {
            if (bits[i] == 1) {
                // Start of a two-bit character
                i += 2;
            } else {
                // One-bit character
                i += 1;
            }
        }
        return i == n - 1;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} bits
 * @return {boolean}
 */
var isOneBitCharacter = function(bits) {
    const n = bits.length;
    let i = 0;
    while (i < n - 1) {
        if (bits[i] === 1) {
            i += 2;
        } else {
            i += 1;
        }
    }
    return i === n - 1;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        n = len(bits)
        i = 0
        while i < n - 1:
            if bits[i] == 1:
                i += 2
            else:
                i += 1
        return i == n - 1
```

---

### Go

```go
package main

// isOneBitCharacter returns true if the last character must be a one-bit character.
func isOneBitCharacter(bits []int) bool {
    n := len(bits)
    i := 0
    for i < n-1 {
        if bits[i] == 1 {
            i += 2
        } else {
            i += 1
        }
    }
    return i == n-1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the important lines in a simple, step-by-step way. The logic is identical across all languages, so I'll describe the behavior and give language-specific pointers.

1. **Initialize variables**

   ```pseudo
   n = bits.length
   i = 0
   ```

   * `n` is the total number of bits.
   * `i` is my pointer that scans the array from left to right.

2. **Main loop condition**

   ```pseudo
   while i < n - 1:
   ```

   * I loop while `i` is strictly less than `n-1`. I stop before the last element because I want to see whether the last element remains to be consumed. If a two-bit character consumes it, we will exit with `i >= n`.

3. **If current bit is 1**

   ```pseudo
   if bits[i] == 1:
       i += 2
   ```

   * A `1` can only be the start of a two-bit character (`10` or `11`). Therefore, I skip the next bit too. Move the pointer two steps forward.

   * Language notes:

     * C++: `if (bits[i] == 1) i += 2;`
     * Java: `if (bits[i] == 1) i += 2;`
     * JS: `if (bits[i] === 1) i += 2;`
     * Python: `if bits[i] == 1: i += 2`
     * Go: `if bits[i] == 1 { i += 2 }`

4. **If current bit is 0**

   ```pseudo
   else:
       i += 1
   ```

   * `0` is a one-bit character. Move pointer one step forward.

5. **After loop — final check**

   ```pseudo
   return i == n - 1
   ```

   * If the pointer is exactly on the last index (`n-1`) after decoding earlier characters, it means the last `0` is still unconsumed — a one-bit character → return `true`.
   * If the pointer is `n`, it means the last `0` was consumed as part of a two-bit character → return `false`.

### Why `while i < n - 1`?

* We stop the simulation when `i >= n - 1` because the last index is the one we need to check. If `i` becomes `n`, that means a previous `1` consumed the last `0` (it formed a two-bit char ending at the last index), so the last bit is not a standalone one-bit character.

---

## Examples

1. `bits = [1, 0, 0]`

   * Process: `i=0` → `bits[0]=1` → skip two (`i=2`). Loop stops. `i == n-1` (2==2) → `true`.

2. `bits = [1, 1, 1, 0]`

   * Process: `i=0` → `bits[0]=1` → `i=2`. Then `bits[2]=1` → `i=4`. `i` is past last index (`4 > 3`) → `false`.

3. `bits = [0]`

   * `i=0`, loop condition `i < n-1` is `0 < 0` false, so directly `i == n-1` true → `true`.

---

## How to use / Run locally

### C++

1. Put the solution class into a `.cpp` file along with a `main()` that reads input and calls the function, or run on an online judge.
2. Compile:

   ```bash
   g++ -std=c++17 solution.cpp -O2 -o solution
   ./solution
   ```

   (If using LeetCode, simply paste the class in the online editor.)

### Java

1. Save as `Solution.java` and add a `main` for local testing if required.
2. Compile & run:

   ```bash
   javac Solution.java
   java Solution
   ```

### JavaScript (Node)

1. Save the function in `solution.js` and add test harness code to call the function.
2. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save as `solution.py`. Wrap the class usage in a `main()` for local tests or run in interactive judge.
2. Run:

   ```bash
   python3 solution.py
   ```

### Go

1. Save code in `main.go`, include a `main()` and call `isOneBitCharacter` for local tests.
2. Run:

   ```bash
   go run main.go
   ```

---

## Notes & Optimizations

* This greedy simulation is optimal for this problem: it is linear in time and constant in space.
* There is also an alternative interpretation using parity counting of consecutive `1`s before the trailing zero (count how many consecutive `1`s appear immediately before the last `0`: if that count is even → last `0` is one-bit, else it's part of two-bit). That method is O(n) as well but requires iterating from the end to count consecutive `1`s.
* The left-to-right simulation is clearer when teaching and easy to implement without edge-case pitfalls.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
