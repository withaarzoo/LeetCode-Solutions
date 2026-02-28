class Solution {
public:
    int concatenatedBinary(int n) {
        const int MOD = 1e9 + 7;
        long long ans = 0;      // To store result
        int bitLength = 0;      // Current number of bits needed
        
        for (int i = 1; i <= n; i++) {
            
            // If i is power of 2, increase bit length
            if ((i & (i - 1)) == 0) {
                bitLength++;
            }
            
            // Shift left to make space for new number
            ans = ((ans << bitLength) % MOD + i) % MOD;
        }
        
        return (int)ans;
    }
};