/* C++ (O(n log n) with std::set) */
class Solution {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        int n = rains.size();
        vector<int> ans(n, 1);
        unordered_map<int,int> last;   // lake -> last day index it rained
        set<int> dryDays;             // indices of dry days (sorted)

        for (int i = 0; i < n; ++i) {
            if (rains[i] > 0) {
                int lake = rains[i];
                ans[i] = -1;  // rainy day must be -1
                if (last.count(lake)) {
                    int prevDay = last[lake];
                    // find earliest dry day strictly after prevDay
                    auto it = dryDays.lower_bound(prevDay + 1);
                    if (it == dryDays.end()) {
                        return {}; // impossible to avoid flood
                    }
                    ans[*it] = lake;   // dry that lake on this dry day
                    dryDays.erase(it); // remove used dry day
                }
                last[lake] = i; // update last rainy day for lake
            } else {
                dryDays.insert(i); // available dry day
                // ans[i] stays 1 unless later assigned to dry a lake
            }
        }
        return ans;
    }
};
