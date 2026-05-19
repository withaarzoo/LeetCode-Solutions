# 2540. Minimum Common Value

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

This LeetCode problem asks us to find the minimum common value between two sorted integer arrays.

We are given:

* `nums1`
* `nums2`

Both arrays are already sorted in non-decreasing order.

The goal is simple:

* Return the smallest integer that appears in both arrays
* If no common value exists, return `-1`

This is a classic Two Pointers problem in Data Structures and Algorithms because the sorted order allows us to compare elements efficiently without checking every pair.

This problem is popular in coding interviews because it tests:

* Array traversal
* Sorting observations
* Two pointer technique
* Time complexity optimization

---

## Constraints

| Constraint                                | Value                   |
| ----------------------------------------- | ----------------------- |
| `1 <= nums1.length, nums2.length <= 10^5` | Large input size        |
| `1 <= nums1[i], nums2[j] <= 10^9`         | Large integer values    |
| Arrays are sorted                         | In non-decreasing order |

---

## Intuition

The first thing I noticed was that both arrays are already sorted.

That immediately made brute force feel unnecessary.

If I compare every value from `nums1` with every value from `nums2`, the solution becomes too slow for large inputs.

Since the arrays are sorted, I realized I can move through them in a single pass using two pointers.

The idea is simple:

* If one value is smaller, move that pointer forward
* If both values become equal, that is the smallest common value

Because the arrays are sorted from left to right, the first match I find is automatically the answer.

---

## Approach

I used the Two Pointers algorithm.

Step-by-step process:

1. Create two pointers:

   * One for `nums1`
   * One for `nums2`

2. Start both pointers at index `0`

3. Compare current values:

   * If values are equal:

     * Return that value immediately
   * If value in `nums1` is smaller:

     * Move pointer of `nums1`
   * Otherwise:

     * Move pointer of `nums2`

4. Continue until one array finishes

5. If no match is found:

   * Return `-1`

This approach works efficiently because every element is visited at most once.

---

## Data Structures Used

| Data Structure   | Why It Was Used                     |
| ---------------- | ----------------------------------- |
| Arrays           | Input storage                       |
| Integer Pointers | To traverse both arrays efficiently |

No extra hash map, set, or additional array is needed.

---

## Operations & Behavior Summary

The algorithm performs these operations:

1. Start scanning both sorted arrays together
2. Compare current elements
3. Skip smaller values because they cannot match later
4. Stop immediately when a common value is found
5. Return `-1` if traversal finishes without a match

This creates an optimized linear-time solution.

---

## Complexity

| Type             | Complexity | Explanation                                                                   |
| ---------------- | ---------- | ----------------------------------------------------------------------------- |
| Time Complexity  | `O(n + m)` | `n` is size of `nums1`, `m` is size of `nums2`. Each pointer moves only once. |
| Space Complexity | `O(1)`     | No extra data structure is used apart from pointers.                          |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int getCommon(vector<int>& nums1, vector<int>& nums2) {
        
        // Pointer for nums1
        int i = 0;

        // Pointer for nums2
        int j = 0;

        // Traverse both arrays together
        while (i < nums1.size() && j < nums2.size()) {

            // If both values are equal,
            // this is the minimum common value
            if (nums1[i] == nums2[j]) {
                return nums1[i];
            }

            // Move the pointer with smaller value
            // because it cannot match later
            if (nums1[i] < nums2[j]) {
                i++;
            } else {
                j++;
            }
        }

        // No common value found
        return -1;
    }
};
```

### Java

```java
class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        
        // Pointer for nums1
        int i = 0;

        // Pointer for nums2
        int j = 0;

        // Traverse both arrays
        while (i < nums1.length && j < nums2.length) {

            // Common value found
            if (nums1[i] == nums2[j]) {
                return nums1[i];
            }

            // Move the smaller value forward
            if (nums1[i] < nums2[j]) {
                i++;
            } else {
                j++;
            }
        }

        // No common element exists
        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var getCommon = function(nums1, nums2) {
    
    // Pointer for nums1
    let i = 0;

    // Pointer for nums2
    let j = 0;

    // Traverse both arrays
    while (i < nums1.length && j < nums2.length) {

        // If values are equal,
        // return the common value
        if (nums1[i] === nums2[j]) {
            return nums1[i];
        }

        // Move pointer having smaller value
        if (nums1[i] < nums2[j]) {
            i++;
        } else {
            j++;
        }
    }

    // No common value found
    return -1;
};
```

### Python3

```python
class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        
        # Pointer for nums1
        i = 0

        # Pointer for nums2
        j = 0

        # Traverse both arrays together
        while i < len(nums1) and j < len(nums2):

            # If both values are same,
            # return the common value
            if nums1[i] == nums2[j]:
                return nums1[i]

            # Move the pointer with smaller value
            if nums1[i] < nums2[j]:
                i += 1
            else:
                j += 1

        # No common value exists
        return -1
```

### Go

```go
func getCommon(nums1 []int, nums2 []int) int {
    
    // Pointer for nums1
    i := 0

    // Pointer for nums2
    j := 0

    // Traverse both arrays
    for i < len(nums1) && j < len(nums2) {

        // If both values are equal,
        // return the answer
        if nums1[i] == nums2[j] {
            return nums1[i]
        }

        // Move the pointer with smaller value
        if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }

    // No common value found
    return -1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

Only syntax changes.

### Step 1 — Create Two Pointers

I create:

* One pointer for `nums1`
* One pointer for `nums2`

These pointers help me track the current position in both arrays.

Both start from index `0`.

---

### Step 2 — Traverse Both Arrays Together

I continue looping while:

* Pointer of `nums1` is inside the array
* Pointer of `nums2` is inside the array

The moment one pointer reaches the end, comparison becomes impossible.

---

### Step 3 — Compare Current Values

At every step, I compare:

* `nums1[i]`
* `nums2[j]`

Three situations are possible.

#### Case 1: Both Values Are Equal

This means I found a common value.

Because arrays are sorted, this is automatically the minimum common value.

So I immediately return it.

---

#### Case 2: First Value Is Smaller

Example:

```text
nums1[i] = 2
nums2[j] = 5
```

Since arrays are sorted, `2` can never match `5` later.

So I move pointer `i` forward.

---

#### Case 3: Second Value Is Smaller

Example:

```text
nums1[i] = 10
nums2[j] = 7
```

Now `7` is smaller.

It also cannot become equal later unless I move forward.

So I move pointer `j`.

---

### Step 4 — Handle No Common Value

If the loop finishes and no equality is found, I return `-1`.

That means the arrays do not share any value.

---

## Examples

### Example 1

Input:

```text
nums1 = [1,2,3]
nums2 = [2,4]
```

Output:

```text
2
```

Trace:

* Compare `1` and `2`
* Move first pointer
* Compare `2` and `2`
* Match found

---

### Example 2

Input:

```text
nums1 = [1,2,3,6]
nums2 = [2,3,4,5]
```

Output:

```text
2
```

Trace:

* `1 < 2`, move first pointer
* `2 == 2`, return `2`

---

### Example 3

Input:

```text
nums1 = [1,3,5]
nums2 = [2,4,6]
```

Output:

```text
-1
```

Trace:

* No values ever match
* Traversal ends
* Return `-1`

---

## How to Use / Run Locally

### C++

Compile and run:

```bash
g++ main.cpp -o main
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
node index.js
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

* The Two Pointers approach is the best solution for sorted arrays
* A brute force solution would be too slow for large inputs
* A hash set approach also works, but it uses extra memory
* Since arrays are already sorted, using pointers is more efficient
* The algorithm stops early as soon as a match is found
* This problem is commonly asked in coding interviews for array and pointer concepts

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
