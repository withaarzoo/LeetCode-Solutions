# Sum of GCD of Formed Pairs | LeetCode 3867 Solution (C++, Java, JavaScript, Python, Go)

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
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

LeetCode 3867, **Sum of GCD of Formed Pairs**, asks us to process an integer array in multiple stages.

First, we build a new array called `prefixGcd`. For every index, we find the maximum value seen so far in the original array and calculate the Greatest Common Divisor (GCD) between the current element and that maximum.

Once the new array is ready, we sort it in non-decreasing order. Then we repeatedly form pairs by taking the smallest unused value and the largest unused value. For every pair, we calculate their GCD and add it to the final answer.

If the array contains an odd number of elements, the middle element remains unused and should be ignored.

The goal is to return the total sum of the GCD values of all formed pairs.

This solution uses simple simulation, sorting, two pointers, and the Euclidean Algorithm for GCD, making it both clean and efficient.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= nums.length <= 10^5` | Array size |
| `1 <= nums[i] <= 10^9` | Element value |

---

## Intuition

The first thing I noticed was that computing the prefix maximum doesn't need to be repeated for every position. I can simply maintain one running maximum while traversing the array from left to right.

Once I have the prefix maximum, calculating the GCD for the current element becomes straightforward.

The second half of the problem is even simpler because the pairing rule is already fixed. After sorting the array, the smallest and largest unused elements are always available at the two ends.

That naturally leads to a two-pointer solution.

Instead of thinking about many different pairing combinations, I simply follow the order given in the problem statement.

---

## Approach

My solution follows these steps.

1. Create an empty array called `prefixGcd`.
2. Traverse the original array once.
3. Keep updating the maximum value seen so far.
4. Store the GCD between the current number and the running maximum.
5. Sort the `prefixGcd` array.
6. Use two pointers.
   - One starts from the beginning.
   - The other starts from the end.
7. Compute the GCD of every pair.
8. Add each GCD to the final answer.
9. Stop when both pointers meet or cross.

Since sorting is the most expensive operation, it dominates the overall time complexity.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Array / Vector | Stores the generated `prefixGcd` values |
| Two Pointers | Efficiently pairs the smallest and largest values after sorting |
| Integer Variables | Store the running prefix maximum and the final answer |

No advanced data structures like heaps, trees, or hash maps are needed for this problem.

---

## Operations & Behavior Summary

The algorithm works in four simple stages.

1. Traverse the original array.
2. Maintain the running prefix maximum.
3. Build the `prefixGcd` array using the Euclidean GCD algorithm.
4. Sort the array.
5. Pair the smallest and largest remaining values.
6. Compute the GCD for every pair.
7. Add every result to the answer.
8. Ignore the middle element automatically if the array length is odd.
9. Return the final sum.

---

## Complexity

| Operation | Complexity | Explanation |
|-----------|------------|-------------|
| Time Complexity | **O(n log n)** | Building the array takes `O(n)`, sorting takes `O(n log n)`, and pairing takes `O(n)`. Sorting is the dominant operation. |
| Space Complexity | **O(n)** | An additional array is used to store the computed `prefixGcd` values. |

Where:

- `n` = number of elements in `nums`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    long long gcdSum(vector<int>& nums) {
        int n = nums.size();

        // Store the gcd values for every prefix
        vector<int> prefixGcd(n);

        // Running maximum of the prefix
        int prefixMax = 0;

        // Build the prefixGcd array
        for (int i = 0; i < n; i++) {
            prefixMax = max(prefixMax, nums[i]);
            prefixGcd[i] = gcd(nums[i], prefixMax);
        }

        // Sort so that smallest and largest can be paired
        sort(prefixGcd.begin(), prefixGcd.end());

        long long ans = 0;

        // Pair smallest with largest
        int left = 0;
        int right = n - 1;

        while (left < right) {
            // Add gcd of the current pair
            ans += gcd(prefixGcd[left], prefixGcd[right]);

            left++;
            right--;
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    // Euclidean algorithm for gcd
    private int gcd(int a, int b) {
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    public long gcdSum(int[] nums) {
        int n = nums.length;

        // Store prefix gcd values
        int[] prefixGcd = new int[n];

        // Running prefix maximum
        int prefixMax = 0;

        // Build prefixGcd
        for (int i = 0; i < n; i++) {
            prefixMax = Math.max(prefixMax, nums[i]);
            prefixGcd[i] = gcd(nums[i], prefixMax);
        }

        // Sort the array
        java.util.Arrays.sort(prefixGcd);

        long ans = 0;

        // Pair smallest with largest
        int left = 0;
        int right = n - 1;

        while (left < right) {
            ans += gcd(prefixGcd[left], prefixGcd[right]);
            left++;
            right--;
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
var gcdSum = function(nums) {

    // Euclidean algorithm
    const gcd = (a, b) => {
        while (b !== 0) {
            let t = a % b;
            a = b;
            b = t;
        }
        return a;
    };

    const n = nums.length;

    // Store prefix gcd values
    const prefixGcd = new Array(n);

    // Running prefix maximum
    let prefixMax = 0;

    // Build prefixGcd
    for (let i = 0; i < n; i++) {
        prefixMax = Math.max(prefixMax, nums[i]);
        prefixGcd[i] = gcd(nums[i], prefixMax);
    }

    // Sort in ascending order
    prefixGcd.sort((a, b) => a - b);

    let ans = 0;

    // Pair smallest with largest
    let left = 0;
    let right = n - 1;

    while (left < right) {
        ans += gcd(prefixGcd[left], prefixGcd[right]);
        left++;
        right--;
    }

    return ans;
};
```

### Python3

```python
from math import gcd

class Solution:
    def gcdSum(self, nums: list[int]) -> int:
        n = len(nums)

        # Store prefix gcd values
        prefix_gcd = []

        # Running prefix maximum
        prefix_max = 0

        # Build prefix_gcd
        for x in nums:
            prefix_max = max(prefix_max, x)
            prefix_gcd.append(gcd(x, prefix_max))

        # Sort the array
        prefix_gcd.sort()

        ans = 0

        # Pair smallest with largest
        left = 0
        right = n - 1

        while left < right:
            ans += gcd(prefix_gcd[left], prefix_gcd[right])
            left += 1
            right -= 1

        return ans
```

### Go

```go
import "sort"

// Euclidean algorithm for gcd
func gcd(a, b int) int {
 for b != 0 {
  a, b = b, a%b
 }
 return a
}

func gcdSum(nums []int) int64 {

 n := len(nums)

 // Store prefix gcd values
 prefixGcd := make([]int, n)

 // Running prefix maximum
 prefixMax := 0

 // Build prefixGcd
 for i, x := range nums {
  if x > prefixMax {
   prefixMax = x
  }
  prefixGcd[i] = gcd(x, prefixMax)
 }

 // Sort the array
 sort.Ints(prefixGcd)

 var ans int64 = 0

 // Pair smallest with largest
 left := 0
 right := n - 1

 for left < right {
  ans += int64(gcd(prefixGcd[left], prefixGcd[right]))
  left++
  right--
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in every language. Only the syntax changes.

The first step is building the `prefixGcd` array.

Instead of searching for the maximum element from the beginning every time, I maintain one running maximum while moving through the array. This makes every update happen in constant time.

After updating the running maximum, I immediately calculate the GCD between the current number and that maximum. The computed value is stored inside the new array.

Once every value has been processed, the new array is complete.

Next, I sort the array in ascending order.

Sorting is important because the problem requires pairing the smallest unused value with the largest unused value. Once the array is sorted, those elements are always located at opposite ends.

To avoid removing elements repeatedly, I use two pointers.

The left pointer starts at the beginning.

The right pointer starts at the end.

During every iteration:

- Compute the GCD of both values.
- Add it to the answer.
- Move both pointers toward the center.

The process continues until both pointers meet.

If the array contains an odd number of elements, one value naturally remains in the middle. Since the loop stops before processing it, no extra condition is needed to ignore that element.

Although the implementation syntax differs slightly across C++, Java, JavaScript, Python3, and Go, every version performs the exact same sequence of operations.

---

## Examples

### Example 1

**Input**

```
nums = [2,6,4]
```

**Output**

```
2
```

**Trace**

```
prefixGcd = [2,6,2]

Sorted:

[2,2,6]

Pairs:

(2,6)

GCD = 2

Answer = 2
```

---

### Example 2

**Input**

```
nums = [3,6,2,8]
```

**Output**

```
5
```

**Trace**

```
prefixGcd = [3,6,2,8]

Sorted:

[2,3,6,8]

Pairs:

(2,8) -> 2
(3,6) -> 3

Total = 5
```

---

### Example 3

**Input**

```
nums = [5]
```

**Output**

```
0
```

**Trace**

```
Only one element exists.

No pair can be formed.

Answer = 0
```

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -std=c++17 -O2
```

Run

```bash
./a.out
```

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

Build

```bash
go build solution.go
```

---

## Notes & Optimizations

- The Euclidean Algorithm computes the Greatest Common Divisor efficiently.
- Maintaining a running prefix maximum avoids repeatedly scanning previous elements.
- The two-pointer technique is the simplest way to follow the pairing rule after sorting.
- Sorting contributes the largest portion of the running time.
- The algorithm automatically handles arrays with odd lengths because the middle element is never paired.
- Using a 64-bit integer for the answer is important since the total sum can become larger than a standard 32-bit integer.
- This approach is both simple to understand and efficient enough for the maximum constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
