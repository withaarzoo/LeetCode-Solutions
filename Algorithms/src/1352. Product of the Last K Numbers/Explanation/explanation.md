# Product of the Last K Numbers - LeetCode 1352 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

This is LeetCode problem 1352 titled Product of the Last K Numbers. You need to design a data structure that accepts a continuous stream of integers and can quickly return the product of the last k numbers in that stream.

The class is called ProductOfNumbers. It has three methods:

- A constructor that starts with an empty stream.
- An add method that appends a new integer to the stream.
- A getProduct method that returns the product of the last k numbers currently in the stream.

You are guaranteed that the stream always has at least k numbers when getProduct is called, and that any product will fit inside a 32-bit integer.

The goal is to make both add and getProduct run in constant time so the solution stays efficient even with tens of thousands of operations.

## Constraints

- 0 <= num <= 100
- 1 <= k <= 4 * 10^4
- At most 4 * 10^4 calls will be made to add and getProduct combined
- The product of any contiguous sequence of numbers will always fit in a 32-bit integer

## Intuition

When I first looked at the Product of the Last K Numbers problem, the naive way of multiplying the last k elements every time felt too slow. With up to 40,000 operations, anything slower than constant time would time out.

I also noticed that zeros completely change the game. Once a zero appears, every product that includes it becomes zero. That observation pushed me toward keeping a running list of prefix products, but only for the numbers that came after the most recent zero. This way I can answer any product query with a single division while still handling zeros cleanly.

## Approach

I maintain a list that stores prefix products of the numbers added since the last zero. The list always begins with the value 1 so the empty product is well-defined.

When a new number arrives:

- If the number is zero, I clear the entire list and put 1 back in. Everything before this zero is no longer useful.
- If the number is non-zero, I multiply the last stored prefix by the new number and append the result.

When getProduct(k) is called I check the current length of the prefix list. The length is always one more than the count of numbers stored after the last zero. If k is larger than or equal to that count, the window must contain a zero, so I return 0. Otherwise I simply divide the newest prefix by the prefix that sits k positions earlier.

Both operations finish in constant time and never overflow because the problem guarantees the product fits in a 32-bit integer.

## Data Structures Used

I use a dynamic list (vector in C++, ArrayList in Java, array in JavaScript/TypeScript, list in Python, slice in Go) to store the prefix products.

This structure was chosen because:

- Appending a new prefix is O(1) amortized.
- Accessing any earlier prefix by index is O(1).
- Clearing the list when a zero appears is also O(1) in practice for the sizes involved.

No other data structures are needed. A simple array of running products is enough to answer every query instantly.

## Operations & Behavior Summary

1. Start with a list containing only the value 1.
2. On every add(num):
   - If num is 0, reset the list to [1].
   - Otherwise multiply the last value by num and append it.
3. On every getProduct(k):
   - If k is greater than or equal to the current list length, return 0.
   - Otherwise return (last prefix) / (prefix located k steps before the end).

That is the entire algorithm. It stays simple, fast, and correct for every test case.

## Complexity

| Operation          | Time Complexity | Space Complexity | Explanation |
|--------------------|-----------------|------------------|-------------|
| add                | O(1)            | O(1) amortized   | One multiplication and one append, or a full reset when zero appears |
| getProduct         | O(1)            | O(1)             | Two index lookups and one division |
| Overall (n calls)  | O(n)            | O(n)             | n is the maximum number of consecutive non-zero numbers; the list never grows larger than that |

## Multi-language Solutions

### C++
```cpp
class ProductOfNumbers {
public:
    vector<int> pre;
    ProductOfNumbers() {
        pre = {1};
    }
    
    void add(int num) {
        if (num == 0) {
            pre = {1};
        } else {
            pre.push_back(pre.back() * num);
        }
    }
    
    int getProduct(int k) {
        if (k >= pre.size()) return 0;
        return pre.back() / pre[pre.size() - 1 - k];
    }
};
```

### Java
```java
class ProductOfNumbers {
    List<Integer> pre;
    public ProductOfNumbers() {
        pre = new ArrayList<>();
        pre.add(1);
    }
    
    public void add(int num) {
        if (num == 0) {
            pre = new ArrayList<>();
            pre.add(1);
        } else {
            pre.add(pre.get(pre.size() - 1) * num);
        }
    }
    
    public int getProduct(int k) {
        if (k >= pre.size()) return 0;
        return pre.get(pre.size() - 1) / pre.get(pre.size() - 1 - k);
    }
}
```

### JavaScript
```javascript
var ProductOfNumbers = function() {
    this.pre = [1];
};

ProductOfNumbers.prototype.add = function(num) {
    if (num === 0) {
        this.pre = [1];
    } else {
        this.pre.push(this.pre[this.pre.length - 1] * num);
    }
};

ProductOfNumbers.prototype.getProduct = function(k) {
    if (k >= this.pre.length) return 0;
    return this.pre[this.pre.length - 1] / this.pre[this.pre.length - 1 - k];
};
```

### TypeScript
```typescript
class ProductOfNumbers {
    pre: number[];
    constructor() {
        this.pre = [1];
    }

    add(num: number): void {
        if (num === 0) {
            this.pre = [1];
        } else {
            this.pre.push(this.pre[this.pre.length - 1] * num);
        }
    }

    getProduct(k: number): number {
        if (k >= this.pre.length) return 0;
        return this.pre[this.pre.length - 1] / this.pre[this.pre.length - 1 - k];
    }
}
```

### Python3
```python
class ProductOfNumbers:

    def __init__(self):
        self.pre = [1]

    def add(self, num: int) -> None:
        if num == 0:
            self.pre = [1]
        else:
            self.pre.append(self.pre[-1] * num)

    def getProduct(self, k: int) -> int:
        if k >= len(self.pre):
            return 0
        return self.pre[-1] // self.pre[-1 - k]
```

### Go
```go
type ProductOfNumbers struct {
    pre []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{pre: []int{1}}
}


func (this *ProductOfNumbers) Add(num int)  {
    if num == 0 {
        this.pre = []int{1}
    } else {
        this.pre = append(this.pre, this.pre[len(this.pre)-1]*num)
    }
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    if k >= len(this.pre) {
        return 0
    }
    return this.pre[len(this.pre)-1] / this.pre[len(this.pre)-1-k]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The logic is identical across all six languages. Only the syntax changes.

I begin by creating a list and placing the number 1 inside it. That single 1 represents the product of zero numbers and makes later divisions safe.

When add is called I first test whether the incoming number is zero.  
If it is, every earlier product is now useless because any future window that reaches past this zero will contain a zero. So I throw the old list away and start a brand-new list that again contains only 1.  
If the number is not zero I take the most recent prefix, multiply it by the new number, and store the result. The new value is exactly the product of every number from the last zero up to the current position.

For getProduct I look at the length of the list. The length is always one larger than the number of real numbers stored after the last zero.  
If the requested k is greater than or equal to that length, the window of k numbers must include at least one zero, so the answer is zero.  
Otherwise the product of the last k numbers is simply the newest prefix divided by the prefix that sits k positions earlier. All the unwanted earlier factors cancel out, leaving only the product I need.

Because the problem states that every product fits inside a 32-bit integer, I never have to worry about overflow in any of the languages. The same reasoning works in C++, Java, JavaScript, TypeScript, Python and Go; only the way lists and integer division are written differs.

Edge cases are handled automatically:  
- Adding a zero resets everything.  
- Asking for a product larger than the current non-zero streak returns zero.  
- The first few adds after a zero correctly rebuild the prefix list from scratch.

## Examples

**Example 1**

Input sequence of operations:  
["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]  
[[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]

Step-by-step trace:  
- add(3) → list becomes [1, 3]  
- add(0) → list resets to [1]  
- add(2) → list becomes [1, 2]  
- add(5) → list becomes [1, 2, 10]  
- add(4) → list becomes [1, 2, 10, 40]  
- getProduct(2) → 40 / 10 = 20  
- getProduct(3) → 40 / 2 = 20  
- getProduct(4) → 4 numbers needed but only 3 exist after last zero → 0  
- add(8) → list becomes [1, 2, 10, 40, 320]  
- getProduct(2) → 320 / 40 = 8  

Output: [null, null, null, null, null, null, 20, 20, 0, null, 8]

**Example 2**

Operations: add(1), add(2), add(3), getProduct(3)  
List after three adds: [1, 1, 2, 6]  
getProduct(3) returns 6 / 1 = 6  
The product of 1 * 2 * 3 is correctly recovered by the single division.

**Example 3**

Operations: add(0), getProduct(1)  
After add(0) the list is [1].  
getProduct(1) sees that 1 is not less than the list length, so it returns 0, which is correct because the only number present is zero.

## How to Use / Run Locally

Clone the repository and navigate into the folder that contains the solution files.

**C++**  
g++ -std=c++17 product_of_last_k_numbers.cpp -o solution  
./solution

**Java**  
javac ProductOfNumbers.java  
java ProductOfNumbers

**JavaScript**  
node product_of_last_k_numbers.js

**TypeScript**  
tsc product_of_last_k_numbers.ts  
node product_of_last_k_numbers.js

**Python3**  
python3 product_of_last_k_numbers.py

**Go**  
go run product_of_last_k_numbers.go

Each file contains a simple driver that creates the object, calls a few add and getProduct operations, and prints the results so you can verify the behavior yourself.

## Notes & Optimizations

The key optimization is resetting the prefix list the moment a zero appears. Without that reset the division would eventually try to divide by zero or produce incorrect results.

An alternative approach is to store every number and multiply the last k elements on each query. That solution is correct but runs in O(k) time and will time out on the larger test cases.

Because the problem guarantees that the product always fits in a 32-bit signed integer, none of the languages need special big-integer handling. In languages where integer division truncates toward zero the results stay exact because every prefix is an integer multiple of earlier prefixes.

If the constraints ever allowed larger products, the same idea could be extended with modular arithmetic or big-integer libraries, but that is unnecessary for the current problem limits.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
