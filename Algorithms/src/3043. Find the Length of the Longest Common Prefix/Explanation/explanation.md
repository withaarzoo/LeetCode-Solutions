# 3043. Find the Length of the Longest Common Prefix

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

In this LeetCode problem, we are given two arrays of positive integers: `arr1` and `arr2`.

The task is to find the length of the longest common prefix between any pair of numbers where:

* one number comes from `arr1`
* the other number comes from `arr2`

A prefix means digits starting from the left side of a number.

Example:

* `123` is a prefix of `12345`
* `56` is not a prefix of `12345`

The final answer is not the prefix itself. We only need the length of the longest matching prefix.

This is a classic hashing and prefix-processing problem that appears frequently in competitive programming and DSA interview preparation.

---

## Constraints

| Constraint                                  | Value               |
| ------------------------------------------- | ------------------- |
| `1 <= arr1.length, arr2.length <= 5 * 10^4` | Array size limit    |
| `1 <= arr1[i], arr2[i] <= 10^8`             | Integer value limit |

---

## Intuition

When I first looked at this problem, my instinct was to compare every pair of numbers and check how many starting digits matched.

But that would be very slow because both arrays can contain up to `50,000` elements.

Then I noticed something important:

Every number automatically contains all of its prefixes.

For example:

`12345`

* `12345`
* `1234`
* `123`
* `12`
* `1`

So instead of comparing every pair directly, I can generate all prefixes from the first array and store them in a hash set.

Then for every number in the second array, I keep removing digits from the end until I find a prefix that exists in the set.

This makes the solution much faster and cleaner.

---

## Approach

First, I generate every possible prefix from all numbers inside `arr1`.

I store those prefixes inside a hash set because lookup operations are very fast.

Then I iterate through every number in `arr2`.

For each number:

1. Check whether the current number exists in the prefix set
2. If not, remove the last digit
3. Continue until:

   * a prefix is found
   * or the number becomes `0`

Whenever I find a valid prefix, I calculate its digit length and update the maximum answer.

Since I check from largest prefix to smallest prefix, the first match is automatically the longest one for that number.

---

## Data Structures Used

| Data Structure    | Purpose                                                   |
| ----------------- | --------------------------------------------------------- |
| Hash Set          | Stores all prefixes generated from `arr1` for fast lookup |
| Integer Variables | Used for prefix generation and digit removal              |

The hash set is the key optimization in this solution because it avoids repeated pairwise comparisons.

---

## Operations & Behavior Summary

The algorithm works in two major stages.

### Stage 1: Build Prefix Storage

For every number in `arr1`:

* Add the full number into the set
* Remove the last digit
* Add the new prefix
* Repeat until the number becomes `0`

### Stage 2: Search Common Prefixes

For every number in `arr2`:

* Check whether the current value exists in the set
* If yes:

  * compute digit length
  * update answer
  * stop checking smaller prefixes
* Otherwise:

  * remove the last digit
  * continue searching

At the end, return the maximum prefix length found.

---

## Complexity

| Type             | Complexity       | Explanation                                              |
| ---------------- | ---------------- | -------------------------------------------------------- |
| Time Complexity  | `O((n + m) * k)` | `n` and `m` are array sizes, `k` is maximum digit length |
| Space Complexity | `O(n * k)`       | Extra space is used for storing prefixes in a hash set   |

Since integer values are at most `10^8`, the maximum number of digits is small, which keeps the solution efficient.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int longestCommonPrefix(vector<int>& arr1, vector<int>& arr2) {
        
        // Hash set to store all prefixes from arr1
        unordered_set<int> prefixes;

        // Generate all prefixes from arr1
        for (int num : arr1) {

            int x = num;

            // Keep removing last digit
            while (x > 0) {

                // Store current prefix
                prefixes.insert(x);

                // Remove last digit
                x /= 10;
            }
        }

        int ans = 0;

        // Check every number from arr2
        for (int num : arr2) {

            int x = num;

            // Keep shortening the number
            while (x > 0) {

                // If prefix exists in arr1 prefixes
                if (prefixes.count(x)) {

                    // Convert to string to get digit length
                    ans = max(ans, (int)to_string(x).size());

                    // No need to check smaller prefixes
                    break;
                }

                // Remove last digit
                x /= 10;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        
        // Hash set to store all prefixes from arr1
        HashSet<Integer> prefixes = new HashSet<>();

        // Generate all prefixes
        for (int num : arr1) {

            int x = num;

            // Keep removing last digit
            while (x > 0) {

                // Store current prefix
                prefixes.add(x);

                // Remove last digit
                x /= 10;
            }
        }

        int ans = 0;

        // Process arr2
        for (int num : arr2) {

            int x = num;

            // Keep checking prefixes
            while (x > 0) {

                // Prefix found
                if (prefixes.contains(x)) {

                    // Update maximum length
                    ans = Math.max(ans, String.valueOf(x).length());

                    // Stop because larger prefix already found
                    break;
                }

                // Remove last digit
                x /= 10;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @return {number}
 */
var longestCommonPrefix = function(arr1, arr2) {
    
    // Set to store all prefixes from arr1
    const prefixes = new Set();

    // Generate prefixes
    for (let num of arr1) {

        let x = num;

        // Keep removing last digit
        while (x > 0) {

            // Store current prefix
            prefixes.add(x);

            // Remove last digit
            x = Math.floor(x / 10);
        }
    }

    let ans = 0;

    // Check numbers from arr2
    for (let num of arr2) {

        let x = num;

        // Try all prefixes
        while (x > 0) {

            // Prefix found
            if (prefixes.has(x)) {

                // Update answer using digit count
                ans = Math.max(ans, x.toString().length);

                // Stop because this is the longest for current number
                break;
            }

            // Remove last digit
            x = Math.floor(x / 10);
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        
        # Set to store all prefixes from arr1
        prefixes = set()

        # Generate prefixes from arr1
        for num in arr1:

            x = num

            # Keep removing last digit
            while x > 0:

                # Store current prefix
                prefixes.add(x)

                # Remove last digit
                x //= 10

        ans = 0

        # Process arr2
        for num in arr2:

            x = num

            # Try all prefixes
            while x > 0:

                # Prefix exists
                if x in prefixes:

                    # Update answer with digit length
                    ans = max(ans, len(str(x)))

                    # Stop because longer prefix already found
                    break

                # Remove last digit
                x //= 10

        return ans
```

### Go

```go
func longestCommonPrefix(arr1 []int, arr2 []int) int {
    
    // Map used as a hash set to store prefixes
    prefixes := make(map[int]bool)

    // Generate all prefixes from arr1
    for _, num := range arr1 {

        x := num

        // Keep removing last digit
        for x > 0 {

            // Store current prefix
            prefixes[x] = true

            // Remove last digit
            x /= 10
        }
    }

    ans := 0

    // Process arr2
    for _, num := range arr2 {

        x := num

        // Keep checking prefixes
        for x > 0 {

            // Prefix found
            if prefixes[x] {

                // Count digits manually
                length := 0
                temp := x

                for temp > 0 {
                    length++
                    temp /= 10
                }

                // Update answer
                if length > ans {
                    ans = length
                }

                // Stop because larger prefix already found
                break
            }

            // Remove last digit
            x /= 10
        }
    }

    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

Only syntax changes.

### Step 1: Create a Prefix Storage

The first thing I do is create a hash-based structure.

Depending on the language:

* C++ uses `unordered_set`
* Java uses `HashSet`
* JavaScript uses `Set`
* Python uses `set`
* Go uses `map[int]bool`

This structure stores every prefix from `arr1`.

Fast lookup is important because we need to check prefixes repeatedly.

---

### Step 2: Generate All Prefixes

For every number:

* store the current number
* divide by `10`
* repeat

Example:

Starting with `12345`

After repeatedly removing digits:

* `12345`
* `1234`
* `123`
* `12`
* `1`

These are all valid prefixes.

This preprocessing step is what makes the algorithm efficient.

Without this step, we would need nested comparisons between both arrays.

---

### Step 3: Search Prefixes for arr2

Now I process every number from `arr2`.

I again remove digits one by one.

But this time, I check whether the current value exists inside the prefix set.

Example:

For `12399`

* check `12399`
* check `1239`
* check `123`
* found

So the common prefix length is `3`.

---

### Step 4: Update Maximum Answer

Whenever a match is found:

* calculate digit length
* compare with current answer
* update if larger

Then immediately stop checking smaller prefixes.

Why?

Because we are already checking from longest to shortest.

The first valid match is automatically the best match for that number.

---

### Step 5: Return Final Result

After processing every number:

* the stored maximum value becomes the answer

If no prefix is found at all:

* return `0`

---

## Examples

### Example 1

#### Input

```text
arr1 = [1,10,100]
arr2 = [1000]
```

#### Output

```text
3
```

#### Explanation

Possible matches:

* `1` matches with `1000`
* `10` matches with `1000`
* `100` matches with `1000`

The longest common prefix is `100`.

Length = `3`

---

### Example 2

#### Input

```text
arr1 = [1,2,3]
arr2 = [4,4,4]
```

#### Output

```text
0
```

#### Explanation

No numbers share a starting prefix.

So the answer is `0`.

---

### Example 3

#### Input

```text
arr1 = [5655359]
arr2 = [56554]
```

#### Output

```text
3
```

#### Explanation

Common prefixes:

* `5`
* `56`
* `565`

Longest common prefix length = `3`

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
javac Solution.java
```

Run:

```bash
java Solution
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

or

```bash
python3 main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* A brute-force solution would compare every pair of numbers directly.
* That approach can become extremely slow for large inputs.
* Using a hash set improves lookup speed significantly.
* Prefix generation using integer division is faster than converting everything into strings repeatedly.
* The algorithm works efficiently because numbers contain at most a few digits.
* This is a good example of combining hashing with prefix processing.

Possible alternative approaches:

* Trie (Prefix Tree)
* String-based prefix matching
* Sorting and adjacent comparison techniques

The hash set solution is simpler and performs very well for the given constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
