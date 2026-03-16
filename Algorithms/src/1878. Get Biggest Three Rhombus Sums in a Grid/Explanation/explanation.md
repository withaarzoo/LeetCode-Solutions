# 1878. Get Biggest Three Rhombus Sums in a Grid

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

We are given an `m x n` integer grid. A **rhombus** is a diamond shaped figure whose four corners lie on grid cells. The rhombus must be a square rotated 45 degrees.

The **rhombus sum** is defined as the **sum of the border cells** of the rhombus.

Our task is to:

Return the **three largest distinct rhombus sums** present in the grid. If fewer than three distinct sums exist, return all of them.

Important points:

* Only border cells are counted
* Rhombus of size `0` is just a single cell
* Returned values must be **distinct**

---

## Constraints

```
1 <= m, n <= 50
1 <= grid[i][j] <= 10^5
```

Because the grid size is relatively small, exploring every possible rhombus centered at each cell is feasible.

---

## Intuition

When I started thinking about this problem, I noticed that every rhombus can be uniquely identified by:

1. A **center cell**
2. A **distance (size)** from the center

For example:

* size = 0 → single cell
* size = 1 → small diamond
* size = 2 → bigger diamond

So the idea is simple:

1. Treat every cell as the **center** of a rhombus.
2. Expand the rhombus outward layer by layer.
3. For each layer, calculate the **border sum**.
4. Store the sum in a set to keep values unique.
5. Finally return the **top 3 largest sums**.

---

## Approach

Step 1: Iterate through every cell `(r, c)` in the grid.

Step 2: Add the **size 0 rhombus**, which is simply the cell itself.

Step 3: Determine the **maximum rhombus size** possible from that center.

```
maxSize = min(r, c, m-1-r, n-1-c)
```

This ensures the rhombus stays inside the grid.

Step 4: For each possible size `k`:

Traverse the four edges:

1. Top → Right
2. Right → Bottom
3. Bottom → Left
4. Left → Top

Each edge moves diagonally.

Step 5: Add the computed border sum to a **set**.

Step 6: Convert the set to a list, sort descending, and return the **first three values**.

---

## Data Structures Used

Set

Purpose:

* Maintain **unique rhombus sums**
* Automatically remove duplicates

Vector / List / Array

Purpose:

* Store sorted results
* Extract top three values

---

## Operations & Behavior Summary

| Operation        | Purpose                          |
| ---------------- | -------------------------------- |
| Iterate Grid     | Try every cell as rhombus center |
| Compute Size     | Ensure rhombus stays inside grid |
| Traverse Borders | Calculate rhombus border sum     |
| Insert Into Set  | Maintain distinct sums           |
| Sort Results     | Get largest values               |

---

## Complexity

### Time Complexity

```
O(m * n * min(m,n))
```

Where:

* `m` = number of rows
* `n` = number of columns

For each cell we try expanding rhombus sizes.

Worst case:

```
50 * 50 * 25 ≈ 62,500 iterations
```

Which is efficient.

### Space Complexity

```
O(k)
```

Where `k` is the number of distinct rhombus sums stored in the set.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> getBiggestThree(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();

        set<int> sums;

        for(int r = 0; r < m; r++){
            for(int c = 0; c < n; c++){

                sums.insert(grid[r][c]);

                int maxSize = min({r, c, m-1-r, n-1-c});

                for(int k = 1; k <= maxSize; k++){

                    int sum = 0;

                    for(int i = 0; i < k; i++)
                        sum += grid[r-k+i][c+i];

                    for(int i = 0; i < k; i++)
                        sum += grid[r+i][c+k-i];

                    for(int i = 0; i < k; i++)
                        sum += grid[r+k-i][c-i];

                    for(int i = 0; i < k; i++)
                        sum += grid[r-i][c-k+i];

                    sums.insert(sum);
                }
            }
        }

        vector<int> res(sums.begin(), sums.end());
        sort(res.rbegin(), res.rend());

        if(res.size() > 3) res.resize(3);

        return res;
    }
};
```

### Java

```java
class Solution {
    public int[] getBiggestThree(int[][] grid) {

        int m = grid.length;
        int n = grid[0].length;

        TreeSet<Integer> set = new TreeSet<>();

        for(int r=0;r<m;r++){
            for(int c=0;c<n;c++){

                set.add(grid[r][c]);

                int maxSize = Math.min(Math.min(r,c), Math.min(m-1-r,n-1-c));

                for(int k=1;k<=maxSize;k++){

                    int sum=0;

                    for(int i=0;i<k;i++)
                        sum+=grid[r-k+i][c+i];

                    for(int i=0;i<k;i++)
                        sum+=grid[r+i][c+k-i];

                    for(int i=0;i<k;i++)
                        sum+=grid[r+k-i][c-i];

                    for(int i=0;i<k;i++)
                        sum+=grid[r-i][c-k+i];

                    set.add(sum);
                }
            }
        }

        List<Integer> list=new ArrayList<>(set);
        Collections.sort(list,Collections.reverseOrder());

        int size=Math.min(3,list.size());

        int[] ans=new int[size];

        for(int i=0;i<size;i++)
            ans[i]=list.get(i);

        return ans;
    }
}
```

### JavaScript

```javascript
var getBiggestThree = function(grid) {

    let m = grid.length
    let n = grid[0].length

    let set = new Set()

    for(let r=0;r<m;r++){
        for(let c=0;c<n;c++){

            set.add(grid[r][c])

            let maxSize = Math.min(r,c,m-1-r,n-1-c)

            for(let k=1;k<=maxSize;k++){

                let sum = 0

                for(let i=0;i<k;i++)
                    sum += grid[r-k+i][c+i]

                for(let i=0;i<k;i++)
                    sum += grid[r+i][c+k-i]

                for(let i=0;i<k;i++)
                    sum += grid[r+k-i][c-i]

                for(let i=0;i<k;i++)
                    sum += grid[r-i][c-k+i]

                set.add(sum)
            }
        }
    }

    let res = Array.from(set)

    res.sort((a,b)=>b-a)

    return res.slice(0,3)
}
```

### Python3

```python
class Solution:
    def getBiggestThree(self, grid):

        m=len(grid)
        n=len(grid[0])

        sums=set()

        for r in range(m):
            for c in range(n):

                sums.add(grid[r][c])

                maxSize=min(r,c,m-1-r,n-1-c)

                for k in range(1,maxSize+1):

                    s=0

                    for i in range(k):
                        s+=grid[r-k+i][c+i]

                    for i in range(k):
                        s+=grid[r+i][c+k-i]

                    for i in range(k):
                        s+=grid[r+k-i][c-i]

                    for i in range(k):
                        s+=grid[r-i][c-k+i]

                    sums.add(s)

        res=sorted(sums,reverse=True)

        return res[:3]
```

### Go

```go
func getBiggestThree(grid [][]int) []int {

    m := len(grid)
    n := len(grid[0])

    sums := map[int]bool{}

    for r:=0;r<m;r++{
        for c:=0;c<n;c++{

            sums[grid[r][c]] = true

            maxSize := min(min(r,c), min(m-1-r,n-1-c))

            for k:=1;k<=maxSize;k++{

                sum := 0

                for i:=0;i<k;i++{
                    sum += grid[r-k+i][c+i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r+i][c+k-i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r+k-i][c-i]
                }

                for i:=0;i<k;i++{
                    sum += grid[r-i][c-k+i]
                }

                sums[sum] = true
            }
        }
    }

    res := []int{}

    for k := range sums{
        res = append(res,k)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(res)))

    if len(res) > 3 {
        res = res[:3]
    }

    return res
}

func min(a,b int) int{
    if a<b {return a}
    return b
}
```

---

## Step-by-step Detailed Explanation

1. Iterate through every cell in the grid.

2. Treat each cell as the **center of a rhombus**.

3. Insert the value of the cell itself because a rhombus of size `0` is valid.

4. Calculate the maximum rhombus size that can fit without leaving the grid.

5. For each possible rhombus size:

   * Move from **top corner to right corner** diagonally.
   * Move from **right corner to bottom corner** diagonally.
   * Move from **bottom corner to left corner** diagonally.
   * Move from **left corner to top corner** diagonally.

6. Sum all these border values.

7. Insert the sum into a set.

8. After processing all centers, convert the set to a list.

9. Sort the list in descending order.

10. Return the first three values.

---

## Examples

Example 1

```
Input:
[[3,4,5,1,3],
 [3,3,4,2,3],
 [20,30,200,40,10],
 [1,5,5,4,1],
 [4,3,2,2,5]]

Output:
[228,216,211]
```

Example 2

```
Input:
[[1,2,3],[4,5,6],[7,8,9]]

Output:
[20,9,8]
```

---

## How to use / Run locally

Example using Python

```
python solution.py
```

Example using C++

```
g++ solution.cpp
./a.out
```

Example using Go

```
go run solution.go
```

---

## Notes & Optimizations

Possible improvement:

Use **diagonal prefix sums** to compute rhombus borders faster.

This reduces repeated calculations and makes the solution closer to **O(n²)**.

However, the current implementation is already efficient for the given constraints.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
