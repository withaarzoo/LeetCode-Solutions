class Solution
{
public:
    int minimumPairRemoval(vector<int> &nums)
    {
        int n = nums.size();
        if (n <= 1)
            return 0;

        vector<long long> a(nums.begin(), nums.end());
        vector<int> prev(n), next(n);
        vector<bool> removed(n, false);

        for (int i = 0; i < n; i++)
        {
            prev[i] = i - 1;
            next[i] = (i + 1 < n) ? i + 1 : -1;
        }

        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq;
        for (int i = 0; i < n - 1; i++)
            pq.push({a[i] + a[i + 1], i});

        int bad = 0;
        for (int i = 0; i < n - 1; i++)
            if (a[i] > a[i + 1])
                bad++;

        int ops = 0;

        while (bad > 0)
        {
            auto [sum, i] = pq.top();
            pq.pop();

            if (removed[i] || next[i] == -1)
                continue;
            int j = next[i];
            if (removed[j] || a[i] + a[j] != sum)
                continue;

            int pi = prev[i];
            int nj = next[j];

            if (pi != -1 && a[pi] > a[i])
                bad--;
            if (a[i] > a[j])
                bad--;
            if (nj != -1 && a[j] > a[nj])
                bad--;

            a[i] = sum;
            removed[j] = true;
            next[i] = nj;
            if (nj != -1)
                prev[nj] = i;

            if (pi != -1 && a[pi] > a[i])
                bad++;
            if (nj != -1 && a[i] > a[nj])
                bad++;

            if (pi != -1)
                pq.push({a[pi] + a[i], pi});
            if (nj != -1)
                pq.push({a[i] + a[nj], i});

            ops++;
        }

        return ops;
    }
};
