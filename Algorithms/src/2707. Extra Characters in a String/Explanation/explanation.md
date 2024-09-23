# Minimum Extra Characters in String Problem - Step by Step Explanation

This README walks through the solution for calculating the minimum number of extra characters required to split a string into valid words from a given dictionary. We explore this in C++, Java, JavaScript, Python, and Go. Each step provides a code snippet followed by detailed explanations.

---

### C++ Solution

1. **Convert Dictionary to Unordered Set:**
   - Convert the list of dictionary words into an unordered set for fast lookups (O(1) time complexity).

2. **DP Array Initialization:**
   - Create a DP array `dp[]` where `dp[i]` represents the minimum number of extra characters required to parse the substring `s[0:i]`.
   - Initialize each entry of `dp[]` to the maximum possible number of extra characters, which is `n` (the length of the string `s`).
   - Set the base case `dp[0] = 0` (no extra characters for an empty string).

3. **Iterate Over String:**
   - For each index `i` from 1 to `n`, check all possible substrings ending at `i` by looping through each `j` from 0 to `i-1`.
   - Extract the substring `s[j:i]` and check if it exists in the dictionary.

4. **Update DP Array:**
   - If a valid substring is found, update `dp[i]` by considering `dp[j]` (no extra characters needed for this substring).
   - If no valid substring is found, treat the current character as an extra character and update `dp[i]` by incrementing `dp[i-1]`.

5. **Return Result:**
   - The result is stored in `dp[n]`, which represents the minimum extra characters needed for the entire string.

---

### Java Solution

1. **Create Dictionary Set:**
   - Convert the dictionary array into a set for fast O(1) lookups.

2. **Initialize DP Array:**
   - Create a `dp[]` array of size `n+1`, where `dp[i]` stores the minimum number of extra characters required to process the first `i` characters of the string.
   - Set all values in `dp[]` to `n` (maximum possible extra characters) and set `dp[0] = 0`.

3. **Check All Substrings:**
   - Loop over each position `i` in the string.
   - For each `i`, check all possible substrings `s[j:i]` where `j < i` and check if the substring is present in the dictionary.

4. **Update DP Array:**
   - If a valid substring is found, update `dp[i]` with the minimum between `dp[i]` and `dp[j]`.
   - If no valid substring is found, consider the current character as an extra and update `dp[i]` by adding 1 to `dp[i-1]`.

5. **Final Output:**
   - The minimum extra characters for the entire string is stored in `dp[n]`.

---

### JavaScript Solution

1. **Convert Dictionary to Set:**
   - Use a `Set` to store dictionary words for O(1) lookups.

2. **Initialize DP Array:**
   - Create an array `dp[]` of size `n+1` and initialize it with `n` (the maximum number of extra characters).
   - Set the base case `dp[0] = 0`.

3. **Check Possible Substrings:**
   - Loop through each index `i` from 1 to `n`, and for each `i`, check all possible substrings `s[j:i]` where `j < i`.
   - Use `Set.has()` to check if the substring exists in the dictionary.

4. **Update DP Array:**
   - If a valid substring is found, update `dp[i]` by considering `dp[j]`.
   - Otherwise, treat the current character as extra and update `dp[i]` to `dp[i-1] + 1`.

5. **Final Answer:**
   - The answer is stored in `dp[n]`.

---

### Python Solution

1. **Convert Dictionary to Set:**
   - Convert the dictionary list into a set for O(1) time complexity during lookup.

2. **Initialize DP Array:**
   - Create a DP array `dp[]` where `dp[i]` represents the minimum extra characters needed for the substring `s[0:i]`.
   - Initialize `dp[]` with the value `n`, and set `dp[0] = 0` (base case).

3. **Iterate Over Substrings:**
   - Loop through each index `i` from 1 to `n` and check all substrings `s[j:i]` where `j < i`.
   - If the substring exists in the dictionary, update `dp[i]`.

4. **Update for Extra Characters:**
   - If no valid substring is found, consider the current character `s[i-1]` as extra and update `dp[i] = dp[i-1] + 1`.

5. **Return Result:**
   - The final result is found in `dp[n]`.

---

### Go Solution

1. **Create Dictionary Map:**
   - Create a map to store dictionary words for O(1) lookups.

2. **Initialize DP Array:**
   - Create a DP array `dp[]` where `dp[i]` stores the minimum number of extra characters required up to index `i`.
   - Initialize the DP array to `n` (maximum possible extra characters) and set `dp[0] = 0`.

3. **Iterate Over Substrings:**
   - Loop through each index `i` from 1 to `n` and check all substrings `s[j:i]`.
   - If the substring is in the dictionary, update `dp[i]` by considering `dp[j]`.

4. **Handle Extra Characters:**
   - If no valid substring is found, consider the current character `s[i-1]` as extra and update `dp[i] = dp[i-1] + 1`.

5. **Final Result:**
   - The result is stored in `dp[n]`.
