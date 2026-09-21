class Solution { 
public: 
    vector<long long> resultArray(vector<int>& nums, int k) { 
        vector<long long> result(k, 0);
        vector<long long> curr(k, 0);
        for (int num : nums) {
            int m = num % k;
            vector<long long> next(k, 0);
            next[m] = 1;
            for (int prev = 0; prev < k; prev++) {
                if (curr[prev] > 0) {
                    int nr = (int)((prev * 1LL * m) % k);
                    next[nr] += curr[prev];
                }
            }
            for (int r = 0; r < k; r++) {
                result[r] += next[r];
            }
            curr = next;
        }
        return result;
    } 
};