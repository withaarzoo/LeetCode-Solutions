# 868. Binary Gap

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given a positive integer `n`, return the longest distance between two adjacent `1`s in the binary representation of `n`.

Two `1`s are considered adjacent if there are only `0`s between them (possibly zero `0`s).

The distance between two `1`s is the difference between their bit positions.

If there are no two adjacent `1`s, return `0`.

---

## Constraints

* 1 <= n <= 10^9

---

## Intuition

When I first looked at this problem, I thought in a simple way.

I just need to convert the number into binary (mentally or using bit operations). Then I need to check the positions of `1`s.

Every time I see a `1`, I compare its position with the previous `1` I saw. I calculate the distance and keep track of the maximum.

So instead of converting the number into a string, I directly use bit manipulation. That makes the solution faster and cleaner.

---

## Approach

Here is exactly how I solved it:

1. I create a variable `lastPosition` and set it to `-1`. This stores the index of the previous `1`.
2. I create `maxDistance` to store the final answer.
3. I start checking bits from right to left.
4. While `n > 0`:

   * Check if the last bit is `1` using `(n & 1)`.
   * If yes:

     * If I have seen a `1` before, calculate distance.
     * Update `maxDistance`.
     * Update `lastPosition`.
   * Shift the number right using `n >>= 1`.
5. Return `maxDistance`.

This way I scan each bit exactly once.

---

## Data Structures Used

I did not use any extra data structure.

Only simple variables:

* Integer for last position
* Integer for max distance
* Integer for current bit index

So space usage is constant.

---

## Operations & Behavior Summary

* Bitwise AND (`n & 1`) → Check if current bit is 1
* Right Shift (`n >>= 1`) → Move to next bit
* Subtraction → Calculate distance between bit positions
* Max comparison → Track largest gap

---

## Complexity

**Time Complexity:** O(log n)

* Because we process each bit of `n` once.
* A number `n` has at most log2(n) bits.

**Space Complexity:** O(1)

* Only a few integer variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int binaryGap(int n) {
        int lastPosition = -1;
        int maxDistance = 0;
        int currentPosition = 0;
        
        while (n > 0) {
            if (n & 1) {
                if (lastPosition != -1) {
                    maxDistance = max(maxDistance, currentPosition - lastPosition);
                }
                lastPosition = currentPosition;
            }
            n >>= 1;
            currentPosition++;
        }
        return maxDistance;
    }
};
```

### Java

```java
class Solution {
    public int binaryGap(int n) {
        int lastPosition = -1;
        int maxDistance = 0;
        int currentPosition = 0;
        
        while (n > 0) {
            if ((n & 1) == 1) {
                if (lastPosition != -1) {
                    maxDistance = Math.max(maxDistance, currentPosition - lastPosition);
                }
                lastPosition = currentPosition;
            }
            n >>= 1;
            currentPosition++;
        }
        return maxDistance;
    }
}
```

### JavaScript

```javascript
var binaryGap = function(n) {
    let lastPosition = -1;
    let maxDistance = 0;
    let currentPosition = 0;
    
    while (n > 0) {
        if ((n & 1) === 1) {
            if (lastPosition !== -1) {
                maxDistance = Math.max(maxDistance, currentPosition - lastPosition);
            }
            lastPosition = currentPosition;
        }
        n = n >> 1;
        currentPosition++;
    }
    return maxDistance;
};
```

### Python3

```python
class Solution:
    def binaryGap(self, n: int) -> int:
        last_position = -1
        max_distance = 0
        current_position = 0
        
        while n > 0:
            if n & 1:
                if last_position != -1:
                    max_distance = max(max_distance, current_position - last_position)
                last_position = current_position
            n >>= 1
            current_position += 1
        
        return max_distance
```

### Go

```go
func binaryGap(n int) int {
    lastPosition := -1
    maxDistance := 0
    currentPosition := 0
    
    for n > 0 {
        if n&1 == 1 {
            if lastPosition != -1 {
                if currentPosition-lastPosition > maxDistance {
                    maxDistance = currentPosition - lastPosition
                }
            }
            lastPosition = currentPosition
        }
        n >>= 1
        currentPosition++
    }
    return maxDistance
}
```

---

## Step-by-step Detailed Explanation

Let me explain the logic clearly.

1. `lastPosition = -1`

   * This means I have not seen any `1` yet.

2. `maxDistance = 0`

   * This will store the longest gap.

3. `currentPosition = 0`

   * This tracks the current bit index.

4. `while (n > 0)`

   * I keep checking bits until the number becomes zero.

5. `if (n & 1)`

   * This checks if the current bit is 1.

6. If I already saw a `1` before:

   * I calculate distance: `currentPosition - lastPosition`.
   * Update maximum.

7. Update `lastPosition` to current index.

8. Shift right: `n >>= 1`

   * Move to next bit.

9. Increase `currentPosition`.

10. Finally return `maxDistance`.

---

## Examples

### Example 1

Input: n = 22
Binary: 10110
Output: 2

### Example 2

Input: n = 8
Binary: 1000
Output: 0

### Example 3

Input: n = 5
Binary: 101
Output: 2

---

## How to use / Run locally

### C++

* Compile: `g++ file.cpp`
* Run: `./a.out`

### Java

* Compile: `javac Solution.java`
* Run: `java Solution`

### Python

* Run: `python file.py`

### JavaScript

* Run: `node file.js`

### Go

* Run: `go run file.go`

---

## Notes & Optimizations

* No need to convert number to binary string.
* Bit manipulation is faster and memory efficient.
* Only one traversal of bits.
* Works efficiently for large inputs up to 10^9.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
