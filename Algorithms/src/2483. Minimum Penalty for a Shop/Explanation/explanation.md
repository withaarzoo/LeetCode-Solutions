# Problem Title

**2483. Minimum Penalty for a Shop**

---

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

You are given a string `customers` where:

* `'Y'` means customers visited the shop at that hour
* `'N'` means no customers visited the shop at that hour

The shop can be closed at any hour `j` (0 ≤ j ≤ n).

Penalty rules:

* Shop **open + no customer (`N`)** → penalty +1
* Shop **closed + customer (`Y`)** → penalty +1

Your task is to return the **earliest hour** at which the shop should close to get the **minimum total penalty**.

---

## Constraints

* `1 ≤ customers.length ≤ 10^5`
* `customers` contains only `'Y'` and `'N'`

---

## Intuition

I thought about **why penalty happens**.

Penalty only occurs in two cases:

1. Shop is open and no customer comes (`N`)
2. Shop is closed and customers still come (`Y`)

If I close the shop at hour `j`:

* Before `j` → shop is open → count `'N'`
* From `j` onward → shop is closed → count `'Y'`

So the total penalty becomes:

```
(number of 'N' before j) + (number of 'Y' after j)
```

Now the problem becomes:

> Try every possible closing hour and pick the one with the **minimum penalty**.

---

## Approach

1. Count the total number of `'Y'` in the string.
2. Assume the shop closes at hour `0`.
3. Move hour by hour:

   * If the current hour has `'N'`, opening causes penalty.
   * If the current hour has `'Y'`, closing avoids penalty.
4. Track the minimum penalty and earliest hour.
5. Return the best closing time.

This approach works in **one pass**.

---

## Data Structures Used

* Primitive variables only (`int`, `string`)
* No extra arrays or hash maps

---

## Operations & Behavior Summary

* Single traversal of the string
* Update penalties dynamically
* Compare and store minimum penalty
* Prefer earlier hour when penalties are equal

---

## Complexity

* **Time Complexity:** `O(n)`
  Where `n` is the length of the customers string.

* **Space Complexity:** `O(1)`
  No additional memory used apart from variables.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int bestClosingTime(string customers) {
        int totalY = 0;
        for (char c : customers)
            if (c == 'Y') totalY++;

        int openPenalty = 0;
        int closedPenalty = totalY;
        int minPenalty = closedPenalty;
        int answer = 0;

        for (int i = 0; i < customers.size(); i++) {
            if (customers[i] == 'N')
                openPenalty++;
            else
                closedPenalty--;

            int currentPenalty = openPenalty + closedPenalty;
            if (currentPenalty < minPenalty) {
                minPenalty = currentPenalty;
                answer = i + 1;
            }
        }
        return answer;
    }
};
```

---

### Java

```java
class Solution {
    public int bestClosingTime(String customers) {
        int totalY = 0;
        for (char c : customers.toCharArray())
            if (c == 'Y') totalY++;

        int openPenalty = 0;
        int closedPenalty = totalY;
        int minPenalty = closedPenalty;
        int answer = 0;

        for (int i = 0; i < customers.length(); i++) {
            if (customers.charAt(i) == 'N')
                openPenalty++;
            else
                closedPenalty--;

            int currentPenalty = openPenalty + closedPenalty;
            if (currentPenalty < minPenalty) {
                minPenalty = currentPenalty;
                answer = i + 1;
            }
        }
        return answer;
    }
}
```

---

### JavaScript

```javascript
var bestClosingTime = function(customers) {
    let totalY = 0;
    for (let c of customers)
        if (c === 'Y') totalY++;

    let openPenalty = 0;
    let closedPenalty = totalY;
    let minPenalty = closedPenalty;
    let answer = 0;

    for (let i = 0; i < customers.length; i++) {
        if (customers[i] === 'N')
            openPenalty++;
        else
            closedPenalty--;

        let currentPenalty = openPenalty + closedPenalty;
        if (currentPenalty < minPenalty) {
            minPenalty = currentPenalty;
            answer = i + 1;
        }
    }
    return answer;
};
```

---

### Python3

```python
class Solution:
    def bestClosingTime(self, customers: str) -> int:
        totalY = customers.count('Y')

        openPenalty = 0
        closedPenalty = totalY
        minPenalty = closedPenalty
        answer = 0

        for i, c in enumerate(customers):
            if c == 'N':
                openPenalty += 1
            else:
                closedPenalty -= 1

            currentPenalty = openPenalty + closedPenalty
            if currentPenalty < minPenalty:
                minPenalty = currentPenalty
                answer = i + 1

        return answer
```

---

### Go

```go
func bestClosingTime(customers string) int {
    totalY := 0
    for _, c := range customers {
        if c == 'Y' {
            totalY++
        }
    }

    openPenalty := 0
    closedPenalty := totalY
    minPenalty := closedPenalty
    answer := 0

    for i, c := range customers {
        if c == 'N' {
            openPenalty++
        } else {
            closedPenalty--
        }

        currentPenalty := openPenalty + closedPenalty
        if currentPenalty < minPenalty {
            minPenalty = currentPenalty
            answer = i + 1
        }
    }
    return answer
}
```

---

## Step-by-step Detailed Explanation

1. Count total `'Y'` → penalty if shop closes at hour `0`
2. Initialize open and closed penalties
3. Iterate through each hour
4. Update penalties based on `'N'` or `'Y'`
5. Track minimum penalty and earliest hour
6. Return final answer

---

## Examples

**Input:** `customers = "YYNY"`
**Output:** `2`

**Input:** `customers = "NNNNN"`
**Output:** `0`

**Input:** `customers = "YYYY"`
**Output:** `4`

---

## How to use / Run locally

1. Clone the repository
2. Open the solution file in your preferred language
3. Compile and run using standard language commands:

   * C++: `g++ solution.cpp && ./a.out`
   * Java: `javac Solution.java && java Solution`
   * Python: `python solution.py`
   * Go: `go run solution.go`

---

## Notes & Optimizations

* Single-pass solution
* No prefix arrays needed
* Works efficiently for large inputs
* Interview-friendly logic
* Easy to explain and debug

---

## Author

* **Md Aarzoo Islam**
  [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
