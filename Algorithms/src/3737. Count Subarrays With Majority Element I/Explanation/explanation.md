# 3737. Count Subarrays With Majority Element I

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

In this LeetCode problem, I am given an integer array `nums` and an integer `target`.

My task is to count how many subarrays have `target` as the majority element. A majority element means the value appears strictly more than half of the length of that subarray.

So for every subarray, I need to check whether:

`count(target) > length / 2`

The output is the total number of such valid subarrays.

This is a classic subarray counting problem, and the key idea is to track the frequency of `target` while checking all possible subarrays.

## Constraints

* `1 <= nums.length <= 1000`
* `1 <= nums[i] <= 10^9`
* `1 <= target <= 10^9`

## Intuition

My first thought was that the majority condition is very simple once I know two things: the length of the subarray and how many times `target` appears in it.

I do not need to track every number inside the subarray. I only care about one value: `target`.

That made the problem feel much easier. Since the array length is only up to 1000, I realized a direct `O(n^2)` solution is safe. I can try every subarray, keep a running count of `target`, and test the majority condition right away.

The important observation is this:

* if `2 * count(target) > subarray_length`, then `target` is the majority element

That removes the need for extra work and keeps the logic clean.

## Approach

I use two nested loops to generate all subarrays.

1. I fix the left boundary of the subarray.
2. I extend the right boundary one step at a time.
3. While extending, I keep a running count of how many times `target` appears.
4. For every subarray, I compute its length.
5. I check whether `2 * target_count > length`.
6. If yes, I increase the answer.

This works because every subarray is checked exactly once, and I update the target count in constant time instead of recounting from scratch.

This is a simple brute force approach, but it is efficient enough for the given constraints.

## Data Structures Used

* **Array / List**: to store the input numbers.
* **Integer variables**: to track:

  * the current answer
  * the target count in the current subarray
  * the current subarray length

No extra hash map, set, or prefix structure is needed for this solution.

## Operations & Behavior Summary

* Start from each index as the left boundary.
* Expand to the right one element at a time.
* Update the frequency of `target` only when the newly added element equals `target`.
* Compute the current subarray length from the two pointers.
* Check the majority condition using `2 * count(target) > length`.
* Add 1 to the result whenever the condition is true.

This is the full logic of the solution in plain English.

## Complexity

| Metric           | Complexity | Explanation                                                                                |
| ---------------- | ---------: | ------------------------------------------------------------------------------------------ |
| Time Complexity  |   `O(n^2)` | I try every subarray by fixing the left boundary and expanding the right boundary.         |
| Space Complexity |     `O(1)` | I only use a few integer variables and no extra data structure that grows with input size. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countMajoritySubarrays(vector<int>& nums, int target) {
        int n = nums.size();
        int ans = 0;

        // Try every possible starting index
        for (int left = 0; left < n; left++) {
            int countTarget = 0;

            // Extend the subarray one element at a time
            for (int right = left; right < n; right++) {

                // Update target frequency for the current subarray
                if (nums[right] == target)
                    countTarget++;

                // Current subarray length
                int len = right - left + 1;

                // Target is a majority if it appears more than half the time
                if (2 * countTarget > len)
                    ans++;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int countMajoritySubarrays(int[] nums, int target) {
        int n = nums.length;
        int ans = 0;

        // Try every possible starting index
        for (int left = 0; left < n; left++) {
            int countTarget = 0;

            // Extend the subarray one element at a time
            for (int right = left; right < n; right++) {

                // Update target frequency
                if (nums[right] == target)
                    countTarget++;

                // Current subarray length
                int len = right - left + 1;

                // Check majority condition
                if (2 * countTarget > len)
                    ans++;
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
 * @param {number} target
 * @return {number}
 */
var countMajoritySubarrays = function(nums, target) {
    const n = nums.length;
    let ans = 0;

    // Try every possible starting index
    for (let left = 0; left < n; left++) {
        let countTarget = 0;

        // Extend the subarray
        for (let right = left; right < n; right++) {

            // Update target frequency
            if (nums[right] === target)
                countTarget++;

            // Current subarray length
            const len = right - left + 1;

            // Check majority condition
            if (2 * countTarget > len)
                ans++;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        n = len(nums)
        ans = 0

        # Try every possible starting index
        for left in range(n):
            count_target = 0

            # Extend the subarray
            for right in range(left, n):

                # Update target frequency
                if nums[right] == target:
                    count_target += 1

                # Current subarray length
                length = right - left + 1

                # Check majority condition
                if 2 * count_target > length:
                    ans += 1

        return ans
```

### Go

```go
func countMajoritySubarrays(nums []int, target int) int {
 n := len(nums)
 ans := 0

 // Try every possible starting index
 for left := 0; left < n; left++ {
  countTarget := 0

  // Extend the subarray
  for right := left; right < n; right++ {

   // Update target frequency
   if nums[right] == target {
    countTarget++
   }

   // Current subarray length
   length := right - left + 1

   // Check majority condition
   if 2*countTarget > length {
    ans++
   }
  }
 }

 return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages, so I explain the idea once and it applies everywhere.

First, I store the length of the array and create an answer variable. The answer starts at zero because I have not found any valid subarray yet.

Then I pick a left boundary. This represents where the subarray starts. For each left boundary, I reset the count of `target` because I am beginning a new set of subarrays.

After that, I move the right boundary from the left boundary to the end of the array. Every time I move right by one step, I include one more number in the current subarray.

If that new number equals `target`, I increase the count. I do not scan the whole subarray again because that would waste time.

Next, I compute the subarray length using the two pointers. This is just `right - left + 1`.

Now I check the majority rule. If `2 * count(target)` is greater than the length, then `target` appears more than half the time, so this subarray is valid.

Whenever that happens, I add one to the answer.

At the end, after checking all possible start and end positions, the answer contains the total number of valid subarrays.

The same reasoning fits C++, Java, JavaScript, Python3, and Go because each language can handle the same two-pointer brute force structure easily.

## Examples

### Example 1

**Input:** `nums = [1,2,2,3], target = 2`

**Output:** `5`

**Trace:**

* `[2]` → target appears 1 out of 1, valid
* `[2]` → target appears 1 out of 1, valid
* `[2,2]` → target appears 2 out of 2, valid
* `[1,2,2]` → target appears 2 out of 3, valid
* `[2,2,3]` → target appears 2 out of 3, valid

So the total count is 5.

### Example 2

**Input:** `nums = [1,1,1,1], target = 1`

**Output:** `10`

**Trace:**
Every subarray is made only of `1`, so `1` is always the majority element.

The number of subarrays in an array of length 4 is:

`4 * 5 / 2 = 10`

So the answer is 10.

### Example 3

**Input:** `nums = [1,2,3], target = 4`

**Output:** `0`

**Trace:**
The value `4` never appears in the array, so it cannot be the majority element in any subarray.

## How to Use / Run Locally

### C++

Save the file as `main.cpp`, add the solution class, and compile it with:

```bash
g++ -std=c++17 -O2 -o main main.cpp
./main
```

### Java

Save the file as `Solution.java`, then compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript

Save the file as `solution.js`, then run:

```bash
node solution.js
```

### Python3

Save the file as `solution.py`, then run:

```bash
python3 solution.py
```

### Go

Save the file as `main.go`, then run:

```bash
go run main.go
```

## Notes & Optimizations

This problem looks like it might need a clever data structure at first, but the constraints are small enough that a clean brute force approach is enough.

A few useful points:

* I only track `target`, not every number in the subarray.
* I use `2 * count > length` to avoid floating-point division.
* The solution is easy to understand and hard to get wrong.
* If the input size were much larger, I would need a more advanced idea, but here `O(n^2)` is the right tradeoff.

This makes the solution simple, readable, and reliable for competitive programming practice.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
