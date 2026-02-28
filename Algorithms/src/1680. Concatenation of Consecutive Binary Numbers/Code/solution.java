class Solution {
    public int concatenatedBinary(int n) {
        final int MOD = 1_000_000_007;
        long ans = 0;
        int bitLength = 0;
        
        for (int i = 1; i <= n; i++) {
            
            // If i is power of 2, bit length increases
            if ((i & (i - 1)) == 0) {
                bitLength++;
            }
            
            // Shift and add
            ans = ((ans << bitLength) % MOD + i) % MOD;
        }
        
        return (int) ans;
    }
}