# Minimum Distance Between Three Equal Elements I

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

You are given an integer array `nums`.

A tuple `(i, j, k)` of 3 distinct indices is considered good if:

```text
nums[i] == nums[j] == nums[k]
```

The distance of a good tuple is:

```text
abs(i - j) + abs(j - k) + abs(k - i)
```

Your task is to return the minimum possible distance among all good tuples.

If no valid tuple exists, return `-1`.

## Constraints

```text
1 <= nums.length <= 100
1 <= nums[i] <= nums.length
```

## Intuition

I noticed that for any valid tuple `(i, j, k)` where:

```text
i < j < k
```

The formula becomes:

```text
|i - j| + |j - k| + |k - i|
= (j - i) + (k - j) + (k - i)
= 2 * (k - i)
```

This means I do not really need the middle index value separately.

I only need to minimize the difference between the first and third chosen index.

So for every number, I store all of its indices.
Then I check every consecutive group of 3 indices.

## Approach

1. Create a hashmap where:

   * Key = array value
   * Value = list of indices where that value appears

2. Traverse the array and store indices.

3. For each value:

   * If it appears less than 3 times, skip it.
   * Otherwise, check every consecutive group of 3 indices.

4. For indices `[a, b, c]`, distance becomes:

```text
2 * (c - a)
```

1. Keep track of the minimum answer.

2. If no valid tuple exists, return `-1`.

## Data Structures Used

* HashMap / Dictionary
* Array / List / Vector

The hashmap is used to group indices of the same value together.

## Operations & Behavior Summary

| Operation                | Purpose                            |
| ------------------------ | ---------------------------------- |
| Store indices in hashmap | Group same values together         |
| Traverse index list      | Find valid groups of 3 occurrences |
| Compute distance         | Use `2 * (last - first)`           |
| Track minimum            | Keep smallest valid distance       |

## Complexity

* Time Complexity: `O(n)`

  * `n` is the size of the array.
  * I traverse the array once and then process stored indices once.

* Space Complexity: `O(n)`

  * Extra space is used for storing indices inside the hashmap.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumDistance(vector<int>& nums) {
        unordered_map<int, vector<int>> pos;
        
        // Store all indices for each value
        for (int i = 0; i < nums.size(); i++) {
            pos[nums[i]].push_back(i);
        }
        
        int ans = INT_MAX;
        
        // Process each value's index list
        for (auto& entry : pos) {
            vector<int>& indices = entry.second;
            
            if (indices.size() < 3) continue;
            
            for (int i = 0; i + 2 < indices.size(); i++) {
                int distance = 2 * (indices[i + 2] - indices[i]);
                ans = min(ans, distance);
            }
        }
        
        return ans == INT_MAX ? -1 : ans;
    }
};
```

### Java

```java
class Solution {
    public int minimumDistance(int[] nums) {
        Map<Integer, List<Integer>> map = new HashMap<>();
        
        // Store all indices for each value
        for (int i = 0; i < nums.length; i++) {
            map.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }
        
        int ans = Integer.MAX_VALUE;
        
        // Process each value's indices
        for (List<Integer> indices : map.values()) {
            if (indices.size() < 3) continue;
            
            for (int i = 0; i + 2 < indices.size(); i++) {
                int distance = 2 * (indices.get(i + 2) - indices.get(i));
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
    const map = new Map();
    
    // Store all indices for each value
    for (let i = 0; i < nums.length; i++) {
        if (!map.has(nums[i])) {
            map.set(nums[i], []);
        }
        map.get(nums[i]).push(i);
    }
    
    let ans = Infinity;
    
    // Process each value's indices
    for (const indices of map.values()) {
        if (indices.length < 3) continue;
        
        for (let i = 0; i + 2 < indices.length; i++) {
            const distance = 2 * (indices[i + 2] - indices[i]);
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
        pos = {}
        
        # Store all indices for each value
        for i, num in enumerate(nums):
            if num not in pos:
                pos[num] = []
            pos[num].append(i)
        
        ans = float('inf')
        
        # Process each value's indices
        for indices in pos.values():
            if len(indices) < 3:
                continue
            
            for i in range(len(indices) - 2):
                distance = 2 * (indices[i + 2] - indices[i])
                ans = min(ans, distance)
        
        return -1 if ans == float('inf') else ans
```

### Go

```go
func minimumDistance(nums []int) int {
    pos := make(map[int][]int)
    
    // Store all indices for each value
    for i, num := range nums {
        pos[num] = append(pos[num], i)
    }
    
    ans := int(^uint(0) >> 1)
    
    // Process each value's indices
    for _, indices := range pos {
        if len(indices) < 3 {
            continue
        }
        
        for i := 0; i+2 < len(indices); i++ {
            distance := 2 * (indices[i+2] - indices[i])
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

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

### Step 1: Create a hashmap

I first create a hashmap to store indices of every value.

Example:

```text
nums = [1, 2, 1, 1, 3]
```

Hashmap becomes:

```text
1 -> [0, 2, 3]
2 -> [1]
3 -> [4]
```

---

### Step 2: Store all indices

While traversing the array, I store the current index for its value.

```text
nums[i] -> push index i
```

---

### Step 3: Skip values with less than 3 occurrences

A valid tuple needs exactly 3 equal values.

So if any value appears less than 3 times, I skip it.

---

### Step 4: Check every consecutive group of 3 indices

Suppose:

```text
indices = [0, 2, 3, 7]
```

Then possible groups are:

```text
[0, 2, 3]
[2, 3, 7]
```

For each group:

```text
distance = 2 * (lastIndex - firstIndex)
```

Example:

```text
2 * (3 - 0) = 6
2 * (7 - 2) = 10
```

Minimum is `6`.

---

### Step 5: Keep updating minimum answer

I keep comparing every valid distance with the current answer.

```text
ans = min(ans, currentDistance)
```

---

### Step 6: Return final result

If I found at least one valid tuple, I return the minimum distance.

Otherwise, I return:

```text
-1
```

## Examples

### Example 1

```text
Input: nums = [1,2,1,1,3]
Output: 6
```

Explanation:

```text
Indices of 1 = [0, 2, 3]
Distance = 2 * (3 - 0) = 6
```

### Example 2

```text
Input: nums = [1,1,2,3,2,1,2]
Output: 8
```

Explanation:

```text
Indices of 2 = [2, 4, 6]
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
g++ main.cpp -o main
./main
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

* I simplified the original formula into `2 * (k - i)`.
* This avoids recalculating all absolute values repeatedly.
* I only check consecutive groups of 3 indices because they always give the minimum possible span.
* The solution is efficient and works in linear time.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
