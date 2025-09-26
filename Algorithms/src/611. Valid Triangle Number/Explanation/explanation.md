# 611. Valid Triangle Number

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#cpp)
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

Given an integer array `nums`, return the number of triplets chosen from the array that can form triangles if taken as side lengths. In other words, count the number of index triplets `(i, j, k)` with `i < j < k` such that `nums[i]`, `nums[j]`, `nums[k]` can be side lengths of a triangle.

---

## Constraints

* `1 <= nums.length <= 1000`
* `0 <= nums[i] <= 1000`

---

## Intuition

I thought about the triangle inequality: for three numbers `a, b, c` with `a ≤ b ≤ c`, they form a triangle iff `a + b > c`. Sorting makes this easier.

---

## Approach

1. Sort `nums` ascending.
2. Pick `nums[k]` as the largest side.
3. Use two pointers (`i` left, `j` right) to count pairs that satisfy `nums[i] + nums[j] > nums[k]`.
4. Return the total count.

---

## Data Structures Used

* Array (sorted in-place).
* Two pointers (`i`, `j`) and a counter.

---

## Operations & Behavior Summary

* Sorting to enable ordered comparisons.
* Two-pointer scan for counting valid pairs.
* Avoids triple nested loops.

---

## Complexity

* **Time Complexity:** `O(n^2)` where `n` is array length.
* **Space Complexity:** `O(1)` extra space.

---

## Multi-language Solutions

### Cpp

```c++
#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int triangleNumber(vector<int>& nums) {
        int n = nums.size();
        if (n < 3) return 0;
        sort(nums.begin(), nums.end());
        int count = 0;
        for (int k = n - 1; k >= 2; --k) {
            int i = 0, j = k - 1;
            while (i < j) {
                if (nums[i] + nums[j] > nums[k]) {
                    count += j - i;
                    --j;
                } else {
                    ++i;
                }
            }
        }
        return count;
    }
};
```

### Java

```java
import java.util.Arrays;

class Solution {
    public int triangleNumber(int[] nums) {
        int n = nums.length;
        if (n < 3) return 0;
        Arrays.sort(nums);
        int count = 0;
        for (int k = n - 1; k >= 2; k--) {
            int i = 0, j = k - 1;
            while (i < j) {
                if (nums[i] + nums[j] > nums[k]) {
                    count += j - i;
                    j--;
                } else {
                    i++;
                }
            }
        }
        return count;
    }
}
```

### JavaScript

```javascript
var triangleNumber = function(nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    let count = 0;
    for (let k = n - 1; k >= 2; k--) {
        let i = 0, j = k - 1;
        while (i < j) {
            if (nums[i] + nums[j] > nums[k]) {
                count += j - i;
                j--;
            } else {
                i++;
            }
        }
    }
    return count;
};
```

### Python3

```python
from typing import List

class Solution:
    def triangleNumber(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 3:
            return 0
        nums.sort()
        count = 0
        for k in range(n - 1, 1, -1):
            i, j = 0, k - 1
            while i < j:
                if nums[i] + nums[j] > nums[k]:
                    count += j - i
                    j -= 1
                else:
                    i += 1
        return count
```

### Go

```go
import "sort"

func triangleNumber(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    sort.Ints(nums)
    count := 0
    for k := n - 1; k >= 2; k-- {
        i, j := 0, k - 1
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                count += j - i
                j--
            } else {
                i++
            }
        }
    }
    return count
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

* **Step 1:** Sort the array. This ensures that if the smallest `i` works with `j`, then all bigger numbers between `i` and `j-1` will also work.
* **Step 2:** Pick `k` as the largest side.
* **Step 3:** Move two pointers inside the range `[0..k-1]`.

  * If `nums[i] + nums[j] > nums[k]`, add `j-i` to result.
  * If not, increment `i`.
* **Step 4:** Continue until all possible `k` are checked.

---

## Examples

* Example 1: `nums = [2,2,3,4]` → Output: `3`.
* Example 2: `nums = [4,2,3,4]` → Output: `4`.

---

## How to use / Run locally

* Copy the solution into a file in your language of choice.
* Compile or run as shown in examples (e.g., `python3 file.py`, `g++ file.cpp`, `javac Solution.java`, `go run main.go`).
* Run the function/class with test arrays.

---

## Notes & Optimizations

* Sorting + two pointers is the core optimization.
* Brute force `O(n^3)` fails for `n=1000`.
* Edge cases with zeros or duplicates are naturally handled.

---

## Author

[Aarzoo](https://bento.me/withaarzoo)
