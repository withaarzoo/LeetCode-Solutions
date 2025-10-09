#include <vector>
#include <climits>
using namespace std;

class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size();
        int m = mana.size();
        if (m == 0) return 0LL;

        // prefix sums of skills (use long long)
        vector<long long> pref(n);
        for (int i = 0; i < n; ++i) {
            pref[i] = skill[i] + (i ? pref[i-1] : 0LL);
        }

        long long S = 0LL; // S_j (start time of current potion on wizard 0)
        // compute S_j iteratively using the derived formula
        for (int j = 1; j < m; ++j) {
            long long prev = (long long)mana[j-1];
            long long cur  = (long long)mana[j];
            long long best = LLONG_MIN;
            for (int i = 0; i < n; ++i) {
                long long prev_pref = (i ? pref[i-1] : 0LL);
                long long cand = pref[i] * prev - prev_pref * cur;
                if (cand > best) best = cand;
            }
            S += best;
        }

        // final finish time = S_{m-1} + total skill sum * mana[last]
        long long ans = S + pref[n-1] * (long long)mana[m-1];
        return ans;
    }
};
