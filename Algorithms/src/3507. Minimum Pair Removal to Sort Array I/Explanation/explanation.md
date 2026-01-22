# Minimum Pair Removal to Sort Array I

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

I am given an integer array `nums`.

I can repeat the following operation any number of times:

* Select the **adjacent pair with the minimum sum**
* If multiple pairs have the same sum, choose the **leftmost one**
* Replace the pair with their **sum**

My task is to return the **minimum number of operations** required to make the array **non-decreasing**.

An array is called non-decreasing if every element is **greater than or equal to** the previous element.

---

## Constraints

* `1 ≤ nums.length ≤ 50`
* `-1000 ≤ nums[i] ≤ 1000`

---

## Intuition

When I first read the problem, I noticed something very important.

I do **not** have freedom to remove any pair I want.
The problem strictly forces me to remove **only the adjacent pair with the minimum sum**.

So I stopped thinking about greedy tricks or math shortcuts.

Because the array size is very small, I realized the safest and cleanest solution is to **simulate exactly what the problem says**.

If I keep applying the given operation until the array becomes non-decreasing, the number of operations I perform is automatically the minimum.

---

## Approach

Here is how I solved the problem step by step:

1. First, I check if the array is already non-decreasing.
2. If it is already sorted, I return `0`.
3. Otherwise, I repeat the following:

   * Scan all adjacent pairs
   * Find the pair with the minimum sum
   * If multiple exist, the leftmost one is picked automatically
4. Replace that pair with their sum.
5. Increase the operation counter.
6. Repeat until the array becomes non-decreasing.
7. Return the total number of operations.

Because the array size keeps shrinking, this process always ends.

---

## Data Structures Used

* Array / List only
* No extra complex data structures required

---

## Operations & Behavior Summary

* Operation always removes **exactly one element**
* Array size decreases by `1` each time
* Order of remaining elements is preserved
* Process is deterministic (no randomness)

---

## Complexity

* **Time Complexity:** `O(n³)`

  * Checking sorted: `O(n)`
  * Finding minimum pair: `O(n)`
  * Maximum operations: `O(n)`
  * `n ≤ 50`, so this is fully acceptable

* **Space Complexity:** `O(1)`

  * Only in-place modifications
  * No extra memory used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumPairRemoval(vector<int>& nums) {
        int ops = 0;

        auto isSorted = [&]() {
            for (int i = 1; i < nums.size(); i++) {
                if (nums[i] < nums[i - 1]) return false;
            }
            return true;
        };

        while (!isSorted()) {
            int minSum = INT_MAX, idx = 0;

            for (int i = 0; i + 1 < nums.size(); i++) {
                int s = nums[i] + nums[i + 1];
                if (s < minSum) {
                    minSum = s;
                    idx = i;
                }
            }

            nums[idx] = minSum;
            nums.erase(nums.begin() + idx + 1);
            ops++;
        }

        return ops;
    }
};
```

---

### Java

```java
class Solution {
    public int minimumPairRemoval(int[] nums) {
        List<Integer> list = new ArrayList<>();
        for (int n : nums) list.add(n);

        int ops = 0;

        while (!isSorted(list)) {
            int minSum = Integer.MAX_VALUE;
            int idx = 0;

            for (int i = 0; i < list.size() - 1; i++) {
                int s = list.get(i) + list.get(i + 1);
                if (s < minSum) {
                    minSum = s;
                    idx = i;
                }
            }

            list.set(idx, minSum);
            list.remove(idx + 1);
            ops++;
        }

        return ops;
    }

    private boolean isSorted(List<Integer> list) {
        for (int i = 1; i < list.size(); i++) {
            if (list.get(i) < list.get(i - 1)) return false;
        }
        return true;
    }
}
```

---

### JavaScript

```javascript
var minimumPairRemoval = function(nums) {
    let ops = 0;

    const isSorted = () => {
        for (let i = 1; i < nums.length; i++) {
            if (nums[i] < nums[i - 1]) return false;
        }
        return true;
    };

    while (!isSorted()) {
        let minSum = Infinity;
        let idx = 0;

        for (let i = 0; i < nums.length - 1; i++) {
            let s = nums[i] + nums[i + 1];
            if (s < minSum) {
                minSum = s;
                idx = i;
            }
        }

        nums[idx] = minSum;
        nums.splice(idx + 1, 1);
        ops++;
    }

    return ops;
};
```

---

### Python3

```python
class Solution:
    def minimumPairRemoval(self, nums):
        ops = 0

        def is_sorted():
            for i in range(1, len(nums)):
                if nums[i] < nums[i - 1]:
                    return False
            return True

        while not is_sorted():
            min_sum = float('inf')
            idx = 0

            for i in range(len(nums) - 1):
                s = nums[i] + nums[i + 1]
                if s < min_sum:
                    min_sum = s
                    idx = i

            nums[idx] = min_sum
            nums.pop(idx + 1)
            ops += 1

        return ops
```

---

### Go

```go
func minimumPairRemoval(nums []int) int {
    ops := 0

    isSorted := func() bool {
        for i := 1; i < len(nums); i++ {
            if nums[i] < nums[i-1] {
                return false
            }
        }
        return true
    }

    for !isSorted() {
        minSum := int(1e9)
        idx := 0

        for i := 0; i < len(nums)-1; i++ {
            s := nums[i] + nums[i+1]
            if s < minSum {
                minSum = s
                idx = i
            }
        }

        nums[idx] = minSum
        nums = append(nums[:idx+1], nums[idx+2:]...)
        ops++
    }

    return ops
}
```

---

## Step-by-step Detailed Explanation

1. I first check whether the array is already non-decreasing.
2. If yes, no operation is needed.
3. Otherwise, I scan all adjacent pairs.
4. I calculate their sums.
5. I choose the smallest sum pair.
6. I replace that pair with their sum.
7. The array becomes smaller.
8. I repeat until the array is sorted.
9. The number of repetitions is my answer.

---

## Examples

### Example 1

Input:

```bash
[5, 2, 3, 1]
```

Operations:

```bash
(3,1) → 4  → [5,2,4]
(2,4) → 6  → [5,6]
```

Output:

```bash
2
```

---

### Example 2

Input:

```bash
[1,2,2]
```

Output:

```bash
0
```

---

## How to Use / Run Locally

* Copy the solution code
* Paste it into your local compiler or online judge
* Run with sample inputs
* No special libraries required

---

## Notes & Optimizations

* Simulation is the safest approach due to strict rules
* Priority Queue optimization is unnecessary due to small constraints
* Code is clean, readable, and interview-friendly

---

## Author

* **Md Aarzoo Islam**
  *[https://bento.me/withaarzoo](https://bento.me/withaarzoo)*
