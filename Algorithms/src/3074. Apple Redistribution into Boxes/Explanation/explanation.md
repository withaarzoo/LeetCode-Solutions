# 🍎 Apple Redistribution into Boxes (LeetCode 3074)

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given two arrays:

* `apple[]` → each value represents the number of apples in a pack
* `capacity[]` → each value represents how many apples a box can hold

My task is to find the **minimum number of boxes** needed to store **all apples**.

Important rule:

* Apples from one pack **can be split** into multiple boxes.

---

## Constraints

* `1 ≤ apple.length ≤ 50`
* `1 ≤ capacity.length ≤ 50`
* `1 ≤ apple[i], capacity[i] ≤ 50`
* It is always guaranteed that redistribution is possible.

---

## Intuition

When I first read the problem, I stopped thinking about individual apple packs.

Because apples **can be split**, I realized:

* Only the **total number of apples** matters.
* I do **not** care which apple goes into which box.

To minimize the number of boxes:

* I should use the **largest capacity boxes first**.

This immediately tells me that a **greedy approach** will work.

---

## Approach

Here is how I solved it step by step:

1. Calculate the **total number of apples**.
2. Sort the `capacity` array in **descending order**.
3. Start picking boxes from the largest capacity.
4. Keep adding capacities until the total capacity is enough.
5. Count how many boxes I used — that is my answer.

---

## Data Structures Used

* Arrays / Lists
* No extra data structures are required.

---

## Operations & Behavior Summary

* Sum all apples → total apples
* Sort box capacities in descending order
* Greedily select boxes until capacity ≥ total apples
* Return the number of boxes used

---

## Complexity

* **Time Complexity:** `O(m log m)`

  * `m` = number of boxes
  * Sorting the capacity array dominates the time

* **Space Complexity:** `O(1)`

  * Only a few variables are used (excluding sort space)

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumBoxes(vector<int>& apple, vector<int>& capacity) {
        int totalApples = 0;
        for (int a : apple) totalApples += a;

        sort(capacity.begin(), capacity.end(), greater<int>());

        int usedCapacity = 0, boxes = 0;
        for (int cap : capacity) {
            usedCapacity += cap;
            boxes++;
            if (usedCapacity >= totalApples) return boxes;
        }
        return boxes;
    }
};
```

---

### Java

```java
class Solution {
    public int minimumBoxes(int[] apple, int[] capacity) {
        int totalApples = 0;
        for (int a : apple) totalApples += a;

        Arrays.sort(capacity);

        int usedCapacity = 0, boxes = 0;
        for (int i = capacity.length - 1; i >= 0; i--) {
            usedCapacity += capacity[i];
            boxes++;
            if (usedCapacity >= totalApples) return boxes;
        }
        return boxes;
    }
}
```

---

### JavaScript

```javascript
var minimumBoxes = function(apple, capacity) {
    let totalApples = apple.reduce((sum, a) => sum + a, 0);

    capacity.sort((a, b) => b - a);

    let usedCapacity = 0;
    let boxes = 0;

    for (let cap of capacity) {
        usedCapacity += cap;
        boxes++;
        if (usedCapacity >= totalApples) return boxes;
    }
    return boxes;
};
```

---

### Python3

```python
class Solution:
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        total_apples = sum(apple)
        capacity.sort(reverse=True)

        used_capacity = 0
        boxes = 0

        for cap in capacity:
            used_capacity += cap
            boxes += 1
            if used_capacity >= total_apples:
                return boxes
        return boxes
```

---

### Go

```go
func minimumBoxes(apple []int, capacity []int) int {
    totalApples := 0
    for _, a := range apple {
        totalApples += a
    }

    sort.Slice(capacity, func(i, j int) bool {
        return capacity[i] > capacity[j]
    })

    usedCapacity := 0
    boxes := 0
    for _, cap := range capacity {
        usedCapacity += cap
        boxes++
        if usedCapacity >= totalApples {
            return boxes
        }
    }
    return boxes
}
```

---

## Step-by-step Detailed Explanation

1. I add all apple packs to get the **total apples**.
2. I sort the box capacities from **largest to smallest**.
3. I start using boxes one by one.
4. After each box, I check if my capacity is enough.
5. As soon as capacity ≥ total apples, I stop.
6. The count of boxes used is the minimum possible.

---

## Examples

### Example 1

```
Input:
apple = [1,3,2]
capacity = [4,3,1,5,2]

Output:
2
```

Explanation: Using boxes with capacity 5 and 4 is enough.

---

### Example 2

```
Input:
apple = [5,5,5]
capacity = [2,4,2,7]

Output:
4
```

Explanation: All boxes are required.

---

## How to Use / Run Locally

1. Copy the solution in your preferred language.
2. Paste it into LeetCode or your local compiler.
3. Run with custom test cases.

---

## Notes & Optimizations

* This greedy approach is optimal because apples can be freely distributed.
* Sorting in descending order ensures minimum boxes.
* No complex data structures are required.

---

## Author

**Md Aarzoo Islam**

* 🔗 [Portfolio / Links](https://bento.me/withaarzoo)
