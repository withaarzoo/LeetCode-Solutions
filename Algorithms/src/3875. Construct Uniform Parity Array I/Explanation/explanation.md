# 3875. Construct Uniform Parity Array I

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

In this problem, I am given an array `nums1` containing `n` distinct integers.

I need to construct another array `nums2` of the same length. Every element of `nums2` must have the same parity, meaning either all elements must be even or all elements must be odd.

For every index `i`, I have exactly two choices:

* Keep the current value: `nums2[i] = nums1[i]`
* Subtract another element from the array: `nums2[i] = nums1[i] - nums1[j]`, where `j != i`

I need to return `true` if it is possible to make all elements in `nums2` either even or odd. Otherwise, I return `false`.

The main idea behind this LeetCode problem is to understand how subtraction affects the parity of integers.

## Constraints

* `1 <= n == nums1.length <= 100`
* `1 <= nums1[i] <= 100`
* `nums1` consists of distinct integers.

## Intuition

I started by looking at what happens to parity when I subtract two numbers.

The important parity rules are:

* Even - Even = Even
* Odd - Odd = Even
* Even - Odd = Odd
* Odd - Even = Odd

This immediately gives me a useful observation.

If every number in `nums1` is already even, I can simply keep every number unchanged. The resulting array will contain only even numbers.

If there is at least one odd number, I can use that odd number to change every even number into an odd number.

For the odd number itself, I simply keep it unchanged.

For example, consider:

`nums1 = [2, 5, 8]`

I can keep `5` as it is.

Then I can subtract `5` from the even numbers:

* `2 - 5 = -3`, which is odd
* `8 - 5 = 3`, which is odd

So I can create an array where every element is odd.

This means a valid construction always exists.

Because of that, the solution does not need to actually construct `nums2`. I can directly return `true`.

## Approach

I use the following reasoning:

1. If all elements are even, I keep every element unchanged.
2. If at least one element is odd, I keep one odd element unchanged.
3. I use that odd element to subtract from every even element.
4. Subtracting an odd number from an even number produces an odd number.
5. Therefore, every element can have odd parity.
6. Since one of these two constructions always works, the answer is always `true`.

The important point is that I do not need to check the array or perform any actual operations.

## Data Structures Used

No data structures are required.

I do not need:

* An extra array
* A hash set
* Sorting
* A stack or queue
* Any other auxiliary data structure

The input array itself is not modified.

## Operations & Behavior Summary

The algorithm can be summarized as:

1. Consider the parity of the input numbers.
2. If all numbers are even, keep them unchanged.
3. Otherwise, choose any odd number as the number to subtract.
4. Keep that odd number unchanged.
5. Subtract the chosen odd number from every even number.
6. All resulting numbers are odd.
7. Therefore, return `true`.

In practice, I do not even need to perform these steps in the code because the reasoning proves that the answer is always possible.

## Complexity

| Complexity       | Result | Explanation                                            |
| ---------------- | ------ | ------------------------------------------------------ |
| Time Complexity  | `O(1)` | I do not need to traverse the `n` elements of `nums1`. |
| Space Complexity | `O(1)` | I do not create any extra array or data structure.     |

Here, `n` represents the number of elements in `nums1`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool uniformArray(vector<int>& nums1) {
        // A valid uniform-parity array can always be constructed,
        // so no traversal or extra data structure is necessary.
        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean uniformArray(int[] nums1) {
        // A valid uniform-parity array always exists,
        // so I can return true without processing the array.
        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums1
 * @return {boolean}
 */
var uniformArray = function(nums1) {
    // A uniform-parity array can always be constructed,
    // so there is no need to inspect or modify the input.
    return true;
};
```

### Python3

```python
class Solution:
    def uniformArray(self, nums1: list[int]) -> bool:
        # A valid array with all elements having the same parity
        # always exists, so I can directly return True.
        return True
```

### Go

```go
func uniformArray(nums1 []int) bool {
 // A uniform-parity array is always possible,
 // so I do not need to inspect the input array.
 return true
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### C++

The C++ solution only needs to return a boolean value.

I do not need to loop through `nums1` because I have already proved that a valid uniform parity array can always be constructed.

If all numbers are even, keeping them unchanged gives an all-even array.

If there is an odd number, I can keep that number unchanged and subtract it from every even number. This makes all the resulting values odd.

Because one of these cases must always happen, the C++ function returns `true`.

The vector parameter is only the input required by the LeetCode function signature. I do not need to modify it or create another vector.

### Java

The Java solution follows exactly the same reasoning.

The input is provided as an integer array.

I do not need to inspect every element because the mathematical observation already guarantees that the construction is possible.

If the input contains only even numbers, I can keep every value as it is.

If the input contains an odd number, I can use that odd number to turn every even number into an odd number through subtraction.

The method therefore returns `true` directly.

### JavaScript

The JavaScript implementation also needs no iteration or additional data structure.

The input is an array of numbers, but I do not need to modify it.

The important part is the parity property:

`even - odd = odd`

So if an odd value exists, it can be used to make every even value odd. The odd value itself can remain unchanged.

If no odd value exists, every value is already even.

Therefore, the function can directly return `true`.

### Python3

The Python solution uses the same constant-time observation.

The list `nums1` is received as input, but I do not need to access its individual elements.

There are only two possibilities:

* All values are even.
* At least one value is odd.

The first case already gives an all-even result.

In the second case, one odd value can be kept unchanged and used to transform all even values into odd values.

So the method always returns `True`.

### Go

The Go solution follows the same simple logic.

The input is received as a slice of integers.

No loop is necessary because the result does not depend on the exact values in the array. The parity rules guarantee that a valid construction exists for every valid input.

Therefore, the function returns `true` immediately.

## Examples

### Example 1

**Input:**

```text
nums1 = [2, 3]
```

**Output:**

```text
true
```

I can keep `3` unchanged because it is odd.

For `2`, I can subtract `3`:

`2 - 3 = -1`

Now the possible `nums2` is:

`[-1, 3]`

Both values are odd, so the answer is `true`.

### Example 2

**Input:**

```text
nums1 = [4, 6]
```

**Output:**

```text
true
```

Both values are already even.

I can simply keep them unchanged:

`nums2 = [4, 6]`

Every element is even, so the answer is `true`.

### Example 3

**Input:**

```text
nums1 = [2, 5, 8]
```

**Output:**

```text
true
```

I can keep `5` unchanged.

Then I can subtract `5` from the even values:

`2 - 5 = -3`

`8 - 5 = 3`

This gives:

`nums2 = [-3, 5, 3]`

All elements are odd, so the construction is valid.

## How to Use / Run Locally

The code in this repository is designed for the LeetCode function format. To run it locally, place the corresponding solution inside a small driver program and provide your own test cases.

### C++

Save the solution in a file such as `solution.cpp`.

Compile it using:

```bash
```

Run the compiled program using:

```bash
```

### Java

Save the solution in `Solution.java`.

Compile it using:

```bash
```

Run it using:

```bash
```

### JavaScript

Save the solution in `solution.js`.

Run it with Node.js using:

```bash
```

### Python3

Save the solution in `solution.py`.

Run it using:

```bash
```

### Go

Save the solution in a Go file such as `solution.go`.

Run it using:

```bash
```

The exact LeetCode submission does not require a custom `main` function because LeetCode calls the required method automatically.

## Notes & Optimizations

The most important optimization here is recognizing that no computation is actually necessary.

A common first approach might be to iterate through the array, count even and odd values, and then decide which parity can be created. That would still work, but it is unnecessary.

The stronger observation is that:

* If there are no odd numbers, the array is already uniformly even.
* If there is at least one odd number, that odd number can be used to turn every even number into an odd number.
* The chosen odd number itself can simply be left unchanged.

So the answer is always `true`.

There is also no issue with negative results. For example:

`2 - 5 = -3`

`-3` is still odd, so the sign does not affect the parity.

The constraint that `j != i` also does not cause a problem. When using an odd number to transform other elements, I use it only for the other indices. For the odd element itself, I choose the first operation and keep it unchanged.

This makes the final solution both simpler and faster than explicitly constructing the resulting array.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
