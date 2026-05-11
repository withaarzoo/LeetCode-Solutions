# 2553. Separate the Digits in an Array

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

In this LeetCode problem, we are given an array of positive integers.
The task is to separate every digit from each number and return all digits inside a single array while keeping the original order unchanged.

For example:

```text
Input: [13,25,83,77]
Output: [1,3,2,5,8,3,7,7]
```

The problem is mainly about array processing, digit extraction, and maintaining order correctly.

This is a beginner-friendly DSA problem that helps improve understanding of loops, strings, arrays, and basic number handling.

---

## Constraints

| Constraint                 | Value                 |
| -------------------------- | --------------------- |
| `1 <= nums.length <= 1000` | Array size            |
| `1 <= nums[i] <= 10^5`     | Value of each integer |

---

## Intuition

The first thing I noticed was that every number simply needs to be broken into smaller pieces called digits.

There is no complicated algorithm here.
The main challenge is preserving the exact order of digits.

I realized that converting each number into a string would make the solution much simpler. Once a number becomes a string, every character directly represents a digit.

Then I can just loop through the characters and store them into the final result array.

This approach is clean, readable, and avoids unnecessary mathematical operations.

---

## Approach

First, I create an empty result array.

Then I iterate through every number in the input array.

For each number:

1. Convert the number into a string
2. Traverse every character of that string
3. Convert the character back into an integer digit
4. Push the digit into the result array

After processing all numbers, return the final array.

This solution naturally keeps digits in the same order as the original numbers.

---

## Data Structures Used

| Data Structure        | Why It Was Used                            |
| --------------------- | ------------------------------------------ |
| Array / Vector / List | Stores the final separated digits          |
| String                | Makes digit extraction simple and readable |

---

## Operations & Behavior Summary

The algorithm works in these stages:

1. Start with an empty answer array
2. Read each number from the input
3. Convert the number into string format
4. Read every digit character one by one
5. Convert character digits into integer digits
6. Store them into the result array
7. Return the completed array

This process continues until all numbers are fully processed.

---

## Complexity

| Type             | Complexity |
| ---------------- | ---------- |
| Time Complexity  | `O(n * k)` |
| Space Complexity | `O(n * k)` |

### Explanation

* `n` = number of integers in the array
* `k` = average number of digits per integer

The algorithm visits every digit exactly once.

Extra space is used for storing the final separated digits array.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> separateDigits(vector<int>& nums) {
        
        // Final array that will store all separated digits
        vector<int> result;

        // Traverse every number in the input array
        for (int num : nums) {

            // Convert number into string so digits become easy to access
            string s = to_string(num);

            // Traverse every character in the string
            for (char ch : s) {

                // Convert character digit into integer and store it
                result.push_back(ch - '0');
            }
        }

        // Return final separated digits array
        return result;
    }
};
```

### Java

```java
class Solution {
    public int[] separateDigits(int[] nums) {
        
        // Dynamic list to store digits temporarily
        List<Integer> list = new ArrayList<>();

        // Traverse every number
        for (int num : nums) {

            // Convert number to string
            String s = String.valueOf(num);

            // Traverse each character of the string
            for (char ch : s.toCharArray()) {

                // Convert character into integer digit
                list.add(ch - '0');
            }
        }

        // Convert List<Integer> into int[]
        int[] result = new int[list.size()];

        for (int i = 0; i < list.size(); i++) {
            result[i] = list.get(i);
        }

        // Return final answer
        return result;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number[]}
 */
var separateDigits = function(nums) {
    
    // Final array to store separated digits
    let result = [];

    // Traverse every number in nums
    for (let num of nums) {

        // Convert number into string
        let str = num.toString();

        // Traverse every character
        for (let ch of str) {

            // Convert character into number and store it
            result.push(Number(ch));
        }
    }

    // Return final array
    return result;
};
```

### Python3

```python
class Solution:
    def separateDigits(self, nums: List[int]) -> List[int]:
        
        # Final array to store separated digits
        result = []

        # Traverse every number
        for num in nums:

            # Convert number into string
            s = str(num)

            # Traverse every character in string
            for ch in s:

                # Convert character back to integer and store it
                result.append(int(ch))

        # Return final answer
        return result
```

### Go

```go
func separateDigits(nums []int) []int {
    
    // Final array to store separated digits
    result := []int{}

    // Traverse every number
    for _, num := range nums {

        // Convert number into string
        str := strconv.Itoa(num)

        // Traverse every character in string
        for _, ch := range str {

            // Convert character digit into integer
            result = append(result, int(ch-'0'))
        }
    }

    // Return final answer
    return result
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays almost identical in every programming language.

First, an empty result container is created.
This container stores all separated digits.

Then the algorithm loops through every number in the input array.

Each number is converted into a string because strings allow direct access to individual digits.

For example:

```text
123 -> "123"
```

Now every character can be accessed separately:

```text
'1', '2', '3'
```

Each character digit is converted back into an integer and added to the final answer array.

The process repeats until every number has been processed.

### Why string conversion works well

There are two common ways to solve this problem:

1. Mathematical digit extraction
2. String conversion

I prefer the string approach because it keeps digits in the correct left-to-right order automatically.

If mathematical operations like `% 10` were used, digits would come out in reverse order:

```text
123 -> 3,2,1
```

Then extra reversing logic would be needed.

Using strings makes the implementation shorter, cleaner, and easier for beginners to understand.

### Language behavior differences

#### C++

Uses `to_string()` for conversion and `vector<int>` for storing digits.

#### Java

Uses `String.valueOf()` and an `ArrayList<Integer>` before converting to `int[]`.

#### JavaScript

Uses `toString()` and dynamically pushes digits into an array.

#### Python3

Uses `str()` conversion and appends digits into a list.

#### Go

Uses `strconv.Itoa()` because Go does not automatically convert integers into strings.

---

## Examples

### Example 1

```text
Input: nums = [13,25,83,77]

Output: [1,3,2,5,8,3,7,7]
```

### Trace

```text
13 -> [1,3]
25 -> [2,5]
83 -> [8,3]
77 -> [7,7]
```

Final array:

```text
[1,3,2,5,8,3,7,7]
```

---

### Example 2

```text
Input: nums = [7,1,3,9]

Output: [7,1,3,9]
```

### Trace

Each number already contains only one digit.

```text
7 -> [7]
1 -> [1]
3 -> [3]
9 -> [9]
```

Final array remains the same.

---

### Example 3

```text
Input: nums = [100,56]

Output: [1,0,0,5,6]
```

### Trace

```text
100 -> [1,0,0]
56 -> [5,6]
```

Final array:

```text
[1,0,0,5,6]
```

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

* This problem does not require advanced algorithms.
* String conversion is the cleanest solution for readability.
* Mathematical extraction can also solve the problem, but it usually needs reversing logic.
* Since constraints are small, performance is already efficient enough.
* Order preservation is the most important detail in this problem.
* Edge cases like single-digit numbers work automatically without extra handling.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
