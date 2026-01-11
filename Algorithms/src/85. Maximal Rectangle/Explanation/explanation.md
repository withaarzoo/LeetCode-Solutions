# Maximal Rectangle (LeetCode 85)

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a binary matrix filled with `'0'` and `'1'`.
My task is to find the **largest rectangle** that contains **only 1s** and return its **area**.

The rectangle must be continuous and aligned with rows and columns.

---

## Constraints

* Number of rows and columns can be up to `200`
* Each cell contains either `'0'` or `'1'`
* Matrix size can be large, so the solution must be optimized

---

## Intuition

When I first saw this problem, I realized that directly checking all rectangles would be too slow.

Then I noticed something important.

If I look at the matrix **row by row**, each row can act as the **base of a histogram**.

For every column:

* I count how many continuous `1`s are stacked vertically
* This forms a histogram of heights

Once I have a histogram, I already know how to find the **largest rectangle** using a **monotonic stack**.

So my main idea was:

* Convert the 2D matrix into multiple histogram problems
* Solve each histogram efficiently
* Keep track of the maximum area

---

## Approach

1. I create an array `heights` with size equal to the number of columns
2. I iterate through the matrix row by row
3. For each cell:

   * If the value is `'1'`, I increase the height
   * If the value is `'0'`, I reset the height to `0`
4. After processing a row, I treat `heights` as a histogram
5. I compute the largest rectangle in that histogram using a **monotonic increasing stack**
6. I repeat this for every row and return the maximum area found

---

## Data Structures Used

* Array (for histogram heights)
* Stack (monotonic stack for histogram processing)

---

## Operations & Behavior Summary

* Heights accumulate vertically for consecutive `1`s
* Stack maintains increasing height indices
* When a smaller height appears:

  * I pop from the stack
  * I calculate area using popped height
* Each index is pushed and popped once for efficiency

---

## Complexity

**Time Complexity:**
`O(rows × columns)`
Every cell is processed once, and stack operations are linear.

**Space Complexity:**
`O(columns)`
Extra space is used only for the heights array and stack.

---

## Multi-language Solutions

---

### C++

```cpp
class Solution {
public:
    int maximalRectangle(vector<vector<char>>& matrix) {
        if (matrix.empty()) return 0;

        int cols = matrix[0].size();
        vector<int> heights(cols, 0);
        int maxArea = 0;

        for (auto& row : matrix) {
            for (int j = 0; j < cols; j++) {
                heights[j] = (row[j] == '1') ? heights[j] + 1 : 0;
            }

            stack<int> st;
            heights.push_back(0);

            for (int i = 0; i < heights.size(); i++) {
                while (!st.empty() && heights[st.top()] > heights[i]) {
                    int h = heights[st.top()];
                    st.pop();
                    int w = st.empty() ? i : i - st.top() - 1;
                    maxArea = max(maxArea, h * w);
                }
                st.push(i);
            }
            heights.pop_back();
        }
        return maxArea;
    }
};
```

---

### Java

```java
class Solution {
    public int maximalRectangle(char[][] matrix) {
        if (matrix.length == 0) return 0;

        int cols = matrix[0].length;
        int[] heights = new int[cols];
        int maxArea = 0;

        for (char[] row : matrix) {
            for (int j = 0; j < cols; j++) {
                heights[j] = row[j] == '1' ? heights[j] + 1 : 0;
            }

            Stack<Integer> stack = new Stack<>();
            for (int i = 0; i <= cols; i++) {
                int h = (i == cols) ? 0 : heights[i];
                while (!stack.isEmpty() && heights[stack.peek()] > h) {
                    int height = heights[stack.pop()];
                    int width = stack.isEmpty() ? i : i - stack.peek() - 1;
                    maxArea = Math.max(maxArea, height * width);
                }
                stack.push(i);
            }
        }
        return maxArea;
    }
}
```

---

### JavaScript

```javascript
var maximalRectangle = function(matrix) {
    if (matrix.length === 0) return 0;

    const cols = matrix[0].length;
    let heights = new Array(cols).fill(0);
    let maxArea = 0;

    for (let row of matrix) {
        for (let j = 0; j < cols; j++) {
            heights[j] = row[j] === '1' ? heights[j] + 1 : 0;
        }

        let stack = [];
        heights.push(0);

        for (let i = 0; i < heights.length; i++) {
            while (stack.length && heights[stack[stack.length - 1]] > heights[i]) {
                let h = heights[stack.pop()];
                let w = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
                maxArea = Math.max(maxArea, h * w);
            }
            stack.push(i);
        }
        heights.pop();
    }
    return maxArea;
};
```

---

### Python3

```python
class Solution:
    def maximalRectangle(self, matrix):
        if not matrix:
            return 0

        cols = len(matrix[0])
        heights = [0] * cols
        max_area = 0

        for row in matrix:
            for j in range(cols):
                heights[j] = heights[j] + 1 if row[j] == '1' else 0

            stack = []
            heights.append(0)

            for i in range(len(heights)):
                while stack and heights[stack[-1]] > heights[i]:
                    h = heights[stack.pop()]
                    w = i if not stack else i - stack[-1] - 1
                    max_area = max(max_area, h * w)
                stack.append(i)

            heights.pop()

        return max_area
```

---

### Go

```go
func maximalRectangle(matrix [][]byte) int {
 if len(matrix) == 0 {
  return 0
 }

 cols := len(matrix[0])
 heights := make([]int, cols)
 maxArea := 0

 for _, row := range matrix {
  for j := 0; j < cols; j++ {
   if row[j] == '1' {
    heights[j]++
   } else {
    heights[j] = 0
   }
  }

  stack := []int{}
  h := append(heights, 0)

  for i := 0; i < len(h); i++ {
   for len(stack) > 0 && h[stack[len(stack)-1]] > h[i] {
    height := h[stack[len(stack)-1]]
    stack = stack[:len(stack)-1]
    width := i
    if len(stack) > 0 {
     width = i - stack[len(stack)-1] - 1
    }
    if height*width > maxArea {
     maxArea = height * width
    }
   }
   stack = append(stack, i)
  }
 }
 return maxArea
}
```

---

## Step-by-step Detailed Explanation

1. I scan the matrix row by row
2. I build histogram heights for each column
3. I use a monotonic stack to track increasing heights
4. When a smaller height appears, I calculate area
5. Width is calculated using previous smaller element
6. I repeat this for all rows
7. The largest area found is my answer

---

## Examples

Input:

```bash
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
```

Output:

```bash
6
```

---

## How to Use / Run Locally

1. Clone the repository
2. Choose your language folder
3. Compile or run using standard compiler
4. Test with custom inputs

Example (C++):

```bash
g++ solution.cpp
./a.out
```

---

## Notes & Optimizations

* This solution avoids brute force
* Stack ensures linear time per row
* Sentinel `0` simplifies boundary handling
* Very common interview pattern

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
