#include <vector>
#include <cmath>
#include <algorithm>

class Solution
{
public:
    bool canDistribute(const std::vector<int> &quantities, int maxProducts, int n)
    {
        int storesNeeded = 0;
        for (int quantity : quantities)
        {
            storesNeeded += std::ceil((double)quantity / maxProducts);
            if (storesNeeded > n)
                return false;
        }
        return storesNeeded <= n;
    }

    int minimizedMaximum(int n, std::vector<int> &quantities)
    {
        int low = 1;
        int high = *std::max_element(quantities.begin(), quantities.end());
        int answer = high;

        while (low <= high)
        {
            int mid = low + (high - low) / 2;
            if (canDistribute(quantities, mid, n))
            {
                answer = mid;
                high = mid - 1;
            }
            else
            {
                low = mid + 1;
            }
        }

        return answer;
    }
};