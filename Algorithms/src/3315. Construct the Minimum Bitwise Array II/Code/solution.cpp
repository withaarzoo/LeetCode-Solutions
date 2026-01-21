class Solution
{
public:
    vector<int> minBitwiseArray(vector<int> &nums)
    {
        for (int &p : nums)
        {
            // Find the bit that can be safely removed
            int removable = ((p + 1) & ~p) >> 1;

            // If removable == 0, number is invalid (even prime like 2)
            // In that case flip all bits to force -1
            p ^= removable | (-(removable == 0) & ~p);
        }
        return nums;
    }
};
