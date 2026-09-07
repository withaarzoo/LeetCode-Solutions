class Solution
{
public:
    int distinctSubseqII(string s)
    {
        const long long MOD = 1000000007; // I use this modulo because the answer can be very large.

        long long dp = 1;              // I start with 1 because the empty subsequence is initially the only subsequence.
        vector<long long> last(26, 0); // last[c] stores dp from before the previous occurrence of character c.

        for (char c : s)
        {                        // I process every character once.
            int index = c - 'a'; // I convert the character into an index from 0 to 25.

            long long oldDp = dp; // I save the old count because last[index] must be updated with this value.

            dp = (2 * dp - last[index] + MOD) % MOD; // I double the subsequences and remove duplicates made by c.

            last[index] = oldDp; // I remember the old count for the next occurrence of this character.
        }

        return (dp - 1 + MOD) % MOD; // I remove the empty subsequence because the problem asks for non-empty ones.
    }
};