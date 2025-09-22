# 3005. Count Elements With Maximum Frequency

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

You are given an array `nums` consisting of positive integers.
Return the **total number of elements** in `nums` such that those elements all have the **maximum frequency**.

* The **frequency** of an element is the number of occurrences of that element in the array.
* In short: count each value's frequency, find the maximum frequency `maxFreq`, then return the sum of occurrences of every element whose frequency is `maxFreq`.

Example:
`nums = [1,2,2,3,1,4]` → `1` occurs `2` times, `2` occurs `2` times, others occur `1`. `maxFreq = 2` → elements with frequency 2 are `1` and `2` → total occurrences = `2 + 2 = 4`.

---

## Constraints

* `1 <= nums.length <= 100`
* `1 <= nums[i] <= 100`

> Note: Given these constraints, the number of distinct elements `k` is at most `100`. For larger constraints, the same algorithm still applies.

---

## Intuition

I thought about counting how many times each number appears in the array.
After I have those counts, I only care about the largest count (the maximum frequency).
Then I add together the counts of every number whose frequency equals that maximum. That gives the total number of elements that appear with maximum frequency.

---

## Approach

1. Scan the array once and build a frequency table (value → count).
2. Scan the frequency table to find the maximum frequency `maxFreq`.
3. Sum all counts from the frequency table that are equal to `maxFreq`.
4. Return that sum.

This is straightforward and efficient: one pass to count and then a couple of short scans over the distinct values.

---

## Data Structures Used

* Hash map / dictionary (`unordered_map` in C++, `HashMap` in Java, `Map` in JavaScript, `Counter`/`dict` in Python, `map` in Go) to store frequency of each distinct value.
* Alternative (because `1 <= nums[i] <= 100`): a fixed-size array of length `101` for counts which gives O(1) space relative to value range.

---

## Operations & Behavior Summary

* Build frequency map: O(n) operations (increment counts).
* Determine `maxFreq`: scan the frequency map — O(k) where `k` is number of distinct values.
* Sum counts equal to `maxFreq`: another O(k) scan.
* Return the final sum.

Behavior: stable and deterministic. Works for arrays with all unique values (answer = `n`) or all equal values (answer = `n`).

---

## Complexity

* **Time Complexity:** `O(n)` — `n` is `nums.length`. We scan `nums` once and then inspect the frequency map (size `k ≤ n`).
* **Space Complexity:** `O(k)` — `k` is the number of distinct elements (or `O(1)` if we use a fixed-range array due to value limits).

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxFrequencyElements(vector<int>& nums) {
        unordered_map<int,int> freq;
        // Build frequency map
        for (int x : nums) freq[x]++;

        // Find maximum frequency
        int maxFreq = 0;
        for (auto &p : freq) maxFreq = max(maxFreq, p.second);

        // Sum frequencies of elements that have frequency == maxFreq
        int result = 0;
        for (auto &p : freq) if (p.second == maxFreq) result += p.second;
        return result;
    }
};

// Simple main to run locally
int main() {
    vector<int> nums = {1,2,2,3,1,4};
    Solution s;
    cout << s.maxFrequencyElements(nums) << "\n"; // Expected output: 4
    return 0;
}
```

---

### Java

```java
import java.util.HashMap;
import java.util.Map;

class Solution {
    public int maxFrequencyElements(int[] nums) {
        Map<Integer, Integer> freq = new HashMap<>();
        // Build frequency map
        for (int x : nums) freq.put(x, freq.getOrDefault(x, 0) + 1);

        // Find maximum frequency
        int maxFreq = 0;
        for (int v : freq.values()) maxFreq = Math.max(maxFreq, v);

        // Sum frequencies of elements that have frequency == maxFreq
        int result = 0;
        for (int v : freq.values()) if (v == maxFreq) result += v;
        return result;
    }
}

// Local runner
public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int[] nums = {1,2,2,3,1,4};
        System.out.println(sol.maxFrequencyElements(nums)); // Expected: 4
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxFrequencyElements = function(nums) {
    const freq = new Map();
    // Build frequency map
    for (const x of nums) freq.set(x, (freq.get(x) || 0) + 1);

    // Find max frequency
    let maxFreq = 0;
    for (const v of freq.values()) if (v > maxFreq) maxFreq = v;

    // Sum counts of elements whose frequency == maxFreq
    let result = 0;
    for (const v of freq.values()) if (v === maxFreq) result += v;
    return result;
};

// Local test
console.log(maxFrequencyElements([1,2,2,3,1,4])); // 4
```

---

### Python3

```python
from collections import Counter
from typing import List

class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        # Count frequencies of all numbers
        cnt = Counter(nums)

        # Find the maximum frequency value
        max_freq = max(cnt.values())

        # Sum frequencies of numbers that have frequency == max_freq
        return sum(v for v in cnt.values() if v == max_freq)

# Local test
if __name__ == "__main__":
    print(Solution().maxFrequencyElements([1,2,2,3,1,4]))  # 4
```

---

### Go

```go
package main

import (
    "fmt"
)

func maxFrequencyElements(nums []int) int {
    freq := make(map[int]int)
    // Build frequency map
    for _, x := range nums {
        freq[x]++
    }

    // Find maximum frequency
    maxFreq := 0
    for _, v := range freq {
        if v > maxFreq {
            maxFreq = v
        }
    }

    // Sum frequencies of those elements that have frequency == maxFreq
    result := 0
    for _, v := range freq {
        if v == maxFreq {
            result += v
        }
    }
    return result
}

func main() {
    nums := []int{1,2,2,3,1,4}
    fmt.Println(maxFrequencyElements(nums)) // 4
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the algorithm once and show how each language maps to it. The logic is the same in all languages.

### Key steps (common)

1. **Build frequency map:** For every value `x` in `nums`, increment `freq[x]` by 1.
2. **Find maximum frequency:** Scan `freq` and record the maximum value `maxFreq`.
3. **Sum matching frequencies:** Scan `freq` again and sum every `count` that equals `maxFreq`. Return that sum.

---

### Explanation mapped to C++

```c++
// unordered_map<int,int> freq;
// for (int x : nums) freq[x]++;
```

* `freq[x]++` increases the count for element `x`. If `x` is not present, `operator[]` default-constructs it to `0`.

```c++
int maxFreq = 0;
for (auto &p : freq) maxFreq = max(maxFreq, p.second);
```

* Iterate over key-value pairs, `p.second` is the count. Keep the largest.

```c++
int result = 0;
for (auto &p : freq) if (p.second == maxFreq) result += p.second;
```

* If a value's count equals `maxFreq`, add that count to the final result.

---

### Explanation mapped to Java

```java
Map<Integer, Integer> freq = new HashMap<>();
for (int x : nums) freq.put(x, freq.getOrDefault(x, 0) + 1);
```

* `getOrDefault` is used to safely increment a possibly-missing key.

```java
int maxFreq = 0;
for (int v : freq.values()) maxFreq = Math.max(maxFreq, v);
```

* Find the maximum count across `freq.values()`.

```java
int result = 0;
for (int v : freq.values()) if (v == maxFreq) result += v;
```

* Sum all frequencies equal to the `maxFreq`.

---

### Explanation mapped to JavaScript

```javascript
const freq = new Map();
for (const x of nums) freq.set(x, (freq.get(x) || 0) + 1);
```

* `Map` stores counts. `freq.get(x) || 0` handles missing keys.

```javascript
let maxFreq = 0;
for (const v of freq.values()) if (v > maxFreq) maxFreq = v;
```

* Find largest frequency.

```javascript
let result = 0;
for (const v of freq.values()) if (v === maxFreq) result += v;
```

* Sum counts equal to `maxFreq`.

---

### Explanation mapped to Python

```python
from collections import Counter
cnt = Counter(nums)
```

* `Counter` returns dict-like counts.

```python
max_freq = max(cnt.values())
return sum(v for v in cnt.values() if v == max_freq)
```

* `max` finds the largest count; `sum(...)` adds all counts equal to it.

---

### Explanation mapped to Go

```go
freq := make(map[int]int)
for _, x := range nums { freq[x]++ }
```

* Build frequency map.

```go
maxFreq := 0
for _, v := range freq { if v > maxFreq { maxFreq = v } }
```

* Find maximum.

```go
result := 0
for _, v := range freq { if v == maxFreq { result += v } }
```

* Sum matching counts.

---

## Examples

1. Input: `nums = [1,2,2,3,1,4]`

   * Frequencies: `1->2, 2->2, 3->1, 4->1`
   * `maxFreq = 2` → elements `1` and `2` → answer `2 + 2 = 4`
   * Output: `4`

2. Input: `nums = [1,2,3,4,5]`

   * All frequencies are `1` → `maxFreq = 1` → sum = `1+1+1+1+1 = 5`
   * Output: `5`

3. Edge: `nums = [7]` → Output: `1`

4. Edge: `nums = [2,2,2,2]` → Output: `4`

---

## How to use / Run locally

### C++

* Save as `solution.cpp`. Build and run:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

* Save as `Main.java` (contains `Main` class). Compile and run:

```bash
javac Main.java
java Main
```

### JavaScript (Node.js)

* Save as `solution.js`. Run:

```bash
node solution.js
```

### Python 3

* Save as `solution.py`. Run:

```bash
python3 solution.py
```

### Go

* Save as `solution.go`. Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* Because `1 <= nums[i] <= 100`, you can use an integer array of length `101` instead of a hash map to store counts. That reduces overhead and gives O(1) extra space relative to value range:

  * `int cnt[101] = {0}; for (x : nums) cnt[x]++;`
  * Then find `maxFreq` by scanning `cnt`.
* If constraints were large (e.g., values up to `1e9`), the hash map approach is the correct general approach.
* The algorithm is already optimal in time for this problem (you must examine all elements at least once).
* This solution is simple, readable, and efficient for interview and contest settings.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
