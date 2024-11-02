# Circular Sentence Checker

## Problem Description

Given a sentence, determine if it is **circular**. A sentence is circular if:

1. The last character of each word matches the first character of the next word.
2. The last character of the last word matches the first character of the first word.

### Example

For example, `"leetcode exercises sound delightful"` is a circular sentence because:

- `leetcode`'s last character `'e'` matches `exercises`'s first character `'e'`
- `exercises`'s last character `'s'` matches `sound`'s first character `'s'`
- `sound`'s last character `'d'` matches `delightful`'s first character `'d'`
- `delightful`'s last character `'l'` matches `leetcode`'s first character `'l'`

---

## Solution Walkthrough (Step-by-Step)

### Step 1: Splitting the Sentence

The solution starts by **splitting the sentence into individual words**. Each programming language has its own method for splitting strings. This allows us to access each word separately.

- **C++**: We use a stringstream to extract words from the sentence and store them in a list.
- **Java**: Use `String.split(" ")` to split the sentence by spaces.
- **JavaScript**: Use `sentence.split(" ")` to split by spaces.
- **Python**: Use `sentence.split()` to separate by spaces.
- **Go**: Use `strings.Split(sentence, " ")` to split into words.

### Step 2: Loop through Each Word and Check Circular Condition

After splitting, we use a loop to **check the circular condition**:

1. We iterate through each word and compare its **last character** to the **first character of the next word**.
2. To complete the circular check, we use modular arithmetic to connect the **last word** to the **first word**.

- **C++**: Use an indexed `for` loop and `words[(i + 1) % words.size()]` to access the next word in a circular way.
- **Java**: Use a similar indexed loop with `words[(i + 1) % words.length]`.
- **JavaScript**: Same as Java, using `words[(i + 1) % words.length]`.
- **Python**: Use a `for` loop with `range(len(words))` and `words[(i + 1) % len(words)]`.
- **Go**: Use a loop with `len(words)` to access the next word in a circular way with `words[(i+1)%len(words)]`.

### Step 3: Compare Characters

Within each loop iteration, **retrieve the last character of the current word** and the **first character of the next word**:

1. If these characters **do not match**, return `false` immediately.
2. If all pairs match, the sentence is circular, so return `true`.

- **C++**: Access the last character with `words[i].back()` and the first with `words[(i + 1) % words.size()].front()`.
- **Java**: Use `charAt` to get the last and first characters, e.g., `words[i].charAt(words[i].length() - 1)` and `words[(i + 1) % words.length].charAt(0)`.
- **JavaScript**: Use `charAt` similarly to Java.
- **Python**: Use indexing, e.g., `words[i][-1]` for the last character and `words[(i + 1) % len(words)][0]` for the first.
- **Go**: Use indexing `words[i][len(words[i])-1]` and `words[(i+1)%len(words)][0]` to get characters.

### Step 4: Return Result

The function returns **true** if the sentence is circular (all character checks passed) or **false** if any check failed.

---

## Complexity Analysis

- **Time Complexity**: \(O(n)\), where \(n\) is the length of the sentence. Each word is processed once.
- **Space Complexity**: \(O(m)\), where \(m\) is the number of words in the sentence, as we store words in a list or array.

---

## Summary

This approach effectively checks each pair of words in the sentence for circularity using split operations, loops, and character comparisons. This is done efficiently with minimal space usage for temporary storage of words.
