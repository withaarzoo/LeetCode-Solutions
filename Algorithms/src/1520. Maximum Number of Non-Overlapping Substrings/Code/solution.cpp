class Solution {
public:
    vector<string> maxNumOfSubstrings(string s) {
        int n = s.size();
        // first and last occurrence of each letter
        vector<int> left(26, -1), right(26, -1);
        for (int i = 0; i < n; ++i) {
            int c = s[i] - 'a';
            if (left[c] == -1) left[c] = i;   // record first time we see it
            right[c] = i;                     // always update last time
        }

        // collect every valid closed interval
        vector<pair<int,int>> intervals;
        for (int i = 0; i < 26; ++i) {
            if (left[i] == -1) continue;      // letter never appeared
            int L = left[i], R = right[i];
            bool valid = true;
            // expand right end while scanning the current range
            for (int j = L; j <= R; ++j) {
                int c = s[j] - 'a';
                if (left[c] < L) {            // needs to start earlier -> not closed
                    valid = false;
                    break;
                }
                R = max(R, right[c]);         // push right end if needed
            }
            if (valid) intervals.emplace_back(L, R);
        }

        // sort by ending position so greedy can pick earliest-ending first
        sort(intervals.begin(), intervals.end(),
             [](const pair<int,int>& a, const pair<int,int>& b) {
                 return a.second < b.second;
             });

        // greedy selection of non-overlapping intervals
        vector<string> ans;
        int lastEnd = -1;
        for (auto [L, R] : intervals) {
            if (L > lastEnd) {                // completely after previous piece
                ans.push_back(s.substr(L, R - L + 1));
                lastEnd = R;
            }
        }
        return ans;
    }
};