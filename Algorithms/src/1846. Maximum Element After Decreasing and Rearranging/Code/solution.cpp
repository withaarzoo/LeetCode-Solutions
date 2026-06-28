class Solution
{
public:
    int maximumElementAfterDecrementingAndRearranging(vector<int> &arr)
    {

        // Sort the array so we process numbers from smallest to largest.
        sort(arr.begin(), arr.end());

        // The first element must always become 1.
        arr[0] = 1;

        // Build the largest valid sequence.
        for (int i = 1; i < arr.size(); i++)
        {

            // The current value cannot be larger than previous + 1.
            // If it already satisfies this, it stays unchanged.
            arr[i] = min(arr[i], arr[i - 1] + 1);
        }

        // The last element is the maximum possible value.
        return arr.back();
    }
};