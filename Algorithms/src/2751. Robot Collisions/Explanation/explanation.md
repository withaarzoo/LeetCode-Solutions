# Robot Collisions Problem

This repository contains solutions in multiple programming languages (C++, Java, JavaScript, Python, Go) for the "Robot Collisions" problem on LeetCode. The problem involves simulating the movements of robots on a line and handling collisions based on their health.

## Problem Description

There are `n` robots each with a position, health, and movement direction ('L' for left or 'R' for right). Robots move simultaneously in their directions, and if two robots collide at the same position, the robot with lower health is removed. If both have the same health, both are removed.

### Example

Input:

- Positions: [5, 4, 3, 2, 1]
- Healths: [2, 17, 9, 15, 10]
- Directions: "RRRRR"

Output:

- [2, 17, 9, 15, 10]

Explanation:
No collisions occur since all robots move in the same direction ('R'). Thus, the healths remain unchanged.

## C++ Solution

```cpp
class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int>& positions, vector<int>& healths, string directions) {
        int n = positions.size();
        vector<int> indices(n), result;
        stack<int> stack;

        // Create indices array for sorting
        for (int index = 0; index < n; ++index) {
            indices[index] = index;
        }

        // Sort indices based on positions
        sort(indices.begin(), indices.end(),
             [&](int lhs, int rhs) { return positions[lhs] < positions[rhs]; });

        // Process robots according to sorted order
        for (int currentIndex : indices) {
            if (directions[currentIndex] == 'R') {
                stack.push(currentIndex);
            } else {
                // Handle left-moving robots
                while (!stack.empty() && healths[currentIndex] > 0) {
                    int topIndex = stack.top();
                    stack.pop();

                    if (healths[topIndex] > healths[currentIndex]) {
                        healths[topIndex] -= 1;
                        healths[currentIndex] = 0;
                        stack.push(topIndex);
                    } else if (healths[topIndex] < healths[currentIndex]) {
                        healths[currentIndex] -= 1;
                        healths[topIndex] = 0;
                    } else {
                        healths[currentIndex] = 0;
                        healths[topIndex] = 0;
                    }
                }
            }
        }

        // Collect surviving robot healths
        for (int index = 0; index < n; ++index) {
            if (healths[index] > 0) {
                result.push_back(healths[index]);
            }
        }
        return result;
    }
};
```

### Explanation

1. **Initialization**: Initialize `indices` to keep track of original indices, `stack` to manage right-moving robots, and `result` to store surviving robot healths.
2. **Sorting**: Sort `indices` based on robot positions to process robots in order of their positions.
3. **Processing Robots**: Iterate over sorted indices. If robot moves right ('R'), push its index onto `stack`. If left ('L'), handle collisions with robots in `stack`.
4. **Collision Handling**: Compare healths of colliding robots. Update healths accordingly and mark robots as removed if healths are equal.
5. **Result Collection**: After processing all robots, collect healths of surviving robots in `result`.

---

## Java Solution

```java
import java.util.*;

class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length;
        List<Integer> result = new ArrayList<>();
        int[] indices = new int[n];
        Stack<Integer> stack = new Stack<>();

        // Create indices array for sorting
        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        // Sort indices based on positions
        Arrays.sort(indices, (a, b) -> Integer.compare(positions[a], positions[b]));

        // Process robots according to sorted order
        for (int currentIndex : indices) {
            if (directions.charAt(currentIndex) == 'R') {
                stack.push(currentIndex);
            } else {
                // Handle left-moving robots
                while (!stack.isEmpty() && healths[currentIndex] > 0) {
                    int topIndex = stack.pop();

                    if (healths[topIndex] > healths[currentIndex]) {
                        healths[topIndex] -= 1;
                        healths[currentIndex] = 0;
                        stack.push(topIndex);
                    } else if (healths[topIndex] < healths[currentIndex]) {
                        healths[currentIndex] -= 1;
                        healths[topIndex] = 0;
                    } else {
                        healths[currentIndex] = 0;
                        healths[topIndex] = 0;
                    }
                }
            }
        }

        // Collect surviving robot healths
        for (int i = 0; i < n; i++) {
            if (healths[i] > 0) {
                result.add(healths[i]);
            }
        }
        return result;
    }
}
```

### Explanation

1. **Initialization**: Initialize `indices` to keep track of original indices, `stack` to manage right-moving robots, and `result` to store surviving robot healths.
2. **Sorting**: Sort `indices` based on robot positions to process robots in order of their positions.
3. **Processing Robots**: Iterate over sorted indices. If robot moves right ('R'), push its index onto `stack`. If left ('L'), handle collisions with robots in `stack`.
4. **Collision Handling**: Compare healths of colliding robots. Update healths accordingly and mark robots as removed if healths are equal.
5. **Result Collection**: After processing all robots, collect healths of surviving robots in `result`.

---

## JavaScript Solution

```javascript
var survivedRobotsHealths = function(positions, healths, directions) {
    let n = positions.length;
    let indices = Array.from({length: n}, (_, i) => i);
    let stack = [];
    let result = [];

    // Sort indices based on positions
    indices.sort((a, b) => positions[a] - positions[b]);

    // Process robots according to sorted order
    for (let currentIndex of indices) {
        if (directions[currentIndex] === 'R') {
            stack.push(currentIndex);
        } else {
            // Handle left-moving robots
            while (stack.length > 0 && healths[currentIndex] > 0) {
                let topIndex = stack.pop();

                if (healths[topIndex] > healths[currentIndex]) {
                    healths[topIndex] -= 1;
                    healths[currentIndex] = 0;
                    stack.push(topIndex);
                } else if (healths[topIndex] < healths[currentIndex]) {
                    healths[currentIndex] -= 1;
                    healths[topIndex] = 0;
                } else {
                    healths[currentIndex] = 0;
                    healths[topIndex] = 0;
                }
            }
        }
    }

    // Collect surviving robot healths
    for (let i = 0; i < n; i++) {
        if (healths[i] > 0) {
            result.push(healths[i]);
        }
    }

    return result;
};
```

### Explanation

1. **Initialization**: Initialize `indices` to keep track of original indices, `stack` to manage right-moving robots, and `result` to store surviving robot healths.
2. **Sorting**: Sort `indices` based on robot positions to process robots in order of their positions.
3. **Processing Robots**: Iterate over sorted indices. If robot moves right ('R'), push its index onto `stack`. If left ('L'), handle collisions with robots in `stack`.
4. **Collision Handling**: Compare healths of colliding robots. Update healths accordingly and mark robots as removed if healths are equal.
5. **Result Collection**: After processing all robots, collect healths of surviving robots in `result`.

---

## Python Solution

```python
class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        n = len(positions)
        indices = list(range(n))
        stack = []
        result = []

        # Sort indices based on positions
        indices.sort(key=lambda x: positions[x])

        # Process robots according to sorted order
        for currentIndex in indices:
            if directions[currentIndex] == 'R':
                stack.append(currentIndex)
            else:
                # Handle left-moving robots
                while stack and healths[currentIndex] > 0:
                    topIndex = stack.pop()

                    if healths[topIndex] > healths[currentIndex]:
                        healths[topIndex] -= 1
                        healths[currentIndex] = 0
                        stack.append(topIndex)
                    elif healths[topIndex] < healths[currentIndex]:
                        healths[currentIndex] -= 1
                        healths[topIndex] = 0
                    else:
                        healths[currentIndex] = 0
                        healths[topIndex] = 0

        # Collect surviving robot healths
        for i in range(n):
            if healths[i] > 0:
                result.append(healths[i])

        return result
``

`

### Explanation

1. **Initialization**: Initialize `indices` to keep track of original indices, `stack` to manage right-moving robots, and `result` to store surviving robot healths.
2. **Sorting**: Sort `indices` based on robot positions to process robots in order of their positions.
3. **Processing Robots**: Iterate over sorted indices. If robot moves right ('R'), push its index onto `stack`. If left ('L'), handle collisions with robots in `stack`.
4. **Collision Handling**: Compare healths of colliding robots. Update healths accordingly and mark robots as removed if healths are equal.
5. **Result Collection**: After processing all robots, collect healths of surviving robots in `result`.

---

## Go Solution

```go
package main

import (
    "sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    n := len(positions)
    indices := make([]int, n)
    stack := []int{}
    result := []int{}

    // Create indices array for sorting
    for i := range indices {
        indices[i] = i
    }

    // Sort indices based on positions
    sort.Slice(indices, func(a, b int) bool {
        return positions[a] < positions[b]
    })

    // Process robots according to sorted order
    for _, currentIndex := range indices {
        if directions[currentIndex] == 'R' {
            stack = append(stack, currentIndex)
        } else {
            // Handle left-moving robots
            for len(stack) > 0 && healths[currentIndex] > 0 {
                topIndex := stack[len(stack)-1]
                stack = stack[:len(stack)-1]

                if healths[topIndex] > healths[currentIndex] {
                    healths[topIndex] -= 1
                    healths[currentIndex] = 0
                    stack = append(stack, topIndex)
                } else if healths[topIndex] < healths[currentIndex] {
                    healths[currentIndex] -= 1
                    healths[topIndex] = 0
                } else {
                    healths[currentIndex] = 0
                    healths[topIndex] = 0
                }
            }
        }
    }

    // Collect surviving robot healths
    for i := 0; i < n; i++ {
        if healths[i] > 0 {
            result = append(result, healths[i])
        }
    }

    return result
}
```

### Explanation

1. **Initialization**: Initialize `indices` to keep track of original indices, `stack` to manage right-moving robots, and `result` to store surviving robot healths.
2. **Sorting**: Sort `indices` based on robot positions to process robots in order of their positions.
3. **Processing Robots**: Iterate over sorted indices. If robot moves right ('R'), push its index onto `stack`. If left ('L'), handle collisions with robots in `stack`.
4. **Collision Handling**: Compare healths of colliding robots. Update healths accordingly and mark robots as removed if healths are equal.
5. **Result Collection**: After processing all robots, collect healths of surviving robots in `result`.

---

These solutions provide efficient ways to handle robot collisions and collect the healths of surviving robots using different programming languages. Each solution follows a similar approach but adapts to the specific syntax and idioms of each language.
