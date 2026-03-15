# 1622. Fancy Sequence

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

We need to design a data structure that supports the following operations efficiently:

1. append(val) – add a value to the sequence
2. addAll(inc) – increase every element in the sequence by inc
3. multAll(m) – multiply every element in the sequence by m
4. getIndex(idx) – return the value at index idx

All results must be returned modulo 1e9 + 7.

The number of operations can go up to 100000, so updating the entire array for each operation would be too slow.

---

## Constraints

1 <= val, inc, m <= 100

0 <= idx <= 100000

At most 100000 operations will be called.

All answers must be returned modulo 1000000007.

---

## Intuition

If we apply addAll or multAll directly to every element, each operation becomes O(n).

With up to 100000 operations, this would be too slow.

Instead, we observe that all operations form a linear transformation.

Any number in the array follows this transformation:

value = stored_value * mul + add

Where

mul = total multiplication applied

add = total addition applied

Instead of updating every element, we maintain these two global values.

When we append a new value, we store it in a reversed form so that future transformations still work correctly.

---

## Approach

1. Maintain two global variables

mul = 1

add = 0

1. Store elements in an array but in normalized form.

2. When append(val) is called:

stored = (val - add) * modular_inverse(mul)

This cancels previous operations.

1. For addAll(inc):

add += inc

1. For multAll(m):

mul *= m

add *= m

1. For getIndex(idx):

return stored[idx] * mul + add

All calculations are done using modulo 1e9 + 7.

---

## Data Structures Used

Dynamic Array / List

The array stores normalized values of elements.

Global Variables

mul : cumulative multiplication

add : cumulative addition

---

## Operations & Behavior Summary

append(val)

Stores normalized value in the array.

addAll(inc)

Updates global addition factor.

multAll(m)

Updates multiplication and addition factors.

getIndex(idx)

Returns transformed value if index exists.

Otherwise returns -1.

---

## Complexity

Time Complexity

append : O(log MOD)

addAll : O(1)

multAll : O(1)

getIndex : O(1)

Where MOD = 1e9 + 7

Space Complexity

O(n)

We store all appended elements.

---

## Multi-language Solutions

### C++

```cpp
class Fancy {
private:
    const long long MOD = 1e9 + 7;
    vector<long long> seq;
    long long mul = 1;
    long long add = 0;

    long long modPow(long long a, long long b) {
        long long res = 1;
        while (b) {
            if (b & 1) res = (res * a) % MOD;
            a = (a * a) % MOD;
            b >>= 1;
        }
        return res;
    }

public:

    Fancy() {}

    void append(int val) {
        long long inv = modPow(mul, MOD - 2);
        long long stored = ((val - add + MOD) % MOD * inv) % MOD;
        seq.push_back(stored);
    }

    void addAll(int inc) {
        add = (add + inc) % MOD;
    }

    void multAll(int m) {
        mul = (mul * m) % MOD;
        add = (add * m) % MOD;
    }

    int getIndex(int idx) {
        if (idx >= seq.size()) return -1;
        return (seq[idx] * mul % MOD + add) % MOD;
    }
};
```

### Java

```java
class Fancy {

    static final long MOD = 1000000007;

    ArrayList<Long> seq = new ArrayList<>();
    long mul = 1;
    long add = 0;

    private long modPow(long a, long b) {
        long res = 1;
        while (b > 0) {
            if ((b & 1) == 1) res = (res * a) % MOD;
            a = (a * a) % MOD;
            b >>= 1;
        }
        return res;
    }

    public Fancy() {}

    public void append(int val) {
        long inv = modPow(mul, MOD - 2);
        long stored = ((val - add + MOD) % MOD * inv) % MOD;
        seq.add(stored);
    }

    public void addAll(int inc) {
        add = (add + inc) % MOD;
    }

    public void multAll(int m) {
        mul = (mul * m) % MOD;
        add = (add * m) % MOD;
    }

    public int getIndex(int idx) {
        if (idx >= seq.size()) return -1;
        return (int)((seq.get(idx) * mul % MOD + add) % MOD);
    }
}
```

### JavaScript

```javascript
var Fancy = function() {
    this.MOD = 1000000007n;
    this.seq = [];
    this.mul = 1n;
    this.add = 0n;
};

Fancy.prototype.modPow = function(a, b) {
    let res = 1n;
    while (b > 0n) {
        if (b & 1n) res = (res * a) % this.MOD;
        a = (a * a) % this.MOD;
        b >>= 1n;
    }
    return res;
};

Fancy.prototype.append = function(val) {
    val = BigInt(val);
    let inv = this.modPow(this.mul, this.MOD - 2n);
    let stored = ((val - this.add + this.MOD) % this.MOD * inv) % this.MOD;
    this.seq.push(stored);
};

Fancy.prototype.addAll = function(inc) {
    this.add = (this.add + BigInt(inc)) % this.MOD;
};

Fancy.prototype.multAll = function(m) {
    m = BigInt(m);
    this.mul = (this.mul * m) % this.MOD;
    this.add = (this.add * m) % this.MOD;
};

Fancy.prototype.getIndex = function(idx) {
    if (idx >= this.seq.length) return -1;
    let val = (this.seq[idx] * this.mul + this.add) % this.MOD;
    return Number(val);
};
```

### Python3

```python
class Fancy:

    MOD = 10**9 + 7

    def __init__(self):
        self.seq = []
        self.mul = 1
        self.add = 0

    def mod_pow(self, a, b):
        res = 1
        while b:
            if b & 1:
                res = res * a % self.MOD
            a = a * a % self.MOD
            b >>= 1
        return res

    def append(self, val: int) -> None:
        inv = self.mod_pow(self.mul, self.MOD - 2)
        stored = ((val - self.add) % self.MOD) * inv % self.MOD
        self.seq.append(stored)

    def addAll(self, inc: int) -> None:
        self.add = (self.add + inc) % self.MOD

    def multAll(self, m: int) -> None:
        self.mul = self.mul * m % self.MOD
        self.add = self.add * m % self.MOD

    def getIndex(self, idx: int) -> int:
        if idx >= len(self.seq):
            return -1
        return (self.seq[idx] * self.mul + self.add) % self.MOD
```

### Go

```go
const MOD int64 = 1e9 + 7

type Fancy struct {
    seq []int64
    mul int64
    add int64
}

func modPow(a, b int64) int64 {
    res := int64(1)
    for b > 0 {
        if b&1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        b >>= 1
    }
    return res
}

func Constructor() Fancy {
    return Fancy{[]int64{}, 1, 0}
}

func (this *Fancy) Append(val int) {
    inv := modPow(this.mul, MOD-2)
    stored := ((int64(val)-this.add+MOD)%MOD * inv) % MOD
    this.seq = append(this.seq, stored)
}

func (this *Fancy) AddAll(inc int) {
    this.add = (this.add + int64(inc)) % MOD
}

func (this *Fancy) MultAll(m int) {
    this.mul = this.mul * int64(m) % MOD
    this.add = this.add * int64(m) % MOD
}

func (this *Fancy) GetIndex(idx int) int {
    if idx >= len(this.seq) {
        return -1
    }
    return int((this.seq[idx]*this.mul%MOD + this.add) % MOD)
}
```

---

## Step-by-step Detailed Explanation

1. We store numbers in normalized form.

2. Two global variables represent all operations applied to the sequence.

mul represents cumulative multiplication.

add represents cumulative addition.

1. When append(val) happens, we reverse the transformation.

stored = (val - add) * inverse(mul)

1. addAll only updates the add variable.

2. multAll updates both mul and add.

3. When retrieving an element we apply the transformation again.

result = stored_value * mul + add

All operations use modulo arithmetic.

---

## Examples

Operations

append(2)

Sequence: [2]

addAll(3)

Sequence: [5]

append(7)

Sequence: [5,7]

multAll(2)

Sequence: [10,14]

getIndex(0)

Output: 10

---

## How to use / Run locally

Clone repository

```
git clone <repo-url>
```

Compile C++

```
g++ solution.cpp -o solution
./solution
```

Run Python

```
python solution.py
```

Run Java

```
javac Fancy.java
java Fancy
```

---

## Notes & Optimizations

The main optimization is lazy transformation.

Instead of updating every element after each operation, we store transformation parameters.

This reduces expensive operations from O(n) to O(1).

Modular inverse is used to cancel previous multiplications when inserting new elements.

---

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
