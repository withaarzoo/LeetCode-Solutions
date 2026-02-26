# 1404. Number of Steps to Reduce a Number in Binary Representation to One

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

Given a binary string `s` representing a positive integer, return the number of steps required to reduce it to `1`.

Rules:

* If the number is even, divide it by 2.
* If the number is odd, add 1 to it.

The input is guaranteed to always reach 1.

---

## Constraints

* 1 <= s.length <= 500
* s consists only of characters '0' or '1'
* s[0] == '1'

---

## Intuition

When I first looked at this problem, I realized converting the binary string into an integer is not a good idea because the length can go up to 500. That number can be extremely large.

So instead of converting it, I thought about how division by 2 and addition work in binary.

* If a number ends with 0, it is even. Dividing by 2 just removes the last bit.
* If a number ends with 1, it is odd. Adding 1 causes a carry.

So I decided to simulate the process directly on the string from right to left while handling the carry.

---

## Approach

1. Start from the rightmost bit.
2. Maintain a variable `carry` initialized to 0.
3. Traverse from right to left until index 1.
4. For each bit:

   * Compute current value = bit + carry.
   * If value == 1, it is odd.

     * We need two operations: add 1 and divide by 2.
     * Increase steps by 2.
     * Set carry to 1.
   * Otherwise, it is even.

     * Only divide by 2.
     * Increase steps by 1.
5. After finishing the loop, if carry is still 1, add one extra step.

This way, we simulate all operations without converting to a large integer.

---

## Data Structures Used

* Integer variable for counting steps.
* Integer variable for carry handling.
* No additional data structures are used.

---

## Operations & Behavior Summary

* Odd case (bit + carry == 1):

  * Add 1
  * Divide by 2
  * Total 2 steps

* Even case (bit + carry == 0 or 2):

  * Divide by 2
  * Total 1 step

* Final carry may introduce one additional step.

---

## Complexity

Time Complexity: O(n)

We traverse the string once from right to left.

Space Complexity: O(1)

We only use constant extra variables.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int numSteps(string s) {
        int steps = 0;
        int carry = 0;

        for (int i = s.size() - 1; i > 0; i--) {
            int bit = (s[i] - '0') + carry;

            if (bit == 1) {
                steps += 2;
                carry = 1;
            } else {
                steps += 1;
            }
        }

        return steps + carry;
    }
};
```

### Java

```java
class Solution {
    public int numSteps(String s) {
        int steps = 0;
        int carry = 0;

        for (int i = s.length() - 1; i > 0; i--) {
            int bit = (s.charAt(i) - '0') + carry;

            if (bit == 1) {
                steps += 2;
                carry = 1;
            } else {
                steps += 1;
            }
        }

        return steps + carry;
    }
}
```

### JavaScript

```javascript
var numSteps = function(s) {
    let steps = 0;
    let carry = 0;

    for (let i = s.length - 1; i > 0; i--) {
        let bit = (s[i] - '0') + carry;

        if (bit === 1) {
            steps += 2;
            carry = 1;
        } else {
            steps += 1;
        }
    }

    return steps + carry;
};
```

### Python3

```python
class Solution:
    def numSteps(self, s: str) -> int:
        steps = 0
        carry = 0

        for i in range(len(s) - 1, 0, -1):
            bit = int(s[i]) + carry

            if bit == 1:
                steps += 2
                carry = 1
            else:
                steps += 1

        return steps + carry
```

### Go

```go
func numSteps(s string) int {
    steps := 0
    carry := 0

    for i := len(s) - 1; i > 0; i-- {
        bit := int(s[i]-'0') + carry

        if bit == 1 {
            steps += 2
            carry = 1
        } else {
            steps += 1
        }
    }

    return steps + carry
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Initialize `steps = 0` and `carry = 0`.
2. Start from the last character of the string.
3. For each position except the first:

   * Add the current bit and carry.
   * If result is 1:

     * It means the number is odd.
     * Perform add 1 and divide.
     * Increase steps by 2.
     * Set carry to 1.
   * Otherwise:

     * It means the number is even.
     * Perform divide only.
     * Increase steps by 1.
4. After finishing the loop, add the remaining carry to steps.
5. Return total steps.

---

## Examples

Example 1:

Input: s = "1101"
Output: 6

Example 2:

Input: s = "10"
Output: 1

Example 3:

Input: s = "1"
Output: 0

---

## How to use / Run locally

C++:

* Compile using g++ filename.cpp
* Run using ./a.out

Java:

* Compile using javac Solution.java
* Run using java Solution

Python:

* Run using python filename.py

JavaScript:

* Run using node filename.js

Go:

* Run using go run filename.go

---

## Notes & Optimizations

* No need to convert binary string to integer.
* Avoid BigInteger to prevent overflow.
* One pass solution.
* Constant extra space.
* Efficient even for maximum constraint size.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
