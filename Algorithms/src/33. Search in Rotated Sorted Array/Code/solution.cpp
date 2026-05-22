class Solution
{
public:
    int search(vector<int> &nums, int target)
    {

        // Start pointer
        int left = 0;

        // End pointer
        int right = nums.size() - 1;

        // Continue until search space becomes empty
        while (left <= right)
        {

            // Find middle index safely
            int mid = left + (right - left) / 2;

            // If target is found, return index
            if (nums[mid] == target)
            {
                return mid;
            }

            // Check if left half is sorted
            if (nums[left] <= nums[mid])
            {

                // Check whether target lies inside left sorted half
                if (nums[left] <= target && target < nums[mid])
                {

                    // Move to left half
                    right = mid - 1;
                }
                else
                {

                    // Move to right half
                    left = mid + 1;
                }
            }
            // Otherwise right half must be sorted
            else
            {

                // Check whether target lies inside right sorted half
                if (nums[mid] < target && target <= nums[right])
                {

                    // Move to right half
                    left = mid + 1;
                }
                else
                {

                    // Move to left half
                    right = mid - 1;
                }
            }
        }

        // Target not found
        return -1;
    }
};