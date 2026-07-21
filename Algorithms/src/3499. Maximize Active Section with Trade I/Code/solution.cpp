class Solution {
public:
    int maxActiveSectionsAfterTrade(string s) {
        // Count active sections in the original string
        int ones = 0;
        for (char c : s)
            if (c == '1')
                ones++;

        // Add virtual '1' on both ends
        string t = "1" + s + "1";

        vector<pair<char, int>> runs;

        // Build run-length encoding
        for (char c : t) {
            if (runs.empty() || runs.back().first != c)
                runs.push_back({c, 1});
            else
                runs.back().second++;
        }

        int best = 0;

        // Check every internal 1-block
        for (int i = 1; i + 1 < (int)runs.size(); i++) {
            if (runs[i].first == '1' &&
                runs[i - 1].first == '0' &&
                runs[i + 1].first == '0') {

                // Gain equals left zero length + right zero length
                best = max(best, runs[i - 1].second + runs[i + 1].second);
            }
        }

        return ones + best;
    }
};