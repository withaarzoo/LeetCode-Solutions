# 1984. Minimum Difference Between Highest and Lowest of K Scores

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

I am given an integer array `nums` where each value represents a student’s score.
I am also given an integer `k`.

My task is to **select exactly `k` students** such that the **difference between the highest and lowest score among them is minimum**.

Finally, I return that **minimum possible difference**.

---

## Constraints

* `1 <= k <= nums.length <= 1000`
* `0 <= nums[i] <= 10^5`

---

## Intuition

When I read the problem, I understood that the difference will be smallest when the selected scores are **as close to each other as possible**.

So I thought:

* If I **sort the array**, similar scores will come together.
* Then, if I pick **any `k` consecutive elements**, their max and min difference will be minimal compared to random picks.

This idea simplified the whole problem.

---

## Approach

1. If `k == 1`, I directly return `0` because highest and lowest will be the same.
2. I sort the `nums` array.
3. I slide a window of size `k` across the sorted array.
4. For each window:

   * Lowest score = `nums[i]`
   * Highest score = `nums[i + k - 1]`
   * Difference = highest − lowest
5. I keep track of the **minimum difference**.
6. I return the final answer.

---

## Data Structures Used

* Array (for storing scores)
* No extra data structure is required

---

## Operations & Behavior Summary

* Sorting groups similar scores together
* Sliding window checks only valid `k`-sized groups
* Ensures optimal solution without checking all combinations

---

## Complexity

**Time Complexity:**
`O(n log n)`

* `n` = number of students
* Sorting dominates the runtime

**Space Complexity:**
`O(1)`

* No extra space used (ignoring sorting space)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDifference(vector<int>& nums, int k) {
        if (k == 1) return 0;

        sort(nums.begin(), nums.end());
        int minDiff = INT_MAX;

        for (int i = 0; i + k - 1 < nums.size(); i++) {
            minDiff = min(minDiff, nums[i + k - 1] - nums[i]);
        }
        return minDiff;
    }
};
```

---

### Java

```java
class Solution {
    public int minimumDifference(int[] nums, int k) {
        if (k == 1) return 0;

        Arrays.sort(nums);
        int minDiff = Integer.MAX_VALUE;

        for (int i = 0; i + k - 1 < nums.length; i++) {
            minDiff = Math.min(minDiff, nums[i + k - 1] - nums[i]);
        }
        return minDiff;
    }
}
```

---

### JavaScript

```javascript
var minimumDifference = function(nums, k) {
    if (k === 1) return 0;

    nums.sort((a, b) => a - b);
    let minDiff = Infinity;

    for (let i = 0; i + k - 1 < nums.length; i++) {
        minDiff = Math.min(minDiff, nums[i + k - 1] - nums[i]);
    }
    return minDiff;
};
```

---

### Python3

```python
class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        if k == 1:
            return 0

        nums.sort()
        min_diff = float('inf')

        for i in range(len(nums) - k + 1):
            min_diff = min(min_diff, nums[i + k - 1] - nums[i])

        return min_diff
```

---

### Go

```go
func minimumDifference(nums []int, k int) int {
    if k == 1 {
        return 0
    }

    sort.Ints(nums)
    minDiff := math.MaxInt32

    for i := 0; i + k - 1 < len(nums); i++ {
        diff := nums[i + k - 1] - nums[i]
        if diff < minDiff {
            minDiff = diff
        }
    }
    return minDiff
}
```

---

## Step-by-step Detailed Explanation (All Languages)

1. I check if `k == 1`

   * One score means no difference.
2. I sort the array

   * This makes nearby scores closer.
3. I use a sliding window of size `k`

   * Each window represents one valid selection.
4. I calculate:

   * `difference = max - min`
5. I store the minimum difference found.
6. I return the answer.

This logic is same in **C++, Java, JavaScript, Python, and Go**.

---

## Examples

**Example 1**

```bash
Input: nums = [90], k = 1
Output: 0
```

**Example 2**

```bash
Input: nums = [9,4,1,7], k = 2
Output: 2
```

---

## How to use / Run locally

1. Copy the code for your preferred language.
2. Paste it into your local compiler or LeetCode editor.
3. Run with sample inputs to test.
4. Submit on LeetCode.

---

## Notes & Optimizations

* Sorting is unavoidable for optimal solution.
* Sliding window avoids unnecessary combinations.
* Works efficiently within given constraints.
* Very common **interview sliding window + sorting** pattern.

---

## Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
