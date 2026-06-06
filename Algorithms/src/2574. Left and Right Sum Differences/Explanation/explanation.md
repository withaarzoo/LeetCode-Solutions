# 2574. Left and Right Sum Differences | Prefix Sum Solution in C++, Java, JavaScript, Python, and Go

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

Given a 0-indexed integer array `nums`, we need to create a new array where each position contains the absolute difference between:

* The sum of all elements to the left of the current index
* The sum of all elements to the right of the current index

For every index:

```text
answer[i] = |leftSum[i] - rightSum[i]|
```

If there are no elements on either side, that sum is considered `0`.

The goal is to return the resulting array efficiently.

This problem is a classic Prefix Sum problem on LeetCode and is a good exercise for learning running sums and array traversal techniques.

## Constraints

| Constraint    | Value                    |
| ------------- | ------------------------ |
| Array Length  | 1 <= nums.length <= 1000 |
| Element Value | 1 <= nums[i] <= 100000   |

## Intuition

My first thought was to calculate the left sum and right sum separately for every index.

That works, but it repeatedly scans parts of the array and wastes time.

Then I noticed that if I already know the total sum of the array, I can keep track of:

* A running left sum
* A running right sum

As I move through the array, I update both values and immediately compute the answer for the current position.

This avoids recalculating sums again and again.

## Approach

1. Calculate the total sum of the entire array.
2. Store it in a variable called `rightSum`.
3. Initialize `leftSum` as `0`.
4. Traverse the array from left to right.
5. Remove the current element from `rightSum`.
6. Now:

   * `leftSum` represents all elements on the left.
   * `rightSum` represents all elements on the right.
7. Compute the absolute difference.
8. Store the result.
9. Add the current element to `leftSum`.
10. Continue until all positions are processed.

This gives an efficient linear-time solution.

## Data Structures Used

| Data Structure    | Purpose                           |
| ----------------- | --------------------------------- |
| Array / Vector    | Stores the final answer           |
| Integer Variables | Track running left and right sums |

No additional complex data structures are needed.

## Operations & Behavior Summary

1. Find the total sum of all elements.
2. Start with an empty left side.
3. Move through the array one element at a time.
4. Remove the current element from the remaining right-side sum.
5. Compare left-side and right-side sums.
6. Store the absolute difference.
7. Add the current value to the left-side sum.
8. Repeat until the array ends.
9. Return the answer array.

## Complexity

| Metric           | Complexity       | Explanation                                                     |
| ---------------- | ---------------- | --------------------------------------------------------------- |
| Time Complexity  | O(n)             | The array is traversed a constant number of times               |
| Space Complexity | O(1) Extra Space | Only a few variables are used besides the required output array |

Where:

* `n` = length of the input array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> leftRightDifference(vector<int>& nums) {
        
        int n = nums.size();

        // Store the sum of all elements initially
        int rightSum = 0;
        for (int num : nums) {
            rightSum += num;
        }

        // Sum of elements on the left side
        int leftSum = 0;

        // Result array
        vector<int> ans(n);

        for (int i = 0; i < n; i++) {

            // Remove current element so rightSum becomes
            // the sum of elements strictly to the right
            rightSum -= nums[i];

            // Store absolute difference
            ans[i] = abs(leftSum - rightSum);

            // Add current element to leftSum for next index
            leftSum += nums[i];
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[] leftRightDifference(int[] nums) {
        
        int n = nums.length;

        // Store total array sum
        int rightSum = 0;
        for (int num : nums) {
            rightSum += num;
        }

        // Sum of elements on the left
        int leftSum = 0;

        // Result array
        int[] ans = new int[n];

        for (int i = 0; i < n; i++) {

            // Remove current element from right side sum
            rightSum -= nums[i];

            // Store absolute difference
            ans[i] = Math.abs(leftSum - rightSum);

            // Include current element in left side sum
            leftSum += nums[i];
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var leftRightDifference = function(nums) {
    
    const n = nums.length;

    // Calculate total array sum
    let rightSum = 0;
    for (const num of nums) {
        rightSum += num;
    }

    // Sum of elements on the left side
    let leftSum = 0;

    // Result array
    const ans = new Array(n);

    for (let i = 0; i < n; i++) {

        // Remove current element from right side sum
        rightSum -= nums[i];

        // Store absolute difference
        ans[i] = Math.abs(leftSum - rightSum);

        // Add current element to left side sum
        leftSum += nums[i];
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        
        # Total sum of all elements
        right_sum = sum(nums)

        # Sum of elements on the left side
        left_sum = 0

        # Result array
        ans = [0] * len(nums)

        for i in range(len(nums)):

            # Remove current element so right_sum contains
            # only elements to the right
            right_sum -= nums[i]

            # Store absolute difference
            ans[i] = abs(left_sum - right_sum)

            # Add current element to left_sum
            left_sum += nums[i]

        return ans
```

### Go

```go
func leftRightDifference(nums []int) []int {
    
    n := len(nums)

    // Calculate total array sum
    rightSum := 0
    for _, num := range nums {
        rightSum += num
    }

    // Sum of elements on the left side
    leftSum := 0

    // Result array
    ans := make([]int, n)

    for i := 0; i < n; i++ {

        // Remove current element from right side sum
        rightSum -= nums[i]

        // Store absolute difference
        diff := leftSum - rightSum
        if diff < 0 {
            diff = -diff
        }
        ans[i] = diff

        // Add current element to left side sum
        leftSum += nums[i]
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

First, calculate the total sum of the array.

This total sum initially represents everything on the right side.

Next, create a variable called `leftSum` and set it to `0`.

At the beginning, there are no elements to the left of index `0`.

Now start traversing the array.

For every element:

### Step 1: Update Right Sum

Remove the current element from the total remaining sum.

This is important because the problem asks for elements strictly to the right of the current index.

After subtraction, the remaining value becomes the true right-side sum.

### Step 2: Calculate Difference

Compute:

```text
abs(leftSum - rightSum)
```

Store this value in the answer array.

### Step 3: Update Left Sum

Add the current element to `leftSum`.

This prepares the left-side sum for the next index.

### Step 4: Move Forward

Continue the same process until the last element is processed.

By maintaining two running sums, we avoid repeated calculations and achieve optimal performance.

### Why This Works

At every index:

* `leftSum` contains the sum of all previous elements.
* `rightSum` contains the sum of all upcoming elements.

Because both values are always accurate, their absolute difference gives the correct answer.

### Edge Cases

#### Single Element Array

Input:

```text
[1]
```

There are no elements on either side.

```text
leftSum = 0
rightSum = 0
```

Answer:

```text
[0]
```

#### Two Elements

Input:

```text
[5, 3]
```

Output:

```text
[3, 5]
```

The logic still works without any special handling.

## Examples

### Example 1

Input:

```text
nums = [10,4,8,3]
```

Output:

```text
[15,1,11,22]
```

Trace:

```text
Index 0:
Left = 0
Right = 15
Difference = 15

Index 1:
Left = 10
Right = 11
Difference = 1

Index 2:
Left = 14
Right = 3
Difference = 11

Index 3:
Left = 22
Right = 0
Difference = 22
```

### Example 2

Input:

```text
nums = [1]
```

Output:

```text
[0]
```

Trace:

```text
Left = 0
Right = 0
Difference = 0
```

### Example 3

Input:

```text
nums = [5,2,1]
```

Output:

```text
[3,4,7]
```

Trace:

```text
Index 0:
Left = 0
Right = 3
Difference = 3

Index 1:
Left = 5
Right = 1
Difference = 4

Index 2:
Left = 7
Right = 0
Difference = 7
```

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
```

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

### JavaScript

Run:

```bash
node solution.js
```

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

## Notes & Optimizations

* This is a Prefix Sum style problem.
* A brute-force approach would compute left and right sums separately for every index and would take O(n²) time.
* Using running sums reduces the complexity to O(n).
* No extra prefix array is required.
* The solution uses constant extra memory apart from the required output array.
* This approach is commonly used in array processing and Prefix Sum interview questions.
* The algorithm works efficiently even for the maximum constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
