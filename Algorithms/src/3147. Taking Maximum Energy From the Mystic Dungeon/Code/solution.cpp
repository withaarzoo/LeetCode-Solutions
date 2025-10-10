class Solution {
public:
    int maximumEnergy(vector<int>& energy, int k) {
        int n = energy.size();
        long long ans = LLONG_MIN; // store as long long to avoid overflow
        // For each residue class mod k
        for (int r = 0; r < k; ++r) {
            long long cur = 0;
            // compute last index in this class: r + t*k where t = (n-1 - r)/k
            int last = r + ((n - 1 - r) / k) * k;
            for (int i = last; i >= r; i -= k) {
                cur += energy[i];        // suffix sum starting at i
                ans = max(ans, cur);     // update global maximum
            }
        }
        return (int)ans;
    }
};
