# 2078. Two Furthest Houses With Different Colors

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

You are given an integer array `colors` where `colors[i]` represents the color of the `i-th` house.

Your task is to find the maximum distance between two houses having different colors.

The distance between house `i` and house `j` is:

```text
abs(i - j)
```

You need to return the maximum possible distance.

## Constraints

```text
2 <= colors.length <= 100
0 <= colors[i] <= 100
At least two houses have different colors.
```

## Intuition

I thought about the maximum possible distance between two houses.

To get the farthest distance, I should always try to use one of the ends of the array.

That means:

* Either I use the first house and search from the end for a different color.
* Or I use the last house and search from the beginning for a different color.

Because the answer must include one of the two ends, I only need these two checks.

## Approach

1. Store the size of the array in `n`.
2. Start from the end of the array.
3. Find the first house whose color is different from `colors[0]`.
4. Its index itself becomes the distance from the first house.
5. Then start from the beginning of the array.
6. Find the first house whose color is different from `colors[n - 1]`.
7. Calculate the distance from the last house.
8. Return the maximum of both distances.

## Data Structures Used

* Array / List
* Integer variables for storing:

  * Size of array
  * Current answer
  * Loop indexes

No extra data structure is needed.

## Operations & Behavior Summary

| Operation                   | Purpose                                                   |
| --------------------------- | --------------------------------------------------------- |
| Traverse from right to left | Find farthest house with different color from first house |
| Traverse from left to right | Find farthest house with different color from last house  |
| Compare distances           | Store the maximum valid answer                            |

## Complexity

* Time Complexity: `O(n)`

  * I scan the array at most two times.
  * Here, `n` is the number of houses.

* Space Complexity: `O(1)`

  * I only use a few extra variables.
  * No additional array, map, or set is used.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxDistance(vector<int>& colors) {
        int n = colors.size();
        int ans = 0;

        // Find farthest house from the first house
        for (int i = n - 1; i >= 0; i--) {
            if (colors[i] != colors[0]) {
                ans = max(ans, i);
                break;
            }
        }

        // Find farthest house from the last house
        for (int i = 0; i < n; i++) {
            if (colors[i] != colors[n - 1]) {
                ans = max(ans, n - 1 - i);
                break;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int maxDistance(int[] colors) {
        int n = colors.length;
        int ans = 0;

        // Find farthest house from the first house
        for (int i = n - 1; i >= 0; i--) {
            if (colors[i] != colors[0]) {
                ans = Math.max(ans, i);
                break;
            }
        }

        // Find farthest house from the last house
        for (int i = 0; i < n; i++) {
            if (colors[i] != colors[n - 1]) {
                ans = Math.max(ans, n - 1 - i);
                break;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} colors
 * @return {number}
 */
var maxDistance = function(colors) {
    const n = colors.length;
    let ans = 0;

    // Find farthest house from the first house
    for (let i = n - 1; i >= 0; i--) {
        if (colors[i] !== colors[0]) {
            ans = Math.max(ans, i);
            break;
        }
    }

    // Find farthest house from the last house
    for (let i = 0; i < n; i++) {
        if (colors[i] !== colors[n - 1]) {
            ans = Math.max(ans, n - 1 - i);
            break;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        n = len(colors)
        ans = 0

        # Find farthest house from the first house
        for i in range(n - 1, -1, -1):
            if colors[i] != colors[0]:
                ans = max(ans, i)
                break

        # Find farthest house from the last house
        for i in range(n):
            if colors[i] != colors[n - 1]:
                ans = max(ans, n - 1 - i)
                break

        return ans
```

### Go

```go
func maxDistance(colors []int) int {
    n := len(colors)
    ans := 0

    // Find farthest house from the first house
    for i := n - 1; i >= 0; i-- {
        if colors[i] != colors[0] {
            if i > ans {
                ans = i
            }
            break
        }
    }

    // Find farthest house from the last house
    for i := 0; i < n; i++ {
        if colors[i] != colors[n-1] {
            distance := n - 1 - i
            if distance > ans {
                ans = distance
            }
            break
        }
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Let:

```text
colors = [1,1,1,6,1,1,1]
```

Here:

* First house color = `1`
* Last house color = `1`

### First Traversal

I start from the end and compare every house with the first house.

```text
Index 6 -> color 1 -> same
Index 5 -> color 1 -> same
Index 4 -> color 1 -> same
Index 3 -> color 6 -> different
```

So distance becomes:

```text
3 - 0 = 3
```

### Second Traversal

Now I compare every house with the last house.

```text
Index 0 -> color 1 -> same
Index 1 -> color 1 -> same
Index 2 -> color 1 -> same
Index 3 -> color 6 -> different
```

So distance becomes:

```text
6 - 3 = 3
```

Final answer:

```text
max(3, 3) = 3
```

### Important Lines

```cpp
if (colors[i] != colors[0])
```

This checks whether the current house has a different color from the first house.

```cpp
if (colors[i] != colors[n - 1])
```

This checks whether the current house has a different color from the last house.

```cpp
ans = max(ans, i);
```

Since the first house index is `0`, the index `i` itself is the distance.

```cpp
ans = max(ans, n - 1 - i);
```

This calculates the distance from the last house.

## Examples

### Example 1

```text
Input: colors = [1,1,1,6,1,1,1]
Output: 3
```

### Example 2

```text
Input: colors = [1,8,3,8,3]
Output: 4
```

### Example 3

```text
Input: colors = [0,1]
Output: 1
```

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
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

## Notes & Optimizations

* I do not need nested loops.
* A brute force solution would check every pair and take `O(n^2)` time.
* This optimized approach only takes `O(n)` time.
* Since I only care about the farthest valid pair, checking from both ends is enough.
* No extra memory is required.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
