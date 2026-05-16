class Solution
{
public:
    int findMin(vector<int> &nums)
    {

        // Initialize binary search boundaries
        int left = 0;
        int right = nums.size() - 1;

        // Continue until both pointers meet
        while (left < right)
        {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // Case 1:
            // Minimum lies on left side including mid
            if (nums[mid] < nums[right])
            {
                right = mid;
            }

            // Case 2:
            // Minimum lies strictly on right side
            else if (nums[mid] > nums[right])
            {
                left = mid + 1;
            }

            // Case 3:
            // Duplicates found, cannot decide direction
            // Safely shrink search space
            else
            {
                right--;
            }
        }

        // Left now points to minimum element
        return nums[left];
    }
};