class Solution
{
public:
    int maxSumDivThree(vector<int> &nums)
    {
        long long sum = 0;
        // Use large numbers as "infinity"
        const int INF = 1e9;
        int r1_min1 = INF, r1_min2 = INF; // two smallest numbers with remainder 1
        int r2_min1 = INF, r2_min2 = INF; // two smallest numbers with remainder 2

        for (int x : nums)
        {
            sum += x;
            int r = x % 3;
            if (r == 1)
            {
                // Update two smallest remainder-1 numbers
                if (x < r1_min1)
                {
                    r1_min2 = r1_min1;
                    r1_min1 = x;
                }
                else if (x < r1_min2)
                {
                    r1_min2 = x;
                }
            }
            else if (r == 2)
            {
                // Update two smallest remainder-2 numbers
                if (x < r2_min1)
                {
                    r2_min2 = r2_min1;
                    r2_min1 = x;
                }
                else if (x < r2_min2)
                {
                    r2_min2 = x;
                }
            }
        }

        int mod = sum % 3;
        if (mod == 0)
            return (int)sum;

        long long removeCost = 1e18; // very large

        if (mod == 1)
        {
            // Option 1: remove one remainder-1 number
            if (r1_min1 != INF)
                removeCost = min(removeCost, (long long)r1_min1);
            // Option 2: remove two remainder-2 numbers
            if (r2_min2 != INF)
                removeCost = min(removeCost, (long long)r2_min1 + r2_min2);
        }
        else
        { // mod == 2
            // Option 1: remove one remainder-2 number
            if (r2_min1 != INF)
                removeCost = min(removeCost, (long long)r2_min1);
            // Option 2: remove two remainder-1 numbers
            if (r1_min2 != INF)
                removeCost = min(removeCost, (long long)r1_min1 + r1_min2);
        }

        if (removeCost >= 1e18)
            return 0; // no possible subset, return 0
        return (int)(sum - removeCost);
    }
};
