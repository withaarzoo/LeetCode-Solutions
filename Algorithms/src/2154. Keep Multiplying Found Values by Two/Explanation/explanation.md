# 2154. Keep Multiplying Found Values by Two

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

You are given an array of integers `nums` and an integer `original`. While `original` exists in `nums`, multiply it by 2 (i.e., `original = original * 2`) and repeat. Stop when the new `original` is not present in `nums`. Return the final value of `original`.

Example:

* `nums = [5,3,6,1,12], original = 3` → sequence: 3 → 6 → 12 → 24 → **24** (24 not found) → return `24`.

---

## Constraints

* `1 <= nums.length <= 1000`
* `1 <= nums[i], original <= 1000`

These constraints mean `nums` is reasonably small (≤ 1000), and all values are positive integers bounded by 1000.

---

## Intuition

I thought: the only operation I need to do repeatedly is “check whether the current `original` is present in `nums`”. If I can check membership quickly, I can repeatedly double `original` until it stops being present. A hash-based set gives O(1) average membership checks, so I convert `nums` into a set and loop `while original in set: original *= 2`. That’s it.

---

## Approach

1. Convert `nums` into a hash set (or equivalent) for constant-time membership checks.
2. While the set contains `original`, multiply `original` by 2.
3. When `original` is not in the set, return `original`.

This approach is simple, clear, and efficient.

---

## Data Structures Used

* **Hash set** (C++: `unordered_set`, Java: `HashSet`, JavaScript: `Set`, Python: `set`, Go: `map[int]bool`) for O(1) average-time membership queries.

---

## Operations & Behavior Summary

* Build set from input array: O(n).
* Repeatedly check membership and double `original`. Each check is O(1) on average.
* Doubling continues until `original` is not found in the set.
* Return the final `original`.

Edge cases:

* If `original` is not in `nums` initially, return it immediately.
* Duplicate numbers in `nums` do not affect the result because the set ignores duplicates.

---

## Complexity

* **Time Complexity:** `O(n + t)` where:

  * `n` = `nums.length` (time to build the set),
  * `t` = number of times we double `original` until it stops being found. `t` is bounded because values in `nums` are finite and small. In practice `t` is small.
* **Space Complexity:** `O(n)` for the hash set.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <unordered_set>
using namespace std;

class Solution {
public:
    int findFinalValue(vector<int>& nums, int original) {
        // Build set for O(1) membership checks
        unordered_set<int> s(nums.begin(), nums.end());
        // While original is found in the set, double it
        while (s.count(original)) {
            original *= 2;
        }
        return original;
    }
};
```

### Java

```java
import java.util.HashSet;
import java.util.Set;

class Solution {
    public int findFinalValue(int[] nums, int original) {
        // Build a HashSet for O(1) average lookup
        Set<Integer> set = new HashSet<>();
        for (int x : nums) set.add(x);
        // Keep doubling original while it's present
        while (set.contains(original)) {
            original *= 2;
        }
        return original;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function(nums, original) {
    // Use Set for constant-time membership checks
    const s = new Set(nums);
    // While original exists, double it
    while (s.has(original)) {
        original *= 2;
    }
    return original;
};
```

### Python3

```python3
class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        # Use set for O(1) average-time membership queries
        s = set(nums)
        # Keep doubling while original present
        while original in s:
            original *= 2
        return original
```

### Go

```go
func findFinalValue(nums []int, original int) int {
    // Use map[int]bool as a set
    seen := make(map[int]bool, len(nums))
    for _, v := range nums {
        seen[v] = true
    }
    // While original is in seen, double it
    for seen[original] {
        original *= 2
    }
    return original
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the shared logic once, then point out language-specific details.

**Common algorithm (shared logic)**:

1. Convert `nums` into a set called `s` (or `seen`) so membership checks are O(1) on average.
2. Check if `original` is in the set:

   * If yes → `original = original * 2`, then repeat step 2.
   * If no → stop and return `original`.

**Why this works**:

* Because each step only depends on whether the current `original` exists in `nums`. We don't need order or counts — just membership.
* Using a set avoids repeated O(n) searches; otherwise each membership test would be linear and total complexity could become O(n * t).

### C++ (line-by-line)

```c++
unordered_set<int> s(nums.begin(), nums.end());
```

* Creates an unordered_set containing all elements of `nums` in O(n) time.

```c++
while (s.count(original)) {
    original *= 2;
}
```

* `s.count(original)` returns 1 if `original` is in the set, else 0.
* While it's present, we double `original` and check again.
* When loop exits, `original` isn't in `s` so we return it.

### Java (line-by-line)

```java
Set<Integer> set = new HashSet<>();
for (int x : nums) set.add(x);
```

* Build a HashSet from the array.

```java
while (set.contains(original)) {
    original *= 2;
}
```

* `contains` checks membership in expected O(1) time.
* Double until not found.

### JavaScript (line-by-line)

```js
const s = new Set(nums);
```

* Create a Set of numbers from `nums`.

```js
while (s.has(original)) {
    original *= 2;
}
```

* `has` checks presence. Double while present.

### Python3 (line-by-line)

```py
s = set(nums)
while original in s:
    original *= 2
```

* `set(nums)` is O(n). `in` is O(1) average. Keep doubling until `original` not in the set.

### Go (line-by-line)

```go
seen := make(map[int]bool, len(nums))
for _, v := range nums {
    seen[v] = true
}
```

* Go has no built-in set; `map[int]bool` is used as a set.

```go
for seen[original] {
    original *= 2
}
```

* `seen[original]` returns `false` if absent, `true` if present. Loop doubles `original` while present.

---

## Examples

1. Example 1:

   * Input: `nums = [5,3,6,1,12]`, `original = 3`
   * Process: `3` (found) → `6` (found) → `12` (found) → `24` (not found) → return `24`
   * Output: `24`

2. Example 2:

   * Input: `nums = [2,7,9]`, `original = 4`
   * Process: `4` (not found) → return `4`
   * Output: `4`

3. Edge:

   * Input: `nums = [1,2,4,8]`, `original = 1`
   * Process: `1 → 2 → 4 → 8 → 16` (stop) → return `16`

---

## How to use / Run locally

### C++

1. Place the `Solution` class in a file `solution.cpp`.
2. Add a `main()` that reads input or testcases and calls `Solution().findFinalValue(...)`.
3. Compile: `g++ -std=c++17 solution.cpp -O2 -o solution`
4. Run: `./solution`

### Java

1. Save `Solution` as `Solution.java`.
2. Add `public static void main(String[] args)` to call `findFinalValue`.
3. Compile: `javac Solution.java`
4. Run: `java Solution`

### JavaScript (Node.js)

1. Save function code to `solution.js` and export or call it with test inputs.
2. Run: `node solution.js`

### Python3

1. Save the `Solution` class in `solution.py`.
2. Add test code below (e.g., `print(Solution().findFinalValue([5,3,6,1,12], 3))`).
3. Run: `python3 solution.py`

### Go

1. Create file `main.go`, include `findFinalValue` function and `main` that calls it and prints result.
2. Build: `go build -o solution main.go`
3. Run: `./solution`

---

## Notes & Optimizations

* Using a set is the most straightforward and performant approach here.
* Because `nums` length ≤ 1000 and values ≤ 1000, this solution is more than adequate.
* If you were constrained on memory and `nums` were sorted, you could binary-search for each `original` in a sorted array; but set-based solution is simpler and faster in practice.
* Avoid looping through `nums` for every check (that would be O(n * t)). Always prefer O(1) membership checks if possible.
* There is no concern with integer overflow given constraints `original <= 1000` and typical languages — but if doubling could exceed language limits, consider using `long`/`long long` or check bounds.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
