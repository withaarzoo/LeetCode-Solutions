# 788. Rotated Digits

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

This problem asks us to count how many numbers between 1 and n are "good" after rotating each digit by 180 degrees.

When a number is rotated:

* Some digits remain the same
* Some digits change into different digits
* Some digits become invalid

A number is considered valid only if all its digits still form a valid number after rotation.
A number is considered good only if it is valid and also changes into a different number.

So the goal is simple:
Given an integer n, return how many numbers from 1 to n are good after rotation.

This problem is commonly known as a digit-based validation problem in competitive programming and is often used to test number manipulation and logic building skills.

---

## Constraints

* 1 ≤ n ≤ 10^4

---

## Intuition

When I first looked at this problem, I focused on how digits behave after rotation.

I quickly noticed that:

* 0, 1, 8 stay the same
* 2 ↔ 5 and 6 ↔ 9 swap with each other
* 3, 4, 7 become invalid

That means a number is only valid if it doesn’t contain 3, 4, or 7.

But just being valid is not enough.
At least one digit must change after rotation. Otherwise, the number remains the same and is not counted as good.

This gave me a clear direction:
Check every number and classify it based on its digits.

---

## Approach

I iterate through all numbers from 1 to n.

For each number:

1. Extract each digit one by one.
2. Check if the digit is invalid (3, 4, 7). If yes, reject the number immediately.
3. Check if the digit changes (2, 5, 6, 9). If yes, mark that the number changes.
4. Continue until all digits are processed.

At the end:

* If the number is valid and has at least one changing digit, I count it.

Finally, I return the total count.

This is a brute-force approach, but since n is small (≤ 10^4), it works efficiently.

---

## Data Structures Used

* Integer variables
  Used to store numbers, digits, and counters.

* Boolean flags
  Used to track:

  * Whether a number is valid
  * Whether it changes after rotation

No complex data structures are required for this problem.

---

## Operations & Behavior Summary

For every number from 1 to n:

* Break the number into digits
* For each digit:

  * If it is 3, 4, or 7 → mark invalid and stop
  * If it is 2, 5, 6, or 9 → mark as changing
* After processing all digits:

  * If valid AND changing → increment count

Return the final count

---

## Complexity

| Type             | Complexity | Explanation                                              |
| ---------------- | ---------- | -------------------------------------------------------- |
| Time Complexity  | O(n * d)   | n numbers, and each number has up to d digits (log10(n)) |
| Space Complexity | O(1)       | Only a few variables are used, no extra memory needed    |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int rotatedDigits(int n) {
        int count = 0; // this will store the number of good numbers
        
        for (int i = 1; i <= n; i++) {
            int num = i;
            bool isValid = true;   // assume number is valid initially
            bool hasChange = false; // to check if at least one digit changes
            
            while (num > 0) {
                int digit = num % 10; // extract last digit
                
                // if digit is invalid after rotation
                if (digit == 3 || digit == 4 || digit == 7) {
                    isValid = false;
                    break; // no need to check further
                }
                
                // if digit changes after rotation
                if (digit == 2 || digit == 5 || digit == 6 || digit == 9) {
                    hasChange = true;
                }
                
                num /= 10; // remove last digit
            }
            
            // count only if valid and has at least one changing digit
            if (isValid && hasChange) {
                count++;
            }
        }
        
        return count;
    }
};
```

### Java

```java
class Solution {
    public int rotatedDigits(int n) {
        int count = 0; // total good numbers
        
        for (int i = 1; i <= n; i++) {
            int num = i;
            boolean isValid = true;   // assume valid
            boolean hasChange = false; // check if it changes
            
            while (num > 0) {
                int digit = num % 10; // extract last digit
                
                // invalid digits
                if (digit == 3 || digit == 4 || digit == 7) {
                    isValid = false;
                    break;
                }
                
                // digits that change
                if (digit == 2 || digit == 5 || digit == 6 || digit == 9) {
                    hasChange = true;
                }
                
                num /= 10; // remove last digit
            }
            
            if (isValid && hasChange) {
                count++;
            }
        }
        
        return count;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var rotatedDigits = function(n) {
    let count = 0; // total good numbers
    
    for (let i = 1; i <= n; i++) {
        let num = i;
        let isValid = true;   // assume valid
        let hasChange = false; // track if it changes
        
        while (num > 0) {
            let digit = num % 10; // get last digit
            
            // invalid digits
            if (digit === 3 || digit === 4 || digit === 7) {
                isValid = false;
                break;
            }
            
            // changing digits
            if (digit === 2 || digit === 5 || digit === 6 || digit === 9) {
                hasChange = true;
            }
            
            num = Math.floor(num / 10); // remove last digit
        }
        
        if (isValid && hasChange) {
            count++;
        }
    }
    
    return count;
};
```

### Python3

```python
class Solution:
    def rotatedDigits(self, n: int) -> int:
        count = 0  # total good numbers
        
        for i in range(1, n + 1):
            num = i
            isValid = True   # assume valid
            hasChange = False  # check if it changes
            
            while num > 0:
                digit = num % 10  # extract last digit
                
                # invalid digits
                if digit in [3, 4, 7]:
                    isValid = False
                    break
                
                # digits that change
                if digit in [2, 5, 6, 9]:
                    hasChange = True
                
                num //= 10  # remove last digit
            
            if isValid and hasChange:
                count += 1
        
        return count
```

### Go

```go
func rotatedDigits(n int) int {
    count := 0 // total good numbers
    
    for i := 1; i <= n; i++ {
        num := i
        isValid := true   // assume valid
        hasChange := false // check if it changes
        
        for num > 0 {
            digit := num % 10 // extract last digit
            
            // invalid digits
            if digit == 3 || digit == 4 || digit == 7 {
                isValid = false
                break
            }
            
            // digits that change
            if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
                hasChange = true
            }
            
            num /= 10 // remove last digit
        }
        
        if isValid && hasChange {
            count++
        }
    }
    
    return count
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains the same across all languages.

I start by initializing a counter to store how many good numbers I find.

Then I loop from 1 to n:

* For each number, I copy it into another variable so I can extract digits.

I use two flags:

* One to check if the number is valid
* One to check if at least one digit changes

Now I process digits:

* Extract the last digit using modulo
* Check:

  * If digit is 3, 4, or 7 → mark invalid and break
  * If digit is 2, 5, 6, or 9 → mark as changing
* Remove the last digit using division

After all digits are checked:

* If valid and changing → increment the count

This exact logic works in all five languages. Only syntax differs.

Edge case:

* If n = 1 → answer is 0 because 1 does not change after rotation

---

## Examples

Example 1
Input: n = 10
Output: 4

Explanation:
Good numbers are 2, 5, 6, 9
Each of these changes to a different valid number after rotation

---

Example 2
Input: n = 1
Output: 0

Explanation:
1 remains the same after rotation, so it is not good

---

Example 3
Input: n = 20
Output: 9

Explanation:
Valid and changing numbers include:
2, 5, 6, 9, 12, 15, 16, 19, 20

---

## How to Use / Run Locally

C++:

* Save file as solution.cpp
* Compile: g++ solution.cpp -o solution
* Run: ./solution

Java:

* Save file as Solution.java
* Compile: javac Solution.java
* Run: java Solution

JavaScript:

* Save file as solution.js
* Run: node solution.js

Python:

* Save file as solution.py
* Run: python solution.py

Go:

* Save file as solution.go
* Run: go run solution.go

---

## Notes & Optimizations

* This is a brute-force solution, but efficient enough due to small constraints.
* Avoid using string conversion. Digit extraction using modulo is faster.
* Always check invalid digits first to exit early.
* Make sure at least one digit changes, otherwise do not count the number.
* This problem is a good example of digit manipulation and condition-based filtering.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
