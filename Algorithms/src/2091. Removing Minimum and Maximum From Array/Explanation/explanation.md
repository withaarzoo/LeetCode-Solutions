# 2091. Removing Minimum and Maximum From Array

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

This problem asks us to remove both the minimum and maximum elements from an array using the minimum possible number of deletions.

The array contains distinct integers, so there is exactly one minimum element and one maximum element.

A deletion can only happen in two ways:

* Remove an element from the front of the array.
* Remove an element from the back of the array.

The goal is to find the minimum number of front and back deletions needed to remove both the minimum and maximum values.

The main input is the integer array `nums`, and the output is a single integer representing the minimum number of deletions.

This LeetCode problem can be solved efficiently by finding the positions of the minimum and maximum elements and checking all meaningful deletion strategies.

## Constraints

| Constraint     | Value                               |
| -------------- | ----------------------------------- |
| Array length   | `1 <= nums.length <= 10^5`          |
| Element value  | `-10^5 <= nums[i] <= 10^5`          |
| Array elements | All integers in `nums` are distinct |

Because the array can contain up to `10^5` elements, checking every possible deletion sequence would be unnecessary and inefficient. A linear scan gives the optimal solution.

## Intuition

My first observation was that I do not need to care about the exact values of most elements in the array.

I only need to know where the minimum and maximum elements are located.

Once I know their positions, there are only three useful ways to remove both of them:

1. Remove both by deleting elements from the front.
2. Remove both by deleting elements from the back.
3. Remove one from the front and the other from the back.

So instead of simulating different deletion orders, I can calculate the cost of these possibilities and return the smallest answer.

That is the key observation behind this minimum deletions array solution.

## Approach

I use the following steps:

1. Scan the array and find the index of the minimum element.
2. Find the index of the maximum element during the same scan.
3. Calculate how many deletions are needed if both elements are removed from the front.
4. Calculate how many deletions are needed if both elements are removed from the back.
5. Calculate the cost of removing one element from the front and the other from the back.
6. Return the minimum among all possible strategies.

For removing both elements from the front, I need to delete elements until I reach the farther index.

For removing both elements from the back, I need to delete elements until I reach the smaller index.

For the mixed approach, I check both possibilities because either the minimum or maximum element can be removed from the front first.

This gives an efficient `O(n)` solution for LeetCode 2091: Removing Minimum and Maximum From Array.

## Data Structures Used

No extra data structure is needed.

I only use a few variables:

* `minIndex` to store the position of the minimum element.
* `maxIndex` to store the position of the maximum element.
* Variables to store the deletion count for each possible strategy.

The original array is scanned directly, so the solution uses constant extra space.

## Operations & Behavior Summary

The algorithm works in the following order:

1. Start with the first element as the current minimum and maximum.
2. Scan the rest of the array.
3. Update the minimum index whenever a smaller value is found.
4. Update the maximum index whenever a larger value is found.
5. Calculate the cost of deleting from the front only.
6. Calculate the cost of deleting from the back only.
7. Calculate both mixed front-and-back deletion options.
8. Take the smallest result.

In simple terms, the algorithm finds where the two important elements are and checks every meaningful direction for removing them.

## Complexity

| Type             | Complexity | Explanation                                                                                                     |
| ---------------- | ---------- | --------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)`     | The array is scanned once to find the minimum and maximum element positions. Here, `n` is the length of `nums`. |
| Space Complexity | `O(1)`     | Only a fixed number of variables are used. No extra array, stack, queue, or hash map is needed.                 |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDeletions(vector<int>& nums) {
        // Store the total number of elements in the array.
        int n = nums.size();

        // Start by assuming the first element is both minimum and maximum.
        int minIndex = 0;
        int maxIndex = 0;

        // Find the positions of the minimum and maximum elements.
        for (int i = 1; i < n; i++) {
            // Update the minimum index if a smaller value is found.
            if (nums[i] < nums[minIndex]) {
                minIndex = i;
            }

            // Update the maximum index if a larger value is found.
            if (nums[i] > nums[maxIndex]) {
                maxIndex = i;
            }
        }

        // Remove everything from the front up to the farther special element.
        int removeFromFront = max(minIndex, maxIndex) + 1;

        // Remove everything from the back up to the farther special element.
        int removeFromBack = n - min(minIndex, maxIndex);

        // Remove one special element from the front and the other from the back.
        int removeFromBothSides = min(
            minIndex + 1 + (n - maxIndex),
            maxIndex + 1 + (n - minIndex)
        );

        // Return the minimum deletions among all possible strategies.
        return min(removeFromFront, min(removeFromBack, removeFromBothSides));
    }
};
```

### Java

```java
class Solution {
    public int minimumDeletions(int[] nums) {
        // Store the total number of elements in the array.
        int n = nums.length;

        // Start by assuming the first element is both minimum and maximum.
        int minIndex = 0;
        int maxIndex = 0;

        // Find the positions of the minimum and maximum elements.
        for (int i = 1; i < n; i++) {
            // Update the minimum index if a smaller value is found.
            if (nums[i] < nums[minIndex]) {
                minIndex = i;
            }

            // Update the maximum index if a larger value is found.
            if (nums[i] > nums[maxIndex]) {
                maxIndex = i;
            }
        }

        // Remove everything from the front up to the farther special element.
        int removeFromFront = Math.max(minIndex, maxIndex) + 1;

        // Remove everything from the back up to the farther special element.
        int removeFromBack = n - Math.min(minIndex, maxIndex);

        // Calculate both ways of removing one element from each side.
        int removeFromBothSides = Math.min(
            minIndex + 1 + (n - maxIndex),
            maxIndex + 1 + (n - minIndex)
        );

        // Return the minimum deletions among all possible strategies.
        return Math.min(
            removeFromFront,
            Math.min(removeFromBack, removeFromBothSides)
        );
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDeletions = function(nums) {
    // Store the total number of elements in the array.
    const n = nums.length;

    // Start by assuming the first element is both minimum and maximum.
    let minIndex = 0;
    let maxIndex = 0;

    // Find the positions of the minimum and maximum elements.
    for (let i = 1; i < n; i++) {
        // Update the minimum index if a smaller value is found.
        if (nums[i] < nums[minIndex]) {
            minIndex = i;
        }

        // Update the maximum index if a larger value is found.
        if (nums[i] > nums[maxIndex]) {
            maxIndex = i;
        }
    }

    // Remove everything from the front up to the farther special element.
    const removeFromFront = Math.max(minIndex, maxIndex) + 1;

    // Remove everything from the back up to the farther special element.
    const removeFromBack = n - Math.min(minIndex, maxIndex);

    // Calculate both ways of removing one element from each side.
    const removeFromBothSides = Math.min(
        minIndex + 1 + (n - maxIndex),
        maxIndex + 1 + (n - minIndex)
    );

    // Return the minimum deletions among all possible strategies.
    return Math.min(
        removeFromFront,
        removeFromBack,
        removeFromBothSides
    );
};
```

### Python3

```python
class Solution:
    def minimumDeletions(self, nums: List[int]) -> int:
        # Store the total number of elements in the array.
        n = len(nums)

        # Start by assuming the first element is both minimum and maximum.
        min_index = 0
        max_index = 0

        # Find the positions of the minimum and maximum elements.
        for i in range(1, n):
            # Update the minimum index if a smaller value is found.
            if nums[i] < nums[min_index]:
                min_index = i

            # Update the maximum index if a larger value is found.
            if nums[i] > nums[max_index]:
                max_index = i

        # Remove everything from the front up to the farther special element.
        remove_from_front = max(min_index, max_index) + 1

        # Remove everything from the back up to the farther special element.
        remove_from_back = n - min(min_index, max_index)

        # Calculate both ways of removing one element from each side.
        remove_from_both_sides = min(
            min_index + 1 + (n - max_index),
            max_index + 1 + (n - min_index)
        )

        # Return the minimum deletions among all possible strategies.
        return min(
            remove_from_front,
            remove_from_back,
            remove_from_both_sides
        )
```

### Go

```go
func minimumDeletions(nums []int) int {
 // Store the total number of elements in the array.
 n := len(nums)

 // Start by assuming the first element is both minimum and maximum.
 minIndex := 0
 maxIndex := 0

 // Find the positions of the minimum and maximum elements.
 for i := 1; i < n; i++ {
  // Update the minimum index if a smaller value is found.
  if nums[i] < nums[minIndex] {
   minIndex = i
  }

  // Update the maximum index if a larger value is found.
  if nums[i] > nums[maxIndex] {
   maxIndex = i
  }
 }

 // Remove everything from the front up to the farther special element.
 removeFromFront := max(minIndex, maxIndex) + 1

 // Remove everything from the back up to the farther special element.
 removeFromBack := n - min(minIndex, maxIndex)

 // Calculate both ways of removing one element from each side.
 removeFromBothSides := min(
  minIndex+1+(n-maxIndex),
  maxIndex+1+(n-minIndex),
 )

 // Return the minimum deletions among all possible strategies.
 return min(
  removeFromFront,
  min(removeFromBack, removeFromBothSides),
 )
}

// min returns the smaller of two integers.
func min(a, b int) int {
 if a < b {
  return a
 }
 return b
}

// max returns the larger of two integers.
func max(a, b int) int {
 if a > b {
  return a
 }
 return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in C++, Java, JavaScript, Python3, and Go. Only the syntax changes.

### Step 1: Find the minimum and maximum positions

I first scan the array and keep track of the indices of the smallest and largest values.

I do not need to sort the array because sorting would change the original positions and also take more time than necessary.

During one traversal:

* If the current value is smaller than the known minimum, I update the minimum index.
* If the current value is larger than the known maximum, I update the maximum index.

After this scan, I know exactly where both required elements are.

### Step 2: Remove both elements from the front

If both the minimum and maximum are removed from the front, I must reach the one that appears farther from the beginning.

For example, if one element is at index `2` and the other is at index `5`, deleting from the front requires removing the first `6` elements.

This is why the calculation uses the larger index plus one.

The extra one is necessary because array indexing starts from zero.

### Step 3: Remove both elements from the back

For this case, I delete elements only from the end of the array.

The element with the smaller index is farther from the back, so I need to delete enough elements to reach that position.

The number of deletions is based on the array length and the smaller of the two indices.

### Step 4: Remove one element from each side

There are two possible mixed strategies.

The first strategy removes:

* The minimum element from the front.
* The maximum element from the back.

The second strategy removes:

* The maximum element from the front.
* The minimum element from the back.

Both need to be checked because one can require fewer deletions than the other.

### Step 5: Return the smallest answer

At this point, I have checked all three meaningful cases:

* Front only.
* Back only.
* Both sides.

I return the smallest deletion count.

This works because any optimal solution must remove the minimum and maximum elements through one of these strategies. There is no need to simulate every possible sequence of front and back deletions.

### Language-specific behavior

The algorithm behaves the same in all five languages.

In C++, Java, JavaScript, Python3, and Go, the main differences are only:

* How arrays are represented.
* How minimum and maximum values are calculated.
* How functions and variables are written.

The underlying greedy approach and time complexity remain exactly the same.

## Examples

### Example 1

**Input:**

```text
nums = [2,10,7,5,4,1,8,6]
```

**Expected Output:**

```text
5
```

**Trace:**

* Minimum value `1` is at index `5`.
* Maximum value `10` is at index `1`.

Possible strategies:

* Remove both from the front: `6` deletions.
* Remove both from the back: `7` deletions.
* Remove maximum from the front and minimum from the back: `5` deletions.

The minimum answer is `5`.

### Example 2

**Input:**

```text
nums = [0,-4,19,1,8,-2,-3,5]
```

**Expected Output:**

```text
3
```

**Trace:**

* Minimum value `-4` is at index `1`.
* Maximum value `19` is at index `2`.

Both elements are close to the front.

Removing the first three elements removes both the minimum and maximum values.

So the answer is `3`.

### Example 3

**Input:**

```text
nums = [101]
```

**Expected Output:**

```text
1
```

**Trace:**

There is only one element in the array.

That element is both the minimum and maximum value.

Removing it with one deletion solves the problem.

So the answer is `1`.

## How to Use / Run Locally

Before running any solution, create a source file for the language you want to use and paste the corresponding solution code into it.

### C++

Save the code in a file named `solution.cpp`.

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it on macOS or Linux:

```bash
./solution
```

Run it on Windows:

```bash
solution.exe
```

### Java

Save the code in a file named `Solution.java`.

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

Make sure the Java Development Kit is installed and available from your terminal.

### JavaScript

Save the code in a file named `solution.js`.

Run it with Node.js:

```bash
node solution.js
```

You need Node.js installed on your system.

### Python3

Save the code in a file named `solution.py`.

Run it with:

```bash
python3 solution.py
```

On some Windows systems, the command may be:

```bash
python solution.py
```

Make sure Python 3 is installed.

### Go

Save the code in a file named `solution.go`.

Run it directly:

```bash
go run solution.go
```

Or build an executable:

```bash
go build solution.go
```

Then run the generated executable.

## Notes & Optimizations

The most important optimization is avoiding unnecessary simulation.

A direct simulation of every possible sequence of front and back deletions would make the solution more complicated without providing any benefit.

Only the positions of the minimum and maximum elements matter.

Sorting is also unnecessary. Although sorting could identify the minimum and maximum values, it would take `O(n log n)` time and would not directly preserve the original positions needed for the deletion calculation.

The edge case where the array contains only one element is handled naturally. The single element is both the minimum and maximum, and the answer is one deletion.

Because all array elements are distinct, there is no need to handle multiple positions for the same minimum or maximum value.

The final solution is optimal with `O(n)` time and `O(1)` extra space.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
