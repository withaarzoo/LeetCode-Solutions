# LeetCode Problem 1701: Average Waiting Time - Step-by-Step Solutions

## Problem Description

Given an array `customers` where `customers[i] = [arrivali, timei]`:

- `arrivali` is the arrival time of the ith customer.
- `timei` is the time needed to prepare the order of the ith customer.

The chef starts preparing the order as soon as the previous order is finished. The goal is to calculate the average waiting time for all customers.

## C++ Solution

### Step-by-Step Explanation

1. **Initialize Variables**:
   - `total_waiting_time` to accumulate the total waiting time.
   - `current_time` to track when the chef is available to start the next order.

    ```cpp
    long long total_waiting_time = 0;
    int current_time = 0;
    ```

2. **Iterate Through Each Customer**:
   - Extract `arrival_time` and `cooking_time` for each customer.
   - Update `current_time` to the maximum of the current time or the customer's arrival time plus the cooking time.
   - Calculate the waiting time for the current customer and add it to `total_waiting_time`.

    ```cpp
    for (const auto& customer : customers) {
        int arrival_time = customer[0];
        int cooking_time = customer[1];
        
        current_time = max(current_time, arrival_time) + cooking_time;
        total_waiting_time += current_time - arrival_time;
    }
    ```

3. **Calculate and Return Average Waiting Time**:
   - Divide `total_waiting_time` by the number of customers to get the average waiting time.

    ```cpp
    return (double)total_waiting_time / customers.size();
    ```

### Complete Code

```cpp
class Solution {
public:
    double averageWaitingTime(vector<vector<int>>& customers) {
        long long total_waiting_time = 0;
        int current_time = 0;
        
        for (const auto& customer : customers) {
            int arrival_time = customer[0];
            int cooking_time = customer[1];
            
            current_time = max(current_time, arrival_time) + cooking_time;
            total_waiting_time += current_time - arrival_time;
        }
        
        return (double)total_waiting_time / customers.size();
    }
};
```

## Java Solution

### Step-by-Step Explanation

1. **Initialize Variables**:
   - `totalWaitingTime` to accumulate the total waiting time.
   - `currentTime` to track when the chef is available to start the next order.

    ```java
    long totalWaitingTime = 0;
    int currentTime = 0;
    ```

2. **Iterate Through Each Customer**:
   - Extract `arrivalTime` and `cookingTime` for each customer.
   - Update `currentTime` to the maximum of the current time or the customer's arrival time plus the cooking time.
   - Calculate the waiting time for the current customer and add it to `totalWaitingTime`.

    ```java
    for (int[] customer : customers) {
        int arrivalTime = customer[0];
        int cookingTime = customer[1];
        
        currentTime = Math.max(currentTime, arrivalTime) + cookingTime;
        totalWaitingTime += currentTime - arrivalTime;
    }
    ```

3. **Calculate and Return Average Waiting Time**:
   - Divide `totalWaitingTime` by the number of customers to get the average waiting time.

    ```java
    return (double) totalWaitingTime / customers.length;
    ```

### Complete Code

```java
class Solution {
    public double averageWaitingTime(int[][] customers) {
        long totalWaitingTime = 0;
        int currentTime = 0;
        
        for (int[] customer : customers) {
            int arrivalTime = customer[0];
            int cookingTime = customer[1];
            
            currentTime = Math.max(currentTime, arrivalTime) + cookingTime;
            totalWaitingTime += currentTime - arrivalTime;
        }
        
        return (double) totalWaitingTime / customers.length;
    }
}
```

## JavaScript Solution

### Step-by-Step Explanation

1. **Initialize Variables**:
   - `totalWaitingTime` to accumulate the total waiting time.
   - `currentTime` to track when the chef is available to start the next order.

    ```javascript
    let totalWaitingTime = 0;
    let currentTime = 0;
    ```

2. **Iterate Through Each Customer**:
   - Extract `arrivalTime` and `cookingTime` for each customer.
   - Update `currentTime` to the maximum of the current time or the customer's arrival time plus the cooking time.
   - Calculate the waiting time for the current customer and add it to `totalWaitingTime`.

    ```javascript
    for (const customer of customers) {
        let arrivalTime = customer[0];
        let cookingTime = customer[1];
        
        currentTime = Math.max(currentTime, arrivalTime) + cookingTime;
        totalWaitingTime += currentTime - arrivalTime;
    }
    ```

3. **Calculate and Return Average Waiting Time**:
   - Divide `totalWaitingTime` by the number of customers to get the average waiting time.

    ```javascript
    return totalWaitingTime / customers.length;
    ```

### Complete Code

```javascript
var averageWaitingTime = function(customers) {
    let totalWaitingTime = 0;
    let currentTime = 0;
    
    for (const customer of customers) {
        let arrivalTime = customer[0];
        let cookingTime = customer[1];
        
        currentTime = Math.max(currentTime, arrivalTime) + cookingTime;
        totalWaitingTime += currentTime - arrivalTime;
    }
    
    return totalWaitingTime / customers.length;
};
```

## Python Solution

### Step-by-Step Explanation

1. **Initialize Variables**:
   - `total_waiting_time` to accumulate the total waiting time.
   - `current_time` to track when the chef is available to start the next order.

    ```python
    total_waiting_time = 0
    current_time = 0
    ```

2. **Iterate Through Each Customer**:
   - Extract `arrival_time` and `cooking_time` for each customer.
   - Update `current_time` to the maximum of the current time or the customer's arrival time plus the cooking time.
   - Calculate the waiting time for the current customer and add it to `total_waiting_time`.

    ```python
    for arrival_time, cooking_time in customers:
        current_time = max(current_time, arrival_time) + cooking_time
        total_waiting_time += current_time - arrival_time
    ```

3. **Calculate and Return Average Waiting Time**:
   - Divide `total_waiting_time` by the number of customers to get the average waiting time.

    ```python
    return total_waiting_time / len(customers)
    ```

### Complete Code

```python
class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        total_waiting_time = 0
        current_time = 0
        
        for arrival_time, cooking_time in customers:
            current_time = max(current_time, arrival_time) + cooking_time
            total_waiting_time += current_time - arrival_time
        
        return total_waiting_time / len(customers)
```

## Go Solution

### Step-by-Step Explanation

1. **Initialize Variables**:
   - `totalWaitingTime` to accumulate the total waiting time.
   - `currentTime` to track when the chef is available to start the next order.

    ```go
    totalWaitingTime := 0
    currentTime := 0
    ```

2. **Iterate Through Each Customer**:
   - Extract `arrivalTime` and `cookingTime` for each customer.
   - Update `currentTime` to the maximum of the current time or the customer's arrival time plus the cooking time.
   - Calculate the waiting time for the current customer and add it to `totalWaitingTime`.

    ```go
    for _, customer := range customers {
        arrivalTime := customer[0]
        cookingTime := customer[1]
        
        currentTime = max(currentTime, arrivalTime) + cookingTime
        totalWaitingTime += currentTime - arrivalTime
    }
    ```

3. **Calculate and Return Average Waiting Time**:
   - Divide `totalWaitingTime` by the number of customers to get the average waiting time.

    ```go
    return float64(totalWaitingTime) / float64(len(customers))
    ```

### Complete Code

```go
func averageWaitingTime(customers [][]int) float64 {
    totalWaitingTime := 0
    currentTime := 0
    
    for _, customer := range customers {
        arrivalTime := customer[0]
        cookingTime := customer[1]
        
        currentTime = max(currentTime, arrivalTime) + cookingTime
        totalWaitingTime += currentTime - arrivalTime
    }
    
    return float64(totalWaitingTime) / float64(len(customers))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

---

These solutions should help you understand the step-by-step process of solving the LeetCode problem 1701 "Average Waiting Time" in multiple programming languages. Each code snippet is accompanied by a detailed explanation to guide you through the logic and implementation.
