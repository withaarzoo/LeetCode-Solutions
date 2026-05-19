class Solution
{
public:
    int getCommon(vector<int> &nums1, vector<int> &nums2)
    {

        // Pointer for nums1
        int i = 0;

        // Pointer for nums2
        int j = 0;

        // Traverse both arrays together
        while (i < nums1.size() && j < nums2.size())
        {

            // If both values are equal,
            // this is the minimum common value
            if (nums1[i] == nums2[j])
            {
                return nums1[i];
            }

            // Move the pointer with smaller value
            // because it cannot match later
            if (nums1[i] < nums2[j])
            {
                i++;
            }
            else
            {
                j++;
            }
        }

        // No common value found
        return -1;
    }
};