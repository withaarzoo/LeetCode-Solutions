# Minimum Operations to Make a Uni-Value Grid (LeetCode 2033)

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

In this problem, we are given a 2D grid of integers and a number `x`.
In one operation, we can either add `x` or subtract `x` from any element in the grid.

Our goal is to make all elements in the grid equal using the minimum number of operations.

If it is not possible to make all elements equal, we must return `-1`.

This is a classic grid transformation problem that involves understanding number patterns and optimizing operations.

---

## Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 10^5`
* `1 <= m * n <= 10^5`
* `1 <= x, grid[i][j] <= 10^4`

---

## Intuition

When I first looked at the problem, I realized something important.

Since I can only add or subtract `x`, I can only move numbers in steps of size `x`.

So I asked myself:
Can all numbers even reach the same value?

The key idea:
If two numbers give different remainders when divided by `x`, then they can never become equal.

Once that condition is satisfied, the problem becomes:
Find a number such that the total number of operations needed to convert all elements into it is minimum.

This is where the median comes into play.

In many problems involving minimizing total distance, the median gives the optimal answer.

---

## Approach

Here is how I solved it step by step:

1. Convert the 2D grid into a 1D list
2. Check if all elements have the same remainder when divided by `x`

   * If not, return `-1`
3. Sort the list
4. Choose the median element as the target value
5. Calculate the total number of operations needed to convert each element into the median
6. Return the total operations

---

## Data Structures Used

* **Array / List**

  * Used to store all grid elements in a flattened format
  * Makes sorting and traversal easier

* **Sorting**

  * Required to find the median efficiently

---

## Operations & Behavior Summary

* Flatten the grid into a single list
* Check feasibility using modulo operation
* Sort the list
* Select the median
* Calculate total operations using absolute difference divided by `x`
* Return result

---

## Complexity

| Type             | Complexity |
| ---------------- | ---------- |
| Time Complexity  | O(n log n) |
| Space Complexity | O(n)       |

**Explanation:**

* `n` is the total number of elements in the grid
* Sorting takes O(n log n)
* Extra space is used to store the flattened array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minOperations(vector<vector<int>>& grid, int x) {
        vector<int> nums;

        // Step 1: Flatten grid
        for (auto &row : grid) {
            for (int val : row) {
                nums.push_back(val);
            }
        }

        // Step 2: Check feasibility
        int rem = nums[0] % x;
        for (int num : nums) {
            if (num % x != rem) return -1;
        }

        // Step 3: Sort
        sort(nums.begin(), nums.end());

        // Step 4: Median
        int median = nums[nums.size() / 2];

        // Step 5: Calculate operations
        int ops = 0;
        for (int num : nums) {
            ops += abs(num - median) / x;
        }

        return ops;
    }
};
```

### Java

```java
class Solution {
    public int minOperations(int[][] grid, int x) {
        List<Integer> nums = new ArrayList<>();

        // Flatten grid
        for (int[] row : grid) {
            for (int val : row) {
                nums.add(val);
            }
        }

        // Check feasibility
        int rem = nums.get(0) % x;
        for (int num : nums) {
            if (num % x != rem) return -1;
        }

        // Sort
        Collections.sort(nums);

        // Median
        int median = nums.get(nums.size() / 2);

        // Count operations
        int ops = 0;
        for (int num : nums) {
            ops += Math.abs(num - median) / x;
        }

        return ops;
    }
}
```

### JavaScript

```javascript
var minOperations = function(grid, x) {
    let nums = [];

    // Flatten grid
    for (let row of grid) {
        for (let val of row) {
            nums.push(val);
        }
    }

    // Check feasibility
    let rem = nums[0] % x;
    for (let num of nums) {
        if (num % x !== rem) return -1;
    }

    // Sort
    nums.sort((a, b) => a - b);

    // Median
    let median = nums[Math.floor(nums.length / 2)];

    // Count operations
    let ops = 0;
    for (let num of nums) {
        ops += Math.abs(num - median) / x;
    }

    return ops;
};
```

### Python3

```python
class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        nums = []

        # Flatten grid
        for row in grid:
            nums.extend(row)

        # Check feasibility
        rem = nums[0] % x
        for num in nums:
            if num % x != rem:
                return -1

        # Sort
        nums.sort()

        # Median
        median = nums[len(nums) // 2]

        # Count operations
        ops = 0
        for num in nums:
            ops += abs(num - median) // x

        return ops
```

### Go

```go
func minOperations(grid [][]int, x int) int {
    nums := []int{}

    // Flatten grid
    for _, row := range grid {
        for _, val := range row {
            nums = append(nums, val)
        }
    }

    // Check feasibility
    rem := nums[0] % x
    for _, num := range nums {
        if num%x != rem {
            return -1
        }
    }

    // Sort
    sort.Ints(nums)

    // Median
    median := nums[len(nums)/2]

    // Count operations
    ops := 0
    for _, num := range nums {
        if num > median {
            ops += (num - median) / x
        } else {
            ops += (median - num) / x
        }
    }

    return ops
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

First, I take all elements from the 2D grid and store them in a single list.
This simplifies processing and avoids dealing with row and column indices separately.

Next, I check if making the grid uni-value is even possible.
I pick the first element and compute its remainder when divided by `x`.

Then I compare this remainder with every other element in the list.
If any element has a different remainder, I immediately return `-1`.

If all elements pass this check, I move forward.

Now I sort the list.
Sorting helps me find the median easily.

I pick the middle element as the target value.

Then I go through each element and calculate how many steps it takes to reach the median.
Each step changes the value by `x`, so I divide the difference by `x`.

Finally, I add all these steps together and return the result.

This logic works the same way in all programming languages mentioned above.

---

## Examples

### Example 1

Input:

```
grid = [[2,4],[6,8]], x = 2
```

Output:

```
4
```

Explanation:

* Target value = 4
* Steps:

  * 2 → 4 (1 step)
  * 6 → 4 (1 step)
  * 8 → 4 (2 steps)
* Total = 4

---

### Example 2

Input:

```
grid = [[1,5],[2,3]], x = 1
```

Output:

```
5
```

Explanation:

* Target value = 3
* Calculate total steps to reach 3

---

### Example 3

Input:

```
grid = [[1,2],[3,4]], x = 2
```

Output:

```
-1
```

Explanation:

* Remainders differ when divided by `x`
* Impossible to make all elements equal

---

## How to Use / Run Locally

### C++

1. Save the file as `solution.cpp`
2. Compile:

   ```
   g++ solution.cpp -o solution
   ```

3. Run:

   ```
   ./solution
   ```

---

### Java

1. Save as `Solution.java`
2. Compile:

   ```
   javac Solution.java
   ```

3. Run:

   ```
   java Solution
   ```

---

### JavaScript

1. Save as `solution.js`
2. Run:

   ```
   node solution.js
   ```

---

### Python

1. Save as `solution.py`
2. Run:

   ```
   python solution.py
   ```

---

### Go

1. Save as `solution.go`
2. Run:

   ```
   go run solution.go
   ```

---

## Notes & Optimizations

* Always check feasibility first using modulo
* Median gives minimum operations in absolute difference problems
* Sorting is necessary unless using advanced selection algorithms
* Works efficiently within given constraints
* Avoid unnecessary extra space if possible

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
