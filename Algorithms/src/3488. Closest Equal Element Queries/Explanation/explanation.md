# Closest Equal Element Queries

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

You are given:

* A circular array `nums`
* An array `queries`

For every query index `i`, you need to find the minimum circular distance between `queries[i]` and any other index having the same value.

If there is no other index with the same value, return `-1` for that query.

A circular array means:

* Moving left from index `0` goes to index `n - 1`
* Moving right from index `n - 1` goes to index `0`

## Constraints

```text
1 <= queries.length <= nums.length <= 10^5
1 <= nums[i] <= 10^6
0 <= queries[i] < nums.length
```

## Intuition

I thought about grouping all indices of the same value together.

For example:

```text
1 -> [0, 2, 4]
3 -> [1, 5]
```

Now for any query index, I already know all the places where the same value appears.

The nearest equal element will always be either:

* The previous occurrence
* The next occurrence

Because the positions are already sorted naturally while traversing the array.

So instead of checking every possible index, I only check the previous and next equal occurrence.

This makes the solution very efficient.

## Approach

1. Create a hashmap where:

   * Key = number in `nums`
   * Value = list of all indices where it appears

2. For every group of equal values:

   * If the value appears only once, answer remains `-1`
   * Otherwise:

     * Find previous occurrence
     * Find next occurrence
     * Compute circular distance for both
     * Store the minimum distance

3. Precompute answer for every index.

4. For each query, directly return the stored answer.

## Data Structures Used

* HashMap / Dictionary

  * To store all positions of each value
* Array

  * To store precomputed answer for each index
* Vector / List

  * To store all occurrence indices for a value

## Operations & Behavior Summary

| Operation           | Purpose                      |
| ------------------- | ---------------------------- |
| Store indices       | Group same values together   |
| Previous occurrence | Check left-side equal value  |
| Next occurrence     | Check right-side equal value |
| Circular distance   | Handle wrap-around movement  |
| Precompute answers  | Answer queries in O(1)       |

## Complexity

* Time Complexity: `O(n + q)`

  * `n` = length of `nums`
  * `q` = length of `queries`

* Space Complexity: `O(n)`

  * Extra hashmap for storing indices
  * Extra answer array

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> solveQueries(vector<int>& nums, vector<int>& queries) {
        int n = nums.size();

        unordered_map<int, vector<int>> positions;

        for (int i = 0; i < n; i++) {
            positions[nums[i]].push_back(i);
        }

        vector<int> answer(n, -1);

        for (auto& entry : positions) {
            vector<int>& pos = entry.second;
            int m = pos.size();

            if (m == 1) continue;

            for (int i = 0; i < m; i++) {
                int curr = pos[i];
                int prev = pos[(i - 1 + m) % m];
                int next = pos[(i + 1) % m];

                int distPrev = abs(curr - prev);
                distPrev = min(distPrev, n - distPrev);

                int distNext = abs(curr - next);
                distNext = min(distNext, n - distNext);

                answer[curr] = min(distPrev, distNext);
            }
        }

        vector<int> result;

        for (int idx : queries) {
            result.push_back(answer[idx]);
        }

        return result;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> solveQueries(int[] nums, int[] queries) {
        int n = nums.length;

        Map<Integer, List<Integer>> positions = new HashMap<>();

        for (int i = 0; i < n; i++) {
            positions.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        int[] answer = new int[n];
        Arrays.fill(answer, -1);

        for (List<Integer> pos : positions.values()) {
            int m = pos.size();

            if (m == 1) continue;

            for (int i = 0; i < m; i++) {
                int curr = pos.get(i);
                int prev = pos.get((i - 1 + m) % m);
                int next = pos.get((i + 1) % m);

                int distPrev = Math.abs(curr - prev);
                distPrev = Math.min(distPrev, n - distPrev);

                int distNext = Math.abs(curr - next);
                distNext = Math.min(distNext, n - distNext);

                answer[curr] = Math.min(distPrev, distNext);
            }
        }

        List<Integer> result = new ArrayList<>();

        for (int idx : queries) {
            result.add(answer[idx]);
        }

        return result;
    }
}
```

### JavaScript

```javascript
var solveQueries = function(nums, queries) {
    const n = nums.length;
    const positions = new Map();

    for (let i = 0; i < n; i++) {
        if (!positions.has(nums[i])) {
            positions.set(nums[i], []);
        }
        positions.get(nums[i]).push(i);
    }

    const answer = new Array(n).fill(-1);

    for (const pos of positions.values()) {
        const m = pos.length;

        if (m === 1) continue;

        for (let i = 0; i < m; i++) {
            const curr = pos[i];
            const prev = pos[(i - 1 + m) % m];
            const next = pos[(i + 1) % m];

            let distPrev = Math.abs(curr - prev);
            distPrev = Math.min(distPrev, n - distPrev);

            let distNext = Math.abs(curr - next);
            distNext = Math.min(distNext, n - distNext);

            answer[curr] = Math.min(distPrev, distNext);
        }
    }

    return queries.map(idx => answer[idx]);
};
```

### Python3

```python
class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)

        positions = {}

        for i, num in enumerate(nums):
            if num not in positions:
                positions[num] = []
            positions[num].append(i)

        answer = [-1] * n

        for pos in positions.values():
            m = len(pos)

            if m == 1:
                continue

            for i in range(m):
                curr = pos[i]
                prev_idx = pos[(i - 1 + m) % m]
                next_idx = pos[(i + 1) % m]

                dist_prev = abs(curr - prev_idx)
                dist_prev = min(dist_prev, n - dist_prev)

                dist_next = abs(curr - next_idx)
                dist_next = min(dist_next, n - dist_next)

                answer[curr] = min(dist_prev, dist_next)

        return [answer[idx] for idx in queries]
```

### Go

```go
func solveQueries(nums []int, queries []int) []int {
    n := len(nums)

    positions := make(map[int][]int)

    for i, num := range nums {
        positions[num] = append(positions[num], i)
    }

    answer := make([]int, n)

    for i := 0; i < n; i++ {
        answer[i] = -1
    }

    for _, pos := range positions {
        m := len(pos)

        if m == 1 {
            continue
        }

        for i := 0; i < m; i++ {
            curr := pos[i]
            prev := pos[(i-1+m)%m]
            next := pos[(i+1)%m]

            distPrev := abs(curr - prev)
            distPrev = min(distPrev, n-distPrev)

            distNext := abs(curr - next)
            distNext = min(distNext, n-distNext)

            answer[curr] = min(distPrev, distNext)
        }
    }

    result := make([]int, len(queries))

    for i, idx := range queries {
        result[i] = answer[idx]
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Store all positions of each value

```text
nums = [1,3,1,4,1,3,2]
```

Stored as:

```text
1 -> [0,2,4]
3 -> [1,5]
4 -> [3]
2 -> [6]
```

This helps quickly find all equal elements.

### Step 2: Skip unique values

If a value appears only once, then there is no other equal element.

```text
4 -> [3]
2 -> [6]
```

Their answer will remain `-1`.

### Step 3: Find previous and next occurrence

For:

```text
1 -> [0,2,4]
```

At index `2`:

* Previous occurrence = `0`
* Next occurrence = `4`

At index `0`:

* Previous occurrence = `4`
* Next occurrence = `2`

Modulo indexing helps us handle circular movement.

### Step 4: Compute circular distance

Formula:

```text
normalDistance = abs(a - b)
circularDistance = n - normalDistance
actualDistance = min(normalDistance, circularDistance)
```

Example:

```text
n = 7
a = 0
b = 5
```

```text
normalDistance = 5
circularDistance = 2
actualDistance = 2
```

### Step 5: Store answer for every index

For each occurrence, compare:

* Distance to previous equal index
* Distance to next equal index

Store the smaller one.

Then every query can be answered instantly.

## Examples

### Example 1

```text
Input:
nums = [1,3,1,4,1,3,2]
queries = [0,3,5]

Output:
[2,-1,3]
```

Explanation:

* Query index `0` -> nearest equal value is at index `2`, distance = `2`
* Query index `3` -> no equal value exists, answer = `-1`
* Query index `5` -> nearest equal value is at index `1`, circular distance = `3`

### Example 2

```text
Input:
nums = [1,2,3,4]
queries = [0,1,2,3]

Output:
[-1,-1,-1,-1]
```

## How to use / Run locally

### C++

```bash
g++ main.cpp -o main
./main
```

### Java

```bash
javac Main.java
java Main
```

### JavaScript

```bash
node main.js
```

### Python3

```bash
python main.py
```

### Go

```bash
go run main.go
```

## Notes & Optimizations

* I only check previous and next equal occurrence.
* I do not compare with every equal index.
* This reduces the solution from `O(n^2)` to `O(n)`.
* Precomputing answers makes queries very fast.
* Circular distance is handled using:

```text
min(abs(a - b), n - abs(a - b))
```

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
