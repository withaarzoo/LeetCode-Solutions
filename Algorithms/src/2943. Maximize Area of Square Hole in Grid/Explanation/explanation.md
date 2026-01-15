# Problem Title

**2943. Maximize Area of Square Hole in Grid**

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

I am given a grid formed by horizontal and vertical bars.

* The grid has `n + 2` horizontal bars
* The grid has `m + 2` vertical bars
* Some bars are removable
* Other bars are fixed and cannot be removed

By removing **some consecutive bars**, I can create a **square-shaped hole** inside the grid.

My task is to **find the maximum possible area of such a square hole**.

---

## Constraints

* 1 ≤ n, m ≤ 10⁹
* 1 ≤ length of hBars, vBars ≤ 100
* Bars are distinct
* Bars are indexed starting from 1

Important observation
The grid is very large, but the number of removable bars is very small.

---

## Intuition

When I read this problem, I ignored the actual grid size.

I focused only on the removable bars.

If I remove:

* `k` **consecutive horizontal bars**, I create a vertical gap of `k + 1`
* `k` **consecutive vertical bars**, I create a horizontal gap of `k + 1`

A square needs **equal height and width**, so the square side length depends on the **smaller gap**.

Once I know the side length, the area is simply
`side × side`.

---

## Approach

1. Sort `hBars` and `vBars`
2. Find the **longest consecutive sequence** in `hBars`
3. Find the **longest consecutive sequence** in `vBars`
4. Convert sequences into gaps by adding `1`
5. Take the minimum gap
6. Square it to get the answer

---

## Data Structures Used

* Arrays / Lists for bar positions
* Simple integer variables for tracking lengths

No advanced data structures are needed.

---

## Operations & Behavior Summary

* Sorting helps identify consecutive bars easily
* Consecutive bars increase the possible hole size
* Non-consecutive bars reset the count
* Square size is limited by the smaller direction

---

## Complexity

**Time Complexity**
O(H log H + V log V)
H = number of horizontal bars
V = number of vertical bars

Sorting dominates the runtime.

**Space Complexity**
O(1)
Only constant extra variables are used.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int getMaxGap(vector<int>& bars) {
        sort(bars.begin(), bars.end());
        int maxLen = 1, curLen = 1;

        for (int i = 1; i < bars.size(); i++) {
            if (bars[i] == bars[i - 1] + 1) {
                curLen++;
            } else {
                curLen = 1;
            }
            maxLen = max(maxLen, curLen);
        }
        return maxLen;
    }

    int maximizeSquareHoleArea(int n, int m, vector<int>& hBars, vector<int>& vBars) {
        int hGap = getMaxGap(hBars) + 1;
        int vGap = getMaxGap(vBars) + 1;
        int side = min(hGap, vGap);
        return side * side;
    }
};
```

---

### Java

```java
class Solution {

    private int getMaxGap(int[] bars) {
        Arrays.sort(bars);
        int maxLen = 1, curLen = 1;

        for (int i = 1; i < bars.length; i++) {
            if (bars[i] == bars[i - 1] + 1) {
                curLen++;
            } else {
                curLen = 1;
            }
            maxLen = Math.max(maxLen, curLen);
        }
        return maxLen;
    }

    public int maximizeSquareHoleArea(int n, int m, int[] hBars, int[] vBars) {
        int hGap = getMaxGap(hBars) + 1;
        int vGap = getMaxGap(vBars) + 1;
        int side = Math.min(hGap, vGap);
        return side * side;
    }
}
```

---

### JavaScript

```javascript
var maximizeSquareHoleArea = function(n, m, hBars, vBars) {

    const getMaxGap = (bars) => {
        bars.sort((a, b) => a - b);
        let maxLen = 1, curLen = 1;

        for (let i = 1; i < bars.length; i++) {
            if (bars[i] === bars[i - 1] + 1) {
                curLen++;
            } else {
                curLen = 1;
            }
            maxLen = Math.max(maxLen, curLen);
        }
        return maxLen;
    };

    let hGap = getMaxGap(hBars) + 1;
    let vGap = getMaxGap(vBars) + 1;
    let side = Math.min(hGap, vGap);

    return side * side;
};
```

---

### Python3

```python
class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: List[int], vBars: List[int]) -> int:

        def get_max_gap(bars):
            bars.sort()
            max_len = 1
            cur_len = 1

            for i in range(1, len(bars)):
                if bars[i] == bars[i - 1] + 1:
                    cur_len += 1
                else:
                    cur_len = 1
                max_len = max(max_len, cur_len)

            return max_len

        h_gap = get_max_gap(hBars) + 1
        v_gap = get_max_gap(vBars) + 1
        side = min(h_gap, v_gap)

        return side * side
```

---

### Go

```go
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {

    getMaxGap := func(bars []int) int {
        sort.Ints(bars)
        maxLen, curLen := 1, 1

        for i := 1; i < len(bars); i++ {
            if bars[i] == bars[i-1]+1 {
                curLen++
            } else {
                curLen = 1
            }
            if curLen > maxLen {
                maxLen = curLen
            }
        }
        return maxLen
    }

    hGap := getMaxGap(hBars) + 1
    vGap := getMaxGap(vBars) + 1
    side := min(hGap, vGap)

    return side * side
}
```

---

## Step-by-step Detailed Explanation

1. I sort the bars so consecutive values come together
2. I scan the array and count how many bars are consecutive
3. When the sequence breaks, I reset the counter
4. The longest sequence tells me how many bars I can remove together
5. Removing `k` bars creates a gap of `k + 1`
6. I repeat this for both directions
7. I take the smaller gap to form a square
8. I square it to get the maximum area

---

## Examples

Input
n = 2
m = 1
hBars = [2,3]
vBars = [2]

Output
4

Explanation

* Horizontal gap = 3
* Vertical gap = 2
* Square side = 2
* Area = 4

---

## How to use / Run locally

1. Clone the repository
2. Open the solution file for your language
3. Run it using the respective compiler or interpreter

Example for C++

```bash
g++ solution.cpp
./a.out
```

---

## Notes & Optimizations

* Grid size does not matter
* Only removable bars matter
* Sorting is unavoidable but efficient due to small input size
* Works within all constraints

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
