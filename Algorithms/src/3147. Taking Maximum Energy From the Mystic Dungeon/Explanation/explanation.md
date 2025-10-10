# 3147. Taking Maximum Energy From the Mystic Dungeon

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

In a line of `n` magicians, each magician `i` has an integer `energy[i]` (which can be positive or negative). Because of a curse, when I absorb energy from magician `i`, I am instantly teleported to magician `i + k`. I choose a starting magician `s`, then visit indices `s, s+k, s+2k, ...` until I move outside the array. I must absorb the energy at every magician I land on. Return the **maximum** possible total energy I can gain by choosing the optimal starting index.

In short: partition indices by `index % k` (each forms a chain). For each chain I can start at any node — choose the best starting point (suffix) and take the maximum among all chains.

---

## Constraints

* `1 <= energy.length <= 10^5`
* `-1000 <= energy[i] <= 1000`
* `1 <= k <= energy.length - 1`

---

## Intuition

I thought about what teleporting by `k` steps means: indexes that are congruent modulo `k` are linked (they form an independent chain). If I pick any starting index `s`, I will only visit elements in the chain `s, s+k, s+2k, ...`. Choosing a starting index inside a chain is equivalent to choosing a **suffix** of that chain (start somewhere and take all elements until the chain end). So I only need to compute, for each chain (residue `r = 0..k-1`), the maximum suffix sum. The answer is the maximum suffix sum among all chains.

---

## Approach

1. Let `n = energy.length`.
2. For each residue `r` from `0` to `k-1`:

   * Build the chain of indices `r, r+k, r+2k, ...` (implicitly).
   * Compute the last index in this chain that is `< n`.
   * Walk the chain **backwards** from `last` to `r` stepping by `-k`, keeping a running sum `cur`. When walking backward, `cur` equals the suffix sum for the current index.
   * Track the maximum `cur` seen for this chain and update the global maximum `ans`.
3. Return `ans`.

I visit each element exactly once → O(n) time. I only use a few scalar variables → O(1) extra space.

---

## Data Structures Used

* Input array `energy` (given).
* Scalar counters and accumulators (`n`, `r`, `cur`, `ans`, `last`).

No additional collections (no extra arrays, no heaps, no maps).

---

## Operations & Behavior Summary

* Partition indices implicitly by `index % k`.
* For each partition, compute suffix sums by traversing from the chain end backwards (this lets me compute every suffix sum in linear time per chain).
* Keep the maximum across all suffix sums and return it.

Behavior with all-negative arrays: because I initialize `ans` to a very small value, the algorithm correctly returns the largest (least negative) single suffix when every number is negative.

---

## Complexity

* **Time Complexity:** `O(n)` where `n = energy.length`. Each element is visited exactly once across all chains.
* **Space Complexity:** `O(1)` extra space (only constant additional variables are used).

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maximumEnergy(vector<int>& energy, int k) {
        int n = energy.size();
        long long ans = LLONG_MIN; // use 64-bit accumulator for safety
        // For each residue class modulo k
        for (int r = 0; r < k; ++r) {
            long long cur = 0;
            // compute last index in this residue class
            int last = r + ((n - 1 - r) / k) * k;
            for (int i = last; i >= r; i -= k) {
                cur += energy[i];      // suffix sum starting at i
                ans = max(ans, cur);   // update global maximum
            }
        }
        return (int)ans;
    }
};

// Example driver (optional)
// int main() {
//     vector<int> energy = {5,2,-10,-5,1};
//     int k = 3;
//     Solution sol;
//     cout << sol.maximumEnergy(energy, k) << endl; // prints 3
//     return 0;
// }
```

### Java

```java
class Solution {
    public int maximumEnergy(int[] energy, int k) {
        int n = energy.length;
        long ans = Long.MIN_VALUE; // use long to safely accumulate sums
        for (int r = 0; r < k; ++r) {
            long cur = 0;
            int last = r + ((n - 1 - r) / k) * k; // last index in this class
            for (int i = last; i >= r; i -= k) {
                cur += energy[i];    // suffix sum starting at i
                ans = Math.max(ans, cur);
            }
        }
        return (int) ans;
    }
}

// Example driver (optional)
// public class Main {
//     public static void main(String[] args) {
//         int[] energy = {5,2,-10,-5,1};
//         int k = 3;
//         Solution s = new Solution();
//         System.out.println(s.maximumEnergy(energy, k)); // 3
//     }
// }
```

### JavaScript

```javascript
/**
 * @param {number[]} energy
 * @param {number} k
 * @return {number}
 */
var maximumEnergy = function(energy, k) {
    const n = energy.length;
    let ans = -Infinity;
    for (let r = 0; r < k; ++r) {
        let cur = 0;
        const last = r + Math.floor((n - 1 - r) / k) * k; // last index in class
        for (let i = last; i >= r; i -= k) {
            cur += energy[i];   // suffix sum starting at i
            if (cur > ans) ans = cur;
        }
    }
    return ans;
};

// Example usage:
// console.log(maximumEnergy([5,2,-10,-5,1], 3)); // 3
```

### Python3

```python
from typing import List

class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        n = len(energy)
        ans = -10**18  # safe initial very small number
        for r in range(k):
            cur = 0
            last = r + ((n - 1 - r) // k) * k
            i = last
            while i >= r:
                cur += energy[i]   # suffix sum starting at i
                ans = max(ans, cur)
                i -= k
        return ans

# Example usage:
# s = Solution()
# print(s.maximumEnergy([5,2,-10,-5,1], 3))  # 3
```

### Go

```go
package main

import "fmt"

func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    var ans int64 = -1 << 62 // very small initial value
    for r := 0; r < k; r++ {
        var cur int64 = 0
        last := r + ((n-1 - r)/k)*k // last index in this residue class
        for i := last; i >= r; i -= k {
            cur += int64(energy[i])   // suffix sum starting at i
            if cur > ans {
                ans = cur
            }
        }
    }
    return int(ans)
}

// Example driver
// func main() {
//     energy := []int{5,2,-10,-5,1}
//     k := 3
//     fmt.Println(maximumEnergy(energy, k)) // 3
// }
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the main idea and then point out small language-specific notes.

### Main idea (applies to all languages)

1. **Compute `n`** — the length of the `energy` array.
2. **Initialize `ans`** with a very small value. This ensures the algorithm works when all numbers are negative.
3. **Loop `r` from `0` to `k-1`**. Each `r` is a residue class modulo `k` and defines a chain: `r, r+k, r+2k, ...`.
4. **Compute `last`** — the largest index `< n` that belongs to this chain. This is done by:

   ```
   last = r + ((n - 1 - r) / k) * k
   ```

   Explanation: `(n-1 - r)/k` is how many full k-steps you can take starting from `r` without exceeding `n-1`. Multiply by `k` and add `r` to get the actual last index.
5. **Traverse from `last` back to `r` stepping by `-k`**:

   * Keep a running `cur` (starting at 0).
   * Add `energy[i]` to `cur` at each step; because we're going backward, `cur` becomes the suffix sum for index `i`.
   * Update `ans = max(ans, cur)`.
6. **Return `ans`** — the maximum suffix sum among all chains.

### Why backward traversal?

If a chain is `[a0, a1, a2, ... , at]` (in increasing index order), the suffix sum starting at position `i` equals `ai + a(i+1) + ... + at`. If I iterate from `at` to `a0` and keep adding, after visiting `ai` the accumulator holds exactly that suffix sum. This is efficient (each element visited once).

### Language considerations

* **C++:** Use `long long` (here `long long ans`), `LLONG_MIN` for initialization to safely handle negative sums.
* **Java:** Use `long` for `ans` and `cur` to avoid overflow while summing many ints; cast to `int` at the end if desired.
* **JavaScript:** Numbers are safe for this problem range; initialize `ans` to `-Infinity`.
* **Python3:** Use a very large negative int (e.g., `-10**18`) as Python's ints are unbounded anyway.
* **Go:** Use `int64` for accumulators (`cur` and `ans`) then cast to `int` at return time (the constraints guarantee safe range).

---

## Examples

1. **Example 1**

   * Input: `energy = [5,2,-10,-5,1]`, `k = 3`
   * Chains:

     * `r=0`: indices `0,3` → values `[5, -5]` → suffix sums `[-5, 0]` → best = `0`.
     * `r=1`: indices `1,4` → values `[2, 1]` → suffix sums `[1, 3]` → best = `3`.
     * `r=2`: indices `2` → `[-10]` → best = `-10`.
   * Answer: `3` (start at index 1 → 2 + 1 = 3).

2. **Example 2**

   * Input: `energy = [-2, -3, -1]`, `k = 2`
   * Chains:

     * `r=0`: indices `0,2` → `[-2, -1]` → suffix sums `[-1, -3]` → best = `-1`.
     * `r=1`: indices `1` → `[-3]` → best = `-3`.
   * Answer: `-1` (start at index 2).

---

## How to use / Run locally

### C++

1. Save the C++ solution into `solution.cpp`.
2. Add a `main()` driver if needed (an example is commented inside the snippet).
3. Compile and run:

   ```bash
   g++ -std=c++17 solution.cpp -O2 -o solution
   ./solution
   ```

### Java

1. Save the class `Solution` and the optional `Main` driver.
2. Compile & run:

   ```bash
   javac Main.java
   java Main
   ```

### JavaScript (Node.js)

1. Save the JS snippet to `solution.js` and add a test call (commented example).
2. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save the Python class to `solution.py` with a test driver or use in LeetCode style.
2. Run:

   ```bash
   python3 solution.py
   ```

### Go

1. Save the Go code to `main.go` and uncomment the `main()` driver if necessary.
2. Build & run:

   ```bash
   go run main.go
   ```

---

## Notes & Optimizations

* The solution is already optimal: O(n) time and O(1) additional space.
* We avoid building explicit chains or extra arrays — we compute indexes arithmetic to navigate chains in-place.
* Using 64-bit accumulators protects against intermediate sums exceeding 32-bit range if `n` is large and values near bounds.
* The algorithm handles negative-only arrays correctly (returns the largest single suffix among all chains).

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
