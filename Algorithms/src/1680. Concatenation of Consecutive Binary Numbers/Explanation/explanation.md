# Problem Title

1. Concatenation of Consecutive Binary Numbers

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

## Problem Summary

Given an integer n, return the decimal value of the binary string formed by concatenating the binary representations of 1 to n in order. Since the result can be very large, return it modulo 10^9 + 7.

Instead of building the full binary string explicitly, we need an optimized mathematical and bit manipulation based approach.

## Constraints

1 <= n <= 10^5

## Intuition

When I first read the problem, I thought about directly building the binary string from 1 to n. But that would be inefficient because the string becomes extremely large.

Then I realized something important. If I already have a number and I want to append another number in binary form, I do not need to use strings. I can use bit shifting.

Appending a binary number is equivalent to:

1. Left shifting the current result by the number of bits in the new number.
2. Adding the new number.

Also, I noticed that the number of bits increases whenever we reach a power of 2.

For example:
1 has 1 bit
2 has 2 bits
4 has 3 bits
8 has 4 bits

So I only need to track when the bit length increases.

## Approach

1. Initialize result as 0.
2. Maintain a variable bitLength to track number of bits.
3. Loop from 1 to n.
4. If the current number is a power of 2, increase bitLength.
5. Left shift result by bitLength.
6. Add the current number.
7. Take modulo 10^9 + 7 at every step.
8. Return the final result.

## Data Structures Used

No additional data structures are used.
Only primitive variables are used.

## Operations & Behavior Summary

For each number from 1 to n:

* Check if it is a power of 2 using (i & (i - 1)) == 0
* Increase bit length when required
* Perform left shift
* Add the current number
* Apply modulo

This ensures correctness and efficiency.

## Complexity

Time Complexity: O(n)
We iterate once from 1 to n.

Space Complexity: O(1)
We use only a few variables.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int concatenatedBinary(int n) {
        const int MOD = 1e9 + 7;
        long long ans = 0;
        int bitLength = 0;
        
        for (int i = 1; i <= n; i++) {
            if ((i & (i - 1)) == 0) {
                bitLength++;
            }
            ans = ((ans << bitLength) % MOD + i) % MOD;
        }
        
        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int concatenatedBinary(int n) {
        final int MOD = 1_000_000_007;
        long ans = 0;
        int bitLength = 0;
        
        for (int i = 1; i <= n; i++) {
            if ((i & (i - 1)) == 0) {
                bitLength++;
            }
            ans = ((ans << bitLength) % MOD + i) % MOD;
        }
        
        return (int) ans;
    }
}
```

### JavaScript

```javascript
var concatenatedBinary = function(n) {
    const MOD = 1000000007n;
    let ans = 0n;
    let bitLength = 0n;
    
    for (let i = 1n; i <= BigInt(n); i++) {
        if ((i & (i - 1n)) === 0n) {
            bitLength++;
        }
        ans = ((ans << bitLength) % MOD + i) % MOD;
    }
    
    return Number(ans);
};
```

### Python3

```python
class Solution:
    def concatenatedBinary(self, n: int) -> int:
        MOD = 10**9 + 7
        ans = 0
        bit_length = 0
        
        for i in range(1, n + 1):
            if (i & (i - 1)) == 0:
                bit_length += 1
            ans = ((ans << bit_length) + i) % MOD
        
        return ans
```

### Go

```go
func concatenatedBinary(n int) int {
    const MOD = 1000000007
    var ans int64 = 0
    bitLength := 0
    
    for i := 1; i <= n; i++ {
        if (i & (i - 1)) == 0 {
            bitLength++
        }
        ans = ((ans << bitLength) % MOD + int64(i)) % MOD
    }
    
    return int(ans)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Initialize ans = 0.
2. Initialize bitLength = 0.
3. Iterate from i = 1 to n.
4. Check if i is a power of 2 using (i & (i - 1)) == 0.
5. If true, increment bitLength.
6. Left shift ans by bitLength.
7. Add i to ans.
8. Take modulo at every step.
9. Continue until n.
10. Return ans.

The logic remains identical across all languages.

## Examples

Example 1
Input: n = 1
Output: 1

Example 2
Input: n = 3
Binary formed: 1, 10, 11
Concatenated: 11011
Output: 27

Example 3
Input: n = 12
Output: 505379714

## How to use / Run locally

C++

* Compile using g++ filename.cpp
* Run ./a.out

Java

* Compile using javac Solution.java
* Run using java Solution

Python

* Run using python filename.py

JavaScript

* Run using node filename.js

Go

* Run using go run filename.go

## Notes & Optimizations

* Avoid building large strings.
* Always apply modulo at each step to prevent overflow.
* Detect power of 2 efficiently using bit manipulation.
* Use long or int64 where necessary to avoid overflow.

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
