# Josephus Problem Solution

The Josephus problem is a classic mathematical problem that can be solved recursively. The problem is as follows:

Given `n` people standing in a circle and a counting number `k`, the first person is killed, then every `k`-th person is killed in clockwise order. This continues until only one person remains, who is the winner. The goal is to find the position of the winner.

Here's a step-by-step explanation of the provided code solutions in various programming languages:

## C++ Solution

```cpp
class Solution {
public:
    int findTheWinner(int n, int k) {
        return josephus(n, k) + 1; // +1 to convert 0-indexed to 1-indexed
    }
    
    int josephus(int n, int k) {
        if (n == 1) {
            return 0; // Base case: only one person left
        }
        return (josephus(n - 1, k) + k) % n; // Recursive case
    }
};
```

1. The `findTheWinner` function takes the number of people `n` and the counting number `k` as input.
2. It calls the `josephus` function, which implements the recursive solution to the Josephus problem.
3. The `josephus` function has a base case where `n == 1`, which means there is only one person left. In this case, it returns `0` (the 0-indexed position of the winner).
4. In the recursive case, the function calculates the position of the winner for `n-1` people and adds `k` to it. The result is then taken modulo `n` to get the position of the winner in the current circle of `n` people.
5. The `findTheWinner` function adds `1` to the result of `josephus` to convert the 0-indexed position to a 1-indexed position, as required by the problem statement.

## Java Solution

```java
class Solution {
    public int findTheWinner(int n, int k) {
        return josephus(n, k) + 1; // +1 to convert 0-indexed to 1-indexed
    }
    
    private int josephus(int n, int k) {
        if (n == 1) {
            return 0; // Base case: only one person left
        }
        return (josephus(n - 1, k) + k) % n; // Recursive case
    }
}
```

The Java solution is almost identical to the C++ solution, with the following differences:

1. The `josephus` function is marked as `private` to encapsulate the recursive implementation.
2. The method names and class structure are adapted to Java conventions.

## JavaScript Solution

```javascript
var findTheWinner = function(n, k) {
    return josephus(n, k) + 1; // +1 to convert 0-indexed to 1-indexed
};

function josephus(n, k) {
    if (n === 1) {
        return 0; // Base case: only one person left
    }
    return (josephus(n - 1, k) + k) % n; // Recursive case
}
```

The JavaScript solution is also very similar to the C++ and Java solutions, with the following differences:

1. The `findTheWinner` function is defined as a standalone function expression.
2. The `josephus` function is defined as a standalone function declaration.
3. The comparison for the base case uses `===` instead of `==`.

## Python Solution

```python
class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        return self.josephus(n, k) + 1 # +1 to convert 0-indexed to 1-indexed
    
    def josephus(self, n: int, k: int) -> int:
        if n == 1:
            return 0 # Base case: only one person left
        return (self.josephus(n - 1, k) + k) % n # Recursive case
```

The Python solution is also very similar to the previous solutions, with the following differences:

1. The solution is implemented as a class with two methods: `findTheWinner` and `josephus`.
2. The method signatures use type annotations for the input parameters and return values.
3. The comparison for the base case uses `==` instead of `===`.

## Go Solution

```go
func findTheWinner(n int, k int) int {
    return josephus(n, k) + 1 // +1 to convert 0-indexed to 1-indexed
}

func josephus(n int, k int) int {
    if n == 1 {
        return 0 // Base case: only one person left
    }
    return (josephus(n-1, k) + k) % n // Recursive case
}
```

The Go solution is very similar to the previous solutions, with the following differences:

1. The `findTheWinner` and `josephus` functions are defined as standalone functions, not as part of a class or object.
2. The function signatures do not use type annotations, as Go can infer the types from the function parameters.

All the solutions follow the same recursive logic to solve the Josephus problem. The base case is when there is only one person left, and the recursive case calculates the position of the winner for the current circle of `n` people based on the position of the winner for the previous circle of `n-1` people.

The final result is adjusted by adding `1` to convert the 0-indexed position to a 1-indexed position, as required by the problem statement.
