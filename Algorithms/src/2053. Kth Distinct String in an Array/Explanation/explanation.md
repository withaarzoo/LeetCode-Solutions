# Finding the K-th Distinct String in an Array

This README provides a step-by-step explanation of how to find the k-th distinct string in an array using different programming languages: C++, Java, JavaScript, Python, and Go.

## C++ Code

### Step-by-Step Explanation

1. **Create an unordered map** to count the occurrences of each string in the array.
2. **Create a vector** to store distinct strings.
3. **Count occurrences** of each string by iterating through the array.
4. **Collect distinct strings** by iterating through the array again and checking if the count of each string is exactly one.
5. **Check if k** is within the range of distinct strings.
6. **Return the k-th distinct string** if it exists, otherwise return an empty string.

## Java Code

### Step-by-Step Explanation

1. **Create a HashMap** to count the occurrences of each string.
2. **Create a List** to store distinct strings.
3. **Loop through the array** to count occurrences of each string, using the `getOrDefault` method to handle new strings.
4. **Loop through the array again** to collect distinct strings by checking if the count of each string is exactly one.
5. **Check if the k-th distinct string** exists by comparing k with the size of the distinct list.
6. **Return the k-th distinct string** if it exists, otherwise return an empty string.

## JavaScript Code

### Step-by-Step Explanation

1. **Create a Map** to count the occurrences of each string.
2. **Create an array** to store distinct strings.
3. **Count occurrences** of each string by iterating through the array and updating the Map.
4. **Collect distinct strings** by iterating through the array again and checking if the count of each string is exactly one.
5. **Check if k** is within the range of distinct strings.
6. **Return the k-th distinct string** if it exists, otherwise return an empty string.

## Python Code

### Step-by-Step Explanation

1. **Create a dictionary** to store the frequency of each string.
2. **Create a list** to store distinct strings.
3. **Iterate through the array** to count occurrences of each string, using the `get` method to handle new strings.
4. **Iterate through the array again** to collect distinct strings by checking if the count of each string is exactly one.
5. **Check if the k-th distinct string** exists by comparing k with the length of the distinct list.
6. **Return the k-th distinct string** if it exists, otherwise return an empty string.

## Go Code

### Step-by-Step Explanation

1. **Create a map** to count the occurrences of each string.
2. **Create a slice** to store distinct strings.
3. **Loop through the array** to count occurrences of each string.
4. **Loop through the array again** to collect distinct strings by checking if the count of each string is exactly one.
5. **Check if k** is within the range of distinct strings.
6. **Return the k-th distinct string** if it exists, otherwise return an empty string.
