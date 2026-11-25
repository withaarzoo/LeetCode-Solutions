class Solution {
public:
    int smallestUniqueSubarray(vector<int>& nums) {
        int n = nums.size();
        const long long MOD1 = 1000000007LL, MOD2 = 1000000009LL;
        const long long BASE1 = 1000003LL, BASE2 = 1000033LL;
        vector<long long> h1(n + 1), h2(n + 1), p1(n + 1), p2(n + 1);
        p1[0] = p2[0] = 1;
        for (int i = 0; i < n; i++) {
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1;
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2;
            p1[i + 1] = p1[i] * BASE1 % MOD1;
            p2[i + 1] = p2[i] * BASE2 % MOD2;
        }
        auto hasUnique = [&](int len) -> bool {
            map<pair<long long, long long>, int> cnt;
            for (int i = 0; i + len <= n; i++) {
                long long a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
                long long b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
                cnt[{a, b}]++;
            }
            for (auto& e : cnt) if (e.second == 1) return true;
            return false;
        };
        int lo = 1, hi = n, ans = n;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (hasUnique(mid)) {
                ans = mid;
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return ans;
    }
};