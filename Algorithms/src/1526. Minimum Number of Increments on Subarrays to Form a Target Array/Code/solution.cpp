#include <vector>
using namespace std;

class Solution
{
public:
    int minNumberOperations(vector<int> &target)
    {
        if (target.empty())
            return 0;
        int n = target.size();
        long long ans = target[0]; // start with operations needed for first element
        for (int i = 1; i < n; ++i)
        {
            if (target[i] > target[i - 1])
            {
                ans += (target[i] - target[i - 1]); // only positive increases add operations
            }
        }
        return (int)ans; // fits in 32-bit as per problem statement, but we used long long to be safe
    }
};
