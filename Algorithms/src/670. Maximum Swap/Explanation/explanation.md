# Maximum Swap Solution - Step-by-Step Explanation

This README explains the **Maximum Swap** problem-solving process in different programming languages: **C++**, **Java**, **JavaScript**, **Python**, and **Go**.

## Problem Recap

Given a number, we can swap two digits at most once to maximize the number. We want to find the maximum possible number after at most one swap.

---

## C++ Code Walkthrough

### Step 1: Convert the Number to a String

- We convert the integer to a string to manipulate individual digits easily.
  
### Step 2: Track Last Occurrence of Each Digit

- Create an array to track the last occurrence of each digit from 0 to 9.

### Step 3: Traverse Each Digit

- Traverse through the number from left to right and for each digit, check if there is a larger digit later in the array.

### Step 4: Swap the Digits

- If a larger digit is found later, swap the current digit with the larger one to maximize the number.

### Step 5: Return the New Number

- Convert the modified string back to an integer and return the result.

---

## Java Code Walkthrough

### Step 1: Convert Number to Character Array

- Convert the integer number into a character array so that we can easily swap individual digits.

### Step 2: Track Last Occurrence of Digits

- Maintain an array to store the last position of each digit from 0 to 9 as they appear in the number.

### Step 3: Iterate Over the Digits

- Iterate through each digit of the number. For each digit, check for a larger digit that appears later in the array.

### Step 4: Perform the Swap

- If a larger digit is found later in the number, perform the swap to get the largest possible number.

### Step 5: Return the Final Number

- After performing the swap (if any), return the new number as the result.

---

## JavaScript Code Walkthrough

### Step 1: Convert the Number to an Array

- Convert the integer into a string and then into an array of characters (digits) for easier manipulation.

### Step 2: Record Last Occurrences of Each Digit

- Create an array to store the last occurrence of each digit (0-9) in the number.

### Step 3: Loop Through the Digits

- Iterate through each digit in the array and check if a larger digit appears later.

### Step 4: Swap and Maximize the Number

- Swap the current digit with a larger digit found later if possible to maximize the value.

### Step 5: Convert the Array Back to a Number

- After making the swap, join the array back into a string and convert it into an integer to return the final result.

---

## Python Code Walkthrough

### Step 1: Convert Number to List of Digits

- Convert the number into a list of characters (digits) to manipulate it more easily.

### Step 2: Track Last Appearance of Each Digit

- Use a dictionary to store the last index of each digit (0-9) as it appears in the number.

### Step 3: Traverse the Digits

- Traverse each digit and for each digit, check if there's a larger one appearing later in the list.

### Step 4: Swap Digits to Maximize Value

- Swap the current digit with a larger digit found later in the list to get the maximum possible number.

### Step 5: Convert Back to an Integer

- After the swap, convert the modified list of digits back into a string, then into an integer, and return the result.

---

## Go Code Walkthrough

### Step 1: Convert the Number to a Rune Slice

- Convert the integer to a string and then into a slice of runes (characters) for easy manipulation.

### Step 2: Record Last Occurrences of Each Digit

- Maintain an array to store the last occurrence of each digit from 0 to 9 as they appear in the number.

### Step 3: Loop Through Each Digit

- Iterate through the digits and check if there is a larger digit available later in the slice.

### Step 4: Perform the Swap

- Swap the current digit with the largest one found later in the slice to maximize the number.

### Step 5: Convert the Rune Slice Back to a Number

- After performing the swap, convert the rune slice back into a string, then convert it into an integer and return the result.

---

Each language follows a similar structure with slight variations in how the number is represented and manipulated. The key steps remain the same:

1. **Convert the number** to a manipulable format (string, array, or list).
2. **Track the last occurrence** of each digit to know where swaps are possible.
3. **Iterate through digits** and find an opportunity to swap to maximize the value.
4. **Perform the swap** and return the maximized number.
