# Problem Title

**2043. Simple Bank System**

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I need to implement a `Bank` class that simulates a simple banking system. The bank is initialized with `n` accounts; balances are given as a 0-indexed array `balance`, but account numbers are **1-indexed** in the API. The class supports three operations:

* `transfer(account1, account2, money)` — move `money` from `account1` to `account2` if both accounts exist and `account1` has enough balance.
* `deposit(account, money)` — add `money` to `account` if it exists.
* `withdraw(account, money)` — subtract `money` from `account` if it exists and has enough balance.

Each operation returns `true` if it succeeds, otherwise `false`.

---

## Constraints

* `n == balance.length`
* `1 <= n, account, account1, account2 <= 10^5`
* `0 <= balance[i], money <= 10^12`
* At most `10^4` calls will be made to each function (`transfer`, `deposit`, `withdraw`).

Important details:

* Accounts in calls are 1-indexed. Internally I will use 0-indexed array indices.
* Use integer types that can hold up to `10^12` (e.g., `long long` in C++, `long` in Java, `int64`/`int` in Go depending on usage, `int` in Python is unbounded).

---

## Intuition

I thought about this problem like managing a simple array of balances. Each account maps to an index in an array. For every operation, I only need to:

1. Verify accounts exist (bounds check).
2. Verify funds are sufficient (for withdraw/transfer).
3. Update the balances if valid.

All checks and updates are constant-time. So each method performs O(1) work.

---

## Approach

1. Store a copy of the initial `balance` array internally as the bank's state (0-indexed).
2. For `transfer`:

   * Validate both account indices (1..n).
   * Convert to 0-based indices.
   * Check if `account1` has at least `money`.
   * Subtract from `account1`, add to `account2`.
3. For `deposit` and `withdraw`:

   * Validate account index.
   * For `withdraw`, check sufficient funds.
   * Update balance and return `true` on success.
4. Return `false` for any invalid account index or insufficient funds.

This is simple, direct, and efficient.

---

## Data Structures Used

* A single dynamic array (vector, slice, list, or basic array) to hold balances:

  * C++: `vector<long long>`
  * Java: `long[]`
  * JavaScript: `Array` of numbers
  * Python3: `List[int]`
  * Go: `[]int64`

No other auxiliary structures are necessary.

---

## Operations & Behavior Summary

* `Bank(balance)` — initialize internal array copy.
* `transfer(account1, account2, money)`:

  * Returns `false` if either account is invalid or `account1` lacks funds.
  * Otherwise performs the transfer and returns `true`.
* `deposit(account, money)`:

  * Returns `false` if account invalid; otherwise adds money and returns `true`.
* `withdraw(account, money)`:

  * Returns `false` if account invalid or insufficient funds; otherwise subtracts money and returns `true`.

All account parameters are 1-indexed in the API.

---

## Complexity

* **Time Complexity:**

  * Each operation (`transfer`, `deposit`, `withdraw`) runs in **O(1)** time because they do only a constant number of checks and updates.
* **Space Complexity:**

  * **O(n)** to store the balance array (where `n` is the number of accounts). No additional structures used.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Bank {
private:
    vector<long long> bal; // 0-indexed balances

    inline bool valid(int account) {
        return account >= 1 && account <= (int)bal.size();
    }

public:
    Bank(vector<long long>& balance) : bal(balance) { }

    bool transfer(int account1, int account2, long long money) {
        if (!valid(account1) || !valid(account2)) return false;
        int a = account1 - 1;
        int b = account2 - 1;
        if (bal[a] < money) return false;
        bal[a] -= money;
        bal[b] += money;
        return true;
    }

    bool deposit(int account, long long money) {
        if (!valid(account)) return false;
        bal[account - 1] += money;
        return true;
    }

    bool withdraw(int account, long long money) {
        if (!valid(account)) return false;
        int a = account - 1;
        if (bal[a] < money) return false;
        bal[a] -= money;
        return true;
    }
};
```

---

### Java

```java
class Bank {
    private long[] bal; // 0-indexed balances

    private boolean valid(int account) {
        return account >= 1 && account <= bal.length;
    }

    public Bank(long[] balance) {
        this.bal = new long[balance.length];
        System.arraycopy(balance, 0, this.bal, 0, balance.length);
    }

    public boolean transfer(int account1, int account2, long money) {
        if (!valid(account1) || !valid(account2)) return false;
        int a = account1 - 1;
        int b = account2 - 1;
        if (bal[a] < money) return false;
        bal[a] -= money;
        bal[b] += money;
        return true;
    }

    public boolean deposit(int account, long money) {
        if (!valid(account)) return false;
        bal[account - 1] += money;
        return true;
    }

    public boolean withdraw(int account, long money) {
        if (!valid(account)) return false;
        int a = account - 1;
        if (bal[a] < money) return false;
        bal[a] -= money;
        return true;
    }
}
```

---

### JavaScript

```javascript
var Bank = function(balance) {
    this.bal = balance.slice(); // copy to avoid external mutations
};

Bank.prototype.transfer = function(account1, account2, money) {
    const n = this.bal.length;
    if (account1 < 1 || account1 > n || account2 < 1 || account2 > n) return false;
    const a = account1 - 1, b = account2 - 1;
    if (this.bal[a] < money) return false;
    this.bal[a] -= money;
    this.bal[b] += money;
    return true;
};

Bank.prototype.deposit = function(account, money) {
    const n = this.bal.length;
    if (account < 1 || account > n) return false;
    this.bal[account - 1] += money;
    return true;
};

Bank.prototype.withdraw = function(account, money) {
    const n = this.bal.length;
    if (account < 1 || account > n) return false;
    const a = account - 1;
    if (this.bal[a] < money) return false;
    this.bal[a] -= money;
    return true;
};
```

---

### Python3

```python
from typing import List

class Bank:
    def __init__(self, balance: List[int]):
        self.bal = balance[:]  # copy to internal state

    def _valid(self, account: int) -> bool:
        return 1 <= account <= len(self.bal)

    def transfer(self, account1: int, account2: int, money: int) -> bool:
        if not (self._valid(account1) and self._valid(account2)):
            return False
        a, b = account1 - 1, account2 - 1
        if self.bal[a] < money:
            return False
        self.bal[a] -= money
        self.bal[b] += money
        return True

    def deposit(self, account: int, money: int) -> bool:
        if not self._valid(account):
            return False
        self.bal[account - 1] += money
        return True

    def withdraw(self, account: int, money: int) -> bool:
        if not self._valid(account):
            return False
        a = account - 1
        if self.bal[a] < money:
            return False
        self.bal[a] -= money
        return True
```

---

### Go

```go
package main

type Bank struct {
    bal []int64
}

func Constructor(balance []int64) Bank {
    b := make([]int64, len(balance))
    copy(b, balance)
    return Bank{bal: b}
}

func (this *Bank) valid(account int) bool {
    return account >= 1 && account <= len(this.bal)
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if !this.valid(account1) || !this.valid(account2) {
        return false
    }
    a := account1 - 1
    b := account2 - 1
    if this.bal[a] < money {
        return false
    }
    this.bal[a] -= money
    this.bal[b] += money
    return true
}

func (this *Bank) Deposit(account int, money int64) bool {
    if !this.valid(account) {
        return false
    }
    this.bal[account-1] += money
    return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
    if !this.valid(account) {
        return false
    }
    a := account - 1
    if this.bal[a] < money {
        return false
    }
    this.bal[a] -= money
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I will explain the Python version step-by-step (other languages use the same logic; names / syntax differ):

1. `class Bank:` — I define a `Bank` class to encapsulate balances and methods.
2. `def __init__(self, balance: List[int]):` — Constructor receives the initial balances list.
3. `self.bal = balance[:]` — I make a shallow copy so external changes to the original list won't affect the bank's internal state.
4. `def _valid(self, account: int) -> bool:` — Helper method to validate 1-indexed account numbers.

   * `return 1 <= account <= len(self.bal)`
5. `def transfer(self, account1, account2, money):`

   * First check both accounts exist: `if not (self._valid(account1) and self._valid(account2)): return False`
   * Convert to 0-based indices: `a, b = account1 - 1, account2 - 1`
   * Check sufficient funds: `if self.bal[a] < money: return False`
   * Update balances: `self.bal[a] -= money; self.bal[b] += money`
   * Return `True` on success.
6. `def deposit(self, account, money):`

   * Validate account: `if not self._valid(account): return False`
   * Add money: `self.bal[account - 1] += money`
   * Return `True`.
7. `def withdraw(self, account, money):`

   * Validate account.
   * Check funds: `if self.bal[a] < money: return False`
   * Subtract money: `self.bal[a] -= money`
   * Return `True`.

C++/Java/JS/Go versions follow the exact same logic:

* validate indices,
* check funds for withdraw/transfer,
* update array in O(1).

---

## Examples

**Example 1 (from problem statement):**

Input:

```
["Bank","withdraw","transfer","deposit","transfer","withdraw"]
[[[10,100,20,50,30],[3,10],[5,1,20],[5,20],[3,4,15],[10,50]]]
```

Explanation of operations:

1. `Bank([10,100,20,50,30])` — initialize accounts 1..5 with balances.
2. `withdraw(3, 10)` — succeeds: account 3 had 20 -> becomes 10.
3. `transfer(5, 1, 20)` — succeeds: account 5 had 30 -> becomes 10; account 1 becomes 30.
4. `deposit(5, 20)` — succeeds: account 5 becomes 30.
5. `transfer(3, 4, 15)` — fails: account 3 has only 10 left.
6. `withdraw(10, 50)` — fails: account 10 does not exist.

Result: `[null, true, true, true, false, false]`

---

## How to use / Run locally

### C++

* Create a `.cpp` file with the class code.
* Compile and run with a test harness (e.g., LeetCode runner or custom main).

### Java

* Place the `Bank` class in a `.java` file.
* Use the LeetCode runner or write a `main` to instantiate and call methods.

### JavaScript

* Use in Node or browser environment. Example:

```js
const bank = new Bank([10,100,20,50,30]);
console.log(bank.withdraw(3,10)); // true
```

### Python3

* Save class in a `.py` file and import or run interactively:

```py
bank = Bank([10,100,20,50,30])
print(bank.withdraw(3,10))  # True
```

### Go

* Put the code into `main` package and use the constructor `Constructor` to create the object.
* Run `go run file.go` after adding a small test.

> On LeetCode, submit the class directly — the judge will instantiate and call methods.

---

## Notes & Optimizations

* I used data types that can hold values up to `10^12` (e.g., `long long` in C++, `long` in Java, `int64` in Go). In Python integers are unbounded.
* I copy the input balance array in the constructor to avoid side effects from external modification.
* Every method does constant-time checks and updates: O(1) time per operation and O(n) memory overall.
* No synchronization or concurrency control is included because the problem doesn't require thread-safety. If the system were concurrent, we'd need mutexes/locks or atomic operations.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
