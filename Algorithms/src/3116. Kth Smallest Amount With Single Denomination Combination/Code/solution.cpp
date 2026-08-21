class Solution
{
public:
    long long findKthSmallest(vector<int> &coins, int k)
    {
        // Sort coins so smaller denominations are processed first.
        sort(coins.begin(), coins.end());

        // Store only useful coins because some denominations can be redundant.
        vector<long long> useful;

        for (int coin : coins)
        {
            bool redundant = false;

            // If an already kept coin divides this coin, all multiples of this
            // coin are already covered by that smaller coin.
            for (long long prev : useful)
            {
                if (coin % prev == 0)
                {
                    redundant = true;
                    break;
                }
            }

            // Keep the coin only if it adds new possible amounts.
            if (!redundant)
            {
                useful.push_back(coin);
            }
        }

        // The kth multiple of the smallest coin is always a valid upper bound.
        long long high = useful[0] * 1LL * k;
        long long low = 1;

        int m = useful.size();
        int totalMasks = 1 << m;

        // lcms[mask] stores the LCM of all coins selected by this mask.
        vector<long long> lcms(totalMasks, 1);

        // signs[mask] is +1 for odd subset sizes and -1 for even subset sizes.
        vector<int> signs(totalMasks, 1);

        // Precompute the LCM and inclusion-exclusion sign for every subset.
        for (int mask = 1; mask < totalMasks; ++mask)
        {
            long long currentLCM = 1;
            int bits = 0;

            for (int i = 0; i < m; ++i)
            {
                // Include useful[i] when its bit is set.
                if (mask & (1 << i))
                {
                    long long g = std::gcd(currentLCM, useful[i]);

                    // Divide before multiplying to reduce the risk of overflow.
                    currentLCM = currentLCM / g;

                    // If the LCM becomes larger than high, its contribution
                    // will always be zero, so I cap it at high + 1.
                    if (currentLCM > high / useful[i])
                    {
                        currentLCM = high + 1;
                        break;
                    }

                    currentLCM *= useful[i];
                    ++bits;
                }
            }

            // Save the subset LCM for future count() calls.
            lcms[mask] = currentLCM;

            // Odd subsets are added and even subsets are subtracted.
            signs[mask] = (bits % 2 == 1) ? 1 : -1;
        }

        // Count how many valid amounts are less than or equal to x.
        auto count = [&](long long x)
        {
            long long result = 0;

            for (int mask = 1; mask < totalMasks; ++mask)
            {
                // An LCM larger than x cannot divide any number up to x.
                if (lcms[mask] <= x)
                {
                    result += signs[mask] * (x / lcms[mask]);
                }
            }

            return result;
        };

        // Find the smallest value that contains at least k valid amounts.
        while (low < high)
        {
            long long mid = low + (high - low) / 2;

            // If mid already contains k or more valid amounts,
            // the answer can be mid or something smaller.
            if (count(mid) >= k)
            {
                high = mid;
            }
            else
            {
                // Otherwise, I need to search to the right.
                low = mid + 1;
            }
        }

        // low is the kth smallest valid amount.
        return low;
    }
};