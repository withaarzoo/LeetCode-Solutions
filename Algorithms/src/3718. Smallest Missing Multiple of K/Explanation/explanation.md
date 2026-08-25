# 3718. Smallest Missing Multiple of K

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

The goal of this LeetCode problem is to find the smallest positive multiple of `k` that does not appear in the integer array `nums`.

A multiple of `k` is any positive number that can be divided evenly by `k`. For example, if `k = 4`, the positive multiples are `4`, `8`, `12`, `16`, and so on.

I need to check these multiples in increasing order and return the first one that is missing from `nums`.

The input contains an integer array `nums` and an integer `k`. The output is a single integer representing the smallest positive multiple of `k` that is not present in the array.

This solution uses a hash set to make membership checks fast, giving an efficient `O(n)` average-case solution for the Smallest Missing Multiple of K problem.

## Constraints

| Constraint   | Value                     |
| ------------ | ------------------------- |
| Array length | `1 <= nums.length <= 100` |
| Array values | `1 <= nums[i] <= 100`     |
| Value of `k` | `1 <= k <= 100`           |

## Intuition

My first observation was that I do not need to sort the array or check every positive number.

I only care about numbers that are multiples of `k`.

So if `k = 3`, I only need to check:

`3, 6, 9, 12, 15, ...`

My first instinct was to start from `k` and keep moving to the next multiple until I find one that is not present in `nums`.

The only thing I needed was a fast way to check whether a number exists in the array. A hash set works well here because it allows average `O(1)` lookup.

Once I find the first missing multiple, I can return it immediately because I am checking the multiples in increasing order.

## Approach

I use the following steps to solve the problem:

1. Store every value from `nums` in a hash set.
2. Start with `k`, which is the smallest positive multiple of `k`.
3. Check whether the current multiple exists in the hash set.
4. If it exists, add `k` to move to the next multiple.
5. Keep repeating this process until I find a multiple that is missing.
6. Return that first missing multiple.

For example, if `nums = [8, 2, 3, 4, 6]` and `k = 2`, I check:

`2 -> present`

`4 -> present`

`6 -> present`

`8 -> present`

`10 -> missing`

So the answer is `10`.

## Data Structures Used

### Hash Set

I use a hash set to store all values from `nums`.

This is useful because I only need to know whether a particular multiple exists or not. I do not need its position or frequency.

A hash set gives average `O(1)` lookup time, which makes checking each multiple fast.

Language-specific implementations include:

* C++: `unordered_set`
* Java: `HashSet`
* JavaScript: `Set`
* Python3: `set`
* Go: `map[int]bool`

## Operations & Behavior Summary

The algorithm works like this:

1. Read all numbers from the input array.
2. Put them into a hash set.
3. Set the current value to `k`.
4. Check whether the current value exists in the set.
5. If it exists, move to the next multiple by adding `k`.
6. Repeat until the current multiple does not exist.
7. Return the missing multiple.

The important part is that I always check multiples in increasing order:

`k, 2k, 3k, 4k, ...`

Because of this order, the first missing value is guaranteed to be the smallest positive multiple of `k` that is absent from `nums`.

## Complexity

| Complexity       | Value          | Explanation                                                                                                                                                   |
| ---------------- | -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` average | `n` is the number of elements in `nums`. Building the hash set takes `O(n)`, and checking multiples takes at most `O(n)` additional checks in the worst case. |
| Space Complexity | `O(n)`         | The hash set stores up to `n` distinct values from `nums`.                                                                                                    |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int missingMultiple(vector<int>& nums, int k) {
        // Store every number so I can check whether a multiple exists quickly.
        unordered_set<int> present(nums.begin(), nums.end());

        // Start from the smallest positive multiple of k.
        int multiple = k;

        // Keep checking multiples until I find one that is missing.
        while (present.count(multiple)) {
            // Move to the next positive multiple of k.
            multiple += k;
        }

        // This is the smallest multiple of k that does not exist in nums.
        return multiple;
    }
};
```

### Java

```java
import java.util.HashSet;
import java.util.Set;

class Solution {
    public int missingMultiple(int[] nums, int k) {
        // Store every number so checking whether a value exists is fast.
        Set<Integer> present = new HashSet<>();

        // Add all values from nums into the hash set.
        for (int num : nums) {
            present.add(num);
        }

        // Start with the smallest positive multiple of k.
        int multiple = k;

        // Continue while the current multiple is already present.
        while (present.contains(multiple)) {
            // Move to the next positive multiple of k.
            multiple += k;
        }

        // Return the first missing multiple.
        return multiple;
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
var missingMultiple = function(nums, k) {
    // Store all values so membership checks are fast.
    const present = new Set(nums);

    // Start with the smallest positive multiple of k.
    let multiple = k;

    // Keep checking consecutive multiples of k.
    while (present.has(multiple)) {
        // Move to the next positive multiple.
        multiple += k;
    }

    // Return the first multiple that is not present.
    return multiple;
};
```

### Python3

```python
from typing import List

class Solution:
    def missingMultiple(self, nums: List[int], k: int) -> int:
        # Store all values so checking whether a multiple exists is fast.
        present = set(nums)

        # Start with the smallest positive multiple of k.
        multiple = k

        # Continue while the current multiple already exists.
        while multiple in present:
            # Move to the next positive multiple of k.
            multiple += k

        # Return the first missing multiple.
        return multiple
```

### Go

```go
func missingMultiple(nums []int, k int) int {
    // Create a map that works like a hash set for fast membership checks.
    present := make(map[int]bool)

    // Store every value from nums in the map.
    for _, num := range nums {
        present[num] = true
    }

    // Start with the smallest positive multiple of k.
    multiple := k

    // Keep checking multiples while the current one exists.
    for present[multiple] {
        // Move to the next positive multiple of k.
        multiple += k
    }

    // Return the first multiple that is missing.
    return multiple
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages. Only the syntax used to create and access a hash set changes.

### Step 1: Store the array values

I first store every number from `nums` in a hash-based data structure.

This avoids searching through the whole array every time I want to check a multiple.

Without a hash set, checking whether a number exists could take `O(n)` time for each multiple. With a hash set, the lookup takes average `O(1)` time.

### Step 2: Start from the smallest positive multiple

I start with `k`.

There is no need to start from `0` because the problem asks for a positive multiple.

There is also no reason to check numbers smaller than `k`, because none of them can be a positive multiple of `k`.

### Step 3: Check whether the current multiple exists

I check whether the current multiple is already present in the hash set.

If it exists, it cannot be the answer, so I move forward.

The membership check is different in syntax across languages:

* C++ uses a lookup in `unordered_set`.
* Java uses `contains`.
* JavaScript uses `has`.
* Python uses the `in` operator.
* Go uses a map lookup.

The behavior is the same in every version.

### Step 4: Move to the next multiple

When the current multiple exists, I add `k`.

For example, when `k = 5`, the sequence becomes:

`5 -> 10 -> 15 -> 20 -> ...`

I do not increase the value by `1`, because that would make me check numbers that are not multiples of `k`.

Adding `k` keeps every checked value valid.

### Step 5: Return the first missing multiple

The loop stops as soon as I find a multiple that does not exist in the hash set.

At that point, I return it immediately.

This works because all smaller positive multiples have already been checked. If the current value is the first missing one in increasing order, it must be the smallest missing multiple of `k`.

### Edge Cases

If `k` itself is missing, I return `k` immediately.

For example:

`nums = [1, 4, 7, 10]`

`k = 3`

Since `3` is not present, the answer is `3`.

Duplicate values also do not affect the solution. A hash set only needs to know whether a value exists at least once.

If several consecutive multiples are present, I simply continue checking until I find the first missing one.

## Examples

### Example 1

**Input**

```text
nums = [8, 2, 3, 4, 6]
k = 2
```

**Expected Output**

```text
10
```

**Trace**

* Check `2` -> present
* Check `4` -> present
* Check `6` -> present
* Check `8` -> present
* Check `10` -> missing

The smallest missing multiple of `2` is `10`.

### Example 2

**Input**

```text
nums = [1, 4, 7, 10, 15]
k = 5
```

**Expected Output**

```text
5
```

**Trace**

* Check `5` -> missing

The first positive multiple of `5` is already missing, so I return `5`.

### Example 3

**Input**

```text
nums = [3, 6, 9, 12, 20]
k = 3
```

**Expected Output**

```text
15
```

**Trace**

* Check `3` -> present
* Check `6` -> present
* Check `9` -> present
* Check `12` -> present
* Check `15` -> missing

The smallest missing multiple of `3` is `15`.

## How to Use / Run Locally

### C++

Save the solution in a file such as `solution.cpp`.

Compile it using:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run the compiled program:

```bash
./solution
```

On Windows, you may run:

```bash
solution.exe
```

### Java

Save the solution in a file named `Solution.java`.

Compile it using:

```bash
javac Solution.java
```

Run it using:

```bash
java Solution
```

Make sure the class and file name follow Java's naming rules if you add a standalone test program.

### JavaScript

Save the solution in a file such as `solution.js`.

Run it with Node.js:

```bash
node solution.js
```

Make sure Node.js is installed on your system before running the file.

### Python3

Save the solution in a file such as `solution.py`.

Run it using:

```bash
python3 solution.py
```

On some systems, this command may also work:

```bash
python solution.py
```

### Go

Save the solution in a file such as `solution.go`.

Run it directly with:

```bash
go run solution.go
```

Or build an executable using:

```bash
go build solution.go
```

Then run the generated executable.

## Notes & Optimizations

The hash set approach is already efficient for this problem.

Sorting the array would also allow me to search for values, but sorting takes `O(n log n)` time and is unnecessary because I only need fast membership checks.

I could repeatedly scan the original array for every multiple, but that could lead to `O(n²)` time in the worst case. Using a hash set avoids that repeated work.

The constraints are small, but the hash set solution still gives a clean and scalable approach.

One useful observation is that the answer must appear after checking at most `n + 1` multiples. Since the array contains only `n` elements, it cannot contain every one of the first `n + 1` distinct multiples.

This makes the hash set solution a simple and efficient choice for the LeetCode 3718 Smallest Missing Multiple of K problem.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
