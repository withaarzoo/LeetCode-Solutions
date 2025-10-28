# 3354. Make Array Elements Equal to Zero

---

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given an integer array `nums`. You must choose a starting index `curr` such that `nums[curr] == 0`, and choose an initial movement direction (left or right). Then you repeat the following until `curr` is out of range:

* If `nums[curr] == 0`, move one step in the current direction.
* Else (`nums[curr] > 0`):

  * Decrement `nums[curr]` by 1.
  * Reverse your movement direction (left <> right).
  * Take one step in the new direction.

A selection (starting index + direction) is **valid** if after the process ends every element of `nums` becomes `0`. Return the number of valid selections.

## Constraints

* `1 <= nums.length <= 100`
* `0 <= nums[i] <= 100`
* There is at least one element `i` where `nums[i] == 0`.

These limits are small enough that a clear simulation is acceptable in performance and easy to reason about.

---

## Intuition

I thought about what the process actually does. Starting from a zero cell, I either move left or right. When I land on a cell with value `> 0`, I reduce it by `1`, reverse direction, and step once. When I land on a `0`, I keep going in the same direction. This is deterministic and easy to simulate. Since `n <= 100` and values are `<= 100`, the total number of decrements across the whole array is bounded by `sum(nums)` and simulation for each candidate is efficient. So I decided to try every zero index and simulate both starting directions (left and right) to check if the array becomes all zeros.

---

## Approach

1. Find all indices `i` where `nums[i] == 0`.
2. For each such index `i`, simulate the described process twice:

   * Once starting with direction left (dir = -1).
   * Once starting with direction right (dir = +1).
3. During simulation, work on a copy of `nums` so that each trial is independent.
4. Walk using `curr` and `dir` until `curr` goes out of bounds. Apply the exact rules inside the loop.
5. After the walker leaves the array, check whether all elements in the copy are `0`. If yes, this (start, direction) is valid.
6. Sum all valid trials and return the total.

This straightforward simulation is easy to implement and reason about.

---

## Data Structures Used

* Plain arrays / vectors / slices depending on language.
* A copy of the input array for each simulation trial so we can modify values without affecting other trials.

No special data structure is required.

---

## Operations & Behavior Summary

* Copying the array: O(n)
* Simulating walker steps: Each decrement reduces the global sum by 1. The number of non-trivial steps (decrements) is at most `sum(nums)`; zero-steps only move the cursor and are bounded by array traversal behavior. With given constraints direct simulation is safe and simple.

---

## Complexity

* **Time Complexity:** O(n * S)

  * `n` = number of elements in `nums`.
  * `S` = total sum of elements in `nums` (worst-case number of decrements).
  * Explanation: For each zero starting index (at most `n`) and for both directions (2), we simulate at most `O(S)` steps. Given constraints (n <= 100 and elements <= 100), this is efficient.

* **Space Complexity:** O(n)

  * We make a copy of the array for each simulation. This uses O(n) extra memory. We can also reuse a single buffer and reset it between runs if we want a small micro-optimization.

---

## Multi-language Solutions

Below are polished, ready-to-run implementations in multiple languages. They follow the same simulation approach.

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    // simulate from start index with direction dir (-1 left, +1 right)
    bool simulate(const vector<int>& nums, int start, int dir) {
        int n = nums.size();
        vector<int> a = nums;              // copy to mutate
        int curr = start;
        while (curr >= 0 && curr < n) {
            if (a[curr] == 0) {
                curr += dir;              // move in same direction
            } else {
                a[curr]--;                // decrement
                dir = -dir;               // reverse direction
                curr += dir;              // step in new direction
            }
        }
        // check if all zero
        for (int v : a) if (v != 0) return false;
        return true;
    }

    int countValidSelections(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            if (nums[i] != 0) continue;   // start must be a zero
            if (simulate(nums, i, -1)) ++ans; // left
            if (simulate(nums, i, +1)) ++ans; // right
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    // helper: simulate starting at start with dir (-1 left, +1 right)
    private boolean simulate(int[] nums, int start, int dir) {
        int n = nums.length;
        int[] a = nums.clone(); // copy to mutate
        int curr = start;
        while (curr >= 0 && curr < n) {
            if (a[curr] == 0) {
                curr += dir; // move same direction
            } else {
                a[curr]--;   // decrement
                dir = -dir;  // reverse direction
                curr += dir; // step in new direction
            }
        }
        // check all zero
        for (int v : a) if (v != 0) return false;
        return true;
    }

    public int countValidSelections(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            if (nums[i] != 0) continue;
            if (simulate(nums, i, -1)) ans++; // left
            if (simulate(nums, i, +1)) ans++; // right
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var countValidSelections = function(nums) {
    const n = nums.length;

    const simulate = (arr, start, dir) => {
        const a = arr.slice(); // copy
        let curr = start;
        while (curr >= 0 && curr < n) {
            if (a[curr] === 0) {
                curr += dir; // move same direction
            } else {
                a[curr] -= 1; // decrement
                dir = -dir;   // reverse
                curr += dir;  // step in new direction
            }
        }
        // check all zeros
        for (let v of a) if (v !== 0) return false;
        return true;
    };

    let ans = 0;
    for (let i = 0; i < n; ++i) {
        if (nums[i] !== 0) continue;
        if (simulate(nums, i, -1)) ans++; // left
        if (simulate(nums, i, +1)) ans++; // right
    }
    return ans;
};
```

### Python3

```python3
class Solution:
    def simulate(self, nums: list[int], start: int, dir: int) -> bool:
        n = len(nums)
        a = nums.copy()   # copy so original remains unchanged
        curr = start
        while 0 <= curr < n:
            if a[curr] == 0:
                curr += dir    # move same direction
            else:
                a[curr] -= 1   # decrement
                dir = -dir     # reverse direction
                curr += dir    # step in new direction
        # check all zero
        return all(v == 0 for v in a)

    def countValidSelections(self, nums: list[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(n):
            if nums[i] != 0:
                continue
            if self.simulate(nums, i, -1):
                ans += 1   # start i, go left
            if self.simulate(nums, i, +1):
                ans += 1   # start i, go right
        return ans
```

### Go

```go
package main

// countValidSelections simulates the process for each zero index and both directions
func simulate(nums []int, start int, dir int) bool {
    n := len(nums)
    a := make([]int, n)
    copy(a, nums)
    curr := start
    for curr >= 0 && curr < n {
        if a[curr] == 0 {
            curr += dir // move same direction
        } else {
            a[curr]--    // decrement
            dir = -dir   // reverse direction
            curr += dir  // step in new direction
        }
    }
    for _, v := range a {
        if v != 0 {
            return false
        }
    }
    return true
}

func countValidSelections(nums []int) int {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            continue
        }
        if simulate(nums, i, -1) {
            ans++ // left
        }
        if simulate(nums, i, 1) {
            ans++ // right
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Below I explain the main idea and map the algorithm to the key lines that implement it. The logic is the same across languages — only syntax changes.

1. **Identify valid start positions**

   * We only try indices `i` where `nums[i] == 0`.
   * In code: loop over `i` and `continue` if `nums[i] != 0`.

2. **Simulate**

   * Make a copy of the original array so we can mutate it safely.
   * Initialize `curr = start` and `dir = -1` (left) or `dir = 1` (right).
   * While `curr` is in bounds:

     * If `a[curr] == 0` then `curr += dir` (keep moving).
     * Else `a[curr] -= 1`, `dir = -dir`, `curr += dir` (decrement, reverse, step).
   * After loop, verify all elements in the copy are zero.

3. **Count results**

   * If the simulation returned `true`, increment the answer.
   * Repeat for both directions.

4. **Return** the final count.

The code blocks in "Multi-language Solutions" show the exact implementation per language with comments.

---

## Examples

* **Example 1**

  * Input: `nums = [1,0,2,0,3]`
  * Output: `2`
  * Explanation: Starting at index `3` (0-based) and moving left, or starting at index `3` moving right lead to all zeros by the time the walker leaves the array.

* **Example 2**

  * Input: `nums = [2,3,4,0,4,1,0]`
  * Output: `0`
  * Explanation: No valid starting index + direction results in all zeros.

---

## How to use / Run locally

* **C++**: Put the `Solution` class in a file, compile with `g++ -std=c++17 file.cpp` and call the method from `main()` with test vectors.
* **Java**: Place the `Solution` class in a file `Solution.java`, add a `main` to test the method, compile `javac Solution.java` and run `java Solution`.
* **JavaScript (Node)**: Put the function in a `.js` file and call it with test arrays. Run with `node file.js`.
* **Python3**: Save the class into a `.py` file, add test code at the bottom to instantiate `Solution()` and call `countValidSelections`. Run `python3 file.py`.
* **Go**: Place code in `main` package, add `main()` that calls `countValidSelections` with test input, run `go run file.go`.

Example (Python quick test):

```python
if __name__ == "__main__":
    s = Solution()
    print(s.countValidSelections([1,0,2,0,3]))  # expected 2
```

---

## Notes & Optimizations

* The straightforward simulation is sufficient given the constraints. If `n` and values were much larger, we would need to analyze the walker's pattern for pruning or use more clever invariants.
* Micro-optimizations:

  * Reuse a single buffer and restore it between simulations to avoid repeated allocation.
  * Early exit during simulation if current partial state proves that some index cannot reach zero (not trivial to detect in constant time).
* The important correctness idea is: every time we decrement, the global sum decreases by 1. This bounds the number of non-trivial steps.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
