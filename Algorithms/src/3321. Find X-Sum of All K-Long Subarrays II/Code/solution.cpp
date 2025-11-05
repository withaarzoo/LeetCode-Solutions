class Solution
{
public:
    using P = pair<int, int>; // (freq, value)
    struct Cmp
    {
        bool operator()(const P &a, const P &b) const
        {
            if (a.first != b.first)
                return a.first > b.first; // freq desc
            return a.second > b.second;   // value desc
        }
    };

    vector<long long> findXSum(vector<int> &nums, int k, int x)
    {
        int n = (int)nums.size();
        vector<long long> ans(n - k + 1);

        unordered_map<int, int> cnt;
        cnt.reserve(n * 2);

        set<P, Cmp> top, rest; // top = best x; rest = others
        long long topSum = 0;

        auto pull = [&](int v, int f)
        {
            P key{f, v};
            auto it = top.find(key);
            if (it != top.end())
            {
                topSum -= 1LL * v * f;
                top.erase(it);
            }
            else
            {
                auto jt = rest.find(key);
                if (jt != rest.end())
                    rest.erase(jt);
            }
        };
        auto pushToTop = [&](int v, int f)
        {
            top.insert({f, v});
            topSum += 1LL * v * f;
        };

        auto insertVal = [&](int v)
        {
            int f = cnt[v];
            if (f)
                pull(v, f);
            ++f;
            cnt[v] = f;
            pushToTop(v, f);
            if ((int)top.size() > x)
            {
                auto it = prev(top.end()); // smallest inside top
                topSum -= 1LL * it->first * it->second;
                rest.insert(*it);
                top.erase(it);
            }
        };

        auto eraseVal = [&](int v)
        {
            auto itc = cnt.find(v);
            if (itc == cnt.end())
                return;
            int f = itc->second;
            pull(v, f);
            if (--f == 0)
            {
                cnt.erase(itc);
            }
            else
            {
                cnt[v] = f;
                rest.insert({f, v}); // rank worsened
            }
            if ((int)top.size() < x && !rest.empty())
            {
                auto best = rest.begin(); // best among rest
                topSum += 1LL * best->first * best->second;
                top.insert(*best);
                rest.erase(best);
            }
        };

        for (int i = 0; i < k; ++i)
            insertVal(nums[i]);
        ans[0] = topSum;
        for (int i = k; i < n; ++i)
        {
            eraseVal(nums[i - k]);
            insertVal(nums[i]);
            ans[i - k + 1] = topSum;
        }
        return ans;
    }
};
