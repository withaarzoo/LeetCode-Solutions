# LeetCode 3524. Find X Value of Array I Solution

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

This problem asks you to count the number of ways you can leave a contiguous subarray after removing a non-overlapping prefix and suffix from the given array. The remaining subarray must stay non-empty. For every possible remainder x from 0 to k-1, you need to count how many such subarrays have a product that leaves remainder x when divided by k.

In short, you are given an array of positive integers nums and a positive integer k. You must return an array of size k where the value at index x tells how many non-empty contiguous subarrays have product % k equal to x.

The operation of removing a prefix and a suffix is just another way of saying “pick any contiguous subarray”. Because the product can become very large and the array can be long, a smart way of tracking products modulo k is required.

## Constraints

- 1 <= nums[i] <= 10^9
- 1 <= nums.length <= 10^5
- 1 <= k <= 5

These constraints make a plain O(n²) solution too slow. The small value of k is the key that lets us use an efficient dynamic-programming approach.

## Intuition

The first thing I noticed is that removing a non-overlapping prefix and suffix simply leaves a contiguous subarray. So the real task is to count, for every possible remainder, how many subarrays have product % k equal to that remainder.

With n up to 1e5, checking every subarray directly is impossible. But k is at most 5. That means there are only a handful of possible remainders. I realized I could keep track of how many subarrays ending at the current index produce each remainder. When I add a new number, I only need to multiply those previous remainders by the new number (modulo k) and update the counts. This keeps everything linear in n and quadratic only in the tiny value of k.

## Approach

I maintain a small array called curr of size k. curr[r] stores how many subarrays that end at the previous position have product % k equal to r.

I also keep a result array of size k that will hold the final counts.

For every number in the array I do the following:

1. Reduce the current number modulo k. Call it m.
2. Create a fresh array next of size k.
3. The single-element subarray that contains only the current number contributes 1 to next[m].
4. For every remainder that already exists in curr, multiply it by m (modulo k) and add the corresponding count into the correct bucket of next.
5. Add every value from next into the matching index of the result array.
6. Replace curr with next so the next iteration can continue from the current ending position.

When the loop finishes, the result array contains the answer for every possible remainder.

## Data Structures Used

- Two fixed-size arrays (or vectors) of length k: one stores the counts for subarrays ending at the previous index, the other stores the counts for the current index.
- One result array of length k that accumulates the final answer.

Because k is at most 5, these arrays use almost no memory and give constant-time access. No maps, sets, or other heavy structures are needed.

## Operations & Behavior Summary

- Initialize result and curr arrays of size k to all zeros.
- For each number in nums:
  - Compute m = number % k.
  - Create a new next array of size k filled with zeros.
  - Set next[m] = 1 (the single-element subarray).
  - For every previous remainder prev that has a positive count:
    - Compute new remainder = (prev * m) % k.
    - Add the previous count to next[new remainder].
  - Add every entry of next into the corresponding entry of result.
  - Set curr = next.
- Return the result array.

This process examines every possible subarray exactly once—when its rightmost element is processed—and records its product modulo k.

## Complexity

| Type              | Complexity     | Explanation |
|-------------------|----------------|-------------|
| Time Complexity   | O(n * k²)      | For each of the n elements we examine up to k previous remainders and write into up to k new remainders. Since k ≤ 5 the k² factor is tiny. |
| Space Complexity  | O(k)           | Only a few arrays of length k are used. No extra space that grows with n is required. |

## Multi-language Solutions

### C++
```cpp
class Solution { 
public: 
    vector<long long> resultArray(vector<int>& nums, int k) { 
        vector<long long> result(k, 0);
        vector<long long> curr(k, 0);
        for (int num : nums) {
            int m = num % k;
            vector<long long> next(k, 0);
            next[m] = 1;
            for (int prev = 0; prev < k; prev++) {
                if (curr[prev] > 0) {
                    int nr = (int)((prev * 1LL * m) % k);
                    next[nr] += curr[prev];
                }
            }
            for (int r = 0; r < k; r++) {
                result[r] += next[r];
            }
            curr = next;
        }
        return result;
    } 
};
```

### Java
```java
class Solution { 
    public long[] resultArray(int[] nums, int k) { 
        long[] result = new long[k];
        long[] curr = new long[k];
        for (int num : nums) {
            int m = num % k;
            long[] next = new long[k];
            next[m] = 1;
            for (int prev = 0; prev < k; prev++) {
                if (curr[prev] > 0) {
                    int nr = (int)((prev * 1L * m) % k);
                    next[nr] += curr[prev];
                }
            }
            for (int r = 0; r < k; r++) {
                result[r] += next[r];
            }
            curr = next;
        }
        return result;
    } 
}
```

### JavaScript
```javascript
/** 
 * @param {number[]} nums 
 * @param {number} k 
 * @return {number[]} 
 */ 
var resultArray = function(nums, k) { 
    let result = new Array(k).fill(0);
    let curr = new Array(k).fill(0);
    for (let num of nums) {
        let m = num % k;
        let next = new Array(k).fill(0);
        next[m] = 1;
        for (let prev = 0; prev < k; prev++) {
            if (curr[prev] > 0) {
                let nr = (prev * m) % k;
                next[nr] += curr[prev];
            }
        }
        for (let r = 0; r < k; r++) {
            result[r] += next[r];
        }
        curr = next;
    }
    return result;
};
```

### TypeScript
```typescript
function resultArray(nums: number[], k: number): number[] {
    let result = new Array(k).fill(0);
    let curr = new Array(k).fill(0);
    for (let num of nums) {
        let m = num % k;
        let next = new Array(k).fill(0);
        next[m] = 1;
        for (let prev = 0; prev < k; prev++) {
            if (curr[prev] > 0) {
                let nr = (prev * m) % k;
                next[nr] += curr[prev];
            }
        }
        for (let r = 0; r < k; r++) {
            result[r] += next[r];
        }
        curr = next;
    }
    return result;
};
```

### Python3
```python
class Solution:
    def resultArray(self, nums: List[int], k: int) -> List[int]:
        result = [0] * k
        curr = [0] * k
        for num in nums:
            m = num % k
            next = [0] * k
            next[m] = 1
            for prev in range(k):
                if curr[prev] > 0:
                    nr = (prev * m) % k
                    next[nr] += curr[prev]
            for r in range(k):
                result[r] += next[r]
            curr = next
        return result
```

### Go
```go
func resultArray(nums []int, k int) []int64 {
    result := make([]int64, k)
    curr := make([]int64, k)
    for _, num := range nums {
        m := num % k
        next := make([]int64, k)
        next[m] = 1
        for prev := 0; prev < k; prev++ {
            if curr[prev] > 0 {
                nr := (prev * m) % k
                next[nr] += curr[prev]
            }
        }
        for r := 0; r < k; r++ {
            result[r] += next[r]
        }
        curr = next
    }
    return result
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. The only differences are syntax details such as how arrays are declared, how long integers are written, and how modulo is computed.

We start by creating a result array of size k filled with zeros. This will hold the final counts for every remainder from 0 to k-1.

We also create a curr array of size k filled with zeros. At the beginning no subarrays have been seen, so every count is zero.

Then we walk through the input array one number at a time. For the current number we first reduce it modulo k. Call the result m. From now on we only work with m.

We create a brand-new array called next, also of size k and initially zero. The first action is to set next[m] = 1. This records the subarray that consists of only the current element.

Next we look at every possible remainder that already exists in curr. If a remainder prev has a positive count, it means there are that many subarrays ending at the previous index whose product was congruent to prev. Multiplying each of those products by the new number m produces a new product congruent to (prev * m) % k. We therefore add the previous count into the correct bucket of next.

After the inner loop finishes, next contains the exact number of subarrays that end at the current index for every possible remainder. We simply add every entry of next into the matching entry of the result array. That is how the global answer grows.

Finally we assign next to curr. The next iteration will treat the current ending position as the new “previous” position.

When the outer loop ends, every non-empty contiguous subarray has been processed exactly once—at the moment its rightmost element was examined—and its product modulo k has been recorded in the result array. We return that array.

Edge cases are handled naturally: a single-element array works because of the next[m] = 1 line; the whole array is covered when both the prefix and suffix removed are empty; and because we always work with remainders, overflow never becomes an issue as long as intermediate multiplications are performed carefully (using 64-bit integers in C++ and Java).

## Examples

**Example 1**

Input: nums = [1, 2, 3, 4, 5], k = 3  
Output: [9, 2, 4]

Trace:  
- When we reach 1, only the subarray [1] is recorded (remainder 1).  
- When we reach 2, we get [2] (remainder 2) and [1,2] (remainder 2).  
- When we reach 3, we get [3] (remainder 0), [2,3] (remainder 0), [1,2,3] (remainder 0).  
- Continuing this way for 4 and 5 produces the final counts 9, 2 and 4 for remainders 0, 1 and 2.

**Example 2**

Input: nums = [1, 2, 4, 8, 16, 32], k = 4  
Output: [18, 1, 2, 0]

Trace: Most subarrays contain a factor of 4 or higher powers of 2, so they become 0 modulo 4. Only a few short prefixes produce the other remainders, matching the output.

**Example 3**

Input: nums = [1, 1, 2, 1, 1], k = 2  
Output: [9, 6]

Trace: The number 2 is the only even number. Subarrays that include it become even (remainder 0). Subarrays that avoid it stay odd (remainder 1). Counting all of them gives the reported answer.

## How to Use / Run Locally

1. Copy the code for the language you want into a file (for example solution.cpp, Solution.java, solution.js, etc.).
2. Make sure you have the corresponding compiler or runtime installed.
3. For C++:  
   g++ -std=c++17 solution.cpp -o solution && ./solution
4. For Java:  
   javac Solution.java && java Solution
5. For JavaScript:  
   node solution.js
6. For TypeScript:  
   tsc solution.ts && node solution.js
7. For Python3:  
   python3 solution.py
8. For Go:  
   go run solution.go

You will need to add a small main function or test harness that creates the input arrays and prints the returned result, because the provided snippets contain only the core function.

## Notes & Optimizations

- The solution already runs in the best practical complexity given the constraints. Because k is tiny, even the k² factor is negligible.
- An alternative O(n * k) approach is possible by carefully updating only the necessary remainders, but the current version is simpler and still fast enough.
- Watch out for integer overflow when multiplying remainders in languages that do not promote automatically. Using 64-bit integers for the multiplication step avoids the problem.
- Empty arrays never appear because the problem guarantees at least one element and the remaining subarray must stay non-empty.
- If k were larger (for example 10^9), a completely different technique would be required; the current method relies heavily on k being very small.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)