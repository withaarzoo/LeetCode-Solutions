class Solution {
    public int reverseBits(int n) {
        int result = 0;
        
        for (int i = 0; i < 32; i++) {
            
            // Shift result left
            result <<= 1;
            
            // Add last bit of n
            result |= (n & 1);
            
            // Logical right shift (important for negative numbers)
            n >>>= 1;
        }
        
        return result;
    }
}
