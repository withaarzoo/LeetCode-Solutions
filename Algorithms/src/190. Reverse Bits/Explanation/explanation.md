# 190. Reverse Bits

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

Given a 32-bit unsigned integer, I need to reverse its bits and return the resulting number.

This means if the binary representation of the number is:

00000010100101000001111010011100

It should become:

00111001011110000010100101000000

The goal is to reverse all 32 bits efficiently.

---

## Constraints

* Input is a 32-bit unsigned integer.
* The integer contains exactly 32 bits.
* The solution should be efficient.

---

## Intuition

When I first saw this problem, I thought very simply.

I do not need strings.
I do not need extra arrays.

Bits are just 0 and 1.

So I thought:
If I keep taking the last bit from the number,
and keep building a new number from left to right,
then automatically the bits will get reversed.

So I decided to:

* Extract last bit using n & 1
* Shift result left
* Add the extracted bit
* Shift n right
* Repeat 32 times

That is the entire idea.

---

## Approach

Step by step what I did:

1. Initialize result = 0.
2. Run a loop 32 times.
3. In every iteration:

   * Shift result left by 1.
   * Extract last bit of n using (n & 1).
   * Add it to result.
   * Shift n right.
4. After 32 iterations, result contains reversed bits.
5. Return result.

Why 32 times?
Because the integer has exactly 32 bits.

---

## Data Structures Used

No extra data structures.

Only one integer variable:

* result

So space usage is constant.

---

## Operations & Behavior Summary

* Bit extraction using AND operator (&)
* Left shift to build reversed number
* Right shift to process next bit
* Loop runs fixed 32 times

This ensures efficient and clean bit manipulation.

---

## Complexity

**Time Complexity:** O(1)

The loop runs exactly 32 times.
Since 32 is constant, time complexity is constant.

**Space Complexity:** O(1)

We only use one extra variable.
No extra arrays or memory.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int reverseBits(int n) {
        unsigned int result = 0;
        
        for (int i = 0; i < 32; i++) {
            result <<= 1;            // Shift result left
            result |= (n & 1);       // Add last bit of n
            n >>= 1;                 // Shift n right
        }
        
        return result;
    }
};
```

---

### Java

```java
class Solution {
    public int reverseBits(int n) {
        int result = 0;
        
        for (int i = 0; i < 32; i++) {
            result <<= 1;
            result |= (n & 1);
            n >>>= 1;   // Logical shift
        }
        
        return result;
    }
}
```

---

### JavaScript

```javascript
var reverseBits = function(n) {
    let result = 0;
    
    for (let i = 0; i < 32; i++) {
        result = result << 1;
        result = result | (n & 1);
        n = n >>> 1;
    }
    
    return result >>> 0;
};
```

---

### Python3

```python
class Solution:
    def reverseBits(self, n: int) -> int:
        result = 0
        
        for _ in range(32):
            result <<= 1
            result |= (n & 1)
            n >>= 1
        
        return result
```

---

### Go

```go
func reverseBits(n int) int {
    result := 0
    
    for i := 0; i < 32; i++ {
        result <<= 1
        result |= (n & 1)
        n >>= 1
    }
    
    return result
}
```

---

## Step-by-step Detailed Explanation

Let us understand the core logic.

result <<= 1

This shifts result one bit to the left.
It creates space at the rightmost position.

result |= (n & 1)

n & 1 extracts the last bit of n.
If last bit is 1, it returns 1.
If last bit is 0, it returns 0.

We add this bit into result.

n >>= 1

This shifts n to the right.
Now the next bit becomes the last bit.

After 32 iterations, all bits are reversed.

---

## Examples

Example 1:

Input: 43261596
Output: 964176192

Example 2:

Input: 2147483644
Output: 1073741822

---

## How to use / Run locally

C++

* Compile using: g++ filename.cpp
* Run: ./a.out

Java

* Compile: javac Solution.java
* Run: java Solution

Python

* Run: python filename.py

JavaScript

* Run using Node.js: node filename.js

Go

* Run: go run filename.go

---

## Notes & Optimizations

If this function is called many times,
I can optimize further using a lookup table.

Idea:

* Precompute reversed values of all 8-bit numbers (0 to 255).
* Break 32-bit number into four 8-bit parts.
* Reverse each part using lookup.
* Combine them back.

This reduces repeated computation.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
