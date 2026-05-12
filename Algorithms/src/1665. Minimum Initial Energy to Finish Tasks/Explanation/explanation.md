# 1665. Minimum Initial Energy to Finish Tasks

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

In this LeetCode greedy algorithm problem, we are given a list of tasks where each task contains:

* `actual[i]` → the amount of energy spent after completing the task
* `minimum[i]` → the minimum energy required before starting the task

The goal is to find the minimum initial energy needed to finish all tasks in any order.

This is a classic greedy sorting problem in Data Structures and Algorithms (DSA). The main challenge is deciding the best order to complete tasks so the starting energy stays as small as possible.

Input:

* A 2D array called `tasks`

Output:

* The minimum initial energy required to complete every task

This problem is popular in coding interviews because it tests:

* Greedy strategy
* Sorting logic
* Optimization thinking
* Simulation-based problem solving

---

## Constraints

| Constraint                           | Value             |
| ------------------------------------ | ----------------- |
| `1 <= tasks.length <= 10^5`          | Large input size  |
| `1 <= actual_i <= minimum_i <= 10^4` | Valid task values |

---

## Intuition

The first thing I noticed was that task order matters a lot.

If I complete a task with a huge minimum requirement too late, I may not have enough energy left to even start it.

So instead of focusing only on energy loss (`actual`), I focused on how demanding a task is before starting it.

I realized that tasks with a larger:

```text
minimum - actual
```

should usually come first.

Why?

Because those tasks require high starting energy but consume relatively less afterward. Finishing them earlier prevents problems later.

That observation leads directly to a greedy sorting approach.

---

## Approach

I used a greedy algorithm with sorting.

Step-by-step process:

1. Sort all tasks in descending order of:

   ```text
   minimum - actual
   ```

2. Keep track of:

   * current available energy
   * total initial energy needed

3. Process tasks one by one.

4. Before starting a task:

   * if current energy is smaller than the required minimum,
   * add the missing amount

5. Complete the task and subtract its actual cost.

This guarantees that:

* every task can be started
* the starting energy stays minimum

This approach is both fast and optimal for large inputs.

---

## Data Structures Used

| Data Structure        | Why It Was Used                           |
| --------------------- | ----------------------------------------- |
| Array / Vector / List | To store all tasks                        |
| Sorting Algorithm     | To process tasks in the best greedy order |
| Integer Variables     | To track current energy and answer        |

No advanced data structures are needed for this problem.

---

## Operations & Behavior Summary

Here is what the algorithm does internally:

1. Read all tasks
2. Compare tasks using:

   ```text
   minimum - actual
   ```

3. Sort tasks from largest difference to smallest
4. Start with zero energy
5. For each task:

   * check if energy is enough
   * add extra energy if required
   * perform the task
   * reduce energy
6. Return the total added energy

This behaves like a simulation after sorting.

---

## Complexity

| Type             | Complexity         | Explanation                                   |
| ---------------- | ------------------ | --------------------------------------------- |
| Time Complexity  | `O(n log n)`       | Sorting all tasks takes `O(n log n)`          |
| Space Complexity | `O(1)` extra space | Only a few variables are used besides sorting |

Where:

* `n` = number of tasks

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumEffort(vector<vector<int>>& tasks) {

        // Sort tasks by (minimum - actual) in descending order
        sort(tasks.begin(), tasks.end(), [](vector<int>& a, vector<int>& b) {
            return (a[1] - a[0]) > (b[1] - b[0]);
        });

        int answer = 0; // Minimum initial energy required
        int energy = 0; // Current available energy

        // Process tasks one by one
        for (auto& task : tasks) {

            int actual = task[0];
            int minimum = task[1];

            // If current energy is not enough,
            // add the missing amount
            if (energy < minimum) {

                int need = minimum - energy;

                answer += need;
                energy += need;
            }

            // Complete the task
            energy -= actual;
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int minimumEffort(int[][] tasks) {

        // Sort by (minimum - actual) in descending order
        Arrays.sort(tasks, (a, b) -> (b[1] - b[0]) - (a[1] - a[0]));

        int answer = 0; // Minimum starting energy
        int energy = 0; // Current energy

        // Process every task
        for (int[] task : tasks) {

            int actual = task[0];
            int minimum = task[1];

            // Add extra energy if needed
            if (energy < minimum) {

                int need = minimum - energy;

                answer += need;
                energy += need;
            }

            // Finish the task
            energy -= actual;
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} tasks
 * @return {number}
 */
var minimumEffort = function(tasks) {

    // Sort by (minimum - actual) descending
    tasks.sort((a, b) => (b[1] - b[0]) - (a[1] - a[0]));

    let answer = 0; // Minimum initial energy
    let energy = 0; // Current energy

    // Process all tasks
    for (let [actual, minimum] of tasks) {

        // If energy is insufficient,
        // increase it
        if (energy < minimum) {

            let need = minimum - energy;

            answer += need;
            energy += need;
        }

        // Complete the task
        energy -= actual;
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def minimumEffort(self, tasks: List[List[int]]) -> int:

        # Sort by (minimum - actual) in descending order
        tasks.sort(key=lambda x: (x[1] - x[0]), reverse=True)

        answer = 0  # Minimum initial energy
        energy = 0  # Current available energy

        # Process every task
        for actual, minimum in tasks:

            # If current energy is less than required,
            # add the missing amount
            if energy < minimum:

                need = minimum - energy

                answer += need
                energy += need

            # Finish the task
            energy -= actual

        return answer
```

### Go

```go
func minimumEffort(tasks [][]int) int {

    // Sort by (minimum - actual) in descending order
    sort.Slice(tasks, func(i, j int) bool {
        return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
    })

    answer := 0 // Minimum initial energy
    energy := 0 // Current energy

    // Process all tasks
    for _, task := range tasks {

        actual := task[0]
        minimum := task[1]

        // If energy is not enough,
        // increase it
        if energy < minimum {

            need := minimum - energy

            answer += need
            energy += need
        }

        // Complete the task
        energy -= actual
    }

    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same across all five languages.

Only syntax changes.

### Step 1 — Sort the Tasks

The most important part of the solution is sorting.

I sort tasks using:

```text
minimum - actual
```

in descending order.

This means:

* tasks with higher starting requirements are handled earlier
* risky tasks are completed before energy becomes too low

Without sorting, the answer may become larger than necessary.

---

### Step 2 — Track Current Energy

I keep two variables:

| Variable | Purpose                         |
| -------- | ------------------------------- |
| `energy` | Current available energy        |
| `answer` | Minimum initial energy required |

At the beginning:

* both values start at `0`

---

### Step 3 — Check Task Requirement

Before starting a task:

```text
if energy < minimum
```

then I do not have enough energy to start.

So I calculate:

```text
minimum - energy
```

and add that amount.

This increases:

* current energy
* total starting energy

---

### Step 4 — Complete the Task

After starting the task successfully:

```text
energy -= actual
```

because the task consumes energy.

---

### Why This Greedy Strategy Works

Suppose two tasks are:

```text
A = [1, 10]
B = [5, 6]
```

If I do B first:

* energy drops faster
* task A may become impossible later

If I do A first:

* high requirement gets handled early
* remaining tasks become easier

That is why sorting by:

```text
minimum - actual
```

works correctly.

---

### Language Differences

#### C++

* Uses `sort()` with a custom lambda comparator
* Fast and memory efficient

#### Java

* Uses `Arrays.sort()` with a comparator
* Very clean for interview settings

#### JavaScript

* Uses built-in `.sort()`
* Comparator returns descending difference order

#### Python3

* Uses `sort(key=..., reverse=True)`
* Most compact implementation

#### Go

* Uses `sort.Slice()`
* Comparator handles custom sorting logic

The algorithm itself remains identical in every language.

---

## Examples

### Example 1

Input:

```text
tasks = [[1,2],[2,4],[4,8]]
```

Output:

```text
8
```

Explanation:

Sorted order:

```text
[[4,8],[2,4],[1,2]]
```

Start with energy `8`

Process:

* Complete `[4,8]` → remaining `4`
* Complete `[2,4]` → remaining `2`
* Complete `[1,2]` → remaining `1`

Minimum starting energy = `8`

---

### Example 2

Input:

```text
tasks = [[1,3],[2,4],[10,11],[10,12],[8,9]]
```

Output:

```text
32
```

Explanation:

The greedy order prevents high minimum tasks from failing later.

The algorithm keeps adding energy only when absolutely necessary.

---

### Example 3

Input:

```text
tasks = [[1,7],[2,8],[3,9],[4,10]]
```

Output:

```text
19
```

Explanation:

Tasks with bigger minimum requirements are processed first.

This minimizes the required initial energy.

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

* Sorting is the key optimization here.
* A brute force approach checking every order would be impossible because permutations grow extremely fast.
* This greedy solution handles large constraints efficiently.
* The algorithm works well even when:

  * all tasks have the same values
  * one task has a very large minimum requirement
  * task count reaches `10^5`

Possible alternative ideas:

* Binary search on answer
* Priority queue simulation

But those approaches are either slower or more complicated than necessary.

The greedy sorting method is the cleanest and most optimized solution.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
