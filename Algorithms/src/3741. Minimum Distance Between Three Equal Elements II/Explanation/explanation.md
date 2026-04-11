# Minimum Distance Between Three Equal Elements II

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

You are given an integer array `nums`.

A tuple `(i, j, k)` is considered good if:

* `i`, `j`, and `k` are distinct indices
* `nums[i] == nums[j] == nums[k]`

The distance of a good tuple is:

```text
|i - j| + |j - k| + |k - i|
```

You need to return the minimum possible distance among all good tuples.

If no such tuple exists, return `-1`.

## Constraints

```text
1 <= nums.length <= 10^5
1 <= nums[i] <= nums.length
```

## Intuition

I noticed that for any valid tuple `(i, j, k)` where:

```text
i < j < k
```

The distance formula becomes:

```text
|i-j| + |j-k| + |k-i|
```

Which simplifies to:

```text
(j-i) + (k-j) + (k-i)
= 2 * (k-i)
```

So I do not need to care about the middle index separately.

I only need the smallest possible difference between the first and third index.

For every number, I store all its indices.
Then I check every consecutive group of 3 indices because that gives the minimum possible range.

## Approach

1. Create a hashmap where:

   * key = number in the array
   * value = list of indices where the number appears

2. Traverse the array and store each index.

3. For every value:

   * If it appears less than 3 times, skip it.
   * Otherwise, check every consecutive triple of indices.

4. Suppose the three indices are:

```text
[a, b, c]
```

Then the distance becomes:

```text
2 * (c - a)
```

1. Keep updating the minimum answer.

2. If no valid triple is found, return `-1`.

## Data Structures Used

* HashMap / Dictionary

  * Used to store all indices of each value.

* Dynamic Array / List / Vector

  * Used to store positions of the same value.

## Operations & Behavior Summary

| Operation                 | Purpose                                   |
| ------------------------- | ----------------------------------------- |
| Store indices             | Save positions of each value              |
| Skip small lists          | Ignore values appearing less than 3 times |
| Check consecutive triples | Find minimum span                         |
| Compute distance          | Use `2 * (lastIndex - firstIndex)`        |
| Update answer             | Keep track of minimum distance            |

## Complexity

* Time Complexity: `O(n)`

  * `n` is the length of the array.
  * Every index is stored once and processed once.

* Space Complexity: `O(n)`

  * Extra hashmap is used to store indices.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDistance(vector<int>& nums) {
        unordered_map<int, vector<int>> positions;

        for (int i = 0; i < nums.size(); i++) {
            positions[nums[i]].push_back(i);
        }

        int ans = INT_MAX;

        for (auto& entry : positions) {
            vector<int>& idx = entry.second;

            if (idx.size() < 3) continue;

            for (int i = 0; i + 2 < idx.size(); i++) {
                int distance = 2 * (idx[i + 2] - idx[i]);
                ans = min(ans, distance);
            }
        }

        return (ans == INT_MAX) ? -1 : ans;
    }
};
```

### Java

```java
class Solution {
    public int minimumDistance(int[] nums) {
        Map<Integer, List<Integer>> positions = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            positions.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        int ans = Integer.MAX_VALUE;

        for (List<Integer> idx : positions.values()) {
            if (idx.size() < 3) continue;

            for (int i = 0; i + 2 < idx.size(); i++) {
                int distance = 2 * (idx.get(i + 2) - idx.get(i));
                ans = Math.min(ans, distance);
            }
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDistance = function(nums) {
    const positions = new Map();

    for (let i = 0; i < nums.length; i++) {
        if (!positions.has(nums[i])) {
            positions.set(nums[i], []);
        }
        positions.get(nums[i]).push(i);
    }

    let ans = Infinity;

    for (const idx of positions.values()) {
        if (idx.length < 3) continue;

        for (let i = 0; i + 2 < idx.length; i++) {
            const distance = 2 * (idx[i + 2] - idx[i]);
            ans = Math.min(ans, distance);
        }
    }

    return ans === Infinity ? -1 : ans;
};
```

### Python3

```python
class Solution:
    def minimumDistance(self, nums: List[int]) -> int:
        positions = {}

        for i, num in enumerate(nums):
            if num not in positions:
                positions[num] = []
            positions[num].append(i)

        ans = float('inf')

        for idx in positions.values():
            if len(idx) < 3:
                continue

            for i in range(len(idx) - 2):
                distance = 2 * (idx[i + 2] - idx[i])
                ans = min(ans, distance)

        return -1 if ans == float('inf') else ans
```

### Go

```go
func minimumDistance(nums []int) int {
    positions := make(map[int][]int)

    for i, num := range nums {
        positions[num] = append(positions[num], i)
    }

    ans := int(^uint(0) >> 1)

    for _, idx := range positions {
        if len(idx) < 3 {
            continue
        }

        for i := 0; i+2 < len(idx); i++ {
            distance := 2 * (idx[i+2] - idx[i])
            if distance < ans {
                ans = distance
            }
        }
    }

    if ans == int(^uint(0)>>1) {
        return -1
    }

    return ans
}
```

## Step-by-step Detailed Explanation

### C++

```cpp
unordered_map<int, vector<int>> positions;
```

I create a hashmap where:

* key = number
* value = list of indices

---

```cpp
for (int i = 0; i < nums.size(); i++) {
    positions[nums[i]].push_back(i);
}
```

I store every index of each value.

Example:

```text
nums = [1,2,1,1,3]
```

Stored as:

```text
1 -> [0,2,3]
2 -> [1]
3 -> [4]
```

---

```cpp
if (idx.size() < 3) continue;
```

If a number appears less than 3 times, it cannot form a valid tuple.

---

```cpp
for (int i = 0; i + 2 < idx.size(); i++) {
    int distance = 2 * (idx[i + 2] - idx[i]);
}
```

I only check consecutive triples.

Example:

```text
[1,4,7,10]
```

Triples:

```text
(1,4,7)
(4,7,10)
```

---

```cpp
ans = min(ans, distance);
```

I keep track of the minimum answer.

### Java

The logic is exactly the same.

* Use `HashMap<Integer, List<Integer>>`
* Store indices for every value
* Check consecutive triples only
* Update the minimum answer

### JavaScript

The logic is exactly the same.

* Use `Map()`
* Store indices in arrays
* Traverse consecutive groups of 3
* Keep the minimum answer

### Python3

The logic is exactly the same.

* Use dictionary
* Store positions in lists
* Check every consecutive triple
* Return minimum distance

### Go

The logic is exactly the same.

* Use map with slices
* Store all indices
* Check consecutive triples
* Return smallest answer

## Examples

### Example 1

```text
Input: nums = [1,2,1,1,3]
Output: 6
```

Explanation:

```text
Indices of 1 = [0,2,3]
Distance = 2 * (3 - 0) = 6
```

### Example 2

```text
Input: nums = [1,1,2,3,2,1,2]
Output: 8
```

Explanation:

```text
Indices of 2 = [2,4,6]
Distance = 2 * (6 - 2) = 8
```

### Example 3

```text
Input: nums = [1]
Output: -1
```

Explanation:

```text
No value appears 3 times.
```

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* I do not need to check all possible triples.
* Checking all triples would be too slow.
* Consecutive triples always give the minimum range.
* This makes the solution linear.
* The solution works efficiently even for `10^5` elements.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
