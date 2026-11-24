class Solution {
public:
    int maxLength(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            long long prod = 1;
            int g = 0;
            long long lc = 1;
            for (int j = i; j < n; j++) {
                int x = nums[j];
                if (g == 0) g = x;
                else {
                    int a = g, b = x;
                    while (b) {
                        int t = b;
                        b = a % b;
                        a = t;
                    }
                    g = a;
                }
                long long gg = lc;
                int b = x;
                while (b) {
                    long long t = b;
                    b = gg % b;
                    gg = t;
                }
                if (lc / gg > LLONG_MAX / x) break;
                lc = lc / gg * x;
                if (prod > LLONG_MAX / x) break;
                prod *= x;
                if (prod == lc * g) ans = max(ans, j - i + 1);
            }
        }
        return ans;
    }
};