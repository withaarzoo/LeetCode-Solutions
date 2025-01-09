# Prefix Count Solution – Step-by-Step Explanation

This README provides a detailed explanation of how the **Prefix Count** problem is solved in different programming languages: C++, Java, JavaScript, Python, and Go. Each explanation follows a structured, step-by-step approach for better understanding.

---

## **C++ Code – Step-by-Step Explanation**

1. **Define the Function**: Start by creating a function `prefixCount` that takes a vector of strings (`words`) and a string (`pref`) as inputs.
2. **Initialize a Counter**: Declare an integer variable `count` and set it to 0. This will track how many words start with the prefix.
3. **Iterate Through Words**: Use a `for` loop to go through each word in the `words` vector.
4. **Extract the Prefix**: For each word, extract the substring of the first `pref.size()` characters.
5. **Compare Prefix**: Check if the extracted substring is equal to the given prefix.
6. **Update Counter**: If the prefix matches, increment the counter by 1.
7. **Return the Result**: After processing all words, return the value of the counter.

---

## **Java Code – Step-by-Step Explanation**

1. **Define the Function**: Create a method `prefixCount` that takes an array of strings (`words`) and a string (`pref`) as inputs.
2. **Initialize a Counter**: Declare an integer variable `count` and initialize it to 0. This will count matching words.
3. **Loop Through Words**: Use a `for` loop to iterate through each word in the `words` array.
4. **Check Prefix**: For each word, use the `startsWith` method to check if the word starts with the prefix `pref`.
5. **Update Counter**: If the word starts with the prefix, increment the `count` variable.
6. **Return the Count**: After iterating through all words, return the value of `count`.

---

## **JavaScript Code – Step-by-Step Explanation**

1. **Define the Function**: Create a function `prefixCount` that takes an array of strings (`words`) and a string (`pref`) as inputs.
2. **Initialize a Counter**: Declare a variable `count` and initialize it to 0. This variable will hold the number of matching words.
3. **Iterate Through Words**: Use a `for...of` loop to go through each word in the `words` array.
4. **Check Prefix**: Use the `startsWith` method to check if the current word begins with the prefix `pref`.
5. **Increment Counter**: If the word starts with the prefix, increment the `count` variable by 1.
6. **Return the Result**: After processing all the words, return the value of `count`.

---

## **Python Code – Step-by-Step Explanation**

1. **Define the Function**: Write a function `prefixCount` that takes a list of strings (`words`) and a string (`pref`) as inputs.
2. **Initialize a Counter**: Start with a variable `count` set to 0. This will count the number of words with the desired prefix.
3. **Iterate Through Words**: Use a `for` loop to go through each word in the `words` list.
4. **Check Prefix**: Use the `startswith` method to check if the current word begins with the prefix `pref`.
5. **Increment Counter**: If the word matches the prefix, increase the `count` by 1.
6. **Return the Count**: After checking all the words, return the value of `count`.

---

## **Go Code – Step-by-Step Explanation**

1. **Define the Function**: Create a function `prefixCount` that takes a slice of strings (`words`) and a string (`pref`) as parameters.
2. **Initialize a Counter**: Declare an integer variable `count` and set it to 0. This will track the number of matching words.
3. **Loop Through Words**: Use a `for` loop with the `range` keyword to iterate over each word in the `words` slice.
4. **Check Prefix**: For each word, check if its first `len(pref)` characters match the prefix using slicing.
5. **Update Counter**: If the prefix matches, increment the `count` variable.
6. **Return the Result**: After iterating through all words, return the value of `count`.
