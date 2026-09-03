# 3876. Construct Uniform Parity Array II

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

I need to determine whether I can construct another array `nums2` of the same length where every element has the same parity. In other words, all elements in `nums2` must be either odd or all even.

For each index `i`, I can choose one of two operations:

1. Keep the original value:
   `nums2[i] = nums1[i]`

2. Subtract another array value:
   `nums2[i] = nums1[i] - nums1[j]`

   Here, `j != i`, and the resulting value must be at least `1`.

The goal is not to actually construct `nums2`. I only need to return `true` if such an array is possible, otherwise return `false`.

The key part of the problem is understanding how subtraction affects parity.

* Even - Even = Even
* Odd - Odd = Even
* Even - Odd = Odd
* Odd - Even = Odd

## Constraints

| Constraint                         | Description                                           |
| ---------------------------------- | ----------------------------------------------------- |
| `1 <= n == nums1.length <= 10^5`   | The array can contain up to 100,000 elements          |
| `1 <= nums1[i] <= 10^9`            | Every value is positive and can be as large as `10^9` |
| `nums1` contains distinct integers | No two elements have the same value                   |

## Intuition

I first focused on the parity of the numbers instead of trying every possible subtraction.

If the array contains only even numbers, I do not need to change anything. The array is already uniform, so the answer is `true`.

The interesting case is when there is at least one odd number.

I noticed that the smallest odd number can never be changed into an even number. To make an odd number even, I would need to subtract an odd number smaller than it. But because it is the smallest odd number, no such value exists.

So, when an odd number exists, the only realistic target parity is odd.

Now I only need to worry about the even numbers.

An even number becomes odd if I subtract an odd number from it. The best odd number to use is the smallest odd number because it is the easiest one to subtract while keeping the result at least `1`.

Therefore, I only need to compare the smallest odd number with the smallest even number.

If:

`smallestOdd < smallestEven`

then every even number is also larger than the smallest odd number. I can subtract the smallest odd number from every even value and make all values odd.

Otherwise, it is impossible.

## Approach

I solve the problem with one pass through the array.

1. I find the smallest odd number.
2. I find the smallest even number.
3. If there is no odd number, all values are already even, so I return `true`.
4. If an odd number exists, I check whether the smallest odd number is smaller than the smallest even number.
5. If it is, every even number can subtract that smallest odd number and become odd.
6. Otherwise, the smallest even number cannot be converted into an odd number, so I return `false`.

This gives me a simple `O(n)` solution without modifying the input array.

## Data Structures Used

No extra data structure is required.

I only use two variables:

* `minOdd` — stores the smallest odd number found in the array.
* `minEven` — stores the smallest even number found in the array.

This is enough because the entire decision depends only on these two values.

## Operations & Behavior Summary

The algorithm can be viewed as the following plain-English pseudocode:

```text
Start with no smallest odd or even value.

For every number in nums1:
    If the number is odd:
        update the smallest odd number.
    Otherwise:
        update the smallest even number.

If there is no odd number:
    return true.

If the smallest odd number is smaller than the smallest even number:
    return true.

Otherwise:
    return false.
```

The important observation is that I never need to test every possible pair of elements.

The smallest odd number is the best candidate for changing even values into odd values. If it cannot be used on the smallest even number, no larger odd number can work either.

## Complexity

| Complexity       | Cost   | Explanation                                                                                          |
| ---------------- | ------ | ---------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I scan all `n` elements exactly once.                                                                |
| Space Complexity | `O(1)` | I only store `minOdd` and `minEven`, so no extra array, map, set, or other data structure is needed. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool uniformArray(vector<int>& nums1) {
        // Store the smallest odd and even values found so far.
        int minOdd = INT_MAX;
        int minEven = INT_MAX;

        // Scan every number once because only the minimum values matter.
        for (int x : nums1) {
            if (x % 2 == 0) {
                // An even value may need to become odd.
                minEven = min(minEven, x);
            } else {
                // The smallest odd value is the best possible value to subtract.
                minOdd = min(minOdd, x);
            }
        }

        // If there is no odd number, every element is already even.
        if (minOdd == INT_MAX) {
            return true;
        }

        // Every even number must be larger than the smallest odd number.
        // Then subtracting minOdd makes every even number odd.
        return minOdd < minEven;
    }
};
```

### Java

```java
class Solution {
    public boolean uniformArray(int[] nums1) {
        // Use a large value as "not found yet" for both parities.
        int minOdd = Integer.MAX_VALUE;
        int minEven = Integer.MAX_VALUE;

        // Scan the array once to find the smallest odd and even values.
        for (int x : nums1) {
            if (x % 2 == 0) {
                // Keep the smallest even number because it is the hardest one to change.
                minEven = Math.min(minEven, x);
            } else {
                // Keep the smallest odd number because it is the best number to subtract.
                minOdd = Math.min(minOdd, x);
            }
        }

        // If there is no odd number, all elements are already even.
        if (minOdd == Integer.MAX_VALUE) {
            return true;
        }

        // Every even value must be greater than the smallest odd value.
        // Subtracting that odd value then makes every even value odd.
        return minOdd < minEven;
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
    // Use Infinity to represent that no odd or even value has been found yet.
    let minOdd = Infinity;
    let minEven = Infinity;

    // Scan every value once to find the minimum value of each parity.
    for (const x of nums1) {
        if (x % 2 === 0) {
            // The smallest even value is the hardest even value to convert.
            minEven = Math.min(minEven, x);
        } else {
            // The smallest odd value is the best value available for subtraction.
            minOdd = Math.min(minOdd, x);
        }
    }

    // If there are no odd values, all numbers are already even.
    if (minOdd === Infinity) {
        return true;
    }

    // Every even value must be larger than minOdd so it can subtract minOdd.
    return minOdd < minEven;
};
```

### Python3

```python
class Solution:
    def uniformArray(self, nums1: list[int]) -> bool:
        # Use infinity to represent that no odd or even value has been found yet.
        min_odd = float('inf')
        min_even = float('inf')

        # Scan the array once because only the smallest values of each parity matter.
        for x in nums1:
            if x % 2 == 0:
                # The smallest even value is the hardest one to convert to odd.
                min_even = min(min_even, x)
            else:
                # The smallest odd value is the best value to subtract.
                min_odd = min(min_odd, x)

        # If there is no odd number, all elements are already even.
        if min_odd == float('inf'):
            return True

        # Every even value must be greater than the smallest odd value.
        # Then subtracting min_odd makes that even value odd.
        return min_odd < min_even
```

### Go

```go
func uniformArray(nums1 []int) bool {
 // Use a value larger than every possible input as the initial minimum.
 const inf = int(1e18)

 // Store the smallest odd and even values found in the array.
 minOdd := inf
 minEven := inf

 // Scan the array once because only the two minimum values are needed.
 for _, x := range nums1 {
  if x%2 == 0 {
   // Keep the smallest even value because it is the hardest to change.
   if x < minEven {
    minEven = x
   }
  } else {
   // Keep the smallest odd value because it is the best subtraction value.
   if x < minOdd {
    minOdd = x
   }
  }
 }

 // If no odd value exists, all numbers are already even.
 if minOdd == inf {
  return true
 }

 // Every even value must be larger than the smallest odd value.
 // Then subtracting minOdd changes each even value into an odd value.
 return minOdd < minEven
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The implementation in all five languages follows the same idea. The syntax changes, but the algorithm stays exactly the same.

### Step 1: Store the minimum odd and even values

I start with two variables:

```text
minOdd
minEven
```

Both are initially set to a very large value.

This lets me update them whenever I find an actual odd or even number.

For example, if:

```text
nums1 = [7, 4, 10, 3, 8]
```

after scanning the array:

```text
minOdd = 3
minEven = 4
```

I do not need to remember the other values.

### Step 2: Check every element

For each value, I check whether it is even or odd.

A number is even when:

```text
number % 2 == 0
```

Otherwise, it is odd.

If the number is even, I compare it with `minEven` and keep the smaller value.

If the number is odd, I compare it with `minOdd` and keep the smaller value.

This works the same way in C++, Java, JavaScript, Python3, and Go.

### Step 3: Handle an array containing only even numbers

After the scan, if `minOdd` was never updated, the array contains no odd values.

That means every element is already even.

For example:

```text
[2, 4, 6, 8]
```

I can simply keep every element unchanged.

So the answer is:

```text
true
```

### Step 4: Understand why an existing odd number forces an odd result

Suppose the smallest odd number is `3`.

To make `3` even, I would need to subtract an odd number smaller than `3`.

But there is no smaller odd number in the array.

I could subtract an even number, but:

```text
odd - even = odd
```

So the smallest odd value cannot become even.

Because of this, whenever an odd number exists, I must make the final array all odd.

### Step 5: Convert even numbers into odd numbers

To change an even number into an odd number, I need to subtract an odd number.

For example:

```text
4 - 1 = 3
10 - 1 = 9
```

Both results are odd.

The smallest odd number is the best candidate because it is smaller than every other odd number.

If the smallest even number is greater than this smallest odd number, then every even number is also greater than it.

For example:

```text
nums1 = [1, 4, 7, 10]
```

Here:

```text
minOdd = 1
minEven = 4
```

Since:

```text
1 < 4
```

I can transform the even values:

```text
4  -> 4 - 1  = 3
10 -> 10 - 1 = 9
```

The odd values can remain unchanged:

```text
1 -> 1
7 -> 7
```

So the resulting array can be:

```text
[1, 3, 7, 9]
```

All values are odd.

### Step 6: Detect the impossible case

Now consider:

```text
nums1 = [2, 3]
```

The smallest values are:

```text
minEven = 2
minOdd = 3
```

The smallest even number is `2`.

To make `2` odd, I would need to subtract an odd number smaller than `2`.

But the only odd number is `3`.

Using it would give:

```text
2 - 3 = -1
```

This is not allowed because the result must be at least `1`.

Therefore, I cannot make the array all odd.

I also cannot make `3` even because it is the smallest odd number.

So the answer is `false`.

### Step 7: Final comparison

The entire decision becomes:

```text
If there is no odd number:
    true

Otherwise:
    true if minOdd < minEven
    false otherwise
```

The same condition is used in all five implementations.

There is no need for sorting because sorting the array would take `O(n log n)`, while I can find both minimum values directly in `O(n)`.

There is also no need for a set, map, or second array. The input contains distinct integers, but the solution does not need to use that property explicitly.

## Examples

### Example 1

Input:

```text
nums1 = [1, 4, 7]
```

The smallest odd number is:

```text
1
```

The smallest even number is:

```text
4
```

Since:

```text
1 < 4
```

I can subtract `1` from `4`:

```text
4 - 1 = 3
```

The other values can stay unchanged:

```text
[1, 3, 7]
```

All values are odd.

Output:

```text
true
```

### Example 2

Input:

```text
nums1 = [2, 3]
```

The smallest even number is `2`.

The smallest odd number is `3`.

Since:

```text
3 > 2
```

the smallest even value cannot subtract a smaller odd number.

So I cannot make all elements odd.

I also cannot make all elements even because the smallest odd value cannot become even.

Output:

```text
false
```

### Example 3

Input:

```text
nums1 = [4, 6]
```

There are no odd numbers.

Both values are already even:

```text
[4, 6]
```

No operation is required.

Output:

```text
true
```

## How to Use / Run Locally

The LeetCode versions are written using the platform's required class and function format. To run them locally, place the solution inside a small driver program that creates an input array and calls the function.

### C++

Save the solution as `solution.cpp`.

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, the generated executable can be run with:

```bash
solution.exe
```

### Java

Save the solution as `Solution.java`.

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

A local Java program should contain a `main` method to create the test input and call `uniformArray`.

### JavaScript

Save the solution as `solution.js`.

Run it with Node.js:

```bash
node solution.js
```

You can create an input array in the file and call the function to print the result.

### Python3

Save the solution as `solution.py`.

Run it with:

```bash
python3 solution.py
```

You can create a test array and call `Solution().uniformArray(nums1)` from a small driver section.

### Go

Save the solution as `solution.go`.

Run it directly with:

```bash
go run solution.go
```

For a standalone Go program, add a `main` function that creates the input array and calls `uniformArray`.

## Notes & Optimizations

The main optimization is avoiding sorting.

A sorting-based solution could find the required values after sorting, but that would take `O(n log n)` time. I only need the smallest odd and smallest even values, so a single scan is enough.

The important edge cases are:

* If all numbers are even, the answer is `true`.
* If there is an odd number but no smaller odd number for the smallest even value, the answer is `false`.
* A single-element array is always valid because I can keep its only element unchanged.
* The input values are distinct, so I never have to deal with choosing the same index for the subtraction operation.

The final condition is very small:

```text
No odd number -> true
Otherwise -> minOdd < minEven
```

This gives an optimal `O(n)` time and `O(1)` extra space solution for the problem.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
