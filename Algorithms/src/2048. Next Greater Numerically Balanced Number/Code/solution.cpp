class Solution
{
public:
    // Helper to test if x is numerically balanced
    bool isBalanced(int x)
    {
        int cnt[10] = {0};
        int t = x;
        while (t > 0)
        {
            cnt[t % 10]++;
            t /= 10;
        }
        if (cnt[0] > 0)
            return false; // digit 0 must not appear
        for (int d = 1; d <= 9; ++d)
        {
            if (cnt[d] != 0 && cnt[d] != d)
                return false;
        }
        return true;
    }

    int nextBeautifulNumber(int n)
    {
        int x = n + 1;
        while (true)
        {
            if (isBalanced(x))
                return x;
            ++x;
        }
        // unreachable
        return -1;
    }
};
