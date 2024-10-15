# Separate Black and White Balls Solution

In this problem, the goal is to group all black balls (represented by `1`s) to the right and all white balls (represented by `0`s) to the left, using the fewest adjacent swaps. Below is the step-by-step breakdown of how the solution is approached and implemented in various languages, including C++, Java, JavaScript, Python, and Go.

---

## C++ Solution Explanation

1. **Initialization**:
   - We start by initializing two variables: one to count the number of black balls (`1`s) encountered, and another to track the total number of swaps needed.

2. **Traverse the String**:
   - Loop through each character in the input string.
   - When we find a white ball (`0`), it will need to swap with all the black balls to its left. So, we add the number of black balls encountered so far to the total swap count.
   - When we find a black ball (`1`), increment the black ball count.

3. **Return Result**:
   - After iterating through the entire string, return the accumulated swap count as the minimum number of steps required.

---

## Java Solution Explanation

1. **Initialize Variables**:
   - We define a variable to keep track of the total swaps needed (`res`) and a counter for the number of black balls (`i`).

2. **Loop Through String**:
   - Loop through each character in the string.
   - For each white ball (`0`), calculate how far it needs to move left by subtracting the current index from the number of black balls encountered, and add this value to the total result.
   - When a black ball (`1`) is found, increment the black ball counter.

3. **Final Result**:
   - Return the total count of swaps as the answer.

---

## JavaScript Solution Explanation

1. **Set Up Variables**:
   - We start by initializing two variables: one for counting the number of black balls (`1`s), and another to accumulate the total number of swaps required.

2. **Iterate Over the String**:
   - Traverse through each character in the string.
   - If the character is a white ball (`0`), add the number of black balls encountered so far to the total number of swaps.
   - If the character is a black ball (`1`), increment the black ball counter.

3. **Output Result**:
   - Once the loop is complete, return the total number of swaps as the result.

---

## Python Solution Explanation

1. **Initialize Variables**:
   - We initialize two variables: `blackCount` to track the number of black balls, and `ans` to store the total number of swaps required.

2. **Process Each Character**:
   - Traverse the input string character by character.
   - If we find a white ball (`0`), we add the number of black balls already encountered (`blackCount`) to the total swap count.
   - If a black ball (`1`) is found, increment the black ball counter.

3. **Return the Answer**:
   - After the traversal is finished, return the total swap count stored in `ans`.

---

## Go Solution Explanation

1. **Initialize Counters**:
   - Declare two variables: one to track the number of black balls encountered (`blackCount`), and another to store the total number of swaps required (`ans`).

2. **Traverse the String**:
   - Iterate through each character in the string.
   - For each white ball (`0`), add the number of black balls encountered so far to the total number of swaps.
   - For each black ball (`1`), increment the black ball counter.

3. **Return Result**:
   - After processing all the characters, return the accumulated number of swaps as the result.

---

### Conclusion

Each solution, whether written in C++, Java, JavaScript, Python, or Go, follows the same logic:

1. **Count the number of black balls** encountered.
2. **For each white ball** encountered, add the number of black balls on its left to the total number of swaps.
3. **Return the total number of swaps** at the end.

These solutions are optimized for performance, running in linear time, \( O(n) \), where \( n \) is the length of the string. They also use constant space, \( O(1) \), making them efficient for large inputs.
