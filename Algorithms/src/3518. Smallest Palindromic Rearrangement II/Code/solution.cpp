class Solution
{
public:
    string smallestPalindrome(string s, int k)
    {
        // Store frequencies of each character in the string
        vector<int> freq(26, 0);
        for (char c : s)
        {
            freq[c - 'a']++;
        }

        // Isolate frequencies meant only for the first half of the palindrome
        vector<int> half(26, 0);
        string mid = "";
        int m = 0;

        for (int i = 0; i < 26; ++i)
        {
            // The character with an odd frequency goes in the exact middle
            if (freq[i] % 2 != 0)
            {
                mid += (char)(i + 'a');
            }
            half[i] = freq[i] / 2;
            m += half[i];
        }

        // Helper lambda to calculate permutations of remaining characters
        auto get_ways = [&](const vector<int> &f, long long target_k)
        {
            long long ways = 1;
            int curr_len = 0;
            for (int count : f)
            {
                if (count > 0)
                {
                    curr_len += count;
                    long long n = curr_len;
                    long long r = count;

                    // Optimize nCr calculation by choosing the smaller r
                    if (r > n - r)
                        r = n - r;
                    long long cur_nCr = 1;

                    // Calculate nCr iteratively and cap it at target_k to prevent overflow
                    for (int i = 1; i <= r; ++i)
                    {
                        cur_nCr = cur_nCr * (n - i + 1) / i;
                        if (cur_nCr > target_k)
                        {
                            cur_nCr = target_k + 1;
                            break;
                        }
                    }
                    ways *= cur_nCr;
                    if (ways > target_k)
                        return target_k + 1;
                }
            }
            return ways;
        };

        // If the total valid permutations are less than k, return empty string
        if (get_ways(half, k) < k)
        {
            return "";
        }

        string first_half = "";
        // Build the first half character by character
        for (int i = 0; i < m; ++i)
        {
            for (int c = 0; c < 26; ++c)
            {
                if (half[c] > 0)
                {
                    // Try using the current character
                    half[c]--;
                    long long ways = get_ways(half, k);

                    // If permutations are enough, lock in this character
                    if (ways >= k)
                    {
                        first_half += (char)(c + 'a');
                        break;
                    }
                    else
                    {
                        // Otherwise, shrink k and restore character to try the next one
                        k -= ways;
                        half[c]++;
                    }
                }
            }
        }

        // Assemble final palindrome: first half + mid + reversed first half
        string res = first_half + mid;
        for (int i = m - 1; i >= 0; --i)
        {
            res += first_half[i];
        }
        return res;
    }
};