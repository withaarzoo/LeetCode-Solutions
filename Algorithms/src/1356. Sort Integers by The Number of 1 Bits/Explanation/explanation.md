# 1356. Sort Integers by The Number of 1 Bits

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

Given an integer array `arr`, sort the integers in ascending order based on the number of 1 bits in their binary representation.

If two integers have the same number of 1 bits, then sort them in ascending numerical order.

Return the sorted array.

---

## Constraints

* 1 <= arr.length <= 500
* 0 <= arr[i] <= 10^4

---

## Intuition

When I first looked at this problem, I understood that normal sorting is not enough.

I need to sort numbers based on:

1. The number of 1 bits in their binary form.
2. If two numbers have the same number of 1 bits, then sort them normally.

So I thought this is clearly a custom sorting problem.

Instead of manually building complex logic, I can use a sorting function with a custom comparator or key function.

---

## Approach

Here is how I solved it step by step:

1. For each number, count how many 1 bits it has.
2. Sort the array using custom logic:

   * First compare by number of set bits.
   * If equal, compare by actual number.
3. Return the sorted array.

Most languages provide built-in optimized functions to count set bits, so I used them directly.

---

## Data Structures Used

* Input array
* Sorting mechanism with custom comparator

No extra data structures are required.

---

## Operations & Behavior Summary

* Convert number to binary or use built-in bit counting.
* Compare numbers based on bit count.
* If bit counts match, compare numerical values.
* Sorting ensures stable and correct ordering.

---

## Complexity

Time Complexity: O(n log n)

* n is the number of elements in the array.
* Sorting takes O(n log n).
* Counting bits takes O(1) since numbers are small (maximum 14 bits).

Space Complexity: O(1)

* No extra data structures are used.
* Sorting is done in-place (ignoring recursion stack or internal sorting memory).

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> sortByBits(vector<int>& arr) {
        sort(arr.begin(), arr.end(), [](int a, int b) {
            int bitsA = __builtin_popcount(a);
            int bitsB = __builtin_popcount(b);

            if (bitsA != bitsB)
                return bitsA < bitsB;

            return a < b;
        });
        return arr;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int[] sortByBits(int[] arr) {
        Integer[] nums = Arrays.stream(arr).boxed().toArray(Integer[]::new);

        Arrays.sort(nums, (a, b) -> {
            int bitsA = Integer.bitCount(a);
            int bitsB = Integer.bitCount(b);

            if (bitsA != bitsB)
                return bitsA - bitsB;

            return a - b;
        });

        for (int i = 0; i < arr.length; i++)
            arr[i] = nums[i];

        return arr;
    }
}
```

### JavaScript

```javascript
var sortByBits = function(arr) {
    return arr.sort((a, b) => {
        const bitsA = a.toString(2).split('1').length - 1;
        const bitsB = b.toString(2).split('1').length - 1;

        if (bitsA !== bitsB)
            return bitsA - bitsB;

        return a - b;
    });
};
```

### Python3

```python
from typing import List

class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        return sorted(arr, key=lambda x: (bin(x).count('1'), x))
```

### Go

```go
import (
    "math/bits"
    "sort"
)

func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        bitsI := bits.OnesCount(uint(arr[i]))
        bitsJ := bits.OnesCount(uint(arr[j]))

        if bitsI != bitsJ {
            return bitsI < bitsJ
        }
        return arr[i] < arr[j]
    })
    return arr
}
```

---

## Step-by-step Detailed Explanation

### Step 1: Count set bits

Each number is converted to binary representation or passed to a built-in bit counting function.

For example:

* 5 in binary is 101, which has 2 set bits.
* 3 in binary is 11, which has 2 set bits.

### Step 2: Custom comparison

While sorting:

* If two numbers have different set bit counts, the one with fewer bits comes first.
* If they have the same set bit count, the smaller number comes first.

### Step 3: Language-specific implementation

* C++ uses __builtin_popcount for fast bit counting.
* Java uses Integer.bitCount.
* JavaScript converts to binary string and counts '1'.
* Python uses tuple-based sorting with key.
* Go uses bits.OnesCount from math/bits package.

### Step 4: Final result

After sorting with this logic, the array automatically satisfies both conditions.

---

## Examples

Example 1:
Input: [0,1,2,3,4,5,6,7,8]
Output: [0,1,2,4,8,3,5,6,7]

Example 2:
Input: [1024,512,256,128,64,32,16,8,4,2,1]
Output: [1,2,4,8,16,32,64,128,256,512,1024]

---

## How to use / Run locally

1. Copy the solution code for your preferred language.
2. Paste it into your LeetCode editor or local IDE.
3. Compile and run with sample test cases.
4. Verify the output.

---

## Notes & Optimizations

* Always prefer built-in bit counting functions as they are optimized at hardware level.
* Avoid manually shifting bits unless required.
* Sorting dominates the time complexity.
* Since maximum value is 10^4, bit counting is constant time.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
