class Solution
{
public:
    vector<int> maxActiveSectionsAfterTrade(string s, vector<vector<int>> &queries)
    {
        int n = s.length(); // Get the size of the original string
        int total_ones = 0; // Track the total number of 1s initially present in the entire string

        // Count all the '1's in the string since our trades will only ADD to this base count
        for (char c : s)
        {
            if (c == '1')
                total_ones++;
        }

        // Arrays to represent compressed segments of identical characters
        vector<int> type;    // Stores 0 for a 0-segment, 1 for a 1-segment
        vector<int> start;   // Stores the starting index of the segment in the original string
        vector<int> end_idx; // Stores the ending index of the segment in the original string

        // Group consecutive identical characters into segments
        for (int i = 0; i < n;)
        {
            int j = i;
            // Advance j as long as characters match the start of the current segment
            while (j < n && s[j] == s[i])
            {
                j++;
            }
            // Record the segment metadata
            type.push_back(s[i] - '0');
            start.push_back(i);
            end_idx.push_back(j - 1);
            i = j; // Move to the start of the next segment
        }

        int N = type.size(); // Total number of compressed segments

        // Lookup array to map any string index to its corresponding segment ID in O(1) time
        vector<int> pos_to_seg(n);
        for (int i = 0; i < N; i++)
        {
            for (int j = start[i]; j <= end_idx[i]; j++)
            {
                pos_to_seg[j] = i;
            }
        }

        // Precalculate the maximum potential gain for every segment
        vector<int> ans(N, 0);
        for (int i = 1; i < N - 1; i++)
        {
            // Only 1-segments surrounded by 0s yield a gain
            if (type[i] == 1)
            {
                // The gain is exactly the length of the left 0-segment + the length of the right 0-segment
                ans[i] = (end_idx[i - 1] - start[i - 1] + 1) + (end_idx[i + 1] - start[i + 1] + 1);
            }
        }

        // Precompute logarithms for the Sparse Table RMQ
        vector<int> log_table(N + 1, 0);
        for (int i = 2; i <= N; i++)
        {
            log_table[i] = log_table[i / 2] + 1;
        }

        int K = log_table[N] + 1; // Maximum power of 2 needed
        // Build the Sparse Table to answer range maximum queries in O(1)
        vector<vector<int>> st(K, vector<int>(N, 0));

        // Base case: intervals of length 2^0 = 1
        for (int i = 0; i < N; i++)
        {
            st[0][i] = ans[i];
        }

        // Dynamic programming to build larger intervals from smaller ones
        for (int j = 1; j < K; j++)
        {
            for (int i = 0; i + (1 << j) <= N; i++)
            {
                st[j][i] = max(st[j - 1][i], st[j - 1][i + (1 << (j - 1))]);
            }
        }

        // Helper lambda to query the maximum gain in a range of segments
        auto queryRMQ = [&](int L_idx, int R_idx)
        {
            if (L_idx > R_idx)
                return 0; // Invalid range protection
            int j = log_table[R_idx - L_idx + 1];
            // Combine two overlapping intervals of size 2^j to cover the whole query range
            return max(st[j][L_idx], st[j][R_idx - (1 << j) + 1]);
        };

        // Helper lambda to manually evaluate edge segments that might be partially chopped by query boundaries
        auto eval = [&](int i, int L, int R, int segL, int segR)
        {
            // The 1-segment must be strictly inside the query bounds to have 0s on both sides
            if (i <= segL || i >= segR)
                return 0;
            // Ignore if it's not a 1-segment
            if (type[i] == 0)
                return 0;

            int left_len = 0;
            // If the left 0-segment crosses the query's left bound, truncate it
            if (i - 1 == segL)
                left_len = max(0, end_idx[i - 1] - L + 1);
            else
                left_len = end_idx[i - 1] - start[i - 1] + 1; // Otherwise take full length

            int right_len = 0;
            // If the right 0-segment crosses the query's right bound, truncate it
            if (i + 1 == segR)
                right_len = max(0, R - start[i + 1] + 1);
            else
                right_len = end_idx[i + 1] - start[i + 1] + 1; // Otherwise take full length

            return left_len + right_len;
        };

        vector<int> res; // Array to hold answers for all queries

        for (const auto &q : queries)
        {
            int L = q[0]; // Query left bound
            int R = q[1]; // Query right bound

            // Find which compressed segments L and R fall into
            int segL = pos_to_seg[L];
            int segR = pos_to_seg[R];

            // If the query spans less than 3 segments, no 0-1-0 trade is possible
            if (segR - segL < 2)
            {
                res.push_back(total_ones);
                continue;
            }

            int max_gain = 0;
            // Check the 1-segments closest to the query boundaries (they might be partially truncated)
            max_gain = max(max_gain, eval(segL + 1, L, R, segL, segR));
            max_gain = max(max_gain, eval(segR - 1, L, R, segL, segR));

            // For all segments safely trapped in the middle, their 0s are fully intact. Fast query them!
            if (segL + 2 <= segR - 2)
            {
                max_gain = max(max_gain, queryRMQ(segL + 2, segR - 2));
            }

            // Final answer is the baseline 1s plus the maximum extra 1s we squeezed out
            res.push_back(total_ones + max_gain);
        }

        return res;
    }
};