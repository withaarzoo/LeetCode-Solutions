# Stone Removal Game - LeetCode 3360 Solution

## Table of Contents
- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

Alice and Bob play a simple stone removal game. There is a pile of n stones. Alice always starts first and must remove exactly 10 stones on her first turn. After that, every player must remove exactly one fewer stone than the previous opponent. The player who cannot make a valid move loses the game.

You are given a positive integer n. Your task is to return true if Alice wins the game and false if Bob wins, assuming both play optimally according to the fixed rules.

This is a classic simulation-style DSA problem that appears in coding interviews and online judges. Understanding the turn-based removal pattern is the key to solving the Stone Removal Game efficiently.

## Constraints

- 1 <= n <= 50

The constraint is very small, which means a straightforward simulation approach is both correct and fast enough.

## Intuition

The first thing I noticed is that Alice’s opening move is fixed at 10 stones. After that the required number of stones simply decreases by 1 on every turn. Because n never goes beyond 50, the longest possible game is just the sequence 10 + 9 + 8 + … + 1, which is only 10 moves. That makes a direct simulation the most natural and reliable way to decide the winner.

There is no need for complicated game theory or dynamic programming. We can simply keep subtracting the current required amount and switch players until someone cannot move.

## Approach

I start with the required removal amount set to 10 and a boolean flag that tracks whose turn it is (Alice starts as true).

While the remaining stones are greater than or equal to the current required amount, I subtract that amount from n, decrease the required amount by one, and flip the turn flag.

When the loop ends, the player whose turn it is cannot make the move and therefore loses. If it is still Alice’s turn she loses (return false). If it is Bob’s turn then Alice wins (return true).

This simple loop correctly handles every possible value of n within the given constraints.

## Data Structures Used

No special data structures are required. The solution only uses a few primitive variables:

- An integer to store the current number of stones that must be removed
- A boolean flag to track whose turn it is
- The input integer n itself

Keeping the solution free of arrays, maps or recursion makes it both fast and easy to understand for beginners.

## Operations & Behavior Summary

1. Initialize the removal amount to 10 and set Alice’s turn to true.
2. Check whether the remaining stones are enough for the current required amount.
3. If yes, subtract the amount, reduce the required amount by 1, and switch the player.
4. Repeat until a player faces a required amount larger than the stones left.
5. The player who cannot move loses. Return the opposite of the current turn flag.

The entire process is a pure simulation of the game rules.

## Complexity

| Type              | Complexity | Explanation                                                                 |
|-------------------|------------|-----------------------------------------------------------------------------|
| Time Complexity   | O(1)       | The loop runs at most 10 times (the sequence from 10 down to 1).            |
| Space Complexity  | O(1)       | Only a constant number of integer and boolean variables are used.          |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    bool canAliceWin(int n) {
        int remove = 10;
        bool aliceTurn = true;
        while (n >= remove) {
            n -= remove;
            remove--;
            aliceTurn = !aliceTurn;
        }
        return !aliceTurn;
    }
};
```

### Java
```java
class Solution {
    public boolean canAliceWin(int n) {
        int remove = 10;
        boolean aliceTurn = true;
        while (n >= remove) {
            n -= remove;
            remove--;
            aliceTurn = !aliceTurn;
        }
        return !aliceTurn;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} n
 * @return {boolean}
 */
var canAliceWin = function(n) {
    let remove = 10;
    let aliceTurn = true;
    while (n >= remove) {
        n -= remove;
        remove--;
        aliceTurn = !aliceTurn;
    }
    return !aliceTurn;
};
```

### TypeScript
```typescript
function canAliceWin(n: number): boolean {
    let remove = 10;
    let aliceTurn = true;
    while (n >= remove) {
        n -= remove;
        remove--;
        aliceTurn = !aliceTurn;
    }
    return !aliceTurn;
};
```

### Python3
```python
class Solution:
    def canAliceWin(self, n: int) -> bool:
        remove = 10
        alice_turn = True
        while n >= remove:
            n -= remove
            remove -= 1
            alice_turn = not alice_turn
        return not alice_turn
```

### Go
```go
func canAliceWin(n int) bool {
    remove := 10
    aliceTurn := true
    for n >= remove {
        n -= remove
        remove--
        aliceTurn = !aliceTurn
    }
    return !aliceTurn
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

All six language versions follow exactly the same logic. The only differences are syntax and naming conventions.

We begin by setting the first required removal to 10 because the problem statement forces Alice to take exactly 10 stones on her opening move. A boolean flag (or equivalent) starts as true to indicate it is Alice’s turn.

The main loop continues as long as the remaining stones are greater than or equal to the current required amount. Inside the loop three actions always happen in the same order:

- Subtract the required amount from the remaining stones.
- Decrease the required amount by one so the next player has a smaller target.
- Flip the turn flag so the other player becomes active.

When the loop condition fails, the current player cannot take the required number of stones and therefore loses. Returning the negation of the turn flag gives the correct answer: false when Alice cannot move, true when Bob cannot move.

Edge cases are handled naturally. If n is less than 10 at the very beginning, the loop never runs and Alice loses immediately. If the stones run out exactly on a player’s move, that player still loses because the next required amount will be larger than zero remaining stones.

Because the maximum number of iterations is fixed and tiny, the same code works efficiently in every language without any further tuning.

## Examples

Example 1  
Input: n = 12  
Output: true  

Alice removes 10 stones, leaving 2. Bob is required to remove 9 stones but only 2 remain, so Bob loses and Alice wins.

Example 2  
Input: n = 1  
Output: false  

Alice is required to remove 10 stones but only 1 is available, so Alice loses immediately.

Example 3  
Input: n = 19  
Trace:  
- Alice takes 10 → 9 left  
- Bob takes 9 → 0 left  
- Alice would need 8 but 0 remain → Alice loses  
Output: false

## How to Use / Run Locally

C++  
1. Save the code in a file named main.cpp  
2. Compile: g++ -std=c++17 main.cpp -o main  
3. Run: ./main  

Java  
1. Save the code in Solution.java  
2. Compile: javac Solution.java  
3. Run with a small driver that creates a Solution object and calls canAliceWin  

JavaScript  
1. Save the code in solution.js  
2. Run with Node: node solution.js  

TypeScript  
1. Save the code in solution.ts  
2. Compile: tsc solution.ts  
3. Run the generated JavaScript file with Node  

Python3  
1. Save the code in solution.py  
2. Run: python3 solution.py  

Go  
1. Save the code in main.go  
2. Run: go run main.go  

In every language you can add a simple main or test function that prints the result for a few sample values of n to verify the solution.

## Notes & Optimizations

Because the constraint on n is only 50, the simulation is already optimal. There is no practical benefit from trying to find a closed-form mathematical formula, although one could be derived by calculating the triangular numbers.

The same loop works for any larger n as long as the removal sequence stays the same. If the problem ever changed the starting removal amount or allowed variable moves, a different approach would be needed.

Watch out for the exact edge case when n equals the exact sum of the first k terms of the sequence; the player who would have to move next loses.

## Author
[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
