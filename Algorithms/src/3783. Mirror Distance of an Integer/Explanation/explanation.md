# Mirror Distance of an Integer

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

You are given an integer `n`.

The mirror distance of `n` is defined as:

```text
abs(n - reverse(n))
```

Here, `reverse(n)` means reversing the digits of `n`.

Your task is to return the mirror distance of the integer.

### Example

```text
Input: n = 25
Output: 27
```

Explanation:

```text
reverse(25) = 52
abs(25 - 52) = 27
```

## Constraints

```text
1 <= n <= 10^9
```

## Intuition

I thought about what the problem is really asking.

I only need to do two things:

1. Reverse the digits of the number.
2. Find the absolute difference between the original number and the reversed number.

For example:

```text
n = 120
reverse(n) = 21
answer = abs(120 - 21) = 99
```

So the main task is reversing the number correctly.

## Approach

1. Create a variable called `rev` to store the reversed number.
2. Store `n` in another variable called `temp`.
3. Use a loop while `temp > 0`.
4. In each iteration:

   * Get the last digit using `temp % 10`
   * Add it to `rev`
   * Remove the last digit from `temp`
5. After the loop ends, calculate:

```text
abs(n - rev)
```

1. Return the final answer.

## Data Structures Used

No special data structure is needed.

I only use:

* Integer variables
* A loop

## Operations & Behavior Summary

| Operation          | Purpose                    |
| ------------------ | -------------------------- |
| `temp % 10`        | Gets the last digit        |
| `temp / 10`        | Removes the last digit     |
| `rev * 10 + digit` | Builds the reversed number |
| `abs(n - rev)`     | Finds the mirror distance  |

## Complexity

* Time Complexity: `O(d)`

  * `d` is the number of digits in the integer.
* Space Complexity: `O(1)`

  * No extra data structure is used.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int mirrorDistance(int n) {
        int rev = 0;
        int temp = n;

        // Reverse the digits of n
        while (temp > 0) {
            int digit = temp % 10;
            rev = rev * 10 + digit;
            temp /= 10;
        }

        // Return the absolute difference
        return abs(n - rev);
    }
};
```

### Java

```java
class Solution {
    public int mirrorDistance(int n) {
        int rev = 0;
        int temp = n;

        // Reverse the digits of n
        while (temp > 0) {
            int digit = temp % 10;
            rev = rev * 10 + digit;
            temp /= 10;
        }

        // Return the absolute difference
        return Math.abs(n - rev);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {number}
 */
var mirrorDistance = function(n) {
    let rev = 0;
    let temp = n;

    // Reverse the digits of n
    while (temp > 0) {
        let digit = temp % 10;
        rev = rev * 10 + digit;
        temp = Math.floor(temp / 10);
    }

    // Return the absolute difference
    return Math.abs(n - rev);
};
```

### Python3

```python
class Solution:
    def mirrorDistance(self, n: int) -> int:
        rev = 0
        temp = n

        # Reverse the digits of n
        while temp > 0:
            digit = temp % 10
            rev = rev * 10 + digit
            temp //= 10

        # Return the absolute difference
        return abs(n - rev)
```

### Go

```go
func mirrorDistance(n int) int {
    rev := 0
    temp := n

    // Reverse the digits of n
    for temp > 0 {
        digit := temp % 10
        rev = rev*10 + digit
        temp /= 10
    }

    // Return the absolute difference
    if n > rev {
        return n - rev
    }

    return rev - n
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Suppose:

```text
n = 250
```

Initially:

```text
rev = 0
temp = 250
```

### First Iteration

```text
digit = 250 % 10 = 0
rev = 0 * 10 + 0 = 0
temp = 250 / 10 = 25
```

### Second Iteration

```text
digit = 25 % 10 = 5
rev = 0 * 10 + 5 = 5
temp = 25 / 10 = 2
```

### Third Iteration

```text
digit = 2 % 10 = 2
rev = 5 * 10 + 2 = 52
temp = 2 / 10 = 0
```

Now the reversed number becomes:

```text
rev = 52
```

Final answer:

```text
abs(250 - 52) = 198
```

So the output is:

```text
198
```

## Examples

### Example 1

```text
Input: n = 25
Output: 27
```

Explanation:

```text
reverse(25) = 52
abs(25 - 52) = 27
```

### Example 2

```text
Input: n = 10
Output: 9
```

Explanation:

```text
reverse(10) = 1
abs(10 - 1) = 9
```

### Example 3

```text
Input: n = 7
Output: 0
```

Explanation:

```text
reverse(7) = 7
abs(7 - 7) = 0
```

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node main.js
```

### Python3

```bash
python main.py
```

### Go

```bash
go run main.go
```

## Notes & Optimizations

* I do not use strings because reversing digits mathematically is more efficient.
* The solution only needs a few integer variables.
* Since the maximum value of `n` is `10^9`, integer overflow is not an issue here.
* The algorithm is already optimal because every digit is visited only once.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
