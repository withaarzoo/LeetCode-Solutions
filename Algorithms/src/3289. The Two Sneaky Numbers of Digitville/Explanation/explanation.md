# Two Sneaky Numbers of Digitville (LeetCode 3289)

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

In Digitville there was a list `nums` containing integers from `0` to `n-1`. Each number was supposed to appear **exactly once**, but **two** distinct numbers sneaked in and appear an extra time each. The array length is `n + 2`. Our task is to find the two numbers that appear twice and return them in any order.

**Goal:** Return an array (or list) of size two containing the two repeated numbers.

## Constraints

* `2 <= n <= 100`
* `nums.length == n + 2`
* `0 <= nums[i] < n`
* Input guarantees **exactly two** distinct repeated elements

## Intuition

I thought about how the array is constructed: there are `n` distinct possible values (`0` to `n-1`), but the array length is `n+2`. That means exactly two values appear twice and the rest appear once. If I scan the array and mark which values I've already seen, every time I see a value that was already marked it's one of the sneaky numbers. So I can detect both by a single pass.

## Approach

1. Compute `n = nums.length - 2`. Valid values lie in `[0, n-1]`.
2. Create a boolean/flag array `seen` of size `n`, initialized to `false`.
3. Iterate through the `nums` array once.

   * If `seen[x] == false`, set `seen[x] = true`.
   * If `seen[x] == true`, append `x` to `result` — we found a sneaky number.
4. Stop early once we have collected two repeated numbers and return them.

This is simple, easy to reason about, and efficient for the given constraints.

## Data Structures Used

* `seen` boolean/flag array of length `n` to remember which numbers have been seen once.
* `result` list/array to store the two sneaky numbers.

## Operations & Behavior Summary

* Single linear pass over `nums` (O(n) time).
* Mark numbers as seen using constant-time indexing `seen[x]` (O(1) per operation).
* Stop early when both duplicates are found.

## Complexity

* **Time Complexity:** `O(n)` where `n = nums.length - 2`. We visit each element at most once.
* **Space Complexity:** `O(n)` extra space for the `seen` array. The result uses `O(1)` additional space beyond that.

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    vector<int> getSneakyNumbers(vector<int>& nums) {
        int n = (int)nums.size() - 2;          // original range size
        vector<char> seen(n, 0);              // boolean marks (char to save space)
        vector<int> res;
        res.reserve(2);
        for (int x : nums) {
            if (seen[x]) {
                res.push_back(x);             // found a repeated number
            } else {
                seen[x] = 1;                 // mark seen
            }
            if ((int)res.size() == 2) break; // early exit once we have both
        }
        return res;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int[] getSneakyNumbers(int[] nums) {
        int n = nums.length - 2;             // original range size
        boolean[] seen = new boolean[n];     // mark seen values
        int[] res = new int[2];
        int idx = 0;
        for (int x : nums) {
            if (seen[x]) {
                res[idx++] = x;              // repeated -> add to result
                if (idx == 2) break;         // early stop once we have both
            } else {
                seen[x] = true;             // mark seen
            }
        }
        return res;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getSneakyNumbers = function(nums) {
    const n = nums.length - 2;             // original range size
    const seen = new Array(n).fill(false); // boolean marks
    const res = [];
    for (const x of nums) {
        if (seen[x]) {
            res.push(x);                   // duplicate found
            if (res.length === 2) break;   // stop early if both found
        } else {
            seen[x] = true;               // mark as seen
        }
    }
    return res;
};
```

### Python3

```python3
from typing import List

class Solution:
    def getSneakyNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums) - 2                   # original range size
        seen = [False] * n                  # boolean marks for 0..n-1
        res = []
        for x in nums:
            if seen[x]:
                res.append(x)               # found a duplicate
                if len(res) == 2:
                    break                   # early exit
            else:
                seen[x] = True              # mark seen
        return res
```

### Go

```go
package main

func getSneakyNumbers(nums []int) []int {
    n := len(nums) - 2           // original range size
    seen := make([]bool, n)      // marks for 0..n-1
    res := make([]int, 0, 2)
    for _, x := range nums {
        if seen[x] {
            res = append(res, x) // this value is repeated
            if len(res) == 2 {
                break            // both found, stop early
            }
        } else {
            seen[x] = true      // mark seen first time
        }
    }
    return res
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic line-by-line in a language-neutral way, then map to each language.

1. **Compute `n`**

   * `n = nums.length - 2` because the array length is `n + 2` by problem statement. Valid values are `0..n-1`.

2. **Create `seen` array**

   * `seen` has size `n`. All entries start as `false`/`0`. It will track whether a number has been seen once already.

3. **Initialize result container**

   * Prepare an array/list `res` to store the two sneaky numbers. Reserve capacity of 2 when possible.

4. **Single pass over `nums`**

   * For each `x` in `nums`:

     * If `seen[x]` is `false`: set it to `true` (this is the first time we see `x`).
     * If `seen[x]` is `true`: this is the second time we see `x` (so we found a sneaky number) — append `x` to `res`.
     * After appending, if `res` has size 2, we can break early because we have both sneaky numbers.

5. **Return `res`**

   * The `res` list/array now contains the two values that appear twice.

### Mapping to languages (notes)

* **C++**: Use `vector<char>` for seen flags to minimize memory; `vector<int>` for result and `reserve(2)` to avoid reallocations.
* **Java**: Use `boolean[] seen = new boolean[n];` and an `int[] res = new int[2];` to fill results directly.
* **JavaScript**: Use `new Array(n).fill(false)` for `seen` and a dynamic array for `res`.
* **Python3**: Use a `list` for `seen` with `False` initial values and `append` duplicates to `res`.
* **Go**: Use `make([]bool, n)` for `seen` and `append` to build the result slice.

## Examples

1. Input: `nums = [0,1,1,0]` -> Output: `[0,1]` (order may be `[1,0]` depending on discovery order)
2. Input: `nums = [0,3,2,1,3,2]` -> Output: `[2,3]`
3. Input: `nums = [7,1,5,4,3,4,6,0,9,5,8,2]` -> Output: `[4,5]`

## How to use / Run locally

* **C++**: Put the code inside a `.cpp` file and call `getSneakyNumbers` from `main` with a test vector. Compile with `g++ -std=c++17 file.cpp -o run` and execute `./run`.
* **Java**: Place the `Solution` class in a `Solution.java` file, create a `main` method to test it, then compile and run: `javac Solution.java` and `java Solution`.
* **JavaScript (Node.js)**: Put the function in a `.js` file, add test calls and `console.log(getSneakyNumbers([...]))`. Run with `node file.js`.
* **Python3**: Put the class in a `.py` file and call the method with a test list. Run with `python3 file.py`.
* **Go**: Put the function in a `.go` file, create a `main` to call it and print results. Run with `go run file.go`.

## Notes & Optimizations

* The boolean `seen` array uses `O(n)` extra space. For small `n` (<= 100) this is perfectly acceptable.
* An alternate in-place trick (for some related problems) is to mark elements by negation or by adding `n` — but those techniques modify the input and require careful handling; they are not necessary here.
* Since the problem guarantees exactly two duplicates, early exit once we find both is safe and slightly faster.

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
