class Solution {
public:
    int minimumOneBitOperations(int n) {
        // Inverse Gray code: ans = n ^ (n>>1) ^ (n>>2) ^ ...
        int ans = 0;
        while (n) {
            ans ^= n;   // accumulate current bits
            n >>= 1;    // move to next higher influence
        }
        return ans;
    }
};
