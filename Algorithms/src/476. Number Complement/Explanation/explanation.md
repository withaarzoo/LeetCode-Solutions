# Find Complement of a Number

This README explains how to find the complement of a given number across various programming languages. Each step in the code is broken down to help you understand the underlying logic.

## C++ Implementation

### Step-by-Step Explanation

1. **Initialize the Mask:**
   - A variable `mask` is initialized to 0. This will eventually represent a bitmask where all bits are set to 1, corresponding to the binary length of the input number `num`.

2. **Copy the Input Number:**
   - Create a temporary variable `temp` initialized with the value of `num`. This is used to determine the binary length of `num`.

3. **Generate the Bitmask:**
   - The loop runs while `temp` is not 0. In each iteration:
     - Shift the current bits in `mask` to the left by 1 position to make room for the new bit.
     - Set the least significant bit of `mask` to 1.
     - Right shift `temp` by 1 to move to the next bit in the binary representation.

4. **Compute the Complement:**
   - Use XOR between `num` and `mask` to flip all bits of `num`, producing its complement.

## Java Implementation

### Step-by-Step Explanation

1. **Initialize the Mask:**
   - A variable `mask` is initialized to 0, which will be used to create a bitmask where all bits are set to 1.

2. **Create a Temporary Copy:**
   - Copy the input number `num` into a temporary variable `temp` to determine the number of bits in `num`.

3. **Create the Bitmask:**
   - Loop until `temp` becomes 0:
     - Left shift `mask` by 1 to make space for the new bit.
     - OR the mask with 1 to set the least significant bit to 1.
     - Right shift `temp` by 1 to move to the next bit.

4. **Find the Complement:**
   - XOR `num` with the mask to flip all bits and get the complement.

## JavaScript Implementation

### Step-by-Step Explanation

1. **Initialize the Mask:**
   - The `mask` is initialized to 0. It will eventually have all bits set to 1 within the range of `num`.

2. **Create a Temporary Variable:**
   - Store the value of `num` in a temporary variable `temp` for manipulation.

3. **Create the Bitmask:**
   - Loop until `temp` is 0:
     - Shift `mask` to the left by 1 and set the least significant bit to 1.
     - Right shift `temp` by 1 to move to the next bit.

4. **Calculate the Complement:**
   - XOR `num` with the mask to flip all bits, effectively finding the complement.

## Python Implementation

### Step-by-Step Explanation

1. **Initialize the Mask:**
   - A variable `mask` is initialized to 0. This mask will create a number with all bits set to 1, matching the length of `num`'s binary representation.

2. **Temporary Variable Setup:**
   - Initialize `temp` to `num` to help determine the binary length of `num`.

3. **Generate the Mask:**
   - While `temp` is not 0:
     - Left shift `mask` by 1 to make space for the next bit and set the last bit to 1.
     - Right shift `temp` by 1 to move to the next bit.

4. **Find the Complement:**
   - XOR `num` with `mask` to flip all bits and compute the complement.

## Go Implementation

### Step-by-Step Explanation

1. **Initialize the Mask:**
   - The `mask` is initialized to 0. It will be used to cover all the bits of `num`.

2. **Create a Temporary Variable:**
   - Store `num` in a temporary variable `temp`, which will help in determining the bit length of `num`.

3. **Create the Bitmask:**
   - The loop continues while `temp` is not 0:
     - Shift the `mask` to the left by 1 bit and set the rightmost bit to 1.
     - Right shift `temp` by 1 bit to reduce its length.

4. **Compute the Complement:**
   - XOR `num` with `mask` to flip all bits and find the complement.
