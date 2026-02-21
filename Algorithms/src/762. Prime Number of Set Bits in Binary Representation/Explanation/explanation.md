# 762. Prime Number of Set Bits in Binary Representation

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
* Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given two integers `left` and `right`, return the count of numbers in the inclusive range `[left, right]` whose number of set bits (1s in binary representation) is a **prime number**.

A set bit means how many `1`s are present in the binary form of a number.

Example:
21 in binary is `10101` → It has 3 set bits → 3 is prime.

---

## Constraints

* 1 <= left <= right <= 10^6
* 0 <= right - left <= 10^4

---

## Intuition

When I first read the problem, I realized I need to do three things for every number in the range:

1. Convert number to binary
2. Count number of 1s
3. Check if that count is prime

Then I thought carefully.

Since `right <= 10^6`, the maximum binary length is less than 20 bits.

So maximum possible set bits is at most 20.

That means I only need to check prime numbers up to 20.

Prime numbers up to 20 are:
2, 3, 5, 7, 11, 13, 17, 19

So I can store these in a set for quick lookup.

That makes the solution simple and very efficient.

---

## Approach

Step 1: Store small prime numbers (up to 20) in a set.

Step 2: Initialize answer counter = 0.

Step 3: Loop from left to right.

Step 4: For each number:

* Count set bits.
* If count is in prime set → increment answer.

Step 5: Return answer.

---

## Data Structures Used

* HashSet / Set → To store prime numbers.
* Simple integer variable for counting.

No complex data structure required.

---

## Operations & Behavior Summary

For each number in range:

* Bit counting is constant time (max 20 bits).
* Prime check is O(1) using hash set.

Total work depends only on range size.

---

## Complexity

### Time Complexity: O(n)

Where n = right - left + 1

For each number, we count bits and check prime in constant time.

### Space Complexity: O(1)

Only small fixed prime set is used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countPrimeSetBits(int left, int right) {
        unordered_set<int> primes = {2,3,5,7,11,13,17,19};
        int ans = 0;
        
        for(int num = left; num <= right; num++) {
            int setBits = __builtin_popcount(num);
            if(primes.count(setBits)) {
                ans++;
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int countPrimeSetBits(int left, int right) {
        java.util.Set<Integer> primeSet = new java.util.HashSet<>();
        int[] primes = {2,3,5,7,11,13,17,19};
        
        for(int p : primes) primeSet.add(p);
        
        int ans = 0;
        
        for(int num = left; num <= right; num++) {
            int setBits = Integer.bitCount(num);
            if(primeSet.contains(setBits)) {
                ans++;
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
var countPrimeSetBits = function(left, right) {
    const primeSet = new Set([2,3,5,7,11,13,17,19]);
    let ans = 0;

    for(let num = left; num <= right; num++) {
        let setBits = num.toString(2).split('1').length - 1;
        if(primeSet.has(setBits)) {
            ans++;
        }
    }
    return ans;
};
```

### Python3

```python
class Solution:
    def countPrimeSetBits(self, left: int, right: int) -> int:
        primes = {2,3,5,7,11,13,17,19}
        ans = 0
        
        for num in range(left, right + 1):
            set_bits = bin(num).count('1')
            if set_bits in primes:
                ans += 1
        
        return ans
```

### Go

```go
import "math/bits"

func countPrimeSetBits(left int, right int) int {
    primeSet := map[int]bool{
        2:true, 3:true, 5:true, 7:true,
        11:true, 13:true, 17:true, 19:true,
    }
    
    ans := 0
    
    for num := left; num <= right; num++ {
        setBits := bits.OnesCount(uint(num))
        if primeSet[setBits] {
            ans++
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. Create a prime set containing primes up to 20.
2. Initialize ans = 0.
3. Start loop from left to right.
4. For each number:

   * Count set bits using built-in function.
   * If setBits is in prime set → increment ans.
5. Return ans.

Each language uses its own built-in bit counting function for efficiency.

---

## Examples

Example 1:
Input: left = 6, right = 10
Output: 4

Explanation:
6 → 110 → 2 set bits → prime
7 → 111 → 3 set bits → prime
8 → 1000 → 1 set bit → not prime
9 → 1001 → 2 set bits → prime
10 → 1010 → 2 set bits → prime

Total = 4

Example 2:
Input: left = 10, right = 15
Output: 5

---

## How to use / Run locally

C++

* Compile using g++ filename.cpp
* Run using ./a.out

Java

* Compile using javac Solution.java
* Run using java Solution

Python

* Run using python filename.py

JavaScript

* Run using node filename.js

Go

* Run using go run filename.go

---

## Notes & Optimizations

* Maximum set bits possible is 20.
* So prime checking is constant time.
* Built-in bit counting functions make solution very fast.
* No need for sieve or large prime computation.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
