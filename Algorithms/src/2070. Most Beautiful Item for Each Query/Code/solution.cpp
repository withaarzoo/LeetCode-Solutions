#include <vector>
#include <algorithm>
using namespace std;

class Solution
{
public:
    vector<int> maximumBeauty(vector<vector<int>> &items, vector<int> &queries)
    {
        // Sort items by price
        sort(items.begin(), items.end());

        // Prepare price-beauty pairs with running max beauty
        vector<pair<int, int>> priceBeauty;
        int maxBeauty = 0;
        for (auto &item : items)
        {
            maxBeauty = max(maxBeauty, item[1]);
            priceBeauty.push_back({item[0], maxBeauty});
        }

        // Process each query
        vector<int> result;
        for (int query : queries)
        {
            auto it = upper_bound(priceBeauty.begin(), priceBeauty.end(), make_pair(query, INT_MAX)) - 1;
            result.push_back(it >= priceBeauty.begin() ? it->second : 0);
        }
        return result;
    }
};
