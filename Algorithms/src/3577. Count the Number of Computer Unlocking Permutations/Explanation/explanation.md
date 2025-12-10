# 3577. Count the Number of Computer Unlocking Permutations

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

I am given an array `complexity` of length `n`.
There are `n` computers labeled from `0` to `n-1`.

* Each computer `i` has a password with complexity `complexity[i]`.
* Computer `0`’s password is **already decrypted** (we start from it).
* For any computer `i > 0`, I can decrypt its password using some computer `j` if:

  * `j < i` and
  * `complexity[j] < complexity[i]`

I need to count how many permutations of `[0, 1, 2, ..., n-1]` are valid unlocking orders, **starting from computer 0 as the only initially unlocked one**.
The answer should be given modulo `1e9 + 7`.

Key note from the problem/editorial: computer **label 0’s password** is decrypted first (so in every valid permutation, index `0` must appear first).

---

## Constraints

* `2 <= complexity.length <= 10^6`
* `1 <= complexity[i] <= 10^9`
* Result is taken modulo `1_000_000_007`
* Need an `O(n)` or close to `O(n)` solution (because `n` can be up to 1e6)

---

## Intuition

When I first read the problem, I thought:

* Every computer needs a **smaller index** with **smaller complexity** to unlock it.
* But we already have computer `0` unlocked.

Then I noticed an important fact:

* If there exists some computer with complexity **smaller** than `complexity[0]` and index `> 0`,
  that computer can never be unlocked.
  Why? Because there is no `j < i` with smaller complexity than it.
* Also, if any computer (other than 0) has the **same** complexity as `complexity[0]`,
  then that one also can’t be unlocked, because we need **strictly smaller** complexity.

So for the whole unlocking process to be possible:

> `complexity[0]` must be the **unique minimum** value in the entire array.

Once this condition is true, computer `0` can unlock **every** other computer directly because:
`complexity[0] < complexity[i]` for all `i > 0`.

That means:

* The first position in the permutation is **fixed** as `0`.
* The remaining `n-1` indices can be arranged in **any order**.

Number of such permutations is simply: **`(n-1)!`**, taken modulo `1e9 + 7`.

---

## Approach

In simple steps, my approach is:

1. Let `n` be the length of `complexity`.
2. Scan the whole array once:

   * Find `minVal` = smallest complexity.
   * Count how many times `minVal` appears (`cntMin`).
3. Check the validity:

   * If `complexity[0] != minVal` → computer 0 is **not** the smallest; impossible ⇒ answer `0`.
   * If `cntMin != 1` → smallest complexity is **not unique**; some other index with min value can’t be unlocked ⇒ answer `0`.
4. If both checks pass:

   * Computer 0 is **unique smallest**.
   * It can unlock all other computers.
   * The remaining `n-1` computers (`1..n-1`) can be permuted however we want.
5. Compute `(n - 1)!` modulo `1e9 + 7`:

   * Initialize `ans = 1`.
   * Loop `i` from `2` to `n-1`:

     * `ans = (ans * i) % MOD`.
6. Return `ans`.

---

## Data Structures Used

* Only **basic variables**:

  * `int` / `long` / `long long` / `int64` for factorial computation.
  * `int` values for `minVal` and `cntMin`.
* No extra arrays, sets, maps etc.

So memory usage is constant, apart from the input array itself.

---

## Operations & Behavior Summary

1. **Single pass for minimum value:**

   * Compare each element with current `minVal`.
   * Track how many times min appears.

2. **Validation checks:**

   * Ensure `complexity[0]` equals global minimum.
   * Ensure the minimum appears exactly once in the array.

3. **Factorial computation:**

   * Multiply numbers from `2` to `n-1`.
   * Apply modulo at each step to avoid overflow.

4. **Result:**

   * If invalid configuration → return `0`.
   * Otherwise → return `(n-1)! % MOD`.

---

## Complexity

* **Time Complexity:** `O(n)`

  * `n` is the length of `complexity`.
  * One scan to find min and count, and one loop to compute factorial up to `n-1`.

* **Space Complexity:** `O(1)`

  * I only use a constant number of extra variables, no extra data structures.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countPermutations(vector<int>& complexity) {
        const int MOD = 1'000'000'007;
        int n = (int)complexity.size();
        
        // 1. Find global minimum and its frequency
        int minVal = complexity[0];
        int cntMin = 0;
        for (int x : complexity) {
            if (x < minVal) {
                minVal = x;
                cntMin = 1;
            } else if (x == minVal) {
                cntMin++;
            }
        }
        
        // 2. Check if index 0 has unique minimum
        if (complexity[0] != minVal || cntMin != 1) {
            return 0;
        }
        
        // 3. Compute (n - 1)! % MOD
        long long ans = 1;
        for (int i = 2; i <= n - 1; ++i) {
            ans = (ans * i) % MOD;
        }
        return (int)ans;
    }
};
```

---

### Java

```java
class Solution {
    public int countPermutations(int[] complexity) {
        final int MOD = 1_000_000_007;
        int n = complexity.length;

        // 1. Find global minimum and its frequency
        int minVal = complexity[0];
        int cntMin = 0;
        for (int x : complexity) {
            if (x < minVal) {
                minVal = x;
                cntMin = 1;
            } else if (x == minVal) {
                cntMin++;
            }
        }

        // 2. Check if index 0 has unique minimum
        if (complexity[0] != minVal || cntMin != 1) {
            return 0;
        }

        // 3. Compute (n - 1)! % MOD
        long ans = 1;
        for (int i = 2; i <= n - 1; i++) {
            ans = (ans * i) % MOD;
        }
        return (int) ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} complexity
 * @return {number}
 */
var countPermutations = function(complexity) {
    const MOD = 1_000_000_007n;
    const n = complexity.length;

    // 1. Find global minimum and its frequency
    let minVal = complexity[0];
    let cntMin = 0;
    for (const x of complexity) {
        if (x < minVal) {
            minVal = x;
            cntMin = 1;
        } else if (x === minVal) {
            cntMin++;
        }
    }

    // 2. Check if index 0 has unique minimum
    if (complexity[0] !== minVal || cntMin !== 1) {
        return 0;
    }

    // 3. Compute (n - 1)! % MOD using BigInt
    let ans = 1n;
    for (let i = 2n; i <= BigInt(n - 1); i++) {
        ans = (ans * i) % MOD;
    }

    return Number(ans);
};
```

---

### Python3

```python
class Solution:
    def countPermutations(self, complexity: List[int]) -> int:
        MOD = 10**9 + 7
        n = len(complexity)

        # 1. Find global minimum and its count
        min_val = min(complexity)
        cnt_min = complexity.count(min_val)

        # 2. Check if index 0 has unique minimum
        if complexity[0] != min_val or cnt_min != 1:
            return 0

        # 3. Compute (n - 1)! % MOD
        ans = 1
        for i in range(2, n):
            ans = (ans * i) % MOD
        return ans
```

---

### Go

```go
package main

func countPermutations(complexity []int) int {
 const MOD int64 = 1_000_000_007
 n := len(complexity)

 // 1. Find global minimum and its count
 minVal := complexity[0]
 cntMin := 0
 for _, x := range complexity {
  if x < minVal {
   minVal = x
   cntMin = 1
  } else if x == minVal {
   cntMin++
  }
 }

 // 2. Check if index 0 has unique minimum
 if complexity[0] != minVal || cntMin != 1 {
  return 0
 }

 // 3. Compute (n - 1)! % MOD
 ans := int64(1)
 for i := 2; i <= n-1; i++ {
  ans = (ans * int64(i)) % MOD
 }
 return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Here I’ll explain the main idea once, and then small notes for each language.

### Core Logic (common to all languages)

1. **Find the global minimum and its frequency**

   ```text
   minVal = complexity[0]
   cntMin = 0
   for x in complexity:
       if x < minVal:
           minVal = x
           cntMin = 1
       else if x == minVal:
           cntMin++
   ```

   * I start assuming the first element is the minimum.
   * If I find a strictly smaller value, I update `minVal` and reset `cntMin` to 1.
   * If I find the same as current min, I increase `cntMin`.

2. **Check if computer 0 is the unique minimum**

   ```text
   if complexity[0] != minVal or cntMin != 1:
       return 0
   ```

   * If index `0` is not holding the smallest complexity, some smaller-complexity computer can never be unlocked.
   * If the smallest value appears more than once, then any other index with that value cannot be unlocked (no strictly smaller complexity exists).
   * In both cases, there is **no valid permutation**, so answer is `0`.

3. **Compute factorial `(n - 1)!` under modulo**

   ```text
   ans = 1
   for i from 2 to n - 1:
       ans = (ans * i) % MOD
   return ans
   ```

   * First position in permutation is fixed as `0`.
   * Remaining `n-1` positions can be filled with computers `1..n-1` in any order.
   * Number of ways = permutations of `n-1` distinct items = `(n-1)!`.

### Language-specific notes

* **C++**

  * Uses `long long` for `ans` to avoid overflow before modulo.
  * `MOD` is an `int`, multiplication and modulo are done in `long long`.

* **Java**

  * Uses `long` for `ans`.
  * Cast back to `int` when returning.

* **JavaScript**

  * Uses `BigInt` because factorial can grow very large.
  * `MOD` is `1_000_000_007n` as a BigInt.
  * Convert the final BigInt result back to Number with `Number(ans)`.

* **Python3**

  * Python integers are unbounded, but I still take `% MOD` at each step.
  * Simple `for` loop from `2` to `n-1`.

* **Go**

  * Uses `int64` for safe multiplication.
  * Keeps `MOD` as `int64` and converts at the end to `int`.

---

## Examples

### Example 1

```text
Input: complexity = [1, 2, 3]
```

* `minVal = 1`, appears only once, and at index `0`.
* Valid condition, so answer is `(3 - 1)! = 2! = 2`.
* Possible valid permutations:

  * [0, 1, 2]
  * [0, 2, 1]

```text
Output: 2
```

---

### Example 2

```text
Input: complexity = [3, 3, 3, 4, 4, 4]
```

* `minVal = 3`, appears three times.
* `cntMin != 1` ⇒ more than one computer has the smallest complexity.
* Any computer with min complexity and index > 0 cannot be unlocked.
* So no valid permutation exists.

```text
Output: 0
```

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Inside `main.cpp`, include LeetCode `Solution` class and create your own tests in `main()`.

---

### Java

```bash
javac Solution.java
java Solution
```

Wrap the class in a test harness or run on LeetCode’s online judge.

---

### JavaScript (Node.js)

```bash
node main.js
```

Where `main.js` exports/uses `countPermutations` and you call it with sample arrays.

---

### Python3

```bash
python3 main.py
```

Create a `Solution` object and call `countPermutations` with desired test cases.

---

### Go

```bash
go run main.go
```

In `main.go`, define the `countPermutations` function and call it from `main()` with some sample inputs.

---

## Notes & Optimizations

* The key trick is to realize that:

  * Either there are **0** valid permutations, or there are **exactly `(n-1)!`**.
  * There is no “in-between” count.
* We reduce the whole problem to:

  * A simple minimum check.
  * A factorial computation.
* The solution is `O(n)` and uses `O(1)` extra space, so it works even for `n = 10^6`.
* Since factorial grows very fast, always keep modulo in the loop to avoid overflow.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
