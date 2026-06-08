# 2161. Partition Array According to Given Pivot

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

Given an integer array `nums` and an integer `pivot`, the goal is to rearrange the array according to three rules:

1. All elements smaller than `pivot` must appear before elements greater than `pivot`.
2. All elements equal to `pivot` must appear between the smaller and greater groups.
3. The relative order of elements inside the smaller group and inside the greater group must remain unchanged.

The task is to return the newly arranged array.

This is a classic Array and Two Pointers style problem that focuses on stable partitioning while preserving ordering.

---

## Constraints

| Constraint                   | Value               |
| ---------------------------- | ------------------- |
| `1 <= nums.length <= 100000` | Array size          |
| `-10^6 <= nums[i] <= 10^6`   | Element value range |
| `pivot` exists in `nums`     | Guaranteed          |

---

## Intuition

The first thing I noticed was that the problem does not ask me to sort the array.

Instead, it only wants three groups:

* Values smaller than the pivot
* Values equal to the pivot
* Values greater than the pivot

The order inside each group must stay the same.

That immediately suggests collecting elements into separate containers while scanning the array from left to right. Since elements are inserted in their original order, the required stability is automatically preserved.

After collecting all three groups, I can simply join them together.

---

## Approach

1. Create three separate containers.

   * One for values smaller than the pivot.
   * One for values equal to the pivot.
   * One for values greater than the pivot.

2. Traverse the array once.

3. For each element:

   * Put it into the appropriate container.

4. Combine the containers in this order:

   * Smaller elements
   * Equal elements
   * Greater elements

5. Return the final array.

This guarantees that all conditions from the problem statement are satisfied.

---

## Data Structures Used

### Dynamic Array / List

Used to store:

* Elements smaller than the pivot
* Elements equal to the pivot
* Elements greater than the pivot

Why?

* Fast insertion at the end
* Preserves original ordering
* Easy to combine into the final answer

---

## Operations & Behavior Summary

The algorithm performs the following actions:

1. Start with three empty collections.
2. Read each number exactly once.
3. Compare the number with the pivot.
4. Place the number into one of three groups.
5. After traversal:

   * Append the equal group after the smaller group.
   * Append the greater group after the equal group.
6. Return the resulting sequence.

This creates a stable partition of the array.

---

## Complexity

| Metric           | Complexity | Explanation                                |
| ---------------- | ---------- | ------------------------------------------ |
| Time Complexity  | O(n)       | Each element is processed exactly once     |
| Space Complexity | O(n)       | Extra storage is used for the three groups |

Where:

* `n` = length of the input array `nums`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> pivotArray(vector<int>& nums, int pivot) {
        // Store elements smaller than pivot
        vector<int> smaller;

        // Store elements equal to pivot
        vector<int> equal;

        // Store elements greater than pivot
        vector<int> greater;

        // Classify every element into one of three groups
        for (int num : nums) {
            if (num < pivot) {
                smaller.push_back(num);
            }
            else if (num == pivot) {
                equal.push_back(num);
            }
            else {
                greater.push_back(num);
            }
        }

        // Append equal elements after smaller elements
        smaller.insert(smaller.end(), equal.begin(), equal.end());

        // Append greater elements at the end
        smaller.insert(smaller.end(), greater.begin(), greater.end());

        // Return the final partitioned array
        return smaller;
    }
};
```

### Java

```java
class Solution {
    public int[] pivotArray(int[] nums, int pivot) {

        // Lists to store three groups
        List<Integer> smaller = new ArrayList<>();
        List<Integer> equal = new ArrayList<>();
        List<Integer> greater = new ArrayList<>();

        // Classify each element
        for (int num : nums) {
            if (num < pivot) {
                smaller.add(num);
            } else if (num == pivot) {
                equal.add(num);
            } else {
                greater.add(num);
            }
        }

        // Result array of same size
        int[] result = new int[nums.length];
        int index = 0;

        // Add smaller elements
        for (int num : smaller) {
            result[index++] = num;
        }

        // Add equal elements
        for (int num : equal) {
            result[index++] = num;
        }

        // Add greater elements
        for (int num : greater) {
            result[index++] = num;
        }

        // Return final answer
        return result;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} pivot
 * @return {number[]}
 */
var pivotArray = function(nums, pivot) {

    // Store elements smaller than pivot
    const smaller = [];

    // Store elements equal to pivot
    const equal = [];

    // Store elements greater than pivot
    const greater = [];

    // Classify every element
    for (const num of nums) {
        if (num < pivot) {
            smaller.push(num);
        } else if (num === pivot) {
            equal.push(num);
        } else {
            greater.push(num);
        }
    }

    // Combine all three groups in required order
    return [...smaller, ...equal, ...greater];
};
```

### Python3

```python
class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:

        # Store elements smaller than pivot
        smaller = []

        # Store elements equal to pivot
        equal = []

        # Store elements greater than pivot
        greater = []

        # Classify each element
        for num in nums:
            if num < pivot:
                smaller.append(num)
            elif num == pivot:
                equal.append(num)
            else:
                greater.append(num)

        # Return all groups in required order
        return smaller + equal + greater
```

### Go

```go
func pivotArray(nums []int, pivot int) []int {

    // Store elements smaller than pivot
    smaller := []int{}

    // Store elements equal to pivot
    equal := []int{}

    // Store elements greater than pivot
    greater := []int{}

    // Classify every element
    for _, num := range nums {
        if num < pivot {
            smaller = append(smaller, num)
        } else if num == pivot {
            equal = append(equal, num)
        } else {
            greater = append(greater, num)
        }
    }

    // Build final answer
    result := []int{}

    result = append(result, smaller...)
    result = append(result, equal...)
    result = append(result, greater...)

    return result
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Step 1: Create Three Groups

I begin by creating three separate collections.

One collection stores values less than the pivot.

One collection stores values equal to the pivot.

One collection stores values greater than the pivot.

Keeping them separate makes the solution easy to understand and guarantees stability.

---

### Step 2: Traverse the Array

I scan every element from left to right.

This ensures I see elements in their original order.

Since I never rearrange elements inside a group, the original ordering is preserved automatically.

---

### Step 3: Compare with Pivot

For every number:

* If it is smaller than the pivot, place it in the smaller group.
* If it is equal to the pivot, place it in the equal group.
* If it is greater than the pivot, place it in the greater group.

Every element belongs to exactly one of these categories.

---

### Step 4: Merge the Groups

Once the traversal is complete:

1. Add all smaller elements first.
2. Add all pivot elements next.
3. Add all greater elements last.

This arrangement exactly matches the required output format.

---

### Step 5: Return the Result

The combined collection becomes the final answer.

At this point:

* All smaller elements are on the left.
* All pivot values are in the middle.
* All greater elements are on the right.
* Relative ordering remains unchanged.

---

### Why Stability Matters

Consider:

Input:

```text
[9, 5, 3]
```

All three values are smaller than the pivot.

The output must remain:

```text
[9, 5, 3]
```

and not:

```text
[3, 5, 9]
```

The same rule applies to the greater-than-pivot section.

Because the algorithm processes elements from left to right and appends them directly into groups, stability is preserved naturally.

---

## Examples

### Example 1

Input

```text
nums = [9,12,5,10,14,3,10]
pivot = 10
```

Processing

```text
smaller = [9,5,3]
equal   = [10,10]
greater = [12,14]
```

Output

```text
[9,5,3,10,10,12,14]
```

---

### Example 2

Input

```text
nums = [-3,4,3,2]
pivot = 2
```

Processing

```text
smaller = [-3]
equal   = [2]
greater = [4,3]
```

Output

```text
[-3,2,4,3]
```

---

### Example 3

Input

```text
nums = [5,5,5,5]
pivot = 5
```

Processing

```text
smaller = []
equal   = [5,5,5,5]
greater = []
```

Output

```text
[5,5,5,5]
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

---

### Go

Run

```bash
go run main.go
```

---

## Notes & Optimizations

* This solution is already optimal in terms of time complexity.
* Every element is processed only once.
* Stable ordering is preserved automatically.
* Using three separate containers keeps the implementation simple and readable.
* An in-place solution would be much harder because preserving relative order is required.
* For competitive programming and interview settings, this approach is usually the cleanest and most reliable choice.

### Edge Cases

* All elements are equal to the pivot.
* All elements except the pivot are smaller.
* All elements except the pivot are greater.
* Negative numbers are present.
* Duplicate values appear multiple times.

The algorithm handles all of these cases naturally.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
