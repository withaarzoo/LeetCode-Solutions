# 2075. Decode the Slanted Ciphertext

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

Given an encoded string `encodedText` and an integer `rows`, I need to decode the original message.

The encoded text was created using a matrix with a fixed number of rows. The original string was placed diagonally from top-left to bottom-right. After that, the entire matrix was read row by row to generate the encoded text.

My task is to reconstruct the original text.

## Constraints

* `0 <= encodedText.length <= 10^6`
* `encodedText` consists of lowercase English letters and spaces.
* `encodedText` is a valid encoding of some original text.
* `1 <= rows <= 1000`
* The original text does not have trailing spaces.

## Intuition

I thought about how the matrix was created.

Since the encoded text is formed by reading the matrix row by row, I can easily calculate the number of columns:

```text
cols = encodedText.length / rows
```

Now I know the matrix dimensions.

Instead of actually building the whole matrix, I can directly access any character using:

```text
index = row * cols + col
```

Then I just simulate the original diagonal movement.

For every column in the first row:

* Start at `(0, startCol)`
* Move diagonally down-right
* Collect characters

At the end, I remove trailing spaces.

## Approach

1. Find the total number of columns.
2. Start from every column in the first row.
3. For each starting column:

   * Move diagonally down-right.
   * Append characters into the result.
4. Continue until the row or column goes out of bounds.
5. Remove trailing spaces from the final answer.

## Data Structures Used

* String / StringBuilder / Array for storing the decoded result.
* No extra matrix is required.

## Operations & Behavior Summary

| Operation            | Purpose                                    |
| -------------------- | ------------------------------------------ |
| `cols = n / rows`    | Finds total columns in the matrix          |
| `row * cols + col`   | Converts matrix position into string index |
| Diagonal traversal   | Reconstructs original message              |
| Trim trailing spaces | Removes unnecessary ending spaces          |

## Complexity

* Time Complexity: `O(n)`

  * Every character is visited at most once.
* Space Complexity: `O(n)`

  * Extra space is used for storing the decoded string.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string decodeCiphertext(string encodedText, int rows) {
        if (encodedText.empty()) return "";

        int n = encodedText.size();
        int cols = n / rows;

        string result;

        for (int startCol = 0; startCol < cols; startCol++) {
            int row = 0;
            int col = startCol;

            while (row < rows && col < cols) {
                result += encodedText[row * cols + col];
                row++;
                col++;
            }
        }

        while (!result.empty() && result.back() == ' ') {
            result.pop_back();
        }

        return result;
    }
};
```

### Java

```java
class Solution {
    public String decodeCiphertext(String encodedText, int rows) {
        if (encodedText.length() == 0) return "";

        int n = encodedText.length();
        int cols = n / rows;

        StringBuilder result = new StringBuilder();

        for (int startCol = 0; startCol < cols; startCol++) {
            int row = 0;
            int col = startCol;

            while (row < rows && col < cols) {
                result.append(encodedText.charAt(row * cols + col));
                row++;
                col++;
            }
        }

        while (result.length() > 0 && result.charAt(result.length() - 1) == ' ') {
            result.deleteCharAt(result.length() - 1);
        }

        return result.toString();
    }
}
```

### JavaScript

```javascript
var decodeCiphertext = function(encodedText, rows) {
    if (encodedText.length === 0) return "";

    const n = encodedText.length;
    const cols = Math.floor(n / rows);

    const result = [];

    for (let startCol = 0; startCol < cols; startCol++) {
        let row = 0;
        let col = startCol;

        while (row < rows && col < cols) {
            result.push(encodedText[row * cols + col]);
            row++;
            col++;
        }
    }

    return result.join('').replace(/\s+$/, '');
};
```

### Python3

```python
class Solution:
    def decodeCiphertext(self, encodedText: str, rows: int) -> str:
        if not encodedText:
            return ""

        n = len(encodedText)
        cols = n // rows

        result = []

        for start_col in range(cols):
            row = 0
            col = start_col

            while row < rows and col < cols:
                result.append(encodedText[row * cols + col])
                row += 1
                col += 1

        return ''.join(result).rstrip()
```

### Go

```go
func decodeCiphertext(encodedText string, rows int) string {
    if len(encodedText) == 0 {
        return ""
    }

    n := len(encodedText)
    cols := n / rows

    result := make([]byte, 0)

    for startCol := 0; startCol < cols; startCol++ {
        row, col := 0, startCol

        for row < rows && col < cols {
            index := row*cols + col
            result = append(result, encodedText[index])
            row++
            col++
        }
    }

    end := len(result) - 1
    for end >= 0 && result[end] == ' ' {
        end--
    }

    return string(result[:end+1])
}
```

## Step-by-step Detailed Explanation

### C++

* First, I check if the string is empty.
* Then I calculate the number of columns.
* I use two nested loops:

  * Outer loop starts from every column.
  * Inner loop moves diagonally.
* I collect characters directly from the encoded string.
* Finally, I remove trailing spaces.

### Java

* I use `StringBuilder` because appending strings repeatedly is faster.
* I calculate the number of columns.
* I move diagonally through the virtual matrix.
* After collecting everything, I remove spaces from the end.

### JavaScript

* I store characters inside an array.
* Later I join the array into a string.
* I use regex to remove trailing spaces.

### Python3

* I use a list because appending to a list is efficient.
* After collecting characters, I use `.rstrip()` to remove ending spaces.

### Go

* I use a byte slice to build the result.
* I manually remove trailing spaces by decreasing the ending index.

## Examples

### Example 1

```text
Input:
encodedText = "ch   ie   pr"
rows = 3

Output:
"cipher"
```

### Example 2

```text
Input:
encodedText = "iveo    eed   l te   olc"
rows = 4

Output:
"i love leetcode"
```

### Example 3

```text
Input:
encodedText = "coding"
rows = 1

Output:
"coding"
```

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* I do not create an actual matrix because that would use extra memory.
* I directly calculate the index of each character.
* This keeps the solution fast and memory efficient.
* The solution works well even when the input size is very large.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
