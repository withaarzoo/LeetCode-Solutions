class Solution {
public:
    int reverseBits(int n) {
        unsigned int result = 0;
        
        // We iterate exactly 32 times (32 bits)
        for (int i = 0; i < 32; i++) {
            
            // Shift result left to make space for next bit
            result <<= 1;
            
            // Get last bit of n and add to result
            result |= (n & 1);
            
            // Shift n right to process next bit
            n >>= 1;
        }
        
        return result;
    }
};
