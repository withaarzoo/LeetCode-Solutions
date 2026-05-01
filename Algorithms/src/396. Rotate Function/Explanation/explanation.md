# Rotate Function (LeetCode 396) – Optimal O(n) Solution Explained

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

This problem is about finding the maximum value of a special function after rotating an array.

You are given an integer array `nums`. You can rotate it clockwise any number of times. For each rotation, you calculate a value like this:

F(k) = 0 * arr[0] + 1 * arr[1] + 2 * arr[2] + ... + (n - 1) * arr[n - 1]

Here, `arr` is the array after rotating `nums` by `k` positions.

Your goal is to return the maximum value among all possible rotations.

In short:

* Input: integer array
* Output: maximum rotation function value

This is a classic array rotation + optimization problem that appears often in coding interviews.

## Constraints

* 1 ≤ n ≤ 100000
* -100 ≤ nums[i] ≤ 100
* Result fits in a 32-bit integer

## Intuition

At first, I thought about rotating the array every time and recalculating the function. But that clearly takes too long because each rotation is O(n), and there are n rotations.

So I started looking for a pattern.

When I compared two consecutive rotations, I noticed something interesting:

* Most elements just shift one step to the right
* Only one element jumps from the end to the front

That means instead of recalculating everything, I can reuse the previous result and adjust it.

That’s where the idea of a recurrence relation comes in.

## Approach

Here’s how I solved it step by step:

1. First, calculate the total sum of all elements in the array.
2. Then compute F(0), which is the value without any rotation.
3. Now instead of rotating the array, I update the value using a formula.

The key formula I use is:

F(k) = F(k-1) + total_sum - n * last_element

Where:

* `total_sum` is the sum of all elements
* `last_element` is the element that moves from the end to the front

4. I repeat this for all rotations from 1 to n-1.
5. While doing this, I keep track of the maximum value.

This avoids unnecessary recomputation and gives an optimal solution.

## Data Structures Used

* Integer variables
  Used to store sum, current function value, and result

No extra data structures like arrays or maps are needed. Everything is done in-place using variables.

## Operations & Behavior Summary

* Compute total sum of array elements
* Compute initial rotation value F(0)
* Loop through the array to simulate rotations
* Update value using mathematical relation
* Track the maximum result during iteration
* Return the maximum value at the end

## Complexity

| Type             | Complexity | Explanation                                         |
| ---------------- | ---------- | --------------------------------------------------- |
| Time Complexity  | O(n)       | I loop through the array a constant number of times |
| Space Complexity | O(1)       | No extra space is used apart from a few variables   |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxRotateFunction(vector<int>& nums) {
        int n = nums.size();
        
        long sum = 0;     // Total sum of elements
        long F = 0;       // Initial rotation value F(0)
        
        // Step 1: Calculate total sum and F(0)
        for(int i = 0; i < n; i++) {
            sum += nums[i];          // accumulate total sum
            F += (long)i * nums[i];  // compute F(0)
        }
        
        long result = F; // store max result
        
        // Step 2: Use recurrence relation to compute next values
        for(int k = 1; k < n; k++) {
            // Transition from F(k-1) to F(k)
            F = F + sum - (long)n * nums[n - k];
            
            result = max(result, F); // update maximum
        }
        
        return (int)result;
    }
};
```

### Java

```java
class Solution {
    public int maxRotateFunction(int[] nums) {
        int n = nums.length;
        
        long sum = 0; // total sum
        long F = 0;   // F(0)
        
        // Step 1: compute sum and F(0)
        for(int i = 0; i < n; i++) {
            sum += nums[i];
            F += (long)i * nums[i];
        }
        
        long result = F;
        
        // Step 2: compute next rotations
        for(int k = 1; k < n; k++) {
            F = F + sum - (long)n * nums[n - k];
            result = Math.max(result, F);
        }
        
        return (int)result;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxRotateFunction = function(nums) {
    let n = nums.length;
    
    let sum = 0; // total sum
    let F = 0;   // F(0)
    
    // Step 1: compute sum and F(0)
    for(let i = 0; i < n; i++) {
        sum += nums[i];
        F += i * nums[i];
    }
    
    let result = F;
    
    // Step 2: compute next rotations
    for(let k = 1; k < n; k++) {
        F = F + sum - n * nums[n - k];
        result = Math.max(result, F);
    }
    
    return result;
};
```

### Python3

```python
class Solution:
    def maxRotateFunction(self, nums: List[int]) -> int:
        n = len(nums)
        
        total_sum = 0  # total sum
        F = 0          # F(0)
        
        # Step 1: compute sum and F(0)
        for i in range(n):
            total_sum += nums[i]
            F += i * nums[i]
        
        result = F
        
        # Step 2: compute next rotations
        for k in range(1, n):
            F = F + total_sum - n * nums[n - k]
            result = max(result, F)
        
        return result
```

### Go

```go
func maxRotateFunction(nums []int) int {
    n := len(nums)
    
    var sum int64 = 0 // total sum
    var F int64 = 0   // F(0)
    
    // Step 1: compute sum and F(0)
    for i := 0; i < n; i++ {
        sum += int64(nums[i])
        F += int64(i) * int64(nums[i])
    }
    
    result := F
    
    // Step 2: compute next rotations
    for k := 1; k < n; k++ {
        F = F + sum - int64(n)*int64(nums[n-k])
        
        if F > result {
            result = F
        }
    }
    
    return int(result)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same across all languages, so I’ll explain it once clearly.

First, I calculate the total sum of all elements. This is needed because every rotation adds this sum to the function value.

Then I calculate F(0). This is done by multiplying each element with its index and adding everything together.

Now comes the main part.

Instead of rotating the array, I simulate the effect of rotation.

When the array rotates:

* Every element moves one position forward
* The last element goes to index 0

So:

* All elements gain +1 in their index contribution
* The last element loses its contribution because it moves to index 0

That’s why I subtract n times the last element.

So each step:

* Add total sum
* Subtract n multiplied by the element that moved to the front

I repeat this for every rotation and keep updating the maximum value.

This avoids rebuilding the array and keeps the solution efficient.

## Examples

### Example 1

Input:
nums = [4, 3, 2, 6]

Output:
26

Explanation:
F(0) = 0*4 + 1*3 + 2*2 + 3*6 = 25
F(1) = 16
F(2) = 23
F(3) = 26

Maximum is 26

---

### Example 2

Input:
nums = [100]

Output:
0

Explanation:
Only one element, so result is always 0

---

### Example 3

Input:
nums = [1, 2, 3, 4, 5]

Output:
40

Explanation:
Best rotation gives the highest weighted sum

## How to Use / Run Locally

### C++

1. Save the file as `solution.cpp`
2. Compile using:
   g++ solution.cpp -o solution
3. Run:
   ./solution

### Java

1. Save as `Solution.java`
2. Compile:
   javac Solution.java
3. Run:
   java Solution

### JavaScript (Node.js)

1. Save as `solution.js`
2. Run:
   node solution.js

### Python3

1. Save as `solution.py`
2. Run:
   python solution.py

### Go

1. Save as `solution.go`
2. Run:
   go run solution.go

## Notes & Optimizations

* The brute force approach takes O(n²), which will fail for large inputs
* The optimized solution uses a mathematical trick to reduce time complexity to O(n)
* Works with negative numbers as well
* Make sure to use long or 64-bit integers to avoid overflow during calculations
* No need to actually rotate the array, which saves both time and space

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
