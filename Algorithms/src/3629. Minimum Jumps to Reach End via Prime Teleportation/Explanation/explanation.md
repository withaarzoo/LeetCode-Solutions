# Minimum Jumps to Reach End via Prime Teleportation

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

This problem asks us to find the minimum number of jumps needed to move from index `0` to index `n - 1` in an integer array.

From any index, we can do two types of moves:

1. Move to the left or right adjacent index
2. Use prime teleportation

Prime teleportation works only if the current number is prime. If it is prime, we can instantly jump to any index whose value is divisible by that prime number.

The goal is to return the minimum number of jumps required to reach the last index.

This is a classic shortest path style DSA problem that combines:

* BFS (Breadth First Search)
* Prime factorization
* Sieve of Eratosthenes
* Graph traversal optimization

---

## Constraints

| Constraint                 | Value                                                          |
| -------------------------- | -------------------------------------------------------------- |
| `1 <= nums.length <= 10^5` | Array size can be large                                        |
| `1 <= nums[i] <= 10^6`     | Values can be large enough to require optimized prime handling |

---

## Intuition

The first thing I noticed was that every move costs exactly one jump.

That usually points directly toward BFS because BFS always finds the shortest path in an unweighted graph.

The array itself behaves like a graph:

* each index is a node
* adjacent indices are connected
* teleportation creates extra edges

The tricky part is teleportation.

If I try checking every divisible index every time, the solution becomes too slow.

So instead of repeatedly searching the entire array, I preprocess everything first.

I store:

* every prime factor
* and all indices divisible by that factor

That way teleportation becomes fast lookup instead of repeated scanning.

---

## Approach

First, I preprocess the smallest prime factor for every number using a sieve.

This helps me:

* quickly check whether a number is prime
* quickly factorize numbers

Next, I create a mapping:

* prime factor → list of indices divisible by that factor

Then I run BFS from index `0`.

For every current index:

* move left if possible
* move right if possible
* if current value is prime, teleport to all valid indices

To avoid repeated work:

* once a teleport list is used, I clear it

This optimization is very important because otherwise the same teleport edges may be processed many times.

The first time BFS reaches the last index, that answer is guaranteed to be minimum.

---

## Data Structures Used

| Data Structure              | Purpose                                          |
| --------------------------- | ------------------------------------------------ |
| Queue                       | Used for BFS traversal                           |
| Hash Map / Dictionary       | Stores prime factor → indices mapping            |
| Visited / Distance Array    | Prevents revisiting nodes                        |
| Smallest Prime Factor Array | Helps with fast prime factorization              |
| Set                         | Avoids duplicate prime factors while factorizing |

---

## Operations & Behavior Summary

1. Find the maximum value in the array
2. Build the smallest prime factor sieve
3. Factorize every number
4. Store indices based on prime divisibility
5. Start BFS from index `0`
6. Try adjacent moves
7. If current number is prime:

   * teleport to all divisible indices
8. Mark visited indices
9. Clear teleport list after use
10. Return answer once last index is reached

---

## Complexity

| Type             | Complexity                 | Explanation                                          |
| ---------------- | -------------------------- | ---------------------------------------------------- |
| Time Complexity  | `O(n log M + M log log M)` | `n` is array size and `M` is maximum value in `nums` |
| Space Complexity | `O(n + M)`                 | Extra memory used for sieve, BFS queue, and mappings |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minJumps(vector<int>& nums) {
        int n = nums.size();

        // If already at last index
        if (n == 1) return 0;

        // Find maximum value for sieve
        int mx = *max_element(nums.begin(), nums.end());

        // Smallest prime factor array
        vector<int> spf(mx + 1);

        // Initialize SPF
        for (int i = 0; i <= mx; i++) {
            spf[i] = i;
        }

        // Build sieve for smallest prime factor
        for (int i = 2; i * i <= mx; i++) {
            if (spf[i] == i) {
                for (int j = i * i; j <= mx; j += i) {
                    if (spf[j] == j) {
                        spf[j] = i;
                    }
                }
            }
        }

        // Map prime factor -> indices divisible by it
        unordered_map<int, vector<int>> mp;

        // Build factor mapping
        for (int i = 0; i < n; i++) {
            int x = nums[i];

            unordered_set<int> used;

            // Get all unique prime factors
            while (x > 1) {
                int p = spf[x];

                if (!used.count(p)) {
                    mp[p].push_back(i);
                    used.insert(p);
                }

                x /= p;
            }
        }

        // BFS queue
        queue<int> q;

        // Distance array
        vector<int> dist(n, -1);

        q.push(0);
        dist[0] = 0;

        while (!q.empty()) {
            int i = q.front();
            q.pop();

            int steps = dist[i];

            // Reached end
            if (i == n - 1) {
                return steps;
            }

            // Move left
            if (i - 1 >= 0 && dist[i - 1] == -1) {
                dist[i - 1] = steps + 1;
                q.push(i - 1);
            }

            // Move right
            if (i + 1 < n && dist[i + 1] == -1) {
                dist[i + 1] = steps + 1;
                q.push(i + 1);
            }

            int val = nums[i];

            // Check if current number is prime
            if (val > 1 && spf[val] == val) {

                // Teleport to all divisible indices
                for (int nxt : mp[val]) {
                    if (dist[nxt] == -1) {
                        dist[nxt] = steps + 1;
                        q.push(nxt);
                    }
                }

                // Clear so we never process again
                mp[val].clear();
            }
        }

        return -1;
    }
};
```

### Java

```java
class Solution {
    public int minJumps(int[] nums) {
        int n = nums.length;

        // Already at destination
        if (n == 1) return 0;

        int mx = 0;

        // Find maximum value
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        // Smallest prime factor array
        int[] spf = new int[mx + 1];

        // Initialize SPF
        for (int i = 0; i <= mx; i++) {
            spf[i] = i;
        }

        // Sieve preprocessing
        for (int i = 2; i * i <= mx; i++) {
            if (spf[i] == i) {
                for (int j = i * i; j <= mx; j += i) {
                    if (spf[j] == j) {
                        spf[j] = i;
                    }
                }
            }
        }

        // Prime factor -> indices mapping
        Map<Integer, List<Integer>> map = new HashMap<>();

        for (int i = 0; i < n; i++) {
            int x = nums[i];

            Set<Integer> used = new HashSet<>();

            // Extract unique prime factors
            while (x > 1) {
                int p = spf[x];

                if (!used.contains(p)) {
                    map.computeIfAbsent(p, k -> new ArrayList<>()).add(i);
                    used.add(p);
                }

                x /= p;
            }
        }

        // BFS queue
        Queue<Integer> q = new LinkedList<>();

        // Distance array
        int[] dist = new int[n];

        Arrays.fill(dist, -1);

        q.offer(0);
        dist[0] = 0;

        while (!q.isEmpty()) {
            int i = q.poll();

            int steps = dist[i];

            // Reached end
            if (i == n - 1) {
                return steps;
            }

            // Move left
            if (i - 1 >= 0 && dist[i - 1] == -1) {
                dist[i - 1] = steps + 1;
                q.offer(i - 1);
            }

            // Move right
            if (i + 1 < n && dist[i + 1] == -1) {
                dist[i + 1] = steps + 1;
                q.offer(i + 1);
            }

            int val = nums[i];

            // Current value must be prime
            if (val > 1 && spf[val] == val) {

                List<Integer> list = map.getOrDefault(val, new ArrayList<>());

                // Teleport moves
                for (int nxt : list) {
                    if (dist[nxt] == -1) {
                        dist[nxt] = steps + 1;
                        q.offer(nxt);
                    }
                }

                // Clear after use
                list.clear();
            }
        }

        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minJumps = function(nums) {

    const n = nums.length;

    // Already at destination
    if (n === 1) return 0;

    // Find maximum value
    const mx = Math.max(...nums);

    // Smallest prime factor array
    const spf = Array(mx + 1).fill(0);

    // Initialize SPF
    for (let i = 0; i <= mx; i++) {
        spf[i] = i;
    }

    // Sieve preprocessing
    for (let i = 2; i * i <= mx; i++) {
        if (spf[i] === i) {
            for (let j = i * i; j <= mx; j += i) {
                if (spf[j] === j) {
                    spf[j] = i;
                }
            }
        }
    }

    // Prime factor -> indices
    const mp = new Map();

    for (let i = 0; i < n; i++) {

        let x = nums[i];

        const used = new Set();

        // Extract unique prime factors
        while (x > 1) {

            const p = spf[x];

            if (!used.has(p)) {

                if (!mp.has(p)) {
                    mp.set(p, []);
                }

                mp.get(p).push(i);

                used.add(p);
            }

            x = Math.floor(x / p);
        }
    }

    // BFS queue
    const q = [0];

    // Distance array
    const dist = Array(n).fill(-1);

    dist[0] = 0;

    let front = 0;

    while (front < q.length) {

        const i = q[front++];

        const steps = dist[i];

        // Reached destination
        if (i === n - 1) {
            return steps;
        }

        // Move left
        if (i - 1 >= 0 && dist[i - 1] === -1) {
            dist[i - 1] = steps + 1;
            q.push(i - 1);
        }

        // Move right
        if (i + 1 < n && dist[i + 1] === -1) {
            dist[i + 1] = steps + 1;
            q.push(i + 1);
        }

        const val = nums[i];

        // Current value must be prime
        if (val > 1 && spf[val] === val) {

            const list = mp.get(val) || [];

            // Teleport moves
            for (const nxt of list) {

                if (dist[nxt] === -1) {
                    dist[nxt] = steps + 1;
                    q.push(nxt);
                }
            }

            // Prevent repeated processing
            mp.set(val, []);
        }
    }

    return -1;
};
```

### Python3

```python
class Solution:
    def minJumps(self, nums: List[int]) -> int:

        n = len(nums)

        # Already at destination
        if n == 1:
            return 0

        mx = max(nums)

        # Smallest prime factor array
        spf = list(range(mx + 1))

        # Build sieve
        for i in range(2, int(mx ** 0.5) + 1):

            if spf[i] == i:

                for j in range(i * i, mx + 1, i):

                    if spf[j] == j:
                        spf[j] = i

        # Prime factor -> indices mapping
        mp = {}

        for i, val in enumerate(nums):

            x = val

            used = set()

            # Extract unique prime factors
            while x > 1:

                p = spf[x]

                if p not in used:

                    if p not in mp:
                        mp[p] = []

                    mp[p].append(i)

                    used.add(p)

                x //= p

        # BFS queue
        q = deque([0])

        # Distance array
        dist = [-1] * n

        dist[0] = 0

        while q:

            i = q.popleft()

            steps = dist[i]

            # Reached destination
            if i == n - 1:
                return steps

            # Move left
            if i - 1 >= 0 and dist[i - 1] == -1:
                dist[i - 1] = steps + 1
                q.append(i - 1)

            # Move right
            if i + 1 < n and dist[i + 1] == -1:
                dist[i + 1] = steps + 1
                q.append(i + 1)

            val = nums[i]

            # Teleport possible only if current value is prime
            if val > 1 and spf[val] == val:

                for nxt in mp.get(val, []):

                    if dist[nxt] == -1:
                        dist[nxt] = steps + 1
                        q.append(nxt)

                # Clear to avoid repeated work
                mp[val] = []

        return -1
```

### Go

```go
func minJumps(nums []int) int {

 n := len(nums)

 // Already at destination
 if n == 1 {
  return 0
 }

 // Find maximum value
 mx := 0

 for _, x := range nums {
  if x > mx {
   mx = x
  }
 }

 // Smallest prime factor array
 spf := make([]int, mx+1)

 // Initialize SPF
 for i := 0; i <= mx; i++ {
  spf[i] = i
 }

 // Build sieve
 for i := 2; i*i <= mx; i++ {

  if spf[i] == i {

   for j := i * i; j <= mx; j += i {

    if spf[j] == j {
     spf[j] = i
    }
   }
  }
 }

 // Prime factor -> indices mapping
 mp := map[int][]int{}

 for i, val := range nums {

  x := val

  used := map[int]bool{}

  // Extract unique prime factors
  for x > 1 {

   p := spf[x]

   if !used[p] {

    mp[p] = append(mp[p], i)

    used[p] = true
   }

   x /= p
  }
 }

 // BFS queue
 q := []int{0}

 // Distance array
 dist := make([]int, n)

 for i := 0; i < n; i++ {
  dist[i] = -1
 }

 dist[0] = 0

 front := 0

 for front < len(q) {

  i := q[front]
  front++

  steps := dist[i]

  // Reached destination
  if i == n-1 {
   return steps
  }

  // Move left
  if i-1 >= 0 && dist[i-1] == -1 {

   dist[i-1] = steps + 1

   q = append(q, i-1)
  }

  // Move right
  if i+1 < n && dist[i+1] == -1 {

   dist[i+1] = steps + 1

   q = append(q, i+1)
  }

  val := nums[i]

  // Teleport allowed only if value is prime
  if val > 1 && spf[val] == val {

   for _, nxt := range mp[val] {

    if dist[nxt] == -1 {

     dist[nxt] = steps + 1

     q = append(q, nxt)
    }
   }

   // Clear after use
   mp[val] = []int{}
  }
 }

 return -1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same across all five languages. Only syntax changes.

First, the solution creates a smallest prime factor array.

This array helps identify:

* whether a number is prime
* what prime factors a number contains

Instead of doing slow prime checking repeatedly, the sieve preprocesses everything once.

After that, the algorithm factorizes every number in the array.

For every prime factor found:

* the current index gets stored inside a mapping

Example:

```text
Prime 2 → indices divisible by 2
Prime 3 → indices divisible by 3
Prime 5 → indices divisible by 5
```

Now the graph connections are ready.

Then BFS begins.

The queue stores indices to process.

At every step:

* move left
* move right
* use teleportation if current value is prime

The visited array prevents infinite loops and repeated processing.

One important optimization is clearing teleport lists after they are used once.

Without this:

* the same teleportation edges may be processed repeatedly
* performance becomes much slower

Since BFS explores level by level:

* the first time we reach the last index
* that path is automatically the shortest

That is why BFS is the perfect algorithm for this problem.

---

## Examples

### Example 1

Input:

```text
nums = [1,2,4,6]
```

Output:

```text
2
```

Explanation:

* Start at index `0`
* Move to index `1`
* `nums[1] = 2` which is prime
* Teleport to index `3` because `6 % 2 == 0`

Total jumps = `2`

---

### Example 2

Input:

```text
nums = [2,3,4,7,9]
```

Output:

```text
2
```

Explanation:

* Start at index `0`
* Move to index `1`
* `nums[1] = 3` which is prime
* Teleport to index `4` because `9 % 3 == 0`

Total jumps = `2`

---

### Example 3

Input:

```text
nums = [4,6,5,8]
```

Output:

```text
3
```

Explanation:

No useful teleportation exists.

Path:

* `0 → 1 → 2 → 3`

Total jumps = `3`

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
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

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

* BFS is required because every move has equal cost
* Clearing teleport lists avoids repeated traversal
* Using smallest prime factor preprocessing is much faster than repeated trial division
* The graph is built implicitly instead of explicitly storing all edges
* This solution easily handles large constraints up to `10^5`
* Prime factor preprocessing is the key optimization in this problem
* An unoptimized teleportation search would cause Time Limit Exceeded (TLE)

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
