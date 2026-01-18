# 🧩 Problem Title

**1895. Largest Magic Square**

---

## 📑 Table of Contents

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

## 📘 Problem Summary

I am given a **2D grid of integers**.

A **magic square** is a square subgrid where:

* All **row sums** are equal
* All **column sums** are equal
* Both **diagonal sums** are equal

The size of a magic square is its **side length (k × k)**.

My task is to **find the largest possible magic square** that exists inside the given grid and return its size.

Important note:

* A **1×1 grid is always a magic square**

---

## 📌 Constraints

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 50`
* `1 <= grid[i][j] <= 10^6`

---

## 💡 Intuition

When I first read the problem, I understood that checking every square directly would be slow.

So I thought:

* Bigger magic squares are more important than smaller ones
* If I find a valid large square, I can stop early
* Recalculating row and column sums repeatedly is inefficient

That’s when I decided to use **prefix sums** to make sum calculation fast.

My strategy was simple:

* Try all possible square sizes from **largest to smallest**
* For each square, check rows, columns, and diagonals efficiently
* Return the first valid size I find

---

## 🛠️ Approach

1. I precompute:

   * Row prefix sums
   * Column prefix sums

2. I iterate over possible square sizes `k`

   * Start from `min(m, n)`
   * Move down to size `2`

3. For every top-left position of a `k × k` square:

   * Take the first row sum as the **target**
   * Check all rows
   * Check all columns
   * Check both diagonals

4. If everything matches:

   * I immediately return `k`

5. If no square larger than `1` is valid:

   * I return `1`

---

## 🧱 Data Structures Used

* 2D arrays for:

  * Row prefix sums
  * Column prefix sums
* Simple integer variables for comparisons

No extra complex data structures are used.

---

## 🔁 Operations & Behavior Summary

* Prefix sum helps fetch row/column sum in **O(1)**
* Nested loops scan valid square regions
* Early break improves performance
* Diagonal sums are computed directly

---

## ⏱️ Complexity

### Time Complexity

**O(m × n × min(m, n))**

* `m × n` for grid traversal
* `min(m, n)` for possible square sizes

This is efficient enough because grid size is limited to 50.

### Space Complexity

**O(m × n)**

* Used for row and column prefix sum arrays

---

## 🌐 Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int largestMagicSquare(vector<vector<int>>& grid) {
        int m = grid.size(), n = grid[0].size();
        vector<vector<int>> row(m, vector<int>(n+1,0)), col(m+1, vector<int>(n,0));

        for(int i=0;i<m;i++)
            for(int j=0;j<n;j++){
                row[i][j+1] = row[i][j] + grid[i][j];
                col[i+1][j] = col[i][j] + grid[i][j];
            }

        for(int k=min(m,n); k>=2; k--){
            for(int i=0;i+k<=m;i++){
                for(int j=0;j+k<=n;j++){
                    int target = row[i][j+k] - row[i][j];
                    bool ok = true;

                    for(int r=i;r<i+k && ok;r++)
                        if(row[r][j+k]-row[r][j]!=target) ok=false;

                    for(int c=j;c<j+k && ok;c++)
                        if(col[i+k][c]-col[i][c]!=target) ok=false;

                    int d1=0,d2=0;
                    for(int x=0;x<k;x++){
                        d1+=grid[i+x][j+x];
                        d2+=grid[i+x][j+k-1-x];
                    }

                    if(ok && d1==target && d2==target) return k;
                }
            }
        }
        return 1;
    }
};
```

---

### Java

```java
class Solution {
    public int largestMagicSquare(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[][] row = new int[m][n+1];
        int[][] col = new int[m+1][n];

        for(int i=0;i<m;i++)
            for(int j=0;j<n;j++){
                row[i][j+1] = row[i][j] + grid[i][j];
                col[i+1][j] = col[i][j] + grid[i][j];
            }

        for(int k=Math.min(m,n);k>=2;k--){
            for(int i=0;i+k<=m;i++){
                for(int j=0;j+k<=n;j++){
                    int target = row[i][j+k] - row[i][j];
                    boolean ok = true;

                    for(int r=i;r<i+k && ok;r++)
                        if(row[r][j+k]-row[r][j]!=target) ok=false;

                    for(int c=j;c<j+k && ok;c++)
                        if(col[i+k][c]-col[i][c]!=target) ok=false;

                    int d1=0,d2=0;
                    for(int x=0;x<k;x++){
                        d1+=grid[i+x][j+x];
                        d2+=grid[i+x][j+k-1-x];
                    }

                    if(ok && d1==target && d2==target) return k;
                }
            }
        }
        return 1;
    }
}
```

---

### JavaScript

```javascript
var largestMagicSquare = function(grid) {
    const m = grid.length, n = grid[0].length;
    const row = Array.from({length:m},()=>Array(n+1).fill(0));
    const col = Array.from({length:m+1},()=>Array(n).fill(0));

    for(let i=0;i<m;i++)
        for(let j=0;j<n;j++){
            row[i][j+1]=row[i][j]+grid[i][j];
            col[i+1][j]=col[i][j]+grid[i][j];
        }

    for(let k=Math.min(m,n);k>=2;k--){
        for(let i=0;i+k<=m;i++){
            for(let j=0;j+k<=n;j++){
                let target=row[i][j+k]-row[i][j];
                let ok=true;

                for(let r=i;r<i+k && ok;r++)
                    if(row[r][j+k]-row[r][j]!==target) ok=false;

                for(let c=j;c<j+k && ok;c++)
                    if(col[i+k][c]-col[i][c]!==target) ok=false;

                let d1=0,d2=0;
                for(let x=0;x<k;x++){
                    d1+=grid[i+x][j+x];
                    d2+=grid[i+x][j+k-1-x];
                }

                if(ok && d1===target && d2===target) return k;
            }
        }
    }
    return 1;
};
```

---

### Python3

```python
class Solution:
    def largestMagicSquare(self, grid):
        m, n = len(grid), len(grid[0])
        row = [[0]*(n+1) for _ in range(m)]
        col = [[0]*n for _ in range(m+1)]

        for i in range(m):
            for j in range(n):
                row[i][j+1] = row[i][j] + grid[i][j]
                col[i+1][j] = col[i][j] + grid[i][j]

        for k in range(min(m,n),1,-1):
            for i in range(m-k+1):
                for j in range(n-k+1):
                    target = row[i][j+k] - row[i][j]
                    ok = True

                    for r in range(i,i+k):
                        if row[r][j+k]-row[r][j] != target:
                            ok = False

                    for c in range(j,j+k):
                        if col[i+k][c]-col[i][c] != target:
                            ok = False

                    d1 = d2 = 0
                    for x in range(k):
                        d1 += grid[i+x][j+x]
                        d2 += grid[i+x][j+k-1-x]

                    if ok and d1==target and d2==target:
                        return k
        return 1
```

---

### Go

```go
func largestMagicSquare(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    row := make([][]int, m)
    col := make([][]int, m+1)

    for i := range row {
        row[i] = make([]int, n+1)
    }
    for i := range col {
        col[i] = make([]int, n)
    }

    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            row[i][j+1]=row[i][j]+grid[i][j]
            col[i+1][j]=col[i][j]+grid[i][j]
        }
    }

    for k:=min(m,n);k>=2;k--{
        for i:=0;i+k<=m;i++{
            for j:=0;j+k<=n;j++{
                target:=row[i][j+k]-row[i][j]
                ok:=true

                for r:=i;r<i+k;r++{
                    if row[r][j+k]-row[r][j]!=target { ok=false }
                }

                for c:=j;c<j+k;c++{
                    if col[i+k][c]-col[i][c]!=target { ok=false }
                }

                d1,d2:=0,0
                for x:=0;x<k;x++{
                    d1+=grid[i+x][j+x]
                    d2+=grid[i+x][j+k-1-x]
                }

                if ok && d1==target && d2==target {
                    return k
                }
            }
        }
    }
    return 1
}
```

---

## 🧠 Step-by-step Detailed Explanation

* Prefix sums allow instant row and column sum checks
* Largest square is tested first to avoid unnecessary work
* Each square validates:

  * All rows
  * All columns
  * Both diagonals
* Early exit makes the solution efficient

---

## 🧪 Examples

**Input**

```bash
[[7,1,4,5,6],
 [2,5,1,6,4],
 [1,5,4,3,2],
 [1,2,7,3,4]]
```

**Output**

```bash
3
```

---

## ▶️ How to use / Run locally

1. Clone the repository
2. Open the file in your preferred language
3. Run using standard compiler or interpreter
4. Test with custom inputs

---

## 📝 Notes & Optimizations

* Prefix sums are the key optimization
* Early break reduces unnecessary checks
* Works comfortably within constraints

---

## 👤 Author

* **Md Aarzoo Islam**
  👉 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
