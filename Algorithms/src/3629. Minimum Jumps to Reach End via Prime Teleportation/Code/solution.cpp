class Solution
{
public:
    int minJumps(vector<int> &nums)
    {
        int n = nums.size();

        // If already at last index
        if (n == 1)
            return 0;

        // Find maximum value for sieve
        int mx = *max_element(nums.begin(), nums.end());

        // Smallest prime factor array
        vector<int> spf(mx + 1);

        // Initialize SPF
        for (int i = 0; i <= mx; i++)
        {
            spf[i] = i;
        }

        // Build sieve for smallest prime factor
        for (int i = 2; i * i <= mx; i++)
        {
            if (spf[i] == i)
            {
                for (int j = i * i; j <= mx; j += i)
                {
                    if (spf[j] == j)
                    {
                        spf[j] = i;
                    }
                }
            }
        }

        // Map prime factor -> indices divisible by it
        unordered_map<int, vector<int>> mp;

        // Build factor mapping
        for (int i = 0; i < n; i++)
        {
            int x = nums[i];

            unordered_set<int> used;

            // Get all unique prime factors
            while (x > 1)
            {
                int p = spf[x];

                if (!used.count(p))
                {
                    mp[p].push_back(i);
                    used.insert(p);
                }

                x /= p;
            }
        }

        // BFS queue
        queue<int> q;

        // Distance array
        vector<int> dist(n, -1);

        q.push(0);
        dist[0] = 0;

        while (!q.empty())
        {
            int i = q.front();
            q.pop();

            int steps = dist[i];

            // Reached end
            if (i == n - 1)
            {
                return steps;
            }

            // Move left
            if (i - 1 >= 0 && dist[i - 1] == -1)
            {
                dist[i - 1] = steps + 1;
                q.push(i - 1);
            }

            // Move right
            if (i + 1 < n && dist[i + 1] == -1)
            {
                dist[i + 1] = steps + 1;
                q.push(i + 1);
            }

            int val = nums[i];

            // Check if current number is prime
            if (val > 1 && spf[val] == val)
            {

                // Teleport to all divisible indices
                for (int nxt : mp[val])
                {
                    if (dist[nxt] == -1)
                    {
                        dist[nxt] = steps + 1;
                        q.push(nxt);
                    }
                }

                // Clear so we never process again
                mp[val].clear();
            }
        }

        return -1;
    }
};