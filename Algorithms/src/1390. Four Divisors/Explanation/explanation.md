# LeetCode 1390 – Four Divisors

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

Given an integer array `nums`, I have to find all numbers that have **exactly four divisors**.
For each such number, I calculate the **sum of its divisors**.
Finally, I return the **total sum** of all those divisor sums.

If no number has exactly four divisors, I return `0`.

---

## Constraints

* 1 ≤ nums.length ≤ 10⁴
* 1 ≤ nums[i] ≤ 10⁵

---

## Intuition

When I first read the problem, I focused on one key question:
what kind of numbers have **exactly four divisors**?

After thinking and testing small numbers, I noticed:

* Prime numbers have only 2 divisors
* Perfect squares of primes have 3 divisors
* Numbers like `21 = 3 × 7` have exactly 4 divisors

So instead of checking all divisors blindly, I decided to **count divisors smartly** and stop early whenever the count becomes more than 4.

This makes the solution fast and efficient.

---

## Approach

1. I iterate through every number in the array.
2. For each number, I loop from `1` to `sqrt(number)`.
3. If `d` divides the number, then:

   * `d` is a divisor
   * `number / d` is also a divisor
4. I keep track of:

   * total divisor count
   * sum of divisors
5. If divisor count becomes more than 4, I stop checking that number.
6. If the final divisor count is exactly 4, I add its divisor sum to the answer.
7. After processing all numbers, I return the final sum.

---

## Data Structures Used

* Only basic variables
* No extra arrays, maps, or sets

---

## Operations & Behavior Summary

* Uses divisor pairs to reduce checks
* Stops early when divisor count exceeds 4
* Avoids unnecessary calculations
* Works efficiently within given constraints

---

## Complexity

**Time Complexity:**
O(n × √k)
where

* n = number of elements in the array
* k = maximum value in the array

**Space Complexity:**
O(1)
No extra memory used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int sumFourDivisors(vector<int>& nums) {
        int total = 0;
        for (int num : nums) {
            int cnt = 0, sum = 0;
            for (int d = 1; d * d <= num; d++) {
                if (num % d == 0) {
                    int other = num / d;
                    cnt++; sum += d;
                    if (other != d) {
                        cnt++; sum += other;
                    }
                    if (cnt > 4) break;
                }
            }
            if (cnt == 4) total += sum;
        }
        return total;
    }
};
```

### Java

```java
class Solution {
    public int sumFourDivisors(int[] nums) {
        int total = 0;
        for (int num : nums) {
            int cnt = 0, sum = 0;
            for (int d = 1; d * d <= num; d++) {
                if (num % d == 0) {
                    int other = num / d;
                    cnt++; sum += d;
                    if (other != d) {
                        cnt++; sum += other;
                    }
                    if (cnt > 4) break;
                }
            }
            if (cnt == 4) total += sum;
        }
        return total;
    }
}
```

### JavaScript

```javascript
var sumFourDivisors = function(nums) {
    let total = 0;
    for (let num of nums) {
        let cnt = 0, sum = 0;
        for (let d = 1; d * d <= num; d++) {
            if (num % d === 0) {
                let other = num / d;
                cnt++; sum += d;
                if (other !== d) {
                    cnt++; sum += other;
                }
                if (cnt > 4) break;
            }
        }
        if (cnt === 4) total += sum;
    }
    return total;
};
```

### Python3

```python
class Solution:
    def sumFourDivisors(self, nums):
        total = 0
        for num in nums:
            cnt = 0
            div_sum = 0
            d = 1
            while d * d <= num:
                if num % d == 0:
                    other = num // d
                    cnt += 1
                    div_sum += d
                    if other != d:
                        cnt += 1
                        div_sum += other
                    if cnt > 4:
                        break
                d += 1
            if cnt == 4:
                total += div_sum
        return total
```

### Go

```go
func sumFourDivisors(nums []int) int {
    total := 0
    for _, num := range nums {
        cnt, sum := 0, 0
        for d := 1; d*d <= num; d++ {
            if num%d == 0 {
                other := num / d
                cnt++
                sum += d
                if other != d {
                    cnt++
                    sum += other
                }
                if cnt > 4 {
                    break
                }
            }
        }
        if cnt == 4 {
            total += sum
        }
    }
    return total
}
```

---

## Step-by-step Detailed Explanation

* I start with one number.
* I check divisors only up to its square root.
* Every divisor gives me two values.
* I count and sum them carefully.
* If count crosses 4, I stop immediately.
* If count is exactly 4, I add the sum to the result.
* I repeat this for all numbers.

---

## Examples

Input
`[21, 4, 7]`

Explanation

* 21 → divisors: 1, 3, 7, 21 → sum = 32
* 4 → only 3 divisors
* 7 → only 2 divisors

Output
`32`

---

## How to use / Run locally

1. Copy the code for your preferred language
2. Paste it into LeetCode or your local IDE
3. Run with test cases
4. Verify output

---

## Notes & Optimizations

* Early stopping improves performance
* No extra memory used
* Works well within constraints
* Interview friendly and easy to explain

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
