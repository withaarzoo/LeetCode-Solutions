# Minimum Absolute Distance Between Mirror Pairs

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

A mirror pair is a pair of indices `(i, j)` such that:

* `0 <= i < j < nums.length`
* `reverse(nums[i]) == nums[j]`

The reverse of a number means reversing its digits.

Examples:

* `reverse(12) = 21`
* `reverse(120) = 21`
* `reverse(33) = 33`

We need to return the minimum absolute distance between any valid mirror pair.

If no mirror pair exists, return `-1`.

## Constraints

* `1 <= nums.length <= 10^5`
* `1 <= nums[i] <= 10^9`

## Intuition

At first, I thought about checking every pair `(i, j)`.

For every index `i`, I could compare it with every later index `j` and check whether `reverse(nums[i]) == nums[j]`.

But that would take `O(n^2)` time, which is too slow for `10^5` elements.

So I needed something faster.

I noticed that while scanning the array from left to right:

* If I already know the reversed value of previous numbers
* Then I can quickly check whether the current number forms a mirror pair

A hash map is perfect for this.

## Approach

1. Create a hash map.
2. Traverse the array from left to right.
3. For each number:

   * Check if it already exists in the map.
   * If yes, then a previous number's reverse matches the current number.
   * Update the minimum distance.
4. Reverse the current number.
5. Store the reversed value as key and the current index as value.
6. Keep updating the latest index because closer indices help us get a smaller answer.
7. If no pair is found, return `-1`.

## Data Structures Used

* Hash Map / Dictionary

  * Key = reversed number
  * Value = latest index where this reversed number was seen

## Operations & Behavior Summary

| Operation            | Purpose                                          |
| -------------------- | ------------------------------------------------ |
| Reverse Number       | Finds the reversed form of current number        |
| Hash Map Lookup      | Checks if current number already exists          |
| Hash Map Update      | Stores reverse(current number) with latest index |
| Distance Calculation | Computes `i - previousIndex`                     |
| Answer Update        | Keeps track of minimum distance                  |

## Complexity

* Time Complexity: `O(n * d)`

  * `n` = size of array
  * `d` = number of digits in each number
  * Since a number has at most 10 digits, this is practically `O(n)`

* Space Complexity: `O(n)`

  * We use a hash map to store reversed values and indices.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int reverseNum(int x) {
        int rev = 0;

        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x /= 10;
        }

        return rev;
    }

    int minMirrorPairDistance(vector<int>& nums) {
        unordered_map<int, int> lastIndex;
        int ans = INT_MAX;

        for (int i = 0; i < nums.size(); i++) {
            if (lastIndex.count(nums[i])) {
                ans = min(ans, i - lastIndex[nums[i]]);
            }

            int rev = reverseNum(nums[i]);
            lastIndex[rev] = i;
        }

        return (ans == INT_MAX) ? -1 : ans;
    }
};
```

### Java

```java
class Solution {
    private int reverseNum(int x) {
        int rev = 0;

        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x /= 10;
        }

        return rev;
    }

    public int minMirrorPairDistance(int[] nums) {
        HashMap<Integer, Integer> lastIndex = new HashMap<>();
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < nums.length; i++) {
            if (lastIndex.containsKey(nums[i])) {
                ans = Math.min(ans, i - lastIndex.get(nums[i]));
            }

            int rev = reverseNum(nums[i]);
            lastIndex.put(rev, i);
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

### JavaScript

```javascript
var minMirrorPairDistance = function(nums) {
    const reverseNum = (x) => {
        let rev = 0;

        while (x > 0) {
            rev = rev * 10 + (x % 10);
            x = Math.floor(x / 10);
        }

        return rev;
    };

    const lastIndex = new Map();
    let ans = Infinity;

    for (let i = 0; i < nums.length; i++) {
        if (lastIndex.has(nums[i])) {
            ans = Math.min(ans, i - lastIndex.get(nums[i]));
        }

        const rev = reverseNum(nums[i]);
        lastIndex.set(rev, i);
    }

    return ans === Infinity ? -1 : ans;
};
```

### Python3

```python
class Solution:
    def reverseNum(self, x: int) -> int:
        rev = 0

        while x > 0:
            rev = rev * 10 + (x % 10)
            x //= 10

        return rev

    def minMirrorPairDistance(self, nums: List[int]) -> int:
        last_index = {}
        ans = float('inf')

        for i, num in enumerate(nums):
            if num in last_index:
                ans = min(ans, i - last_index[num])

            rev = self.reverseNum(num)
            last_index[rev] = i

        return -1 if ans == float('inf') else ans
```

### Go

```go
func reverseNum(x int) int {
    rev := 0

    for x > 0 {
        rev = rev*10 + (x % 10)
        x /= 10
    }

    return rev
}

func minMirrorPairDistance(nums []int) int {
    lastIndex := make(map[int]int)
    ans := int(^uint(0) >> 1)

    for i, num := range nums {
        if prevIndex, exists := lastIndex[num]; exists {
            if i-prevIndex < ans {
                ans = i - prevIndex
            }
        }

        rev := reverseNum(num)
        lastIndex[rev] = i
    }

    if ans == int(^uint(0)>>1) {
        return -1
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Suppose:

```text
nums = [12, 21, 45, 33, 54]
```

Initially:

```text
lastIndex = {}
ans = infinity
```

### Step 1

Current number = `12`

* `12` is not present in map
* Reverse of `12` = `21`
* Store:

```text
lastIndex[21] = 0
```

Map becomes:

```text
{21: 0}
```

### Step 2

Current number = `21`

* `21` exists in map
* Previous index = `0`
* Distance = `1 - 0 = 1`
* Update answer:

```text
ans = 1
```

Reverse of `21` is `12`

Store:

```text
lastIndex[12] = 1
```

Map becomes:

```text
{21: 0, 12: 1}
```

### Step 3

Current number = `45`

* `45` is not present
* Reverse of `45` = `54`

Store:

```text
lastIndex[54] = 2
```

### Step 4

Current number = `33`

* `33` is not present
* Reverse of `33` = `33`

Store:

```text
lastIndex[33] = 3
```

### Step 5

Current number = `54`

* `54` exists in map
* Previous index = `2`
* Distance = `4 - 2 = 2`
* Minimum answer remains:

```text
min(1, 2) = 1
```

Final answer = `1`

## Examples

### Example 1

```text
Input: nums = [12,21,45,33,54]
Output: 1
```

### Example 2

```text
Input: nums = [120,21]
Output: 1
```

### Example 3

```text
Input: nums = [21,120]
Output: -1
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

* Brute force solution takes `O(n^2)` time.
* Using a hash map reduces the complexity to nearly `O(n)`.
* Storing only the latest index is important because it gives the minimum possible future distance.
* Reversing a number takes at most 10 operations because `nums[i] <= 10^9`.
* This solution easily works within the given constraints.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
