class Solution {
    public int findComplement(int num) {
        // Step 1: Initialize a variable 'mask' with a value of 0.
        // This will be used to create a bitmask where all bits are set to 1.
        int mask = 0;

        // Step 2: Create a copy of the original number 'num' and store it in 'temp'.
        // This copy will be used to determine the number of bits in 'num'.
        int temp = num;

        // Step 3: Loop to create the bitmask.
        // The idea is to create a mask with all bits set to 1 that has the same number
        // of bits as 'num'.
        // Example: If num = 5 (binary 101), we want mask to become 111.
        while (temp != 0) {
            // Left shift 'mask' by 1 to make space for the new bit.
            // OR the current mask with 1 to set the least significant bit to 1.
            mask = (mask << 1) | 1;

            // Right shift 'temp' by 1 to check the next bit in the original number.
            temp >>= 1;
        }

        // Step 4: XOR the original number 'num' with the mask.
        // XOR operation will flip the bits of 'num' where the corresponding bit in
        // 'mask' is 1.
        // This gives us the complement of the number.
        return num ^ mask;
    }
}
