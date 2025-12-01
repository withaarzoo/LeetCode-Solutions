class Solution
{
public:
    long long maxRunTime(int n, vector<int> &batteries)
    {
        // Calculate total sum of all battery minutes
        long long total = 0;
        for (int b : batteries)
            total += b;

        // Maximum possible time per computer cannot exceed total / n
        long long low = 0;
        long long high = total / n; // upper bound for answer

        // Binary search for maximum feasible time
        while (low < high)
        {
            long long mid = low + (high - low + 1) / 2; // upper mid

            // Check if we can run all n computers for 'mid' minutes
            long long usable = 0;
            for (int b : batteries)
            {
                // Each battery contributes at most 'mid' minutes
                usable += min<long long>(b, mid);
                // Small early cut: if already enough, no need to continue
                if (usable >= mid * n)
                    break;
            }

            if (usable >= mid * n)
            {
                // 'mid' minutes is possible, try longer
                low = mid;
            }
            else
            {
                // 'mid' minutes is not possible, try shorter
                high = mid - 1;
            }
        }

        return low;
    }
};
