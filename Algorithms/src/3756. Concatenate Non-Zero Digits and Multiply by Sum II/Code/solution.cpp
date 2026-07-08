class Solution
{
public:
    vector<int> sumAndMultiply(string s, vector<vector<int>> &queries)
    {
        // I use long long so multiplication does not overflow before taking modulo.
        const long long MOD = 1000000007LL;
        int n = s.size();

        // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
        vector<int> nonZeroCount(n + 1, 0);

        // I store only non-zero digits because zeroes are removed in every query.
        vector<int> digits;

        // I build the prefix count and compressed digit list in one pass.
        for (int i = 0; i < n; i++)
        {
            // I carry the previous count forward first.
            nonZeroCount[i + 1] = nonZeroCount[i];

            // Only non-zero digits belong to the compressed sequence.
            if (s[i] != '0')
            {
                nonZeroCount[i + 1]++;
                digits.push_back(s[i] - '0');
            }
        }

        int k = digits.size();

        // prefixValue[i] stores the first i compressed digits as a number modulo MOD.
        vector<long long> prefixValue(k + 1, 0);

        // prefixSum[i] stores the sum of the first i compressed digits.
        vector<long long> prefixSum(k + 1, 0);

        // power10[i] stores 10^i modulo MOD for removing an earlier prefix.
        vector<long long> power10(k + 1, 1);

        // I build all prefix information over the compressed digits.
        for (int i = 0; i < k; i++)
        {
            // Appending a digit means multiplying the old number by 10, then adding it.
            prefixValue[i + 1] = (prefixValue[i] * 10 + digits[i]) % MOD;

            // The digit sum is a normal prefix sum.
            prefixSum[i + 1] = prefixSum[i] + digits[i];

            // I precompute the next power of 10 for constant-time range extraction.
            power10[i + 1] = (power10[i] * 10) % MOD;
        }

        // I reserve the final size so the vector does not need repeated reallocations.
        vector<int> answer;
        answer.reserve(queries.size());

        // Each query is now answered in constant time.
        for (const auto &query : queries)
        {
            int l = query[0];
            int r = query[1];

            // left is the number of non-zero digits strictly before l.
            int left = nonZeroCount[l];

            // right is the number of non-zero digits up to and including r.
            int right = nonZeroCount[r + 1];

            // This is the number of digits in the compressed query range.
            int len = right - left;

            // I remove the earlier compressed prefix after shifting it by len places.
            long long x = (prefixValue[right] - (prefixValue[left] * power10[len]) % MOD + MOD) % MOD;

            // I get the digit sum using a normal prefix-sum subtraction.
            long long sum = prefixSum[right] - prefixSum[left];

            // I multiply the compressed number by its digit sum under modulo.
            answer.push_back((int)((x * sum) % MOD));
        }

        return answer;
    }
};