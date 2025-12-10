class Solution
{
public:
    int countPermutations(vector<int> &complexity)
    {
        const int MOD = 1'000'000'007;
        int n = (int)complexity.size();

        // 1. Find global minimum and how many times it appears
        int minVal = complexity[0];
        int cntMin = 0;
        for (int x : complexity)
        {
            if (x < minVal)
            {
                minVal = x;
                cntMin = 1;
            }
            else if (x == minVal)
            {
                cntMin++;
            }
        }

        // 2. Check that index 0 has the unique minimum
        if (complexity[0] != minVal || cntMin != 1)
        {
            return 0;
        }

        // 3. Compute (n - 1)! % MOD
        long long ans = 1;
        for (int i = 2; i <= n - 1; ++i)
        {
            ans = (ans * i) % MOD;
        }
        return (int)ans;
    }
};
