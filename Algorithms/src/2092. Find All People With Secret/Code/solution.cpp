class Solution
{
public:
    vector<int> parent;

    int find(int x)
    {
        if (parent[x] == x)
            return x;
        return parent[x] = find(parent[x]);
    }

    void unite(int x, int y)
    {
        x = find(x);
        y = find(y);
        if (x != y)
            parent[y] = x;
    }

    vector<int> findAllPeople(int n, vector<vector<int>> &meetings, int firstPerson)
    {
        sort(meetings.begin(), meetings.end(),
             [](auto &a, auto &b)
             { return a[2] < b[2]; });

        parent.resize(n);
        for (int i = 0; i < n; i++)
            parent[i] = i;

        vector<bool> knows(n, false);
        knows[0] = knows[firstPerson] = true;

        int i = 0;
        while (i < meetings.size())
        {
            int time = meetings[i][2];
            vector<int> people;

            int j = i;
            while (j < meetings.size() && meetings[j][2] == time)
            {
                unite(meetings[j][0], meetings[j][1]);
                people.push_back(meetings[j][0]);
                people.push_back(meetings[j][1]);
                j++;
            }

            unordered_map<int, bool> hasSecret;
            for (int p : people)
            {
                if (knows[p])
                    hasSecret[find(p)] = true;
            }

            for (int p : people)
            {
                if (hasSecret[find(p)])
                {
                    knows[p] = true;
                }
                else
                {
                    parent[p] = p; // reset
                }
            }

            i = j;
        }

        vector<int> ans;
        for (int i = 0; i < n; i++)
        {
            if (knows[i])
                ans.push_back(i);
        }
        return ans;
    }
};
