# Rotate String - Solution Explanation

This repository provides step-by-step explanations for solving the "Rotate String" problem across multiple programming languages: C++, Java, JavaScript, Python, and Go.

## Problem Summary

Given two strings, `s` and `goal`, determine if `s` can become `goal` after a certain number of rotations. A rotation involves moving the leftmost character to the rightmost position, one character at a time.

For example, if `s = "abcde"`:

- After one rotation: `"bcdea"`
- After two rotations: `"cdeab"`

The goal is to check if such rotations can make `s` equal to `goal`.

## Approach and Solution

### Step 1: Check Lengths

- First, verify if the lengths of `s` and `goal` are equal.
- If they are not, then `s` cannot be rotated to become `goal`, so we immediately return `false`.

### Step 2: Concatenate `s` with Itself

- Concatenate `s` with itself to create a new string, `doubled = s + s`.
- This combined string includes all possible rotations of `s` as substrings. By doubling `s`, each possible rotation is represented as a contiguous substring within `doubled`.

### Step 3: Check for `goal` in `doubled`

- Use a substring search to see if `goal` is contained within `doubled`.
- If `goal` is found within `doubled`, then `s` can indeed be rotated to match `goal`, and we return `true`.
- If `goal` is not found in `doubled`, then it's impossible to match `s` to `goal` via rotations, so we return `false`.

---

## Detailed Code Explanations

### C++ Code Explanation

1. **Check Lengths**: First, check if the length of `s` is equal to `goal`. If they are different, return `false`.
2. **Concatenate Strings**: Create a new string by concatenating `s` with itself.
3. **Search for Substring**: Use the `find` function to check if `goal` is a substring of the concatenated string. If `goal` is found, return `true`; otherwise, return `false`.

### Java Code Explanation

1. **Check Lengths**: First, compare the lengths of `s` and `goal`. If they don’t match, return `false`.
2. **Concatenate Strings**: Create a new string by combining `s` with itself.
3. **Search for Substring**: Use the `contains` method to see if `goal` is a substring of the concatenated string. If `goal` exists in this combined string, return `true`; otherwise, return `false`.

### JavaScript Code Explanation

1. **Check Lengths**: Begin by comparing the lengths of `s` and `goal`. If they are different, immediately return `false`.
2. **Concatenate Strings**: Construct a new string by concatenating `s` with itself.
3. **Search for Substring**: Use the `includes` method to determine if `goal` is a substring of the concatenated string. If `goal` is found, return `true`; otherwise, return `false`.

### Python Code Explanation

1. **Check Lengths**: Start by checking if the length of `s` matches the length of `goal`. If not, return `false`.
2. **Concatenate Strings**: Create a new string by concatenating `s` with itself.
3. **Search for Substring**: Use the `in` keyword to check if `goal` is a substring of the concatenated string. If `goal` is found within, return `true`; otherwise, return `false`.

### Go Code Explanation

1. **Check Lengths**: First, verify if `s` and `goal` have the same length. If they don’t, return `false`.
2. **Concatenate Strings**: Create a new string by concatenating `s` with itself.
3. **Search for Substring**: Use `strings.Contains` to check if `goal` is within the concatenated string. If `goal` is a substring, return `true`; otherwise, return `false`.

---

## Complexity Analysis

- **Time Complexity**: \(O(n)\), where \(n\) is the length of `s` or `goal`, because checking if `goal` is a substring of the doubled string takes linear time.
- **Space Complexity**: \(O(n)\), as we create a concatenated string of size \(2 \times n\).

## Conclusion

This approach provides an efficient solution by leveraging string concatenation and substring search instead of generating each rotation explicitly. The method is consistent across different languages, offering a uniform understanding and solution to the problem.
