# 2948. Make Lexicographically Smallest Array by Swapping Elements

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

This repository contains an optimized solution for LeetCode Problem 2948, **Make Lexicographically Smallest Array by Swapping Elements**.

We are given a `0-indexed` array of positive integers called `nums` and an integer called `limit`.

In one operation, we can choose any two indices `i` and `j` and swap `nums[i]` with `nums[j]` if:

`|nums[i] - nums[j]| <= limit`

We can perform this operation any number of times.

The goal is to return the **lexicographically smallest array** that can be created using valid swaps.

The important part of this DSA problem is understanding that some elements can be connected indirectly. Two values may not be able to swap directly, but they can still move between positions through a sequence of valid swaps.

This solution uses a **sorting and connected components approach** to efficiently find groups of values that can be rearranged with each other.

## Constraints

| Constraint    | Value                      |
| ------------- | -------------------------- |
| `nums.length` | `1 <= nums.length <= 10^5` |
| `nums[i]`     | `1 <= nums[i] <= 10^9`     |
| `limit`       | `1 <= limit <= 10^9`       |

Because the array size can be as large as `10^5`, checking every possible pair or simulating swaps directly would be too slow.

## Intuition

My first thought was to look at the condition for a valid swap.

A swap depends only on the difference between two values. The positions themselves do not decide whether a swap is allowed.

Then I noticed an important pattern.

Suppose the sorted values are:

`1, 3, 5`

and `limit = 2`.

`1` can swap with `3`, and `3` can swap with `5`.

Even though `1` and `5` cannot swap directly, they are still connected through `3`. That means all three values can effectively belong to the same swappable group.

So instead of thinking about individual swaps, I thought about groups of values that are connected through valid swaps.

After sorting the array, I only need to check the difference between consecutive values. If the difference is within `limit`, they belong to the same group. If the difference is greater than `limit`, a new group starts.

Once I know a group, I can make the answer lexicographically smallest by placing its smallest values at its smallest original indices.

## Approach

I use the following sorting-based approach:

1. Store every element together with its original index.
2. Sort all elements by their values.
3. Scan the sorted elements from left to right.
4. Keep consecutive values in the same group while their difference is less than or equal to `limit`.
5. Start a new group when the difference becomes greater than `limit`.
6. Collect all original indices belonging to the current group.
7. Sort those indices in increasing order.
8. The values in the current group are already sorted because the full array was sorted earlier.
9. Place the smallest group value at the smallest index, the next smallest value at the next smallest index, and continue.
10. Repeat this process for every group.

This greedy assignment works because lexicographical order always gives priority to earlier positions.

If a smaller value can be placed at an earlier index, that is always better than placing a larger value there.

## Data Structures Used

### Value and Index Pairs

I store each number together with its original index.

This helps me sort the values while still remembering where each value came from.

For example:

`[1, 7, 6]`

can be represented as:

`[(1, 0), (7, 1), (6, 2)]`

### Sorted Array of Pairs

Sorting by value makes connected values appear next to each other.

This allows me to find valid swap groups by checking consecutive differences instead of checking every possible pair.

### Index List

For each connected group, I collect all original indices and sort them.

This lets me place the smallest available values into the earliest possible positions.

### Result Array

I use a separate result array to build the lexicographically smallest answer.

Each group fills only the positions that belong to that group.

## Operations & Behavior Summary

The algorithm works like this:

```text
Store each number with its original index

Sort all pairs by value

Start from the first sorted element

While consecutive values differ by at most limit:
    Keep them in the same group

Collect all original indices of the group

Sort those indices

Since group values are already sorted:
    Put the smallest value at the smallest index
    Put the next smallest value at the next smallest index
    Continue until the group is finished

Move to the next group

Return the final array
```

The algorithm does not simulate individual swaps.

Instead, it finds which elements can eventually move between the same set of positions and directly creates the smallest possible arrangement.

## Complexity

| Complexity       | Value        | Explanation                                                                                                        |
| ---------------- | ------------ | ------------------------------------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n log n)` | Sorting the elements takes `O(n log n)`. Sorting the indices of the groups also stays within `O(n log n)` overall. |
| Space Complexity | `O(n)`       | Extra space is used for value-index pairs, the result array, and temporary index lists.                            |

Here, `n` is the length of the `nums` array.

This complexity is efficient enough for the maximum constraint of `10^5` elements.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> lexicographicallySmallestArray(vector<int>& nums, int limit) {
        int n = nums.size(); // Store the size so I can use it throughout the solution.
        
        // Store every value together with its original position.
        vector<pair<int, int>> elements;
        for (int i = 0; i < n; i++) {
            elements.push_back({nums[i], i});
        }
        
        // Sort by value so connected values appear next to each other.
        sort(elements.begin(), elements.end());
        
        // This will store the lexicographically smallest possible result.
        vector<int> answer(n);
        
        int start = 0; // Marks the first element of the current connected group.
        
        while (start < n) {
            int end = start; // Expand this group as far as valid swaps allow.
            
            // Consecutive sorted values belong to the same group if their
            // difference is at most limit, which means they are connected.
            while (end + 1 < n &&
                   (long long)elements[end + 1].first - elements[end].first <= limit) {
                end++;
            }
            
            vector<int> indices; // Store the original positions of this group.
            
            // Collect every original index that belongs to the current group.
            for (int i = start; i <= end; i++) {
                indices.push_back(elements[i].second);
            }
            
            // Sort positions so smaller values can be placed earlier.
            sort(indices.begin(), indices.end());
            
            // Values are already sorted in elements[start...end].
            // Assign them to the sorted original indices.
            for (int i = 0; i < (int)indices.size(); i++) {
                answer[indices[i]] = elements[start + i].first;
            }
            
            // Start processing the next connected group.
            start = end + 1;
        }
        
        return answer; // Return the lexicographically smallest arrangement.
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int[] lexicographicallySmallestArray(int[] nums, int limit) {
        int n = nums.length; // Store the array size for easier use.
        
        // Each pair stores {value, originalIndex}.
        int[][] elements = new int[n][2];
        
        for (int i = 0; i < n; i++) {
            elements[i][0] = nums[i]; // Store the value.
            elements[i][1] = i;       // Store its original position.
        }
        
        // Sort all pairs by their values.
        Arrays.sort(elements, (a, b) -> Integer.compare(a[0], b[0]));
        
        // Store the final lexicographically smallest array.
        int[] answer = new int[n];
        
        int start = 0; // First position of the current connected group.
        
        while (start < n) {
            int end = start; // Expand the current group.
            
            // Keep consecutive values in the same group while they can
            // be connected through valid swaps.
            while (end + 1 < n &&
                   (long) elements[end + 1][0] - elements[end][0] <= limit) {
                end++;
            }
            
            // Store all original indices belonging to this group.
            int[] indices = new int[end - start + 1];
            
            for (int i = start; i <= end; i++) {
                indices[i - start] = elements[i][1];
            }
            
            // Sort positions so smaller values go to earlier indices.
            Arrays.sort(indices);
            
            // The values in the current group are already sorted because
            // the complete elements array was sorted by value.
            for (int i = 0; i < indices.length; i++) {
                answer[indices[i]] = elements[start + i][0];
            }
            
            // Move to the next connected group.
            start = end + 1;
        }
        
        return answer; // Return the final array.
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} limit
 * @return {number[]}
 */
var lexicographicallySmallestArray = function(nums, limit) {
    const n = nums.length; // Store the array size.
    
    // Store every value with its original index.
    const elements = nums.map((value, index) => [value, index]);
    
    // Sort by value so connected values become consecutive.
    elements.sort((a, b) => a[0] - b[0]);
    
    // Store the lexicographically smallest result.
    const answer = new Array(n);
    
    let start = 0; // First element of the current connected group.
    
    while (start < n) {
        let end = start; // Expand the current group.
        
        // Consecutive values stay in the same group when their difference
        // is at most limit, so they are connected through valid swaps.
        while (
            end + 1 < n &&
            elements[end + 1][0] - elements[end][0] <= limit
        ) {
            end++;
        }
        
        // Collect all original indices of the current group.
        const indices = [];
        
        for (let i = start; i <= end; i++) {
            indices.push(elements[i][1]);
        }
        
        // Sort positions so smaller values are placed earlier.
        indices.sort((a, b) => a - b);
        
        // Values are already sorted inside the current group.
        for (let i = 0; i < indices.length; i++) {
            answer[indices[i]] = elements[start + i][0];
        }
        
        // Move to the next connected group.
        start = end + 1;
    }
    
    return answer; // Return the lexicographically smallest array.
};
```

### Python3

```python
from typing import List

class Solution:
    def lexicographicallySmallestArray(self, nums: List[int], limit: int) -> List[int]:
        n = len(nums)  # Store the size of the array.
        
        # Store each value together with its original index.
        elements = [(nums[i], i) for i in range(n)]
        
        # Sort by value so connected values become consecutive.
        elements.sort()
        
        # Store the lexicographically smallest result.
        answer = [0] * n
        
        start = 0  # First element of the current connected group.
        
        while start < n:
            end = start  # Expand the current group.
            
            # Consecutive values belong to the same group when their
            # difference is at most limit.
            while (
                end + 1 < n
                and elements[end + 1][0] - elements[end][0] <= limit
            ):
                end += 1
            
            # Collect all original indices belonging to this group.
            indices = []
            
            for i in range(start, end + 1):
                indices.append(elements[i][1])
            
            # Sort positions so smaller values go to earlier positions.
            indices.sort()
            
            # The values are already sorted because elements was sorted.
            for i, index in enumerate(indices):
                answer[index] = elements[start + i][0]
            
            # Move to the next connected group.
            start = end + 1
        
        return answer  # Return the final lexicographically smallest array.
```

### Go

```go
import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
 n := len(nums) // Store the size of the input array.

 // Each pair stores a value and its original index.
 type Pair struct {
  value int
  index int
 }

 // Create an array containing all values with their original positions.
 elements := make([]Pair, n)
 for i := 0; i < n; i++ {
  elements[i] = Pair{nums[i], i}
 }

 // Sort by value so connected values appear next to each other.
 sort.Slice(elements, func(i, j int) bool {
  return elements[i].value < elements[j].value
 })

 // Store the lexicographically smallest result.
 answer := make([]int, n)

 start := 0 // Marks the beginning of the current connected group.

 for start < n {
  end := start // Expand the current connected group.

  // Keep consecutive values together while their difference is
  // at most limit, meaning they are connected through swaps.
  for end+1 < n &&
   int64(elements[end+1].value)-int64(elements[end].value) <= int64(limit) {
   end++
  }

  // Collect all original indices belonging to this group.
  indices := make([]int, 0, end-start+1)
  for i := start; i <= end; i++ {
   indices = append(indices, elements[i].index)
  }

  // Sort positions so smaller values can be placed first.
  sort.Ints(indices)

  // Values in elements[start:end+1] are already sorted.
  for i := 0; i < len(indices); i++ {
   answer[indices[i]] = elements[start+i].value
  }

  // Continue with the next connected group.
  start = end + 1
 }

 return answer // Return the lexicographically smallest arrangement.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The same algorithm is used in C++, Java, JavaScript, Python3, and Go. The syntax and built-in sorting functions are different, but the logic remains exactly the same.

### Step 1: Keep track of values and original positions

I cannot simply sort `nums` because I need to know the original positions of the elements.

So I store each element as:

`(value, originalIndex)`

For example:

```text
nums = [1, 7, 6, 18, 2, 1]
```

becomes:

```text
(1, 0)
(7, 1)
(6, 2)
(18, 3)
(2, 4)
(1, 5)
```

The value tells me how elements are connected. The original index tells me which positions can receive those values later.

Without storing indices, I would lose the information needed to reconstruct the final array.

### Step 2: Sort everything by value

After sorting, the previous example becomes:

```text
(1, 0)
(1, 5)
(2, 4)
(6, 2)
(7, 1)
(18, 3)
```

Now similar values are close together.

This makes it possible to find connected groups by checking neighboring values.

I do not need to compare every value with every other value, which would be too expensive for large inputs.

### Step 3: Find the connected groups

Suppose `limit = 3`.

I compare consecutive sorted values:

```text
1 and 1 -> difference = 0
1 and 2 -> difference = 1
2 and 6 -> difference = 4
```

The first three values belong to the same group because each consecutive difference is at most `3`.

But the difference between `2` and `6` is greater than `3`, so a new group must start.

The groups are:

```text
[1, 1, 2]
[6, 7]
[18]
```

This works because values can connect indirectly.

For example, if `1` can swap with `3` and `3` can swap with `5`, all three values belong to the same connected group even if `1` and `5` cannot directly swap.

### Step 4: Collect original indices for each group

Consider this group:

```text
(1, 0)
(1, 5)
(2, 4)
```

Its values are already sorted:

```text
[1, 1, 2]
```

Its original indices are:

```text
[0, 5, 4]
```

I sort those indices:

```text
[0, 4, 5]
```

Now I know the earliest positions where values from this group can appear.

### Step 5: Assign smaller values to smaller indices

The sorted values are:

```text
[1, 1, 2]
```

The sorted positions are:

```text
[0, 4, 5]
```

So I assign:

```text
index 0 -> 1
index 4 -> 1
index 5 -> 2
```

This is the most important greedy step.

Lexicographical comparison starts from index `0`.

If two arrays differ at the first position, the array with the smaller value at that position is lexicographically smaller.

So whenever multiple values can be rearranged among multiple positions, I always place the smallest value at the smallest available index.

If I placed a larger value earlier, the answer could not be lexicographically optimal.

### Step 6: Process every group independently

Once one group is finished, I move to the next group.

Values from different groups cannot cross the gap between them because the sorted consecutive difference is greater than `limit`.

That means each group can be processed independently.

### Language-specific behavior

**C++** uses a `vector` of pairs and `sort()` from the standard library.

**Java** can use a two-dimensional array or custom pair-like objects together with `Arrays.sort()`.

**JavaScript** uses arrays containing `[value, index]` pairs and custom comparator functions with `sort()`.

**Python3** can use tuples because Python naturally sorts tuples by their first value.

**Go** can use a struct containing the value and original index, then sort it using `sort.Slice()`.

The algorithm and complexity stay the same in every language.

## Examples

### Example 1

**Input:**

```text
nums = [1, 5, 3, 9, 8]
limit = 2
```

**Expected Output:**

```text
[1, 3, 5, 8, 9]
```

**How it works:**

After sorting by value:

```text
1, 3, 5, 8, 9
```

The differences are:

```text
3 - 1 = 2
5 - 3 = 2
8 - 5 = 3
9 - 8 = 1
```

So the groups are:

```text
[1, 3, 5]
[8, 9]
```

For the first group, the values can move among their original positions, so sorting both the values and positions gives the smallest possible arrangement.

The same process works for the second group.

The final result is:

```text
[1, 3, 5, 8, 9]
```

### Example 2

**Input:**

```text
nums = [1, 7, 6, 18, 2, 1]
limit = 3
```

**Expected Output:**

```text
[1, 6, 7, 18, 1, 2]
```

**Sorted values with original indices:**

```text
(1, 0)
(1, 5)
(2, 4)
(6, 2)
(7, 1)
(18, 3)
```

The groups are:

```text
[1, 1, 2]
[6, 7]
[18]
```

For the first group:

```text
Values:    [1, 1, 2]
Positions: [0, 4, 5]
```

For the second group:

```text
Values:    [6, 7]
Positions: [1, 2]
```

The value `18` stays alone.

After assigning every group, the result becomes:

```text
[1, 6, 7, 18, 1, 2]
```

### Example 3

**Input:**

```text
nums = [1, 7, 28, 19, 10]
limit = 3
```

**Expected Output:**

```text
[1, 7, 28, 19, 10]
```

**How it works:**

After sorting:

```text
1, 7, 10, 19, 28
```

The differences show that some values cannot connect with the others through valid swaps.

No rearrangement can produce a lexicographically smaller valid array.

So the original array remains the answer:

```text
[1, 7, 28, 19, 10]
```

## How to Use / Run Locally

### C++

Make sure a C++ compiler such as `g++` is installed.

Create a file such as:

```text
solution.cpp
```

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it on Linux or macOS:

```bash
./solution
```

On Windows:

```bash
solution.exe
```

### Java

Make sure the JDK is installed.

Create a file containing the Java solution.

If you are testing locally with your own `main` method, compile it with:

```bash
javac Main.java
```

Run it with:

```bash
java Main
```

For LeetCode, only the required `Solution` class and method need to be submitted.

### JavaScript

Make sure Node.js is installed.

Create a file such as:

```text
solution.js
```

Run it with:

```bash
node solution.js
```

If you are submitting directly to LeetCode, paste the required function into the JavaScript editor.

### Python3

Make sure Python 3 is installed.

Create a file such as:

```text
solution.py
```

Run it with:

```bash
python solution.py
```

Depending on your system, you may need:

```bash
python3 solution.py
```

For LeetCode, submit the required `Solution` class.

### Go

Make sure Go is installed.

Create a file such as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

For local testing, add a `main` function. For LeetCode submission, use the required function format provided by the platform.

## Notes & Optimizations

The main optimization is avoiding direct swap simulation.

The number of possible swaps can be very large, so trying every operation would not be practical.

Sorting changes the problem into a grouping problem. After sorting, a gap larger than `limit` clearly separates two groups that cannot connect through valid swaps.

One important edge case is duplicate values. Their difference is `0`, so they always belong to the same connected group when they appear together.

Another important case is when the entire array forms one connected group. In that situation, all values can effectively be rearranged among all their positions, and the answer becomes the fully sorted array.

If every element belongs to a separate group, no useful swaps are possible and the original array remains unchanged.

A DSU or Union-Find solution is also possible. However, sorting is simpler here because connected components can be found by checking consecutive sorted values. This avoids building unnecessary edges between pairs of elements.

The final sorting-based solution runs in `O(n log n)` time and uses `O(n)` extra space, making it suitable for the full constraints of LeetCode 2948.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
