class Solution
{
public:
    vector<string> validateCoupons(vector<string> &code,
                                   vector<string> &businessLine,
                                   vector<bool> &isActive)
    {

        // Priority order for business lines
        unordered_map<string, int> priority = {
            {"electronics", 0},
            {"grocery", 1},
            {"pharmacy", 2},
            {"restaurant", 3}};

        vector<pair<int, string>> validCoupons;

        for (int i = 0; i < code.size(); i++)
        {

            // Check active status
            if (!isActive[i])
                continue;

            // Check business line validity
            if (priority.find(businessLine[i]) == priority.end())
                continue;

            // Check code validity
            if (code[i].empty())
                continue;

            bool ok = true;
            for (char c : code[i])
            {
                if (!isalnum(c) && c != '_')
                {
                    ok = false;
                    break;
                }
            }
            if (!ok)
                continue;

            validCoupons.push_back({priority[businessLine[i]], code[i]});
        }

        // Sort by business priority, then by code
        sort(validCoupons.begin(), validCoupons.end(),
             [](auto &a, auto &b)
             {
                 if (a.first == b.first)
                     return a.second < b.second;
                 return a.first < b.first;
             });

        vector<string> result;
        for (auto &p : validCoupons)
        {
            result.push_back(p.second);
        }

        return result;
    }
};
