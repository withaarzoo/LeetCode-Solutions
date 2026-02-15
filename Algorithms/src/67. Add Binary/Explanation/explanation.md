# 67. Add Binary

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

Given two binary strings `a` and `b`, return their sum as a binary string.

We are not allowed to convert them directly into integers because the length can be large. So we must simulate binary addition manually.

---

## Constraints

* 1 <= a.length, b.length <= 10^4
* a and b contain only characters '0' or '1'
* No leading zeros except for the number zero itself

---

## Intuition

When I saw this problem, I immediately thought about how we add numbers by hand.

Binary addition works the same way as decimal addition, but digits are only 0 and 1.

Rules:

* 0 + 0 = 0
* 1 + 0 = 1
* 1 + 1 = 0 and carry 1

So I decided to:

* Start from the rightmost digit
* Keep a carry
* Add digits step by step

This avoids overflow and works for very large strings.

---

## Approach

1. Start two pointers from the end of both strings.
2. Maintain a variable called carry.
3. Loop while:

   * Either pointer is valid
   * Or carry is not zero
4. Add digits and carry.
5. Append (sum % 2) to result.
6. Update carry = sum // 2.
7. Reverse the result at the end.

---

## Data Structures Used

* String / StringBuilder
* Dynamic array (in JS and Python)
* Simple integer variable for carry

No extra heavy data structures are required.

---

## Operations & Behavior Summary

* Traverse both strings from right to left
* Simulate binary addition
* Maintain carry
* Reverse final result
* Return final binary string

---

## Complexity

Let n = length of a
Let m = length of b

Time Complexity: O(max(n, m))
We traverse both strings once.

Space Complexity: O(max(n, m))
We store the result string.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string addBinary(string a, string b) {
        int i = a.size() - 1;
        int j = b.size() - 1;
        int carry = 0;
        string result = "";

        while (i >= 0 || j >= 0 || carry) {
            int sum = carry;

            if (i >= 0) sum += a[i--] - '0';
            if (j >= 0) sum += b[j--] - '0';

            result += (sum % 2) + '0';
            carry = sum / 2;
        }

        reverse(result.begin(), result.end());
        return result;
    }
};
```

---

### Java

```java
class Solution {
    public String addBinary(String a, String b) {
        int i = a.length() - 1;
        int j = b.length() - 1;
        int carry = 0;
        StringBuilder result = new StringBuilder();

        while (i >= 0 || j >= 0 || carry != 0) {
            int sum = carry;

            if (i >= 0) sum += a.charAt(i--) - '0';
            if (j >= 0) sum += b.charAt(j--) - '0';

            result.append(sum % 2);
            carry = sum / 2;
        }

        return result.reverse().toString();
    }
}
```

---

### JavaScript

```javascript
var addBinary = function(a, b) {
    let i = a.length - 1;
    let j = b.length - 1;
    let carry = 0;
    let result = [];

    while (i >= 0 || j >= 0 || carry > 0) {
        let sum = carry;

        if (i >= 0) sum += a[i--] - '0';
        if (j >= 0) sum += b[j--] - '0';

        result.push(sum % 2);
        carry = Math.floor(sum / 2);
    }

    return result.reverse().join("");
};
```

---

### Python3

```python
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        i = len(a) - 1
        j = len(b) - 1
        carry = 0
        result = []

        while i >= 0 or j >= 0 or carry:
            total = carry

            if i >= 0:
                total += int(a[i])
                i -= 1
            if j >= 0:
                total += int(b[j])
                j -= 1

            result.append(str(total % 2))
            carry = total // 2

        return "".join(result[::-1])
```

---

### Go

```go
func addBinary(a string, b string) string {
    i := len(a) - 1
    j := len(b) - 1
    carry := 0
    result := []byte{}

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry

        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        result = append(result, byte(sum%2)+'0')
        carry = sum / 2
    }

    for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
        result[l], result[r] = result[r], result[l]
    }

    return string(result)
}
```

---

## Step-by-step Detailed Explanation

1. Initialize two pointers at the end of both strings.
2. Create a carry variable.
3. Loop while at least one pointer is valid or carry exists.
4. Add carry to sum.
5. Add digit from a if pointer is valid.
6. Add digit from b if pointer is valid.
7. Append sum % 2 to result.
8. Update carry as sum / 2.
9. Move pointers left.
10. Reverse the result.

This simulates manual binary addition.

---

## Examples

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

---

## How to use / Run locally

C++:

* Compile using g++
* Run the executable

Java:

* Compile using javac
* Run using java

Python:

* Run using python3 filename.py

JavaScript:

* Run using node filename.js

Go:

* Run using go run filename.go

---

## Notes & Optimizations

* We never convert strings into integers.
* This avoids overflow.
* Works efficiently for length up to 10^4.
* Time optimal solution.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
