# Minimum Score by Changing Two Elements - LeetCode 2567 Solution

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

You are given an integer array called nums. The low score of the array is the smallest absolute difference between any two numbers. The high score is the largest absolute difference between any two numbers. The score of the array is simply the sum of the high score and the low score.

Your task is to change exactly two elements of the array (you can change them to any integer values) so that the final score becomes as small as possible. Return that minimum possible score.

This is a classic array manipulation problem that appears in coding interviews and online judges under the name Minimum Score by Changing Two Elements (LeetCode 2567).

## Constraints

- 3 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9

The upper limit on the array size means any solution slower than O(n log n) will likely time out, so sorting-based ideas are preferred.

## Intuition

When I first read the problem I noticed that the score is high + low. The low score can always be forced to zero because I am allowed to change two numbers and I can set both of them to the same value. That means the real challenge is only to make the high score (the overall range of the array) as small as possible.

After I change two numbers, n-2 original numbers stay untouched. The smallest range I can keep is therefore the tightest window of n-2 consecutive values in the sorted array. Looking only at the extremes gives me exactly three candidates, and the answer is the smallest of those three ranges.

## Approach

1. Sort the array in non-decreasing order.
2. After sorting, the three possible ways to discard two extreme values are:
   - Discard the two smallest numbers → remaining range is nums[n-1] - nums[2]
   - Discard the two largest numbers → remaining range is nums[n-3] - nums[0]
   - Discard one smallest and one largest → remaining range is nums[n-2] - nums[1]
3. Return the minimum of those three values. Because the two changed numbers can be placed anywhere inside that range (and can even be equal), the low score becomes zero and the total score equals the chosen range.

This greedy observation after sorting solves the Minimum Score by Changing Two Elements problem cleanly.

## Data Structures Used

- The input array itself. We sort it in place (or create a sorted copy depending on language constraints). No extra heavy data structures are required; the solution stays memory-light.

## Operations & Behavior Summary

- Sort the entire array once.
- Compute three simple differences using the ends of the sorted array.
- Take the minimum of those three differences.
- That minimum is the answer.

The algorithm never builds additional maps, sets, or heaps; everything is done with direct index arithmetic after the sort.

## Complexity

| Complexity       | Value      | Explanation                                                                 |
|------------------|------------|-----------------------------------------------------------------------------|
| Time Complexity  | O(n log n) | Dominated by the sorting step. n is the length of nums.                     |
| Space Complexity | O(1) or O(log n) | In-place sorting uses constant extra memory (or the recursion stack of the sort). |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minimizeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        int n = nums.size();
        return min({nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1]});
    }
};
```

### Java
```java
class Solution {
    public int minimizeSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        return Math.min(Math.min(nums[n-1] - nums[2], nums[n-3] - nums[0]), nums[n-2] - nums[1]);
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minimizeSum = function(nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    return Math.min(nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1]);
};
```

### TypeScript
```typescript
function minimizeSum(nums: number[]): number {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    return Math.min(nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1]);
};
```

### Python3
```python
class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        nums.sort()
        n = len(nums)
        return min(nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1])
```

### Go
```go
func minimizeSum(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    return min(nums[n-1]-nums[2], min(nums[n-3]-nums[0], nums[n-2]-nums[1]))
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages, so the reasoning is the same.

First the array is sorted. After sorting, nums[0] holds the smallest value and nums[n-1] holds the largest. Every later calculation depends on this ordered view.

I then look at the three candidate ranges:

- nums[n-1] - nums[2] keeps everything except the two leftmost (smallest) elements.
- nums[n-3] - nums[0] keeps everything except the two rightmost (largest) elements.
- nums[n-2] - nums[1] keeps the middle portion after removing one value from each end.

Any other pair of removals would leave a wider or equal range, so these three cover the optimum. Taking the minimum of them gives the smallest possible high score. Because I can still change the two discarded positions to any values that lie inside the chosen range, I can make two numbers equal and force the low score to zero. Therefore the returned value is already the minimum score.

Edge cases are handled automatically: when n = 3 the three expressions all evaluate to zero, which matches the expected answer of making every remaining number identical.

## Examples

**Example 1**  
Input: nums = [1, 4, 7, 8, 5]  
Sorted: [1, 4, 5, 7, 8]  
Three ranges: 8-5 = 3, 5-1 = 4, 7-4 = 3  
Minimum = 3  
Output: 3

**Example 2**  
Input: nums = [1, 4, 3]  
Sorted: [1, 3, 4]  
Three ranges all become 0  
Output: 0

**Example 3**  
Input: nums = [10, 20, 30, 40, 50, 60]  
Sorted: [10, 20, 30, 40, 50, 60]  
Three ranges: 60-30 = 30, 40-10 = 30, 50-20 = 30  
Minimum = 30  
Output: 30

## How to Use / Run Locally

**C++**  
Compile with g++ -std=c++17 and run the executable. Paste the function into a main that reads the array and prints the result.

**Java**  
Place the class in a file named Solution.java, compile with javac, and run with java. Feed the array through a simple driver.

**JavaScript**  
Copy the function into a Node.js file and call it with a test array. Run with node filename.js.

**TypeScript**  
Compile with tsc then run the generated JavaScript, or use ts-node for direct execution.

**Python3**  
Save the class in a .py file and call the method from a short test harness. Run with python3.

**Go**  
Put the function in a package main file, add a main that constructs the slice and prints the answer, then run with go run.

## Notes & Optimizations

The solution is already optimal for both time and space under the given constraints. An alternative would be to try every possible pair of indices to change, but that would be O(n²) and fail for n = 10^5. Sorting once and examining only three candidates is the practical way to solve Minimum Score by Changing Two Elements.

If the problem ever allowed changing only one element or zero elements, the same idea extends by simply considering more windows, but the current statement fixes the number of changes at two.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
