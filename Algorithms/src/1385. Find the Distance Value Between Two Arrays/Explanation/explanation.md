# Find the Distance Value Between Two Arrays - LeetCode 1385 Solution

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

You are given two integer arrays called arr1 and arr2, along with an integer d. Your task is to find the distance value between these two arrays.

The distance value is simply the count of elements in arr1 that have no element in arr2 within a distance of d. In other words, for every number x in arr1, check if there exists any number y in arr2 such that the absolute difference |x - y| is less than or equal to d. If no such y exists, then x contributes 1 to the answer.

This is a classic LeetCode easy problem that tests your understanding of searching and sorting techniques in DSA.

## Constraints

- 1 <= arr1.length, arr2.length <= 500
- -1000 <= arr1[i], arr2[j] <= 1000
- 0 <= d <= 100

## Intuition

When I first read the problem, I realized a brute-force double loop would work because the array sizes are small. But the real insight came from noticing that for every element in arr1 I only care about the closest elements in arr2. If I sort arr2, I can use binary search to quickly find the nearest values and check only those. This turns an O(n * m) solution into a much cleaner O(n log m) approach, which is the standard way to solve the find the distance value between two arrays problem efficiently.

## Approach

First, sort the second array so that its values are in increasing order.  
Then, for every number in the first array, perform a binary search on the sorted second array to locate the position where that number would be inserted.  
Look at the element at that position (if it exists) and the element just before it (if it exists). These two are the closest candidates.  
If either of them lies within distance d, the current number from arr1 is invalid.  
If neither is close enough, increase the answer by one.  
After checking every element of arr1, return the final count.

This binary search approach is the most common and efficient way to solve LeetCode 1385.

## Data Structures Used

- Arrays: The input arrays themselves. No extra arrays are created.
- Binary search (via lower_bound or manual implementation): Used to locate the closest elements quickly after sorting.

Sorting is done in-place (or with language-provided temporary space), so we avoid heavy extra memory.

## Operations & Behavior Summary

1. Sort arr2 in ascending order.  
2. Initialize a counter to zero.  
3. For each value x in arr1:  
   - Binary search to find the first index in arr2 that is greater than or equal to x.  
   - Check the value at that index (if within bounds).  
   - Check the value immediately before that index (if it exists).  
   - If neither value is within distance d of x, increment the counter.  
4. Return the counter.

## Complexity

| Type              | Complexity          | Explanation |
|-------------------|---------------------|-------------|
| Time Complexity   | O(m log m + n log m) | Sorting arr2 takes O(m log m). Each of the n elements in arr1 needs one binary search of O(log m). |
| Space Complexity  | O(1) or O(m)        | Extra space is only what the sorting algorithm needs. No additional data structures of our own are allocated. |

Here n = length of arr1 and m = length of arr2.

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    int findTheDistanceValue(vector<int>& arr1, vector<int>& arr2, int d) {
        sort(arr2.begin(), arr2.end());
        int ans = 0;
        for (int x : arr1) {
            auto it = lower_bound(arr2.begin(), arr2.end(), x);
            bool ok = true;
            if (it != arr2.end() && abs(*it - x) <= d) ok = false;
            if (it != arr2.begin() && abs(*(it - 1) - x) <= d) ok = false;
            if (ok) ans++;
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int findTheDistanceValue(int[] arr1, int[] arr2, int d) {
        Arrays.sort(arr2);
        int ans = 0;
        for (int x : arr1) {
            int idx = Arrays.binarySearch(arr2, x);
            if (idx < 0) idx = -idx - 1;
            boolean ok = true;
            if (idx < arr2.length && Math.abs(arr2[idx] - x) <= d) ok = false;
            if (idx > 0 && Math.abs(arr2[idx - 1] - x) <= d) ok = false;
            if (ok) ans++;
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @param {number} d
 * @return {number}
 */
var findTheDistanceValue = function(arr1, arr2, d) {
    arr2.sort((a, b) => a - b);
    let ans = 0;
    for (let x of arr1) {
        let left = 0, right = arr2.length;
        while (left < right) {
            let mid = (left + right) >> 1;
            if (arr2[mid] < x) left = mid + 1;
            else right = mid;
        }
        let ok = true;
        if (left < arr2.length && Math.abs(arr2[left] - x) <= d) ok = false;
        if (left > 0 && Math.abs(arr2[left - 1] - x) <= d) ok = false;
        if (ok) ans++;
    }
    return ans;
};
```

### TypeScript
```typescript
function findTheDistanceValue(arr1: number[], arr2: number[], d: number): number {
    arr2.sort((a, b) => a - b);
    let ans = 0;
    for (let x of arr1) {
        let left = 0, right = arr2.length;
        while (left < right) {
            let mid = (left + right) >> 1;
            if (arr2[mid] < x) left = mid + 1;
            else right = mid;
        }
        let ok = true;
        if (left < arr2.length && Math.abs(arr2[left] - x) <= d) ok = false;
        if (left > 0 && Math.abs(arr2[left - 1] - x) <= d) ok = false;
        if (ok) ans++;
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        arr2.sort()
        ans = 0
        for x in arr1:
            left, right = 0, len(arr2)
            while left < right:
                mid = (left + right) // 2
                if arr2[mid] < x:
                    left = mid + 1
                else:
                    right = mid
            ok = True
            if left < len(arr2) and abs(arr2[left] - x) <= d:
                ok = False
            if left > 0 and abs(arr2[left - 1] - x) <= d:
                ok = False
            if ok:
                ans += 1
        return ans
```

### Go
```go
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    sort.Ints(arr2)
    ans := 0
    for _, x := range arr1 {
        left, right := 0, len(arr2)
        for left < right {
            mid := (left + right) / 2
            if arr2[mid] < x {
                left = mid + 1
            } else {
                right = mid
            }
        }
        ok := true
        if left < len(arr2) && abs(arr2[left]-x) <= d {
            ok = false
        }
        if left > 0 && abs(arr2[left-1]-x) <= d {
            ok = false
        }
        if ok {
            ans++
        }
    }
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea stays the same across all six languages. We first sort the second array. After sorting, binary search becomes reliable because the values are ordered.

For every number x taken from arr1 we locate the insertion point in the sorted arr2. This insertion point tells us exactly where the closest larger-or-equal value sits. We then also look one step to the left to catch the closest smaller value. Checking only these two neighbors is enough; any farther element would be even more distant.

If both neighbors (when they exist) have an absolute difference greater than d, we safely count x. Otherwise we ignore it.

In C++ we use the standard lower_bound function which returns an iterator.  
In Java we use Arrays.binarySearch and convert a negative insertion point into a proper index.  
In JavaScript, TypeScript, Python and Go we write a short manual binary search that finds the leftmost position where the value is not smaller than x. The rest of the logic is identical: inspect the two candidates and decide whether to count the current x.

Edge cases such as empty neighbors (when x is smaller than everything or larger than everything) are handled naturally by the boundary checks. Because array lengths are at most 500, the solution stays fast even with the logarithmic factor.

## Examples

**Example 1**  
Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2  
Output: 2  

Trace:  
- 4 → closest values in sorted arr2 [1,8,9,10] are 1 and 8. Both differences > 2 → count it.  
- 5 → same neighbors, both differences > 2 → count it.  
- 8 → 8 itself and 9 are within distance 2 → skip.  
Final answer is 2.

**Example 2**  
Input: arr1 = [1,4,2,3], arr2 = [-4,-3,6,10,20,30], d = 3  
Output: 2  

Trace:  
Sorted arr2 = [-4,-3,6,10,20,30].  
1 and 2 have no close neighbors within 3, while 3 and 4 are close to 6. Answer becomes 2.

**Example 3**  
Input: arr1 = [2,1,100,3], arr2 = [-5,-2,10,-3,7], d = 6  
Output: 1  

Only 100 stays far from every value in arr2.

## How to Use / Run Locally

1. Copy the solution code for your preferred language into a file (for example Solution.cpp, Solution.java, etc.).  
2. Add a simple main / driver function that creates the two arrays and the value of d, then calls the method and prints the result.  
3. Compile and run:  
   - C++: `g++ Solution.cpp -o solution && ./solution`  
   - Java: `javac Solution.java && java Solution`  
   - JavaScript: `node solution.js`  
   - TypeScript: `tsc solution.ts && node solution.js`  
   - Python: `python3 solution.py`  
   - Go: `go run solution.go`  

Replace the hard-coded arrays with any test case you want to try.

## Notes & Optimizations

A pure double loop also passes because of the tiny constraints, but the binary-search version is cleaner and teaches a useful pattern that appears in many harder problems.  

If the arrays were much larger, the same approach would still scale well.  

You can also sort both arrays and use two pointers, but binary search is simpler to implement and equally efficient for this problem size.

Watch out for the absolute difference calculation and for the boundary cases when the binary search lands at the very beginning or the very end of arr2.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
