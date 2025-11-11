# 474. Ones and Zeroes

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

You are given an array of binary strings `strs` and two integers `m` and `n`.
Return the size of the **largest subset** of `strs` that contains **at most** `m` zeros and **at most** `n` ones in total.

This is a 0/1 decision for each string—either I include it or I don’t—and each string has a “cost” of `(zeros, ones)` and a “value” of `+1` item in the subset.

---

## Constraints

* `1 <= strs.length <= 600`
* `1 <= strs[i].length <= 100`
* `strs[i]` consists only of `'0'` and `'1'`
* `0 <= m, n <= 100`

---

## Intuition

I looked at every string and thought: only the **count of 0s and 1s** matters, not their order.
This feels like a **knapsack** problem but with **two capacities**: one for zeros (`m`) and one for ones (`n`).
Each string is an item with cost `(z, o)` and value `1`. I want the maximum number of items without exceeding the capacities. That’s a classic **0/1 knapsack in 2D**.

---

## Approach

1. For each string, count zeros `z` and ones `o`.
2. Use a DP table `dp[z][o]` meaning: the best (max count) subset using **at most** `z` zeros and `o` ones.
3. For every string `(z, o)`, iterate the DP **backwards**:

   * for `i = m..z` and `j = n..o`:
     `dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1)`
4. The reverse order guarantees each string is used **at most once** (0/1 behavior).
5. The answer is `dp[m][n]`.

---

## Data Structures Used

* A 2D integer array `dp` of size `(m+1) x (n+1)` storing the best achievable subset size under each `(zeros, ones)` budget.

---

## Operations & Behavior Summary

* **Count (per string):** count `'0'`s (then ones are `len - zeros`).
* **State transition:** try to include the current string if capacity allows, otherwise keep previous best.
* **Backward iteration:** prevents reusing the same string multiple times in one pass.

---

## Complexity

* **Time Complexity:** `O(L * m * n)`
  `L` = number of strings. For each string I may touch the whole `m x n` DP grid.
* **Space Complexity:** `O(m * n)`
  Single 2D DP table, updated in place. No extra per-string DP layers.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int findMaxForm(vector<string>& strs, int m, int n) {
        // dp[i][j] = max strings using at most i zeros and j ones
        vector<vector<int>> dp(m + 1, vector<int>(n + 1, 0));

        for (const string& s : strs) {
            int z = 0, o = 0;
            for (char c : s) (c == '0') ? ++z : ++o;

            // 0/1 knapsack: iterate backwards to avoid reusing s
            for (int i = m; i >= z; --i) {
                for (int j = n; j >= o; --j) {
                    dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1);
                }
            }
        }
        return dp[m][n];
    }
};
```

### Java

```java
class Solution {
    public int findMaxForm(String[] strs, int m, int n) {
        // dp[i][j] = max strings using at most i zeros and j ones
        int[][] dp = new int[m + 1][n + 1];

        for (String s : strs) {
            int z = 0, o = 0;
            for (char c : s.toCharArray()) {
                if (c == '0') z++;
                else o++;
            }

            // 0/1 knapsack: iterate backwards
            for (int i = m; i >= z; i--) {
                for (int j = n; j >= o; j--) {
                    dp[i][j] = Math.max(dp[i][j], dp[i - z][j - o] + 1);
                }
            }
        }
        return dp[m][n];
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} strs
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var findMaxForm = function(strs, m, n) {
  // dp[i][j] = max strings using at most i zeros and j ones
  const dp = Array.from({ length: m + 1 }, () => Array(n + 1).fill(0));

  for (const s of strs) {
    let z = 0, o = 0;
    for (const ch of s) (ch === '0') ? z++ : o++;

    // Backward loops ensure 0/1 usage
    for (let i = m; i >= z; i--) {
      for (let j = n; j >= o; j--) {
        dp[i][j] = Math.max(dp[i][j], dp[i - z][j - o] + 1);
      }
    }
  }
  return dp[m][n];
};
```

### Python3

```python
from typing import List

class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        # dp[i][j] = max strings using at most i zeros and j ones
        dp = [[0] * (n + 1) for _ in range(m + 1)]

        for s in strs:
            z = s.count('0')
            o = len(s) - z

            # 0/1 knapsack: iterate backwards
            for i in range(m, z - 1, -1):
                for j in range(n, o - 1, -1):
                    dp[i][j] = max(dp[i][j], dp[i - z][j - o] + 1)

        return dp[m][n]
```

### Go

```go
package main

func findMaxForm(strs []string, m int, n int) int {
    // dp[i][j] = max strings using at most i zeros and j ones
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for _, s := range strs {
        z, o := 0, 0
        for _, ch := range s {
            if ch == '0' {
                z++
            } else {
                o++
            }
        }

        // 0/1 knapsack: iterate backwards
        for i := m; i >= z; i-- {
            for j := n; j >= o; j-- {
                if dp[i-z][j-o]+1 > dp[i][j] {
                    dp[i][j] = dp[i-z][j-o] + 1
                }
            }
        }
    }
    return dp[m][n]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across languages; here’s the flow you’ll see in every file:

1. **Define DP meaning**
   `dp[i][j]` = the maximum number of strings I can choose using ≤ `i` zeros and ≤ `j` ones.

2. **Initialize**
   Start with all zeros: choosing nothing yields `0`.

3. **For every string `s`**

   * Count `z` = number of `'0'`s.
     `o` = number of `'1'`s = `len(s) - z`.

4. **Backward update (core trick)**

   * For `i` from `m` down to `z`, and `j` from `n` down to `o`:

     * Either skip the string: keep `dp[i][j]`.
     * Or take it: `dp[i - z][j - o] + 1`.
     * Choose the maximum.
   * Going **backwards** ensures each string is used at most once.

5. **Answer**
   After processing all strings, return `dp[m][n]`.

---

## Examples

**Example 1**

```
Input:  strs = ["10","0001","111001","1","0"], m = 5, n = 3
Output: 4
Explanation:
Best subset size is 4, e.g., {"10","0001","1","0"} uses 5 zeros and 3 ones.
```

**Example 2**

```
Input:  strs = ["10","0","1"], m = 1, n = 1
Output: 2
Explanation:
Best subset is {"0","1"} ⇒ uses 1 zero and 1 one.
```

---

## How to use / Run locally

**C++**

```bash
g++ -std=c++17 -O2 main.cpp -o app
./app
```

**Java**

```bash
javac Solution.java
java Solution
```

**JavaScript (Node.js)**

```bash
node main.js
```

**Python3**

```bash
python3 main.py
```

**Go**

```bash
go run main.go
```

> Put the respective function into a small driver (read input, call function, print result) if you want to run with custom inputs from the terminal.

---

## Notes & Optimizations

* The 2D DP is already optimal for constraints (`m,n ≤ 100`).
* If memory is tight, `int16`/`short` can be used since the answer ≤ `strs.length` ≤ 600.
* Precomputing `(zeros, ones)` pairs for all strings once can make the code a bit cleaner.
* Forward iteration would turn this into an **unbounded knapsack** by mistake—**always iterate backwards**.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
