# Maximize Consecutive Elements in an Array After Modification – LeetCode 3041 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations &amp; Behavior Summary](#operations--behavior-summary)
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
- [Notes &amp; Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an array of positive integers. You are allowed to increase any element by at most 1 (you can also leave it unchanged). After making these changes, you must pick as many elements as possible so that when you sort the chosen values they form a consecutive sequence. That means each next number is exactly one greater than the previous one, with no gaps and no duplicates.

The task is to return the size of the largest such consecutive group you can create.

This is a classic dynamic programming + sorting problem that appears in coding interviews and competitive programming contests under the name Maximize Consecutive Elements in an Array After Modification (LeetCode 3041).

## Constraints

- 1 &lt;= nums.length &lt;= 10^5
- 1 &lt;= nums[i] &lt;= 10^6

Because the array can be as large as 100 000 elements, any solution slower than O(n log n) will time out.

## Intuition

The first useful observation is that the order of the numbers does not matter for the final consecutive check. Only the values after the possible +1 modification matter. Therefore sorting the array immediately brings related numbers next to each other.

Once the numbers are sorted, each value can contribute either to its own position or to the next higher position. This suggests keeping track of the longest consecutive chain that ends at every possible number. A simple map (or dictionary) can store those lengths while we scan the sorted array from left to right.

## Approach

1. Sort the input array in non-decreasing order.
2. Create an empty map that will store the length of the longest consecutive sequence ending at a given number.
3. Walk through every number in the sorted array.
4. For the current number:
   - First try using it as (number + 1). Look up the best length that already ends at the original number and add 1. Store the result under the key (number + 1).
   - Then try keeping the number as-is. Look up the best length that ends at (number – 1) and add 1. Store the result under the key equal to the number.
5. After both updates, keep the maximum of the two newly written lengths as a candidate answer.
6. When the scan finishes, the maximum value seen is the answer.

The order of the two map updates is important: the +1 version must be written before the original version is overwritten.

## Data Structures Used

- Array / Vector / List – the original input that we sort in place.
- Hash Map (unordered_map in C++, HashMap in Java, Map in JavaScript/TypeScript, dict in Python, map in Go) – stores the longest consecutive length ending at each possible value. It is chosen because we need fast random access by value and the set of values is sparse.

## Operations &amp; Behavior Summary

- Sort the array so that smaller candidates appear before larger ones.
- For every number x:
  - Extend any chain that ends at x by turning the current element into x+1.
  - Extend any chain that ends at x-1 by keeping the current element as x.
- Continuously record the longest length discovered.
- Because the array is sorted, every possible predecessor has already been processed when we look it up.

## Complexity

| Metric            | Value          | Explanation                                                                 |
|-------------------|----------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(n log n)     | Dominated by the initial sort. The linear scan with map operations is average O(n). |
| Space Complexity  | O(n)           | The map may store an entry for almost every distinct value and its +1 neighbour. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int maxSelectedElements(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        unordered_map<int, int> dp;
        int ans = 0;
        for (int num : nums) {
            dp[num + 1] = dp[num] + 1;
            dp[num] = dp[num - 1] + 1;
            ans = max({ans, dp[num], dp[num + 1]});
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, Integer> dp = new HashMap<>();
        int ans = 0;
        for (int num : nums) {
            dp.put(num + 1, dp.getOrDefault(num, 0) + 1);
            dp.put(num, dp.getOrDefault(num - 1, 0) + 1);
            ans = Math.max(ans, Math.max(dp.get(num), dp.get(num + 1)));
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
var maxSelectedElements = function(nums) {
    nums.sort((a, b) => a - b);
    const dp = new Map();
    let ans = 0;
    for (const num of nums) {
        dp.set(num + 1, (dp.get(num) || 0) + 1);
        dp.set(num, (dp.get(num - 1) || 0) + 1);
        ans = Math.max(ans, dp.get(num), dp.get(num + 1));
    }
    return ans;
};
```

### TypeScript
```typescript
function maxSelectedElements(nums: number[]): number {
    nums.sort((a, b) => a - b);
    const dp = new Map<number, number>();
    let ans = 0;
    for (const num of nums) {
        dp.set(num + 1, (dp.get(num) || 0) + 1);
        dp.set(num, (dp.get(num - 1) || 0) + 1);
        ans = Math.max(ans, dp.get(num)!, dp.get(num + 1)!);
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        dp = {}
        ans = 0
        for num in nums:
            dp[num + 1] = dp.get(num, 0) + 1
            dp[num] = dp.get(num - 1, 0) + 1
            ans = max(ans, dp[num], dp[num + 1])
        return ans
```

### Go
```go
func maxSelectedElements(nums []int) int {
    sort.Ints(nums)
    dp := make(map[int]int)
    ans := 0
    for _, num := range nums {
        dp[num+1] = dp[num] + 1
        dp[num] = dp[num-1] + 1
        if dp[num] > ans {
            ans = dp[num]
        }
        if dp[num+1] > ans {
            ans = dp[num+1]
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six implementations follow exactly the same logic. The only differences are language-specific syntax for maps, sorting, and max calculations.

- We begin by sorting the array. This guarantees that when we process a number, every smaller or equal number has already been considered.
- We allocate a map whose keys are possible ending values and whose values are the longest consecutive lengths found so far.
- For each number we first compute the length we can achieve by increasing it. We take whatever length currently ends at the original number and add one, then store that result under the increased key. This must happen before the next line because the next line will change the entry for the original number.
- Immediately afterwards we compute the length obtained by leaving the number unchanged. We look up the length that ends one smaller and add one, writing the result under the original key.
- We compare both lengths with the global answer and keep the larger one.
- When the same number appears several times, the first occurrence creates a chain ending at that number. Later identical numbers can then attach their increased versions to that chain, which is why the update order is critical.
- Gaps in the sorted array automatically produce a length of zero for missing predecessors, so a new chain of length one starts cleanly.
- After the whole array has been processed the global answer holds the size of the longest consecutive sequence that can be formed.

Edge cases such as a single-element array or an array with large gaps are handled naturally by the map returning zero for missing keys.

## Examples

**Example 1**

Input: nums = [2, 1, 5, 1, 1]

After sorting: [1, 1, 1, 2, 5]

- First 1 creates length 1 ending at 1 and length 1 ending at 2.
- Second and third 1s extend the chain ending at 2.
- The 2 extends the chain ending at 1 into a chain of length 2 ending at 2, and also creates a chain ending at 3.
- The 5 starts a new short chain.
- Maximum length obtained is 3 (the sequence 1-2-3).

Output: 3

**Example 2**

Input: nums = [1, 4, 7, 10]

After sorting the numbers are far apart. Each number can only form a chain of length 1 because no predecessor exists within distance 1.

Output: 1

**Example 3**

Input: nums = [3, 3, 3, 4]

Possible consecutive sequences of length 3 can be formed (for example 3-4-5 by increasing two of the 3s).

Output: 3

## How to Use / Run Locally

**C++**
```
g++ -std=c++17 solution.cpp -o solution
./solution
```

**Java**
```
javac Solution.java
java Solution
```

**JavaScript**
```
node solution.js
```

**TypeScript**
```
tsc solution.ts
node solution.js
```

**Python3**
```
python3 solution.py
```

**Go**
```
go run solution.go
```

Replace the empty function bodies with the code from the sections above, then feed the sample inputs through standard input or hard-code them for a quick test.

## Notes &amp; Optimizations

- The map-based solution uses O(n) extra space. An O(1)-space version exists that only keeps the lengths of the two most recent possible endings, but it is more error-prone to implement.
- Because the values themselves go up to 10^6, a pure frequency array would waste memory; the sparse map is the practical choice.
- The algorithm never creates duplicate values in the final consecutive sequence because each map entry records a strictly increasing chain.
- If the problem later allowed decreasing by 1 as well, the same framework could be extended by adding a third possible ending, but that is outside the current statement.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
