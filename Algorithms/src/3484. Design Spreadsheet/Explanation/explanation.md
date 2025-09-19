# Design Spreadsheet — LeetCode 3484

> A clean, professional README you can put in a GitHub repo for my solution to **LeetCode 3484 — Design Spreadsheet**.
> It contains the problem summary, my intuition, approach, complexity analysis, a step-by-step explanation (in my voice), and working solutions in **C++, Java, JavaScript, Python3, and Go**.
> Each code file is self-contained and includes comments.

---

## Problem summary

A spreadsheet has 26 columns labeled `A` to `Z` and a given number of rows. Each cell holds an integer (initially `0`). Implement a `Spreadsheet` class that supports:

* `Spreadsheet(int rows)` — initialize spreadsheet with the given rows (columns fixed `A`–`Z`)
* `void setCell(String cell, int value)` — set the specified cell (e.g., `"A1"`) to `value`
* `void resetCell(String cell)` — reset the specified cell to `0` (remove any explicit value)
* `int getValue(String formula)` — evaluate a formula of the form `=X+Y`, where `X` and `Y` are either non-negative integer literals or cell references. Return the sum (cells not explicitly set count as `0`)

Important: The formula is always exactly one `=` and one `+`. Number of rows `rows` is `1 <= rows <= 10^3`. Value ranges are within `0..10^5`. Up to `10^4` calls total.

---

## Intuition

I thought of the spreadsheet as a sparse storage: most cells remain `0` so I only store cells that I explicitly set. When `getValue` is called, I just parse the two operands in `=X+Y`. Each operand is either:

* a plain number → parse it, or
* a cell reference like `A1` → look it up in my map and treat missing entries as `0`.

To map a cell reference to a key, I convert the column letter `A..Z` to an index `0..25` and the 1-indexed row to a 0-based index, then combine them (e.g., `key = colIndex * rows + rowIndex`). This gives a unique integer key per cell and makes lookups fast.

---

## Approach

1. Represent cell with an integer key: `key = col * rows + row`.
2. Keep a hash map `key -> value` for all explicitly set cells.
3. `setCell(cell, value)` writes to the map.
4. `resetCell(cell)` removes the map entry (so it becomes `0`).
5. `getValue(formula)`:

   * Remove leading `=`.
   * Split by `+` into two operands.
   * For each operand: if it starts with a digit, parse integer; otherwise convert cell ref to `key` and return `map.get(key, 0)`.
   * Return the sum.

This is simple and efficient given the constraints and the limited formula format.

---

## Complexity

* **Time Complexity:**

  * `setCell` / `resetCell`: average **O(1)** for map operations.
  * `getValue`: **O(L)** where `L` is the formula length (parsing two operands). Practically constant; formula strings are short.
* **Space Complexity:**

  * **O(k)** where `k` is the number of explicitly set cells (only those are stored).

---

## Example

Input (sequence of calls):

```text
Spreadsheet(3)
getValue("=5+7")    -> 12
setCell("A1", 10)
getValue("=A1+6")   -> 16
setCell("B2", 15)
getValue("=A1+B2")  -> 25
resetCell("A1")
getValue("=A1+B2")  -> 15
```

---

## Implementation (multi-language)

> Save each language solution into a file named `Spreadsheet.<ext>` (or as appropriate). Example run instructions are given below.

---

### C++ (file: `Spreadsheet.cpp`)

```c++
#include <unordered_map>
#include <string>

using namespace std;

class Spreadsheet {
private:
    int rows;
    unordered_map<int,int> vals; // key -> value

    // Convert a cell like "A1" to integer key
    int keyFromCell(const string &cell) {
        int col = cell[0] - 'A';
        int row = stoi(cell.substr(1)) - 1; // input is 1-indexed
        return col * rows + row;
    }

    // Evaluate operand: either integer literal or cell reference
    int evalOperand(const string &op) {
        if (isdigit(op[0])) return stoi(op); // numeric literal
        auto it = vals.find(keyFromCell(op));
        return (it == vals.end()) ? 0 : it->second;
    }

public:
    Spreadsheet(int rows) : rows(rows) {}

    void setCell(string cell, int value) {
        vals[keyFromCell(cell)] = value;
    }

    void resetCell(string cell) {
        vals.erase(keyFromCell(cell));
    }

    int getValue(string formula) {
        // formula is like "=X+Y"
        string expr = formula.substr(1); // drop '='
        size_t plus = expr.find('+');
        string a = expr.substr(0, plus);
        string b = expr.substr(plus + 1);
        return evalOperand(a) + evalOperand(b);
    }
};

/**
 * LeetCode will instantiate and invoke the class methods; above implementation is complete.
 */
```

* Compile / Run (example test harness required)

---

### Java (file: `Spreadsheet.java`)

```java
import java.util.HashMap;

class Spreadsheet {
    private int rows;
    private HashMap<Integer, Integer> map;

    public Spreadsheet(int rows) {
        this.rows = rows;
        this.map = new HashMap<>();
    }

    // Convert "A1" to key
    private int keyFromCell(String cell) {
        int col = cell.charAt(0) - 'A';
        int row = Integer.parseInt(cell.substring(1)) - 1;
        return col * rows + row;
    }

    // Evaluate operand either numeric or cell
    private int evalOperand(String op) {
        if (Character.isDigit(op.charAt(0))) {
            return Integer.parseInt(op);
        } else {
            return map.getOrDefault(keyFromCell(op), 0);
        }
    }

    public void setCell(String cell, int value) {
        map.put(keyFromCell(cell), value);
    }

    public void resetCell(String cell) {
        map.remove(keyFromCell(cell));
    }

    public int getValue(String formula) {
        String expr = formula.substring(1); // remove '='
        int plus = expr.indexOf('+');
        String a = expr.substring(0, plus);
        String b = expr.substring(plus + 1);
        return evalOperand(a) + evalOperand(b);
    }
}
```

* Build / Run

* `javac Spreadsheet.java`
* Provide a driver to test (LeetCode provides the harness on the platform).

---

### JavaScript (Node) (file: `Spreadsheet.js`)

```javascript
/**
 * @param {number} rows
 */
var Spreadsheet = function(rows) {
    this.rows = rows;
    this.map = new Map(); // key -> value
};

// helper: convert "A1" to integer key
Spreadsheet.prototype.keyFromCell = function(cell) {
    const col = cell.charCodeAt(0) - 'A'.charCodeAt(0);
    const row = parseInt(cell.slice(1), 10) - 1;
    return col * this.rows + row;
};

// helper: evaluate operand (number or cell)
Spreadsheet.prototype.evalOperand = function(op) {
    if (/^\d/.test(op)) return parseInt(op, 10);
    const key = this.keyFromCell(op);
    return this.map.has(key) ? this.map.get(key) : 0;
};

/** 
 * @param {string} cell 
 * @param {number} value
 * @return {void}
 */
Spreadsheet.prototype.setCell = function(cell, value) {
    this.map.set(this.keyFromCell(cell), value);
};

/** 
 * @param {string} cell
 * @return {void}
 */
Spreadsheet.prototype.resetCell = function(cell) {
    this.map.delete(this.keyFromCell(cell));
};

/** 
 * @param {string} formula
 * @return {number}
 */
Spreadsheet.prototype.getValue = function(formula) {
    const expr = formula.slice(1); // drop '='
    const parts = expr.split('+');
    return this.evalOperand(parts[0]) + this.evalOperand(parts[1]);
};

module.exports = Spreadsheet; // for Node.js tests
```

**Run (Node)**: require it in a test script and call the methods.

---

### Python3 (file: `spreadsheet.py`)

```python
class Spreadsheet:

    def __init__(self, rows: int):
        # number of rows (columns are fixed A-Z)
        self.rows = rows
        # store only explicitly set cells: key -> value
        self._vals = {}

    def _key_from_cell(self, cell: str) -> int:
        # "A1" -> col index 0, row index 0 => key = col * rows + row
        col = ord(cell[0]) - ord('A')
        row = int(cell[1:]) - 1
        return col * self.rows + row

    def _eval_operand(self, op: str) -> int:
        # if operand starts with digit -> literal number; else cell reference
        if op[0].isdigit():
            return int(op)
        return self._vals.get(self._key_from_cell(op), 0)

    def setCell(self, cell: str, value: int) -> None:
        self._vals[self._key_from_cell(cell)] = value

    def resetCell(self, cell: str) -> None:
        self._vals.pop(self._key_from_cell(cell), None)

    def getValue(self, formula: str) -> int:
        # formula is like "=X+Y"
        expr = formula[1:]  # drop '='
        a, b = expr.split('+')
        return self._eval_operand(a) + self._eval_operand(b)
```

**Run (Python)**: import class and call methods in a small test driver. Example:

```python
# test_driver.py
from spreadsheet import Spreadsheet
s = Spreadsheet(3)
print(s.getValue("=5+7"))   # 12
s.setCell("A1", 10)
print(s.getValue("=A1+6"))  # 16
s.setCell("B2", 15)
print(s.getValue("=A1+B2")) # 25
s.resetCell("A1")
print(s.getValue("=A1+B2")) # 15
```

`python3 test_driver.py`

---

## Go (file: `spreadsheet.go`)

```go
package main

import (
 "strconv"
 "strings"
)

// Spreadsheet holds number of rows and a map for explicitly set cells.
type Spreadsheet struct {
 rows int
 mp   map[int]int
}

// Constructor initializes the spreadsheet.
func Constructor(rows int) Spreadsheet {
 return Spreadsheet{
  rows: rows,
  mp:   make(map[int]int),
 }
}

// helper: convert "A1" -> integer key
func keyFromCell(cell string, rows int) int {
 col := int(cell[0] - 'A')
 r, _ := strconv.Atoi(cell[1:]) // input guaranteed valid
 return col*rows + (r - 1)
}

func evalOperand(op string, rows int, mp map[int]int) int {
 // numeric literal?
 if op[0] >= '0' && op[0] <= '9' {
  v, _ := strconv.Atoi(op)
  return v
 }
 k := keyFromCell(op, rows)
 if val, ok := mp[k]; ok {
  return val
 }
 return 0
}

func (this *Spreadsheet) SetCell(cell string, value int) {
 k := keyFromCell(cell, this.rows)
 this.mp[k] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
 k := keyFromCell(cell, this.rows)
 delete(this.mp, k)
}

func (this *Spreadsheet) GetValue(formula string) int {
 expr := formula[1:] // drop '='
 parts := strings.Split(expr, "+")
 return evalOperand(parts[0], this.rows, this.mp) + evalOperand(parts[1], this.rows, this.mp)
}
```

**Build / Run**
`go test` or create a `main` to instantiate and call methods.

---

### Step-by-step detailed explanation (my voice)

I’ll explain the Python solution line-by-line (the other languages implement the same logic).

1. `class Spreadsheet:` — I create a class to keep state between calls.
2. `def __init__(self, rows: int):` — I save `rows` and create a dictionary `_vals` where I store only cells I explicitly set.

   * Reason: storing only set cells is memory efficient because most cells remain `0`.
3. `_key_from_cell(self, cell)`:

   * `col = ord(cell[0]) - ord('A')` converts `'A'` to `0`, `'B'` to `1`, etc.
   * `row = int(cell[1:]) - 1` gets 0-based row index (input uses 1-based rows).
   * `return col * self.rows + row` — combine into a unique integer key for this (col,row).
   * Reason: using a single integer key simplifies map indexing and is fast.
4. `_eval_operand(self, op)`:

   * If the operand starts with a digit, I return `int(op)` (literal).
   * Otherwise, I treat it as a cell reference and return `_vals.get(key, 0)`.
   * Missing entries default to `0`, per problem statement.
5. `setCell(self, cell, value)`:

   * Compute key and set `_vals[key] = value`.
6. `resetCell(self, cell)`:

   * Remove key from `_vals` (pop with default `None`) — effectively makes cell value `0`.
7. `getValue(self, formula)`:

   * I remove the leading `'='` and split on `'+'` into `a` and `b`.
   * I evaluate each operand with `_eval_operand` and return their sum.

Why this works:

* The formula format is fixed and simple (`=X+Y`), so I don't need to implement any complex expression parser or dependency graph.
* Using a hash map provides average O(1) read/write for cells.
* Converting cells to integer keys ensures compact and deterministic mapping.

---

## Notes and tips

* The solution assumes column is a single uppercase letter `A..Z` (per statement). If the problem extended to multi-letter columns (`AA` etc.), `keyFromCell` must parse all letters.
* This implementation purposely avoids storing defaults (0), which saves memory and is faster for sparse usage.
* The approach is ready for the LeetCode constraints and large numbers of calls (10^4).

---

## Files & usage suggestions

* `Spreadsheet.cpp` — compile in a C++ environment and include a test harness.
* `Spreadsheet.java` — compile with `javac` and test.
* `Spreadsheet.js` — `node` supported; export class and test with a script.
* `spreadsheet.py` — import and use directly or run test script.
* `spreadsheet.go` — put in a package or add a `main` to test.

---

## License

Use freely — MIT or any permissive license is fine for this repo.

---

If you want, I can:

* add a sample test harness for each language (so `make test` can run examples), or
* add GitHub Actions to automatically run simple tests.
