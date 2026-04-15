# 2515. Shortest Distance to Target String in a Circular Array

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

You are given a circular string array `words`, a target string `target`, and a starting index `startIndex`.

Since the array is circular:

* Moving right from the last index takes us to the first index.
* Moving left from the first index takes us to the last index.

I need to return the minimum number of steps required to reach the target string from `startIndex`.

If the target string does not exist in the array, I return `-1`.

## Constraints

* `1 <= words.length <= 100`
* `1 <= words[i].length <= 100`
* `words[i]` and `target` contain only lowercase English letters.
* `0 <= startIndex < words.length`

## Intuition

I thought about checking every index in the array.

If the current word matches the target, then I can calculate how far it is from `startIndex`.

Because the array is circular, I have two possible ways to reach that index:

1. Move directly
2. Move around the other side of the circle

So for every matching index:

* Direct distance = `abs(i - startIndex)`
* Circular distance = `n - abs(i - startIndex)`

I take the minimum of these two values.

Then I keep updating the overall minimum answer.

## Approach

1. Store the size of the array in `n`.
2. Initialize the answer with a very large value.
3. Traverse the entire array.
4. If `words[i] == target`:

   * Calculate the normal distance.
   * Calculate the circular distance.
   * Take the minimum of both.
   * Update the final answer.
5. If the target was never found, return `-1`.
6. Otherwise, return the minimum answer.

## Data Structures Used

* Array / List
* Integer variables

No extra data structure is needed.

## Operations & Behavior Summary

| Operation          | Description                                |
| ------------------ | ------------------------------------------ |
| Traverse Array     | Check every index once                     |
| Compare String     | See if current word matches target         |
| Calculate Distance | Find direct distance and circular distance |
| Update Answer      | Keep the smallest valid distance           |
| Return Result      | Return minimum distance or `-1`            |

## Complexity

* Time Complexity: `O(n)`

  * I scan the array once.
  * Here, `n` is the number of elements in `words`.

* Space Complexity: `O(1)`

  * I only use a few extra variables.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int closestTarget(vector<string>& words, string target, int startIndex) {
        int n = words.size();
        int ans = INT_MAX;

        // Traverse the array
        for (int i = 0; i < n; i++) {
            // If target is found
            if (words[i] == target) {
                int diff = abs(i - startIndex);

                // Distance if we go in the circular direction
                int circularDist = n - diff;

                // Update minimum answer
                ans = min(ans, min(diff, circularDist));
            }
        }

        // If target was never found
        return ans == INT_MAX ? -1 : ans;
    }
};
```

### Java

```java
class Solution {
    public int closestTarget(String[] words, String target, int startIndex) {
        int n = words.length;
        int ans = Integer.MAX_VALUE;

        // Traverse the array
        for (int i = 0; i < n; i++) {
            // If current word matches target
            if (words[i].equals(target)) {
                int diff = Math.abs(i - startIndex);

                // Circular distance
                int circularDist = n - diff;

                // Update answer
                ans = Math.min(ans, Math.min(diff, circularDist));
            }
        }

        // If target does not exist
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string[]} words
 * @param {string} target
 * @param {number} startIndex
 * @return {number}
 */
var closestTarget = function(words, target, startIndex) {
    const n = words.length;
    let ans = Infinity;

    // Traverse the array
    for (let i = 0; i < n; i++) {
        // If target is found
        if (words[i] === target) {
            const diff = Math.abs(i - startIndex);

            // Circular distance
            const circularDist = n - diff;

            // Update answer
            ans = Math.min(ans, Math.min(diff, circularDist));
        }
    }

    // If target was not found
    return ans === Infinity ? -1 : ans;
};
```

### Python3

```python
class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        n = len(words)
        ans = float('inf')

        # Traverse the array
        for i in range(n):
            # If current word matches target
            if words[i] == target:
                diff = abs(i - startIndex)

                # Circular distance
                circular_dist = n - diff

                # Update minimum answer
                ans = min(ans, diff, circular_dist)

        # If target does not exist
        return -1 if ans == float('inf') else ans
```

### Go

```go
func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    ans := int(^uint(0) >> 1)

    // Traverse the array
    for i := 0; i < n; i++ {
        // If target is found
        if words[i] == target {
            diff := i - startIndex
            if diff < 0 {
                diff = -diff
            }

            // Circular distance
            circularDist := n - diff

            current := diff
            if circularDist < current {
                current = circularDist
            }

            // Update answer
            if current < ans {
                ans = current
            }
        }
    }

    // If target does not exist
    if ans == int(^uint(0)>>1) {
        return -1
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Store the size of the array

```cpp
int n = words.size();
```

I store the array size because I need it later while calculating circular distance.

---

### Step 2: Initialize the answer

```cpp
int ans = INT_MAX;
```

I use a very large number initially.

Whenever I find a smaller valid distance, I update it.

---

### Step 3: Traverse every index

```cpp
for (int i = 0; i < n; i++)
```

I check every index because the target may appear multiple times.

---

### Step 4: Check whether current word matches target

```cpp
if (words[i] == target)
```

Only matching indices are useful.

If the word is not equal to target, I skip it.

---

### Step 5: Calculate direct distance

```cpp
int diff = abs(i - startIndex);
```

This gives the direct distance between the current index and the starting index.

Example:

```cpp
startIndex = 1
i = 4
diff = abs(4 - 1) = 3
```

---

### Step 6: Calculate circular distance

```cpp
int circularDist = n - diff;
```

Since the array is circular, I can also move in the opposite direction.

Example:

```cpp
n = 5
diff = 3
circularDist = 5 - 3 = 2
```

So moving 2 steps in the opposite direction is better than moving 3 steps directly.

---

### Step 7: Update minimum answer

```cpp
ans = min(ans, min(diff, circularDist));
```

First, I take the better distance for the current target position.

Then I compare it with the global answer.

---

### Step 8: Return the final result

```cpp
return ans == INT_MAX ? -1 : ans;
```

If the answer was never updated, it means target does not exist.

So I return `-1`.

Otherwise, I return the minimum distance.

## Examples

### Example 1

```text
Input:
words = ["hello", "i", "am", "leetcode", "hello"]
target = "hello"
startIndex = 1

Output:
1
```

Explanation:

* From index 1 to index 0 = 1 step
* From index 1 to index 4 = 2 steps
* Minimum answer = 1

### Example 2

```text
Input:
words = ["a", "b", "leetcode"]
target = "leetcode"
startIndex = 0

Output:
1
```

Explanation:

* Direct distance = 2
* Circular distance = 1
* Minimum answer = 1

### Example 3

```text
Input:
words = ["i", "eat", "leetcode"]
target = "ate"
startIndex = 0

Output:
-1
```

Explanation:

* The target does not exist in the array.

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

* I only scan the array once.
* I do not need any extra data structure.
* The circular distance formula makes the solution simple.
* This is the most optimal solution because every element must be checked at least once.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
