# LeetCode 2089. Find Target Indices After Sorting Array - Optimal O(n) Solution

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

You are given a 0-indexed integer array called `nums` and an integer `target`.  

A target index is any index `i` where `nums[i]` equals `target` after the array has been sorted in non-decreasing order.  

Your task is to return a list of all such target indices. The list must be sorted in increasing order. If no target indices exist, return an empty list.

This is a classic LeetCode array problem that tests whether you can find the positions of a value after sorting without actually performing a full sort every time.

## Constraints

- 1 <= nums.length <= 100
- 1 <= nums[i], target <= 100

## Intuition

When I first read the problem, my immediate thought was to sort the array and then scan for the target. That works, but it felt a bit heavy for such small constraints.

I quickly realized I do not need the fully sorted array at all. I only need to know where the target values would land after sorting.  

If I count how many numbers are smaller than the target, I instantly know the starting index of the target block. Counting how many numbers equal the target tells me how long that block is. That single observation turns an O(n log n) sorting solution into a clean O(n) linear scan.

## Approach

I make one pass over the array and maintain two simple counters:

1. `less` – how many elements are strictly smaller than `target`
2. `equal` – how many elements are exactly equal to `target`

After the pass finishes, the target values occupy a continuous range of indices starting at `less` and ending at `less + equal - 1`.  

I then generate every index in that range and return them as the answer. If `equal` is zero, the result is simply an empty list.

This approach never sorts the array and never uses extra memory beyond a couple of integer variables.

## Data Structures Used

- Two integer counters (`less` and `equal`) – constant extra space, extremely fast to update.
- A dynamic list / vector / array for the result – required only to store the final indices. Its size is at most the number of times the target appears.

No hash maps, no priority queues, and no temporary sorted copies are needed.

## Operations & Behavior Summary

1. Initialize `less = 0` and `equal = 0`.
2. Walk through every element `x` in `nums`:
   - If `x < target`, increment `less`.
   - Else if `x == target`, increment `equal`.
3. Create an empty result list.
4. For `i` from 0 to `equal - 1`, append the value `less + i` to the result.
5. Return the result list.

The algorithm is essentially a counting-based range generator that reconstructs the positions the target would occupy after a stable non-decreasing sort.

## Complexity

| Complexity Type | Value | Explanation |
|-----------------|-------|-------------|
| Time Complexity | O(n) | A single linear scan over the array of length n. Generating the answer takes O(k) time where k is the number of target occurrences (k ≤ n). |
| Space Complexity | O(1) extra | Only a few integer variables are used. The output list is not counted as extra space. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<int> targetIndices(vector<int>& nums, int target) {
        int less = 0, equal = 0;
        for (int x : nums) {
            if (x < target) less++;
            else if (x == target) equal++;
        }
        vector<int> ans;
        for (int i = 0; i < equal; i++) {
            ans.push_back(less + i);
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public List<Integer> targetIndices(int[] nums, int target) {
        int less = 0, equal = 0;
        for (int x : nums) {
            if (x < target) less++;
            else if (x == target) equal++;
        }
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < equal; i++) {
            ans.add(less + i);
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var targetIndices = function(nums, target) {
    let less = 0, equal = 0;
    for (const x of nums) {
        if (x < target) less++;
        else if (x === target) equal++;
    }
    const ans = [];
    for (let i = 0; i < equal; i++) {
        ans.push(less + i);
    }
    return ans;
};
```

### TypeScript
```typescript
function targetIndices(nums: number[], target: number): number[] {
    let less = 0, equal = 0;
    for (const x of nums) {
        if (x < target) less++;
        else if (x === target) equal++;
    }
    const ans: number[] = [];
    for (let i = 0; i < equal; i++) {
        ans.push(less + i);
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        less = equal = 0
        for x in nums:
            if x < target:
                less += 1
            elif x == target:
                equal += 1
        return list(range(less, less + equal))
```

### Go
```go
func targetIndices(nums []int, target int) []int {
    less, equal := 0, 0
    for _, x := range nums {
        if x < target {
            less++
        } else if x == target {
            equal++
        }
    }
    ans := make([]int, equal)
    for i := 0; i < equal; i++ {
        ans[i] = less + i
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

Across all six languages the core logic stays identical. Only the syntax for loops, lists, and type declarations changes.

I start by declaring the two counters. In C++, Java, JavaScript, TypeScript and Go I use two separate integer variables. In Python I can write `less = equal = 0` for brevity.

Next comes the single traversal. Every language simply iterates over the input array once. Inside the loop I compare the current element against the target and update the appropriate counter. Numbers larger than the target are ignored because they will appear after the target block in the sorted order and do not affect the indices we care about.

After the loop I know the exact range of indices that will contain the target. I create a result container (vector in C++, ArrayList in Java, array in JavaScript/TypeScript/Go, list in Python) and fill it with the consecutive integers from `less` to `less + equal - 1`.

Edge cases are handled automatically:
- If the target never appears, `equal` stays 0 and the result remains empty.
- If every element equals the target, `less` is 0 and the result contains every index from 0 to n-1.
- Duplicate values are correctly counted because we only care about the final positions after sorting, not the original indices.

The same reasoning applies in every language; only the surface syntax differs.

## Examples

**Example 1**  
Input: nums = [1,2,5,2,3], target = 2  
After a conceptual sort the array becomes [1,2,2,3,5].  
Numbers smaller than 2: one (the value 1) → less = 1  
Numbers equal to 2: two → equal = 2  
Result indices: 1 and 2  
Output: [1,2]

**Example 2**  
Input: nums = [1,2,5,2,3], target = 3  
Numbers smaller than 3: three (1,2,2) → less = 3  
Numbers equal to 3: one → equal = 1  
Result index: 3  
Output: [3]

**Example 3**  
Input: nums = [1,2,5,2,3], target = 5  
Numbers smaller than 5: four → less = 4  
Numbers equal to 5: one → equal = 1  
Result index: 4  
Output: [4]

## How to Use / Run Locally

**C++**  
1. Copy the solution into a file named `solution.cpp`.  
2. Compile: `g++ -std=c++17 solution.cpp -o solution`  
3. Run: `./solution` (you will need to add a small main function that reads input and prints the result).

**Java**  
1. Place the code inside a file `Solution.java`.  
2. Compile: `javac Solution.java`  
3. Run with a test harness that creates a `Solution` object and calls `targetIndices`.

**JavaScript**  
1. Save the function in `solution.js`.  
2. Run with Node: `node solution.js` after adding a few console.log test cases.

**TypeScript**  
1. Save as `solution.ts`.  
2. Compile with `tsc solution.ts` then run the generated JavaScript, or use `ts-node solution.ts`.

**Python3**  
1. Save the class in `solution.py`.  
2. Run: `python3 solution.py` after adding a few print statements that instantiate the class and call the method.

**Go**  
1. Save the function in `solution.go`.  
2. Run: `go run solution.go` (add a main function that tests a few cases).

In every language you can also paste the code directly into the LeetCode editor and run the built-in test cases.

## Notes & Optimizations

Because the constraints are tiny (n ≤ 100), even a full sort would pass easily. The counting approach is still preferable because it is faster, uses less memory, and teaches a useful technique: deriving sorted positions from frequency information.

Alternative approaches include:
- Sorting a copy of the array and then collecting matching indices (O(n log n) time).
- Using binary search on the sorted array to find the first and last occurrence of the target (still requires the sort).

The linear counting method wins on both speed and simplicity for this specific problem. It also generalizes nicely to problems where you need the rank of a value inside a sorted sequence without materializing the sequence.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)