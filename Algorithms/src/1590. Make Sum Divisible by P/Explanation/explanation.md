# 1590. Make Sum Divisible by P

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

I’m given:

* An array `nums` of **positive integers**.
* An integer `p`.

I can **remove one continuous subarray** (maybe empty, but not the whole array).
After removing, the sum of the **remaining elements** should be **divisible by `p`**.

I must return the **minimum length** of such a subarray.
If it's **not possible**, I return `-1`.

---

## Constraints

From the original problem:

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^9`
* `1 <= p <= 10^9`

These constraints tell me:

* I can’t use an `O(n^2)` approach (too slow for `n = 10^5`).
* I must use something close to **O(n)** or **O(n log n)**.

---

## Intuition

I started by thinking about the **total sum** of the array:

* Let `S` be the sum of all elements in `nums`.
* If `S % p == 0`, then the sum is already divisible by `p`.
  So I don’t need to remove anything → answer is `0`.

If it’s **not** divisible, I need to remove some subarray with sum `x` such that:

[
(S - x) % p = 0
]

This means:

[
S % p = x % p
]

So instead of thinking “which subarray can I remove?”, I rephrased it as:

> I need to find the **shortest subarray** whose sum modulo `p` is equal to `S % p`.

This smells like a typical **prefix sum + hashmap** problem where I use remainders.

---

## Approach

Step-by-step, here’s how I solved it:

1. **Compute total sum modulo `p`**

   * I loop through `nums` and compute `total = (total + nums[i]) % p`.
   * Let `need = total`.
   * If `need == 0`, I just return `0`.

2. **Use prefix sums modulo `p`**

   * I maintain a running prefix sum: `prefix = (prefix + nums[i]) % p`.
   * I also maintain a map `lastIndex` that stores:

     * **Key**: remainder value `r` (0 to `p-1`)
     * **Value**: **latest index** where prefix remainder = `r`.

3. **Derive the target remainder**

   * Suppose the prefix remainder at index `i` is `pref`.
   * At some earlier index `j`, the prefix remainder was `pref_j`.
   * The sum of the subarray `(j+1 ... i)` is:
     [
     sub = (pref - pref_j + p) % p
     ]
   * I want `sub % p == need`.
   * So:
     [
     (pref - pref_j) % p = need
     \Rightarrow pref_j = (pref - need + p) % p
     ]
   * So for each `i`, I compute:

     ```text
     target = (prefix - need + p) % p
     ```

     and check if this `target` remainder exists in my map.

4. **Update minimum length**

   * If `target` exists in `lastIndex`, say at index `j`:

     * Then subarray `(j+1 ... i)` is a valid subarray to remove.
     * Its length is `i - j`.
     * I keep track of the **minimum** such length in `ans`.

5. **Maintain latest occurrence**

   * After processing index `i`, I store:

     ```text
     lastIndex[prefix] = i
     ```

   * I always keep the **latest index** for each remainder, because this naturally tends to give **shorter** subarrays when I subtract earlier.

6. **Final result**

   * If `ans == n`, it means the only possible subarray is the whole array (which is not allowed) → return `-1`.
   * Otherwise return `ans`.

---

## Data Structures Used

* **Prefix sum (integer variable)**

  * Stores current prefix sum modulo `p`.

* **Hash map / dictionary / map**

  * Key: remainder `r` (0 to `p-1`), i.e., `prefix % p`.
  * Value: the **latest index** where this remainder was seen.
  * Purpose: quickly find an earlier prefix with a certain remainder to form a subarray with remainder `need`.

In different languages:

* C++: `unordered_map<int, int>`
* Java: `HashMap<Integer, Integer>`
* JavaScript: `Map`
* Python3: `dict`
* Go: `map[int]int`

---

## Operations & Behavior Summary

During a single left-to-right scan of the array:

1. **Update prefix modulo**

   * `prefix = (prefix + nums[i]) % p`

2. **Compute required remainder**

   * `target = (prefix - need + p) % p`
     (or `(prefix - need) % p` where language has safe modulo for negatives like Python)

3. **Lookup in map**

   * If `target` exists:

     * calculate candidate length `len = i - lastIndex[target]`
     * update `ans = min(ans, len)`

4. **Update map**

   * `lastIndex[prefix] = i` (store latest index for this remainder)

5. **Edge setup**

   * `lastIndex[0] = -1` at the beginning to handle subarrays starting from index `0`.

---

## Complexity

* **Time Complexity:** `O(n)`

  * Single pass over the array with constant-time hash map operations on average.
  * `n` is the number of elements in `nums`.

* **Space Complexity:** `O(min(n, p))`

  * At most one entry per remainder (0 to `p-1`), or at most one per prefix.
  * In practice, `O(n)` in the worst case.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int minSubarray(vector<int>& nums, int p) {
        long long total = 0;
        for (int x : nums) {
            total = (total + x) % p;   // total sum modulo p
        }
        
        int need = (int)total;
        if (need == 0) return 0;       // already divisible
        
        int n = nums.size();
        unordered_map<int, int> lastIndex;
        lastIndex.reserve(n * 2);      // small optimization
        lastIndex[0] = -1;             // prefix before start
        
        int ans = n;
        long long prefix = 0;
        
        for (int i = 0; i < n; ++i) {
            prefix = (prefix + nums[i]) % p;
            int prefMod = (int)prefix;
            
            // We want some previous prefix with this target remainder
            int target = prefMod - need;
            if (target < 0) target += p;   // (prefMod - need + p) % p
            
            if (lastIndex.count(target)) {
                ans = min(ans, i - lastIndex[target]);
            }
            
            // Update latest index for this prefix remainder
            lastIndex[prefMod] = i;
        }
        
        return ans == n ? -1 : ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public int minSubarray(int[] nums, int p) {
        long total = 0;
        for (int x : nums) {
            total = (total + x) % p;   // keep modulo
        }
        
        int need = (int) total;
        if (need == 0) return 0;       // already divisible
        
        int n = nums.length;
        Map<Integer, Integer> lastIndex = new HashMap<>();
        lastIndex.put(0, -1);          // prefix before any element
        
        int ans = n;
        long prefix = 0;
        
        for (int i = 0; i < n; i++) {
            prefix = (prefix + nums[i]) % p;
            int prefMod = (int) prefix;
            
            int target = prefMod - need;
            if (target < 0) target += p;   // (prefMod - need + p) % p
            
            if (lastIndex.containsKey(target)) {
                ans = Math.min(ans, i - lastIndex.get(target));
            }
            
            // Store latest index for this remainder
            lastIndex.put(prefMod, i);
        }
        
        return ans == n ? -1 : ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
var minSubarray = function(nums, p) {
    let total = 0;
    for (const x of nums) {
        total = (total + x) % p;   // keep modulo
    }
    
    const need = total;
    if (need === 0) return 0;      // already divisible
    
    const n = nums.length;
    const lastIndex = new Map();
    lastIndex.set(0, -1);          // prefix before index 0
    
    let ans = n;
    let prefix = 0;
    
    for (let i = 0; i < n; i++) {
        prefix = (prefix + nums[i]) % p;
        
        const prefMod = prefix;
        let target = prefMod - need;
        if (target < 0) target += p;   // (prefMod - need + p) % p
        
        if (lastIndex.has(target)) {
            ans = Math.min(ans, i - lastIndex.get(target));
        }
        
        // Store latest index for this remainder
        lastIndex.set(prefMod, i);
    }
    
    return ans === n ? -1 : ans;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def minSubarray(self, nums: List[int], p: int) -> int:
        total = 0
        for x in nums:
            total = (total + x) % p  # total sum modulo p
        
        need = total
        if need == 0:
            return 0  # already divisible
        
        n = len(nums)
        last_index = {0: -1}  # remainder -> latest index
        ans = n
        prefix = 0
        
        for i, x in enumerate(nums):
            prefix = (prefix + x) % p
            # remainder we want to have seen before
            target = (prefix - need) % p  # Python handles negative mod
            
            if target in last_index:
                ans = min(ans, i - last_index[target])
            
            # update latest index of this remainder
            last_index[prefix] = i
        
        return -1 if ans == n else ans
```

---

### Go

```go
package main

func minSubarray(nums []int, p int) int {
    n := len(nums)
    total := 0
    for _, x := range nums {
        total = (total + x) % p
    }
    
    need := total
    if need == 0 {
        return 0
    }
    
    // map remainder -> latest index
    lastIndex := make(map[int]int, n*2)
    lastIndex[0] = -1
    
    ans := n
    prefix := 0
    
    for i, x := range nums {
        prefix = (prefix + x) % p
        prefMod := prefix
        
        target := prefMod - need
        if target < 0 {
            target += p
        }
        
        if j, ok := lastIndex[target]; ok {
            if i-j < ans {
                ans = i - j
            }
        }
        
        lastIndex[prefMod] = i
    }
    
    if ans == n {
        return -1
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The **core logic** is identical in all languages.
I’ll explain the main steps; you can map them to any of the codes above.

### 1. Compute total remainder

```python
total = 0
for x in nums:
    total = (total + x) % p
need = total
if need == 0:
    return 0
```

* I first compute the sum of all elements **modulo `p`**.
* `need` is the remainder I must “remove” from the array.
* If `need` is `0`, it means the total sum is already divisible by `p`, so I don’t need to remove anything.

**Same idea in other languages:**

* C++: `total = (total + x) % p;`
* Java: `total = (total + x) % p;`
* JS: `total = (total + x) % p;`
* Go: `total = (total + x) % p;`

---

### 2. Initialize hashmap / map

```python
n = len(nums)
last_index = {0: -1}
ans = n
prefix = 0
```

* I set `last_index[0] = -1` to represent an imaginary prefix *before* index 0 with sum 0.
* `ans` stores the minimum subarray length found. I start with `n` (worst case).
* `prefix` will store the running prefix sum modulo `p`.

**Other languages:**

* C++: `unordered_map<int, int> lastIndex; lastIndex[0] = -1;`
* Java: `Map<Integer, Integer> lastIndex = new HashMap<>(); lastIndex.put(0, -1);`
* JS: `const lastIndex = new Map(); lastIndex.set(0, -1);`
* Go: `lastIndex := make(map[int]int); lastIndex[0] = -1`

---

### 3. Iterate over the array and update prefix

```python
for i, x in enumerate(nums):
    prefix = (prefix + x) % p
```

* `prefix` now holds `(nums[0] + ... + nums[i]) % p`.

---

### 4. Compute target remainder

```python
    target = (prefix - need) % p
```

* This `target` is what the previous prefix remainder should be so that:

  * subarray `(j+1 ... i)` has remainder equal to `need`.

In languages that don’t handle negative modulo well, I manually adjust:

* C++/Java/JS/Go:

  ```cpp
  int target = prefMod - need;
  if (target < 0) target += p;
  ```

---

### 5. Check if target exists in the map

```python
    if target in last_index:
        ans = min(ans, i - last_index[target])
```

* If I’ve seen this remainder before at some index `j`,

  * subarray `(j+1 ... i)` is a candidate to remove.
  * length = `i - j`.
* I update the minimum answer.

Same in other languages with `if (lastIndex.count(target))`, `if (lastIndex.containsKey(target))`, `if (lastIndex.has(target))`, etc.

---

### 6. Update the map with current prefix

```python
    last_index[prefix] = i
```

* I store `i` as the latest index where `prefix` remainder occurs.
* I overwrite any older index for this remainder; the latest index usually leads to shorter subarrays.

---

### 7. Return final result

```python
return -1 if ans == n else ans
```

* If I never found a proper subarray (other than possibly whole array), `ans` stays `n`, and I return `-1`.
* Otherwise, I return the minimum length I found.

All other languages do the same logic with their own syntax.

---

## Examples

### Example 1

**Input:**

```text
nums = [3,1,4,2], p = 6
```

* Total sum = 3 + 1 + 4 + 2 = 10
* `10 % 6 = 4` → `need = 4`
* We need to remove a subarray with sum `% 6 = 4`.

Check subarrays:

* `[4]` has sum 4 → `4 % 6 = 4` → valid and length 1.
* This is the minimum possible.

**Output:**

```text
1
```

---

### Example 2

**Input:**

```text
nums = [6,3,5,2], p = 9
```

* Sum = 16, `16 % 9 = 7`, so `need = 7`.

The best subarray to remove is `[5,2]` (sum = 7), leaving `[6,3]` which sums to 9 (divisible by 9).

**Output:**

```text
2
```

---

### Example 3

**Input:**

```text
nums = [1,2,3], p = 3
```

* Sum = 6, `6 % 3 = 0`.
* Already divisible; no need to remove anything.

**Output:**

```text
0
```

---

## How to use / Run locally

### C++

1. Save the C++ code in a file, e.g., `solution.cpp`.
2. Add a `main()` function to read input if you’re testing manually, or just use it directly on LeetCode.
3. Compile:

   ```bash
   g++ -std=c++17 solution.cpp -o solution
   ./solution
   ```

### Java

1. Save code in `Solution.java`.
2. Use it directly in LeetCode or add your own main for testing.
3. Compile and run:

   ```bash
   javac Solution.java
   java Solution
   ```

### JavaScript (Node.js)

1. Save the JS function in a file, e.g., `solution.js`.
2. Export or call the function in a small test harness.
3. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save as `solution.py`.
2. Directly run:

   ```bash
   python3 solution.py
   ```

### Go

1. Save as `main.go`.
2. Run:

   ```bash
   go run main.go
   ```

For LeetCode, I simply paste the function body into the editor for the respective language; they handle input/output.

---

## Notes & Optimizations

* I always keep sums **modulo `p`** to:

  * avoid integer overflow,
  * and only work with remainders, which is all I actually need.
* Using `unordered_map` / `HashMap` / `Map` / `dict` / `map` gives average **O(1)** for insert + lookup.
* I always track the **latest index** for each remainder.
  This is safe because for a fixed current index `i`, using a larger `j` (latest) produces a **shorter subarray** `i - j`.
* Edge case `need == 0` is handled at the top to avoid extra work.
* Edge case where the only valid subarray is the entire array is handled by `ans == n` check.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
