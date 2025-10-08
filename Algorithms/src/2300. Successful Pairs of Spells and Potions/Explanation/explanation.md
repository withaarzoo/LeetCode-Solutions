# 2300. Successful Pairs of Spells and Potions

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

We are given two arrays:

* `spells` of length `n`, where `spells[i]` is the strength of the i-th spell.
* `potions` of length `m`, where `potions[j]` is the strength of the j-th potion.
  And an integer `success`.

A pair (spell, potion) is **successful** if `spell * potion >= success`.

Return an array `pairs` of length `n` where `pairs[i]` is the number of potions that will form a successful pair with the i-th spell.

---

## Constraints

* `n == spells.length`
* `m == potions.length`
* `1 <= n, m <= 10^5`
* `1 <= spells[i], potions[i] <= 10^5`
* `1 <= success <= 10^10`

**Important:** Products and the `success` threshold may exceed 32-bit integer ranges, so we must use 64-bit arithmetic where appropriate.

---

## Intuition

I thought: checking each spell with every potion is O(n·m) and will be too slow when `n` and `m` can each be 10^5. If I **sort `potions`**, then for each spell I can compute the minimum potion strength required to reach `success`. Using **binary search** on the sorted `potions` array I can find the first valid potion and then count how many potions from that index to the end are valid. This reduces the per-spell work to `O(log m)`.

---

## Approach

1. Sort `potions` (ascending).
2. For each spell `s`:

   * Compute `need = ceil(success / s)` — the smallest potion value `p` such that `s * p >= success`.
   * Binary search `potions` to find the first index `idx` where `potions[idx] >= need`.
   * The number of successful potions for this spell is `m - idx`.
3. Collect results in an output array in the same order as the spells.

---

## Data Structures Used

* Arrays / Lists for input and output.
* Sorting (in-place) on the `potions` array.
* Binary search (lower bound) on the sorted `potions`.

---

## Operations & Behavior Summary

* Sorting `potions` allows quick range counts via binary search.
* For each spell, using integer arithmetic (`ceil`) keeps calculations exact.
* Use 64-bit integer types (long, long long, int64) to avoid overflow since `success` can be up to `1e10`.

---

## Complexity

* **Time Complexity:** `O(m log m + n log m)`

  * `m` = number of potions, `n` = number of spells.
  * Sorting `potions`: `O(m log m)`
  * For each of the `n` spells we do a `O(log m)` binary search.
* **Space Complexity:** `O(1)` extra (excluding output array). Sorting in-place; only the output array of size `n` is extra.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    // LeetCode-style signature
    vector<int> successfulPairs(vector<int>& spells, vector<int>& potions, long long success) {
        sort(potions.begin(), potions.end());         // sort potions for binary search
        int n = spells.size();
        int m = potions.size();
        vector<int> ans(n, 0);
        for (int i = 0; i < n; ++i) {
            long long s = spells[i];
            // compute ceil(success / s) using integer arithmetic
            long long need = (success + s - 1) / s;
            // lower_bound gives first index where potions[idx] >= need
            auto it = lower_bound(potions.begin(), potions.end(), need);
            ans[i] = m - int(it - potions.begin());
        }
        return ans;
    }
};
```

---

### Java

```java
import java.util.Arrays;

class Solution {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);                   // sort potions
        int n = spells.length;
        int m = potions.length;
        int[] ans = new int[n];

        for (int i = 0; i < n; ++i) {
            long s = spells[i];
            long need = (success + s - 1) / s;  // ceil(success / s)
            int pos = lowerBound(potions, need);
            ans[i] = m - pos;
        }
        return ans;
    }

    // returns first index where arr[index] >= target
    private int lowerBound(int[] arr, long target) {
        int l = 0, r = arr.length;
        while (l < r) {
            int mid = (l + r) >>> 1;
            if ((long)arr[mid] < target) l = mid + 1;
            else r = mid;
        }
        return l;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} spells
 * @param {number[]} potions
 * @param {number} success
 * @return {number[]}
 */
var successfulPairs = function(spells, potions, success) {
    potions.sort((a,b) => a - b);
    const m = potions.length;
    const res = new Array(spells.length);

    for (let i = 0; i < spells.length; ++i) {
        const s = spells[i];
        // compute ceil(success / s) using integer math (avoid floats rounding issues)
        const need = Math.floor((success + s - 1) / s);
        // binary search first index >= need
        let l = 0, r = m;
        while (l < r) {
            let mid = Math.floor((l + r) / 2);
            if (potions[mid] < need) l = mid + 1;
            else r = mid;
        }
        res[i] = m - l;
    }
    return res;
};
```

---

### Python3

```python
from bisect import bisect_left
from typing import List

class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        m = len(potions)
        ans = []
        for s in spells:
            need = (success + s - 1) // s   # ceil(success / s)
            pos = bisect_left(potions, need)
            ans.append(m - pos)
        return ans
```

---

### Go

```go
package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    m := len(potions)
    res := make([]int, len(spells))

    for i, s := range spells {
        s64 := int64(s)
        // ceil(success / s)
        need := (success + s64 - 1) / s64

        // binary search first index with value >= need
        l, r := 0, m
        for l < r {
            mid := (l + r) / 2
            if int64(potions[mid]) < need {
                l = mid + 1
            } else {
                r = mid
            }
        }
        res[i] = m - l
    }
    return res
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll explain the common logic first (applies to every language). Then I’ll briefly explain language-specific parts.

**High-level reasoning (common):**

1. Sort `potions` in ascending order. After sorting, potions that are large enough for a spell form a contiguous suffix in the array.
2. For a spell `s`, we need potions `p` such that `s * p >= success`. So `p >= success / s`. Because of integer arithmetic, we compute `need = ceil(success / s)`. If a potion is `>= need`, it's valid.
3. Use binary search on the sorted `potions` to find the first index `idx` with `potions[idx] >= need`. Then `m - idx` potions succeed with this spell.
4. Repeat for each spell and collect results.

---

### Detailed line-by-line mapping (Python example — easiest to read)

```python
from bisect import bisect_left
from typing import List
```

* `bisect_left` gives the first index where an element can be inserted to keep the list sorted; effectively the lower bound.

```python
potions.sort()
m = len(potions)
ans = []
```

* Sort potions: `O(m log m)`. `m` used for final counts. `ans` stores results.

```python
for s in spells:
    need = (success + s - 1) // s
```

* `need = ceil(success / s)` computed without floating-point errors. Example: if `success = 7` and `s = 5`, need = `(7 + 5 - 1)//5 = 11//5 = 2`.

```python
pos = bisect_left(potions, need)
ans.append(m - pos)
```

* `bisect_left` returns first index of potion >= `need`. `m - pos` is the number of potions from `pos` to `m-1`, inclusive.

**Language-specific notes**

* **C++**: Use `long long` for `success` and intermediate `need`. Use `lower_bound` from `<algorithm>`.
* **Java**: Use `long` for `success` and cast `arr[mid]` to long when comparing. Implement custom `lowerBound` because `Arrays.binarySearch` is less convenient for insertion points.
* **JavaScript**: Numbers are `double` (IEEE 754) but integer operations up to 2^53 are exact — `success <= 1e10` is safe. Use `Math.floor((success + s - 1) / s)` for integer ceil formula.
* **Go**: Use `int64` for `success` and cast potion values to `int64` in comparisons. Use a manual binary search loop.
* In all languages, prefer integer arithmetic for the ceil formula to avoid rounding issues.

---

## Examples

**Example 1**

```
Input:
spells = [5,1,3]
potions = [1,2,3,4,5]
success = 7

Output:
[4,0,3]

Explanation:
- spell 5 -> need ceil(7/5)=2 → potions >=2 => [2,3,4,5] => 4
- spell 1 -> need 7 → potions >=7 => none => 0
- spell 3 -> need ceil(7/3)=3 → potions >=3 => [3,4,5] => 3
```

**Example 2**

```
Input:
spells = [3,1,2]
potions = [8,5,8]
success = 16

Output:
[2,0,2]

Explanation:
- spell 3 -> need ceil(16/3)=6 → potions >=6 => [8,8] => 2
- spell 1 -> need 16 → none => 0
- spell 2 -> need 8 → potions >=8 => [8,8] => 2
```

---

## How to use / Run locally

Below are quick instructions for running sample tests. For each language you can create a small test harness around the solution.

### Python

1. Create `solution.py` with the `Solution` class and add this test harness at the bottom:

```python
if __name__ == "__main__":
    spells = [5,1,3]
    potions = [1,2,3,4,5]
    success = 7
    s = Solution()
    print(s.successfulPairs(spells, potions, success))  # -> [4,0,3]
```

2. Run:

```bash
python3 solution.py
```

### JavaScript (Node)

1. Create `solution.js` with the `successfulPairs` function and add:

```javascript
// test
const spells = [5,1,3];
const potions = [1,2,3,4,5];
const success = 7;
console.log(successfulPairs(spells, potions, success)); // -> [4,0,3]
```

2. Run:

```bash
node solution.js
```

### C++

1. You can adapt the `Solution` class into a full program with `main()` to run tests, or use it on LeetCode. Example quick-run file `main.cpp`:

```c++
#include <bits/stdc++.h>
using namespace std;

// paste Solution class here

int main() {
    Solution sol;
    vector<int> spells = {5,1,3};
    vector<int> potions = {1,2,3,4,5};
    long long success = 7;
    vector<int> ans = sol.successfulPairs(spells, potions, success);
    for (int v : ans) cout << v << " ";
    cout << endl; // -> 4 0 3
    return 0;
}
```

2. Compile and run:

```bash
g++ -std=c++17 main.cpp -O2 -o main
./main
```

### Java

1. Paste the `Solution` class into `Solution.java` and add a `main` to test:

```java
public class SolutionTest {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int[] spells = {5,1,3};
        int[] potions = {1,2,3,4,5};
        long success = 7;
        int[] ans = sol.successfulPairs(spells, potions, success);
        System.out.println(Arrays.toString(ans)); // -> [4,0,3]
    }
}
```

2. Compile and run:

```bash
javac Solution.java SolutionTest.java
java SolutionTest
```

### Go

1. Create `main.go` with the `successfulPairs` function and a main:

```go
package main

import (
    "fmt"
)

func main() {
    spells := []int{5,1,3}
    potions := []int{1,2,3,4,5}
    success := int64(7)
    fmt.Println(successfulPairs(spells, potions, success)) // -> [4 0 3]
}
```

2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* Use 64-bit integers (`long`, `long long`, `int64`) to avoid overflow: `spells[i] * potions[j]` may exceed 32-bit.
* Sorting potions once, then performing `n` binary searches is asymptotically optimal for this approach.
* If memory is tight and `potions` must remain unchanged, sort a copy to preserve original order.
* In languages with only floating-point numeric type (JS), the constraints still fit safely within integer-exact representation (up to 2^53). If you want absolute safety for arbitrarily large numbers use `BigInt` in JS (and adjust arithmetic).
* If many queries with the same spells are expected, consider caching results for repeated spells.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
