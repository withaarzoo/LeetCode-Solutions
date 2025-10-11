# 3186. Maximum Total Damage With Spell Casting

---

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

I am given an array `power`, where each element is the damage value of a spell. If I cast any spell with damage `v`, I **cannot** cast any spell with damage `v-2, v-1, v+1, or v+2`. Each individual spell (array element) can be cast at most once, but there can be multiple spells with the same damage and I may take all of them if I choose that damage value. I must compute the maximum total damage I can cast.

---

## Constraints

* `1 <= power.length <= 10^5`
* `1 <= power[i] <= 10^9`
* Values and frequencies are large → use 64-bit integers for sums.

---

## Intuition

I thought about how casting a damage `v` forbids nearby values `v±1` and `v±2`. But if there are multiple spells with the same damage `v`, taking `v` lets me take **all** those spells (sum = `v * freq[v]`). So I can group spells by damage and make a decision per unique damage value: either I take the entire group's sum or I skip it. Because conflicts exist only for values within distance 2, the earliest non-conflicting value for `v` is any damage `<= v-3`. That leads naturally to a DP on sorted unique damage values.

---

## Approach

1. Count frequency of each damage value: `freq[v]`.
2. For each unique damage value `v`, compute `valueSum[v] = v * freq[v]`.
3. Get the list `vals` of unique damage values and sort ascending.
4. Let `dp[i]` be the maximum damage achievable using `vals[0..i]`.

   * Option 1 (take `vals[i]`): `take = valueSum[i] + dp[j]` where `j` is the last index `< i` with `vals[j] <= vals[i] - 3`. If no such `j`, add 0.
   * Option 2 (skip `vals[i]`): `skip = dp[i-1]`.
   * `dp[i] = max(take, skip)`.
5. Answer is `dp[last]`.
6. Use binary search to find `j` (or maintain a two-pointer for O(m) DP after sorting).

---

## Data Structures Used

* Hash map / dictionary (`value -> frequency`)
* Sorted array of unique damage values (`vals`)
* DP array (`dp`) of length `m` (number of unique values)

---

## Operations & Behavior Summary

* Build frequency map: O(n)
* Sort unique values: O(m log m) (m ≤ n)
* For each unique value compute best using binary search: O(m log m) total, or O(m) with two-pointer
* Use 64-bit integer arithmetic for sums.

---

## Complexity

* **Time Complexity:** `O(n + m log m)` where:

  * `n` = `power.length` (to build frequencies)
  * `m` = number of unique damage values (≤ `n`) (sort cost `O(m log m)` and DP with binary search `O(m log m)`).
    With a two-pointer approach the DP phase can be `O(m)`, so overall `O(n + m log m)`.
* **Space Complexity:** `O(m)` for frequency map, sorted unique array, and DP array.

---

## Multi-language Solutions

Below are complete, clear, and commented implementations. All use the same algorithm: group-by-value, sort unique values, then DP over unique values. Each snippet also contains a tiny test harness that prints an example result.

---

### C++

```c++
/*
 * Solution (C++) - Group by damage value and run DP over unique sorted values.
 * Compile: g++ -std=c++17 -O2 solution.cpp -o solution
 * Run: ./solution
 */

#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    long long maximumTotalDamage(vector<int>& power) {
        // count frequency
        unordered_map<long long, long long> freq;
        for (int v : power) freq[v] += 1;

        // collect unique values and sort them
        vector<long long> vals;
        vals.reserve(freq.size());
        for (auto &p : freq) vals.push_back(p.first);
        sort(vals.begin(), vals.end());

        int m = (int)vals.size();
        if (m == 0) return 0;

        // valueSum[i] = vals[i] * freq[vals[i]]
        vector<long long> valueSum(m);
        for (int i = 0; i < m; ++i) valueSum[i] = vals[i] * freq[vals[i]];

        // dp[i] = best up to i
        vector<long long> dp(m, 0);
        dp[0] = valueSum[0];

        for (int i = 1; i < m; ++i) {
            // need = vals[i] - 3 => last index j with vals[j] <= need is non-conflicting
            long long need = vals[i] - 3;
            int j = (int)(upper_bound(vals.begin(), vals.begin() + i, need) - vals.begin() - 1);
            long long take = valueSum[i] + (j >= 0 ? dp[j] : 0LL);
            long long skip = dp[i - 1];
            dp[i] = max(skip, take);
        }
        return dp[m - 1];
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    vector<int> power = {1, 1, 3, 4}; // Example input
    Solution sol;
    cout << sol.maximumTotalDamage(power) << '\n'; // Expected output: 6
    return 0;
}
```

---

### Java

```java
/*
 * Solution (Java) - Group by damage value and DP over unique sorted values.
 * Compile: javac Solution.java
 * Run: java Solution
 */

import java.util.*;

public class Solution {

    public long maximumTotalDamage(int[] power) {
        if (power == null || power.length == 0) return 0L;

        // frequency map
        HashMap<Long, Long> freq = new HashMap<>();
        for (int v : power) freq.put((long) v, freq.getOrDefault((long) v, 0L) + 1L);

        // unique sorted values
        long[] vals = new long[freq.size()];
        int idx = 0;
        for (Long k : freq.keySet()) vals[idx++] = k;
        Arrays.sort(vals);

        int m = vals.length;
        long[] valueSum = new long[m];
        for (int i = 0; i < m; ++i) valueSum[i] = vals[i] * freq.get(vals[i]);

        long[] dp = new long[m];
        dp[0] = valueSum[0];

        for (int i = 1; i < m; ++i) {
            long need = vals[i] - 3; // values <= need are safe to combine
            int lo = 0, hi = i - 1, j = -1;
            while (lo <= hi) {
                int mid = (lo + hi) >>> 1;
                if (vals[mid] <= need) { j = mid; lo = mid + 1; }
                else hi = mid - 1;
            }
            long take = valueSum[i] + (j >= 0 ? dp[j] : 0L);
            long skip = dp[i - 1];
            dp[i] = Math.max(skip, take);
        }
        return dp[m - 1];
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] power = {1, 1, 3, 4};
        System.out.println(s.maximumTotalDamage(power)); // Expected: 6
    }
}
```

---

### JavaScript

```javascript
/**
 * Solution (JavaScript) - group & DP approach
 * Run: node solution.js
 */

function maximumTotalDamage(power) {
    if (!power || power.length === 0) return 0;

    // frequency map
    const freq = new Map();
    for (const v of power) freq.set(v, (freq.get(v) || 0) + 1);

    // unique sorted values
    const vals = Array.from(freq.keys()).sort((a,b) => a - b);
    const m = vals.length;
    if (m === 0) return 0;

    const valueSum = new Array(m);
    for (let i = 0; i < m; ++i) valueSum[i] = vals[i] * freq.get(vals[i]);

    const dp = new Array(m).fill(0);
    dp[0] = valueSum[0];

    for (let i = 1; i < m; ++i) {
        const need = vals[i] - 3;
        // binary search for last index <= need
        let lo = 0, hi = i - 1, j = -1;
        while (lo <= hi) {
            const mid = (lo + hi) >> 1;
            if (vals[mid] <= need) { j = mid; lo = mid + 1; }
            else hi = mid - 1;
        }
        const take = valueSum[i] + (j >= 0 ? dp[j] : 0);
        const skip = dp[i - 1];
        dp[i] = Math.max(skip, take);
    }
    return dp[m - 1];
}

// Example
console.log(maximumTotalDamage([1,1,3,4])); // Expected: 6
```

---

### Python3

```python
"""
Solution (Python3) - group by value and DP over sorted unique values.
Run: python3 solution.py
"""

from collections import Counter
import bisect

def maximumTotalDamage(power):
    if not power:
        return 0
    cnt = Counter(power)
    vals = sorted(cnt.keys())  # sorted unique damage values
    m = len(vals)
    value_sum = [vals[i] * cnt[vals[i]] for i in range(m)]
    dp = [0] * m
    dp[0] = value_sum[0]
    for i in range(1, m):
        need = vals[i] - 3  # last allowed value <= need
        j = bisect.bisect_right(vals, need, 0, i) - 1
        take = value_sum[i] + (dp[j] if j >= 0 else 0)
        skip = dp[i - 1]
        dp[i] = max(skip, take)
    return dp[-1]

# Example
if __name__ == "__main__":
    print(maximumTotalDamage([1,1,3,4]))  # Expected: 6
```

---

### Go

```go
/*
Solution (Go) - group by value and DP over unique sorted values.
Run: go run solution.go
*/

package main

import (
    "fmt"
    "sort"
)

func maximumTotalDamage(power []int) int64 {
    if len(power) == 0 {
        return 0
    }
    freq := make(map[int64]int64)
    for _, v := range power {
        freq[int64(v)]++
    }

    vals := make([]int64, 0, len(freq))
    for k := range freq { vals = append(vals, k) }
    sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })

    m := len(vals)
    valueSum := make([]int64, m)
    for i := 0; i < m; i++ { valueSum[i] = vals[i] * freq[vals[i]] }

    dp := make([]int64, m)
    dp[0] = valueSum[0]

    for i := 1; i < m; i++ {
        need := vals[i] - 3
        lo, hi, j := 0, i-1, -1
        for lo <= hi {
            mid := (lo + hi) / 2
            if vals[mid] <= need { j = mid; lo = mid + 1 } else { hi = mid - 1 }
        }
        take := valueSum[i]
        if j >= 0 { take += dp[j] }
        skip := dp[i-1]
        if take > skip { dp[i] = take } else { dp[i] = skip }
    }
    return dp[m-1]
}

func main() {
    fmt.Println(maximumTotalDamage([]int{1, 1, 3, 4})) // Expected: 6
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the main common steps, then point out the corresponding lines for each language.

**Common algorithm steps (high-level)**

1. Build `freq` map: count occurrences of each damage.
2. Build `vals`: sorted list of unique damage values.
3. For each unique `v` compute `valueSum = v * freq[v]`.
4. Use a DP array `dp[i]` meaning "best total damage considering `vals[0..i]`".

   * If I pick `vals[i]`, I can add `valueSum[i]` plus `dp[j]` where `j` is the largest index with `vals[j] <= vals[i] - 3`.
   * Else I skip `vals[i]` and keep `dp[i-1]`.
   * Pick maximum of both options.
5. Return `dp[last]`.

**Why `vals[i] - 3`?**
Taking `v` forbids `v-2, v-1, v+1, v+2`. So the largest value that is still safe to combine with `v` is `v-3` or less. We find the rightmost index with value ≤ `v-3` and add its `dp` result.

---

### Line-by-line mapping (concept → code)

**1) Count frequencies**

* C++: `for (int v : power) freq[v] += 1;`
* Java: `for (int v : power) freq.put((long)v, freq.getOrDefault((long)v, 0L) + 1L);`
* JS: `for (const v of power) freq.set(v, (freq.get(v) || 0) + 1);`
* Python: `cnt = Counter(power)`
* Go: `for _, v := range power { freq[int64(v)]++ }`

**2) Build and sort unique values**

* C++: push keys into `vals` then `sort(vals.begin(), vals.end());`
* Java: put keys into `long[] vals` and `Arrays.sort(vals);`
* JS: `const vals = Array.from(freq.keys()).sort((a,b) => a-b);`
* Python: `vals = sorted(cnt.keys())`
* Go: append keys to `vals` and `sort.Slice(vals, ...)`

**3) Compute `valueSum` for each unique**

* C++: `valueSum[i] = vals[i] * freq[vals[i]];`
* Java: `valueSum[i] = vals[i] * freq.get(vals[i]);`
* JS: `valueSum[i] = vals[i] * freq.get(vals[i]);`
* Python: list comprehension `[vals[i] * cnt[vals[i]] for i in range(m)]`
* Go: `valueSum[i] = vals[i] * freq[vals[i]]`

**4) DP base**

* `dp[0] = valueSum[0];` in all languages.

**5) For each i > 0 compute j and dp**

* Find `need = vals[i] - 3`.
* Use binary search to find last index with `vals[j] <= need`.

  * C++: `int j = upper_bound(vals.begin(), vals.begin() + i, need) - vals.begin() - 1;`
  * Java: custom `while` binary search on `vals`.
  * JS: custom `while` binary search.
  * Python: `j = bisect_right(vals, need, 0, i) - 1`
  * Go: custom binary search `for lo <= hi { mid := (lo+hi)/2 ... }`
* `take = valueSum[i] + (j >= 0 ? dp[j] : 0)`
* `skip = dp[i-1]`
* `dp[i] = max(skip, take)`

**6) Return `dp[last]`** — this is the answer.

---

## Examples

**Example 1**

* Input: `power = [1,1,3,4]`
* Explanation: I can take damage `1` (both spells: sum=2) and `4` (sum=4) → total = `6`. Taking `3` conflicts with `1` and `4`.
* Output: `6`

**Example 2**

* Input: `power = [7,1,6,6]`
* Explanation: Best to take `1` (1) and both `6`s (12) → total = `13`. `6` conflicts with `4,5,7,8` so it conflicts with `7`, but `1` is fine.
* Output: `13`

---

## How to use / Run locally

Place the language-specific code into a file:

**C++**

1. Save as `solution.cpp`.
2. Compile: `g++ -std=c++17 -O2 solution.cpp -o solution`
3. Run: `./solution`

**Java**

1. Save as `Solution.java`.
2. Compile: `javac Solution.java`
3. Run: `java Solution`

**JavaScript**

1. Save as `solution.js`.
2. Run: `node solution.js`

**Python3**

1. Save as `solution.py`.
2. Run: `python3 solution.py`

**Go**

1. Save as `solution.go`.
2. Run: `go run solution.go`

To test with custom inputs, modify the example `power` array in the `main`/test area of each file.

---

## Notes & Optimizations

* Use 64-bit integers (`long long` / `long` / `int64`) because `value * freq` can exceed 32-bit.
* `m` (unique values) ≤ `n`. If all values are different, `m = n`.
* Binary search per DP step gives `O(m log m)`. You can convert the DP step to **two-pointer** (sliding pointer that only moves forward) to achieve `O(m)` after sorting because `j` never moves left as `i` increases.
* If the input domain were dense (e.g., small range), a different bucket/array approach might be faster, but here `power[i]` may be up to `10^9`, so map + unique sorting is the correct method.
* This approach chooses to **take all spells** of a damage value if I choose that value — this is optimal because picking some but not all identical damage spells never gives an advantage (they have identical conflict behavior).

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo) (solution & README) — competitive programming style, multi-language implementations, DP + grouping approach.
