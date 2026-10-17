# Minimum Swaps to Sort by Digit Sum | LeetCode 3551 Solution

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

You are given an array of distinct positive integers. The goal is to rearrange the array so that the numbers appear in increasing order based on the sum of their digits. If two numbers have the same digit sum, the smaller number should come first.

You need to return the minimum number of swaps required to turn the given array into this special sorted order. A swap means exchanging the values at any two different positions.

This is a classic minimum swaps problem that can be solved efficiently by reducing it to cycle detection in a permutation.

## Constraints

- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9
- All elements in nums are distinct positive integers

## Intuition

The final order of the array is completely fixed once we decide the sorting rule. Because every number is unique, there is only one correct sorted version.

I realized that finding the minimum swaps to reach a known target order is the same as counting cycles in the permutation that maps the current positions to the target positions. The formula is simple: minimum swaps equals n minus the number of cycles.

That observation turns a sorting-with-swaps problem into a straightforward cycle-counting problem after we build the correct target order.

## Approach

First calculate the digit sum of every number.  
Create a sorted copy of the array where the primary key is the digit sum and the secondary key is the number itself.  

Build a map that records the correct index of every value in this sorted array.  

Now treat the original array as a permutation. For each position, look up where its current value should actually go. Follow these links to discover the cycles.  

Finally return the length of the array minus the number of cycles found. That gives the minimum number of swaps needed.

## Data Structures Used

- A copy of the array used for sorting (to preserve the original order while creating the target).
- A hash map (or unordered_map / Map / dict) that stores the target index of each value.
- A boolean visited array to track which positions have already been examined while counting cycles.

These structures keep the solution clean and run in linear time after the initial sort.

## Operations & Behavior Summary

1. Compute digit sums on the fly while sorting.
2. Sort a copy of the input using the two-key rule (digit sum, then value).
3. Record the sorted position of every number in a map.
4. Walk through every index of the original array.
5. Whenever an unvisited index is found, start a new cycle and keep jumping to the target position of the current value until the cycle closes.
6. Subtract the total number of cycles from the array length to get the answer.

## Complexity

| Type              | Complexity | Explanation |
|-------------------|------------|-------------|
| Time Complexity   | O(n log n) | Sorting the array dominates. Building the map and counting cycles are both O(n). Digit-sum calculation is O(log A) per number where A is the maximum value, which is negligible. |
| Space Complexity  | O(n)       | Extra space is used for the sorted copy, the position map, and the visited array. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int minSwaps(vector<int>& nums) {
        auto digitSum = [](int x) {
            int s = 0;
            while (x) {
                s += x % 10;
                x /= 10;
            }
            return s;
        };
        int n = nums.size();
        vector<int> sortedNums = nums;
        sort(sortedNums.begin(), sortedNums.end(), [&](int a, int b) {
            int sa = digitSum(a), sb = digitSum(b);
            if (sa != sb) return sa < sb;
            return a < b;
        });
        unordered_map<int, int> pos;
        for (int i = 0; i < n; i++) pos[sortedNums[i]] = i;
        vector<bool> visited(n, false);
        int cycles = 0;
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                cycles++;
                int j = i;
                while (!visited[j]) {
                    visited[j] = true;
                    j = pos[nums[j]];
                }
            }
        }
        return n - cycles;
    }
};
```

### Java
```java
class Solution {
    public int minSwaps(int[] nums) {
        int n = nums.length;
        Integer[] sorted = new Integer[n];
        for (int i = 0; i < n; i++) sorted[i] = nums[i];
        Arrays.sort(sorted, (a, b) -> {
            int sa = digitSum(a), sb = digitSum(b);
            if (sa != sb) return sa - sb;
            return a - b;
        });
        Map<Integer, Integer> pos = new HashMap<>();
        for (int i = 0; i < n; i++) pos.put(sorted[i], i);
        boolean[] visited = new boolean[n];
        int cycles = 0;
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                cycles++;
                int j = i;
                while (!visited[j]) {
                    visited[j] = true;
                    j = pos.get(nums[j]);
                }
            }
        }
        return n - cycles;
    }
    private int digitSum(int x) {
        int s = 0;
        while (x > 0) {
            s += x % 10;
            x /= 10;
        }
        return s;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minSwaps = function(nums) {
    const digitSum = (x) => {
        let s = 0;
        while (x > 0) {
            s += x % 10;
            x = Math.floor(x / 10);
        }
        return s;
    };
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => {
        const sa = digitSum(a), sb = digitSum(b);
        if (sa !== sb) return sa - sb;
        return a - b;
    });
    const pos = new Map();
    for (let i = 0; i < n; i++) pos.set(sorted[i], i);
    const visited = new Array(n).fill(false);
    let cycles = 0;
    for (let i = 0; i < n; i++) {
        if (!visited[i]) {
            cycles++;
            let j = i;
            while (!visited[j]) {
                visited[j] = true;
                j = pos.get(nums[j]);
            }
        }
    }
    return n - cycles;
};
```

### TypeScript
```typescript
function minSwaps(nums: number[]): number {
    const digitSum = (x: number): number => {
        let s = 0;
        while (x > 0) {
            s += x % 10;
            x = Math.floor(x / 10);
        }
        return s;
    };
    const n = nums.length;
    const sorted = [...nums].sort((a, b) => {
        const sa = digitSum(a), sb = digitSum(b);
        if (sa !== sb) return sa - sb;
        return a - b;
    });
    const pos = new Map<number, number>();
    for (let i = 0; i < n; i++) pos.set(sorted[i], i);
    const visited = new Array(n).fill(false);
    let cycles = 0;
    for (let i = 0; i < n; i++) {
        if (!visited[i]) {
            cycles++;
            let j = i;
            while (!visited[j]) {
                visited[j] = true;
                j = pos.get(nums[j])!;
            }
        }
    }
    return n - cycles;
};
```

### Python3
```python
class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        def digit_sum(x):
            s = 0
            while x:
                s += x % 10
                x //= 10
            return s
        n = len(nums)
        sorted_nums = sorted(nums, key=lambda x: (digit_sum(x), x))
        pos = {val: i for i, val in enumerate(sorted_nums)}
        visited = [False] * n
        cycles = 0
        for i in range(n):
            if not visited[i]:
                cycles += 1
                j = i
                while not visited[j]:
                    visited[j] = True
                    j = pos[nums[j]]
        return n - cycles
```

### Go
```go
func minSwaps(nums []int) int {
    digitSum := func(x int) int {
        s := 0
        for x > 0 {
            s += x % 10
            x /= 10
        }
        return s
    }
    n := len(nums)
    sortedNums := make([]int, n)
    copy(sortedNums, nums)
    sort.Slice(sortedNums, func(i, j int) bool {
        sa, sb := digitSum(sortedNums[i]), digitSum(sortedNums[j])
        if sa != sb {
            return sa < sb
        }
        return sortedNums[i] < sortedNums[j]
    })
    pos := make(map[int]int)
    for i, v := range sortedNums {
        pos[v] = i
    }
    visited := make([]bool, n)
    cycles := 0
    for i := 0; i < n; i++ {
        if !visited[i] {
            cycles++
            j := i
            for !visited[j] {
                visited[j] = true
                j = pos[nums[j]]
            }
        }
    }
    return n - cycles
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core logic is identical across all six languages. Only syntax and library details change.

I begin by writing a small helper that repeatedly takes the last digit and divides the number by ten until it becomes zero. This produces the digit sum needed for comparisons.

Next I create a sorted version of the input. The comparison first checks digit sums. When the sums are equal it falls back to comparing the actual numbers. After sorting I know the exact final position of every value.

I store those final positions in a map keyed by the number itself. Looking up any value instantly tells me where it belongs.

Then I allocate a visited array of size n. I scan every index. If the current index has not been visited, I have discovered a new cycle. I increment the cycle counter and start following the chain: take the number sitting at the current index, look up its target index in the map, and jump there. I mark every position I visit so I never process the same cycle twice.

When the scan finishes, the answer is simply n minus the number of cycles. This works because each cycle of length k can be sorted with exactly k-1 swaps, and summing over all cycles yields n - number_of_cycles.

Edge cases are handled naturally. An already-sorted array produces n cycles of length 1, so the answer is zero. A single full cycle produces the maximum number of swaps (n-1). Because the numbers are guaranteed distinct, the mapping is always a valid permutation and no extra checks are required.

## Examples

Example 1  
Input: nums = [37, 100]  
Digit sums: 10 and 1  
Sorted order: [100, 37]  
The two elements form one cycle of length 2.  
Minimum swaps = 2 - 1 = 1

Example 2  
Input: nums = [22, 14, 33, 7]  
Digit sums: 4, 5, 6, 7  
The array is already in the required order, producing four cycles of length 1.  
Minimum swaps = 4 - 4 = 0

Example 3  
Input: nums = [18, 43, 34, 16]  
Digit sums: 9, 7, 7, 7  
Sorted order: [16, 34, 43, 18]  
Two cycles appear (one of length 2 for 18 and 16, another of length 2 for 43 and 34).  
Minimum swaps = 4 - 2 = 2

## How to Use / Run Locally

C++  
Save the code in a file named main.cpp. Compile with g++ -std=c++17 main.cpp -o main and run ./main. Add a simple main function that reads the array and prints the result if you want to test interactively.

Java  
Save the code inside a file named Solution.java. Compile with javac Solution.java and run java Solution. You will need a small driver that creates an instance and calls minSwaps.

JavaScript  
Save the code in a file named solution.js. Run it with node solution.js after adding a few console.log test cases at the bottom.

TypeScript  
Save the code in a file named solution.ts. Compile with tsc solution.ts and then run the generated JavaScript with node, or use ts-node solution.ts directly.

Python3  
Save the code in a file named solution.py. Run it with python3 solution.py. Add a few print statements calling the method for quick testing.

Go  
Save the code in a file named main.go. Run it with go run main.go after adding a main function that prepares a sample slice and prints the result.

## Notes & Optimizations

The solution already runs in the optimal O(n log n) time because of the required sort. No further asymptotic improvement is possible under the given constraints.

Because nums[i] can be as large as 10^9, the digit-sum helper must be careful with integer division; all six implementations handle this correctly.

If the problem ever allowed duplicate numbers the cycle approach would break, but the constraints guarantee distinctness, so the mapping stays a pure permutation.

An alternative way would be to generate the sorted array and then solve the assignment problem, but that would be slower and more complicated. Cycle counting remains the cleanest and fastest practical method.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)