# 11. Container With Most Water

---

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

Given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`ᵗʰ line are `(i, 0)` and `(i, height[i])`.
Find two lines that together with the x-axis form a container, such that the container contains the most water. Return the **maximum** amount of water a container can store.

Important: the container cannot be slanted — the area is determined by the shorter of the two chosen lines and the horizontal distance between them.

---

## Constraints

* `n == height.length`
* `2 <= n <= 10^5`
* `0 <= height[i] <= 10^4`

---

## Intuition

I thought about how area is calculated for any pair of lines: it is `min(h[i], h[j]) * (j - i)`. If I tried all pairs, that would be `O(n^2)` and far too slow for `n = 10^5`.
So I started with the widest possible container (the two endpoints). Then I realized: if I move the taller side inward, the height limit (which is the shorter side) doesn't increase, and width decreases — that can't help. Only moving the shorter side can potentially increase the limiting height to get a bigger area. This leads to the two-pointer greedy approach.

---

## Approach

1. Set `left = 0` and `right = n - 1` (start with the outermost lines).
2. Compute area = `min(height[left], height[right]) * (right - left)`. Keep the maximum seen so far.
3. Move the pointer that points to the shorter line inward — this may increase the limiting height and possibly increase area.
4. Repeat steps 2–3 until `left` meets `right`.
5. Return `maxArea`.

This checks only `O(n)` candidate pairs and is optimal.

---

## Data Structures Used

* Two integer pointers/indexes: `left` and `right`.
* A few integer variables to track width, current area, and maximum area.

No extra arrays or collections are used.

---

## Operations & Behavior Summary

* **Compute width**: `right - left`.
* **Compute limiting height**: `min(height[left], height[right])`.
* **Compute area**: `width * limiting height`.
* **Update best**: replace `maxArea` if `area` is higher.
* **Greedy pointer move**: move the pointer at the shorter height inward.

---

## Complexity

* **Time Complexity:** `O(n)` — we make at most one pass across the array with two pointers. `n` is `height.length`.
* **Space Complexity:** `O(1)` — only constant extra memory is used for pointers and counters.

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0;                      // left pointer
        int right = (int)height.size() - 1;// right pointer
        int maxArea = 0;                   // best area found

        while (left < right) {
            int width = right - left;                      // horizontal distance
            int h = min(height[left], height[right]);      // limiting height
            int area = h * width;                          // current area
            maxArea = max(maxArea, area);                  // update best

            // Move the pointer at the smaller height inward.
            if (height[left] < height[right]) ++left;
            else --right;
        }
        return maxArea;
    }
};

// Example main to run locally
int main() {
    vector<int> height = {1,8,6,2,5,4,8,3,7};
    Solution sol;
    cout << sol.maxArea(height) << endl; // expected: 49
    return 0;
}
```

---

### Java

```java
// Save as Main.java to run locally
import java.util.*;

class Solution {
    public int maxArea(int[] height) {
        int left = 0;
        int right = height.length - 1;
        int maxArea = 0;

        while (left < right) {
            int width = right - left;
            int h = Math.min(height[left], height[right]);
            int area = h * width;
            if (area > maxArea) maxArea = area;

            // move the shorter side inward
            if (height[left] < height[right]) left++;
            else right--;
        }
        return maxArea;
    }
}

public class Main {
    public static void main(String[] args) {
        int[] height = {1,8,6,2,5,4,8,3,7};
        Solution sol = new Solution();
        System.out.println(sol.maxArea(height)); // expected: 49
    }
}
```

---

### JavaScript

```javascript
// Run with: node maxArea.js
/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let left = 0;
    let right = height.length - 1;
    let maxArea = 0;

    while (left < right) {
        const width = right - left;
        const h = Math.min(height[left], height[right]);
        const area = h * width;
        if (area > maxArea) maxArea = area;

        if (height[left] < height[right]) left++;
        else right--;
    }
    return maxArea;
};

// Example
const height = [1,8,6,2,5,4,8,3,7];
console.log(maxArea(height)); // expected: 49
```

---

### Python3

```python
# Run with: python3 max_area.py
from typing import List

class Solution:
    def maxArea(self, height: List[int]) -> int:
        left = 0
        right = len(height) - 1
        max_area = 0

        while left < right:
            width = right - left
            h = min(height[left], height[right])
            area = h * width
            if area > max_area:
                max_area = area

            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

        return max_area

# Example usage
if __name__ == "__main__":
    height = [1,8,6,2,5,4,8,3,7]
    sol = Solution()
    print(sol.maxArea(height))  # expected: 49
```

---

### Go

```go
// Save as main.go and run: go run main.go
package main

import "fmt"

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0

    for left < right {
        width := right - left
        h := height[left]
        if height[right] < h {
            h = height[right]
        }
        area := h * width
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return maxArea
}

func main() {
    height := []int{1,8,6,2,5,4,8,3,7}
    fmt.Println(maxArea(height)) // expected: 49
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the critical parts step-by-step with the core loop (same logic across all languages).

**Core idea in code (pseudocode):**

```text
left = 0
right = n - 1
maxArea = 0

while left < right:
    width = right - left
    h = min(height[left], height[right])
    area = h * width
    maxArea = max(maxArea, area)

    if height[left] < height[right]:
        left += 1
    else:
        right -= 1

return maxArea
```

**Line-by-line explanation:**

* `left = 0` and `right = n - 1`
  I initialize two pointers at both ends of the array so I start with the widest container.

* `while left < right:`
  I loop as long as there are two different lines to form a container. When pointers meet, every pair has been considered by this greedy approach.

* `width = right - left`
  This is the horizontal distance between the two vertical lines — an essential multiplier for area.

* `h = min(height[left], height[right])`
  The container's height is limited by the shorter line. Water would spill over the smaller side.

* `area = h * width`
  Compute the water amount for this pair.

* `maxArea = max(maxArea, area)`
  Keep the maximum area encountered so far.

* `if height[left] < height[right]: left += 1 else: right -= 1`
  I move the pointer at the **smaller** height inward because moving the taller one cannot lead to a larger area — the limiting height remains the smaller line and the width shrinks.

**Why this is safe (intuition/proof sketch):**

* Suppose `height[left] < height[right]`. Any area formed by pairing `left` with some `k` where `left < k <= right` will have width smaller than `(right - left)`. Its height is at most `height[left]` unless we find a taller line > `height[left]`. Only by moving `left` inward can we possibly find such a taller line. Moving `right` inward cannot increase the limiting height above `height[left]`, so it cannot produce a larger area than the current pair. So skipping `right--` is safe.

**Edge cases handled:**

* If there are zeros in `height`, the algorithm naturally handles them: `h` becomes zero and area is zero; pointers move and continue.
* Minimum `n = 2` works because loop runs until `left < right` and checks that single pair.

---

## Examples

1. **Example 1**

   * Input: `height = [1,8,6,2,5,4,8,3,7]`
   * Output: `49`
   * Explanation: The best container is between index `1` (height 8) and index `8` (height 7). Area = `min(8,7) * (8 - 1) = 7 * 7 = 49`.

2. **Example 2**

   * Input: `height = [1,1]`
   * Output: `1`
   * Explanation: Only one pair to choose: `min(1,1) * (1 - 0) = 1`.

---

## How to use / Run locally

### C++

1. Save file as `max_area.cpp`.
2. Compile & run:

   ```bash
   g++ -std=c++17 -O2 max_area.cpp -o max_area
   ./max_area
   ```

### Java

1. Save file as `Main.java`.
2. Compile & run:

   ```bash
   javac Main.java
   java Main
   ```

### JavaScript (Node.js)

1. Save file as `maxArea.js`.
2. Run:

   ```bash
   node maxArea.js
   ```

### Python3

1. Save file as `max_area.py`.
2. Run:

   ```bash
   python3 max_area.py
   ```

### Go

1. Save file as `main.go`.
2. Run:

   ```bash
   go run main.go
   ```

Each example program prints the expected result for the provided sample `height` array (`49` for the main sample).

---

## Notes & Optimizations

* This greedy two-pointer solution runs in linear time `O(n)` and is the optimal solution for this problem.
* Avoid trying all pairs (two nested loops) — that is `O(n^2)` and will time out for large `n`.
* The algorithm uses constant extra space (`O(1)`).
* If multiple languages are used, the internal logic remains identical — only syntax changes.

---

## Author

[Md. Aarzoo Islam](https://bento.me/withaarzoo)
