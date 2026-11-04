# Minimum Operations to Make K Non-Overlapping Subarrays Equal - LeetCode 3505 Solution

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

You are given an integer array `nums` and two integers `x` and `k`.

Your task is to choose exactly `k` non-overlapping subarrays, each of length exactly `x`. For every chosen subarray you must make all of its elements equal by repeatedly adding or subtracting 1 from any element. Each such change costs 1.

Return the minimum total cost required to do this.

In other words: pick `k` windows of size `x` that never overlap, turn every window into a constant array (by moving numbers toward a common value), and keep the sum of all those movement costs as small as possible.

## Constraints

- `1 <= n <= 10^5` where `n = nums.length`
- `1 <= x <= n`
- `1 <= k <= 15` and `k * x <= n`
- `-10^6 <= nums[i] <= 10^6`

These limits force us to compute every window cost in roughly logarithmic time and then run a linear DP with a small factor of `k`.

## Intuition

The first useful observation is that the cheapest way to make a whole window equal is to move every element to the median of that window. Moving everything to any other target costs at least as much.

Once we know the cost of every possible window of length `x`, the original problem reduces to a classic selection task: pick `k` of those windows so that they never overlap and the sum of their costs is minimized.

Because `k` is tiny (at most 15) a simple dynamic-programming solution works: walk through the array from left to right and decide, at every position, whether to finish a new window there or keep going.

The only remaining difficulty is computing the median-based cost of every window fast enough. A plain sort for each window would be far too slow, so we need a data structure that can add a number, remove a number, find the median, and obtain the sum of everything below the median in logarithmic time.

## Approach

1. Compress all values in `nums` to ranks from 1 to `m` (the number of distinct numbers). This lets us index Fenwick trees with small integers.

2. Maintain two Fenwick trees while sliding a window of length `x` across the array:
   - one tree stores the count of each rank currently inside the window,
   - the other stores the sum of the actual values of those ranks.

3. After each slide, walk down the count tree to locate the median rank and, at the same time, accumulate the count and the sum of every element strictly smaller than the median. From those three numbers (median, count-below, sum-below) and the total sum of the window we can compute the cost in constant time.

4. Store every window cost in an array of size `n - x + 1`.

5. Run a classic DP:
   - `dp[j][i]` = minimum cost to place exactly `j` non-overlapping windows inside the first `i` elements.
   - Transition: either skip the current position, or end a window that covers positions `[i-x, i-1]`.
   - Because each row depends only on the previous row we keep only two arrays and swap them after every `j`.

6. After `k` rounds the answer sits at `dp[k][n]`.

## Data Structures Used

- **Value compression / ranking**  
  Turns arbitrary integers into dense ranks so Fenwick trees stay small and fast.

- **Two Fenwick trees (Binary Indexed Trees)**  
  One for frequencies, one for value sums. Together they support insert, delete, order-statistic queries and prefix-sum queries in `O(log m)` time.

- **Rolling DP arrays**  
  Two arrays of size `n+1` that store the previous and current DP rows, saving a factor of `k` in memory.

## Operations & Behavior Summary

- Sort and unique the array values, then map every `nums[i]` to its rank.
- Slide a window of length `x`:
  - insert the new element into both Fenwick trees and into the running total,
  - remove the element that just left the window,
  - locate the median by binary-lifting on the frequency tree while collecting the sum of smaller elements,
  - compute the absolute-difference cost and store it.
- Initialise a DP array of zeros (zero windows cost nothing).
- For each of the `k` windows:
  - for every possible ending position `i`, take the minimum of “do not end a window here” and “end a window of length `x` here”.
- Return the final entry of the last DP row.

## Complexity

| Type              | Complexity          | Explanation |
|-------------------|---------------------|-------------|
| Time Complexity   | `O(n log n + n·k)`  | Sorting for compression is `O(n log n)`. Each of the `n` Fenwick updates/queries costs `O(log n)`. The DP fills `k` rows of length `n`. |
| Space Complexity  | `O(n)`              | Rank array, two Fenwick trees, cost array and two DP rows are all linear in `n`. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long minOperations(vector<int>& nums, int x, int k) {
        int n = nums.size();
        vector<int> vals(nums);
        sort(vals.begin(), vals.end());
        vals.erase(unique(vals.begin(), vals.end()), vals.end());
        int m = vals.size();
        vector<int> rk(n);
        for (int i = 0; i < n; i++) {
            rk[i] = lower_bound(vals.begin(), vals.end(), nums[i]) - vals.begin() + 1;
        }
        vector<int> cnt(m + 1, 0);
        vector<long long> sm(m + 1, 0);
        int top = 1;
        while (top * 2 <= m) top *= 2;
        int need = (x + 1) / 2;
        int w = n - x + 1;
        vector<long long> cost(w);
        long long total = 0;
        for (int i = 0; i < n; i++) {
            total += nums[i];
            for (int p = rk[i]; p <= m; p += p & -p) {
                cnt[p] += 1;
                sm[p] += nums[i];
            }
            if (i >= x) {
                total -= nums[i - x];
                for (int p = rk[i - x]; p <= m; p += p & -p) {
                    cnt[p] -= 1;
                    sm[p] -= nums[i - x];
                }
            }
            if (i >= x - 1) {
                int pos = 0, rem = need;
                long long s0 = 0;
                for (int pw = top; pw > 0; pw >>= 1) {
                    int nxt = pos + pw;
                    if (nxt <= m && cnt[nxt] < rem) {
                        pos = nxt;
                        rem -= cnt[nxt];
                        s0 += sm[nxt];
                    }
                }
                long long med = vals[pos];
                long long c0 = need - rem;
                cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0);
            }
        }
        const long long INF = 1000000000000000000LL;
        vector<long long> prev(n + 1, 0), cur(n + 1, 0);
        for (int j = 0; j < k; j++) {
            for (int i = 0; i < x; i++) cur[i] = INF;
            for (int i = x; i <= n; i++) {
                cur[i] = min(cur[i - 1], prev[i - x] + cost[i - x]);
            }
            swap(prev, cur);
        }
        return prev[n];
    }
};
```

### Java

```java
class Solution {
    public long minOperations(int[] nums, int x, int k) {
        int n = nums.length;
        int[] vals = nums.clone();
        Arrays.sort(vals);
        int m = 0;
        for (int i = 0; i < n; i++) {
            if (m == 0 || vals[i] != vals[m - 1]) {
                vals[m++] = vals[i];
            }
        }
        int[] rk = new int[n];
        for (int i = 0; i < n; i++) {
            rk[i] = Arrays.binarySearch(vals, 0, m, nums[i]) + 1;
        }
        int[] cnt = new int[m + 1];
        long[] sm = new long[m + 1];
        int top = 1;
        while (top * 2 <= m) top *= 2;
        int need = (x + 1) / 2;
        int w = n - x + 1;
        long[] cost = new long[w];
        long total = 0;
        for (int i = 0; i < n; i++) {
            total += nums[i];
            for (int p = rk[i]; p <= m; p += p & -p) {
                cnt[p] += 1;
                sm[p] += nums[i];
            }
            if (i >= x) {
                total -= nums[i - x];
                for (int p = rk[i - x]; p <= m; p += p & -p) {
                    cnt[p] -= 1;
                    sm[p] -= nums[i - x];
                }
            }
            if (i >= x - 1) {
                int pos = 0, rem = need;
                long s0 = 0;
                for (int pw = top; pw > 0; pw >>= 1) {
                    int nxt = pos + pw;
                    if (nxt <= m && cnt[nxt] < rem) {
                        pos = nxt;
                        rem -= cnt[nxt];
                        s0 += sm[nxt];
                    }
                }
                long med = vals[pos];
                long c0 = need - rem;
                cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0);
            }
        }
        long INF = 1000000000000000000L;
        long[] prev = new long[n + 1];
        long[] cur = new long[n + 1];
        for (int j = 0; j < k; j++) {
            for (int i = 0; i < x; i++) cur[i] = INF;
            for (int i = x; i <= n; i++) {
                long a = cur[i - 1];
                long b = prev[i - x] + cost[i - x];
                cur[i] = a < b ? a : b;
            }
            long[] t = prev;
            prev = cur;
            cur = t;
        }
        return prev[n];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} x
 * @param {number} k
 * @return {number}
 */
var minOperations = function(nums, x, k) {
    const n = nums.length;
    const sorted = Int32Array.from(nums).sort();
    let m = 0;
    for (let i = 0; i < n; i++) {
        if (m === 0 || sorted[i] !== sorted[m - 1]) {
            sorted[m++] = sorted[i];
        }
    }
    const vals = sorted;
    const rk = new Int32Array(n);
    for (let i = 0; i < n; i++) {
        let lo = 0, hi = m - 1;
        while (lo < hi) {
            const mid = (lo + hi) >> 1;
            if (vals[mid] < nums[i]) lo = mid + 1;
            else hi = mid;
        }
        rk[i] = lo + 1;
    }
    const cnt = new Int32Array(m + 1);
    const sm = new Float64Array(m + 1);
    let top = 1;
    while (top * 2 <= m) top *= 2;
    const need = (x + 1) >> 1;
    const w = n - x + 1;
    const cost = new Float64Array(w);
    let total = 0;
    for (let i = 0; i < n; i++) {
        total += nums[i];
        for (let p = rk[i]; p <= m; p += p & -p) {
            cnt[p] += 1;
            sm[p] += nums[i];
        }
        if (i >= x) {
            total -= nums[i - x];
            for (let p = rk[i - x]; p <= m; p += p & -p) {
                cnt[p] -= 1;
                sm[p] -= nums[i - x];
            }
        }
        if (i >= x - 1) {
            let pos = 0, rem = need, s0 = 0;
            for (let pw = top; pw > 0; pw >>= 1) {
                const nxt = pos + pw;
                if (nxt <= m && cnt[nxt] < rem) {
                    pos = nxt;
                    rem -= cnt[nxt];
                    s0 += sm[nxt];
                }
            }
            const med = vals[pos];
            const c0 = need - rem;
            cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0);
        }
    }
    let prev = new Float64Array(n + 1);
    let cur = new Float64Array(n + 1);
    for (let j = 0; j < k; j++) {
        for (let i = 0; i < x; i++) cur[i] = Infinity;
        for (let i = x; i <= n; i++) {
            const a = cur[i - 1];
            const b = prev[i - x] + cost[i - x];
            cur[i] = a < b ? a : b;
        }
        const t = prev;
        prev = cur;
        cur = t;
    }
    return prev[n];
};
```

### TypeScript

```typescript
function minOperations(nums: number[], x: number, k: number): number {
    const n: number = nums.length;
    const sorted: Int32Array = Int32Array.from(nums).sort();
    let m: number = 0;
    for (let i = 0; i < n; i++) {
        if (m === 0 || sorted[i] !== sorted[m - 1]) {
            sorted[m++] = sorted[i];
        }
    }
    const vals: Int32Array = sorted;
    const rk: Int32Array = new Int32Array(n);
    for (let i = 0; i < n; i++) {
        let lo = 0, hi = m - 1;
        while (lo < hi) {
            const mid = (lo + hi) >> 1;
            if (vals[mid] < nums[i]) lo = mid + 1;
            else hi = mid;
        }
        rk[i] = lo + 1;
    }
    const cnt: Int32Array = new Int32Array(m + 1);
    const sm: Float64Array = new Float64Array(m + 1);
    let top: number = 1;
    while (top * 2 <= m) top *= 2;
    const need: number = (x + 1) >> 1;
    const w: number = n - x + 1;
    const cost: Float64Array = new Float64Array(w);
    let total: number = 0;
    for (let i = 0; i < n; i++) {
        total += nums[i];
        for (let p = rk[i]; p <= m; p += p & -p) {
            cnt[p] += 1;
            sm[p] += nums[i];
        }
        if (i >= x) {
            total -= nums[i - x];
            for (let p = rk[i - x]; p <= m; p += p & -p) {
                cnt[p] -= 1;
                sm[p] -= nums[i - x];
            }
        }
        if (i >= x - 1) {
            let pos = 0, rem = need, s0 = 0;
            for (let pw = top; pw > 0; pw >>= 1) {
                const nxt = pos + pw;
                if (nxt <= m && cnt[nxt] < rem) {
                    pos = nxt;
                    rem -= cnt[nxt];
                    s0 += sm[nxt];
                }
            }
            const med: number = vals[pos];
            const c0: number = need - rem;
            cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0);
        }
    }
    let prev: Float64Array = new Float64Array(n + 1);
    let cur: Float64Array = new Float64Array(n + 1);
    for (let j = 0; j < k; j++) {
        for (let i = 0; i < x; i++) cur[i] = Infinity;
        for (let i = x; i <= n; i++) {
            const a = cur[i - 1];
            const b = prev[i - x] + cost[i - x];
            cur[i] = a < b ? a : b;
        }
        const t = prev;
        prev = cur;
        cur = t;
    }
    return prev[n];
};
```

### Python3

```python
class Solution:
    def minOperations(self, nums: List[int], x: int, k: int) -> int:
        n = len(nums)
        vals = sorted(set(nums))
        m = len(vals)
        pos_of = {v: i + 1 for i, v in enumerate(vals)}
        rk = [pos_of[v] for v in nums]
        cnt = [0] * (m + 1)
        sm = [0] * (m + 1)
        top = 1
        while top * 2 <= m:
            top *= 2
        need = (x + 1) // 2
        w = n - x + 1
        cost = [0] * w
        total = 0
        for i in range(n):
            v = nums[i]
            total += v
            p = rk[i]
            while p <= m:
                cnt[p] += 1
                sm[p] += v
                p += p & -p
            if i >= x:
                u = nums[i - x]
                total -= u
                p = rk[i - x]
                while p <= m:
                    cnt[p] -= 1
                    sm[p] -= u
                    p += p & -p
            if i >= x - 1:
                pos = 0
                rem = need
                s0 = 0
                pw = top
                while pw:
                    nxt = pos + pw
                    if nxt <= m and cnt[nxt] < rem:
                        pos = nxt
                        rem -= cnt[nxt]
                        s0 += sm[nxt]
                    pw >>= 1
                med = vals[pos]
                c0 = need - rem
                cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0)
        INF = 1 << 60
        prev = [0] * (n + 1)
        for _ in range(k):
            cur = [INF] * (n + 1)
            for i in range(x, n + 1):
                a = cur[i - 1]
                b = prev[i - x] + cost[i - x]
                cur[i] = a if a < b else b
            prev = cur
        return prev[n]
```

### Go

```go
func minOperations(nums []int, x int, k int) int64 {
    n := len(nums)
    vals := make([]int, n)
    copy(vals, nums)
    sort.Ints(vals)
    m := 0
    for i := 0; i < n; i++ {
        if m == 0 || vals[i] != vals[m-1] {
            vals[m] = vals[i]
            m++
        }
    }
    vals = vals[:m]
    rk := make([]int, n)
    for i, v := range nums {
        rk[i] = sort.SearchInts(vals, v) + 1
    }
    cnt := make([]int, m+1)
    sm := make([]int, m+1)
    top := 1
    for top*2 <= m {
        top *= 2
    }
    need := (x + 1) / 2
    w := n - x + 1
    cost := make([]int, w)
    total := 0
    for i := 0; i < n; i++ {
        total += nums[i]
        for p := rk[i]; p <= m; p += p & -p {
            cnt[p]++
            sm[p] += nums[i]
        }
        if i >= x {
            total -= nums[i-x]
            for p := rk[i-x]; p <= m; p += p & -p {
                cnt[p]--
                sm[p] -= nums[i-x]
            }
        }
        if i >= x-1 {
            pos, rem, s0 := 0, need, 0
            for pw := top; pw > 0; pw >>= 1 {
                nxt := pos + pw
                if nxt <= m && cnt[nxt] < rem {
                    pos = nxt
                    rem -= cnt[nxt]
                    s0 += sm[nxt]
                }
            }
            med := vals[pos]
            c0 := need - rem
            cost[i-x+1] = med*c0 - s0 + (total - s0) - med*(x-c0)
        }
    }
    const inf = 1 << 60
    prev := make([]int, n+1)
    cur := make([]int, n+1)
    for j := 0; j < k; j++ {
        for i := 0; i < x; i++ {
            cur[i] = inf
        }
        for i := x; i <= n; i++ {
            a := cur[i-1]
            b := prev[i-x] + cost[i-x]
            if a < b {
                cur[i] = a
            } else {
                cur[i] = b
            }
        }
        prev, cur = cur, prev
    }
    return int64(prev[n])
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow the same logic; only the language syntax differs.

**Value compression**  
Copy the array, sort it, remove duplicates. The remaining sorted list becomes the dictionary of ranks. Every original number is replaced by its 1-based index in that list. This step is identical in every language (binary search or a map is used to obtain the rank).

**Fenwick trees**  
Two arrays of size `m+1` are allocated. Updates add or subtract 1 (frequency) and the actual value (sum) while jumping by the lowest set bit. The same update code appears in every language; only the loop syntax changes.

**Finding the median**  
A binary-lifting walk starts from the highest power of two that does not exceed `m`. At each step we try to jump forward if the block we would add still contains fewer elements than we still need. While jumping we also accumulate the sum of those blocks. When the walk finishes we know both the median value and the sum of everything strictly smaller than it. This walk is written almost identically in C++, Java, Go, JavaScript, TypeScript and Python.

**Cost formula**  
Let `c0` be the number of elements below the median and `s0` their sum.  
Cost = `(median · c0 - s0) + (total - s0 - median · (x - c0))`.  
Elements equal to the median contribute zero, so it does not matter which side we assign them to.

**Dynamic programming**  
Two arrays `prev` and `cur` of length `n+1` are used.  
`prev` starts filled with zeros.  
For each of the `k` required windows we fill `cur`:

- positions shorter than `x` receive a huge sentinel (impossible),
- every later position takes the minimum of “skip” (`cur[i-1]`) and “place a window ending here” (`prev[i-x] + cost[i-x]`).  
After the row is finished the two arrays are swapped.  
After exactly `k` iterations the answer is `prev[n]`.

Language-specific notes:

- JavaScript / TypeScript use `Float64Array` for the sum tree and for the DP arrays so that large intermediate values stay exact.
- Python builds the rank map with a dictionary for brevity.
- Go and Java perform the unique step in-place with a write pointer.

## Examples

**Example 1**  
Input: `nums = [1, 2, 3, 4, 5]`, `x = 2`, `k = 2`  
Possible non-overlapping pairs of windows:  

- indices [0,1] and [2,3] → costs 1 + 1 = 2  
- indices [0,1] and [3,4] → costs 1 + 1 = 2  
- indices [1,2] and [3,4] → costs 1 + 1 = 2  
Answer: 2

**Example 2**  
Input: `nums = [5, 1, 5, 1, 5]`, `x = 3`, `k = 1`  
Only one window is needed. The three possible windows have costs:  

- [5,1,5] → median 5, cost 4  
- [1,5,1] → median 1, cost 4  
- [5,1,5] → median 5, cost 4  
Answer: 4

**Example 3**  
Input: `nums = [10, 3, 8, 5, 2]`, `x = 2`, `k = 2`  
Cheapest pair of non-overlapping windows of length 2 yields total cost 6.

## How to Use / Run Locally

1. Copy the solution of the language you prefer into a file (`solution.cpp`, `Solution.java`, `solution.js`, \ldots).
2. Compile / run with the usual commands:

- **C++**  
  `g++ -std=c++17 solution.cpp -o sol && ./sol`

- **Java**  
  `javac Solution.java && java Solution`

- **JavaScript**  
  `node solution.js`

- **TypeScript**  
  `tsc solution.ts && node solution.js`

- **Python3**  
  `python3 solution.py`

- **Go**  
  `go run solution.go`

Add a small `main` that reads the array, `x` and `k` from standard input if you want to test interactively.

## Notes & Optimizations

- The constraint `k · x ≤ n` guarantees a feasible solution always exists, so the DP never returns the sentinel infinity.
- Using two rolling arrays instead of a full `k × n` table reduces memory from `O(nk)` to `O(n)`.
- The same median-cost idea works for any odd or even window length; we simply take the lower median when the length is even.
- If `k` were larger (hundreds) a different approach such as convex-hull optimization or a min-cost flow would be needed; with `k ≤ 15` the plain DP is already optimal.
- Fenwick trees can be replaced by two balanced BSTs (or policy-based data structures in C++) that also support order statistics; the asymptotic complexity stays the same.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
