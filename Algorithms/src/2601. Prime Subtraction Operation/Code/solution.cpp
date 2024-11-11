#include <vector>
#include <bitset>
#include <algorithm>
#include <cmath>
using namespace std;

class Solution
{
public:
    vector<int> prime;

    void sieve(int M)
    {
        bitset<1001> isPrime;
        isPrime.set();
        isPrime[0] = isPrime[1] = 0;

        for (int p = 2; p * p <= M; ++p)
        {
            if (isPrime[p])
            {
                for (int j = p * p; j <= M; j += p)
                {
                    isPrime[j] = 0;
                }
            }
        }

        for (int p = 2; p <= M; ++p)
        {
            if (isPrime[p])
                prime.push_back(p);
        }
    }

    bool primeSubOperation(vector<int> &nums)
    {
        int n = nums.size(), M = *max_element(nums.begin(), nums.end());
        sieve(M);

        for (int i = n - 2; i >= 0; i--)
        {
            if (nums[i] >= nums[i + 1])
            {
                auto it = upper_bound(prime.begin(), prime.end(), nums[i] - nums[i + 1]);
                if (it == prime.end())
                    return false;
                nums[i] -= *it;
            }
            if (nums[i] <= 0)
                return false;
        }
        return true;
    }
};