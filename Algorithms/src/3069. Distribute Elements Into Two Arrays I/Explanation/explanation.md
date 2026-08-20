# 3069. Distribute Elements Into Two Arrays I

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This repository contains a solution for LeetCode 3069, **Distribute Elements Into Two Arrays I**.

The problem gives an array of distinct integers called `nums`. I need to distribute all of its elements between two new arrays, `arr1` and `arr2`.

The first number always goes into `arr1`, and the second number always goes into `arr2`.

For every remaining number, I compare the last element of both arrays:

* If the last element of `arr1` is greater than the last element of `arr2`, I append the current number to `arr1`.
* Otherwise, I append it to `arr2`.

After processing every element, I return the result formed by concatenating `arr1` followed by `arr2`.

This is a straightforward array simulation problem. The main observation is that every decision depends only on the current last elements of the two arrays.

## Constraints

| Constraint                | Description                                        |
| ------------------------- | -------------------------------------------------- |
| `3 <= n <= 50`            | The input array contains between 3 and 50 elements |
| `1 <= nums[i] <= 100`     | Every element is between 1 and 100                 |
| All elements are distinct | No duplicate values exist in `nums`                |

## Intuition

My first instinct was to simulate exactly what the problem describes.

I noticed that I do not need to inspect every element already stored in `arr1` or `arr2`. For each new number, the problem only asks me to compare the last element of both arrays.

So I can start by placing the first two elements into their required arrays and then process the rest one by one.

At every step, I only need two values: the current ending value of `arr1` and the current ending value of `arr2`.

That makes this LeetCode array problem a simple linear simulation.

## Approach

I use the following steps:

1. Create two arrays named `arr1` and `arr2`.
2. Put `nums[0]` into `arr1`.
3. Put `nums[1]` into `arr2`.
4. Start from the third element of `nums`.
5. Compare the last element of `arr1` with the last element of `arr2`.
6. If the last value of `arr1` is greater, append the current number to `arr1`.
7. Otherwise, append the current number to `arr2`.
8. Repeat until every number has been processed.
9. Concatenate `arr1` and `arr2`.
10. Return the final array.

This approach directly follows the rules given in the problem, so there is no need for sorting, dynamic programming, recursion, or any complex algorithm.

## Data Structures Used

### `arr1`

I use this array to store numbers assigned to the first group. I need its last element because future decisions depend on it.

### `arr2`

I use this array to store numbers assigned to the second group. Its last element is also needed for every comparison.

### Result Array

The final answer contains all elements of `arr1` followed by all elements of `arr2`.

Depending on the programming language, the final result can either be created separately or returned using array concatenation.

## Operations & Behavior Summary

The algorithm behaves like this:

1. Initialize two arrays.
2. Send the first number to `arr1`.
3. Send the second number to `arr2`.
4. Take the next unprocessed number.
5. Look at the last value currently stored in both arrays.
6. If `arr1` ends with a larger value, add the number to `arr1`.
7. Otherwise, add it to `arr2`.
8. Continue until the input is finished.
9. Join `arr1` and `arr2`.
10. Return the combined result.

For example, if the input is:

`[5, 4, 3, 8]`

The process starts with:

`arr1 = [5]`

`arr2 = [4]`

For `3`, I compare `5` and `4`. Since `5 > 4`, I add `3` to `arr1`.

Now:

`arr1 = [5, 3]`

`arr2 = [4]`

For `8`, I compare `3` and `4`. Since `3` is not greater than `4`, I add `8` to `arr2`.

The final result is:

`[5, 3, 4, 8]`

## Complexity

| Complexity       | Value  | Explanation                                                                                               |
| ---------------- | ------ | --------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)` | I process each of the `n` elements once, and concatenating the two arrays also takes linear time overall. |
| Space Complexity | `O(n)` | I store the distributed elements in `arr1` and `arr2`, which together contain all `n` input elements.     |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> resultArray(vector<int>& nums) {
        // I create two arrays because the problem asks me to distribute nums into them.
        vector<int> arr1, arr2;

        // I reserve space to avoid unnecessary resizing while elements are being added.
        arr1.reserve(nums.size());
        arr2.reserve(nums.size());

        // The first element must always go into arr1.
        arr1.push_back(nums[0]);

        // The second element must always go into arr2.
        arr2.push_back(nums[1]);

        // I process every remaining element according to the last values of both arrays.
        for (int i = 2; i < nums.size(); i++) {
            // If arr1 ends with a larger value, the current number goes into arr1.
            if (arr1.back() > arr2.back()) {
                arr1.push_back(nums[i]);
            } else {
                // Otherwise, including when both values are equal, it goes into arr2.
                arr2.push_back(nums[i]);
            }
        }

        // I create the final result with enough space for all n elements.
        vector<int> result;
        result.reserve(nums.size());

        // I append arr1 first because the required result starts with arr1.
        result.insert(result.end(), arr1.begin(), arr1.end());

        // I append arr2 after arr1 to complete the required concatenation.
        result.insert(result.end(), arr2.begin(), arr2.end());

        // I return the final distributed array.
        return result;
    }
};
```

### Java

```java
class Solution {
    public int[] resultArray(int[] nums) {
        // I create arr1 and arr2 to simulate the two arrays from the problem.
        int[] arr1 = new int[nums.length];
        int[] arr2 = new int[nums.length];

        // I keep separate sizes because normal Java arrays do not grow automatically.
        int size1 = 1;
        int size2 = 1;

        // The first number must go into arr1.
        arr1[0] = nums[0];

        // The second number must go into arr2.
        arr2[0] = nums[1];

        // I process every remaining number starting from index 2.
        for (int i = 2; i < nums.length; i++) {
            // I compare the current last elements of arr1 and arr2.
            if (arr1[size1 - 1] > arr2[size2 - 1]) {
                // I add the number to arr1 and increase its size.
                arr1[size1++] = nums[i];
            } else {
                // Otherwise, I add the number to arr2 and increase its size.
                arr2[size2++] = nums[i];
            }
        }

        // I create the final array because LeetCode expects an int[] result.
        int[] result = new int[nums.length];

        // I copy all used elements of arr1 first.
        System.arraycopy(arr1, 0, result, 0, size1);

        // I copy all used elements of arr2 immediately after arr1.
        System.arraycopy(arr2, 0, result, size1, size2);

        // I return the concatenated result.
        return result;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var resultArray = function(nums) {
    // I create two arrays to simulate arr1 and arr2 from the problem.
    const arr1 = [nums[0]];
    const arr2 = [nums[1]];

    // I process every number after the first two required operations.
    for (let i = 2; i < nums.length; i++) {
        // If arr1 has the greater last element, I add nums[i] to arr1.
        if (arr1[arr1.length - 1] > arr2[arr2.length - 1]) {
            arr1.push(nums[i]);
        } else {
            // Otherwise, I add nums[i] to arr2.
            arr2.push(nums[i]);
        }
    }

    // I return arr1 followed by arr2, exactly as the problem requires.
    return [...arr1, ...arr2];
};
```

### Python3

```python
class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        # I put the first element into arr1 because the first operation is fixed.
        arr1 = [nums[0]]

        # I put the second element into arr2 because the second operation is fixed.
        arr2 = [nums[1]]

        # I process every remaining element one by one.
        for i in range(2, len(nums)):
            # If arr1 ends with a larger value, I add the current number to arr1.
            if arr1[-1] > arr2[-1]:
                arr1.append(nums[i])
            else:
                # Otherwise, I add the current number to arr2.
                arr2.append(nums[i])

        # I concatenate arr1 and arr2 because this is the required final order.
        return arr1 + arr2
```

### Go

```go
func resultArray(nums []int) []int {
    // I create arr1 with the first number because the first operation is fixed.
    arr1 := []int{nums[0]}

    // I create arr2 with the second number because the second operation is fixed.
    arr2 := []int{nums[1]}

    // I process every remaining number starting from index 2.
    for i := 2; i < len(nums); i++ {
        // I compare the last element currently stored in both arrays.
        if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
            // I append nums[i] to arr1 when its last value is greater.
            arr1 = append(arr1, nums[i])
        } else {
            // Otherwise, I append nums[i] to arr2.
            arr2 = append(arr2, nums[i])
        }
    }

    // I create the final slice with capacity for all elements to reduce reallocations.
    result := make([]int, 0, len(nums))

    // I append arr1 first because the required result starts with arr1.
    result = append(result, arr1...)

    // I append arr2 after arr1 to complete the concatenation.
    result = append(result, arr2...)

    // I return the final result.
    return result
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. The main difference is how each language handles dynamic arrays and how the final result is created.

### C++

I use `vector<int>` for `arr1` and `arr2`.

A C++ vector can grow dynamically, so I can add elements with `push_back()`. I use `back()` to get the last element of each vector.

For every number starting from index `2`, I compare:

`arr1.back()` and `arr2.back()`

If the first value is greater, the current number goes into `arr1`. Otherwise, it goes into `arr2`.

After processing everything, I append all elements of `arr1` to the result and then append `arr2`.

### Java

In Java, I can use either `ArrayList<Integer>` or normal integer arrays.

Using normal arrays is efficient here because the maximum possible size is already known from `nums.length`. I keep separate size variables to know how many positions in each temporary array are currently used.

The last valid element of `arr1` is found using its current size, and the same idea applies to `arr2`.

After processing all numbers, I copy the used elements of `arr1` first and then the used elements of `arr2` into the final result array.

### JavaScript

JavaScript arrays grow dynamically, so the implementation is direct.

I initialize `arr1` with the first number and `arr2` with the second number.

The last element is accessed using:

`array[array.length - 1]`

I use `push()` to append the current number to the correct array.

Finally, I combine both arrays with array concatenation or the spread operator.

### Python3

Python lists make this solution very short.

I initialize the two lists using the first two elements of `nums`.

Python allows me to access the last element using negative indexing:

`arr1[-1]`

and:

`arr2[-1]`

I use `append()` to add each new number.

At the end, `arr1 + arr2` returns a new list containing `arr1` followed by `arr2`.

### Go

In Go, I use slices for `arr1` and `arr2`.

The last element is accessed with:

`arr1[len(arr1)-1]`

and:

`arr2[len(arr2)-1]`

I use `append()` to grow the slices as numbers are distributed.

Finally, I create the result slice and append all elements from `arr1` followed by all elements from `arr2`.

### Why the Same Logic Works Everywhere

The important part is not the language-specific syntax. The actual algorithm always follows the same rule.

At every step, I only compare the two current last elements.

I never need to search through the full arrays because earlier elements do not affect the next decision. The problem only cares about the elements currently at the end of `arr1` and `arr2`.

That is why the solution remains simple and runs in linear time.

## Examples

### Example 1

**Input:**

`nums = [2,1,3]`

**Expected Output:**

`[2,3,1]`

**Trace:**

* Start: `arr1 = [2]`, `arr2 = [1]`
* Compare `2` and `1`
* Since `2 > 1`, add `3` to `arr1`
* `arr1 = [2,3]`
* `arr2 = [1]`
* Final result: `[2,3,1]`

### Example 2

**Input:**

`nums = [5,4,3,8]`

**Expected Output:**

`[5,3,4,8]`

**Trace:**

* Start: `arr1 = [5]`, `arr2 = [4]`
* For `3`, compare `5 > 4`, so add `3` to `arr1`
* `arr1 = [5,3]`, `arr2 = [4]`
* For `8`, compare `3` and `4`
* Since `3` is not greater than `4`, add `8` to `arr2`
* Final result: `[5,3,4,8]`

### Example 3

**Input:**

`nums = [9,2,7,1,6]`

**Expected Output:**

`[9,7,1,2,6]`

**Trace:**

* Start: `arr1 = [9]`, `arr2 = [2]`
* `9 > 2`, so `7` goes into `arr1`
* Arrays become `[9,7]` and `[2]`
* `7 > 2`, so `1` goes into `arr1`
* Arrays become `[9,7,1]` and `[2]`
* `1` is not greater than `2`, so `6` goes into `arr2`
* Final result: `[9,7,1,2,6]`

## How to Use / Run Locally

### C++

Make sure a C++ compiler such as `g++` is installed.

Save the solution in a file such as:

`solution.cpp`

Compile it:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it:

```bash
./solution
```

On Windows, the generated executable may be named `solution.exe`.

### Java

Make sure the Java Development Kit is installed.

Save the solution in:

`Solution.java`

Compile it:

```bash
javac Solution.java
```

Run it:

```bash
java Solution
```

If the code is copied directly from LeetCode, remember that LeetCode usually provides the test environment and calls the `Solution` class automatically. For local testing, a `main` method may be needed.

### JavaScript

Make sure Node.js is installed.

Save the solution in:

`solution.js`

Run it with:

```bash
node solution.js
```

For local testing, add your own function call and input values because LeetCode normally handles the function call automatically.

### Python3

Make sure Python 3 is installed.

Save the solution in:

`solution.py`

Run it:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

For local testing, add a small test case that creates an instance of `Solution` and calls the required method.

### Go

Make sure Go is installed.

Save the solution in:

`solution.go`

Run it directly:

```bash
go run solution.go
```

Or build an executable first:

```bash
go build solution.go
```

Then run the generated executable.

For LeetCode-style Go code, a local `main` function may be needed when testing outside the LeetCode environment.

## Notes & Optimizations

The input size is small, but the direct simulation is already the best approach for this problem.

There is no need to sort the input because the order of processing matters. Sorting would change the behavior of the algorithm and produce a different result.

There is also no need for a nested loop. Comparing every new element against all previously stored elements would add unnecessary work.

The only values needed to make the next decision are the last elements of `arr1` and `arr2`.

The first two operations are fixed, so the algorithm should always initialize the arrays before starting the main loop.

If the last element of `arr1` is not greater than the last element of `arr2`, the current number must go into `arr2`. This means the `else` condition handles every remaining case.

Because all elements in `nums` are distinct, equality between the values being compared will not occur under the given constraints. Even so, following the exact problem rule with an `if` and `else` keeps the implementation clear.

This solution is a good example of a simple DSA simulation problem where carefully following the problem statement is more important than using a complicated algorithm.

## Author

Md Aarzoo Islam — [Instagram](https://www.instagram.com/code.with.aarzoo/?utm_source=chatgpt.com)
