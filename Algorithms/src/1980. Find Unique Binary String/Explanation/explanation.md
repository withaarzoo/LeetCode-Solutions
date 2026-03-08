# 1980. Find Unique Binary String

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given an array `nums` containing **n unique binary strings**, each of length `n`, we must return a **binary string of length n that does not appear in the array**.

If multiple answers exist, we can return **any one of them**.

The key observation is that although there are `n` strings provided, the total number of binary strings of length `n` is:

```
2^n
```

Therefore many possible strings are not present in the input. Our goal is to construct one efficiently.

---

## Constraints

```
n == nums.length
1 <= n <= 16
nums[i].length == n
nums[i] consists only of '0' and '1'
All strings in nums are unique
```

---

## Intuition

When I first looked at the problem, I noticed something interesting.

There are `n` binary strings and each string also has length `n`.

Instead of generating all `2^n` possible binary strings and checking which one is missing, I thought about how I could **guarantee a different string directly**.

This reminded me of a classic concept called **Cantor's Diagonalization**.

If I build a new string by **flipping the i-th bit of the i-th string**, then the new string will differ from every string in the array at least at one position.

So the generated string **cannot exist in the list**.

---

## Approach

Step-by-step strategy:

1. Let `n` be the number of binary strings.
2. Create an empty result string.
3. Iterate from index `0` to `n-1`.
4. Look at the `i-th` character of the `i-th` string.
5. Flip that bit:

   * If it is `'0'`, append `'1'`
   * If it is `'1'`, append `'0'`
6. Continue this for all positions.
7. Return the constructed string.

This ensures the resulting string differs from each string in the array at least at one index.

---

## Data Structures Used

* String / StringBuilder
* Character comparison

No extra complex data structures are required.

---

## Operations & Behavior Summary

| Operation           | Purpose                              |
| ------------------- | ------------------------------------ |
| Access diagonal bit | Compare each string at position i    |
| Flip bit            | Ensure resulting string is different |
| Append character    | Build final unique string            |

---

## Complexity

**Time Complexity:**

```
O(n)
```

We iterate through the list once and perform constant operations.

`n` = number of binary strings.

**Space Complexity:**

```
O(n)
```

We construct a result string of length `n`.

---

# Multi-language Solutions

## C++

```cpp
class Solution {
public:
    string findDifferentBinaryString(vector<string>& nums) {
        int n = nums.size();
        string result = "";

        for(int i = 0; i < n; i++){
            if(nums[i][i] == '0')
                result += '1';
            else
                result += '0';
        }

        return result;
    }
};
```

---

## Java

```java
class Solution {
    public String findDifferentBinaryString(String[] nums) {
        int n = nums.length;
        StringBuilder result = new StringBuilder();

        for(int i = 0; i < n; i++){
            if(nums[i].charAt(i) == '0')
                result.append('1');
            else
                result.append('0');
        }

        return result.toString();
    }
}
```

---

## JavaScript

```javascript
var findDifferentBinaryString = function(nums) {
    let n = nums.length;
    let result = "";

    for(let i = 0; i < n; i++){
        result += nums[i][i] === '0' ? '1' : '0';
    }

    return result;
};
```

---

## Python3

```python
class Solution:
    def findDifferentBinaryString(self, nums: List[str]) -> str:
        n = len(nums)
        result = []

        for i in range(n):
            if nums[i][i] == '0':
                result.append('1')
            else:
                result.append('0')

        return ''.join(result)
```

---

## Go

```go
func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    result := make([]byte, n)

    for i := 0; i < n; i++ {
        if nums[i][i] == '0' {
            result[i] = '1'
        } else {
            result[i] = '0'
        }
    }

    return string(result)
}
```

---

# Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core idea is to build a string that differs from every given string.

Assume the input is:

```
nums = ["01", "10"]
```

Step 1: Determine `n`

```
n = 2
```

Step 2: Look at diagonal elements

```
nums[0][0]
nums[1][1]
```

Step 3: Flip the bits

```
nums[0][0] = '0' -> '1'
nums[1][1] = '0' -> '1'
```

Step 4: Construct result

```
result = "11"
```

Step 5: Verify uniqueness

```
"11" != "01"
"11" != "10"
```

Thus the string does not appear in the list.

---

# Examples

### Example 1

Input

```
nums = ["01","10"]
```

Output

```
"11"
```

---

### Example 2

Input

```
nums = ["00","01"]
```

Output

```
"11"
```

---

### Example 3

Input

```
nums = ["111","011","001"]
```

Output

```
"101"
```

---

# How to use / Run locally

Clone the repository:

```
git clone https://github.com/your-username/leetcode-solutions
```

Navigate to the project folder:

```
cd leetcode-solutions
```

Compile and run depending on the language.

Example for C++:

```
g++ solution.cpp
./a.out
```

Example for Python:

```
python solution.py
```

---

# Notes & Optimizations

* This solution uses **Cantor's diagonalization technique**.
* It guarantees uniqueness without generating all possible binary strings.
* Avoids brute force search (`O(2^n)`).
* Works efficiently even at the maximum constraint.

---

# Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
