# 1855. Maximum Distance Between a Pair of Values

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

You are given two non-increasing integer arrays `nums1` and `nums2`.

A pair of indices `(i, j)` is valid if:

* `0 <= i < nums1.length`
* `0 <= j < nums2.length`
* `i <= j`
* `nums1[i] <= nums2[j]`

The distance of the pair is:

```text
j - i
```

My task is to return the maximum distance among all valid pairs.

---

## Constraints

```text
1 <= nums1.length, nums2.length <= 100000
1 <= nums1[i], nums2[j] <= 100000
Both nums1 and nums2 are non-increasing
```

---

## Intuition

At first, I thought about checking every possible pair `(i, j)`.

But that would take too much time because both arrays can contain up to `100000` elements.

A brute-force solution would be `O(n * m)`, which is too slow.

Then I noticed that both arrays are already sorted in non-increasing order.

Because of this, I can use the two-pointer technique.

* If `nums1[i] <= nums2[j]`, then the pair is valid.
* I can update the answer and move `j` forward to try for a bigger distance.
* If the pair is invalid, I move `i` forward.

This allows me to solve the problem in linear time.

---

## Approach

1. Start with two pointers:

   * `i = 0` for `nums1`
   * `j = 0` for `nums2`

2. Traverse both arrays while both pointers are inside the valid range.

3. If `i > j`, move `j` forward because valid pairs require `i <= j`.

4. If `nums1[i] <= nums2[j]`:

   * The pair is valid.
   * Update the answer with `j - i`.
   * Move `j` forward.

5. Otherwise:

   * The pair is invalid.
   * Move `i` forward.

6. Continue until one pointer reaches the end.

---

## Data Structures Used

* Two integer pointers: `i` and `j`
* One integer variable: `ans`

No extra arrays or hash maps are used.

---

## Operations & Behavior Summary

| Condition              | Action                              |
| ---------------------- | ----------------------------------- |
| `i > j`                | Move `j` forward                    |
| `nums1[i] <= nums2[j]` | Valid pair, update answer, move `j` |
| `nums1[i] > nums2[j]`  | Invalid pair, move `i`              |

---

## Complexity

* Time Complexity: `O(n + m)`

  * `n` is the size of `nums1`
  * `m` is the size of `nums2`
  * Each pointer moves only once

* Space Complexity: `O(1)`

  * No extra data structure is used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxDistance(vector<int>& nums1, vector<int>& nums2) {
        int i = 0, j = 0;
        int ans = 0;

        while (i < nums1.size() && j < nums2.size()) {

            if (i > j) {
                j++;
                continue;
            }

            if (nums1[i] <= nums2[j]) {
                ans = max(ans, j - i);
                j++;
            } else {
                i++;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int maxDistance(int[] nums1, int[] nums2) {
        int i = 0, j = 0;
        int ans = 0;

        while (i < nums1.length && j < nums2.length) {

            if (i > j) {
                j++;
                continue;
            }

            if (nums1[i] <= nums2[j]) {
                ans = Math.max(ans, j - i);
                j++;
            } else {
                i++;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
var maxDistance = function(nums1, nums2) {
    let i = 0;
    let j = 0;
    let ans = 0;

    while (i < nums1.length && j < nums2.length) {

        if (i > j) {
            j++;
            continue;
        }

        if (nums1[i] <= nums2[j]) {
            ans = Math.max(ans, j - i);
            j++;
        } else {
            i++;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        i = 0
        j = 0
        ans = 0

        while i < len(nums1) and j < len(nums2):

            if i > j:
                j += 1
                continue

            if nums1[i] <= nums2[j]:
                ans = max(ans, j - i)
                j += 1
            else:
                i += 1

        return ans
```

### Go

```go
func maxDistance(nums1 []int, nums2 []int) int {
    i, j := 0, 0
    ans := 0

    for i < len(nums1) && j < len(nums2) {

        if i > j {
            j++
            continue
        }

        if nums1[i] <= nums2[j] {
            if j-i > ans {
                ans = j - i
            }
            j++
        } else {
            i++
        }
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation

### C++

```cpp
int i = 0, j = 0;
int ans = 0;
```

* `i` points to `nums1`
* `j` points to `nums2`
* `ans` stores the maximum distance found so far

```cpp
while (i < nums1.size() && j < nums2.size())
```

I continue while both pointers are valid.

```cpp
if (i > j) {
    j++;
    continue;
}
```

A valid pair needs `i <= j`.

So if `i` becomes larger than `j`, I move `j`.

```cpp
if (nums1[i] <= nums2[j]) {
    ans = max(ans, j - i);
    j++;
}
```

If the pair is valid:

* Update the answer
* Move `j` to search for a bigger distance

```cpp
else {
    i++;
}
```

If the pair is invalid:

* Move `i` to reduce the value from `nums1`

### Java

The Java solution follows exactly the same logic:

* Use two pointers `i` and `j`
* Check if the pair is valid
* Update the answer if valid
* Move `i` or `j` based on the condition

### JavaScript

The JavaScript solution also uses the same two-pointer approach.

* `i` tracks `nums1`
* `j` tracks `nums2`
* Move only one pointer at a time

### Python3

The Python solution is short and clean.

* `i` and `j` are integers
* `ans` stores the maximum distance
* Loop until one pointer reaches the end

### Go

The Go solution is also identical in logic.

* Two pointers
* Constant extra space
* Linear time complexity

---

## Examples

### Example 1

```text
Input:
nums1 = [55,30,5,4,2]
nums2 = [100,20,10,10,5]

Output:
2
```

Valid pair with maximum distance:

```text
(i, j) = (2, 4)
Distance = 4 - 2 = 2
```

### Example 2

```text
Input:
nums1 = [2,2,2]
nums2 = [10,10,1]

Output:
1
```

### Example 3

```text
Input:
nums1 = [30,29,19,5]
nums2 = [25,25,25,25,25]

Output:
2
```

---

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* Brute force would be too slow because it checks every possible pair.
* Two pointers work because both arrays are sorted in non-increasing order.
* Each pointer moves only once.
* No extra memory is needed.
* This is the most optimal solution for this problem.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
