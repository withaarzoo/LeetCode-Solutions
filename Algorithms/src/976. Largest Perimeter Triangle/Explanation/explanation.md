# 976. Largest Perimeter Triangle

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

Given an integer array `nums`, return the largest perimeter of a triangle with non-zero area that can be formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return `0`.

A triangle with sides `a`, `b`, `c` is valid (non-degenerate) if each side is positive and the triangle inequality holds. For sorted sides with `a <= b <= c`, the necessary and sufficient condition is `a + b > c`.

---

## Constraints

* `3 <= nums.length <= 10^4`
* `1 <= nums[i] <= 10^6`

---

## Intuition

I thought about how to get the **largest** perimeter. If I want the biggest sum of three sides that can form a triangle, it makes sense to try the largest sides first. Sorting the array puts the largest values at the end, so I check triples from the largest end. The first valid triple I find will have the maximum possible perimeter.

---

## Approach

1. Sort `nums` in non-decreasing order.
2. Iterate from the end of the sorted array toward the beginning.
3. For each index `i` (treat `nums[i]` as the largest side `c`), take `b = nums[i-1]` and `a = nums[i-2]`.
4. If `a + b > c`, return `a + b + c`.
5. If loop completes with no valid triple, return `0`.

This is a greedy approach — because I try largest triples first, the first valid one has the largest perimeter.

---

## Data Structures Used

* In-place array sort (no additional data structures required).
* Simple indexed access to the array.

---

## Operations & Behavior Summary

* Sorting the array: `O(n log n)`
* Single backward pass checking consecutive triplets: `O(n)`
* Early exit when a valid triple is found (makes average case faster).

---

## Complexity

* **Time Complexity:** `O(n log n)` — `n` is the length of `nums`. Sorting dominates the runtime.
* **Space Complexity:** `O(1)` extra space (if sorting in-place). No additional significant arrays or maps used.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int largestPerimeter(vector<int>& nums) {
        sort(nums.begin(), nums.end());            // O(n log n)
        for (int i = nums.size() - 1; i >= 2; --i) {
            int c = nums[i];
            int b = nums[i - 1];
            int a = nums[i - 2];
            if (a + b > c) return a + b + c;     // valid triangle
        }
        return 0;
    }
};

// Example main to run locally (optional)
int main() {
    vector<int> nums = {2, 1, 2};
    Solution sol;
    cout << sol.largestPerimeter(nums) << endl; // prints 5
    return 0;
}
```

---

### Java

```java
import java.util.Arrays;

class Solution {
    public int largestPerimeter(int[] nums) {
        Arrays.sort(nums);                         // O(n log n)
        for (int i = nums.length - 1; i >= 2; --i) {
            int c = nums[i];
            int b = nums[i - 1];
            int a = nums[i - 2];
            if ((long)a + b > c) return a + b + c; // cast to long safe-guard (optional)
        }
        return 0;
    }

    // Example main to run locally (optional)
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.largestPerimeter(new int[]{2,1,2})); // prints 5
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
var largestPerimeter = function(nums) {
    nums.sort((x, y) => x - y);                 // O(n log n)
    for (let i = nums.length - 1; i >= 2; --i) {
        const c = nums[i];
        const b = nums[i - 1];
        const a = nums[i - 2];
        if (a + b > c) return a + b + c;
    }
    return 0;
};

// Example usage:
// console.log(largestPerimeter([2,1,2])); // 5
```

---

### Python3

```python
from typing import List

class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        nums.sort()                                # O(n log n)
        for i in range(len(nums) - 1, 1, -1):
            c = nums[i]
            b = nums[i - 1]
            a = nums[i - 2]
            if a + b > c:
                return a + b + c
        return 0

# Example usage:
# sol = Solution()
# print(sol.largestPerimeter([2,1,2]))  # 5
```

---

### Go

```go
package main

import (
    "fmt"
    "sort"
)

func largestPerimeter(nums []int) int {
    sort.Ints(nums)                            // O(n log n)
    n := len(nums)
    for i := n - 1; i >= 2; i-- {
        c := nums[i]
        b := nums[i-1]
        a := nums[i-2]
        if a + b > c {
            return a + b + c
        }
    }
    return 0
}

func main() {
    fmt.Println(largestPerimeter([]int{2,1,2})) // prints 5
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the common logic once (same for all languages), then mention language-specific notes.

### Core idea (common)

1. **Sort the array** — I sort `nums` in ascending order so the largest elements are at the end. This makes it straightforward to consider the largest possible triples first.
2. **Iterate from the end** — For each index `i` starting at `len(nums)-1` and moving down, treat `c = nums[i]` as the current largest side, `b = nums[i-1]`, `a = nums[i-2]`.
3. **Check triangle inequality** — If `a + b > c`, then `(a, b, c)` can form a non-degenerate triangle. Because I'm iterating from the largest sides down, the first successful triple has the largest perimeter.
4. **Return perimeter or 0** — If a valid triple is found, return `a + b + c`. If after checking all triples none is valid, return `0`.

### Line-by-line style explanation (example from Python; same mapping applies to other languages)

```python
nums.sort()
```

* I sort the list in-place. Complexity: `O(n log n)`.

```python
for i in range(len(nums) - 1, 1, -1):
```

* Start with the last index and move backwards. Each `i` indicates the largest side `c` for the current triple.

```python
c = nums[i]
b = nums[i - 1]
a = nums[i - 2]
```

* Extract the three consecutive sides from the sorted list. Since the array is sorted in ascending order, `a <= b <= c`.

```python
if a + b > c:
    return a + b + c
```

* Check the triangle condition. If true, return the sum (perimeter). We return immediately because this is the largest possible perimeter among the remaining triples.

```python
return 0
```

* If no valid triple was found, no triangle with non-zero area can be formed, so return `0`.

### Language-specific notes

* **C++**: Use `sort` from `<algorithm>`. Works in-place on `vector<int>`. Use `int` — constraints fit 32-bit signed range (`3 * 10^6` max perimeter).
* **Java**: `Arrays.sort` sorts in-place. If you are paranoid about overflow (not needed here due to constraints), you can cast to `long` when checking sums.
* **JavaScript**: `array.sort((x,y)=>x-y)` to ensure numeric sort. Return value is number.
* **Python**: `list.sort()` sorts in-place. Simple and concise.
* **Go**: `sort.Ints(nums)` sorts in-place. Function returns `int`.

---

## Examples

1. Example 1
   Input: `nums = [2,1,2]`
   Sorted: `[1,2,2]`
   Check triple `(1,2,2)` -> `1 + 2 > 2` -> true -> Perimeter `5`
   Output: `5`

2. Example 2
   Input: `nums = [1,2,1,10]`
   Sorted: `[1,1,2,10]`
   Triples checked: `(1,2,10)`, `(1,1,2)` -> both invalid
   Output: `0`

---

## How to use / Run locally

Create a file with the corresponding code block and run as shown:

**C++ (g++):**

```bash
g++ -std=c++17 -O2 largest_perimeter.cpp -o largest_perimeter
./largest_perimeter
```

**Java:**

```bash
javac Solution.java
java Solution
```

**JavaScript (Node.js):**

```bash
node solution.js
```

**Python3:**

```bash
python3 solution.py
```

**Go:**

```bash
go run solution.go
```

(Each example snippet in this README includes a small `main` / usage example you can copy into a file to test.)

---

## Notes & Optimizations

* Because we only need three sides, sorting and scanning consecutive triples is optimal and simple.
* The first valid triple found when scanning from largest to smallest gives the maximum perimeter (greedy correctness).
* Space usage is minimal — we sort in-place.
* Integer overflow is not a concern under the given constraints (`nums[i] <= 10^6` and max perimeter ≤ `3 * 10^6`), but if constraints increased, check sums with a larger integer type (e.g., `long` in Java/C++).
* Average runtime may be less than `O(n log n)` for arrays where a valid triple occurs among the top elements (we return early).

---

## Author

[Aarzoo Islam](https://bento.me/withaarzoo)
