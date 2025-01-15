# Minimize XOR  

## Problem Overview  

The task is to find the number that minimizes the XOR value between `num1` and the result while ensuring the result has the same number of set bits (`1`s in binary) as `num2`.

---

## Approach  

We aim to match the number of `1`s (set bits) in `num1` to those in `num2` to minimize the XOR value. Here's a step-by-step explanation of the approach for each language.

---

### C++ Code  

1. **Count the Set Bits**  
   - Use the `__builtin_popcount` function to count the number of `1`s in `num1` and `num2`.  

2. **Check for Equality**  
   - If the set bits in `num1` and `num2` are already equal, return `num1` directly.  

3. **Adjust Excess Bits (if needed)**  
   - If `num1` has more set bits than `num2`, remove excess `1`s starting from the least significant bit (LSB).  

4. **Add Missing Bits (if needed)**  
   - If `num1` has fewer set bits than `num2`, add `1`s starting from the least significant positions until the count matches.  

5. **Return the Result**  
   - Once the number of set bits matches, return the result.  

---

### Java Code  

1. **Count the Set Bits**  
   - Use the `Integer.bitCount` method to count the `1`s in `num1` and `num2`.  

2. **Handle Equal Cases**  
   - If `num1` and `num2` already have the same number of set bits, return `num1`.  

3. **Remove Excess Bits**  
   - If `num1` has too many `1`s, iterate over the bits and clear some of them starting from the least significant bit.  

4. **Add Missing Bits**  
   - If `num1` has too few `1`s, iterate and set the least significant unset bits until the set bit count matches.  

5. **Return the Result**  
   - Once adjusted, return the updated value.  

---

### JavaScript Code  

1. **Count the Set Bits**  
   - Convert `num1` and `num2` to binary strings, then count the number of `1`s in each using `split('1').length - 1`.  

2. **Check for Equality**  
   - If the set bit counts are equal, return `num1` immediately.  

3. **Remove Excess `1`s**  
   - If `num1` has extra `1`s, iterate through its binary representation and clear them until the count matches `num2`.  

4. **Add Missing `1`s**  
   - If `num1` has too few `1`s, iterate and set bits in the least significant positions.  

5. **Return the Result**  
   - The adjusted value is returned as the final result.  

---

### Python Code  

1. **Count the Set Bits**  
   - Use Python’s `bin()` function to get the binary representation of `num1` and `num2`. Use `.count('1')` to count the set bits.  

2. **Check for Equality**  
   - If the set bit counts are already equal, return `num1` without modifications.  

3. **Remove Extra Bits**  
   - If `num1` has too many `1`s, clear bits starting from the least significant position until the set bit count matches.  

4. **Add Missing Bits**  
   - If `num1` has too few `1`s, set bits starting from the least significant unset bit.  

5. **Return the Result**  
   - Return the final adjusted value.  

---

### Go Code  

1. **Count the Set Bits**  
   - Use the `bits.OnesCount` function to count the set bits (`1`s) in `num1` and `num2`.  

2. **Check Equality**  
   - If the set bit counts of `num1` and `num2` are equal, return `num1`.  

3. **Remove Extra Bits**  
   - If `num1` has too many `1`s, clear some of them starting from the least significant position.  

4. **Add Missing Bits**  
   - If `num1` has too few `1`s, set bits at the least significant positions.  

5. **Return the Result**  
   - Once adjusted, return the result.  

---

## Complexity  

- **Time Complexity**:  
  - Counting set bits is \(O(\log n)\), and iterating over bits is also \(O(\log n)\).  
  - Overall time complexity is \(O(\log n)\), where \(n\) is the larger of `num1` or `num2`.  

- **Space Complexity**:  
  - We use a constant amount of extra space, so the space complexity is \(O(1)\).  
