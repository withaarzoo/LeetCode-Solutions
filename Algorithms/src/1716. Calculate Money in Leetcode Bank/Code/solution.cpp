class Solution
{
public:
    int totalMoney(int n)
    {
        int w = n / 7; // number of full weeks
        int r = n % 7; // remaining days
        // sum of all full weeks: w*28 + 7 * (0 + 1 + ... + (w-1))
        int fullWeeksSum = w * 28 + 7 * (w * (w - 1) / 2);
        // sum of remaining days: r*(1 + w) + r*(r-1)/2
        int remSum = r * (1 + w) + (r * (r - 1) / 2);
        return fullWeeksSum + remSum;
    }
};
