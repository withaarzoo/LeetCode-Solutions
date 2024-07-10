# Crawler Log Folder Solutions

This repository contains solutions to the "Crawler Log Folder" problem from LeetCode in multiple programming languages. Each solution is explained step-by-step to help you understand the logic and implementation.

## Table of Contents

- [C++ Solution](#c-solution)
- [Java Solution](#java-solution)
- [JavaScript Solution](#javascript-solution)
- [Python Solution](#python-solution)
- [Go Solution](#go-solution)

---

## C++ Solution

### Step-by-Step Explanation

1. **Initialization**:
    - We start by initializing a variable `depth` to `0` to represent the main folder.

    ```cpp
    int depth = 0;
    ```

2. **Processing Each Log Entry**:
    - We loop through each string in the `logs` vector.

    ```cpp
    for (const string& log : logs) {
    ```

3. **Handling `"../"`**:
    - If the log entry is `"../"`, we move to the parent folder. If `depth` is greater than `0`, we decrement it.

    ```cpp
    if (log == "../") {
        if (depth > 0) depth--;
    ```

4. **Handling `"./"`**:
    - If the log entry is `"./"`, we do nothing and stay in the current folder.

    ```cpp
    } else if (log != "./") {
    ```

5. **Handling Other Folders**:
    - For any other folder move (e.g., `"x/"`), we increment the `depth`.

    ```cpp
    depth++;
    ```

6. **Returning the Result**:
    - Finally, we return the `depth`, which represents the minimum number of operations needed to go back to the main folder.

    ```cpp
    return depth;
    ```

### Full C++ Code

```cpp
class Solution {
public:
    int minOperations(vector<string>& logs) {
        int depth = 0;
        for (const string& log : logs) {
            if (log == "../") {
                if (depth > 0) depth--;
            } else if (log != "./") {
                depth++;
            }
        }
        return depth;
    }
};
```

---

## Java Solution

### Step-by-Step Explanation

1. **Initialization**:
    - We start by initializing a variable `depth` to `0` to represent the main folder.

    ```java
    int depth = 0;
    ```

2. **Processing Each Log Entry**:
    - We loop through each string in the `logs` array.

    ```java
    for (String log : logs) {
    ```

3. **Handling `"../"`**:
    - If the log entry is `"../"`, we move to the parent folder. If `depth` is greater than `0`, we decrement it.

    ```java
    if (log.equals("../")) {
        if (depth > 0) depth--;
    ```

4. **Handling `"./"`**:
    - If the log entry is `"./"`, we do nothing and stay in the current folder.

    ```java
    } else if (!log.equals("./")) {
    ```

5. **Handling Other Folders**:
    - For any other folder move (e.g., `"x/"`), we increment the `depth`.

    ```java
    depth++;
    ```

6. **Returning the Result**:
    - Finally, we return the `depth`, which represents the minimum number of operations needed to go back to the main folder.

    ```java
    return depth;
    ```

### Full Java Code

```java
class Solution {
    public int minOperations(String[] logs) {
        int depth = 0;
        for (String log : logs) {
            if (log.equals("../")) {
                if (depth > 0) depth--;
            } else if (!log.equals("./")) {
                depth++;
            }
        }
        return depth;
    }
}
```

---

## JavaScript Solution

### Step-by-Step Explanation

1. **Initialization**:
    - We start by initializing a variable `depth` to `0` to represent the main folder.

    ```javascript
    let depth = 0;
    ```

2. **Processing Each Log Entry**:
    - We loop through each string in the `logs` array.

    ```javascript
    for (const log of logs) {
    ```

3. **Handling `"../"`**:
    - If the log entry is `"../"`, we move to the parent folder. If `depth` is greater than `0`, we decrement it.

    ```javascript
    if (log === "../") {
        if (depth > 0) depth--;
    ```

4. **Handling `"./"`**:
    - If the log entry is `"./"`, we do nothing and stay in the current folder.

    ```javascript
    } else if (log !== "./") {
    ```

5. **Handling Other Folders**:
    - For any other folder move (e.g., `"x/"`), we increment the `depth`.

    ```javascript
    depth++;
    ```

6. **Returning the Result**:
    - Finally, we return the `depth`, which represents the minimum number of operations needed to go back to the main folder.

    ```javascript
    return depth;
    ```

### Full JavaScript Code

```javascript
var minOperations = function(logs) {
    let depth = 0;
    for (const log of logs) {
        if (log === "../") {
            if (depth > 0) depth--;
        } else if (log !== "./") {
            depth++;
        }
    }
    return depth;
};
```

---

## Python Solution

### Step-by-Step Explanation

1. **Initialization**:
    - We start by initializing a variable `depth` to `0` to represent the main folder.

    ```python
    depth = 0
    ```

2. **Processing Each Log Entry**:
    - We loop through each string in the `logs` list.

    ```python
    for log in logs:
    ```

3. **Handling `"../"`**:
    - If the log entry is `"../"`, we move to the parent folder. If `depth` is greater than `0`, we decrement it.

    ```python
    if log == "../":
        if depth > 0:
            depth -= 1
    ```

4. **Handling `"./"`**:
    - If the log entry is `"./"`, we do nothing and stay in the current folder.

    ```python
    elif log != "./":
    ```

5. **Handling Other Folders**:
    - For any other folder move (e.g., `"x/"`), we increment the `depth`.

    ```python
    depth += 1
    ```

6. **Returning the Result**:
    - Finally, we return the `depth`, which represents the minimum number of operations needed to go back to the main folder.

    ```python
    return depth
    ```

### Full Python Code

```python
class Solution:
    def minOperations(self, logs: List[str]) -> int:
        depth = 0
        for log in logs:
            if log == "../":
                if depth > 0:
                    depth -= 1
            elif log != "./":
                depth += 1
        return depth
```

---

## Go Solution

### Step-by-Step Explanation

1. **Initialization**:
    - We start by initializing a variable `depth` to `0` to represent the main folder.

    ```go
    depth := 0
    ```

2. **Processing Each Log Entry**:
    - We loop through each string in the `logs` slice.

    ```go
    for _, log := range logs {
    ```

3. **Handling `"../"`**:
    - If the log entry is `"../"`, we move to the parent folder. If `depth` is greater than `0`, we decrement it.

    ```go
    if log == "../" {
        if depth > 0 {
            depth--
        }
    ```

4. **Handling `"./"`**:
    - If the log entry is `"./"`, we do nothing and stay in the current folder.

    ```go
    } else if log != "./" {
    ```

5. **Handling Other Folders**:
    - For any other folder move (e.g., `"x/"`), we increment the `depth`.

    ```go
    depth++
    ```

6. **Returning the Result**:
    - Finally, we return the `depth`, which represents the minimum number of operations needed to go back to the main folder.

    ```go
    return depth
    ```

### Full Go Code

```go
func minOperations(logs []string) int {
    depth := 0
    for _, log := range logs {
        if log == "../" {
            if depth > 0 {
                depth--
            }
        } else if log != "./" {
            depth++
        }
    }
    return depth
}
```