# 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers

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

A decimal number is called deci-binary if each of its digits is either 0 or 1 and it does not contain any leading zeros.

Given a string `n` that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to `n`.

---

## Constraints

* 1 <= n.length <= 10^5
* n consists only of digits
* n does not contain any leading zeros

---

## Intuition

When I first looked at this problem, I tried to understand what a deci-binary number really means.

A deci-binary number can only contain digits 0 or 1.

That means at any digit position, each deci-binary number can contribute at most 1.

So if a digit in the number `n` is 7, that means I need at least 7 deci-binary numbers to build that digit.

Because each deci-binary number can only add 1 at that position.

So the minimum number of deci-binary numbers required will always be equal to the maximum digit present in `n`.

That is the key observation.

---

## Approach

1. I initialize a variable to store the maximum digit.
2. I iterate through each character of the string.
3. I convert the character into an integer digit.
4. I keep updating the maximum digit.
5. If I ever find digit 9, I stop early because 9 is the highest possible digit.
6. I return the maximum digit.

That maximum digit is the answer.

---

## Data Structures Used

* No extra data structure is required.
* Only a single integer variable is used to store the maximum digit.

---

## Operations & Behavior Summary

* Traverse the string once.
* Convert each character into an integer.
* Track the maximum digit.
* Return the maximum digit as the result.

---

## Complexity

Time Complexity: O(n)

* I traverse the string once.
* n is the length of the string.

Space Complexity: O(1)

* I use only one extra variable.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minPartitions(string n) {
        int maxDigit = 0;
        
        for(char c : n) {
            int digit = c - '0';
            
            if(digit > maxDigit) {
                maxDigit = digit;
            }
            
            if(maxDigit == 9) {
                break;
            }
        }
        
        return maxDigit;
    }
};
```

### Java

```java
class Solution {
    public int minPartitions(String n) {
        int maxDigit = 0;
        
        for(int i = 0; i < n.length(); i++) {
            int digit = n.charAt(i) - '0';
            
            if(digit > maxDigit) {
                maxDigit = digit;
            }
            
            if(maxDigit == 9) {
                break;
            }
        }
        
        return maxDigit;
    }
}
```

### JavaScript

```javascript
var minPartitions = function(n) {
    let maxDigit = 0;
    
    for (let i = 0; i < n.length; i++) {
        let digit = n[i] - '0';
        
        if (digit > maxDigit) {
            maxDigit = digit;
        }
        
        if (maxDigit === 9) {
            break;
        }
    }
    
    return maxDigit;
};
```

### Python3

```python
class Solution:
    def minPartitions(self, n: str) -> int:
        max_digit = 0
        
        for ch in n:
            digit = int(ch)
            
            if digit > max_digit:
                max_digit = digit
            
            if max_digit == 9:
                break
        
        return max_digit
```

### Go

```go
func minPartitions(n string) int {
    maxDigit := 0
    
    for i := 0; i < len(n); i++ {
        digit := int(n[i] - '0')
        
        if digit > maxDigit {
            maxDigit = digit
        }
        
        if maxDigit == 9 {
            break
        }
    }
    
    return maxDigit
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Step 1: Initialize a variable

We create a variable called maxDigit and set it to 0. This will store the largest digit found so far.

Step 2: Traverse the string

We loop through each character of the string `n`.

Step 3: Convert character to integer

In C++ and Java, we subtract '0' from the character.
In Python, we use int().
In JavaScript, subtracting '0' automatically converts it to a number.
In Go, we subtract '0' from the byte value.

Step 4: Update maximum

If the current digit is greater than maxDigit, we update maxDigit.

Step 5: Early stopping

If maxDigit becomes 9, we break the loop because 9 is the maximum possible digit.

Step 6: Return result

Finally, we return maxDigit as the answer.

---

## Examples

Example 1
Input: n = "32"
Output: 3
Explanation: The maximum digit is 3.

Example 2
Input: n = "82734"
Output: 8
Explanation: The maximum digit is 8.

Example 3
Input: n = "27346209830709182346"
Output: 9
Explanation: The maximum digit is 9.

---

## How to use / Run locally

C++

1. Save the code in a file named solution.cpp
2. Compile using: g++ solution.cpp
3. Run using: ./a.out

Java

1. Save the file as Solution.java
2. Compile using: javac Solution.java
3. Run using: java Solution

Python

1. Save the file as solution.py
2. Run using: python solution.py

JavaScript

1. Save the file as solution.js
2. Run using: node solution.js

Go

1. Save the file as solution.go
2. Run using: go run solution.go

---

## Notes & Optimizations

* The key insight is that the answer is the maximum digit.
* We do not need to simulate the addition of deci-binary numbers.
* Early stopping when digit 9 is found improves performance in worst cases.
* The solution works efficiently even for length up to 10^5.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
