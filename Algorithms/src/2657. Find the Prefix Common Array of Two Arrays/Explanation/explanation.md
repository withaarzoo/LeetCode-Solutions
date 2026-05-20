# 2657. Find the Prefix Common Array of Two Arrays

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

---

## Problem Summary

In this LeetCode problem, we are given two integer arrays `A` and `B`. Both arrays are permutations of numbers from `1` to `n`.

For every index `i`, we need to find how many numbers are common in:

* `A[0...i]`
* `B[0...i]`

The result should be returned as a new array where each position stores the count of common elements for that prefix.

This problem is mainly about:

* prefix processing
* frequency counting
* array traversal
* hash/frequency based optimization

It is a good beginner-friendly DSA problem for understanding how prefix arrays and counting techniques work together.

---

## Constraints

| Constraint                             | Value                             |
| -------------------------------------- | --------------------------------- |
| `1 <= A.length == B.length == n <= 50` | Array size                        |
| `1 <= A[i], B[i] <= n`                 | Valid values                      |
| `A` and `B` are permutations           | Every number appears exactly once |

---

## Intuition

The first thing I noticed was that every number appears exactly once in both arrays.

That observation makes the problem much easier.

Instead of checking every prefix manually and comparing elements again and again, I realized I could track how many times each number had appeared while moving from left to right.

If a number appears twice overall, that means:

* one appearance came from `A`
* the other came from `B`

At that moment, the number becomes part of the prefix common count.

So the entire problem becomes a simple frequency tracking problem.

---

## Approach

I used a frequency array to store how many times each number has appeared so far.

Then I iterated through both arrays together.

At every index:

1. Add the current value from `A`
2. Add the current value from `B`
3. Check if any number's frequency becomes `2`
4. If yes, increase the common counter
5. Store the current common count in the answer array

Since both arrays are permutations, no number appears more than twice overall.

That makes the logic clean and efficient.

---

## Data Structures Used

| Data Structure  | Purpose                                        |
| --------------- | ---------------------------------------------- |
| Frequency Array | Tracks how many times each number has appeared |
| Result Array    | Stores prefix common counts                    |
| Integer Counter | Keeps running count of common numbers          |

---

## Operations & Behavior Summary

The algorithm works like this:

1. Create a frequency array of size `n + 1`
2. Start traversing both arrays together
3. Insert current elements into frequency tracking
4. Whenever frequency becomes `2`, increase common count
5. Save the common count into the answer array
6. Continue until all indices are processed

This avoids repeated comparisons and keeps the solution linear.

---

## Complexity

| Type             | Complexity | Explanation                   |
| ---------------- | ---------- | ----------------------------- |
| Time Complexity  | `O(n)`     | Each index is processed once  |
| Space Complexity | `O(n)`     | Extra frequency array is used |

Where:

* `n` = length of arrays `A` and `B`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> findThePrefixCommonArray(vector<int>& A, vector<int>& B) {
        
        int n = A.size();
        
        // Frequency array to track how many times
        // each number has appeared so far
        vector<int> freq(n + 1, 0);
        
        // Result array
        vector<int> ans(n);
        
        // Stores current count of common numbers
        int common = 0;
        
        for (int i = 0; i < n; i++) {
            
            // Add current element from A
            freq[A[i]]++;
            
            // If frequency becomes 2,
            // it means number appeared in both arrays
            if (freq[A[i]] == 2) {
                common++;
            }
            
            // Add current element from B
            freq[B[i]]++;
            
            // Same check for B
            if (freq[B[i]] == 2) {
                common++;
            }
            
            // Store answer for current prefix
            ans[i] = common;
        }
        
        return ans;
    }
};
```

### Java

```java
class Solution {
    public int[] findThePrefixCommonArray(int[] A, int[] B) {
        
        int n = A.length;
        
        // Frequency array to count appearances
        int[] freq = new int[n + 1];
        
        // Final answer array
        int[] ans = new int[n];
        
        // Current count of common elements
        int common = 0;
        
        for (int i = 0; i < n; i++) {
            
            // Process current element from A
            freq[A[i]]++;
            
            // If frequency becomes 2,
            // this number exists in both arrays
            if (freq[A[i]] == 2) {
                common++;
            }
            
            // Process current element from B
            freq[B[i]]++;
            
            // Same logic for B
            if (freq[B[i]] == 2) {
                common++;
            }
            
            // Store result for this prefix
            ans[i] = common;
        }
        
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var findThePrefixCommonArray = function(A, B) {
    
    let n = A.length;
    
    // Frequency array
    let freq = new Array(n + 1).fill(0);
    
    // Result array
    let ans = new Array(n);
    
    // Stores current common count
    let common = 0;
    
    for (let i = 0; i < n; i++) {
        
        // Add current value from A
        freq[A[i]]++;
        
        // If count becomes 2,
        // it is now common in both arrays
        if (freq[A[i]] === 2) {
            common++;
        }
        
        // Add current value from B
        freq[B[i]]++;
        
        // Same check for B
        if (freq[B[i]] === 2) {
            common++;
        }
        
        // Save current answer
        ans[i] = common;
    }
    
    return ans;
};
```

### Python3

```python
class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        
        n = len(A)
        
        # Frequency array to count appearances
        freq = [0] * (n + 1)
        
        # Final answer array
        ans = [0] * n
        
        # Stores count of common elements
        common = 0
        
        for i in range(n):
            
            # Add current element from A
            freq[A[i]] += 1
            
            # If frequency becomes 2,
            # number exists in both arrays
            if freq[A[i]] == 2:
                common += 1
            
            # Add current element from B
            freq[B[i]] += 1
            
            # Same logic for B
            if freq[B[i]] == 2:
                common += 1
            
            # Store answer for this prefix
            ans[i] = common
        
        return ans
```

### Go

```go
func findThePrefixCommonArray(A []int, B []int) []int {
    
    n := len(A)
    
    // Frequency array to track appearances
    freq := make([]int, n+1)
    
    // Result array
    ans := make([]int, n)
    
    // Stores count of common elements
    common := 0
    
    for i := 0; i < n; i++ {
        
        // Add current element from A
        freq[A[i]]++
        
        // If frequency becomes 2,
        // number appeared in both arrays
        if freq[A[i]] == 2 {
            common++
        }
        
        // Add current element from B
        freq[B[i]]++
        
        // Same check for B
        if freq[B[i]] == 2 {
            common++
        }
        
        // Store current prefix answer
        ans[i] = common
    }
    
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

First, a frequency array is created.

The purpose of this array is simple:

* index = number
* value = how many times that number has appeared so far

Then a variable called `common` is maintained.

This variable stores how many numbers are currently common between the prefixes of both arrays.

Now the algorithm starts iterating from index `0` to `n - 1`.

For every iteration:

* the current element from `A` is processed
* the current element from `B` is processed

Whenever a number's frequency becomes exactly `2`, it means:

* the number has now appeared once in `A`
* and once in `B`

So that number becomes part of the prefix common array.

The algorithm increases the `common` count immediately.

Then the current value of `common` is stored into the answer array.

The reason this works so well is because both arrays are permutations.

That means:

* no duplicates inside one array
* every number appears exactly twice overall at most

So frequency `2` directly means the number is common.

If this were not a permutation problem, the logic would need additional checks.

---

## Examples

### Example 1

Input:

```text
A = [1,3,2,4]
B = [3,1,2,4]
```

Output:

```text
[0,2,3,4]
```

Explanation:

* At index `0`, no common number exists
* At index `1`, numbers `1` and `3` become common
* At index `2`, number `2` also becomes common
* At index `3`, all numbers become common

---

### Example 2

Input:

```text
A = [2,3,1]
B = [3,1,2]
```

Output:

```text
[0,1,3]
```

Explanation:

* At index `0`, no common numbers
* At index `1`, only `3` is common
* At index `2`, all numbers are common

---

### Example 3

Input:

```text
A = [1,2]
B = [1,2]
```

Output:

```text
[1,2]
```

Explanation:

* At index `0`, number `1` is already common
* At index `1`, both numbers become common

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* This is already an optimal solution for the given constraints.
* Brute force comparison of prefixes would take much more time.
* Using a frequency array keeps the solution clean and fast.
* Since values range from `1` to `n`, array indexing works perfectly.
* A hash map could also work, but an array is faster here.
* The permutation property is the key observation behind the optimization.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
