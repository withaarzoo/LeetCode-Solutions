class Solution
{
public:
    unordered_map<string, unordered_set<char>> rules;
    unordered_set<string> bad; // memo for failed rows

    bool dfs(const string &row, int idx, string &next)
    {
        // If pyramid completed
        if (row.size() == 1)
            return true;

        // If next row is fully built, move up
        if (idx == row.size() - 1)
        {
            if (bad.count(next))
                return false;
            bool ok = dfs(next, 0, *(new string()));
            if (!ok)
                bad.insert(next);
            return ok;
        }

        string key = row.substr(idx, 2);
        if (!rules.count(key))
            return false;

        for (char c : rules[key])
        {
            next.push_back(c);
            if (dfs(row, idx + 1, next))
                return true;
            next.pop_back(); // backtrack
        }
        return false;
    }

    bool pyramidTransition(string bottom, vector<string> &allowed)
    {
        for (auto &s : allowed)
            rules[s.substr(0, 2)].insert(s[2]);

        string next;
        return dfs(bottom, 0, next);
    }
};
