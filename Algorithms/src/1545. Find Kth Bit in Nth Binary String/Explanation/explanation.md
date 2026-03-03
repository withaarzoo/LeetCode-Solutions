# Problem Title

1545. Find Kth Bit in Nth Binary String

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

We are given two integers `n` and `k`.

A binary string `S_n` is defined recursively:

S1 = "0"
S_i = S_(i-1) + "1" + reverse(invert(S_(i-1))) for i > 1

Where:

* `+` means concatenation
* `reverse(x)` reverses the string
* `invert(x)` flips every bit (0 becomes 1 and 1 becomes 0)

Our task is to return the k-th bit (1-indexed) in `S_n`.

---

## Constraints

* 1 <= n <= 20
* 1 <= k <= 2^n - 1

---

## Intuition

When I first saw this problem, I noticed that building the entire string is not practical. The size of the string grows exponentially.

Length of S_n is:

2^n - 1

For n = 20, the string length is more than one million.

So instead of building the full string, I observed the structure carefully:

S_n = Left + Middle + Right

Where:

* Left = S_(n-1)
* Middle = "1"
* Right = reverse(invert(S_(n-1)))

The middle element is always '1'.

This structure is symmetric.

So instead of generating the full string, I can determine where k lies and reduce the problem recursively.

---

## Approach

1. Compute total length = 2^n - 1.
2. Compute middle index = (length + 1) / 2.
3. If k == middle → return '1'.
4. If k < middle → answer is same as findKthBit(n-1, k).
5. If k > middle → it belongs to the right part.

   * Mirror index = length - k + 1.
   * Recursively compute findKthBit(n-1, mirror index).
   * Invert the result.
6. Base case: if n == 1 → return '0'.

This avoids constructing the full string.

---

## Data Structures Used

No extra data structures are required.

Only recursion stack is used.

---

## Operations & Behavior Summary

* Recursive divide and conquer
* Constant time calculations per level
* Depth of recursion at most n (<= 20)

---

## Complexity

Time Complexity: O(n)

At each step, n decreases by 1. Maximum recursion depth is n.

Space Complexity: O(n)

Only recursion stack is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    char findKthBit(int n, int k) {
        if (n == 1) return '0';

        int length = (1 << n) - 1;
        int mid = (length + 1) / 2;

        if (k == mid) return '1';
        else if (k < mid) return findKthBit(n - 1, k);
        else {
            char bit = findKthBit(n - 1, length - k + 1);
            return bit == '0' ? '1' : '0';
        }
    }
};
```

### Java

```java
class Solution {
    public char findKthBit(int n, int k) {
        if (n == 1) return '0';

        int length = (1 << n) - 1;
        int mid = (length + 1) / 2;

        if (k == mid) return '1';
        else if (k < mid) return findKthBit(n - 1, k);
        else {
            char bit = findKthBit(n - 1, length - k + 1);
            return bit == '0' ? '1' : '0';
        }
    }
}
```

### JavaScript

```javascript
var findKthBit = function(n, k) {
    if (n === 1) return '0';

    const length = (1 << n) - 1;
    const mid = Math.floor((length + 1) / 2);

    if (k === mid) return '1';
    else if (k < mid) return findKthBit(n - 1, k);
    else {
        const bit = findKthBit(n - 1, length - k + 1);
        return bit === '0' ? '1' : '0';
    }
};
```

### Python3

```python
class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        if n == 1:
            return "0"

        length = (1 << n) - 1
        mid = (length + 1) // 2

        if k == mid:
            return "1"
        elif k < mid:
            return self.findKthBit(n - 1, k)
        else:
            bit = self.findKthBit(n - 1, length - k + 1)
            return "1" if bit == "0" else "0"
```

### Go

```go
func findKthBit(n int, k int) byte {
    if n == 1 {
        return '0'
    }

    length := (1 << n) - 1
    mid := (length + 1) / 2

    if k == mid {
        return '1'
    } else if k < mid {
        return findKthBit(n-1, k)
    } else {
        bit := findKthBit(n-1, length-k+1)
        if bit == '0' {
            return '1'
        }
        return '0'
    }
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. If n == 1, return '0'. This is the base string.
2. Compute total length = (1 << n) - 1.
3. Compute mid = (length + 1) / 2.
4. If k equals mid, return '1'. Middle element is always '1'.
5. If k is less than mid, recursively solve for n-1.
6. If k is greater than mid:

   * Compute mirror index = length - k + 1.
   * Recursively compute value at mirror index.
   * Invert the result.
7. Continue reducing n until base case.

This works because right half is reversed and inverted version of left half.

---

## Examples

Example 1:
Input: n = 3, k = 1
Output: "0"

Example 2:
Input: n = 4, k = 11
Output: "1"

---

## How to use / Run locally

C++:

* Compile using g++
* Run the executable

Java:

* Compile using javac
* Run using java

Python:

* Run using python3 file.py

JavaScript:

* Run using node file.js

Go:

* Run using go run file.go

---

## Notes & Optimizations

* Do not build the full string.
* Use recursion to reduce problem size.
* Maximum recursion depth is 20.
* Very efficient and safe within constraints.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
