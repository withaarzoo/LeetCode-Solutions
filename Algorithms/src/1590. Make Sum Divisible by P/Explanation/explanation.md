## README: Minimum Subarray Length to Remove for Divisibility

### Problem Description

Given an array of integers `nums` and an integer `p`, the goal is to remove the smallest subarray from `nums` such that the sum of the remaining array is divisible by `p`. The task is to return the length of this subarray, or `-1` if it's not possible.

### Approach (Applicable to All Languages)

1. **Step 1: Calculate Total Sum**  
   Start by calculating the total sum of the array `nums`.

2. **Step 2: Find Remainder of Total Sum**  
   Find the remainder of the total sum when divided by `p`.  
   - If the remainder is 0, the entire array is already divisible by `p`, and no subarray needs to be removed. Thus, return `0`.

3. **Step 3: Track Prefix Sums Using a Hash Map**  
   Use a hash map to track prefix sums' remainders when divided by `p`. This helps in identifying subarrays that can be removed to achieve divisibility by `p`.  
   - Initialize the hash map with `{0: -1}` to handle cases where the entire prefix contributes to divisibility.

4. **Step 4: Iterate Over the Array**  
   Loop through the array, maintaining a running sum (prefix sum) of the elements. For each element:
   - Compute the current prefix sum's remainder when divided by `p`.
   - Calculate the target remainder we need to achieve divisibility by `p`.

5. **Step 5: Check for Subarrays**  
   Check if the target remainder exists in the hash map:
   - If it exists, calculate the length of the subarray that can be removed, and update the minimum length of such subarrays found so far.
   - Update the hash map with the current prefix sum's remainder.

6. **Step 6: Return the Result**  
   After iterating through the array:
   - If a valid subarray was found, return the minimum length of such subarrays.
   - If no valid subarray was found, return `-1`.

---

### C++ Code Explanation

1. **Initialize Variables**: Start by calculating the total sum of the array and finding the remainder when divided by `p`.
2. **Handle Special Case**: If the remainder is 0, return 0 as no subarray needs to be removed.
3. **Track Prefix Sums**: Use an unordered map to track the remainder of prefix sums.
4. **Iterate Over the Array**: For each element, calculate the running sum, find the corresponding target remainder, and check if it exists in the map.
5. **Update Result**: If a valid subarray is found, update the minimum length.
6. **Return Result**: Return the minimum length or `-1` if no valid subarray is found.

---

### Java Code Explanation

1. **Calculate Total Sum**: Calculate the total sum of the elements in the array.
2. **Handle Special Case**: If the total sum is already divisible by `p`, return `0`.
3. **Use HashMap**: Use a HashMap to track the remainders of prefix sums.
4. **Iterate Over Array**: For each element, update the prefix sum and calculate the remainder. Check for the required target remainder in the map.
5. **Track Minimum Length**: Keep track of the minimum length of the subarray that, if removed, makes the sum divisible by `p`.
6. **Return Final Result**: Return `-1` if no such subarray is found, otherwise return the minimum length.

---

### JavaScript Code Explanation

1. **Calculate Total Sum**: Sum up all the numbers in the array and find the remainder of the sum when divided by `p`.
2. **Handle Simple Case**: If the remainder is 0, the entire array is divisible by `p`.
3. **Map for Prefix Modulo**: Create a Map to store the remainders of the prefix sums.
4. **Iterate Through Array**: For each element, update the prefix sum and calculate the remainder. Check for the target remainder in the map.
5. **Update Minimum Length**: Whenever a valid subarray is found, calculate its length and update the minimum length.
6. **Return the Result**: If no valid subarray is found, return `-1`.

---

### Python Code Explanation

1. **Compute Total Sum**: Calculate the total sum of the array and find the remainder when divided by `p`.
2. **Handle Divisibility Case**: If the remainder is zero, return 0.
3. **Prefix Sum Tracking with Dictionary**: Use a dictionary to track remainders of prefix sums and their indices.
4. **Iterate Over Elements**: Compute the running sum and its remainder for each element. Check if the target remainder exists in the dictionary.
5. **Update Minimum Subarray Length**: Whenever a matching remainder is found, update the minimum length of the subarray that needs to be removed.
6. **Return the Minimum Length or -1**: After iterating, return the minimum subarray length or -1 if no valid subarray is found.

---

### Go Code Explanation

1. **Total Sum Calculation**: Start by calculating the total sum of the array and finding the remainder when divided by `p`.
2. **Handle Special Case**: If the remainder is 0, return 0 since no subarray needs to be removed.
3. **Prefix Mod Tracking**: Use a map to track prefix sums and their remainders.
4. **Iterate Through Array**: For each element, update the prefix sum, calculate the current remainder, and check if the target remainder exists in the map.
5. **Update Minimum Length**: If a valid subarray is found, update the minimum length.
6. **Return Result**: Return the minimum length or `-1` if no valid subarray is found.
