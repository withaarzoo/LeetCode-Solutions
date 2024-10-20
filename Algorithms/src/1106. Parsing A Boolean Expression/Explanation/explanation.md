# Parsing a Boolean Expression - Step by Step Explanation

This repository contains multiple implementations of the solution for the LeetCode problem **Parsing a Boolean Expression**. The problem requires parsing and evaluating a Boolean expression based on the given rules for logical operators `AND`, `OR`, and `NOT`. Below is a step-by-step explanation for how the solution is designed and implemented in each language.

## Problem Recap

We are given a boolean expression that evaluates either to `true` or `false`. The expression can be:

1. `t` which evaluates to `true`.
2. `f` which evaluates to `false`.
3. `!(subExpr)` which evaluates to the logical NOT of the inner expression `subExpr`.
4. `&(subExpr1, subExpr2, ..., subExprn)` which evaluates to the logical AND of the inner expressions.
5. `|(subExpr1, subExpr2, ..., subExprn)` which evaluates to the logical OR of the inner expressions.

Our goal is to **evaluate the expression** and return the result as either `true` or `false`.

---

## Approach Overview

1. **Stack-based Solution**:
    - We traverse the expression character by character.
    - When we encounter a closing parenthesis `)`, it signals the end of an inner expression. We pop all the elements inside the parentheses, apply the logical operator (`!`, `&`, or `|`), and push the result back onto the stack.
    - We continue this process until we finish processing the entire string.

2. **Operators and Logic**:
    - `!` (NOT): Inverts the truth value of the sub-expression.
    - `&` (AND): Returns `true` if all the sub-expressions are `true`, otherwise returns `false`.
    - `|` (OR): Returns `true` if at least one sub-expression is `true`.

---

## Step-by-Step Walkthrough for Each Language

---

### C++ Code

1. **Initialize the Stack**:
    - A stack is used to keep track of the characters in the expression.
    - We process each character from the input expression string one by one.

2. **Process Characters**:
    - If a closing parenthesis `)` is encountered, start collecting the sub-expressions within the parentheses until the opening parenthesis `(` is found.
    - Determine the operator (`!`, `&`, or `|`) immediately before the parentheses.

3. **Evaluate Sub-expressions**:
    - For `!`: Apply the NOT operation on the single sub-expression.
    - For `&`: Apply the AND operation across all sub-expressions.
    - For `|`: Apply the OR operation across all sub-expressions.

4. **Push Results Back to Stack**:
    - After evaluating, push the result (`t` or `f`) back onto the stack.

5. **Final Result**:
    - After the entire expression is processed, the top of the stack contains the final result, which is either `true` or `false`.

---

### Java Code

1. **Stack Initialization**:
    - A stack is initialized to store characters during the traversal of the input expression.

2. **Handling Parentheses**:
    - When a closing parenthesis `)` is encountered, we collect the characters from the stack until the corresponding opening parenthesis `(` is found.
    - The operator that precedes the opening parenthesis determines how the sub-expressions should be evaluated.

3. **Logical Operations**:
    - For `!` (NOT), we invert the truth value of the sub-expression.
    - For `&` (AND), we check if all sub-expressions evaluate to `true`; otherwise, the result is `false`.
    - For `|` (OR), we return `true` if at least one sub-expression evaluates to `true`.

4. **Final Push**:
    - After evaluating the sub-expressions, we push the result back onto the stack for further use.

5. **Returning the Result**:
    - Once the entire expression is processed, the top of the stack gives the final result of the Boolean expression.

---

### JavaScript Code

1. **Use of Stack**:
    - A stack is used to track the characters and sub-expressions within the Boolean expression as we iterate through it.

2. **Detecting Sub-expressions**:
    - When encountering a `)`, gather all the sub-expressions inside the parentheses and determine the operator (`!`, `&`, or `|`).

3. **Evaluating the Expression**:
    - Apply the operator to the collected sub-expressions:
        - For `!`, invert the truth value.
        - For `&`, ensure all sub-expressions are `true` to result in `true`.
        - For `|`, check if any sub-expression is `true`.

4. **Update Stack**:
    - Push the evaluation result back to the stack.

5. **Return Final Boolean**:
    - After traversing the entire input, the Boolean result is at the top of the stack and is returned.

---

### Python Code

1. **Stack for Processing**:
    - Initialize a stack to manage the characters of the expression.
    - Each character is processed sequentially.

2. **Sub-expression Handling**:
    - Upon encountering `)`, gather the sub-expressions within the parentheses and then evaluate them based on the operator right before `(`.

3. **Boolean Logic Application**:
    - Apply the corresponding operator:
        - For `!`, invert the value.
        - For `&`, check that all sub-expressions are `true`.
        - For `|`, return `true` if at least one sub-expression is `true`.

4. **Push Result**:
    - After evaluation, push the resulting value back to the stack.

5. **Final Result**:
    - At the end of the iteration, the stack will contain the final result of the Boolean expression, which is returned.

---

### Go Code

1. **Stack Initialization**:
    - Use a stack to keep track of characters and sub-expressions as we parse through the Boolean expression.

2. **Processing Characters**:
    - When a `)` is found, pop elements from the stack until we reach the corresponding `(`.
    - Use the operator found before `(` to evaluate the sub-expressions.

3. **Evaluate Expression**:
    - Based on the operator:
        - For `!`, apply NOT on the single sub-expression.
        - For `&`, apply AND across all sub-expressions.
        - For `|`, apply OR across all sub-expressions.

4. **Push Back Result**:
    - After evaluating, push the result (`t` or `f`) back to the stack.

5. **Final Output**:
    - The result of the Boolean expression is found at the top of the stack at the end of the process.

---

## Conclusion

In all the implementations, the key steps involve:

1. **Stack-based parsing**: Using a stack to manage characters and sub-expressions.
2. **Operator handling**: Determining the logical operation to apply (`!`, `&`, `|`).
3. **Sub-expression evaluation**: Processing sub-expressions within parentheses and pushing the results back to the stack.
4. **Final result**: The top of the stack contains the Boolean result after processing the entire expression.

---

### Time Complexity

- **Time Complexity**: $$O(n)$$, where \(n\) is the length of the expression. Each character in the string is processed once, and the sub-expressions are evaluated linearly.

### Space Complexity

- **Space Complexity**: $$O(n)$$, as we use a stack to store intermediate characters and results during the parsing process.

---

By following this step-by-step guide, you will understand how the Boolean expression is parsed and evaluated using a stack-based approach across multiple programming languages.
