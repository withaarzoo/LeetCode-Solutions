class Solution
{
public:
    long long maxPower(vector<int> &stations, int r, int k)
    {
        const int n = stations.size();

        // 1) Build base power using a difference array.
        vector<long long> diff(n + 1), base(n);
        for (int i = 0; i < n; ++i)
        {
            int L = max(0, i - r);
            int R = min(n, i + r + 1);
            diff[L] += stations[i];
            diff[R] -= stations[i];
        }
        long long run = 0;
        for (int i = 0; i < n; ++i)
        {
            run += diff[i];
            base[i] = run;
        }

        // 2) Binary search on the target minimum T.
        long long lo = 0, hi = accumulate(stations.begin(), stations.end(), 0LL) + k, ans = 0;
        auto can = [&](long long T) -> bool
        {
            vector<long long> add(n + 1, 0);
            long long extra = 0, used = 0;
            for (int i = 0; i < n; ++i)
            {
                extra += add[i];
                long long curr = base[i] + extra;
                if (curr < T)
                {
                    long long need = T - curr;
                    used += need;
                    if (used > k)
                        return false;
                    extra += need;
                    int end = min(n, i + 2 * r + 1);
                    add[end] -= need; // this added coverage stops after 'end-1'
                }
            }
            return true;
        };

        while (lo <= hi)
        {
            long long mid = (lo + hi) >> 1;
            if (can(mid))
            {
                ans = mid;
                lo = mid + 1;
            }
            else
            {
                hi = mid - 1;
            }
        }
        return ans;
    }
};
