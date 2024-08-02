# Minimum Swaps to Group All 1's Together

This document provides a step-by-step explanation of how to solve the problem of finding the minimum number of swaps required to group all 1's together in an array. The explanation is given in C++, Java, JavaScript, Python, and Go.

## Steps

1. **Count Total Number of 1's in the Array**
    - Calculate the total number of 1's in the array. If there are no 1's, return 0 as no swaps are needed.

2. **Initialize Variables**
    - Determine the size of the array.
    - Initialize variables to keep track of the maximum number of 1's in any window (`maxOnesInWindow`) and the current number of 1's in the current window (`currentOnesInWindow`).

3. **Calculate Number of 1's in the Initial Window**
    - Initialize the first window of size equal to the total number of 1's and count the number of 1's in this window.
    - Set `maxOnesInWindow` to the count of 1's in the initial window.

4. **Sliding Window Technique**
    - Slide the window across the array to find the maximum number of 1's in any window of size equal to the total number of 1's. This involves:
        - Subtracting the element that is sliding out of the window.
        - Adding the new element that is sliding into the window.
        - Updating `maxOnesInWindow` if the current window has more 1's.

5. **Calculate Minimum Swaps**
    - The minimum number of swaps needed is the total number of 1's minus the maximum number of 1's found in any window. This gives the number of 0's that need to be swapped.

## Code Explanations

### C++

1. **Include Necessary Libraries**
    - Use `<vector>` for dynamic array management and `<algorithm>` for counting elements.

2. **Class Definition**
    - Define a class `Solution` with a public method `minSwaps` that takes a reference to a vector of integers.

3. **Count Total 1's**
    - Use `std::count` to count the number of 1's in the array.

4. **Handle Edge Case**
    - Return 0 if there are no 1's in the array.

5. **Initialize Variables**
    - Define and initialize variables for the array size, `maxOnesInWindow`, and `currentOnesInWindow`.

6. **Calculate Initial Window**
    - Loop through the first `totalOnes` elements to count the number of 1's.

7. **Sliding Window**
    - Use a for loop to slide the window across the array, adjusting `currentOnesInWindow` and updating `maxOnesInWindow`.

8. **Return Result**
    - Calculate and return the minimum number of swaps needed.

### Java

1. **Class Definition**
    - Define a class `Solution` with a public method `minSwaps` that takes an integer array.

2. **Count Total 1's**
    - Use a for-each loop to count the number of 1's in the array.

3. **Handle Edge Case**
    - Return 0 if there are no 1's in the array.

4. **Initialize Variables**
    - Define and initialize variables for the array length, `maxOnesInWindow`, and `currentOnesInWindow`.

5. **Calculate Initial Window**
    - Loop through the first `totalOnes` elements to count the number of 1's.

6. **Sliding Window**
    - Use a for loop to slide the window across the array, adjusting `currentOnesInWindow` and updating `maxOnesInWindow`.

7. **Return Result**
    - Calculate and return the minimum number of swaps needed.

### JavaScript

1. **Function Definition**
    - Define a function `minSwaps` that takes an array of integers.

2. **Count Total 1's**
    - Use `reduce` to count the number of 1's in the array.

3. **Handle Edge Case**
    - Return 0 if there are no 1's in the array.

4. **Initialize Variables**
    - Define and initialize variables for the array length, `maxOnesInWindow`, and `currentOnesInWindow`.

5. **Calculate Initial Window**
    - Loop through the first `totalOnes` elements to count the number of 1's.

6. **Sliding Window**
    - Use a for loop to slide the window across the array, adjusting `currentOnesInWindow` and updating `maxOnesInWindow`.

7. **Return Result**
    - Calculate and return the minimum number of swaps needed.

### Python

1. **Class Definition**
    - Define a class `Solution` with a method `minSwaps` that takes a list of integers.

2. **Count Total 1's**
    - Use `sum` to count the number of 1's in the array.

3. **Handle Edge Case**
    - Return 0 if there are no 1's in the array.

4. **Initialize Variables**
    - Define and initialize variables for the array length, `maxOnesInWindow`, and `currentOnesInWindow`.

5. **Calculate Initial Window**
    - Loop through the first `totalOnes` elements to count the number of 1's.

6. **Sliding Window**
    - Use a for loop to slide the window across the array, adjusting `currentOnesInWindow` and updating `maxOnesInWindow`.

7. **Return Result**
    - Calculate and return the minimum number of swaps needed.

### Go

1. **Function Definition**
    - Define a function `minSwaps` that takes a slice of integers.

2. **Count Total 1's**
    - Use a for loop to count the number of 1's in the array.

3. **Handle Edge Case**
    - Return 0 if there are no 1's in the array.

4. **Initialize Variables**
    - Define and initialize variables for the array length, `maxOnesInWindow`, and `currentOnesInWindow`.

5. **Calculate Initial Window**
    - Loop through the first `totalOnes` elements to count the number of 1's.

6. **Sliding Window**
    - Use a for loop to slide the window across the array, adjusting `currentOnesInWindow` and updating `maxOnesInWindow`.

7. **Return Result**
    - Calculate and return the minimum number of swaps needed.
