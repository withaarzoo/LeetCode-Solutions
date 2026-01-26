# 1200. Minimum Absolute Difference

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given an array of **distinct integers**.
Your task is to find **all pairs of elements** that have the **minimum absolute difference** among all possible pairs.

Each returned pair must:

* Contain elements from the array
* Be in ascending order `[a, b]` where `a < b`
* Have a difference equal to the minimum absolute difference
* Be returned in ascending order

---

## Constraints

* `2 ≤ arr.length ≤ 10⁵`
* `-10⁶ ≤ arr[i] ≤ 10⁶`
* All elements in `arr` are **distinct**

---

## Intuition

When I first looked at this problem, my main goal was to avoid checking every possible pair because that would be too slow.

I realized something important:

The **minimum absolute difference** can only happen between **numbers that are close to each other**.

So I thought:

* If I **sort the array**, then nearby elements will sit next to each other.
* After sorting, I only need to compare **consecutive elements**.

This single idea reduces the problem from brute force to an efficient solution.

---

## Approach

I solved the problem in these simple steps:

1. Sort the array in ascending order.
2. Traverse the array once to find the **minimum difference** between consecutive elements.
3. Traverse the array again and collect all pairs whose difference equals that minimum.
4. Return the result.

Because the array is sorted, the output is already in the correct order.

---

## Data Structures Used

* Array / List (for input and output)
* No extra complex data structures are required

---

## Operations & Behavior Summary

* Sorting arranges numbers in increasing order
* Consecutive comparison finds the smallest gap
* Matching pairs are stored in a result list
* Output is automatically ordered

---

## Complexity

**Time Complexity:** `O(n log n)`

* `n` is the number of elements
* Sorting dominates the runtime

**Space Complexity:** `O(1)` (excluding output)

* Only constant extra variables are used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<vector<int>> minimumAbsDifference(vector<int>& arr) {
        sort(arr.begin(), arr.end());

        int minDiff = INT_MAX;
        vector<vector<int>> result;

        for (int i = 1; i < arr.size(); i++) {
            minDiff = min(minDiff, arr[i] - arr[i - 1]);
        }

        for (int i = 1; i < arr.size(); i++) {
            if (arr[i] - arr[i - 1] == minDiff) {
                result.push_back({arr[i - 1], arr[i]});
            }
        }

        return result;
    }
};
```

---

### Java

```java
class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        Arrays.sort(arr);

        int minDiff = Integer.MAX_VALUE;
        List<List<Integer>> result = new ArrayList<>();

        for (int i = 1; i < arr.length; i++) {
            minDiff = Math.min(minDiff, arr[i] - arr[i - 1]);
        }

        for (int i = 1; i < arr.length; i++) {
            if (arr[i] - arr[i - 1] == minDiff) {
                result.add(Arrays.asList(arr[i - 1], arr[i]));
            }
        }

        return result;
    }
}
```

---

### JavaScript

```javascript
var minimumAbsDifference = function(arr) {
    arr.sort((a, b) => a - b);

    let minDiff = Infinity;
    let result = [];

    for (let i = 1; i < arr.length; i++) {
        minDiff = Math.min(minDiff, arr[i] - arr[i - 1]);
    }

    for (let i = 1; i < arr.length; i++) {
        if (arr[i] - arr[i - 1] === minDiff) {
            result.push([arr[i - 1], arr[i]]);
        }
    }

    return result;
};
```

---

### Python3

```python
class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()

        min_diff = float('inf')
        result = []

        for i in range(1, len(arr)):
            min_diff = min(min_diff, arr[i] - arr[i - 1])

        for i in range(1, len(arr)):
            if arr[i] - arr[i - 1] == min_diff:
                result.append([arr[i - 1], arr[i]])

        return result
```

---

### Go

```go
func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

    minDiff := math.MaxInt32
    result := [][]int{}

    for i := 1; i < len(arr); i++ {
        diff := arr[i] - arr[i-1]
        if diff < minDiff {
            minDiff = diff
        }
    }

    for i := 1; i < len(arr); i++ {
        if arr[i]-arr[i-1] == minDiff {
            result = append(result, []int{arr[i-1], arr[i]})
        }
    }

    return result
}
```

---

## Step-by-step Detailed Explanation

1. Sorting brings all numbers into order.
2. The smallest difference must be between adjacent elements.
3. First loop finds the minimum difference.
4. Second loop collects all pairs matching that difference.
5. The result is returned in ascending order automatically.

---

## Examples

**Input**

```bash
arr = [4,2,1,3]
```

**Output**

```bash
[[1,2],[2,3],[3,4]]
```

---

**Input**

```bash
arr = [1,3,6,10,15]
```

**Output**

```bash
[[1,3]]
```

---

## How to use / Run locally

1. Copy the code for your preferred language
2. Paste it into your local environment or LeetCode editor
3. Run the program with test inputs
4. Verify output

---

## Notes & Optimizations

* Brute force `O(n²)` is avoided
* Sorting is the key optimization
* Works efficiently for large inputs up to `10⁵`
* Clean, readable, and interview-friendly

---

## Author

* **Md Aarzoo Islam**
  🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
