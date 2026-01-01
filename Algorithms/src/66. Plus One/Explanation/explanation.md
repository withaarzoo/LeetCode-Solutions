# Problem Title

**66. Plus One (LeetCode)**

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
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a large integer represented as an array of digits.
Each element of the array stores a single digit.
The digits are ordered from most significant to least significant.

My task is to **add one to this number** and return the updated digits array.

The number does not contain any leading zeros.

---

## Constraints

* 1 ≤ digits.length ≤ 100
* 0 ≤ digits[i] ≤ 9
* The array does not contain leading zeros

---

## Intuition

I thought about how I normally add 1 to a number using pen and paper.

I always start adding from the **last digit**.

* If the digit is less than 9, I just add 1 and stop.
* If the digit is 9, it becomes 0 and I carry 1 to the left.
* This carry may continue for multiple digits.
* If all digits are 9, I need to add an extra digit at the front.

So the main challenge is **handling carry properly from right to left**.

---

## Approach

1. I start from the last index of the digits array.
2. I add 1 to the current digit.
3. If the digit becomes less than 10, I return the result immediately.
4. If the digit becomes 10, I set it to 0 and move left.
5. If the loop finishes and carry still exists, I add 1 at the beginning.

This approach is simple, fast, and memory efficient.

---

## Data Structures Used

* Array / List to store digits
* No extra complex data structures are required

---

## Operations & Behavior Summary

* Traverse digits from right to left
* Handle carry propagation
* Modify digits in place
* Add new digit only when required

---

## Complexity

**Time Complexity:** O(n)
Here, n is the number of digits. In the worst case, I traverse the full array once.

**Space Complexity:** O(1)
No extra space is used except when adding a new digit at the front.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        for (int i = digits.size() - 1; i >= 0; i--) {
            digits[i]++;
            if (digits[i] < 10) {
                return digits;
            }
            digits[i] = 0;
        }
        digits.insert(digits.begin(), 1);
        return digits;
    }
};
```

---

### Java

```java
class Solution {
    public int[] plusOne(int[] digits) {
        for (int i = digits.length - 1; i >= 0; i--) {
            digits[i]++;
            if (digits[i] < 10) {
                return digits;
            }
            digits[i] = 0;
        }
        int[] result = new int[digits.length + 1];
        result[0] = 1;
        return result;
    }
}
```

---

### JavaScript

```javascript
var plusOne = function(digits) {
    for (let i = digits.length - 1; i >= 0; i--) {
        digits[i]++;
        if (digits[i] < 10) {
            return digits;
        }
        digits[i] = 0;
    }
    digits.unshift(1);
    return digits;
};
```

---

### Python3

```python
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        for i in range(len(digits) - 1, -1, -1):
            digits[i] += 1
            if digits[i] < 10:
                return digits
            digits[i] = 0
        return [1] + digits
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3)

1. I start from the last digit because addition always begins from the right.
2. I add 1 to the current digit.
3. If the value is less than 10, no carry is needed, so I return the result.
4. If the value becomes 10, I reset it to 0 and continue.
5. This process continues until carry stops.
6. If carry never stops, it means the number was all 9s.
7. In that case, I add 1 at the front of the array.

This logic works the same in all languages.

---

## Examples

Input: `[1,2,3]`
Output: `[1,2,4]`

Input: `[4,3,2,1]`
Output: `[4,3,2,2]`

Input: `[9,9,9]`
Output: `[1,0,0,0]`

---

## How to use / Run locally

1. Copy the code for your preferred language.
2. Paste it into the LeetCode editor or your local compiler.
3. Run the program with sample inputs.
4. Verify the output matches expected results.

---

## Notes & Optimizations

* No integer conversion is used, avoiding overflow issues.
* The solution works efficiently for very large numbers.
* In-place modification minimizes memory usage.
* Interview friendly and easy to explain.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
