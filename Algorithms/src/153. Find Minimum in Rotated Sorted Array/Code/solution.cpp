class Solution
{
public:
    int findMin(vector<int> &nums)
    {

        // Left pointer starts from beginning
        int left = 0;

        // Right pointer starts from end
        int right = nums.size() - 1;

        // Continue until both pointers meet
        while (left < right)
        {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // Minimum lies in right half
            if (nums[mid] > nums[right])
            {

                // Ignore left sorted part
                left = mid + 1;
            }
            else
            {

                // Minimum may be at mid or left side
                right = mid;
            }
        }

        // Both pointers point to minimum element
        return nums[left];
    }
};