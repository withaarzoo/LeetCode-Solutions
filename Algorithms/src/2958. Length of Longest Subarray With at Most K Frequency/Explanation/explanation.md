# 2958. Length of Longest Subarray With at Most K Frequency

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

LeetCode 2958, **Length of Longest Subarray With at Most K Frequency**, asks us to find the longest contiguous subarray where no element appears more than `k` times.

We are given an integer array `nums` and an integer `k`.

A subarray is considered good when every value inside that subarray has a frequency less than or equal to `k`.

For example, if `k = 2`, a subarray can contain:

```text
[1, 2, 1, 3, 2]
```

because `1` appears twice, `2` appears twice, and `3` appears once.

But this would not be valid:

```text
[1, 2, 1, 3, 1]
```

because `1` appears three times.

The goal is to return the length of the longest good subarray.

This problem is a good example of the **sliding window**, **two pointers**, and **hash table frequency counting** techniques commonly used in competitive programming and DSA interviews.

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^9`
* `1 <= k <= nums.length`

The array can contain up to `100,000` elements, so an `O(n²)` brute-force solution is not practical.

## Intuition

My first thought was to generate every possible subarray and count how many times each number appears.

The problem with that idea is the number of possible subarrays. For an array of size `n`, there can be around `n²` different subarrays. With `n` up to `10^5`, that approach would be much too slow.

I noticed that the condition depends only on the frequencies of values inside the current subarray.

That made a sliding window a natural fit.

I keep two pointers, `left` and `right`, and treat everything between them as the current window.

Whenever I move `right` forward, I add the new number to a frequency map.

If that number appears more than `k` times, the window is no longer valid. I then move `left` forward until the frequency becomes valid again.

This lets me reuse the work from the previous window instead of checking every subarray from scratch.

## Approach

I use a sliding window with two pointers.

The `right` pointer moves through the array one element at a time.

For every new element:

1. I add it to the frequency map.
2. I check whether its frequency has become greater than `k`.
3. If the window is invalid, I move `left` forward.
4. Every element removed from the left has its frequency decreased.
5. I continue until the current window becomes valid.
6. I calculate the current window length.
7. I update the maximum answer.

The key idea is that the previous window was already valid. When I add one new element, only the frequency of that newly added value can become invalid.

Because both pointers only move forward, the entire array can be processed in linear time.

## Data Structures Used

### Hash Map

I use a hash map to store the frequency of each number inside the current sliding window.

For example:

```text
nums = [1, 2, 1, 3]
```

The frequency map can look like:

```text
1 -> 2
2 -> 1
3 -> 1
```

A hash map is useful here because `nums[i]` can be as large as `10^9`. I do not need to create a huge array based on the actual values.

### Two Pointers

I use `left` and `right` to represent the current window.

The current subarray is always:

```text
nums[left ... right]
```

`right` expands the window, while `left` shrinks it whenever the frequency condition is violated.

## Operations & Behavior Summary

The algorithm can be viewed as this simple process:

```text
Start with an empty window.

Move right through the array.

Add nums[right] to the frequency map.

If nums[right] appears more than k times:
    Move left forward.
    Decrease the frequency of every removed value.
    Continue until the window is valid again.

Calculate the current window length.

Update the maximum answer.

Return the maximum length.
```

The important rule is that after the shrinking step, the current window must always satisfy the condition:

```text
frequency of every value <= k
```

## Complexity

| Operation        | Complexity | Explanation                                                        |
| ---------------- | ---------: | ------------------------------------------------------------------ |
| Time Complexity  |     `O(n)` | Both `left` and `right` move from left to right at most `n` times. |
| Space Complexity |     `O(n)` | The frequency map can contain up to `n` different values.          |

Here, `n` represents the number of elements in `nums`.

The solution is efficient enough for the maximum constraint of `10^5` elements.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxSubarrayLength(vector<int>& nums, int k) {
        unordered_map<int, int> freq; // Stores how many times each value appears in the current window.
        int left = 0;                  // Left boundary of the sliding window.
        int ans = 0;                   // Stores the maximum valid window length found so far.

        for (int right = 0; right < nums.size(); right++) {
            freq[nums[right]]++;       // Add the current element to the window and increase its frequency.

            while (freq[nums[right]] > k) {
                freq[nums[left]]--;    // Remove the leftmost element because the current window is invalid.
                left++;                // Move the left boundary forward to shrink the window.
            }

            ans = max(ans, right - left + 1); // The window is valid, so update the maximum length.
        }

        return ans;                    // Return the length of the longest valid subarray.
    }
};
```

### Java

```java
class Solution {
    public int maxSubarrayLength(int[] nums, int k) {
        HashMap<Integer, Integer> freq = new HashMap<>(); // Stores frequencies of values in the current window.
        int left = 0;                                     // Left boundary of the sliding window.
        int ans = 0;                                      // Stores the longest valid window length.

        for (int right = 0; right < nums.length; right++) {
            freq.put(nums[right], freq.getOrDefault(nums[right], 0) + 1); // Add nums[right] to the window.

            while (freq.get(nums[right]) > k) {
                freq.put(nums[left], freq.get(nums[left]) - 1); // Remove nums[left] from the current window.
                left++;                                         // Move the left boundary forward.
            }

            ans = Math.max(ans, right - left + 1); // Update the answer using the current valid window.
        }

        return ans; // Return the longest valid subarray length.
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSubarrayLength = function(nums, k) {
    const freq = new Map(); // Stores the frequency of each value inside the current window.
    let left = 0;           // Left boundary of the sliding window.
    let ans = 0;            // Stores the longest valid window length.

    for (let right = 0; right < nums.length; right++) {
        // Add nums[right] to the window and increase its frequency.
        freq.set(nums[right], (freq.get(nums[right]) || 0) + 1);

        // Shrink the window while the newly added value appears too many times.
        while (freq.get(nums[right]) > k) {
            // Decrease the frequency of the element leaving from the left.
            freq.set(nums[left], freq.get(nums[left]) - 1);

            // Move the left boundary forward.
            left++;
        }

        // The current window is valid, so update the maximum length.
        ans = Math.max(ans, right - left + 1);
    }

    // Return the length of the longest valid subarray.
    return ans;
};
```

### Python3

```python
class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        freq = {}  # Stores the frequency of each value inside the current window.
        left = 0   # Left boundary of the sliding window.
        ans = 0    # Stores the longest valid window length.

        for right in range(len(nums)):
            # Add nums[right] to the window and increase its frequency.
            freq[nums[right]] = freq.get(nums[right], 0) + 1

            # Shrink the window while the newly added value appears more than k times.
            while freq[nums[right]] > k:
                # Decrease the frequency of the element leaving from the left.
                freq[nums[left]] -= 1

                # Move the left boundary forward.
                left += 1

            # The current window is valid, so update the maximum length.
            ans = max(ans, right - left + 1)

        # Return the length of the longest valid subarray.
        return ans
```

### Go

```go
func maxSubarrayLength(nums []int, k int) int {
 freq := make(map[int]int) // Stores the frequency of each value inside the current window.
 left := 0                  // Left boundary of the sliding window.
 ans := 0                   // Stores the longest valid window length.

 for right := 0; right < len(nums); right++ {
  freq[nums[right]]++ // Add nums[right] to the window and increase its frequency.

  for freq[nums[right]] > k {
   freq[nums[left]]-- // Remove nums[left] because the current window is invalid.
   left++             // Move the left boundary forward.
  }

  if right-left+1 > ans {
   ans = right - left + 1 // Update the maximum length using the valid window.
  }
 }

 return ans // Return the length of the longest valid subarray.
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The five implementations use the same algorithm. The syntax changes between languages, but the actual logic remains the same.

### 1. Create the frequency map

I first create a hash map to track how many times each value appears in the current window.

For example, if the current window is:

```text
[1, 2, 1, 3]
```

the map stores:

```text
1 -> 2
2 -> 1
3 -> 1
```

This gives me the information I need to decide whether the window is valid.

C++ uses `unordered_map`.

Java uses `HashMap`.

JavaScript uses `Map`.

Python uses a dictionary.

Go uses a map.

All of them provide average `O(1)` lookup and update operations.

### 2. Initialize the left pointer

I start `left` at index `0`.

The right pointer will move through the array.

Initially, the window is empty.

As `right` moves forward, the window gradually grows.

### 3. Add the right-side element

At every iteration, I add `nums[right]` to the frequency map.

Suppose:

```text
k = 2
```

and the current frequency is:

```text
5 -> 2
```

If I add another `5`, it becomes:

```text
5 -> 3
```

The window is now invalid because `5` occurs more than twice.

### 4. Shrink the window

When the frequency becomes greater than `k`, I start moving `left`.

For every element removed from the window, I decrease its frequency in the map.

For example:

```text
[1, 2, 3, 1, 2, 3, 1]
 ^
 left
```

If the final `1` causes its frequency to become `3` while `k = 2`, I remove elements from the left until one occurrence of `1` leaves the window.

The window then becomes valid again.

### 5. Why the window can be safely moved forward

I never need to move `right` backward.

I also never need to move `left` backward.

Once an element is outside the window, I do not need to reconsider that earlier window.

This is the main reason the sliding window is efficient.

### 6. Calculate the current length

Once the window is valid, its length is:

```text
right - left + 1
```

The `+1` is necessary because both endpoints are included.

For example:

```text
left = 2
right = 6
```

means the window contains indices:

```text
2, 3, 4, 5, 6
```

So its length is:

```text
6 - 2 + 1 = 5
```

I compare this length with the best answer found so far.

### 7. Why the solution is `O(n)`

At first glance, the nested `while` loop might look like it makes the algorithm `O(n²)`.

It does not.

The important detail is that `left` never moves backward.

Across the complete algorithm, `left` can move at most `n` times.

The `right` pointer also moves exactly `n` times.

So the total number of pointer movements is at most around `2n`, which is still `O(n)`.

### C++ Behavior

The C++ implementation uses `unordered_map<int, int>` for frequency counting.

The map provides average constant-time insertion, lookup, and update operations.

The two-pointer logic is exactly the same as the general sliding-window approach.

### Java Behavior

The Java implementation uses `HashMap<Integer, Integer>`.

Since Java collections work with objects, the integer values are represented using `Integer`.

The frequency update and window movement follow the same logic as the C++ solution.

### JavaScript Behavior

The JavaScript implementation uses `Map`.

`Map` is a good fit because it directly stores key-value pairs and works well when the array values can be large.

The `right` pointer expands the window and the `left` pointer shrinks it when necessary.

### Python3 Behavior

The Python implementation uses a dictionary.

Python dictionaries provide average constant-time access for frequency updates.

The sliding-window logic remains unchanged.

### Go Behavior

The Go implementation uses a map from `int` to `int`.

The map stores each value's frequency, while `left` and `right` control the current window.

The algorithm remains `O(n)` on average.

## Examples

### Example 1

Input:

```text
nums = [1,2,3,1,2,3,1]
k = 2
```

Expected output:

```text
6
```

I start expanding the window from the left.

The first six elements form:

```text
[1,2,3,1,2,3]
```

Their frequencies are:

```text
1 -> 2
2 -> 2
3 -> 2
```

Every value appears at most twice, so the window is valid.

When I add the final `1`, its frequency becomes `3`.

I then move `left` forward until one occurrence of `1` is removed.

The longest valid window has length `6`.

### Example 2

Input:

```text
nums = [1,2,1,2,1,2,1,2]
k = 1
```

Expected output:

```text
2
```

Here every value can appear at most once.

A window such as:

```text
[1,2]
```

is valid.

If I add another `1`, the frequency of `1` becomes `2`, which violates the condition.

So I move `left` forward until that extra occurrence is removed.

The longest possible valid subarray has length `2`.

### Example 3

Input:

```text
nums = [5,5,5,5,5,5,5]
k = 4
```

Expected output:

```text
4
```

There is only one distinct value.

The frequency of `5` can be at most `4`.

So the largest valid window is:

```text
[5,5,5,5]
```

When the fifth `5` enters the window, the frequency becomes `5`.

I move `left` forward and remove one `5`.

The window becomes valid again.

Therefore, the maximum length is `4`.

## How to Use / Run Locally

The solutions in this repository follow the usual LeetCode `class Solution` format. If I want to test them locally, I can add a small driver program with sample input and function calls.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```text
g++ -std=c++17 solution.cpp -o solution
```

Then run:

```text
./solution
```

On Windows, the generated executable can be run with:

```text
solution.exe
```

### Java

Save the solution as:

```text
Solution.java
```

Compile it with:

```text
javac Solution.java
```

Then run:

```text
java Solution
```

For local execution, I need a `main` method that creates test cases and calls `maxSubarrayLength`.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```text
node solution.js
```

A small test harness can be added to create an array, call the function, and print the result.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```text
python3 solution.py
```

A simple test section can call the solution method and print the returned answer.

### Go

Save the solution as:

```text
solution.go
```

Run it with:

```text
go run solution.go
```

For standalone execution, the file needs a `main` function that creates test cases and calls the solution function.

## Notes & Optimizations

The main optimization is replacing brute force with a sliding-window approach.

A brute-force solution can take `O(n²)` time because it may need to examine a large number of subarrays.

The sliding-window solution reduces this to `O(n)` average time by reusing the existing frequency information.

One important edge case is `k = 1`.

In that case, no value can appear twice inside the current window. The algorithm automatically handles this without requiring a separate condition.

Another simple case is when every value already appears at most `k` times in the entire array. Then the complete array is a valid subarray, and the answer is simply `nums.length`.

If all elements are the same, the answer cannot be greater than `k`. The sliding window handles this case by continuously removing elements from the left whenever the frequency becomes too large.

I use a hash map instead of an array for frequencies because the values can be as large as `10^9`. A direct frequency array based on the value itself would waste a huge amount of memory.

The important sliding-window rule to remember is:

```text
Expand from the right.
Shrink from the left when the condition breaks.
Record the best valid window.
```

This pattern is useful for many other DSA problems involving subarrays, frequencies, duplicates, and maximum or minimum window lengths.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
