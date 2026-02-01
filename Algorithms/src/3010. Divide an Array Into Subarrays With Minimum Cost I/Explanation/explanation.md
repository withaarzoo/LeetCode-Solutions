# 3010. Divide an Array Into Subarrays With Minimum Cost I

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an integer array `nums` of length `n`.

The cost of a subarray is defined as the **first element** of that subarray.

My task is to divide the array into **exactly 3 disjoint contiguous subarrays** and return the **minimum possible sum of the costs** of those 3 subarrays.

---

## Constraints

* 3 ≤ n ≤ 50
* 1 ≤ nums[i] ≤ 50

---

## Intuition

When I read the problem carefully, one thing became very clear.

The first subarray **must start from index 0**.
So the cost will **always include `nums[0]`**.

I cannot change that.

Now I only need to choose the starting points of the **second and third subarrays**.

Since the cost of a subarray is its **first element**, I should choose the **smallest possible values** as the starting elements of the second and third subarrays.

So my goal becomes very simple
Take `nums[0]`
From the remaining elements, pick the **two smallest numbers**
Add them together

That will always give the minimum cost.

---

## Approach

1. I take `nums[0]` because it is always included in the answer.
2. I look at the remaining part of the array from index `1` to `n-1`.
3. I sort this remaining part.
4. I take the **two smallest values** from it.
5. I add them with `nums[0]`.
6. That sum is my final answer.

---

## Data Structures Used

* Array
* Sorting (built-in sort functions)

No extra complex data structures are needed.

---

## Operations & Behavior Summary

* First subarray always starts at index 0.
* Cost depends only on the first element of each subarray.
* Picking smaller starting elements reduces total cost.
* Sorting helps quickly find the two smallest remaining values.

---

## Complexity

**Time Complexity**
O(n log n)
Because I sort the array elements (except the first one).

**Space Complexity**
O(1)
Only a few variables are used. No extra memory is required.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumCost(vector<int>& nums) {
        int first = nums[0];
        sort(nums.begin() + 1, nums.end());
        return first + nums[1] + nums[2];
    }
};
```

---

### Java

```java
class Solution {
    public int minimumCost(int[] nums) {
        int first = nums[0];
        Arrays.sort(nums, 1, nums.length);
        return first + nums[1] + nums[2];
    }
}
```

---

### JavaScript

```javascript
var minimumCost = function(nums) {
    let first = nums[0];
    let rest = nums.slice(1).sort((a, b) => a - b);
    return first + rest[0] + rest[1];
};
```

---

### Python3

```python
class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        first = nums[0]
        nums[1:] = sorted(nums[1:])
        return first + nums[1] + nums[2]
```

---

### Go

```go
func minimumCost(nums []int) int {
    first := nums[0]
    sort.Ints(nums[1:])
    return first + nums[1] + nums[2]
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Let’s understand this with a simple example.

Input
`nums = [1, 2, 3, 12]`

Step 1
The first subarray must start from index 0
Cost includes `1`

Step 2
Remaining elements
`[2, 3, 12]`

Step 3
Pick the two smallest values
`2` and `3`

Step 4
Add them
`1 + 2 + 3 = 6`

That is the minimum possible cost.

This logic is exactly the same in all languages.
Only syntax changes, thinking remains the same.

---

## Examples

**Example 1**

Input
`[1, 2, 3, 12]`

Output
`6`

---

**Example 2**

Input
`[5, 4, 3]`

Output
`12`

---

**Example 3**

Input
`[10, 3, 1, 1]`

Output
`12`

---

## How to use / Run locally

1. Copy the code for your preferred language.
2. Paste it into your local editor or online compiler.
3. Call the `minimumCost` function with input array.
4. Run and verify the output.

For LeetCode

* Just paste the code inside the `Solution` class and submit.

---

## Notes & Optimizations

* This problem looks like DP but **does not need DP**.
* Simple greedy logic works perfectly.
* Sorting makes the solution clean and easy to understand.
* Constraints are small, so performance is not an issue.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
