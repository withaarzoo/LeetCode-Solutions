class Solution
{
public:
    int repeatedNTimes(vector<int> &nums)
    {
        unordered_set<int> seen;

        for (int x : nums)
        {
            // If already seen, this is the repeated element
            if (seen.count(x))
            {
                return x;
            }
            // Otherwise, mark it as seen
            seen.insert(x);
        }
        return -1; // Will never reach here due to problem guarantee
    }
};
