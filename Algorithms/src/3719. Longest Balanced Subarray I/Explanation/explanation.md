# 3719. Longest Balanced Subarray I

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

I am given an integer array `nums`.

A subarray is called **balanced** if:

* The number of **distinct even numbers**
* is equal to
* The number of **distinct odd numbers**

My task is to **return the length of the longest balanced subarray**.

Important note:
Only **distinct values matter**, not how many times they appear.

---

## 📏 Constraints

* `1 ≤ nums.length ≤ 1500`
* `1 ≤ nums[i] ≤ 10⁵`

Because the array size is limited, a **quadratic solution is acceptable**.

---

## 💡 Intuition

When I read the problem, the word **“distinct”** immediately caught my attention.

I realized:

* Duplicate numbers do not increase the count.
* I only need to track **unique evens** and **unique odds**.

So I thought:

> If I fix a starting index and keep expanding the subarray,
> I can keep checking how many distinct evens and odds I have.

Whenever both counts match, that subarray is balanced.

---

## 🚀 Approach

Here is exactly how I solved it:

1. I fix a starting index `i`
2. I create two sets:

   * one for distinct even numbers
   * one for distinct odd numbers
3. I expand the subarray from `i` to `j`
4. For every new number:

   * if it is even → add to even set
   * if it is odd → add to odd set
5. If both sets have the same size:

   * update the maximum length
6. Repeat this for all possible starting positions

This guarantees I check **all valid subarrays**.

---

## 🧰 Data Structures Used

* **Hash Set**

  * Stores distinct even numbers
  * Stores distinct odd numbers
  * Ensures O(1) average insert and lookup

---

## 🔄 Operations & Behavior Summary

* Loop through all possible subarrays
* Maintain uniqueness using sets
* Compare distinct even vs odd counts
* Track the maximum valid subarray length

---

## ⏱️ Complexity

**Time Complexity:** `O(n²)`

* Two nested loops over the array
* `n` = length of the array

**Space Complexity:** `O(n)`

* Hash sets store distinct numbers within a subarray

---

## 💻 Multi-language Solutions

---

### ### C++

```cpp
class Solution {
public:
    int longestBalanced(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;

        for (int i = 0; i < n; i++) {
            unordered_set<int> evenSet, oddSet;

            for (int j = i; j < n; j++) {
                if (nums[j] % 2 == 0)
                    evenSet.insert(nums[j]);
                else
                    oddSet.insert(nums[j]);

                if (evenSet.size() == oddSet.size())
                    ans = max(ans, j - i + 1);
            }
        }
        return ans;
    }
};
```

---

### ### Java

```java
class Solution {
    public int longestBalanced(int[] nums) {
        int n = nums.length;
        int ans = 0;

        for (int i = 0; i < n; i++) {
            HashSet<Integer> evenSet = new HashSet<>();
            HashSet<Integer> oddSet = new HashSet<>();

            for (int j = i; j < n; j++) {
                if (nums[j] % 2 == 0)
                    evenSet.add(nums[j]);
                else
                    oddSet.add(nums[j]);

                if (evenSet.size() == oddSet.size())
                    ans = Math.max(ans, j - i + 1);
            }
        }
        return ans;
    }
}
```

---

### ### JavaScript

```javascript
var longestBalanced = function(nums) {
    let ans = 0;

    for (let i = 0; i < nums.length; i++) {
        const evenSet = new Set();
        const oddSet = new Set();

        for (let j = i; j < nums.length; j++) {
            if (nums[j] % 2 === 0)
                evenSet.add(nums[j]);
            else
                oddSet.add(nums[j]);

            if (evenSet.size === oddSet.size)
                ans = Math.max(ans, j - i + 1);
        }
    }
    return ans;
};
```

---

### ### Python3

```python
class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        ans = 0
        n = len(nums)

        for i in range(n):
            even_set = set()
            odd_set = set()

            for j in range(i, n):
                if nums[j] % 2 == 0:
                    even_set.add(nums[j])
                else:
                    odd_set.add(nums[j])

                if len(even_set) == len(odd_set):
                    ans = max(ans, j - i + 1)

        return ans
```

---

### ### Go

```go
func longestBalanced(nums []int) int {
    ans := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        evenSet := make(map[int]bool)
        oddSet := make(map[int]bool)

        for j := i; j < n; j++ {
            if nums[j]%2 == 0 {
                evenSet[nums[j]] = true
            } else {
                oddSet[nums[j]] = true
            }

            if len(evenSet) == len(oddSet) {
                if j-i+1 > ans {
                    ans = j - i + 1
                }
            }
        }
    }
    return ans
}
```

---

## 🪜 Step-by-step Detailed Explanation

1. I start from every index `i`
2. I clear my even and odd sets
3. I extend the subarray to the right
4. I insert numbers into the correct set
5. I compare the sizes of both sets
6. If equal, I update the answer
7. I repeat this for all starting points

This ensures **every possible balanced subarray is checked**.

---

## 🧪 Examples

**Input:** `[2,5,4,3]`
**Output:** `4`
**Explanation:**
Distinct evens → `{2,4}`
Distinct odds → `{5,3}`

Balanced ✔

---

## ▶️ How to use / Run locally

1. Copy the solution in your preferred language
2. Paste it into:

   * LeetCode editor, or
   * Local compiler (VS Code, IntelliJ, etc.)
3. Run with test cases

---

## 📝 Notes & Optimizations

* This solution is optimal for given constraints
* Avoids unnecessary frequency counting
* Clean and interview-friendly
* Uses simple logic and readable code

---

## 👤 Author

* **Md Aarzoo Islam**
  🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
