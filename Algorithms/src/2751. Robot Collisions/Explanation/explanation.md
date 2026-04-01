# 2751. Robot Collisions

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

There are `n` robots placed on a line.

Each robot has:

* A position
* A health value
* A moving direction (`L` or `R`)

All robots start moving at the same speed.

If two robots collide:

* The robot with smaller health gets removed.
* The robot with larger health survives and loses `1` health.
* If both robots have equal health, both robots are removed.

I need to return the health values of the surviving robots in the same order as the input.

## Constraints

```text
1 <= positions.length == healths.length == directions.length <= 10^5
1 <= positions[i], healths[i] <= 10^9
positions[i] are unique
 directions[i] is either 'L' or 'R'
```

## Intuition

I noticed that collisions only happen when:

* One robot is moving right
* Another robot after it is moving left

That means I do not need to compare every pair of robots.

I can sort robots by their positions and process them from left to right.

Whenever I see a robot moving right, I store it.

Whenever I see a robot moving left, I try to collide it with the latest right-moving robot.

This is exactly the type of problem where a stack works very well.

## Approach

1. Store every robot with its original index.
2. Sort robots based on position.
3. Use a stack to keep robots moving right.
4. If current robot is moving right, push it into the stack.
5. If current robot is moving left:

   * Compare it with the top robot from the stack.
   * Keep colliding until:

     * Current robot dies
     * Stack becomes empty
6. At the end, return all robots whose health is greater than `0`.

## Data Structures Used

* Array / Vector / List

  * To store robot indices and final answers.

* Stack

  * To store robots moving right.

* Sorting

  * To process robots in position order.

## Operations & Behavior Summary

| Situation                              | Result                                             |
| -------------------------------------- | -------------------------------------------------- |
| Right robot health < Left robot health | Right robot dies, left robot health decreases by 1 |
| Right robot health > Left robot health | Left robot dies, right robot health decreases by 1 |
| Both health equal                      | Both robots die                                    |
| No right-moving robot available        | Left robot survives                                |

## Complexity

* Time Complexity: `O(n log n)`

  * Sorting takes `O(n log n)`
  * Stack processing takes `O(n)`

* Space Complexity: `O(n)`

  * Extra stack and index array are used.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int>& positions, vector<int>& healths, string directions) {
        int n = positions.size();

        vector<int> indices(n);
        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        sort(indices.begin(), indices.end(), [&](int a, int b) {
            return positions[a] < positions[b];
        });

        stack<int> st;

        for (int idx : indices) {
            if (directions[idx] == 'R') {
                st.push(idx);
            } else {
                while (!st.empty() && healths[idx] > 0) {
                    int topIdx = st.top();

                    if (healths[topIdx] < healths[idx]) {
                        st.pop();
                        healths[idx]--;
                        healths[topIdx] = 0;
                    } else if (healths[topIdx] == healths[idx]) {
                        st.pop();
                        healths[topIdx] = 0;
                        healths[idx] = 0;
                    } else {
                        healths[topIdx]--;
                        healths[idx] = 0;
                    }
                }
            }
        }

        vector<int> result;
        for (int health : healths) {
            if (health > 0) {
                result.push_back(health);
            }
        }

        return result;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length;

        Integer[] indices = new Integer[n];
        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        Arrays.sort(indices, (a, b) -> positions[a] - positions[b]);

        Stack<Integer> stack = new Stack<>();

        for (int idx : indices) {
            if (directions.charAt(idx) == 'R') {
                stack.push(idx);
            } else {
                while (!stack.isEmpty() && healths[idx] > 0) {
                    int topIdx = stack.peek();

                    if (healths[topIdx] < healths[idx]) {
                        stack.pop();
                        healths[idx]--;
                        healths[topIdx] = 0;
                    } else if (healths[topIdx] == healths[idx]) {
                        stack.pop();
                        healths[topIdx] = 0;
                        healths[idx] = 0;
                    } else {
                        healths[topIdx]--;
                        healths[idx] = 0;
                    }
                }
            }
        }

        List<Integer> result = new ArrayList<>();
        for (int health : healths) {
            if (health > 0) {
                result.add(health);
            }
        }

        return result;
    }
}
```

### JavaScript

```javascript
var survivedRobotsHealths = function(positions, healths, directions) {
    const n = positions.length;

    const indices = Array.from({ length: n }, (_, i) => i);
    indices.sort((a, b) => positions[a] - positions[b]);

    const stack = [];

    for (const idx of indices) {
        if (directions[idx] === 'R') {
            stack.push(idx);
        } else {
            while (stack.length > 0 && healths[idx] > 0) {
                const topIdx = stack[stack.length - 1];

                if (healths[topIdx] < healths[idx]) {
                    stack.pop();
                    healths[idx]--;
                    healths[topIdx] = 0;
                } else if (healths[topIdx] === healths[idx]) {
                    stack.pop();
                    healths[topIdx] = 0;
                    healths[idx] = 0;
                } else {
                    healths[topIdx]--;
                    healths[idx] = 0;
                }
            }
        }
    }

    const result = [];

    for (const health of healths) {
        if (health > 0) {
            result.push(health);
        }
    }

    return result;
};
```

### Python3

```python
class Solution:
    def survivedRobotsHealths(self, positions, healths, directions):
        n = len(positions)

        indices = list(range(n))
        indices.sort(key=lambda i: positions[i])

        stack = []

        for idx in indices:
            if directions[idx] == 'R':
                stack.append(idx)
            else:
                while stack and healths[idx] > 0:
                    top_idx = stack[-1]

                    if healths[top_idx] < healths[idx]:
                        stack.pop()
                        healths[idx] -= 1
                        healths[top_idx] = 0
                    elif healths[top_idx] == healths[idx]:
                        stack.pop()
                        healths[top_idx] = 0
                        healths[idx] = 0
                    else:
                        healths[top_idx] -= 1
                        healths[idx] = 0

        result = []
        for health in healths:
            if health > 0:
                result.append(health)

        return result
```

### Go

```go
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    n := len(positions)

    indices := make([]int, n)
    for i := 0; i < n; i++ {
        indices[i] = i
    }

    sort.Slice(indices, func(i, j int) bool {
        return positions[indices[i]] < positions[indices[j]]
    })

    stack := []int{}

    for _, idx := range indices {
        if directions[idx] == 'R' {
            stack = append(stack, idx)
        } else {
            for len(stack) > 0 && healths[idx] > 0 {
                topIdx := stack[len(stack)-1]

                if healths[topIdx] < healths[idx] {
                    stack = stack[:len(stack)-1]
                    healths[idx]--
                    healths[topIdx] = 0
                } else if healths[topIdx] == healths[idx] {
                    stack = stack[:len(stack)-1]
                    healths[topIdx] = 0
                    healths[idx] = 0
                } else {
                    healths[topIdx]--
                    healths[idx] = 0
                }
            }
        }
    }

    result := []int{}
    for _, health := range healths {
        if health > 0 {
            result = append(result, health)
        }
    }

    return result
}
```

## Step-by-step Detailed Explanation

### C++

* Create an index array from `0` to `n-1`
* Sort indices according to robot positions
* Use a stack for robots moving right
* If current robot moves left, compare with stack top
* Update health values after every collision
* At the end, collect all robots with health greater than `0`

### Java

* Store all indices in an Integer array
* Sort the array using positions
* Use Stack to store right-moving robots
* For every left-moving robot, process collisions one by one
* Keep reducing health after every collision
* Return surviving robot healths

### JavaScript

* Build an array of indices
* Sort them by positions
* Use an array as stack
* Push right-moving robots
* Resolve collisions when left-moving robot appears
* Store only remaining health values in answer

### Python3

* Create a sorted index list
* Use a stack for right-moving robots
* Process collisions while stack is not empty
* Reduce health accordingly
* Collect surviving robots in original order

### Go

* Create index slice
* Sort it based on positions
* Use slice as stack
* Compare left-moving robot with latest right-moving robot
* Continue collisions until one robot dies
* Store final surviving healths

## Examples

```text
Input:
positions = [3,5,2,6]
healths = [10,10,15,12]
directions = "RLRL"

Output:
[14]
```

Explanation:

```text
Robot at position 3 and robot at position 5 collide.
Both have same health, so both die.

Robot at position 2 and robot at position 6 collide.
Robot with health 15 survives and becomes 14.
```

---

```text
Input:
positions = [1,2,5,6]
healths = [10,10,11,11]
directions = "RLRL"

Output:
[]
```

Explanation:

```text
First collision removes both robots.
Second collision also removes both robots.
No robot survives.
```

## How to use / Run locally

### C++

```bash
g++ solution.cpp -o solution
./solution
```

### Java

```bash
javac Solution.java
java Solution
```

### JavaScript

```bash
node solution.js
```

### Python3

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

## Notes & Optimizations

* Sorting is necessary because robots are not given in position order.
* Stack helps process only meaningful collisions.
* Every robot is pushed and popped at most once.
* This makes the collision simulation efficient.
* Brute force collision simulation would be too slow for `10^5` robots.

## Author

* [Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
