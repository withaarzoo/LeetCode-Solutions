class Solution {
public:
    int sortableIntegers(vector<int>& nums) {
        int n = nums.size();
        vector<int> sorted = nums;
        sort(sorted.begin(), sorted.end());
        vector<int> divs;
        for (int i = 1; i * i <= n; ++i) {
            if (n % i == 0) {
                divs.push_back(i);
                if (i != n / i) divs.push_back(n / i);
            }
        }
        int ans = 0;
        for (int k : divs) {
            bool ok = true;
            for (int s = 0; s < n && ok; s += k) {
                int desc = 0, pos = -1;
                for (int j = 0; j < k; ++j) {
                    int a = nums[s + j];
                    int b = nums[s + (j + 1) % k];
                    if (a > b) {
                        ++desc;
                        pos = j;
                        if (desc > 1) break;
                    }
                }
                if (desc > 1) {
                    ok = false;
                    break;
                }
                int r = (desc == 0) ? 0 : (pos + 1) % k;
                for (int j = 0; j < k; ++j) {
                    if (nums[s + (r + j) % k] != sorted[s + j]) {
                        ok = false;
                        break;
                    }
                }
            }
            if (ok) ans += k;
        }
        return ans;
    }
};