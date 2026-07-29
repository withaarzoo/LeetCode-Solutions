# K-th Smallest Palindrome: Competitive Programming DSA Solution

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

## Problem Summary

This problem is about finding the **k-th smallest palindrome** that can be formed from a given string.

The key idea is simple: a palindrome is symmetric, so the **first half decides the whole string**. Instead of generating every possible arrangement of the full string, we only work on the left half, count how many valid palindromes can be formed, and use that count to jump directly to the k-th one in lexicographic order.

The input is a string `s` and an integer `k`.

The output is the **k-th lexicographically smallest palindrome** that can be built from the characters of `s`. If fewer than `k` valid palindromes exist, the answer is an empty string.

## Constraints

The exact original problem limits are not included in the attached source, so the implementation below follows these practical assumptions:

| Constraint / Assumption               | Meaning                                               |
| ------------------------------------- | ----------------------------------------------------- |
| Lowercase English letters only        | The solution uses fixed arrays of size 26             |
| `k` is 1-indexed                      | The first valid palindrome is the 1st answer          |
| The string can form a palindrome      | At most one character may have an odd frequency       |
| Large permutation counts are possible | Counting is capped at `k + 1` to avoid overflow       |
| The algorithm must stay efficient     | Brute force generation of all palindromes is not used |

## Intuition

The main observation is that a palindrome is mirror-symmetric.

That means if I know the first half, I already know the second half. The middle character, if it exists, is fixed and does not affect lexicographic choice in the same way.

So the real problem becomes:

1. Build only the first half.
2. Decide each character from left to right.
3. Count how many valid palindromes are still possible after choosing a character.
4. Use that count to skip the branches that do not contain the k-th answer.

This is very similar to finding the k-th permutation, but with repeated characters and palindrome rules.

## Approach

1. Count the frequency of every character in the string.
2. Split the frequency into:

   * the left half counts
   * the middle character, if one character has an odd count
3. Compute how many unique permutations are possible using the left half.
4. If the total number of valid palindromes is smaller than `k`, return an empty string.
5. Build the first half one character at a time:

   * try letters from `'a'` to `'z'`
   * temporarily use one occurrence of that letter
   * count how many palindromes can still be formed
   * if the count is at least `k`, keep the letter
   * otherwise, skip it and reduce `k` by that count
6. Finally, form the complete palindrome as:

   * `first_half + middle + reverse(first_half)`

This is a clean greedy strategy with combinatorics.

## Data Structures Used

* **Frequency array of size 26**
  Used to count lowercase English letters quickly.

* **Half-frequency array**
  Stores only the characters needed for the first half of the palindrome.

* **String / StringBuilder / vector of chars / byte slice**
  Used to build the answer efficiently in each language.

* **Helper counting function**
  Calculates how many unique permutations are possible from the remaining multiset of characters.

## Operations & Behavior Summary

1. Read the string and count all characters.
2. Move half of each frequency into the left-half pool.
3. Save any odd character as the middle character.
4. Count total possible permutations for the left half.
5. If there are not enough permutations, stop early.
6. For each position in the left half:

   * test letters in alphabetical order
   * count how many valid completions each choice gives
   * keep the first letter whose branch contains the k-th answer
7. Mirror the left half to finish the palindrome.

In plain English, the algorithm keeps asking:
“Can I still reach the k-th palindrome if I choose this character right now?”

If yes, it locks that choice. If not, it moves ahead.

## Complexity

| Type             | Complexity   | Explanation                                                                                                                                               |
| ---------------- | ------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n * 26)`  | We build roughly half the string, and at each step we try up to 26 letters. The counting helper is capped early at `k + 1`, so it stays fast in practice. |
| Space Complexity | `O(1)` extra | The solution uses fixed-size arrays of length 26, plus the output string.                                                                                 |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string smallestPalindrome(string s, int k) {
        // Store frequencies of each character in the string
        vector<int> freq(26, 0);
        for (char c : s) {
            freq[c - 'a']++;
        }
        
        // Isolate frequencies meant only for the first half of the palindrome
        vector<int> half(26, 0);
        string mid = "";
        int m = 0;
        
        for (int i = 0; i < 26; ++i) {
            // The character with an odd frequency goes in the exact middle
            if (freq[i] % 2 != 0) {
                mid += (char)(i + 'a');
            }
            half[i] = freq[i] / 2;
            m += half[i];
        }
        
        // Helper lambda to calculate permutations of remaining characters
        auto get_ways = [&](const vector<int>& f, long long target_k) {
            long long ways = 1;
            int curr_len = 0;
            for (int count : f) {
                if (count > 0) {
                    curr_len += count;
                    long long n = curr_len;
                    long long r = count;
                    
                    // Optimize nCr calculation by choosing the smaller r
                    if (r > n - r) r = n - r;
                    long long cur_nCr = 1;
                    
                    // Calculate nCr iteratively and cap it at target_k to prevent overflow
                    for (int i = 1; i <= r; ++i) {
                        cur_nCr = cur_nCr * (n - i + 1) / i;
                        if (cur_nCr > target_k) {
                            cur_nCr = target_k + 1;
                            break;
                        }
                    }
                    ways *= cur_nCr;
                    if (ways > target_k) return target_k + 1;
                }
            }
            return ways;
        };
        
        // If the total valid permutations are less than k, return empty string
        if (get_ways(half, k) < k) {
            return "";
        }
        
        string first_half = "";
        // Build the first half character by character
        for (int i = 0; i < m; ++i) {
            for (int c = 0; c < 26; ++c) {
                if (half[c] > 0) {
                    // Try using the current character
                    half[c]--;
                    long long ways = get_ways(half, k);
                    
                    // If permutations are enough, lock in this character
                    if (ways >= k) {
                        first_half += (char)(c + 'a');
                        break; 
                    } else {
                        // Otherwise, shrink k and restore character to try the next one
                        k -= ways;
                        half[c]++;
                    }
                }
            }
        }
        
        // Assemble final palindrome: first half + mid + reversed first half
        string res = first_half + mid;
        for (int i = m - 1; i >= 0; --i) {
            res += first_half[i];
        }
        return res;
    }
};
```

### Java

```java
class Solution {
    public String smallestPalindrome(String s, int k) {
        // Track global character frequencies
        int[] freq = new int[26];
        for (char c : s.toCharArray()) {
            freq[c - 'a']++;
        }
        
        // Track half frequencies and identify the middle character
        int[] half = new int[26];
        StringBuilder mid = new StringBuilder();
        int m = 0;
        
        for (int i = 0; i < 26; ++i) {
            if (freq[i] % 2 != 0) {
                mid.append((char) (i + 'a'));
            }
            half[i] = freq[i] / 2;
            m += half[i];
        }
        
        // Initial permutation check to see if target k is reachable
        if (getWays(half, k) < k) {
            return "";
        }
        
        StringBuilder firstHalf = new StringBuilder();
        // Construct left half from left to right
        for (int i = 0; i < m; ++i) {
            for (int c = 0; c < 26; ++c) {
                if (half[c] > 0) {
                    // Temporarily claim the character
                    half[c]--;
                    long ways = getWays(half, k);
                    
                    // Target string is within the branches of this character choice
                    if (ways >= k) {
                        firstHalf.append((char) (c + 'a'));
                        break;
                    } else {
                        // Skip character and decrease our k target
                        k -= ways;
                        half[c]++;
                    }
                }
            }
        }
        
        // Mirror the first half to complete the palindrome
        StringBuilder res = new StringBuilder(firstHalf);
        res.append(mid);
        res.append(firstHalf.reverse());
        return res.toString();
    }
    
    // Helper to calculate total permutations of the remaining multiset
    private long getWays(int[] f, long targetK) {
        long ways = 1;
        int currLen = 0;
        for (int count : f) {
            if (count > 0) {
                currLen += count;
                long n = currLen;
                long r = count;
                
                if (r > n - r) r = n - r;
                long curNCr = 1;
                
                // Calculate combinations with an early termination mechanism
                for (int i = 1; i <= r; ++i) {
                    curNCr = curNCr * (n - i + 1) / i;
                    if (curNCr > targetK) {
                        curNCr = targetK + 1;
                        break;
                    }
                }
                ways *= curNCr;
                if (ways > targetK) return targetK + 1;
            }
        }
        return ways;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} k
 * @return {string}
 */
var smallestPalindrome = function(s, k) {
    // Collect character frequencies
    const freq = new Array(26).fill(0);
    for (let i = 0; i < s.length; i++) {
        freq[s.charCodeAt(i) - 97]++;
    }
    
    // Process frequencies for just one side of the palindrome
    const half = new Array(26).fill(0);
    let mid = "";
    let m = 0;
    
    for (let i = 0; i < 26; i++) {
        if (freq[i] % 2 !== 0) {
            mid += String.fromCharCode(i + 97);
        }
        half[i] = Math.floor(freq[i] / 2);
        m += half[i];
    }
    
    // Helper function to figure out multinomial coefficients safely
    const getWays = (f, targetK) => {
        let ways = 1;
        let currLen = 0;
        for (let i = 0; i < 26; i++) {
            const count = f[i];
            if (count > 0) {
                currLen += count;
                let n = currLen;
                let r = count;
                
                if (r > n - r) r = n - r;
                let curNCr = 1;
                
                // Iterative combinations, capped to prevent slow massive math
                for (let j = 1; j <= r; j++) {
                    curNCr = Math.floor(curNCr * (n - j + 1) / j);
                    if (curNCr > targetK) {
                        curNCr = targetK + 1;
                        break;
                    }
                }
                ways *= curNCr;
                if (ways > targetK) return targetK + 1;
            }
        }
        return ways;
    };
    
    // Verify k is valid based on total possible permutations
    if (getWays(half, k) < k) {
        return "";
    }
    
    let firstHalf = "";
    // Build first part sequentially
    for (let i = 0; i < m; i++) {
        for (let c = 0; c < 26; c++) {
            if (half[c] > 0) {
                half[c]--;
                const ways = getWays(half, k);
                
                // Found the right character for this position
                if (ways >= k) {
                    firstHalf += String.fromCharCode(c + 97);
                    break;
                } else {
                    // Update k and try the subsequent alphabet character
                    k -= ways;
                    half[c]++;
                }
            }
        }
    }
    
    // Merge first half, middle, and inverted first half
    return firstHalf + mid + firstHalf.split("").reverse().join("");
};
```

### Python3

```python
class Solution:
    def smallestPalindrome(self, s: str, k: int) -> str:
        from collections import Counter
        import math
        
        # Build raw frequencies of characters from the string
        freq = Counter(s)
        half = {}
        mid = ""
        m = 0
        
        # Calculate exactly half the occurrences and locate the odd character
        for char in "abcdefghijklmnopqrstuvwxyz":
            if freq[char] > 0:
                if freq[char] % 2 != 0:
                    mid += char
                half[char] = freq[char] // 2
                m += half[char]
        
        # Computes combinations using Python's native fast math
        def get_ways(f, target_k):
            ways = 1
            curr_len = 0
            for char in "abcdefghijklmnopqrstuvwxyz":
                count = f.get(char, 0)
                if count > 0:
                    curr_len += count
                    ways *= math.comb(curr_len, count)
                    # Returning early avoids computing huge factorials unnecessarily
                    if ways > target_k:
                        return target_k + 1
            return ways
            
        # Return blank string if k permutations do not exist
        if get_ways(half, k) < k:
            return ""
            
        first_half = []
        # Find exactly the right character for each step of the half length
        for _ in range(m):
            for char in "abcdefghijklmnopqrstuvwxyz":
                if half.get(char, 0) > 0:
                    half[char] -= 1
                    ways = get_ways(half, k)
                    
                    # Accept branch and move to next slot if ways encompasses k
                    if ways >= k:
                        first_half.append(char)
                        break
                    else:
                        # Otherwise, shift k and give back the char to check the next one
                        k -= ways
                        half[char] += 1
                        
        # Tie the pieces together in order
        first_str = "".join(first_half)
        return first_str + mid + first_str[::-1]
```

### Go

```go
func smallestPalindrome(s string, k int) string {
    // Map out the frequency of all items
    freq := make([]int, 26)
    for _, c := range s {
        freq[c-'a']++
    }
    
    // Get parameters for just the first half of the resulting string
    half := make([]int, 26)
    mid := ""
    m := 0
    for i := 0; i < 26; i++ {
        if freq[i] % 2 != 0 {
            mid += string(rune(i + 'a'))
        }
        half[i] = freq[i] / 2
        m += half[i]
    }
    
    // Utility closure to get remaining arrangements
    getWays := func(f []int, targetK int) int {
        ways := 1
        currLen := 0
        for _, count := range f {
            if count > 0 {
                currLen += count
                n := currLen
                r := count
                
                if r > n - r {
                    r = n - r
                }
                curNCr := 1
                
                // Generate nCr sequentially while capping at targetK
                for i := 1; i <= r; i++ {
                    curNCr = curNCr * (n - i + 1) / i
                    if curNCr > targetK {
                        curNCr = targetK + 1
                        break
                    }
                }
                ways *= curNCr
                if ways > targetK {
                    return targetK + 1
                }
            }
        }
        return ways
    }
    
    // Base requirement check to verify enough permutations exist
    if getWays(half, k) < k {
        return ""
    }
    
    firstHalf := []byte{}
    // Create the left half by evaluating each viable letter
    for i := 0; i < m; i++ {
        for c := 0; c < 26; c++ {
            if half[c] > 0 {
                half[c]--
                ways := getWays(half, k)
                
                // Commit to letter if it holds our target permutation
                if ways >= k {
                    firstHalf = append(firstHalf, byte(c+'a'))
                    break
                } else {
                    // Reduce offset and backtrack letter
                    k -= ways
                    half[c]++
                }
            }
        }
    }
    
    // String alignment
    res := string(firstHalf) + mid
    for i := len(firstHalf) - 1; i >= 0; i-- {
        res += string(firstHalf[i])
    }
    return res
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the containers and string-building style change.

First, the code counts how many times each character appears. This is important because palindrome construction depends completely on frequency, not on the original order of the string.

Next, it splits those counts into two parts:

* one part for the left half
* one character for the middle, if needed

If a character appears an odd number of times, one copy goes into the center. The rest are divided equally between the left and right halves.

After that, the helper counting function checks how many unique arrangements are possible with the current remaining characters. The implementation avoids overflow by stopping early once the count becomes larger than `k`. That is enough, because we only need to know whether the count reaches `k` or not.

Then the greedy construction starts. For every position in the left half, the algorithm tries letters from smallest to largest. For each candidate letter, it temporarily uses one copy and checks how many palindromes can still be formed.

If the number of valid completions is at least `k`, that letter is correct for the current position, so the algorithm keeps it.

If not, that branch is skipped, and `k` is reduced by the number of skipped palindromes. This means the answer lies later in lexicographic order.

At the end, the left half is reversed and attached to the middle character. That final mirror step gives the complete palindrome.

The language differences are small:

* **C++** uses `vector<int>`, lambdas, and string concatenation.
* **Java** uses arrays and `StringBuilder`.
* **JavaScript** uses arrays and string operations.
* **Python3** uses `Counter`, dictionaries, and slicing for reversal.
* **Go** uses slices and byte arrays for efficient character handling.

The core idea stays exactly the same in every language.

## Examples

### Example 1

**Input:** `s = "aabb"`, `k = 1`
**Possible palindromes:** `abba`, `baab`
**Expected Output:** `abba`

**Trace:**

* Left half counts become `a:1, b:1`
* Try `'a'` first for the first position
* The branch starting with `'a'` still contains the 1st answer
* So the first half becomes `"ab"`
* Final palindrome is `"abba"`

### Example 2

**Input:** `s = "aabb"`, `k = 2`
**Possible palindromes:** `abba`, `baab`
**Expected Output:** `baab`

**Trace:**

* Try `'a'` first
* The branch with `'a'` gives only the first palindrome
* Since `k = 2`, skip that branch and reduce `k`
* Choose `'b'` for the first character
* The remaining placement leads to `"baab"`

### Example 3

**Input:** `s = "aabbc"`, `k = 1`
**Expected Output:** `abcba`

**Trace:**

* Middle character is `c`
* Left half counts are `a:1, b:1`
* The smallest valid left half is `"ab"`
* Final palindrome is `"ab" + "c" + "ba"` = `"abcba"`

## How to Use / Run Locally

Here is how you can use each version once you paste the actual code into the empty blocks.

### C++

Save the file as `solution.cpp`, then compile and run:

```bash
g++ -std=c++17 -O2 -o solution solution.cpp
./solution
```

### Java

Save the file as `Solution.java`, then compile and run:

```bash
javac Solution.java
java Solution
```

### JavaScript

Save the file as `solution.js`, then run:

```bash
node solution.js
```

### Python3

Save the file as `solution.py`, then run:

```bash
python3 solution.py
```

### Go

Save the file as `main.go`, then run:

```bash
go run main.go
```

If the problem is on an online judge, make sure the function name and input format match the platform’s required template.

## Notes & Optimizations

* This is a much better approach than generating all palindromes and sorting them.
* The greedy method works because lexicographic order can be decided one character at a time.
* The helper function must stop early when counts become larger than `k`; otherwise, large factorial-like numbers can overflow.
* Using fixed arrays of size 26 keeps the solution simple and fast.
* Python’s `math.comb` makes the counting code cleaner, while C++, Java, JavaScript, and Go need manual capped counting.
* The approach assumes lowercase English letters. If the problem allows more characters, the frequency logic must be extended.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
