# Finding Minimum Difference in a Sorted Array (C++, Java, JavaScript, Python, Go)

This code calculates the minimum difference you can achieve by removing elements from a sorted array.

Here's a step-by-step explanation with code snippets in each language:

## 1. Define the Function and Input

Each code snippet defines a function named `minDifference` that takes an integer array `nums` as input.

**C++:**

```c++
class Solution {
public:
    int minDifference(vector<int>& nums) {
        // ...
    }
};
```

**Java:**

```java
class Solution {
    public int minDifference(int[] nums) {
        // ...
    }
}
```

**JavaScript:**

```javascript
var minDifference = function(nums) {
    // ...
};
```

**Python:**

```python
class Solution:
    def minDifference(self, nums: List[int]) -> int:
        // ...
    }
```

**Go:**

```go
func minDifference(nums []int) int {
    // ...
}
```

## 2. Check for Base Case

The code first checks if the array size (`n`) is less than or equal to 4. If so, the minimum difference is 0 (removing elements from such a small array won't make a difference), and the function returns 0.

**C++:**

```c++
int n = nums.size();
if (n <= 4) return 0;
```

**Java:**

```java
int n = nums.length;
if (n <= 4) return 0;
```

**JavaScript:**

```javascript
const n = nums.length;
if (n <= 4) return 0;
```

**Python:**

```python
n = len(nums)
if n <= 4:
    return 0
```

**Go:**

```go
n := len(nums)
if n <= 4 {
    return 0
}
```

## 3. Sort the Array

The code then sorts the input array `nums` in ascending order. Sorting is necessary to efficiently calculate the minimum difference later.

**C++:**

```c++
sort(nums.begin(), nums.end());
```

**Java:**

```java
Arrays.sort(nums);
```

**JavaScript:**

```javascript
nums.sort((a, b) => a - b);
```

**Python:**

```python
nums.sort()
```

**Go:**

```go
sort.Ints(nums)
```

## 4. Calculate and Return Minimum Difference

The core logic involves calculating the minimum difference achievable by removing elements from the sorted array. The code considers four scenarios:

1. Removing the 3 smallest elements.
2. Removing the 2 smallest and 1 largest element.
3. Removing the 1 smallest and 2 largest elements.
4. Removing the 3 largest elements.

It then uses a `min` function (built-in or custom) to find the minimum difference among these four scenarios and returns that value.

**C++:**

```c++
return min({
    nums[n - 1] - nums[3],
    nums[n - 2] - nums[2],
    nums[n - 3] - nums[1],
    nums[n - 4] - nums[0]
});
```

**Java:**

```java
return Math.min(
    Math.min(nums[n - 1] - nums[3], nums[n - 2] - nums[2]),
    Math.min(nums[n - 3] - nums[1], nums[n - 4] - nums[0])
);
```

**JavaScript:**

```javascript
return Math.min(
    nums[n - 1] - nums[3],
    nums[n - 2] - nums[2],
    nums[n - 3] - nums[1],
    nums[n - 4] - nums[0]
);
```

**Python:**

```python
return min(
    nums[n - 1] - nums[3],
    nums[n - 2] - nums[2],
    nums[n - 3] - nums[1],
    nums[n - 4] - nums[0]
);
```

**Go:**

```go
return min(
    nums[n-1]-nums[3],
    nums[n-2]-nums[2],
    nums
