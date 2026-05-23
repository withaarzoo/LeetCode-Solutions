class Solution
{
public:
    bool check(vector<int> &nums)
    {

        int n = nums.size();

        // Stores how many times the order decreases
        int count = 0;

        // Traverse every element
        for (int i = 0; i < n; i++)
        {

            // Compare current element with next element
            // % n is used so last element compares with first
            if (nums[i] > nums[(i + 1) % n])
            {
                count++;
            }

            // More than one decrease means invalid
            if (count > 1)
            {
                return false;
            }
        }

        // Valid sorted and rotated array
        return true;
    }
};