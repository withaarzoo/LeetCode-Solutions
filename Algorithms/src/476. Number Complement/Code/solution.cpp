class Solution
{
public:
    int findComplement(int num)
    {
        // Step 1: Initialize a variable 'mask' to store the bitmask.
        // The bitmask will eventually have all bits set to 1, corresponding to the binary length of 'num'.
        int mask = 0;

        // Step 2: Create a temporary variable 'temp' initialized with the value of 'num'.
        // This will be used to determine the length of 'num' in binary.
        int temp = num;

        // Step 3: Generate the mask by shifting bits to the left and OR-ing with 1.
        // This loop will continue until 'temp' becomes 0.
        while (temp != 0)
        {
            // Shift the current bits in 'mask' to the left by 1 position to make room for the new bit.
            mask = (mask << 1) | 1; // OR-ing with 1 ensures that the least significant bit is set to 1.

            // Right shift 'temp' by 1 bit to move to the next bit in the binary representation.
            temp >>= 1;
        }

        // Step 4: XOR 'num' with 'mask' to get the complement.
        // XOR operation between the number and the mask will flip the bits of 'num', producing its complement.
        return num ^ mask;
    }
};
