class Solution
{
public:
    // Computes the best answer when rides A are taken first
    // and rides B are taken second.
    long long solve(vector<int> &startA, vector<int> &durA,
                    vector<int> &startB, vector<int> &durB)
    {

        int m = startB.size();

        // Store second-category rides as (start, duration)
        vector<pair<int, int>> rides(m);
        for (int i = 0; i < m; i++)
        {
            rides[i] = {startB[i], durB[i]};
        }

        // Sort by opening time
        sort(rides.begin(), rides.end());

        vector<int> starts(m);

        // prefixMinDur[i] = minimum duration in [0..i]
        vector<long long> prefixMinDur(m);

        // suffixMinFinish[i] = minimum (start + duration) in [i..m-1]
        vector<long long> suffixMinFinish(m);

        for (int i = 0; i < m; i++)
        {
            starts[i] = rides[i].first;

            if (i == 0)
                prefixMinDur[i] = rides[i].second;
            else
                prefixMinDur[i] = min(prefixMinDur[i - 1],
                                      (long long)rides[i].second);
        }

        for (int i = m - 1; i >= 0; i--)
        {
            long long finish = (long long)rides[i].first + rides[i].second;

            if (i == m - 1)
                suffixMinFinish[i] = finish;
            else
                suffixMinFinish[i] = min(suffixMinFinish[i + 1], finish);
        }

        long long ans = LLONG_MAX;

        for (int i = 0; i < (int)startA.size(); i++)
        {

            // Finish time after taking first ride
            long long finish1 = (long long)startA[i] + durA[i];

            // First index with start > finish1
            int pos = upper_bound(starts.begin(), starts.end(), finish1) - starts.begin();

            // Rides already open
            if (pos > 0)
            {
                ans = min(ans, finish1 + prefixMinDur[pos - 1]);
            }

            // Rides opening later
            if (pos < m)
            {
                ans = min(ans, suffixMinFinish[pos]);
            }
        }

        return ans;
    }

    int earliestFinishTime(vector<int> &landStartTime, vector<int> &landDuration, vector<int> &waterStartTime, vector<int> &waterDuration)
    {

        long long ans1 = solve(
            landStartTime, landDuration,
            waterStartTime, waterDuration);

        long long ans2 = solve(
            waterStartTime, waterDuration,
            landStartTime, landDuration);

        return (int)min(ans1, ans2);
    }
};