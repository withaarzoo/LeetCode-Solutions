# 2784. Check if Array is Good

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

---

## Problem Summary

In this LeetCode problem, we are given an integer array `nums`.

An array is considered good if it is a permutation of:

```text
[1, 2, 3, ..., n - 1, n, n]
```

This means:

* Numbers from `1` to `n - 1` must appear exactly one time
* The largest number `n` must appear exactly two times

The task is to check whether the given array satisfies this condition or not.

This is a simple array validation and sorting problem that tests understanding of:

* sorting
* frequency checking
* permutations
* edge case handling

The expected output is:

* `true` if the array is good
* `false` otherwise

---

## Constraints

| Constraint                | Value                         |
| ------------------------- | ----------------------------- |
| `1 <= nums.length <= 100` | Array size limit              |
| `1 <= nums[i] <= 200`     | Value range of array elements |

---

## Intuition

The first thing I noticed was that the largest number controls the entire structure of the array.

If the maximum value is `n`, then:

* the array must contain all numbers from `1` to `n - 1`
* and `n` must appear twice

That automatically means:

* the size of the array should always be `n + 1`

After realizing this, the problem became much easier.

I thought sorting the array would help because after sorting, a valid array should always look like this:

```text
[1, 2, 3, ..., n - 1, n, n]
```

So instead of counting frequencies separately, I can directly compare every position with the expected value.

---

## Approach

First, I sort the array in ascending order.

Then I find:

* the size of the array
* the maximum element

After that, I verify:

1. The array length must be equal to `maxElement + 1`
2. Every number from `1` to `maxElement - 1` should appear exactly once
3. The maximum element should appear exactly two times

If any condition fails, I immediately return `false`.

Otherwise, the array is valid.

This approach is simple, clean, and easy to debug.

---

## Data Structures Used

| Data Structure        | Purpose                                 |
| --------------------- | --------------------------------------- |
| Array / Vector / List | Stores the input numbers                |
| Sorting Algorithm     | Helps arrange numbers in expected order |

No extra hash maps or frequency arrays are required in this solution.

---

## Operations & Behavior Summary

1. Sort the array
2. Find the largest element
3. Check if array size matches expected size
4. Verify numbers appear in increasing order
5. Ensure the largest number appears twice
6. Return the final result

This works because a good array always follows a very strict pattern after sorting.

---

## Complexity

| Type             | Complexity   | Explanation                           |
| ---------------- | ------------ | ------------------------------------- |
| Time Complexity  | `O(n log n)` | Sorting the array takes the most time |
| Space Complexity | `O(1)`       | No extra data structures are used     |

Where:

* `n` = size of the input array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool isGood(vector<int>& nums) {
        // Sort the array so numbers come in order
        sort(nums.begin(), nums.end());

        // Total number of elements
        int n = nums.size();

        // Maximum element after sorting
        int mx = nums[n - 1];

        // A good array must have size = mx + 1
        if (n != mx + 1)
            return false;

        // Check first n-1 elements
        // They should be 1, 2, 3, ..., mx
        for (int i = 0; i < n - 1; i++) {

            // Expected value at current position
            if (nums[i] != i + 1)
                return false;
        }

        // Last element must also be mx
        return nums[n - 1] == mx;
    }
};
```

### Java

```java
class Solution {
    public boolean isGood(int[] nums) {

        // Sort the array
        Arrays.sort(nums);

        // Size of array
        int n = nums.length;

        // Largest element
        int mx = nums[n - 1];

        // Valid size should be mx + 1
        if (n != mx + 1) {
            return false;
        }

        // Check sequence 1 to mx
        for (int i = 0; i < n - 1; i++) {

            // Expected number is i + 1
            if (nums[i] != i + 1) {
                return false;
            }
        }

        // Last element should also be mx
        return nums[n - 1] == mx;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var isGood = function(nums) {

    // Sort numbers in ascending order
    nums.sort((a, b) => a - b);

    // Length of array
    let n = nums.length;

    // Maximum element
    let mx = nums[n - 1];

    // Size must be mx + 1
    if (n !== mx + 1) {
        return false;
    }

    // Check numbers from 1 to mx
    for (let i = 0; i < n - 1; i++) {

        // Expected value is i + 1
        if (nums[i] !== i + 1) {
            return false;
        }
    }

    // Last element must also be mx
    return nums[n - 1] === mx;
};
```

### Python3

```python
class Solution:
    def isGood(self, nums: List[int]) -> bool:

        # Sort the array
        nums.sort()

        # Length of array
        n = len(nums)

        # Maximum element
        mx = nums[-1]

        # Size must be mx + 1
        if n != mx + 1:
            return False

        # Check numbers from 1 to mx
        for i in range(n - 1):

            # Expected value is i + 1
            if nums[i] != i + 1:
                return False

        # Last element must also be mx
        return nums[-1] == mx
```

### Go

```go
func isGood(nums []int) bool {

    // Sort the array
    sort.Ints(nums)

    // Length of array
    n := len(nums)

    // Maximum element
    mx := nums[n-1]

    // Size must be mx + 1
    if n != mx+1 {
        return false
    }

    // Check sequence from 1 to mx
    for i := 0; i < n-1; i++ {

        // Expected value is i + 1
        if nums[i] != i+1 {
            return false
        }
    }

    // Last element should also be mx
    return nums[n-1] == mx
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

First, the array is sorted.

Sorting is important because it transforms the array into a predictable order. Without sorting, checking the required pattern would become more complicated.

After sorting:

* the smallest values come first
* the largest value moves to the end

Then the maximum element is stored.

This maximum value represents the candidate value for `n`.

For example:

```text
[1, 2, 3, 3]
```

Here:

* maximum value = `3`
* expected size = `4`

That matches the rule:

```text
n + 1
```

Next, the algorithm checks whether the array length is correct.

If:

```text
length != maxElement + 1
```

then the array cannot be valid.

After that, the algorithm checks every position one by one.

Expected pattern:

```text
index 0 -> 1
index 1 -> 2
index 2 -> 3
...
```

If any number does not match its expected value, the function returns `false`.

Finally, the algorithm ensures the last value is also the maximum number.

That confirms the maximum element appears twice.

This approach avoids unnecessary frequency counting and keeps the implementation short and efficient.

---

## Examples

### Example 1

Input:

```text
nums = [1,3,3,2]
```

Output:

```text
true
```

Explanation:

After sorting:

```text
[1,2,3,3]
```

* Numbers `1` and `2` appear once
* Number `3` appears twice

So the array is good.

---

### Example 2

Input:

```text
nums = [2,1,3]
```

Output:

```text
false
```

Explanation:

Maximum value is `3`.

Expected array size should be:

```text
3 + 1 = 4
```

But current size is only `3`.

So the array is not good.

---

### Example 3

Input:

```text
nums = [1,1]
```

Output:

```text
true
```

Explanation:

Maximum value is `1`.

Expected pattern:

```text
[1,1]
```

The array matches perfectly.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* Sorting is the cleanest approach for this problem
* Constraints are small, so sorting works efficiently
* Another approach could use frequency counting with a hash map
* The sorting approach is easier to understand and implement
* Edge cases like `[1,1]` should always be tested
* Arrays with missing numbers or extra duplicates automatically fail validation

Possible alternative:

* Use a frequency array instead of sorting
* That approach can achieve `O(n)` time complexity
* But it requires additional space

For beginner-friendly competitive programming solutions, the sorting approach is usually the best balance between readability and performance.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
