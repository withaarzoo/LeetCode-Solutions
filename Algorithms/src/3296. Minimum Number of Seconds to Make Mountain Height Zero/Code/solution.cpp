class Solution
{
public:
    // Check if mountain can be reduced within given time
    bool can(long long time, int mountainHeight, vector<int> &workerTimes)
    {

        long long totalHeight = 0;

        for (int t : workerTimes)
        {

            long long left = 0, right = mountainHeight;

            // Binary search max height this worker can reduce
            while (left <= right)
            {
                long long mid = (left + right) / 2;

                long long required = (long long)t * (mid * (mid + 1) / 2);

                if (required <= time)
                {
                    left = mid + 1;
                }
                else
                {
                    right = mid - 1;
                }
            }

            totalHeight += right;

            if (totalHeight >= mountainHeight)
                return true;
        }

        return false;
    }

    long long minNumberOfSeconds(int mountainHeight, vector<int> &workerTimes)
    {

        long long left = 1, right = 1e18;
        long long ans = right;

        while (left <= right)
        {

            long long mid = (left + right) / 2;

            if (can(mid, mountainHeight, workerTimes))
            {
                ans = mid;
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }
        }

        return ans;
    }
};