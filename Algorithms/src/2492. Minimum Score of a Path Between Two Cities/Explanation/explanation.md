# 2492. Minimum Score of a Path Between Two Cities

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

In LeetCode 2492, Minimum Score of a Path Between Two Cities, we are given `n` cities numbered from `1` to `n`.

The cities are connected by bidirectional roads. Each road has a distance.

A road is represented as:

`[a, b, distance]`

This means there is a road between city `a` and city `b` with the given `distance`.

The score of a path is not the total distance of all roads. Instead, it is the minimum road distance used anywhere in that path.

The goal is to find the minimum possible score of any path from city `1` to city `n`.

One important rule makes this graph problem easier: a path can visit the same city multiple times and can also use the same road multiple times.

The problem guarantees that at least one path exists between city `1` and city `n`.

## Constraints

| Constraint          | Value                                                  |
| ------------------- | ------------------------------------------------------ |
| Number of cities    | `2 <= n <= 10^5`                                       |
| Number of roads     | `1 <= roads.length <= 10^5`                            |
| Road format         | `roads[i].length == 3`                                 |
| City numbers        | `1 <= a_i, b_i <= n`                                   |
| Different endpoints | `a_i != b_i`                                           |
| Road distance       | `1 <= distance_i <= 10^4`                              |
| Repeated edges      | No repeated edges                                      |
| Path guarantee      | At least one path exists between city `1` and city `n` |

These constraints mean an efficient graph traversal is required. A slow approach that repeatedly scans all roads would not work well for up to `100,000` cities and `100,000` roads.

## Intuition

My first thought was to find the best path from city `1` to city `n`.

Normally, a weighted graph problem might suggest Dijkstra's algorithm or some kind of shortest path technique. But this problem is different because the score is not the sum of road distances.

The score only depends on the smallest road used in the path.

Then I noticed the most important rule: I can visit cities and roads multiple times.

Because of that, any road inside the connected component containing city `1` can be included in a valid path from city `1` to city `n`.

I can travel from city `1` to that road, use it, and then continue toward city `n`. If necessary, I can go back through cities or roads I already visited.

So I do not need to search for one special route.

I only need to find the smallest road distance in the entire connected component containing city `1`.

Since the problem guarantees that city `n` is reachable from city `1`, both cities are already in the same connected component.

## Approach

I use a graph traversal approach with an adjacency list and BFS.

The same solution can also be written with DFS, but BFS avoids recursion depth problems for large graphs.

The process is simple:

1. Build an undirected adjacency list from the `roads` array.
2. Start BFS from city `1`.
3. Keep a visited array so each city is processed only once.
4. For every road connected to a visited city, compare its distance with the current answer.
5. Store the smallest road distance found so far.
6. Add every unvisited neighboring city to the BFS queue.
7. Continue until the whole connected component of city `1` has been explored.
8. Return the smallest road distance found.

A key detail is that every road distance must be checked, even when the neighboring city has already been visited.

The visited array controls city traversal. It should not decide whether a road can affect the answer.

## Data Structures Used

### Adjacency List

The adjacency list stores every city along with its neighboring cities and road distances.

For a road between `a` and `b`, I store:

* `b` and the distance in the adjacency list of `a`
* `a` and the distance in the adjacency list of `b`

This is needed because all roads are bidirectional.

An adjacency list is a good choice because the graph can have up to `100,000` cities and roads. It uses much less memory than an adjacency matrix.

### Visited Array

The visited array keeps track of cities that have already been discovered.

Without it, the traversal could keep moving between the same connected cities forever because the graph is undirected.

Each city is marked as visited when it is added to the queue.

### Queue

The queue is used for Breadth-First Search.

It stores cities that have been discovered but whose roads have not been fully checked yet.

BFS is useful here because it can explore the complete connected component without recursion.

### Answer Variable

A single variable stores the minimum road distance found during traversal.

It starts with a very large value and is updated whenever a smaller road distance is found.

## Operations & Behavior Summary

The algorithm works in the following order:

1. Create an empty adjacency list for all `n` cities.
2. Read every road.
3. Add the road in both directions because the graph is undirected.
4. Create a visited array.
5. Add city `1` to the BFS queue.
6. Mark city `1` as visited.
7. Start with the answer set to a very large value.
8. Remove one city from the queue.
9. Check every road connected to that city.
10. Update the answer with the smaller value between the current answer and the road distance.
11. If the neighboring city is unvisited, mark it as visited and add it to the queue.
12. Repeat until the queue becomes empty.
13. Return the smallest road distance found.

At the end, every city and road in the connected component containing city `1` has been considered.

## Complexity

| Complexity       | Value      | Explanation                                                                                                |
| ---------------- | ---------- | ---------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n + m)` | Each city is visited at most once, and each road is checked at most twice because the graph is undirected. |
| Space Complexity | `O(n + m)` | The adjacency list stores all roads, while the visited array and BFS queue can store up to `n` cities.     |

Here:

* `n` is the number of cities.
* `m` is the number of roads.

This is the best practical complexity for this graph traversal solution because the algorithm needs to inspect the connected cities and their roads.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minScore(int n, vector<vector<int>>& roads) {
        // graph[city] stores {neighbor, road distance}.
        vector<vector<pair<int, int>>> graph(n + 1);

        // Build the adjacency list.
        // Each road is stored in both directions because the graph is undirected.
        for (const auto& road : roads) {
            int a = road[0];
            int b = road[1];
            int distance = road[2];

            graph[a].push_back({b, distance});
            graph[b].push_back({a, distance});
        }

        // visited prevents processing the same city again and again.
        vector<bool> visited(n + 1, false);

        // I use a queue to perform BFS from city 1.
        queue<int> q;
        q.push(1);
        visited[1] = true;

        // Start with the largest possible integer value.
        int answer = INT_MAX;

        // Visit every city connected to city 1.
        while (!q.empty()) {
            int city = q.front();
            q.pop();

            // Check every road leaving the current city.
            for (const auto& edge : graph[city]) {
                int nextCity = edge.first;
                int distance = edge.second;

                // Every edge in this component can be part of a valid path.
                answer = min(answer, distance);

                // Add an unvisited city so its roads are also checked.
                if (!visited[nextCity]) {
                    visited[nextCity] = true;
                    q.push(nextCity);
                }
            }
        }

        // This is the smallest road distance in city 1's component.
        return answer;
    }
};
```

### Java

```java
class Solution {
    public int minScore(int n, int[][] roads) {
        // graph[city] stores arrays of {neighbor, road distance}.
        List<int[]>[] graph = new ArrayList[n + 1];

        // Create an empty adjacency list for every city.
        for (int city = 1; city <= n; city++) {
            graph[city] = new ArrayList<>();
        }

        // Build the undirected graph by storing every road both ways.
        for (int[] road : roads) {
            int a = road[0];
            int b = road[1];
            int distance = road[2];

            graph[a].add(new int[]{b, distance});
            graph[b].add(new int[]{a, distance});
        }

        // visited prevents the BFS from processing one city repeatedly.
        boolean[] visited = new boolean[n + 1];

        // I use a queue to explore the whole component of city 1.
        Queue<Integer> queue = new ArrayDeque<>();
        queue.offer(1);
        visited[1] = true;

        // Start with the largest possible integer value.
        int answer = Integer.MAX_VALUE;

        // Continue until every reachable city has been processed.
        while (!queue.isEmpty()) {
            int city = queue.poll();

            // Check every road connected to the current city.
            for (int[] edge : graph[city]) {
                int nextCity = edge[0];
                int distance = edge[1];

                // Keep the smallest road found anywhere in this component.
                answer = Math.min(answer, distance);

                // Visit the neighboring city only if it is still unvisited.
                if (!visited[nextCity]) {
                    visited[nextCity] = true;
                    queue.offer(nextCity);
                }
            }
        }

        // Return the minimum road distance in city 1's component.
        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} roads
 * @return {number}
 */
var minScore = function(n, roads) {
    // graph[city] stores [neighbor, road distance].
    const graph = Array.from({ length: n + 1 }, () => []);

    // Build the undirected graph by storing each road in both directions.
    for (const [a, b, distance] of roads) {
        graph[a].push([b, distance]);
        graph[b].push([a, distance]);
    }

    // visited prevents the same city from being processed more than once.
    const visited = new Array(n + 1).fill(false);

    // I use an array as a BFS queue and move a pointer instead of shifting.
    // This keeps every queue operation efficient.
    const queue = [1];
    let front = 0;
    visited[1] = true;

    // Start with the largest possible value.
    let answer = Infinity;

    // Visit every city in the connected component containing city 1.
    while (front < queue.length) {
        const city = queue[front++];

        // Check every road connected to the current city.
        for (const [nextCity, distance] of graph[city]) {
            // Every road in this component can affect the final score.
            answer = Math.min(answer, distance);

            // Add the neighboring city only if it has not been visited.
            if (!visited[nextCity]) {
                visited[nextCity] = true;
                queue.push(nextCity);
            }
        }
    }

    // Return the smallest road distance found in the component.
    return answer;
};
```

### Python3

```python
class Solution:
    def minScore(self, n: int, roads: List[List[int]]) -> int:
        # graph[city] stores pairs of (neighbor, road distance).
        graph = [[] for _ in range(n + 1)]

        # Build the undirected graph by storing each road in both directions.
        for a, b, distance in roads:
            graph[a].append((b, distance))
            graph[b].append((a, distance))

        # visited prevents the same city from being processed repeatedly.
        visited = [False] * (n + 1)

        # I use a deque for efficient BFS operations.
        queue = deque([1])
        visited[1] = True

        # Start with an infinitely large value.
        answer = float("inf")

        # Visit every city connected to city 1.
        while queue:
            city = queue.popleft()

            # Check every road connected to the current city.
            for next_city, distance in graph[city]:
                # Keep the smallest road distance found in this component.
                answer = min(answer, distance)

                # Visit the neighboring city only if it is still unvisited.
                if not visited[next_city]:
                    visited[next_city] = True
                    queue.append(next_city)

        # Return the smallest road distance in city 1's component.
        return answer
```

### Go

```go
func minScore(n int, roads [][]int) int {
    // Edge stores the neighboring city and the road distance.
    type Edge struct {
        city     int
        distance int
    }

    // graph[city] stores all roads connected to that city.
    graph := make([][]Edge, n+1)

    // Build the undirected graph by storing every road in both directions.
    for _, road := range roads {
        a := road[0]
        b := road[1]
        distance := road[2]

        graph[a] = append(graph[a], Edge{b, distance})
        graph[b] = append(graph[b], Edge{a, distance})
    }

    // visited prevents processing the same city more than once.
    visited := make([]bool, n+1)

    // I use a slice as a BFS queue and move a front pointer.
    queue := []int{1}
    front := 0
    visited[1] = true

    // Start with the largest possible integer value.
    answer := int(^uint(0) >> 1)

    // Visit every city in the connected component containing city 1.
    for front < len(queue) {
        city := queue[front]
        front++

        // Check every road connected to the current city.
        for _, edge := range graph[city] {
            // Keep the smallest road distance found in this component.
            if edge.distance < answer {
                answer = edge.distance
            }

            // Add the neighboring city only if it has not been visited.
            if !visited[edge.city] {
                visited[edge.city] = true
                queue = append(queue, edge.city)
            }
        }
    }

    // Return the minimum road distance in city 1's component.
    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages. The main differences come from how each language represents the adjacency list and BFS queue.

### Step 1: Build the graph

I first create an adjacency list with space for cities from `1` to `n`.

The graph is usually created with size `n + 1` because the city numbers start from `1`, not `0`.

For every road:

`[a, b, distance]`

I add:

`b` with its distance to the list of `a`

and:

`a` with its distance to the list of `b`

If I stored the road in only one direction, the traversal would be incorrect because the problem says every road is bidirectional.

In C++, the graph can store pairs containing the neighboring city and road distance.

In Java, each adjacency list can store integer arrays containing the same two values.

In JavaScript, nested arrays work naturally for storing each neighbor and distance.

In Python3, tuples are a simple way to store the neighboring city and road distance.

In Go, a small custom structure can clearly represent each edge.

### Step 2: Prepare the visited array

Next, I create a visited array with `n + 1` positions.

Initially, every city is unvisited.

City `1` is marked as visited before the BFS begins.

Marking a city when it enters the queue is important. If I waited until removing it from the queue, different neighbors could add the same city multiple times.

That would not necessarily change the final answer, but it would waste time and memory.

### Step 3: Start BFS from city 1

I add city `1` to the queue.

I do not need to start from city `n` or from every city in the graph.

Only the connected component containing city `1` matters.

The problem guarantees that city `n` is reachable from city `1`, so city `n` must be inside this same component.

### Step 4: Process the current city

While the queue is not empty, I take one city from it.

Then I look at every road connected to that city.

For each road, I get two values:

* the neighboring city
* the road distance

The order in which cities are visited does not affect the answer.

BFS is only being used to make sure the entire connected component is explored.

### Step 5: Update the minimum score

For every road, I compare its distance with the current answer.

If the road has a smaller distance, I update the answer.

This check happens before or independently from the visited check.

That detail matters.

Suppose a road connects two cities that have both already been discovered. That road can still be the smallest road in the connected component.

If I ignored roads just because their neighboring city was already visited, I could miss the correct answer.

### Step 6: Visit new neighboring cities

After checking the road distance, I check whether the neighboring city has already been visited.

If it has not been visited:

1. Mark it as visited.
2. Add it to the queue.

If it was already visited, I do not add it again.

This keeps the traversal efficient and prevents repeated processing.

### Step 7: Finish the traversal

The BFS ends when the queue becomes empty.

At this point, every city connected to city `1` has been visited, and every road in that connected component has been checked.

The answer now contains the minimum road distance in the component.

That value is returned as the minimum possible score.

### Why this works in every language

The C++, Java, JavaScript, Python3, and Go solutions all follow the same graph algorithm.

The syntax and standard library tools are different, but the behavior stays the same:

* build the undirected graph
* start from city `1`
* visit the full connected component
* inspect every road distance
* return the smallest one

The correctness does not depend on BFS visiting cities in any specific order.

## Examples

### Example 1

Input:

```text
n = 4
roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
```

Expected output:

```text
5
```

The graph contains road distances `9`, `6`, `5`, and `7`.

Starting from city `1`, the BFS can reach all four cities.

The smallest road distance in this connected component is `5`.

A valid path is:

```text
1 -> 2 -> 4
```

The road distances are `9` and `5`.

The score is:

```text
min(9, 5) = 5
```

So the answer is `5`.

### Example 2

Input:

```text
n = 4
roads = [[1,2,2],[1,3,4],[3,4,7]]
```

Expected output:

```text
2
```

The smallest road distance in the connected component is `2`.

A valid path can be:

```text
1 -> 2 -> 1 -> 3 -> 4
```

The road distances are:

```text
2, 2, 4, 7
```

The score is:

```text
min(2, 2, 4, 7) = 2
```

The ability to revisit city `1` is what allows the road with distance `2` to be included.

### Example 3

Input:

```text
n = 5
roads = [[1,2,8],[2,3,3],[3,5,10],[1,4,6]]
```

Expected output:

```text
3
```

Starting from city `1`, all five cities are reachable.

The road distances in the connected component are:

```text
8, 3, 10, 6
```

The smallest distance is `3`.

Even though the road with distance `3` is not the last road before city `5`, it can still be included in a valid path from city `1` to city `5`.

So the minimum possible score is `3`.

## How to Use / Run Locally

Each solution can be tested locally after adding the required imports, the solution code, and a small driver program if needed.

LeetCode normally provides the input handling automatically, so the submitted solution only needs the required class or function.

### C++

Save the code in a file named:

```text
main.cpp
```

Compile it with:

```bash
g++ -std=c++17 main.cpp -o main
```

Run it with:

```bash
./main
```

On Windows Command Prompt, run:

```bash
main.exe
```

### Java

Save the code in:

```text
Main.java
```

Compile it with:

```bash
javac Main.java
```

Run it with:

```bash
java Main
```

Make sure the public class name matches the file name when testing locally.

### JavaScript

Save the code in:

```text
main.js
```

Run it with Node.js:

```bash
node main.js
```

Node.js must be installed on the system before running the file.

### Python3

Save the code in:

```text
main.py
```

Run it with:

```bash
python3 main.py
```

On some systems, the command may be:

```bash
python main.py
```

### Go

Save the code in:

```text
main.go
```

Run it directly with:

```bash
go run main.go
```

Or build an executable first:

```bash
go build main.go
```

Then run the generated executable.

## Notes & Optimizations

The most important observation is that this is not a traditional shortest path problem.

Dijkstra's algorithm is unnecessary because the total path distance does not matter.

A Minimum Spanning Tree approach can also lead to the answer, but it does more work than needed.

Union-Find is another possible solution. I could group all connected cities and then find the smallest road inside the component containing city `1`. However, BFS with an adjacency list is easier to understand and directly explores only the relevant connected component.

DFS also works with the same `O(n + m)` time complexity. The main drawback is recursion depth. With up to `100,000` cities, recursive DFS can cause a stack overflow in some languages.

BFS avoids that problem.

The graph may not be fully connected. This does not matter because roads outside the connected component of city `1` cannot be used in a path from city `1` to city `n`.

The algorithm should not stop as soon as city `n` is found. There may still be a smaller road elsewhere in the same connected component.

The road distance must be checked even if the neighboring city has already been visited. The visited array is only used to control city traversal.

Using an adjacency list instead of an adjacency matrix is necessary for good memory usage. An adjacency matrix would require `O(n^2)` space, which is far too large for `n = 100,000`.

For JavaScript and Go, using an array or slice with a moving front index is better than repeatedly removing the first element. Removing from the front can require shifting other elements and may make queue operations slower.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
