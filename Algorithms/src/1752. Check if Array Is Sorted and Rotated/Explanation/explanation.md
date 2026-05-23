# 1752. Check if Array Is Sorted and Rotated

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

This LeetCode problem asks whether a given array was originally sorted in non-decreasing order and then rotated some number of times.

A rotated sorted array means the numbers were first sorted, and then elements from the beginning were moved to the end.

Example:

```text
Original Sorted Array:
[1,2,3,4,5]

Rotated Version:
[3,4,5,1,2]
```

The task is to return:

* `true` if the array can become sorted after some rotation
* `false` otherwise

This is a popular array checking problem in Data Structures and Algorithms because it tests observation skills more than complicated logic.

---

## Constraints

| Constraint                | Value               |
| ------------------------- | ------------------- |
| `1 <= nums.length <= 100` | Array size          |
| `1 <= nums[i] <= 100`     | Array element value |

---

## Intuition

When I first looked at this problem, I noticed something important.

A sorted array normally increases from left to right. But after rotation, the order can break exactly one time.

For example:

```text
[3,4,5,1,2]
```

Here:

```text
5 > 1
```

This is the only place where the sorted order breaks.

That observation immediately simplifies the problem.

If the array is valid, the number of decreasing points should never be more than one.

So instead of trying every possible rotation, I only need to count how many times:

```text
nums[i] > nums[i + 1]
```

appears.

If it happens more than once, the array cannot be sorted and rotated.

---

## Approach

I used a simple linear scan approach.

Step-by-step strategy:

1. Traverse the array once
2. Compare every element with the next element
3. Count how many times the current number is greater than the next number
4. Use modulo `% n` to compare the last element with the first element
5. If the count becomes greater than `1`, return `false`
6. Otherwise return `true`

This approach works because:

* A sorted array has `0` decreasing points
* A rotated sorted array has exactly `1` decreasing point
* Any more than that means the array is invalid

This gives an optimal `O(n)` solution.

---

## Data Structures Used

| Data Structure   | Why It Was Used                 |
| ---------------- | ------------------------------- |
| Array            | Input storage                   |
| Integer Variable | Used to count decreasing points |

No extra data structures like stacks, queues, maps, or sets are needed.

---

## Operations & Behavior Summary

The algorithm works in these stages:

1. Start a counter at `0`
2. Loop through every index in the array
3. Compare current element with the next element
4. If the current element is bigger:

   * increase the counter
5. Continue checking all pairs
6. If decreasing points become more than one:

   * stop early
   * return `false`
7. If traversal finishes successfully:

   * return `true`

This makes the solution clean, fast, and memory efficient.

---

## Complexity

| Type             | Complexity | Explanation                              |
| ---------------- | ---------- | ---------------------------------------- |
| Time Complexity  | `O(n)`     | The array is traversed once              |
| Space Complexity | `O(1)`     | No extra memory is used except variables |

Where:

* `n` = length of the input array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool check(vector<int>& nums) {
        
        int n = nums.size();
        
        // Stores how many times the order decreases
        int count = 0;

        // Traverse every element
        for (int i = 0; i < n; i++) {
            
            // Compare current element with next element
            // % n is used so last element compares with first
            if (nums[i] > nums[(i + 1) % n]) {
                count++;
            }

            // More than one decrease means invalid
            if (count > 1) {
                return false;
            }
        }

        // Valid sorted and rotated array
        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean check(int[] nums) {
        
        int n = nums.length;

        // Counts how many times order breaks
        int count = 0;

        // Traverse the array
        for (int i = 0; i < n; i++) {

            // Compare current element with next element
            // % n helps compare last element with first
            if (nums[i] > nums[(i + 1) % n]) {
                count++;
            }

            // More than one break means invalid
            if (count > 1) {
                return false;
            }
        }

        // Array is valid
        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var check = function(nums) {
    
    const n = nums.length;

    // Counts how many decreasing points exist
    let count = 0;

    // Traverse the array
    for (let i = 0; i < n; i++) {

        // Compare current with next element
        // % n connects last element to first
        if (nums[i] > nums[(i + 1) % n]) {
            count++;
        }

        // Invalid if order breaks more than once
        if (count > 1) {
            return false;
        }
    }

    // Valid sorted rotated array
    return true;
};
```

### Python3

```python
class Solution:
    def check(self, nums: List[int]) -> bool:
        
        n = len(nums)

        # Counts how many times order decreases
        count = 0

        # Traverse all indices
        for i in range(n):

            # Compare current element with next element
            # % n helps compare last with first
            if nums[i] > nums[(i + 1) % n]:
                count += 1

            # More than one decrease means invalid
            if count > 1:
                return False

        # Valid sorted rotated array
        return True
```

### Go

```go
func check(nums []int) bool {
    
    n := len(nums)

    // Counts how many times order decreases
    count := 0

    // Traverse the array
    for i := 0; i < n; i++ {

        // Compare current element with next element
        // % n connects last element with first
        if nums[i] > nums[(i+1)%n] {
            count++
        }

        // More than one decrease means invalid
        if count > 1 {
            return false
        }
    }

    // Valid sorted rotated array
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

Only syntax changes.

### Step 1 — Find the array size

The first thing needed is the length of the array because the algorithm traverses every element once.

The size is also important for modulo operations.

---

### Step 2 — Create a counter

A variable is used to count how many times the sorted order breaks.

Example:

```text
[1,2,3,4]
```

No break exists.

Example:

```text
[3,4,5,1,2]
```

One break exists:

```text
5 > 1
```

That single break is valid.

---

### Step 3 — Traverse the array

The algorithm checks every element one by one.

For each index:

```text
nums[i]
```

the next element is checked.

---

### Step 4 — Compare with the next element

The important comparison is:

```text
nums[i] > nums[(i + 1) % n]
```

The modulo operation is the key detail here.

It connects the last element back to the first element.

Without modulo:

* the last element would never be checked against the first
* rotated arrays would fail incorrectly

Example:

```text
[3,4,5,1,2]
```

Comparisons become:

```text
3 <= 4
4 <= 5
5 > 1
1 <= 2
2 <= 3
```

Only one decreasing point exists.

So the array is valid.

---

### Step 5 — Count decreasing points

Whenever:

```text
current > next
```

the counter increases.

If the counter becomes greater than one:

```text
count > 1
```

the function immediately returns `false`.

This early stopping improves efficiency slightly because unnecessary checks are avoided.

---

### Step 6 — Return the final result

If the loop finishes and the count never exceeds one, the array satisfies the sorted and rotated condition.

So the function returns `true`.

---

## Examples

### Example 1

#### Input

```text
nums = [3,4,5,1,2]
```

#### Output

```text
true
```

#### Explanation

Order breaks only once:

```text
5 > 1
```

So this array is a valid rotated sorted array.

---

### Example 2

#### Input

```text
nums = [2,1,3,4]
```

#### Output

```text
false
```

#### Explanation

Two decreasing points appear:

```text
2 > 1
4 > 2
```

Since the order breaks more than once, the array is invalid.

---

### Example 3

#### Input

```text
nums = [1,2,3]
```

#### Output

```text
true
```

#### Explanation

The array is already sorted.

No rotation is also considered valid.

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
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run using Node.js:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* The solution uses a greedy observation-based approach
* Only one traversal is needed
* No sorting is performed
* No extra memory allocation is required
* The modulo trick is the most important part of the solution
* This is one of the cleanest possible solutions for this problem

### Edge Cases

#### Single Element

```text
[1]
```

Always valid.

---

#### Already Sorted

```text
[1,2,3,4]
```

Still valid because zero rotations are allowed.

---

#### Duplicate Values

```text
[2,2,2,3,1]
```

Still works correctly because the problem allows duplicates.

---

#### Multiple Breaks

```text
[3,1,2,0]
```

Invalid because sorted order breaks more than once.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
