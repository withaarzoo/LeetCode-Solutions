class Solution {
public:
    long long maximumTotalDamage(vector<int>& power) {
        // Count frequency and compute total per unique value
        unordered_map<long long, long long> freq;
        for (int v : power) freq[v] += 1;
        // extract and sort unique values
        vector<long long> vals;
        vals.reserve(freq.size());
        for (auto &p : freq) vals.push_back(p.first);
        sort(vals.begin(), vals.end());
        int m = vals.size();
        if (m == 0) return 0;
        // valueSum[i] = vals[i] * freq[vals[i]]
        vector<long long> valueSum(m);
        for (int i = 0; i < m; ++i) valueSum[i] = vals[i] * freq[vals[i]];
        // dp[i] = best up to i
        vector<long long> dp(m, 0);
        dp[0] = valueSum[0];
        for (int i = 1; i < m; ++i) {
            // find last index j < i with vals[j] <= vals[i] - 3
            long long need = vals[i] - 3;
            int j = upper_bound(vals.begin(), vals.begin() + i, need) - vals.begin() - 1;
            long long take = valueSum[i] + (j >= 0 ? dp[j] : 0);
            long long skip = dp[i-1];
            dp[i] = max(skip, take);
        }
        return dp[m-1];
    }
};
