# Maximum Subarray With Equal Products - LeetCode 3411 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an array of positive integers called `nums`.  

A subarray is called product equivalent when the product of all its elements equals the LCM of those elements multiplied by their GCD. In simple words:

prod(subarray) == lcm(subarray) * gcd(subarray)

Your task is to return the length of the longest product equivalent subarray that exists inside `nums`.

This is a classic maximum subarray length problem that mixes product, GCD and LCM calculations. It appears in Weekly Contest 431 and is tagged under Array, Math, Sliding Window, Enumeration and Number Theory.

## Constraints

- 2 <= nums.length <= 100
- 1 <= nums[i] <= 10

The small limits (n ≤ 100 and values ≤ 10) make an O(n²) enumeration completely safe and practical.

## Intuition

The first thing I noticed was the tiny constraints. With only 100 elements I can afford to look at every possible subarray.  

I also remembered a basic number-theory fact: for any two positive integers a and b we always have  

a * b == lcm(a, b) * gcd(a, b)

So every length-2 subarray is automatically valid. That means the answer is never smaller than 2.  

The real work is checking longer subarrays. Because the numbers are at most 10, the LCM of any group of them stays tiny (at most 2520). Therefore the right-hand side of the equality (LCM × GCD) can never grow very large. If the running product becomes bigger than that range, the equality is impossible and I can stop extending that particular subarray.

## Approach

I use a simple double loop that tries every left endpoint and every right endpoint:

1. For each possible starting index i I keep three running values: the product so far, the GCD so far and the LCM so far.
2. I extend the right index j one step at a time, updating the three values with the new number nums[j].
3. After each update I check whether product equals LCM × GCD. If it does, I record the current length.
4. If the product is about to overflow a 64-bit integer I break, because the LCM × GCD side stays small and the equality can never hold any longer.
5. After examining every subarray I return the largest length that satisfied the condition.

This brute-force enumeration of all subarrays is both correct and fast enough under the given constraints.

## Data Structures Used

No advanced data structures are required.  

I only maintain a handful of scalar variables:

- a 64-bit integer for the running product  
- an integer for the running GCD  
- a 64-bit integer for the running LCM  
- an integer that stores the best length found so far  

These variables are reset for every new left endpoint, so the extra memory stays constant.

## Operations & Behavior Summary

- Outer loop picks every possible left endpoint i.  
- For each i, three running values (product, GCD, LCM) are initialised with nums[i].  
- Inner loop walks the right endpoint j from i to the end of the array.  
- At each j the three values are updated with nums[j] using the classic Euclidean algorithm for GCD and the formula LCM(a,b) = a / GCD(a,b) * b.  
- Overflow checks protect the 64-bit product and LCM.  
- After a successful update the equality product == LCM × GCD is tested and the answer is updated if needed.  
- When the inner loop finishes (or breaks on overflow) the outer loop moves to the next left endpoint.  
- Finally the largest valid length is returned.

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time       | O(n² log A) | Two nested loops examine every subarray. Each GCD or LCM step costs a few Euclidean iterations that are logarithmic in the maximum value A (A ≤ 10). |
| Space      | O(1)        | Only a constant number of scalar variables are used. No extra arrays or maps are allocated. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maxLength(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            long long prod = 1;
            int g = 0;
            long long lc = 1;
            for (int j = i; j < n; j++) {
                int x = nums[j];
                if (g == 0) g = x;
                else {
                    int a = g, b = x;
                    while (b) {
                        int t = b;
                        b = a % b;
                        a = t;
                    }
                    g = a;
                }
                long long gg = lc;
                int b = x;
                while (b) {
                    long long t = b;
                    b = gg % b;
                    gg = t;
                }
                if (lc / gg > LLONG_MAX / x) break;
                lc = lc / gg * x;
                if (prod > LLONG_MAX / x) break;
                prod *= x;
                if (prod == lc * g) ans = max(ans, j - i + 1);
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int maxLength(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            long prod = 1;
            int g = 0;
            long lc = 1;
            for (int j = i; j < n; j++) {
                int x = nums[j];
                if (g == 0) g = x;
                else {
                    int a = g, b = x;
                    while (b != 0) {
                        int t = b;
                        b = a % b;
                        a = t;
                    }
                    g = a;
                }
                long gg = lc;
                int b = x;
                while (b != 0) {
                    long t = b;
                    b = (int)(gg % b);
                    gg = t;
                }
                if (lc / gg > Long.MAX_VALUE / x) break;
                lc = lc / gg * x;
                if (prod > Long.MAX_VALUE / x) break;
                prod *= x;
                if (prod == lc * g) ans = Math.max(ans, j - i + 1);
            }
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxLength = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 0; i < n; i++) {
        let prod = 1;
        let g = 0;
        let lc = 1;
        for (let j = i; j < n; j++) {
            const x = nums[j];
            if (g === 0) g = x;
            else {
                let a = g, b = x;
                while (b) {
                    const t = b;
                    b = a % b;
                    a = t;
                }
                g = a;
            }
            let gg = lc;
            let b = x;
            while (b) {
                const t = b;
                b = gg % b;
                gg = t;
            }
            if (lc / gg > Number.MAX_SAFE_INTEGER / x) break;
            lc = lc / gg * x;
            if (prod > Number.MAX_SAFE_INTEGER / x) break;
            prod *= x;
            if (prod === lc * g) ans = Math.max(ans, j - i + 1);
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function maxLength(nums: number[]): number {
    const n = nums.length;
    let ans = 0;
    for (let i = 0; i < n; i++) {
        let prod = 1;
        let g = 0;
        let lc = 1;
        for (let j = i; j < n; j++) {
            const x = nums[j];
            if (g === 0) g = x;
            else {
                let a = g, b = x;
                while (b) {
                    const t = b;
                    b = a % b;
                    a = t;
                }
                g = a;
            }
            let gg = lc;
            let b = x;
            while (b) {
                const t = b;
                b = gg % b;
                gg = t;
            }
            if (lc / gg > Number.MAX_SAFE_INTEGER / x) break;
            lc = lc / gg * x;
            if (prod > Number.MAX_SAFE_INTEGER / x) break;
            prod *= x;
            if (prod === lc * g) ans = Math.max(ans, j - i + 1);
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def maxLength(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        for i in range(n):
            prod = 1
            g = 0
            lc = 1
            for j in range(i, n):
                x = nums[j]
                if g == 0:
                    g = x
                else:
                    a, b = g, x
                    while b:
                        a, b = b, a % b
                    g = a
                gg, b = lc, x
                while b:
                    gg, b = b, gg % b
                if lc // gg > (10**18) // x:
                    break
                lc = lc // gg * x
                if prod > (10**18) // x:
                    break
                prod *= x
                if prod == lc * g:
                    ans = max(ans, j - i + 1)
        return ans
```

### Go
```go
func maxLength(nums []int) int {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        prod := int64(1)
        g := 0
        lc := int64(1)
        for j := i; j < n; j++ {
            x := nums[j]
            if g == 0 {
                g = x
            } else {
                a, b := g, x
                for b != 0 {
                    a, b = b, a%b
                }
                g = a
            }
            gg := lc
            b := int64(x)
            for b != 0 {
                gg, b = b, gg%b
            }
            if lc/gg > (1<<62)/int64(x) {
                break
            }
            lc = lc / gg * int64(x)
            if prod > (1<<62)/int64(x) {
                break
            }
            prod *= int64(x)
            if prod == lc*int64(g) {
                if j-i+1 > ans {
                    ans = j - i + 1
                }
            }
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages; only the syntax for 64-bit integers and integer division changes.

1. Record the length of the array as n and set the answer variable to 0.  
2. The outer loop runs from i = 0 to n-1. This is the left endpoint of the current candidate subarray.  
3. Inside the outer loop three variables are reset:  
   - prod starts at 1 (it will be multiplied by the first element on the first iteration),  
   - g starts at 0 (a sentinel that means “no GCD yet”),  
   - lc starts at 1 (it will become the first element on the first iteration).  
4. The inner loop runs from j = i to n-1. This is the right endpoint.  
5. Take the next element x = nums[j].  
6. Update the GCD:  
   - if this is the first element of the subarray, simply set g = x,  
   - otherwise run the Euclidean algorithm (repeated remainder) until the remainder becomes zero. The non-zero value left is the new GCD.  
7. Update the LCM:  
   - compute GCD(lc, x) with the same Euclidean loop,  
   - check that lc / thatGCD * x will not overflow a 64-bit integer,  
   - if it would overflow, break out of the inner loop,  
   - otherwise set lc to the new LCM value.  
8. Update the product:  
   - check that multiplying prod by x will not overflow,  
   - if it would, break,  
   - otherwise multiply.  
9. After the three values are safe, test whether prod equals lc * g.  
   - if the equality holds, update the answer with the length j-i+1.  
10. When the inner loop ends (either by reaching the end of the array or by an overflow break) the outer loop advances to the next left endpoint.  
11. After both loops finish, return the largest length that ever satisfied the equality.

Because every pair of numbers always satisfies the product-equals-LCM-times-GCD rule, the answer is guaranteed to be at least 2. The overflow guards simply stop us from wasting time on subarrays that can never be valid.

## Examples

**Example 1**  
Input: nums = [1,2,1,2,1,1,1]  
Output: 5  

One valid subarray of length 5 is [1,2,1,1,1].  
Product = 2, GCD = 1, LCM = 2 → 2 == 2 * 1, so the condition holds.  
No longer product-equivalent subarray exists, therefore the answer is 5.

**Example 2**  
Input: nums = [2,3,4,5,6]  
Output: 3  

The subarray [3,4,5] works: product = 60, GCD = 1, LCM = 60.  
Several other length-3 subarrays fail the test, and no length-4 or length-5 subarray works, so the maximum is 3.

**Example 3**  
Input: nums = [1,2,3,1,4,5,1]  
Output: 5  

A length-5 window that includes enough 1’s keeps the product small enough to equal LCM × GCD. The algorithm finds this window while scanning all possible left and right endpoints.

## How to Use / Run Locally

**C++**  
```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

**Java**  
```bash
javac Solution.java
java Solution
```

**JavaScript**  
```bash
node solution.js
```

**TypeScript**  
```bash
tsc solution.ts
node solution.js
```

**Python3**  
```bash
python3 solution.py
```

**Go**  
```bash
go run solution.go
```

Replace the empty function bodies with the code you intend to test, then feed sample inputs through standard input or hard-code them inside a small main function.

## Notes & Optimizations

- Length-2 subarrays are always valid, so the answer is never smaller than 2.  
- Because nums[i] ≤ 10 the LCM of any subarray is at most 2520; this fact lets us safely stop multiplying once the product exceeds a few tens of thousands.  
- An O(n²) scan is optimal enough under the given constraints; a sliding-window approach is possible but not necessary and harder to reason about because the “goodness” property is not monotonic.  
- Edge case: when the array contains many 1’s the product stays small for longer, which is why longer valid windows appear in the sample inputs that contain 1.  
- The same algorithm works unchanged in all six languages; only the concrete syntax for 64-bit arithmetic differs.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
