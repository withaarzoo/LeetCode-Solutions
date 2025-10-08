#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    vector<int> successfulPairs(vector<int>& spells, vector<int>& potions, long long success) {
        // sort potions so we can binary-search
        sort(potions.begin(), potions.end());
        int n = spells.size();
        int m = potions.size();
        vector<int> ans(n, 0);

        for (int i = 0; i < n; ++i) {
            long long s = spells[i]; // promote to long long to avoid overflow
            // smallest potion value needed so that s * potion >= success
            long long need = (success + s - 1) / s; // ceil(success / s)
            // find first potion >= need
            auto it = lower_bound(potions.begin(), potions.end(), need);
            ans[i] = m - int(it - potions.begin()); // count of valid potions
        }
        return ans;
    }
};
