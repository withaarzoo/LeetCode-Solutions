#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int maxDistinctElements(vector<int>& nums, int k) {
        int n = nums.size();
        // Build intervals [num - k, num + k] using long long to avoid overflow
        vector<pair<long long,long long>> intervals;
        intervals.reserve(n);
        for (int x : nums) {
            long long l = (long long)x - k;
            long long r = (long long)x + k;
            intervals.emplace_back(l, r);
        }
        // Sort by right endpoint
        sort(intervals.begin(), intervals.end(), [](auto &a, auto &b){
            if (a.second != b.second) return a.second < b.second;
            return a.first < b.first;
        });

        long long last_assigned = LLONG_MIN / 4; // very small
        int ans = 0;
        for (auto &it : intervals) {
            long long l = it.first, r = it.second;
            long long assigned = max(l, last_assigned + 1);
            if (assigned <= r) {
                ans++;
                last_assigned = assigned;
            }
        }
        return ans;
    }
};
