# 📌 Problem Title

**744. Find Smallest Letter Greater Than Target**

---

## 📑 Table of Contents

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

## 🧩 Problem Summary

I am given a sorted array of lowercase English letters and a target character.

My task is to find the **smallest character in the array that is strictly greater than the target**.

If such a character does not exist, then I must return the **first character of the array** (wrap-around rule).

The array is already sorted in non-decreasing order.

---

## ⛓ Constraints

* `2 <= letters.length <= 10^4`
* `letters[i]` is a lowercase English letter
* `letters` is sorted in non-decreasing order
* `letters` contains at least two different characters
* `target` is a lowercase English letter

---

## 💡 Intuition

When I read the problem, the first thing I noticed was that the array is **already sorted**.

That immediately made me think:

> “I don’t need to check every element. I can use Binary Search.”

I also noticed the **wrap-around condition**:

* If no letter is greater than the target, I must return `letters[0]`.

So my thinking became simple:

* Find the **first letter greater than target**
* If not found, return the **first letter**

Binary Search fits perfectly here.

---

## 🚀 Approach

1. I use **Binary Search** on the sorted array.
2. I keep two pointers: `left` and `right`.
3. I check the middle element.
4. If the middle character is **greater than target**:

   * I store it as a possible answer.
   * I move left to find a smaller valid character.
5. If the middle character is **less than or equal to target**:

   * I move right.
6. After the loop:

   * If I found an answer, I return it.
   * Otherwise, I return `letters[0]` (wrap-around).

---

## 🧱 Data Structures Used

* Array (input)
* A few variables for Binary Search

No extra data structures are used.

---

## 🔄 Operations & Behavior Summary

* Binary search reduces search space every step.
* Always keeps track of the smallest valid character.
* Handles wrap-around naturally by defaulting to `letters[0]`.
* Efficient and clean logic.

---

## ⏱ Complexity

**Time Complexity:** `O(log n)`

* `n` is the number of characters in the array.
* Binary Search divides the array in half each step.

**Space Complexity:** `O(1)`

* No extra space is used.

---

## 🌐 Multi-language Solutions

### C++

```cpp
class Solution {
public:
    char nextGreatestLetter(vector<char>& letters, char target) {
        int left = 0, right = letters.size() - 1;
        char answer = letters[0];

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (letters[mid] > target) {
                answer = letters[mid];
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        return answer;
    }
};
```

---

### Java

```java
class Solution {
    public char nextGreatestLetter(char[] letters, char target) {
        int left = 0, right = letters.length - 1;
        char answer = letters[0];

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (letters[mid] > target) {
                answer = letters[mid];
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        return answer;
    }
}
```

---

### JavaScript

```javascript
var nextGreatestLetter = function(letters, target) {
    let left = 0, right = letters.length - 1;
    let answer = letters[0];

    while (left <= right) {
        let mid = Math.floor(left + (right - left) / 2);

        if (letters[mid] > target) {
            answer = letters[mid];
            right = mid - 1;
        } else {
            left = mid + 1;
        }
    }
    return answer;
};
```

---

### Python3

```python
class Solution:
    def nextGreatestLetter(self, letters, target):
        left, right = 0, len(letters) - 1
        answer = letters[0]

        while left <= right:
            mid = (left + right) // 2

            if letters[mid] > target:
                answer = letters[mid]
                right = mid - 1
            else:
                left = mid + 1

        return answer
```

---

### Go

```go
func nextGreatestLetter(letters []byte, target byte) byte {
    left, right := 0, len(letters)-1
    answer := letters[0]

    for left <= right {
        mid := left + (right-left)/2

        if letters[mid] > target {
            answer = letters[mid]
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return answer
}
```

---

## 🧠 Step-by-step Detailed Explanation (All Languages)

1. Start with two pointers at the beginning and end of the array.
2. Find the middle element.
3. Compare the middle character with the target.
4. If it is greater:

   * Save it as a possible answer.
   * Search the left half.
5. If it is smaller or equal:

   * Search the right half.
6. Continue until the search space is exhausted.
7. Return the stored answer or the first character if nothing was found.

---

## 🧪 Examples

**Input:**
`letters = ["c","f","j"], target = "a"`
**Output:**
`"c"`

**Input:**
`letters = ["c","f","j"], target = "c"`
**Output:**
`"f"`

**Input:**
`letters = ["x","x","y","y"], target = "z"`
**Output:**
`"x"`

---

## ▶️ How to use / Run locally

1. Copy the solution code in your preferred language.
2. Paste it into a local compiler or editor.
3. Provide test inputs manually or through test cases.
4. Run and verify the output.

---

## 📝 Notes & Optimizations

* Binary Search is the most optimal solution.
* Linear search would also work but is slower.
* Wrap-around is handled naturally by initializing the answer.
* This solution is **interview-safe and production-ready**.

---

## 👤 Author

* **[Md Aarzoo Islam](https://bento.me/withaarzoo)**
