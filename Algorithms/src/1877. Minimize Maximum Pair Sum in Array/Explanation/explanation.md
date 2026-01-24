# 1877. Minimize Maximum Pair Sum in Array

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given an integer array `nums` of **even length**.

You need to divide the array into `n / 2` pairs such that:

* Every element is used exactly once
* The **maximum pair sum** is as small as possible

A pair sum is defined as:

```bash
pair sum = a + b
```

Your task is to return the **minimum possible value of the maximum pair sum** after pairing the elements optimally.

---

## Constraints

* `2 <= nums.length <= 10^5`
* `nums.length` is even
* `1 <= nums[i] <= 10^5`

---

## Intuition

When I read this problem, the main thing I focused on was **controlling the largest pair sum**.

If I pair two large numbers together, the sum becomes very large. That is bad.

So I thought:

* What if I pair the **largest number with the smallest number**?
* This balances the pair and prevents any single pair from becoming too large.

By doing this for all elements, I can keep the maximum pair sum as small as possible.

---

## Approach

1. Sort the array in ascending order
2. Use two pointers:

   * One at the start (smallest element)
   * One at the end (largest element)
3. Pair the smallest and largest elements
4. Track the maximum pair sum
5. Move both pointers inward
6. Repeat until all elements are paired

Finally, return the maximum pair sum found.

---

## Data Structures Used

* Array (input array)
* Two pointers (integer indices)

No extra data structures are required.

---

## Operations & Behavior Summary

* Sorting helps arrange numbers from smallest to largest
* Pairing smallest with largest balances sums
* Greedy strategy ensures optimal result
* Only one pass after sorting

---

## Complexity

* **Time Complexity:** `O(n log n)`

  * Sorting the array of `n` elements

* **Space Complexity:** `O(1)`

  * No extra space used (sorting is in-place)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minPairSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int left = 0, right = nums.size() - 1;
        int maxPairSum = 0;

        while (left < right) {
            maxPairSum = max(maxPairSum, nums[left] + nums[right]);
            left++;
            right--;
        }
        return maxPairSum;
    }
};
```

---

### Java

```java
class Solution {
    public int minPairSum(int[] nums) {
        Arrays.sort(nums);
        int left = 0, right = nums.length - 1;
        int maxPairSum = 0;

        while (left < right) {
            maxPairSum = Math.max(maxPairSum, nums[left] + nums[right]);
            left++;
            right--;
        }
        return maxPairSum;
    }
}
```

---

### JavaScript

```javascript
var minPairSum = function(nums) {
    nums.sort((a, b) => a - b);
    let left = 0, right = nums.length - 1;
    let maxPairSum = 0;

    while (left < right) {
        maxPairSum = Math.max(maxPairSum, nums[left] + nums[right]);
        left++;
        right--;
    }
    return maxPairSum;
};
```

---

### Python3

```python
class Solution:
    def minPairSum(self, nums: List[int]) -> int:
        nums.sort()
        left, right = 0, len(nums) - 1
        max_pair_sum = 0

        while left < right:
            max_pair_sum = max(max_pair_sum, nums[left] + nums[right])
            left += 1
            right -= 1
        return max_pair_sum
```

---

### Go

```go
func minPairSum(nums []int) int {
    sort.Ints(nums)
    left, right := 0, len(nums)-1
    maxPairSum := 0

    for left < right {
        sum := nums[left] + nums[right]
        if sum > maxPairSum {
            maxPairSum = sum
        }
        left++
        right--
    }
    return maxPairSum
}
```

---

## Step-by-step Detailed Explanation

1. Sort the array so numbers are in increasing order
2. Pair the smallest number with the largest number
3. Calculate their sum
4. Store the maximum sum encountered
5. Move pointers inward and repeat
6. After all pairs are processed, return the stored maximum

This guarantees the smallest possible maximum pair sum.

---

## Examples

### Example 1

```bash
Input: nums = [3,5,2,3]
Output: 7
Explanation:
Pairs: (2,5), (3,3)
Maximum pair sum = 7
```

### Example 2

```bash
Input: nums = [3,5,4,2,4,6]
Output: 8
Explanation:
Pairs: (2,6), (3,5), (4,4)
Maximum pair sum = 8
```

---

## How to Use / Run Locally

1. Copy the solution code for your preferred language
2. Paste it into your local compiler or LeetCode editor
3. Provide the input array
4. Run the program to get the result

---

## Notes & Optimizations

* Greedy pairing is the key idea
* Sorting is unavoidable for optimal pairing
* Two-pointer technique makes solution clean and fast
* Works efficiently even for large inputs

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
