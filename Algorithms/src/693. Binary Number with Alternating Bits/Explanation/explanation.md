# 693. Binary Number with Alternating Bits

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

Given a positive integer `n`, return `true` if and only if it has alternating bits.

A number has alternating bits if no two adjacent bits in its binary representation are the same.

Example:

* 5 → binary 101 → alternating → true
* 7 → binary 111 → not alternating → false

---

## Constraints

* 1 <= n <= 2^31 - 1

---

## Intuition

When I first looked at this problem, I thought very simply.

If the number has alternating bits, then every bit must be different from the previous bit.

So instead of converting the number to a string or using complex tricks, I decided to:

* Check the last bit
* Compare it with the next bit
* Continue this until the number becomes zero

If I ever find two adjacent bits that are the same, I immediately return false.

If I finish checking all bits without finding any equal adjacent bits, then the number is alternating.

---

## Approach

Step 1: Get the last bit using `n & 1`

Step 2: Store it in a variable called `prev`

Step 3: Right shift the number (`n >>= 1`)

Step 4: While `n > 0`

* Extract current bit
* Compare with previous bit
* If same → return false
* Otherwise update previous bit
* Shift again

Step 5: If loop finishes → return true

This ensures we compare every adjacent pair exactly once.

---

## Data Structures Used

No extra data structures are used.

Only integer variables are used for bit comparison.

---

## Operations & Behavior Summary

* Bitwise AND (`&`) to extract last bit
* Right shift (`>>`) to move to next bit
* Simple comparison between adjacent bits
* Early return if invalid pattern found

---

## Complexity

**Time Complexity:** O(k)

Where `k` is the number of bits in `n`.
Since number of bits is approximately log(n), this is very efficient.

**Space Complexity:** O(1)

We only use a few integer variables.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool hasAlternatingBits(int n) {
        int prev = n & 1;
        n >>= 1;

        while (n > 0) {
            int curr = n & 1;

            if (curr == prev) {
                return false;
            }

            prev = curr;
            n >>= 1;
        }

        return true;
    }
};
```

---

### Java

```java
class Solution {
    public boolean hasAlternatingBits(int n) {
        int prev = n & 1;
        n >>= 1;

        while (n > 0) {
            int curr = n & 1;

            if (curr == prev) {
                return false;
            }

            prev = curr;
            n >>= 1;
        }

        return true;
    }
}
```

---

### JavaScript

```javascript
var hasAlternatingBits = function(n) {
    let prev = n & 1;
    n = n >> 1;

    while (n > 0) {
        let curr = n & 1;

        if (curr === prev) {
            return false;
        }

        prev = curr;
        n = n >> 1;
    }

    return true;
};
```

---

### Python3

```python
class Solution:
    def hasAlternatingBits(self, n: int) -> bool:
        prev = n & 1
        n >>= 1

        while n > 0:
            curr = n & 1

            if curr == prev:
                return False

            prev = curr
            n >>= 1

        return True
```

---

### Go

```go
func hasAlternatingBits(n int) bool {
    prev := n & 1
    n >>= 1

    for n > 0 {
        curr := n & 1

        if curr == prev {
            return false
        }

        prev = curr
        n >>= 1
    }

    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Let’s understand using example n = 5

Binary of 5 is 101

Step 1
prev = 5 & 1 → 1
n becomes 10 (binary 10)

Step 2
curr = 10 & 1 → 0
curr != prev → continue
prev = 0
n becomes 1

Step 3
curr = 1 & 1 → 1
curr != prev → continue
prev = 1
n becomes 0

Loop ends → return true

Now example n = 7
Binary 111

prev = 1
n becomes 11

curr = 1
curr == prev → return false immediately

---

## Examples

Input: 5
Output: true
Explanation: 101 is alternating

Input: 7
Output: false
Explanation: 111 is not alternating

Input: 11
Output: false
Explanation: 1011 has two 1s adjacent at the end

---

## How to use / Run locally

C++

* Save file as solution.cpp
* Compile using g++ solution.cpp
* Run using ./a.out

Java

* Save file as Solution.java
* Compile using javac Solution.java
* Run using java Solution

Python

* Save file as solution.py
* Run using python solution.py

Go

* Save file as solution.go
* Run using go run solution.go

---

## Notes & Optimizations

* No need to convert number into binary string
* No need for extra memory
* Early return makes it efficient
* Clean and interview friendly approach

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
