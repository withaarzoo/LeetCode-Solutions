# 3318. Find X-Sum of All K-Long Subarrays I

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

You are given an integer array `nums`, and integers `k` and `x`.
For each contiguous subarray of length `k`, compute its **x-sum**:

1. Count occurrences (frequency) of each distinct value in that subarray.
2. Keep only the **top `x` most frequent values**.

   * If two values have the same frequency, the **larger value** is considered more frequent (i.e., breaks ties).
3. The x-sum is the sum of all elements whose values are in that kept set (equivalently, `sum(value * freq)` for those `x` values).

Return an array `answer` where `answer[i]` is the x-sum of `nums[i .. i+k-1]`.

---

## Constraints

*(LeetCode “I” version — small inputs, friendly to per-window sorting.)*

* `1 ≤ n = nums.length ≤ 50`
* `1 ≤ nums[i] ≤ 50`
* `1 ≤ x ≤ k ≤ n`

---

## Intuition

I thought: instead of summing elements one by one, I can count how many times each value appears in the window.
If I sort values by **(frequency desc, value desc)**, the first `x` items are exactly the ones I should keep.
Then the window’s x-sum is just `Σ value * freq` over those `x` items.
While I slide the window, I only need to update the counts for the element entering and the element leaving.

---

## Approach

1. Use a **sliding window** of size `k`.
2. Maintain a frequency map `freq[value]` for the current window.
3. For each window:

   * Convert `freq` into a list of `(value, freq)` pairs.
   * Sort by `freq` descending, then `value` descending.
   * Take the first `x` pairs and add `value * freq` to form the x-sum.
   * Store it in the answer.
4. Slide:

   * Add `nums[i]` (right end) to `freq`.
   * Remove `nums[i-k]` (left end) from `freq` and delete if its count becomes zero.
5. Repeat for all windows.

This is clean and completely fine for the “I” version constraints.

---

## Data Structures Used

* **Hash map / dictionary**: to keep frequencies within the current window.
* **Array/list of pairs**: to sort `(value, freq)` per window.

---

## Operations & Behavior Summary

* **Add right element**: `freq[add]++`
* **Remove left element**: `freq[rem]--` and delete if it becomes `0`
* **Rank values**: sort `(value, freq)` by `(freq desc, value desc)`
* **Compute x-sum**: sum of `value * freq` for the top `x` pairs

---

## Complexity

* **Time Complexity:** `O(n * m log m)`

  * `n` = length of `nums`
  * `m` = number of **distinct** values in the current window (≤ 50 here).
    For each of the ~`n` windows we sort up to `m` items.
* **Space Complexity:** `O(m)`
  Only the frequency map of the current window plus a temporary list for sorting.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    vector<int> findXSum(vector<int>& nums, int k, int x) {
        int n = nums.size();
        vector<int> ans;
        if (n == 0 || k == 0) return ans;
        ans.reserve(n - k + 1);

        unordered_map<int,int> freq;

        // Build initial window
        for (int i = 0; i < k; ++i) freq[nums[i]]++;

        auto compute = [&](unordered_map<int,int>& f)->int {
            // Collect (value, freq)
            vector<pair<int,int>> items;
            items.reserve(f.size());
            for (auto &p : f) items.push_back({p.first, p.second});
            // Sort by freq desc, value desc
            sort(items.begin(), items.end(), [](const auto& a, const auto& b){
                if (a.second != b.second) return a.second > b.second;
                return a.first > b.first;
            });
            long long sum = 0;
            int take = min<int>(x, (int)items.size());
            for (int i = 0; i < take; ++i) sum += 1LL * items[i].first * items[i].second;
            return (int)sum;
        };

        ans.push_back(compute(freq));

        // Slide the window
        for (int i = k; i < n; ++i) {
            int add = nums[i];
            int rem = nums[i - k];

            freq[add]++;
            if (--freq[rem] == 0) freq.erase(rem);

            ans.push_back(compute(freq));
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int[] findXSum(int[] nums, int k, int x) {
        int n = nums.length;
        if (n == 0 || k == 0) return new int[0];
        int[] ans = new int[n - k + 1];
        Map<Integer, Integer> freq = new HashMap<>();

        // initial window
        for (int i = 0; i < k; i++) {
            freq.put(nums[i], freq.getOrDefault(nums[i], 0) + 1);
        }

        ans[0] = compute(freq, x);

        // slide
        for (int i = k; i < n; i++) {
            int add = nums[i], rem = nums[i - k];

            freq.put(add, freq.getOrDefault(add, 0) + 1);
            int fr = freq.get(rem) - 1;
            if (fr == 0) freq.remove(rem);
            else freq.put(rem, fr);

            ans[i - k + 1] = compute(freq, x);
        }
        return ans;
    }

    private int compute(Map<Integer, Integer> freq, int x) {
        List<int[]> items = new ArrayList<>();
        for (Map.Entry<Integer, Integer> e : freq.entrySet()) {
            items.add(new int[]{e.getKey(), e.getValue()});
        }
        items.sort((a, b) -> {
            if (a[1] != b[1]) return b[1] - a[1]; // freq desc
            return b[0] - a[0];                    // value desc
        });
        long sum = 0;
        int take = Math.min(x, items.size());
        for (int i = 0; i < take; i++) sum += 1L * items.get(i)[0] * items.get(i)[1];
        return (int)sum;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} x
 * @return {number[]}
 */
var findXSum = function(nums, k, x) {
  const n = nums.length;
  if (n === 0 || k === 0) return [];
  const ans = [];
  const freq = new Map();

  // initial window
  for (let i = 0; i < k; i++) {
    freq.set(nums[i], (freq.get(nums[i]) || 0) + 1);
  }
  ans.push(compute(freq, x));

  // slide
  for (let i = k; i < n; i++) {
    const add = nums[i], rem = nums[i - k];

    freq.set(add, (freq.get(add) || 0) + 1);
    const fr = (freq.get(rem) || 0) - 1;
    if (fr === 0) freq.delete(rem);
    else freq.set(rem, fr);

    ans.push(compute(freq, x));
  }
  return ans;

  function compute(map, x) {
    const items = [];
    for (const [v, f] of map.entries()) items.push([v, f]);
    items.sort((a, b) => (a[1] !== b[1] ? b[1] - a[1] : b[0] - a[0]));
    let sum = 0;
    for (let i = 0; i < Math.min(x, items.length); i++) {
      sum += items[i][0] * items[i][1];
    }
    return sum;
  }
};
```

### Python3

```python
from collections import defaultdict
from typing import List

class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        n = len(nums)
        if n == 0 or k == 0:
            return []
        freq = defaultdict(int)

        # initial window
        for i in range(k):
            freq[nums[i]] += 1

        def compute(freq, x):
            items = [(v, f) for v, f in freq.items()]
            items.sort(key=lambda t: (-t[1], -t[0]))  # freq desc, value desc
            total = 0
            for i in range(min(x, len(items))):
                v, f = items[i]
                total += v * f
            return total

        ans = [compute(freq, x)]

        # slide window
        for i in range(k, n):
            add = nums[i]
            rem = nums[i - k]
            freq[add] += 1
            freq[rem] -= 1
            if freq[rem] == 0:
                del freq[rem]
            ans.append(compute(freq, x))

        return ans
```

### Go

```go
package main

import "sort"

func findXSum(nums []int, k int, x int) []int {
 n := len(nums)
 if n == 0 || k == 0 {
  return []int{}
 }
 ans := make([]int, 0, n-k+1)
 freq := map[int]int{}

 // initial window
 for i := 0; i < k; i++ {
  freq[nums[i]]++
 }
 ans = append(ans, computeX(freq, x))

 // slide
 for i := k; i < n; i++ {
  add := nums[i]
  rem := nums[i-k]
  freq[add]++
  freq[rem]--
  if freq[rem] == 0 {
   delete(freq, rem)
  }
  ans = append(ans, computeX(freq, x))
 }
 return ans
}

type pair struct{ val, cnt int }

func computeX(freq map[int]int, x int) int {
 items := make([]pair, 0, len(freq))
 for v, c := range freq {
  items = append(items, pair{v, c})
 }
 sort.Slice(items, func(i, j int) bool {
  if items[i].cnt != items[j].cnt {
   return items[i].cnt > items[j].cnt // freq desc
  }
  return items[i].val > items[j].val   // value desc
 })
 sum := 0
 take := x
 if take > len(items) {
  take = len(items)
 }
 for i := 0; i < take; i++ {
  sum += items[i].val * items[i].cnt
 }
 return sum
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. **Keep frequencies while sliding**

   * Start with the first `k` elements and count them in a map/dictionary.
     Example (Python): `freq[nums[i]] += 1`.
2. **Compute x-sum for this window**

   * Convert the map into a list of `(value, freq)` pairs.
   * Sort by **frequency descending**, then by **value descending**.
     This exactly follows the rule that larger values break ties in frequency.
   * Take the first `x` pairs and add `value * freq` to a running total.
3. **Slide the window by 1**

   * **Add** the element coming in on the right: increment its frequency.
   * **Remove** the element leaving on the left: decrement its frequency and delete if it becomes zero (keeps the map clean).
4. **Repeat**

   * For each new window, recompute the x-sum with the up-to-date frequency map and push it to the result array.

Because constraints are small in this “I” version, recomputing via sorting per window is simple and fast, and the code stays clean.

---

## Examples

**Example 1**

```
Input: nums = [1,1,2,2,3,4,2,3], k = 3, x = 2
Windows:
[1,1,2] -> counts: {1:2, 2:1} -> top2: (1,2),(2,1) -> x-sum = 1*2 + 2*1 = 4
[1,2,2] -> {1:1, 2:2} -> x-sum = 1*1 + 2*2 = 5
[2,2,3] -> {2:2, 3:1} -> x-sum = 2*2 + 3*1 = 7
...
Output: [4,5,7,...]
```

**Example 2 (from problem statement)**

```
nums = [3,8,7,8,7,5], k = 2, x = 2
For every window, x = k so we sum the whole window:
[3,8] -> 11
[8,7] -> 15
[7,8] -> 15
[8,7] -> 15
[7,5] -> 12
Output: [11,15,15,15,12]
```

---

## How to use / Run locally

**C++**

```bash
g++ -std=c++17 -O2 main.cpp -o main && ./main
```

**Java**

```bash
javac Solution.java && java Solution
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

> Each `main.*` should construct `nums`, `k`, `x`, call the function, and print the resulting array for your tests.

---

## Notes & Optimizations

* For the **“I”** version, values and `n` are small ⇒ sorting per window is perfectly fine.
* If you want to push further (for larger constraints), you could maintain two balanced structures:
  one that keeps the **top-x** `(freq, value)` pairs and another for the **rest**, updating them as counts change. That yields near `O(n log m)` total with fewer full sorts. This is more complex and not necessary here.
* Deleting zero-count entries keeps the frequency map compact and speeds up sorting.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
