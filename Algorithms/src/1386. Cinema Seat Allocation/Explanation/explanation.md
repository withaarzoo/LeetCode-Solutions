# 1386. Cinema Seat Allocation - LeetCode DSA Solution

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

LeetCode 1386, **Cinema Seat Allocation**, asks us to find the maximum number of four-person groups that can be seated in a cinema.

The cinema has `n` rows, and every row contains 10 seats numbered from `1` to `10`.

Some seats are already reserved. A four-person group must sit together in the same row, and there are only three possible blocks:

* Seats `2, 3, 4, 5`
* Seats `4, 5, 6, 7`
* Seats `6, 7, 8, 9`

A block can be used only when none of its four seats are reserved.

The goal is to return the maximum number of four-person groups that can be assigned.

The main challenge is that `n` can be as large as `10^9`, so checking every row is not practical. The solution instead processes only the rows that contain relevant reserved seats.

## Constraints

* `1 <= n <= 10^9`
* `1 <= reservedSeats.length <= min(10 * n, 10^4)`
* `reservedSeats[i] = [row_i, seat_i]`
* `1 <= row_i <= n`
* `1 <= seat_i <= 10`
* All entries in `reservedSeats` are distinct.

## Intuition

I first noticed that there are only three possible places where a group can sit:

```text
2 3 4 5
4 5 6 7
6 7 8 9
```

The left block and the right block do not overlap. So if both are available, I can place two groups in the same row.

If that is not possible, I only need to check whether the left, middle, or right block is available. If at least one is free, I can place one group.

The large value of `n` is the main thing I have to handle carefully. Since `n` can reach `10^9`, I cannot create an array for every row.

Instead, I store only the rows that actually appear in `reservedSeats`. Every other row has no relevant reservations and can always hold two groups.

To make the seat checks fast, I represent the reserved seats of each row using a bitmask.

## Approach

I use a hash map to store the reserved seats for each affected row.

For every reservation `[row, seat]`, I mark that seat in the row's bitmask. Seats `1` and `10` can be ignored because they are not part of any valid four-seat group.

After processing all reservations, I know exactly which rows need special handling.

For all rows that do not appear in the map, I immediately add two groups because they are completely free.

For every affected row, I check these three blocks:

```text
Left:    2 3 4 5
Middle:  4 5 6 7
Right:   6 7 8 9
```

If the left and right blocks are both free, I add two groups.

Otherwise, if any of the three blocks is free, I add one group.

If all three blocks are blocked, I add zero groups for that row.

This gives an efficient greedy solution using hashing and bit manipulation.

## Data Structures Used

### Hash Map

I use a hash map where:

* The key is the row number.
* The value is the bitmask containing the reserved seats of that row.

This is important because the number of rows can be extremely large, while the number of reserved seats is limited to `10^4`.

### Bitmask

I use an integer as a compact representation of reserved seats.

Each seat is represented by a bit. Bitwise `AND` lets me quickly check whether a possible group overlaps with any reserved seat.

This makes the seat-block checks constant time.

## Operations & Behavior Summary

The algorithm works in these stages:

1. Create a hash map for rows with relevant reserved seats.
2. Read every `[row, seat]` pair.
3. Ignore seats `1` and `10`.
4. Mark seats `2` through `9` using a bitmask.
5. Count rows that are not present in the map.
6. Add two groups for every completely free row.
7. For each affected row, check the `2-5`, `4-7`, and `6-9` blocks.
8. If `2-5` and `6-9` are both free, add two groups.
9. Otherwise, if any one of the three blocks is free, add one group.
10. Return the final number of groups.

In simple pseudocode:

```text
store reserved seats for each affected row

answer = 2 * number of completely empty rows

for every affected row:
    check seats 2-5
    check seats 4-7
    check seats 6-9

    if 2-5 and 6-9 are free:
        answer += 2
    else if any block is free:
        answer += 1

return answer
```

## Complexity

Let `k` be the number of reserved seats in `reservedSeats`, and let `r` be the number of distinct rows containing relevant reservations.

| Complexity       | Value  | Explanation                                                                                                    |
| ---------------- | ------ | -------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(k)` | I process every reserved seat once and then check each affected row. Since `r <= k`, the total remains `O(k)`. |
| Space Complexity | `O(r)` | I store only rows that contain relevant reserved seats instead of creating storage for all `n` rows.           |

The solution is efficient even when `n = 10^9` because the algorithm never loops through all cinema rows.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxNumberOfFamilies(int n, vector<vector<int>>& reservedSeats) {
        // Store a bitmask for every row that has at least one reserved seat.
        unordered_map<int, int> rows;

        // Process every reserved seat and mark that seat in its row's bitmask.
        for (const auto& seat : reservedSeats) {
            int row = seat[0];
            int col = seat[1];

            // Seats 2 through 9 are represented by bits 1 through 8.
            // Seats 1 and 10 do not belong to any possible group, so I ignore them.
            if (col >= 2 && col <= 9) {
                rows[row] |= (1 << col);
            }
        }

        // Start with the rows that have no reserved seats.
        // Every completely empty row can always fit two groups.
        int answer = 2 * (n - static_cast<int>(rows.size()));

        // These masks represent the three possible four-seat blocks.
        // left  = seats 2,3,4,5
        // middle = seats 4,5,6,7
        // right = seats 6,7,8,9
        const int left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
        const int middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
        const int right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

        // Check only rows that contain reserved seats.
        for (const auto& [row, mask] : rows) {
            // A block is available when none of its seats are reserved.
            bool canLeft = (mask & left) == 0;
            bool canMiddle = (mask & middle) == 0;
            bool canRight = (mask & right) == 0;

            // Left and right blocks do not overlap, so both can be used together.
            if (canLeft && canRight) {
                answer += 2;
            }
            // Otherwise, if any one block is available, I can place one group.
            else if (canLeft || canMiddle || canRight) {
                answer += 1;
            }
            // If no block is available, this row cannot fit a group.
        }

        // Return the maximum number of four-person groups.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public int maxNumberOfFamilies(int n, int[][] reservedSeats) {
        // Store the reserved seats of each affected row as a bitmask.
        HashMap<Integer, Integer> rows = new HashMap<>();

        // Process every reserved seat once.
        for (int[] seat : reservedSeats) {
            int row = seat[0];
            int col = seat[1];

            // Only seats 2 through 9 can affect a four-person group.
            if (col >= 2 && col <= 9) {
                // Set the bit corresponding to this reserved seat.
                rows.put(row, rows.getOrDefault(row, 0) | (1 << col));
            }
        }

        // Rows not present in the map have no useful reserved seats.
        // Every such row can fit two groups.
        int answer = 2 * (n - rows.size());

        // Build masks for the three possible group positions.
        // left = seats 2-5, middle = seats 4-7, right = seats 6-9.
        int left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
        int middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
        int right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

        // Check every row that has relevant reserved seats.
        for (int mask : rows.values()) {
            // Check whether every seat in each block is free.
            boolean canLeft = (mask & left) == 0;
            boolean canMiddle = (mask & middle) == 0;
            boolean canRight = (mask & right) == 0;

            // The left and right blocks are independent, so both can be used.
            if (canLeft && canRight) {
                answer += 2;
            }
            // Otherwise, any available block lets me place one group.
            else if (canLeft || canMiddle || canRight) {
                answer += 1;
            }
            // Otherwise, this row cannot accommodate a group.
        }

        // Return the maximum number of groups.
        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} reservedSeats
 * @return {number}
 */
var maxNumberOfFamilies = function(n, reservedSeats) {
    // Store the reserved seats of every affected row as a bitmask.
    const rows = new Map();

    // Process every reserved seat once.
    for (const [row, col] of reservedSeats) {
        // Seats 1 and 10 never belong to a valid four-seat block.
        if (col >= 2 && col <= 9) {
            // Get the current mask, or use 0 if this is the first seat in the row.
            const currentMask = rows.get(row) || 0;

            // Set the bit corresponding to the reserved seat.
            rows.set(row, currentMask | (1 << col));
        }
    }

    // Every row without relevant reservations can hold two groups.
    let answer = 2 * (n - rows.size);

    // Create masks for seats 2-5, 4-7, and 6-9.
    const left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
    const middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
    const right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

    // Process only rows that contain relevant reserved seats.
    for (const mask of rows.values()) {
        // A block is free when none of its seats are reserved.
        const canLeft = (mask & left) === 0;
        const canMiddle = (mask & middle) === 0;
        const canRight = (mask & right) === 0;

        // Left and right blocks do not overlap, so both groups can fit.
        if (canLeft && canRight) {
            answer += 2;
        }
        // If at least one block is free, one group can fit.
        else if (canLeft || canMiddle || canRight) {
            answer += 1;
        }
        // Otherwise, no group can be placed in this row.
    }

    // Return the maximum number of groups.
    return answer;
};
```

### Python3

```python
from typing import List

class Solution:

    def maxNumberOfFamilies(self, n: int, reservedSeats: List[List[int]]) -> int:
        # Store the reserved seats of each affected row as a bitmask.
        rows = {}

        # Process every reserved seat once.
        for row, col in reservedSeats:
            # Seats 1 and 10 cannot be part of any four-person group.
            if 2 <= col <= 9:
                # Get the current mask, or 0 if this row has not been seen yet.
                current_mask = rows.get(row, 0)

                # Set the bit corresponding to the reserved seat.
                rows[row] = current_mask | (1 << col)

        # Every row not stored in the map has no relevant reservations.
        # Such a row can always contain two groups.
        answer = 2 * (n - len(rows))

        # Create masks for the three possible blocks.
        # left = seats 2-5, middle = seats 4-7, right = seats 6-9.
        left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
        middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
        right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

        # Check only rows that contain relevant reserved seats.
        for mask in rows.values():
            # Check whether each possible block is completely free.
            can_left = (mask & left) == 0
            can_middle = (mask & middle) == 0
            can_right = (mask & right) == 0

            # Left and right do not overlap, so both groups can be placed.
            if can_left and can_right:
                answer += 2

            # Otherwise, any one available block allows one group.
            elif can_left or can_middle or can_right:
                answer += 1

            # If none is available, this row gets zero groups.

        # Return the maximum number of four-person groups.
        return answer
```

### Go

```go
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
 // Store the reserved seats of each affected row as a bitmask.
 rows := make(map[int]int)

 // Process every reserved seat once.
 for _, seat := range reservedSeats {
  row := seat[0]
  col := seat[1]

  // Seats 1 and 10 cannot belong to any valid four-seat block.
  if col >= 2 && col <= 9 {
   // Set the bit corresponding to the reserved seat.
   rows[row] |= 1 << col
  }
 }

 // Every row not stored in the map has no relevant reservations.
 // Each such row can always fit two groups.
 answer := 2 * (n - len(rows))

 // Create masks for the three possible group blocks.
 // left = seats 2-5, middle = seats 4-7, right = seats 6-9.
 left := (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
 middle := (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
 right := (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

 // Check only rows that contain relevant reserved seats.
 for _, mask := range rows {
  // A block is free when none of its seats are reserved.
  canLeft := (mask & left) == 0
  canMiddle := (mask & middle) == 0
  canRight := (mask & right) == 0

  // Left and right blocks do not overlap, so both groups can fit.
  if canLeft && canRight {
   answer += 2
  } else if canLeft || canMiddle || canRight {
   // If any block is free, I can place one group.
   answer++
  }
  // If all three blocks are blocked, this row gets zero groups.
 }

 // Return the maximum number of groups.
 return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages. The main difference is the syntax used for the hash map and bitwise operations.

### 1. Store only affected rows

I create a hash map in each language.

For C++, I use `unordered_map`.

For Java, I use `HashMap`.

For JavaScript, I use `Map`.

For Python, I use a dictionary.

For Go, I use a map.

The key is the row number, and the value is the bitmask of its reserved seats.

This avoids allocating memory for up to `10^9` rows.

### 2. Process the reserved seats

For every entry:

```text
[row, seat]
```

I first check whether the seat is between `2` and `9`.

If the seat is `1` or `10`, I ignore it because neither seat can be part of a four-person group.

For a relevant seat, I set its bit in the row's mask.

Conceptually, the operation is:

```text
rowMask = rowMask OR (1 << seat)
```

The `OR` operation keeps all previously marked reserved seats while adding the new one.

### 3. Count completely free rows

After building the map, suppose there are `r` affected rows.

Then:

```text
n - r
```

rows have no relevant reservations.

Every one of those rows can hold two groups:

```text
2 3 4 5
6 7 8 9
```

So I start the answer with:

```text
2 * (n - r)
```

This is a major optimization because I do not need to inspect those rows individually.

### 4. Build the three seat masks

I create three masks:

```text
Left   = 2, 3, 4, 5
Middle = 4, 5, 6, 7
Right  = 6, 7, 8, 9
```

Each mask represents the seats that must all be free for a group to use that position.

### 5. Check a row using bitwise AND

For an affected row, I check:

```text
rowMask AND leftMask
rowMask AND middleMask
rowMask AND rightMask
```

If the result is zero, there is no reserved seat shared with that block.

For example:

```text
(rowMask & leftMask) == 0
```

means seats `2`, `3`, `4`, and `5` are all available.

This check takes constant time.

### 6. Place two groups when possible

The most important greedy observation is that the left and right blocks do not overlap:

```text
2 3 4 5       6 7 8 9
```

Therefore, when both are available, I can immediately place two groups.

I do not need to consider the middle block in this situation because it overlaps with both of them.

### 7. Otherwise place one group

If I cannot place two groups, I check whether any of the three blocks is available.

If at least one is free, I add one group.

If all three are blocked, I add nothing.

This greedy decision is safe because there can never be more than two groups in one row.

### 8. C++ behavior

In C++, `unordered_map<int, int>` gives average `O(1)` lookup and insertion.

The integer bitmask makes the seat checks very small and fast.

I use a range-based loop to process the stored rows without creating another data structure.

### 9. Java behavior

In Java, `HashMap<Integer, Integer>` stores the affected rows.

The bitwise operations work directly on Java's `int` type, which is more than enough because I only need a few bits for the ten seats.

`getOrDefault` makes it easy to retrieve an existing row mask or start from zero.

### 10. JavaScript behavior

In JavaScript, `Map` is used instead of a normal object.

The bitwise operators work with 32-bit integer operations, which is sufficient because the seat numbers are only from `1` to `10`.

The cinema row count `n` can be `10^9`, which is also safely represented by JavaScript's `Number` type.

### 11. Python3 behavior

Python dictionaries provide the required hash map behavior.

Python integers can grow beyond fixed-size integer limits, although this problem only needs a very small bitmask.

The expression:

```text
mask & block
```

directly checks whether the reserved seats overlap with the selected group block.

### 12. Go behavior

In Go, I use:

```text
map[int]int
```

The key is the row and the value is its bitmask.

Go's integer bitwise operations make the same mask-based approach straightforward.

### 13. Why the solution works

There are only three possible group positions in a row.

The only way to place two groups is to use the left and right blocks because they are the only two blocks that do not overlap.

Therefore, for every affected row, there are only three possible outcomes:

```text
2 groups -> left and right are both free
1 group  -> at least one block is free
0 groups -> all three blocks are blocked
```

That completely covers every possible reservation pattern.

## Examples

### Example 1

Input:

```text
n = 3
reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
```

Output:

```text
4
```

Trace:

* Row `1`: seats `2`, `3`, and `8` are reserved. Only the middle block `4-7` is available, so this row gives `1` group.
* Row `2`: seat `6` is reserved. The left block `2-5` is available, so this row gives `1` group.
* Row `3`: only seats `1` and `10` are reserved. These seats do not affect group placement, so this row gives `2` groups.

Total:

```text
1 + 1 + 2 = 4
```

### Example 2

Input:

```text
n = 2
reservedSeats = [[2,1],[1,8],[2,6]]
```

Output:

```text
2
```

Trace:

* Row `1`: seat `8` is reserved, so the left block `2-5` is available. This row gives `1` group.
* Row `2`: seat `1` does not matter, but seat `6` blocks the left, middle, and right combinations from allowing two groups. One valid block remains, so this row gives `1` group.

Total:

```text
1 + 1 = 2
```

### Example 3

Input:

```text
n = 4
reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
```

Output:

```text
4
```

The algorithm processes only rows `1` and `4`.

The remaining rows have no relevant reservations and can each hold two groups.

The affected rows are checked using the three four-seat blocks, and the final result is `4`.

## How to Use / Run Locally

The repository contains the same algorithm written in C++, Java, JavaScript, Python3, and Go.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

If you are using Windows:

```bash
solution.exe
```

### Java

Save the solution as:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For LeetCode, you normally only need to submit the `Solution` class. LeetCode handles the input and function execution automatically.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```bash
node solution.js
```

You can also paste the function directly into the LeetCode JavaScript editor.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

For LeetCode, paste the `Solution` class into the Python3 editor.

### Go

Save the solution as:

```text
solution.go
```

Run it with:

```bash
go run solution.go
```

For the LeetCode version, use the required function signature provided by the platform.

## Notes & Optimizations

The biggest optimization is avoiding an array of size `n`.

Since `n` can be `10^9`, storing information for every row would waste a huge amount of memory.

I only store rows that appear in the reservation list and contain a seat from `2` through `9`.

Another useful observation is that seats `1` and `10` can be ignored completely. They never participate in any valid four-person block.

The bitmask approach also keeps the row checks constant time. Instead of checking four seats one by one for every possible block, I use a bitwise `AND` operation.

The greedy rule is simple:

* If left and right are free, take both.
* Otherwise, take any available block.
* If none is available, take zero.

There is no need for dynamic programming because the choices inside a row are very limited and the rows do not affect each other.

One alternative solution is to use a set of reserved seat numbers for each affected row instead of bitmasks. That approach is also valid, but the bitmask solution is more compact and makes the block checks very fast.

The most important edge cases are:

* A completely empty row can always fit two groups.
* Reservations on seats `1` and `10` do not affect the answer.
* Reservations can block the left, middle, or right block.
* The middle block cannot be combined with the left or right block because they share seats.
* The left and right blocks can always be combined when both are free.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)

---
