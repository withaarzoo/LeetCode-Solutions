# Solution Explanation for "Find K-th Bit in N-th Binary String"

This README contains step-by-step explanations of the solution in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Each explanation describes how the solution recursively calculates the K-th bit in the N-th binary string without generating the entire string. Instead, it takes advantage of the symmetry and structure of the string.

## C++ Solution

### Step-by-Step Explanation

1. **Base Case Handling**:  
   If `n` equals 1, return '0' immediately, since the first binary string (`S1`) is always "0".

2. **Calculate Length of Sn**:  
   For a given `n`, the total length of the binary string `Sn` is calculated as `2^n - 1`.

3. **Determine Midpoint**:  
   Find the middle position of the binary string `Sn`. The middle position is `(length / 2) + 1`.

4. **Check If K is the Middle Bit**:  
   If `k` equals the middle position, return '1', since the middle bit in any `Sn` is always '1'.

5. **Check First Half**:  
   If `k` is smaller than the middle position, the bit is located in the first half of the binary string. Recursively compute the K-th bit in `Sn-1`.

6. **Check Second Half**:  
   If `k` is greater than the middle position, the bit is located in the second half. This part is the reversed and inverted version of the first half. Recursively compute the K-th bit in `Sn-1`, but invert the result.

---

## Java Solution

### Step-by-Step Explanation

1. **Handle Base Case**:  
   When `n = 1`, return '0' because the first binary string (`S1`) is "0". This is our base case.

2. **Calculate Length of Binary String**:  
   For any given `n`, compute the length of the binary string as `2^n - 1`.

3. **Locate Middle Bit**:  
   The middle position in the binary string is the result of `(length / 2) + 1`.

4. **Check for Middle Element**:  
   If `k` equals the middle position, return '1'. This is because the middle bit is always '1' in all strings.

5. **First Half Check**:  
   If `k` is smaller than the middle, the K-th bit lies in the first half. Make a recursive call to find the K-th bit in `Sn-1`.

6. **Second Half Check**:  
   If `k` is larger than the middle, the K-th bit lies in the second half, which is the reversed and inverted version of the first half. Recursively call for the K-th bit in the first half, and invert the result.

---

## JavaScript Solution

### Step-by-Step Explanation

1. **Handle Base Case**:  
   For the simplest case where `n = 1`, the binary string is "0", so return '0'. This is our base condition for the recursion.

2. **Calculate Length of the Current String**:  
   Calculate the length of the binary string as `2^n - 1`, which represents the total length of `Sn`.

3. **Find the Middle Position**:  
   The middle element in the binary string `Sn` can be found by calculating `(length / 2) + 1`.

4. **Check Middle Element**:  
   If `k` is exactly the middle element, return '1' because the middle bit is always '1' in any of the binary strings.

5. **Recurse for First Half**:  
   If `k` is less than the middle, it means that the bit is in the first half of the string. Recursively find the bit by calling the function on `Sn-1`.

6. **Recurse for Second Half**:  
   If `k` is greater than the middle, the bit lies in the second half. The second half is the inverted and reversed version of the first half. Compute the bit recursively for `Sn-1`, and invert the result.

---

## Python Solution

### Step-by-Step Explanation

1. **Base Case Condition**:  
   The base case is when `n = 1`, where the binary string is "0". In this case, return '0'. This serves as the stopping condition for the recursion.

2. **Calculate Length of Sn**:  
   Compute the length of the binary string `Sn` as `2^n - 1`.

3. **Find the Middle Bit**:  
   Calculate the middle position as `(length // 2) + 1`. This is where the symmetry of the string plays a role.

4. **Check for Middle Bit**:  
   If `k` is equal to the middle position, return '1'. The middle bit in any `Sn` is always '1'.

5. **Handle First Half**:  
   If `k` is less than the middle position, the bit lies in the first half of `Sn`. Recursively compute the K-th bit in `Sn-1`.

6. **Handle Second Half**:  
   If `k` is greater than the middle position, the bit lies in the second half of the string. The second half is an inverted and reversed version of the first half. Recursively compute the bit and invert the result.

---

## Go Solution

### Step-by-Step Explanation

1. **Base Case Check**:  
   When `n = 1`, return '0' since the first binary string (`S1`) is "0". This is the simplest case and serves as the stopping condition.

2. **Calculate the Length of Sn**:  
   For any given `n`, calculate the length of the binary string `Sn` as `2^n - 1`.

3. **Determine the Midpoint**:  
   The middle position of `Sn` can be found by computing `(length / 2) + 1`. This divides the string into two symmetric parts.

4. **Check if K is in the Middle**:  
   If `k` equals the middle position, return '1'. The middle bit is always '1' in every binary string.

5. **Handle First Half**:  
   If `k` is less than the middle, the K-th bit lies in the first half of `Sn`. Recursively compute the bit by reducing the problem to `Sn-1`.

6. **Handle Second Half**:  
   If `k` is greater than the middle, the bit lies in the second half. This half is an inverted and reversed version of the first half. Recursively compute the K-th bit in `Sn-1` and invert the result.

---

## Conclusion

Across all five languages, the logic follows the same recursive pattern:

- If the bit is in the first half, we reduce the problem to `Sn-1`.
- If the bit is in the second half, we also reduce the problem to `Sn-1` but with an inverted result.
- This efficient recursion avoids the need to construct the entire binary string, saving both time and space.
