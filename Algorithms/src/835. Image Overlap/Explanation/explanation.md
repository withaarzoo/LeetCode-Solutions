# LeetCode 835 Image Overlap Solution – Maximum Binary Matrix Overlap

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given two binary square matrices of size n by n. Each cell contains either 0 or 1.  
The task is to translate (slide) one matrix any number of units left, right, up, or down and then place it on top of the other matrix. After the slide, count how many positions contain a 1 in both matrices at the same time. That count is called the overlap.  

Bits that slide outside the matrix borders are simply erased. Rotation is not allowed.  
Return the largest possible overlap you can achieve by choosing the best translation.

This is a classic competitive programming problem that appears in interviews and weekly contests. It tests your ability to reason about relative shifts on binary matrices without rebuilding the full grid every time.

## Constraints

- Both matrices have the same size n  
- 1 <= n <= 30  
- Every cell in both matrices is either 0 or 1  

Because n is small, solutions whose running time is roughly O(n^4) are acceptable.

## Intuition

The only cells that can contribute to the overlap are the cells that hold a 1.  
If I slide one image so that a particular 1 from the first image lands exactly on a 1 from the second image, that pair contributes one to the overlap.  

Every possible pair of 1-cells defines a unique shift (dx, dy).  
If many pairs need the exact same shift, that shift produces a high overlap.  
Therefore the answer is simply the maximum number of pairs that share the same shift.

This observation turns the geometric sliding problem into a simple counting problem over coordinate differences.

## Approach

1. Walk through both matrices once and collect the list of all coordinates that contain a 1.  
2. For every pair of coordinates (one from each list) compute the required shift:  
   dx = row2 - row1  
   dy = col2 - col1  
3. Count how many times each possible (dx, dy) appears.  
4. The highest count is the maximum overlap.

Because coordinate differences range from -(n-1) to +(n-1), a small 2-D array of size roughly 2n by 2n is enough to store the counts. No hash map is required.

## Data Structures Used

- Two lists (or vectors) of coordinate pairs – store the positions of every 1.  
- A 2-D integer array of size 2n × 2n – counts the frequency of each possible shift after an offset of n is added so indices stay non-negative.

These structures keep the algorithm simple and fast for the given constraints.

## Operations & Behavior Summary

- Scan both matrices and record every (row, column) that holds a 1.  
- Initialize a count matrix large enough for all legal shifts.  
- For every pair of 1-positions compute the shift and increment the corresponding cell in the count matrix.  
- Keep a running maximum while updating the counts.  
- Return that maximum value.  

If either matrix contains no 1s the lists stay empty and the answer is correctly reported as 0.

## Complexity

| Type              | Value          | Explanation |
|-------------------|----------------|-------------|
| Time Complexity   | O(n^{2} + k1·k2) | Collecting positions costs O(n^{2}). The double loop over the k1 and k2 ones costs O(k1·k2). In the worst case both are n^{2}, giving O(n^{4}). |
| Space Complexity  | O(n^{2})        | The two position lists and the shift-count array each use O(n^{2}) space in the worst case. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int largestOverlap(vector<vector<int>>& img1, vector<vector<int>>& img2) {
        int n = img1.size();
        // collect every coordinate that holds a 1
        vector<pair<int,int>> A, B;
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (img1[i][j] == 1) A.emplace_back(i, j);
                if (img2[i][j] == 1) B.emplace_back(i, j);
            }
        }
        // count how many times each possible shift appears
        // shifts range from -(n-1) to +(n-1), so offset by n
        vector<vector<int>> cnt(2 * n, vector<int>(2 * n, 0));
        int best = 0;
        for (auto& a : A) {
            for (auto& b : B) {
                int dx = b.first - a.first + n;   // make index non-negative
                int dy = b.second - a.second + n;
                best = max(best, ++cnt[dx][dy]);
            }
        }
        return best;
    }
};
```

### Java

```java
class Solution {
    public int largestOverlap(int[][] img1, int[][] img2) {
        int n = img1.length;
        // collect every coordinate that holds a 1
        List<int[]> A = new ArrayList<>();
        List<int[]> B = new ArrayList<>();
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (img1[i][j] == 1) A.add(new int[]{i, j});
                if (img2[i][j] == 1) B.add(new int[]{i, j});
            }
        }
        // count how many times each possible shift appears
        // shifts range from -(n-1) to +(n-1), so offset by n
        int[][] cnt = new int[2 * n][2 * n];
        int best = 0;
        for (int[] a : A) {
            for (int[] b : B) {
                int dx = b[0] - a[0] + n;   // make index non-negative
                int dy = b[1] - a[1] + n;
                best = Math.max(best, ++cnt[dx][dy]);
            }
        }
        return best;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} img1
 * @param {number[][]} img2
 * @return {number}
 */
var largestOverlap = function(img1, img2) {
    const n = img1.length;
    // collect every coordinate that holds a 1
    const A = [], B = [];
    for (let i = 0; i < n; ++i) {
        for (let j = 0; j < n; ++j) {
            if (img1[i][j] === 1) A.push([i, j]);
            if (img2[i][j] === 1) B.push([i, j]);
        }
    }
    // count how many times each possible shift appears
    // shifts range from -(n-1) to +(n-1), so offset by n
    const cnt = Array.from({length: 2 * n}, () => Array(2 * n).fill(0));
    let best = 0;
    for (const a of A) {
        for (const b of B) {
            const dx = b[0] - a[0] + n;   // make index non-negative
            const dy = b[1] - a[1] + n;
            best = Math.max(best, ++cnt[dx][dy]);
        }
    }
    return best;
};
```

### TypeScript

```typescript
function largestOverlap(img1: number[][], img2: number[][]): number {
    const n = img1.length;
    // collect every coordinate that holds a 1
    const A: number[][] = [], B: number[][] = [];
    for (let i = 0; i < n; ++i) {
        for (let j = 0; j < n; ++j) {
            if (img1[i][j] === 1) A.push([i, j]);
            if (img2[i][j] === 1) B.push([i, j]);
        }
    }
    // count how many times each possible shift appears
    // shifts range from -(n-1) to +(n-1), so offset by n
    const cnt: number[][] = Array.from({length: 2 * n}, () => Array(2 * n).fill(0));
    let best = 0;
    for (const a of A) {
        for (const b of B) {
            const dx = b[0] - a[0] + n;   // make index non-negative
            const dy = b[1] - a[1] + n;
            best = Math.max(best, ++cnt[dx][dy]);
        }
    }
    return best;
}
```

### Python3

```python
class Solution:
    def largestOverlap(self, img1: List[List[int]], img2: List[List[int]]) -> int:
        n = len(img1)
        # collect every coordinate that holds a 1
        A = [(i, j) for i in range(n) for j in range(n) if img1[i][j] == 1]
        B = [(i, j) for i in range(n) for j in range(n) if img2[i][j] == 1]
        # count how many times each possible shift appears
        # shifts range from -(n-1) to +(n-1), so offset by n
        cnt = [[0] * (2 * n) for _ in range(2 * n)]
        best = 0
        for ax, ay in A:
            for bx, by in B:
                dx = bx - ax + n   # make index non-negative
                dy = by - ay + n
                cnt[dx][dy] += 1
                best = max(best, cnt[dx][dy])
        return best
```

### Go

```go
func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    // collect every coordinate that holds a 1
    type pos struct{ x, y int }
    var A, B []pos
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if img1[i][j] == 1 {
                A = append(A, pos{i, j})
            }
            if img2[i][j] == 1 {
                B = append(B, pos{i, j})
            }
        }
    }
    // count how many times each possible shift appears
    // shifts range from -(n-1) to +(n-1), so offset by n
    cnt := make([][]int, 2*n)
    for i := range cnt {
        cnt[i] = make([]int, 2*n)
    }
    best := 0
    for _, a := range A {
        for _, b := range B {
            dx := b.x - a.x + n // make index non-negative
            dy := b.y - a.y + n
            cnt[dx][dy]++
            if cnt[dx][dy] > best {
                best = cnt[dx][dy]
            }
        }
    }
    return best
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical in every language, only the syntax for lists and arrays changes.

First the size n is read. Both matrices are known to be square, so a single variable is enough.

Two containers are prepared to hold the coordinates of the ones. A nested loop walks every cell; whenever a 1 is found the pair (i, j) is stored. After this step the original matrices are no longer needed.

A 2-D count array of side length 2n is allocated and filled with zeros. Possible differences of coordinates lie between -(n-1) and +(n-1). Adding n maps every difference into the safe range [1 \ldots 2n-1].

A double loop then examines every combination of a position from the first list and a position from the second list. The required shift is calculated, the corresponding cell in the count array is incremented, and a global maximum is updated.

When the loops finish the maximum value is returned.  

Edge-case behaviour is automatic: empty lists produce a zero answer; shifts that would move bits outside the matrix never appear in any pair and are therefore ignored.

The same sequence of steps is followed in C++, Java, JavaScript, TypeScript, Python and Go. Only the concrete syntax for vectors, ArrayLists, arrays or slices differs.

## Examples

**Example 1**  
Input  
img1 = [[1,1,0],[0,1,0],[0,1,0]]  
img2 = [[0,0,0],[0,1,1],[0,0,1]]  

Positions of 1s  
img1: (0,0) (0,1) (1,1) (2,1)  
img2: (1,1) (1,2) (2,2)  

The shift (dx=1, dy=1) occurs three times, so the largest overlap is 3.

**Example 2**  
Input  
img1 = [[1]]  
img2 = [[1]]  

Only one pair exists and it needs the zero shift. Overlap = 1.

**Example 3**  
Input  
img1 = [[0]]  
img2 = [[0]]  

No 1s exist, lists are empty, answer = 0.

## How to Use / Run Locally

1. Copy the code of the language you prefer into a file (Solution.cpp, Solution.java, etc.).  
2. Add a small main function or driver that reads the two matrices and prints the result of largestOverlap.  
3. Compile and run:

- C++: `g++ -std=c++17 Solution.cpp -o sol && ./sol`  
- Java: `javac Solution.java && java Solution`  
- JavaScript: `node Solution.js`  
- TypeScript: `tsc Solution.ts && node Solution.js`  
- Python: `python3 Solution.py`  
- Go: `go run Solution.go`

Replace the driver with your own test cases as needed.

## Notes & Optimizations

- Because n ≤ 30 the O(n^{4}) bound is perfectly acceptable.  
- An alternative is to try every possible shift from -(n-1) to +(n-1) and count overlaps directly; the position-pair method is usually cleaner and often faster in practice.  
- Using a hash map instead of a fixed 2-D array also works, but the array version avoids hash overhead and is cache-friendly.  
- The algorithm never rotates the images and correctly discards bits that leave the matrix, matching the problem statement exactly.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
