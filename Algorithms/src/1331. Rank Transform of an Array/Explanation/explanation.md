# 1331. Rank Transform of an Array | LeetCode Solution (C++, Java, JavaScript, Python, Go)

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

The **Rank Transform of an Array** problem asks us to replace every element in an integer array with its rank.

The smallest unique value should receive rank `1`. Larger unique values receive larger ranks, and duplicate values must always share the same rank.

The original order of the array must remain unchanged. Only the values are replaced with their corresponding ranks.

This problem is a good example of combining **sorting**, **hash maps**, and **array manipulation** to build an efficient solution.

---

## Constraints

| Constraint    | Value                     |
| ------------- | ------------------------- |
| Array Length  | `0 <= arr.length <= 10^5` |
| Element Value | `-10^9 <= arr[i] <= 10^9` |

---

## Intuition

The first thing I noticed was that the rank depends only on the relative order of unique values.

Instead of trying to rank elements while keeping their original positions, I realized it would be much easier to sort the values first. Once the unique numbers are sorted, assigning ranks becomes straightforward.

After that, I can simply store every value and its rank in a hash map. Then I only need one more pass through the original array to replace every number with its assigned rank.

This approach is simple, efficient, and works well even for very large arrays.

---

## Approach

I solved the problem in four simple steps.

1. Create a copy of the original array.
2. Sort the copied array in ascending order.
3. Traverse the sorted array and assign ranks only to unique values using a hash map.
4. Traverse the original array and replace every element with the rank stored in the hash map.

Since hash map lookups are very fast, the final replacement step takes linear time.

---

## Data Structures Used

### Array

I create a copy of the original array so I can sort it without changing the original order.

### Hash Map

The hash map stores every unique value along with its assigned rank.

Example:

```
10 → 1
20 → 2
30 → 3
40 → 4
```

Using a hash map allows constant-time rank lookup while rebuilding the final answer.

---

## Operations & Behavior Summary

The algorithm performs the following operations:

1. Copy the original array.
2. Sort the copied array.
3. Visit every value in sorted order.
4. Assign a new rank only if the value has not been seen before.
5. Store the value-to-rank mapping.
6. Traverse the original array again.
7. Replace every value with its stored rank.
8. Return the transformed array.

---

## Complexity

| Complexity       | Value          | Explanation                                                                                    |
| ---------------- | -------------- | ---------------------------------------------------------------------------------------------- |
| Time Complexity  | **O(n log n)** | Sorting takes `O(n log n)`, while building the rank map and replacing values each take `O(n)`. |
| Space Complexity | **O(n)**       | Extra space is used for the copied array and the hash map.                                     |

Where:

- `n` is the number of elements in the array.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> arrayRankTransform(vector<int>& arr) {
        // Create a copy so the original order remains unchanged
        vector<int> sorted = arr;

        // Sort the copied array
        sort(sorted.begin(), sorted.end());

        // Store each unique value with its rank
        unordered_map<int, int> rank;
        int currentRank = 1;

        // Assign ranks only to unique values
        for (int num : sorted) {
            if (!rank.count(num)) {
                rank[num] = currentRank++;
            }
        }

        // Replace every element with its assigned rank
        for (int &num : arr) {
            num = rank[num];
        }

        // Return the transformed array
        return arr;
    }
};
```

### Java

```java
class Solution {
    public int[] arrayRankTransform(int[] arr) {
        // Create a copy of the original array
        int[] sorted = arr.clone();

        // Sort the copied array
        Arrays.sort(sorted);

        // Store each unique value and its rank
        HashMap<Integer, Integer> rank = new HashMap<>();
        int currentRank = 1;

        // Assign ranks to unique values only
        for (int num : sorted) {
            if (!rank.containsKey(num)) {
                rank.put(num, currentRank++);
            }
        }

        // Replace every element with its rank
        for (int i = 0; i < arr.length; i++) {
            arr[i] = rank.get(arr[i]);
        }

        // Return the transformed array
        return arr;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @return {number[]}
 */
var arrayRankTransform = function (arr) {
  // Create a copy of the original array
  const sorted = [...arr];

  // Sort the copied array in increasing order
  sorted.sort((a, b) => a - b);

  // Store value -> rank
  const rank = new Map();
  let currentRank = 1;

  // Assign ranks only once for every unique value
  for (const num of sorted) {
    if (!rank.has(num)) {
      rank.set(num, currentRank++);
    }
  }

  // Replace every element with its rank
  for (let i = 0; i < arr.length; i++) {
    arr[i] = rank.get(arr[i]);
  }

  // Return the transformed array
  return arr;
};
```

### Python3

```python
class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        # Create a sorted copy of the array
        sorted_arr = sorted(arr)

        # Store value -> rank
        rank = {}
        current_rank = 1

        # Assign ranks only to unique values
        for num in sorted_arr:
            if num not in rank:
                rank[num] = current_rank
                current_rank += 1

        # Replace every element with its assigned rank
        return [rank[num] for num in arr]
```

### Go

```go
func arrayRankTransform(arr []int) []int {
 // Create a copy so the original order is preserved
 sorted := append([]int(nil), arr...)

 // Sort the copied array
 sort.Ints(sorted)

 // Store value -> rank
 rank := make(map[int]int)
 currentRank := 1

 // Assign ranks only to unique values
 for _, num := range sorted {
  if _, exists := rank[num]; !exists {
   rank[num] = currentRank
   currentRank++
  }
 }

 // Replace every element with its assigned rank
 for i, num := range arr {
  arr[i] = rank[num]
 }

 // Return the transformed array
 return arr
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is exactly the same in all five languages. Only the syntax changes.

First, I create another array that contains the same values as the original one. This allows me to sort the numbers without losing the original order.

Next, I sort the copied array from smallest to largest.

After sorting, I iterate through the array only once. Every time I find a value that has not appeared before, I assign the next available rank and store it in a hash map.

Duplicate values are ignored because they should all receive the same rank.

Once every unique number has been assigned a rank, I go back to the original array.

For each element, I simply look up its rank inside the hash map and replace the original value.

Since every lookup happens in constant time on average, rebuilding the answer is very efficient.

The behavior is identical in C++, Java, JavaScript, Python3, and Go. Only the built-in syntax for sorting, maps, and loops is different.

---

## Examples

### Example 1

**Input**

```
arr = [40,10,20,30]
```

**Output**

```
[4,1,2,3]
```

**Trace**

```
Sorted Values
10 20 30 40

Assigned Ranks
10 → 1
20 → 2
30 → 3
40 → 4

Final Answer
4 1 2 3
```

---

### Example 2

**Input**

```
arr = [100,100,100]
```

**Output**

```
[1,1,1]
```

**Trace**

```
Sorted Values
100 100 100

Unique Values
100

Assigned Rank
100 → 1

Final Answer
1 1 1
```

---

### Example 3

**Input**

```
arr = [37,12,28,9,100,56,80,5,12]
```

**Output**

```
[5,3,4,2,8,6,7,1,3]
```

**Trace**

```
Sorted Unique Values

5
9
12
28
37
56
80
100

Ranks

5 → 1
9 → 2
12 → 3
28 → 4
37 → 5
56 → 6
80 → 7
100 → 8

Replace every value using the stored ranks.
```

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ main.cpp -o main
```

Run

```bash
./main
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

- Duplicate values must always receive the same rank.
- Sorting only the copied array keeps the original order untouched.
- Using a hash map makes rank lookup very fast.
- This is one of the most common and efficient approaches for solving the Rank Transform of an Array problem.
- The sorting step is unavoidable because ranks depend on the global ordering of unique values.
- The solution handles empty arrays naturally without any extra conditions.
- This approach is suitable for the maximum constraints given in the problem.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
