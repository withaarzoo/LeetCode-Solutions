# 3731. Find Missing Elements - LeetCode Solution

A beginner-friendly solution for **LeetCode 3731: Find Missing Elements** with a simple explanation, optimized approach, complexity analysis, and implementations in **C++, Java, JavaScript, Python, and Go**.

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

In this problem, you are given an array of **unique integers**. The array originally contained every integer within a continuous range, but some numbers have been removed.

The smallest and largest numbers from the original range are still present in the array. Your task is to find every missing number between them and return those numbers in sorted order.

If no numbers are missing, simply return an empty list.

This problem is a good practice for working with **arrays**, **hash sets**, and **range traversal** while keeping the solution efficient.

## Constraints

| Constraint | Value |
| ------------ | ------- |
| `2 <= nums.length <= 100` | Array size |
| `1 <= nums[i] <= 100` | Element value |
| All elements are unique | Yes |

## Intuition

The first thing I noticed was that the smallest and largest numbers are guaranteed to exist in the array. That means I already know the complete range that the original array should have covered.

Instead of trying to rebuild the original array, I only need a fast way to check whether each number in that range exists.

A hash set is perfect for this because it lets me check whether a number is present in constant time.

## Approach

I solved the problem in four simple steps.

1. Find the smallest number in the array.
2. Find the largest number in the array.
3. Store every element inside a hash set.
4. Traverse every number from the smallest value to the largest value.
5. If a number does not exist in the hash set, add it to the answer.
6. Return the final list.

This approach is easy to understand and runs efficiently.

## Data Structures Used

| Data Structure | Purpose |
| --------------- | --------- |
| Array | Stores the given numbers |
| Hash Set | Provides fast lookup to check whether a number already exists |
| List / Vector | Stores all missing numbers before returning the answer |

## Operations & Behavior Summary

The algorithm performs the following operations.

- Read every number from the input.
- Determine the minimum and maximum values.
- Store every number inside a hash set.
- Visit every integer from the minimum value to the maximum value.
- Check whether the current number exists.
- If it does not exist, store it in the answer.
- Return the completed list.

Since the numbers are checked in increasing order, the final answer is already sorted.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n + (max - min + 1))** | Finding the range and building the hash set takes `O(n)`. Traversing the complete range takes `O(max - min + 1)`. |
| Space Complexity | **O(n)** | Extra space is used for the hash set containing every element from the array. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> findMissingElements(vector<int>& nums) {
        // Store every number for O(1) lookup
        unordered_set<int> seen(nums.begin(), nums.end());

        // Find the smallest and largest numbers
        int mn = *min_element(nums.begin(), nums.end());
        int mx = *max_element(nums.begin(), nums.end());

        // Store all missing numbers
        vector<int> ans;

        // Check every value in the original range
        for (int x = mn; x <= mx; x++) {
            // If the value is not present, it is missing
            if (!seen.count(x)) {
                ans.push_back(x);
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> findMissingElements(int[] nums) {
        // Store every number for fast lookup
        HashSet<Integer> seen = new HashSet<>();

        int mn = Integer.MAX_VALUE;
        int mx = Integer.MIN_VALUE;

        // Fill the set and find the minimum and maximum
        for (int num : nums) {
            seen.add(num);
            mn = Math.min(mn, num);
            mx = Math.max(mx, num);
        }

        // Store the answer
        List<Integer> ans = new ArrayList<>();

        // Check every number in the range
        for (int x = mn; x <= mx; x++) {
            // Add numbers that do not exist
            if (!seen.contains(x)) {
                ans.add(x);
            }
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
var findMissingElements = function(nums) {
    // Store all numbers for constant-time lookup
    const seen = new Set(nums);

    // Find the minimum and maximum values
    let mn = Math.min(...nums);
    let mx = Math.max(...nums);

    // Store missing numbers
    const ans = [];

    // Check every value in the range
    for (let x = mn; x <= mx; x++) {
        // If the number is missing, save it
        if (!seen.has(x)) {
            ans.push(x);
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def findMissingElements(self, nums: List[int]) -> List[int]:
        # Store every number for fast lookup
        seen = set(nums)

        # Find the smallest and largest values
        mn = min(nums)
        mx = max(nums)

        # Store the missing numbers
        ans = []

        # Check every number in the range
        for x in range(mn, mx + 1):
            # If the number is missing, add it
            if x not in seen:
                ans.append(x)

        return ans
```

### Go

```go
func findMissingElements(nums []int) []int {
 // Store every number for constant-time lookup
 seen := make(map[int]bool)

 // Initialize minimum and maximum
 mn, mx := nums[0], nums[0]

 // Fill the map and find the range
 for _, num := range nums {
  seen[num] = true

  if num < mn {
   mn = num
  }

  if num > mx {
   mx = num
  }
 }

 // Store missing numbers
 ans := []int{}

 // Check every value in the range
 for x := mn; x <= mx; x++ {
  // If the value is missing, add it
  if !seen[x] {
   ans = append(ans, x)
  }
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in every language. Only the syntax changes.

### C++

The C++ solution uses an `unordered_set` to store every number. This allows constant-time lookups while checking every value inside the original range. A `vector` stores the missing numbers before returning them.

### Java

Java uses a `HashSet<Integer>` for quick lookups and an `ArrayList<Integer>` for the final answer. The algorithm remains exactly the same as the C++ implementation.

### JavaScript

The JavaScript version uses the built-in `Set` object. Every number is inserted into the set, and the solution checks each value in the complete range before adding missing numbers to the result array.

### Python3

Python makes this solution very clean because the built-in `set` data structure provides fast membership checks. A normal list stores the missing numbers in sorted order.

### Go

Go uses a map with boolean values to simulate a hash set. Missing values are collected inside a slice and returned at the end.

No matter which language you choose, the algorithm follows the same sequence of steps.

1. Find the smallest number.
2. Find the largest number.
3. Store every value for fast lookup.
4. Visit every number inside the range.
5. Save numbers that do not exist.
6. Return the final answer.

## Examples

### Example 1

**Input**

```text
nums = [1,4,2,5]
```

**Output**

```text
[3]
```

**Explanation**

The complete range should be `[1,2,3,4,5]`.

The numbers `1`, `2`, `4`, and `5` already exist.

Only `3` is missing.

---

### Example 2

**Input**

```text
nums = [7,8,6,9]
```

**Output**

```text
[]
```

**Explanation**

The complete range is `[6,7,8,9]`.

Every number already exists, so the answer is empty.

---

### Example 3

**Input**

```text
nums = [5,1]
```

**Output**

```text
[2,3,4]
```

**Explanation**

The complete range is `[1,2,3,4,5]`.

The numbers `2`, `3`, and `4` are missing.

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project folder.

```bash
cd <repository-name>
```

### C++

Compile the program.

```bash
g++ solution.cpp -o solution
```

Run it.

```bash
./solution
```

### Java

Compile the file.

```bash
javac Solution.java
```

Run it.

```bash
java Solution
```

### JavaScript

Run using Node.js.

```bash
node solution.js
```

### Python3

Run the program.

```bash
python solution.py
```

or

```bash
python3 solution.py
```

### Go

Run directly.

```bash
go run solution.go
```

Or build an executable.

```bash
go build solution.go
```

## Notes & Optimizations

- The input contains only unique integers, so duplicates never need to be handled.
- A hash set is the simplest way to achieve fast lookups.
- The returned list is automatically sorted because the range is traversed from the smallest value to the largest value.
- An alternative solution would be to sort the array first and compare adjacent elements. However, sorting increases the time complexity to `O(n log n)`, making the hash set approach more efficient for this problem.
- This solution is simple, easy to understand, and performs well within the given constraints.

## Author

**Md Aarzoo Islam**

Instagram: <https://www.instagram.com/code.with.aarzoo/>
