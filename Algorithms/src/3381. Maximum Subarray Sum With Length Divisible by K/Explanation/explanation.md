# 3381. Maximum Subarray Sum With Length Divisible by K

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

I am given:

* An array of integers `nums`.
* An integer `k`.

I need to:

> Return the **maximum sum** of any **subarray** such that the **length** of that subarray is **divisible by `k`**.

If I pick a subarray `nums[l..r]`:

* length = `r - l + 1`
* I only consider this subarray if `(r - l + 1) % k == 0`
* Among all such valid subarrays, I return the **maximum sum**.

---

## Constraints

* `1 <= k <= nums.length <= 2 * 10^5`
* `-10^9 <= nums[i] <= 10^9`

These constraints mean:

* I **cannot** use an `O(n^2)` brute force solution.
* I need something around `O(n)` or `O(n log n)`.

---

## Intuition

My thought process:

1. Whenever I see “maximum subarray sum”, I immediately think about **prefix sums**.

2. For a subarray `nums[l..r]`,
   `sum = prefix[r+1] - prefix[l]`.

3. I also care about the **length** being divisible by `k`:

   ```text
   length = r - l + 1
   (r - l + 1) % k == 0
   (r + 1 - l) % k == 0
   ```

4. Notice that `r + 1` and `l` are **indices of prefix sums**.
   So I need `(r + 1) % k == l % k`.

5. So if two prefix indices have the **same remainder** when divided by `k`,
   then the subarray between them has length divisible by `k`.

6. For each remainder `rem`, if I know the **smallest prefix sum** seen so far at some index with that remainder, then for any future index with the same remainder:

   * The best subarray ending here (with length % k == 0) is:
     `currentPrefix - minPreviousPrefixWithSameRemainder`.

7. So the plan is:

   * Use a running `prefix sum`.
   * Track for each `rem = index % k` the **minimum prefix sum** seen so far.
   * Each time I see the same remainder again, I update my answer.

---

## Approach

1. Let `prefix[i]` be the sum of the first `i` elements (`prefix[0] = 0`).
2. I don’t actually store the whole prefix array; I just maintain a running `prefix` variable.
3. I keep an array `minPref[0..k-1]`:

   * `minPref[rem]` = smallest prefix sum seen at any prefix index whose `index % k == rem`.
4. Initialization:

   * Before processing elements:

     * prefix index = `0`, sum = `0`, and `0 % k = 0`.
     * So `minPref[0] = 0`.
   * Set all other entries of `minPref` to a very large value `INF`.
5. For each element `nums[i]` (0-based):

   * Update `prefix += nums[i]`.
   * Current prefix index is `i + 1`.
   * Compute `rem = (i + 1) % k`.
   * If `minPref[rem]` is not `INF`, then there exists some earlier prefix index with same remainder:

     * Candidate subarray sum = `prefix - minPref[rem]`.
     * Update the answer with this candidate.
   * Then update:

     * `minPref[rem] = min(minPref[rem], prefix)`.
6. At the end, the answer is the maximum subarray sum whose length is divisible by `k`.

---

## Data Structures Used

* `long long` / `long` / `int64` variable for **running prefix sum**.
* An array `minPref` of size `k`:

  * Type: `long long[]`, `long[]`, `number[]`, etc.
  * Purpose: for each remainder `rem`, store the **minimum prefix sum** seen so far at indices whose `index % k == rem`.

No other complex data structures are needed.

---

## Operations & Behavior Summary

During the single pass over `nums`:

For each index `i`:

1. **Update prefix sum**:

   * `prefix += nums[i]`.
2. **Compute remainder of prefix index**:

   * `rem = (i + 1) % k`.
3. **Check existing remainder**:

   * If we already saw some previous index with remainder `rem`, we form a candidate subarray sum:

     * `candidate = prefix - minPref[rem]`.
   * Update the global maximum `ans`.
4. **Update minimum prefix for this remainder**:

   * `minPref[rem] = min(minPref[rem], prefix)`.

So every index:

* Does constant-time math and comparisons.
* No nested loops.

---

## Complexity

Let `n = nums.length` and `k` as given.

* **Time Complexity:** `O(n + k)`

  * `O(k)` to initialize the `minPref` array.
  * `O(n)` to iterate through `nums` once.
* **Space Complexity:** `O(k)`

  * We only store `minPref` of size `k` and a few variables.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <limits>
using namespace std;

class Solution {
public:
    long long maxSubarraySum(vector<int>& nums, int k) {
        int n = nums.size();
        
        // A large number to represent +infinity (greater than any possible prefix sum)
        const long long INF = 4000000000000000000LL; // 4e18
        
        // minPref[r] = minimum prefix sum for any prefix index with (index % k == r)
        vector<long long> minPref(k, INF);
        
        long long prefix = 0;        // running prefix sum
        long long ans = -INF;        // best answer so far
        
        // prefix index 0 has sum 0 and remainder 0
        minPref[0] = 0;
        
        for (int i = 0; i < n; ++i) {
            prefix += (long long)nums[i];   // update running sum
            int rem = (i + 1) % k;          // prefix index is i+1
            
            // If we've seen this remainder before, try forming a subarray
            if (minPref[rem] != INF) {
                long long candidate = prefix - minPref[rem];
                if (candidate > ans) ans = candidate;
            }
            
            // Update minimum prefix for this remainder
            if (prefix < minPref[rem]) {
                minPref[rem] = prefix;
            }
        }
        
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        int n = nums.length;
        
        // Large number as +infinity
        long INF = 4_000_000_000_000_000_000L; // 4e18
        
        long[] minPref = new long[k];
        for (int i = 0; i < k; i++) {
            minPref[i] = INF;
        }
        
        long prefix = 0;   // running prefix sum
        long ans = -INF;   // best answer so far
        
        // prefix index 0 has sum 0 and remainder 0
        minPref[0] = 0;
        
        for (int i = 0; i < n; i++) {
            prefix += nums[i];         // update running sum
            int rem = (i + 1) % k;     // prefix index is i+1
            
            if (minPref[rem] != INF) {
                long candidate = prefix - minPref[rem];
                if (candidate > ans) ans = candidate;
            }
            
            if (prefix < minPref[rem]) {
                minPref[rem] = prefix;
            }
        }
        
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSubarraySum = function(nums, k) {
    const n = nums.length;
    
    // Max possible |prefix| ≤ n * 1e9 = 2e5 * 1e9 = 2e14
    // So 1e15 is safely "infinity" and still inside JS safe integer range (9e15)
    const INF = 1e15;
    
    const minPref = new Array(k).fill(INF);
    
    let prefix = 0;      // running prefix sum
    let ans = -INF;      // best answer so far
    
    // prefix index 0 has sum 0 and remainder 0
    minPref[0] = 0;
    
    for (let i = 0; i < n; i++) {
        prefix += nums[i];       // update running sum
        const rem = (i + 1) % k; // prefix index is i+1
        
        if (minPref[rem] !== INF) {
            const candidate = prefix - minPref[rem];
            if (candidate > ans) ans = candidate;
        }
        
        if (prefix < minPref[rem]) {
            minPref[rem] = prefix;
        }
    }
    
    return ans;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def maxSubarraySum(self, nums: List[int], k: int) -> int:
        n = len(nums)
        
        # Very large number as +infinity
        INF = 10**18
        
        # minPref[r] = smallest prefix sum where prefix index % k == r
        minPref = [INF] * k
        
        prefix = 0          # running prefix sum
        ans = -INF          # best answer so far
        
        # prefix index 0 has sum 0 and remainder 0
        minPref[0] = 0
        
        for i, val in enumerate(nums):
            prefix += val           # update running sum
            rem = (i + 1) % k       # prefix index is i+1
            
            if minPref[rem] != INF:
                candidate = prefix - minPref[rem]
                if candidate > ans:
                    ans = candidate
            
            if prefix < minPref[rem]:
                minPref[rem] = prefix
        
        return ans
```

---

### Go

```go
package main

func maxSubarraySum(nums []int, k int) int64 {
    const INF int64 = 4_000_000_000_000_000_000 // 4e18
    
    n := len(nums)
    
    minPref := make([]int64, k)
    for i := 0; i < k; i++ {
        minPref[i] = INF
    }
    
    var prefix int64 = 0 // running prefix sum
    var ans int64 = -INF // best answer so far
    
    // prefix index 0 has sum 0 and remainder 0
    minPref[0] = 0
    
    for i := 0; i < n; i++ {
        prefix += int64(nums[i])  // update running sum
        rem := (i + 1) % k        // prefix index is i+1
        
        if minPref[rem] != INF {
            candidate := prefix - minPref[rem]
            if candidate > ans {
                ans = candidate
            }
        }
        
        if prefix < minPref[rem] {
            minPref[rem] = prefix
        }
    }
    
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll explain with the C++ version (the others are the same idea, just different syntax).

### 1. Define constants and data structures

```c++
const long long INF = 4000000000000000000LL; // 4e18
vector<long long> minPref(k, INF);

long long prefix = 0;
long long ans = -INF;

minPref[0] = 0;
```

* `INF` acts like +∞, bigger than any possible prefix sum.
* `minPref` is size `k`, initialized with `INF`.
* `prefix` stores the running sum.
* `ans` keeps the maximum subarray sum found so far.
* We know prefix index `0` has sum `0` and `0 % k = 0`, so `minPref[0] = 0`.

---

### 2. Iterate through the array

```c++
for (int i = 0; i < n; ++i) {
    prefix += (long long)nums[i];
    int rem = (i + 1) % k;
```

* At each step, I add `nums[i]` to `prefix`.
* The current prefix index is `i + 1`.
* I compute `rem = (i + 1) % k` — this tells me which remainder class we’re in.

---

### 3. Use previous minimum prefix with same remainder

```c++
    if (minPref[rem] != INF) {
        long long candidate = prefix - minPref[rem];
        if (candidate > ans) ans = candidate;
    }
```

* If `minPref[rem]` is not `INF`, it means I have seen another prefix index `j` earlier such that `j % k == rem`.
* The subarray between that index and this one has length divisible by `k`.
* Its sum is `prefix - minPref[rem]`.
* I compare this with `ans` and keep the larger one.

---

### 4. Update the minimum prefix for this remainder

```c++
    if (prefix < minPref[rem]) {
        minPref[rem] = prefix;
    }
}
```

* Now I update the best starting prefix for this remainder.
* For future indices with same `rem`, this smaller prefix will give them a bigger subarray sum.

---

### 5. Return the final answer

```c++
return ans;
```

* After the loop ends, `ans` holds the maximum subarray sum with length divisible by `k`.

The same logic is implemented in Java, JavaScript, Python, and Go with slight syntax changes.

---

## Examples

### Example 1

```text
Input: nums = [1, 2], k = 1
Output: 3
Explanation:
All subarrays have length divisible by 1.
The best subarray is [1, 2] with sum = 3.
```

### Example 2

```text
Input: nums = [-1, -2, -3, -4], k = 4
Output: -10
Explanation:
Only subarray with length divisible by 4 is the whole array: [-1, -2, -3, -4].
Sum = -10.
```

### Example 3

```text
Input: nums = [-5, 1, 2, -3, 4], k = 2
Output: 4
Explanation:
The maximum subarray with even length is [1, 2, -3, 4].
Length = 4 (divisible by 2), sum = 4.
```

---

## How to use / Run locally

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/<your-repo-name>.git
cd <your-repo-name>
```

### 2. Run per language

#### C++

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

#### Java

```bash
javac Solution.java
java Solution
```

#### JavaScript (Node.js)

```bash
node solution.js
```

#### Python3

```bash
python3 solution.py
```

#### Go

```bash
go run solution.go
```

In each file you can create a `main` function (or quick test harness) that:

* Builds the `nums` array and `k`.
* Calls `maxSubarraySum`.
* Prints the answer.

---

## Notes & Optimizations

* The key optimization is using **prefix sums combined with modular arithmetic** on the **prefix indices** (not on the prefix sums themselves).
* We never store all prefix sums, only:

  * current prefix,
  * and best minimum prefix for each remainder.
* This gives us:

  * **Single pass** over the array.
  * **O(k)** extra memory.
* This solution is optimal given the constraints (`n` up to `2 * 10^5`).

Possible extensions:

* If I needed to find **the actual subarray indices**, I could also store:

  * the index of the prefix where `minPref[rem]` occurs, and reconstruct `(start, end)`.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
