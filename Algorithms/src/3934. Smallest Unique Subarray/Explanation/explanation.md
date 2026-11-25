# Smallest Unique Subarray - LeetCode 3934 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an integer array called `nums`.  

Your task is to find the shortest subarray that does not appear anywhere else in the array. In other words, find the minimum length of a subarray that is unique.

Two subarrays are considered the same only when they have the same length and the exact same values in the same order.

Return that minimum possible length.

This is a classic competitive programming problem that tests your understanding of binary search on the answer combined with rolling hash techniques for fast subarray comparison.

## Constraints

- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^5

The length of the array can be as large as 100,000, so any solution slower than O(n log n) will most likely get Time Limit Exceeded.

## Intuition

When I first read the problem I noticed two important things.

First, the answer is the smallest length L where at least one window of size L appears only once in the whole array.

Second, if a unique subarray of length L exists, then any longer subarray that fully contains it must also be unique. That means once a length works, every larger length also works.

This monotonic behavior is perfect for binary search. I can search for the smallest L instead of checking every possible length one by one.

The only remaining problem is how to check, for a fixed length L, whether any subarray of that length appears exactly once. Comparing every pair of windows directly would be too slow. Rolling hash solves that cleanly.

## Approach

I use binary search on the possible lengths from 1 to n.

For every candidate length mid I do the following:

1. Compute the rolling hash of every subarray of length mid.
2. Count how many times each hash appears using a hash map.
3. If any hash appears exactly once, then a unique subarray of this length exists.

Because of the monotonic property, if the check succeeds I try a smaller length. If it fails I try a larger length.

To make the hash calculation fast I precompute two independent prefix-hash arrays (two different bases and moduli). After that any window hash can be retrieved in constant time.

This combination of binary search + dual rolling hash keeps the solution efficient even for the maximum constraint of 10^5.

## Data Structures Used

- Two prefix-hash arrays and two power arrays  
  These let me calculate the hash of any subarray in O(1) time after O(n) preprocessing.

- Hash map (or dictionary)  
  I use it to count the frequency of each hash pair while checking a fixed length. This tells me instantly whether any window is unique.

- Binary search variables  
  Simple low, high and mid pointers to find the smallest valid length.

## Operations & Behavior Summary

1. Build two independent rolling-hash prefix arrays for the entire input.
2. Set the binary-search range from length 1 to n.
3. While the search range is valid:
   - Pick the middle length.
   - Slide a window of that length across the array.
   - Compute the dual hash of every window and store the counts.
   - If any count equals 1, record this length as a possible answer and search on the left side.
   - Otherwise search on the right side.
4. When the search finishes, the recorded answer is the minimum unique subarray length.

## Complexity

| Type              | Complexity     | Explanation |
|-------------------|----------------|-------------|
| Time Complexity   | O(n log n)     | Binary search runs log n iterations. Each iteration scans O(n) windows and does O(1) hash work per window. |
| Space Complexity  | O(n)           | We store four arrays of size n for the hashes and powers, plus a temporary map of size O(n). |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int smallestUniqueSubarray(vector<int>& nums) {
        int n = nums.size();
        const long long MOD1 = 1000000007LL, MOD2 = 1000000009LL;
        const long long BASE1 = 1000003LL, BASE2 = 1000033LL;
        vector<long long> h1(n + 1), h2(n + 1), p1(n + 1), p2(n + 1);
        p1[0] = p2[0] = 1;
        for (int i = 0; i < n; i++) {
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1;
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2;
            p1[i + 1] = p1[i] * BASE1 % MOD1;
            p2[i + 1] = p2[i] * BASE2 % MOD2;
        }
        auto hasUnique = [&](int len) -> bool {
            map<pair<long long, long long>, int> cnt;
            for (int i = 0; i + len <= n; i++) {
                long long a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
                long long b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
                cnt[{a, b}]++;
            }
            for (auto& e : cnt) if (e.second == 1) return true;
            return false;
        };
        int lo = 1, hi = n, ans = n;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (hasUnique(mid)) {
                ans = mid;
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int smallestUniqueSubarray(int[] nums) {
        int n = nums.length;
        final long MOD1 = 1000000007L, MOD2 = 1000000009L;
        final long BASE1 = 1000003L, BASE2 = 1000033L;
        long[] h1 = new long[n + 1], h2 = new long[n + 1];
        long[] p1 = new long[n + 1], p2 = new long[n + 1];
        p1[0] = p2[0] = 1;
        for (int i = 0; i < n; i++) {
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1;
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2;
            p1[i + 1] = p1[i] * BASE1 % MOD1;
            p2[i + 1] = p2[i] * BASE2 % MOD2;
        }
        int lo = 1, hi = n, ans = n;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (hasUnique(nums, mid, h1, h2, p1, p2, MOD1, MOD2)) {
                ans = mid;
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return ans;
    }
    private boolean hasUnique(int[] nums, int len, long[] h1, long[] h2, long[] p1, long[] p2, long MOD1, long MOD2) {
        Map<Long, Integer> cnt = new HashMap<>();
        int n = nums.length;
        for (int i = 0; i + len <= n; i++) {
            long a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
            long b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
            long key = a * MOD2 + b;
            cnt.put(key, cnt.getOrDefault(key, 0) + 1);
        }
        for (int v : cnt.values()) if (v == 1) return true;
        return false;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var smallestUniqueSubarray = function(nums) {
    const n = nums.length;
    const MOD1 = 1000000007n, MOD2 = 1000000009n;
    const BASE1 = 1000003n, BASE2 = 1000033n;
    const h1 = new Array(n + 1).fill(0n);
    const h2 = new Array(n + 1).fill(0n);
    const p1 = new Array(n + 1).fill(1n);
    const p2 = new Array(n + 1).fill(1n);
    for (let i = 0; i < n; i++) {
        h1[i + 1] = (h1[i] * BASE1 + BigInt(nums[i])) % MOD1;
        h2[i + 1] = (h2[i] * BASE2 + BigInt(nums[i])) % MOD2;
        p1[i + 1] = p1[i] * BASE1 % MOD1;
        p2[i + 1] = p2[i] * BASE2 % MOD2;
    }
    const hasUnique = (len) => {
        const cnt = new Map();
        for (let i = 0; i + len <= n; i++) {
            let a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
            let b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
            const key = a.toString() + "," + b.toString();
            cnt.set(key, (cnt.get(key) || 0) + 1);
        }
        for (const v of cnt.values()) if (v === 1) return true;
        return false;
    };
    let lo = 1, hi = n, ans = n;
    while (lo <= hi) {
        const mid = lo + ((hi - lo) >> 1);
        if (hasUnique(mid)) {
            ans = mid;
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function smallestUniqueSubarray(nums: number[]): number {
    const n = nums.length;
    const MOD1 = 1000000007n, MOD2 = 1000000009n;
    const BASE1 = 1000003n, BASE2 = 1000033n;
    const h1: bigint[] = new Array(n + 1).fill(0n);
    const h2: bigint[] = new Array(n + 1).fill(0n);
    const p1: bigint[] = new Array(n + 1).fill(1n);
    const p2: bigint[] = new Array(n + 1).fill(1n);
    for (let i = 0; i < n; i++) {
        h1[i + 1] = (h1[i] * BASE1 + BigInt(nums[i])) % MOD1;
        h2[i + 1] = (h2[i] * BASE2 + BigInt(nums[i])) % MOD2;
        p1[i + 1] = p1[i] * BASE1 % MOD1;
        p2[i + 1] = p2[i] * BASE2 % MOD2;
    }
    const hasUnique = (len: number): boolean => {
        const cnt = new Map<string, number>();
        for (let i = 0; i + len <= n; i++) {
            let a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
            let b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
            const key = a.toString() + "," + b.toString();
            cnt.set(key, (cnt.get(key) || 0) + 1);
        }
        for (const v of cnt.values()) if (v === 1) return true;
        return false;
    };
    let lo = 1, hi = n, ans = n;
    while (lo <= hi) {
        const mid = lo + ((hi - lo) >> 1);
        if (hasUnique(mid)) {
            ans = mid;
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def smallestUniqueSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        MOD1, MOD2 = 10**9 + 7, 10**9 + 9
        BASE1, BASE2 = 1000003, 1000033
        h1 = [0] * (n + 1)
        h2 = [0] * (n + 1)
        p1 = [1] * (n + 1)
        p2 = [1] * (n + 1)
        for i in range(n):
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2
            p1[i + 1] = p1[i] * BASE1 % MOD1
            p2[i + 1] = p2[i] * BASE2 % MOD2
        def has_unique(length):
            cnt = {}
            for i in range(n - length + 1):
                a = (h1[i + length] - h1[i] * p1[length] % MOD1 + MOD1) % MOD1
                b = (h2[i + length] - h2[i] * p2[length] % MOD2 + MOD2) % MOD2
                key = (a, b)
                cnt[key] = cnt.get(key, 0) + 1
            return any(v == 1 for v in cnt.values())
        lo, hi, ans = 1, n, n
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if has_unique(mid):
                ans = mid
                hi = mid - 1
            else:
                lo = mid + 1
        return ans
```

### Go
```go
func smallestUniqueSubarray(nums []int) int {
    n := len(nums)
    const MOD1, MOD2 = 1000000007, 1000000009
    const BASE1, BASE2 = 1000003, 1000033
    h1 := make([]int64, n+1)
    h2 := make([]int64, n+1)
    p1 := make([]int64, n+1)
    p2 := make([]int64, n+1)
    p1[0], p2[0] = 1, 1
    for i := 0; i < n; i++ {
        h1[i+1] = (h1[i]*BASE1 + int64(nums[i])) % MOD1
        h2[i+1] = (h2[i]*BASE2 + int64(nums[i])) % MOD2
        p1[i+1] = p1[i] * BASE1 % MOD1
        p2[i+1] = p2[i] * BASE2 % MOD2
    }
    hasUnique := func(length int) bool {
        cnt := make(map[[2]int64]int)
        for i := 0; i+length <= n; i++ {
            a := (h1[i+length] - h1[i]*p1[length]%MOD1 + MOD1) % MOD1
            b := (h2[i+length] - h2[i]*p2[length]%MOD2 + MOD2) % MOD2
            cnt[[2]int64{a, b}]++
        }
        for _, v := range cnt {
            if v == 1 {
                return true
            }
        }
        return false
    }
    lo, hi, ans := 1, n, n
    for lo <= hi {
        mid := lo + (hi-lo)/2
        if hasUnique(mid) {
            ans = mid
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six language versions follow the same high-level logic. The only differences are syntax and how maps or dictionaries are written.

I first allocate four arrays: two for the prefix hashes and two for the powers of the bases. I fill them in a single left-to-right pass. This preprocessing costs O(n) and is done only once.

Inside the binary-search loop I never recompute the whole prefix. I just use the already-built arrays to get any window hash with a few arithmetic operations.

For the frequency map:
- In C++ I use `map<pair<long long, long long>, int>` or an unordered map with a custom hash.
- In Java I encode the two 64-bit hashes into one long key.
- In JavaScript / TypeScript I convert the two BigInt hashes into a string key.
- In Python I use a normal dictionary with a tuple key.
- In Go I use a map whose key is a fixed-size array of two int64 values.

After counting, I simply scan the map values. The first time I see a frequency of 1 I know the current length is valid and I can safely try a smaller length.

Edge cases are handled automatically:
- When the array length is 1 the answer is always 1.
- When every element is identical the answer becomes n (the whole array).
- The whole array is always unique, so the binary search is guaranteed to find an answer.

## Examples

**Example 1**  
Input: `nums = [3, 3, 3]`  
Output: `3`  

Length 1 and length 2 windows all repeat. Only the full array of length 3 appears once, so the answer is 3.

**Example 2**  
Input: `nums = [2, 1, 2, 3, 3]`  
Output: `1`  

The single element 1 appears only once. Therefore the minimum unique subarray length is 1.

**Example 3**  
Input: `nums = [1, 1, 2, 2, 1]`  
Output: `2`  

Every length-1 window repeats. Several length-2 windows appear only once (for example [1, 2] and [2, 1]), so the answer is 2.

## How to Use / Run Locally

1. Copy the code of the language you want into a file:
   - C++ → `main.cpp`
   - Java → `Solution.java`
   - JavaScript → `solution.js`
   - TypeScript → `solution.ts`
   - Python → `solution.py`
   - Go → `main.go`

2. Compile / run commands (assuming the function is called from a simple driver):

```bash
# C++
g++ -std=c++17 main.cpp -o main && ./main

# Java
javac Solution.java && java Solution

# JavaScript
node solution.js

# TypeScript
tsc solution.ts && node solution.js

# Python
python3 solution.py

# Go
go run main.go
```

Feel free to wrap the function with a small main method that reads input and prints the result so you can test the examples yourself.

## Notes & Optimizations

- Dual rolling hash (two bases + two moduli) almost eliminates the chance of collisions for the given constraints.
- Using a single 64-bit hash is faster but carries a small theoretical risk of wrong answers. Dual hash is safer for a hard problem.
- The map can be replaced by an unordered map in C++ for a small constant-factor speed-up, but the ordered map is already fast enough.
- An alternative O(n) approach using suffix arrays or Z-algorithm exists, but it is more complex to implement and not necessary under the current constraints.
- Always remember that the whole array is unique, so the binary-search upper bound of n is safe.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
